package tfsec

import (
	"testing"

	"github.com/tfsec/tfsec/internal/app/tfsec/checks/aws"
	"github.com/tfsec/tfsec/internal/app/tfsec/scanner"
)

func TestAWSUnencryptedElasticsearchDomain(t *testing.T) {

	var tests = []struct {
		name                  string
		source                string
		mustIncludeResultCode scanner.RuleID
		mustExcludeResultCode scanner.RuleID
	}{
		{
			name: "check no encrypt_at_rest block aws_elasticsearch_domain",
			source: `
resource "aws_elasticsearch_domain" "my_elasticsearch_domain" {
	
}`,
			mustIncludeResultCode: aws.AWSUnencryptedElasticsearchDomain,
		},
		{
			name: "check false enabled attr aws_elasticsearch_domain",
			source: `
resource "aws_elasticsearch_domain" "my_elasticsearch_domain" {
  domain_name = "domain-foo"

  encrypt_at_rest { }
}`,
			mustIncludeResultCode: aws.AWSUnencryptedElasticsearchDomain,
		},
		{
			name: "check true enabled attr aws_elasticsearch_domain",
			source: `
resource "aws_elasticsearch_domain" "my_elasticsearch_domain" {
  domain_name = "domain-foo"

  encrypt_at_rest {
    enabled = true
  }
}`,
			mustExcludeResultCode: aws.AWSUnencryptedElasticsearchDomain,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results := scanSource(test.source)
			assertCheckCode(t, test.mustIncludeResultCode, test.mustExcludeResultCode, results)
		})
	}

}

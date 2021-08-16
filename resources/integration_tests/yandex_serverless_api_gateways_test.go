package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/yandex-cloud/cq-provider-yandex/resources"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
)

func TestIntegrationServerlessApiGateways(t *testing.T) {
	var tfTmpl = fmt.Sprintf(`
resource "yandex_api_gateway" "cq-apigateway-test-%[1]s" {
  name = "cq-apigateway-test-%[1]s"
  description = "any description"
  labels = {
    label       = "label"
    empty-label = ""
  }
  spec = <<-EOT
openapi: "3.0.0"
info:
  version: 1.0.0
  title: Test API
paths:
  /hello:
    get:
      summary: Say hello
      operationId: hello
      parameters:
        - name: user
          in: query
          description: User name to appear in greetings
          required: false
          schema:
            type: string
            default: 'world'
      responses:
        '200':
          description: Greeting
          content:
            'text/plain':
              schema:
                type: "string"
      x-yc-apigateway-integration:
        type: dummy
        http_code: 200
        http_headers:
          'Content-Type': "text/plain"
        content:
          'text/plain': "Hello again, {user}!\n"
EOT
}
`, suffix)
	testIntegrationHelper(t, resources.ServerlessApiGateways(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "yandex_serverless_api_gateways",
			Filter: func(sq squirrel.SelectBuilder, _ *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": fmt.Sprintf("cq-apigateway-test-%[1]s", suffix),
					},
				},
			},
		}
	}, tfTmpl)
}

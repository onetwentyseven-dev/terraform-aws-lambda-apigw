<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | ~> 5.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_apigatewayv2_integration.lambda](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/apigatewayv2_integration) | resource |
| [aws_apigatewayv2_route.lambda](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/apigatewayv2_route) | resource |
| [aws_lambda_permission.allow_apigateway](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_permission) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_api_execution_arn"></a> [api\_execution\_arn](#input\_api\_execution\_arn) | n/a | `string` | n/a | yes |
| <a name="input_apigw_id"></a> [apigw\_id](#input\_apigw\_id) | n/a | `string` | n/a | yes |
| <a name="input_function_invoke_arn"></a> [function\_invoke\_arn](#input\_function\_invoke\_arn) | n/a | `string` | n/a | yes |
| <a name="input_function_name"></a> [function\_name](#input\_function\_name) | n/a | `string` | n/a | yes |
| <a name="input_routes"></a> [routes](#input\_routes) | n/a | `list(string)` | n/a | yes |

## Outputs

No outputs.
<!-- END_TF_DOCS -->
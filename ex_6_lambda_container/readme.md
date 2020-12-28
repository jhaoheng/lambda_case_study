# Environment variables
- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY
- AWS_SESSION_TOKEN
- AWS_REGION
- AWS_LAMBDA_FUNCTION_TIMEOUT
- AWS_LAMBDA_FUNCTION_VERSION
- AWS_LAMBDA_FUNCTION_NAME
- AWS_LAMBDA_FUNCTION_MEMORY_SIZE

# localhost
1. `GOARCH=amd64 GOOS=linux go build -o ./app/app ./app/main.go`
2. `docker run --rm -v $(pwd)/app:/var/task --name app -p 9000:8080 public.ecr.aws/lambda/go:1 app`
3. `curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"name":"max"}'`

# deploy

> https://docs.aws.amazon.com/lambda/latest/dg/images-create.html#images-create-2

1. login ECR
2. push image to ECR

# deploy with sam
- https://aws.amazon.com/tw/blogs/compute/using-container-image-support-for-aws-lambda-with-aws-sam/
- https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-function-imageconfig.html

- 增加的參數
	- https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-resource-function.html
	- PackageType
	- ImageConfig
		- Command
		- EntryPoint
		- WorkingDirectory
	- ImageUri
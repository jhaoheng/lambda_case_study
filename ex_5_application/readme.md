# doc
- SAM : https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md
- golang event 
    - doc : https://godoc.org/github.com/aws/aws-lambda-go/events
    - example : https://github.com/aws/aws-lambda-go/tree/master/events


# start
1. `make build`
2. 建立 : `sam local start-api --region us-east-1`
3. 測試 : 
    - `curl -X GET http://localhost:3000/hello/get`
    - `curl -X DELETE http://localhost:3000/hello/del`
    - `curl -X POST http://localhost:3000/hello/post -d '{"name":"hello this is your name"}'`
4. `make clean`


# test
## invoke api
- `sam local invoke "apiget"`

## invoke cus-event
- `echo '{"message": "Hey, are you there?" }' | sam local invoke "cus_event" --event -`

## invoke s3 event
- `sam local generate-event s3 put --bucket my-bucket-20200508 --key hello | sam local invoke "s3Event" --event -`


# deploy
1. 設定 bucket polict : https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-template-publishing-applications.html
2. 建立 packaged.yaml : `sam package --template-file template.yaml --output-template-file packaged.yml --s3-bucket {bucket}`
3. 發布版本 : `sam publish --template packaged.yml --region us-east-1`
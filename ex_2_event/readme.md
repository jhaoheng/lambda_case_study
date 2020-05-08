# 目的
- 建立客製化的 event, 傳送給 Lambda
- 在本地端測試

## build

`GOOS=linux go build -o ./lambda/task/main ./app/main.go`

## run localhost test

`docker run --rm -v $(pwd)/lambda/task:/var/task lambci/lambda:go1.x main '{"IsTest":true, "name":"max"}'`

## run localhost with docker-compose

1. `docker-compose -f lambda/docker-compose.yml up -d`
2. 查看 lambda log : `docker logs -f lambda`
3. 調用 lambda, 使用 curl
    - `curl -d '{"IsTest":true}' http://localhost:9001/2015-03-31/functions/myfunction/invocations`
4. 調用 lambda, 使用 awscli : 
    - `docker exec -it awscli /bin/bash`
    - `aws --endpoint-url http://lambda:9001 --region us-east-1 lambda invoke --function-name main --no-sign-request --payload '{"IsTest":true}' /dev/stdout`
    - `exit`
5. 結束 `docker-compose -f ./lambda/docker-compose.yml down`
# 目的
- 在本地端建立 hello Lambda, 並且在本地端測試

## build binary
- `GOOS=linux go build -o ./lambda/task/main ./app/main.go`

## 本地測試 : `docker run` 
- `docker run --rm -v $(pwd)/lambda/task:/var/task lambci/lambda:go1.x main`

## 本地測試 : `docker-compose`
1. `docker-compose -f lambda/docker-compose.yml up -d`
2. 查看 lambda log : `docker logs -f lambda`
3. 調用 lambda, 使用 curl
    - `curl -d '{}' http://localhost:9001/2015-03-31/functions/myfunction/invocations`
4. 調用 lambda, 使用 awscli : 
    - `docker exec -it awscli /bin/bash`
    - `aws --endpoint-url http://lambda:9001 --region us-east-1 lambda invoke --function-name main --no-sign-request --payload '{}' /dev/stdout`
5. 結束 `docker-compose -f ./lambda/docker-compose.yml down`
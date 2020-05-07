1. `cp ../ex_2/lambda/task/main .`
2. `zip main.zip main`
3. 上傳, 透過 console
4. 測試
5. awscli invoke
    1. 確定 `./.aws/credential`
    2. `docker-compose up -d`
    3. `docker exec -it awscli /bin/bash`
    4. `aws lambda list-functions`
    5. `aws lambda invoke --function-name myFunc --cli-binary-format raw-in-base64-out --payload '{ "isTest":true }' /dev/stdout`
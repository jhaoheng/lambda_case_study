# 目的
- 部署 lambda func 到 Cloud, AWS Lambda 

# binary
- main : 為 ex_2 的 binary
- main.zip

# deploy to cloud, AWS Lambda 
## 設定 aws credential
```
[default]
aws_access_key_id=
aws_secret_access_key=
```

## deploy
1. 上傳, 透過 console
2. 從 console 測試
3. awscli invoke
    1. 確定 `./.aws/credential`
    2. `docker-compose up -d`
    3. `docker exec -it awscli /bin/bash`
        - `aws lambda list-functions`
        - `aws lambda invoke --function-name myFunc --cli-binary-format raw-in-base64-out --payload '{ "isTest":true }' /dev/stdout`
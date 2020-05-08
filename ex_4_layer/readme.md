## build 
> 注意 build func 與 layer 必須要在同一個環境

1. `docker run --rm -it -e GO111MODULE=on -v $(pwd)/app:/go/src/app -v $(pwd)/lambda:/lambda -w /go/src golang:1.13.1 /bin/bash`
2. `go build -o /lambda/task/main ./app/main.go`
3. `go build -buildmode=plugin -o /lambda/opt/layer.so ./app/layer/helloLayer.go`

# run local test

```
docker run --rm -v $(pwd)/lambda/task:/var/task:ro,delegated -v $(pwd)/lambda/opt:/opt:ro,delegated lambci/lambda:go1.x main
```
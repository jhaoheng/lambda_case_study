# build 
## 注意
1. 因為 osx golang build plugin 跨平台到 linux 的 c complier 有 issue, 故在 docker complier
2. 注意 func 與 layer 必須要在同一個環境中一起編譯 

## compiler
1. `docker run --rm -it -e GO111MODULE=on -v $(pwd)/app:/go/src/app -v $(pwd)/lambda:/lambda -w /go/src golang:1.13.1 /bin/bash`
2. `go build -o /lambda/task/main ./app/main.go`
3. `go build -buildmode=plugin -o /lambda/opt/layer.so ./app/layer/helloLayer.go`
4. exit

# run local test

```
docker run --rm -v $(pwd)/lambda/task:/var/task:ro,delegated -v $(pwd)/lambda/opt:/opt:ro,delegated lambci/lambda:go1.x main
```
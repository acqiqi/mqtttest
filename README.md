# VhakeIOTHub

`VhakeIotHub 是一款专用于物联网交互的中间件。支持多种协议接入，以及多种硬件协议解析`

#### 1、安装方法：
```$xslt
1.export GO111MODULE=on 
2.export GOPROXY=https://goproxy.io 
3.go mod vendor
4.go run main.go
```

git push -u origin master


screen /dev/tty.wchusbserial14120 115200

GOOS=linux GOARCH=mipsle GOMIPS=softfloat  go build main.go
GOOS=linux GOARCH=amd64 go build main.go
GOOS=windows GOARCH=amd64 go build main.go
go build -ldflags "-w" ./main.go 

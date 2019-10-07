## google proto

### 安装 protoc
下载 
[https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases)

    protoc-3.10.0-linux-x86_64.zip 解压安装到/usr/local
### goprotobuf
    go get github.com/golang/protobuf/protoc-gen-go
    
### 生成go文件
    protoc --go_out=. *.proto
    
    
> 详细 [https://colobu.com/2019/10/03/protobuf-ultimate-tutorial-in-go/](https://colobu.com/2019/10/03/protobuf-ultimate-tutorial-in-go/)
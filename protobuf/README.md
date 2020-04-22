## 安装
1. 安装protoc  
下载地址：https://github.com/protocolbuffers/protobuf/releases
2. 安装go插件
```
go get google.golang.org/protobuf/cmd/protoc-gen-go
```
protoc-gen-go将安装到$GOPATH/bin目录。
## 编译器调用
使用--go_out标志调用编译器，编译器将生成go输出。--go_out标志的参数是您希望编译器输出go源文件的目录。编译器为每一个.proto文件创建一个对应的go源文件，且文件名是将.proto替换为.pb.go。.proto文件中包含一个go_package选项，用于指定生成的go源文件的完整包导入路径，如：
```
option go_package = "example.com/foo/bar";
```
输出文件所在的输出目录的子目录取决于go_package选项和编译器标志：  
默认情况下，输出文件放置在以Go软件包的导入路径命名的目录中。例如，protos/foo.proto 具有上述go_package选项的文件将生成名为的文件 example.com/foo/bar/foo.pb.go。  
如果给--go_opt=paths=source_relative标记protoc，则将输出文件放置在与输入文件相同的相对目录中。例如，该文件protos/foo.proto 生成名为的文件protos/foo.pb.go。    
当您像这样运行编译器时
```
# --proto_path 指定.proto文件路径
protoc --proto_path=src --go_out=build/gen --go_opt=paths=source_relative src/foo.proto src/bar/baz.proto
```
编译器将读取文件src/foo.proto和 src/bar/baz.proto。它产生两个输出文件： build/gen/foo.pb.go和build/gen/bar/baz.pb.go。  
build/gen/bar 必要时，编译器会自动创建目录，但不会创建build或 build/gen; 目录。它们必须已经存在。  

嵌套类型，编译器将生成两个结构体，如下：将生成Foo与Foo_Bar两个struct
```
message Foo {
  message Bar {
  }
}
```
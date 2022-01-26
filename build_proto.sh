for proto in user order
do
  echo "Removing grpc/${proto}/*.go"
  rm -f grpc/${proto}/*.go
  echo "Generating grpc/${proto}/${proto}.pb.go"
#  protoc -I . --go_out=plugins=grpc:. idl/${proto}/${proto}.proto #以前的旧版生成方法
  protoc -I . --go_out=. --go-grpc_out=. idl/${proto}/${proto}.proto
  echo "Generating ${proto}/${proto}.gw.go"
  protoc -I . --grpc-gateway_out=. idl/${proto}/${proto}.proto
done

#$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
#$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

#两个易混参数
#proto文件（非protoc）有两个易混参数，即package和xx_package，xx指的是你的编译语言，比如你要编程成Go语言，对应的就是go_package。
#package
#package参数针对的是protobuf，是proto文件的命名空间，它的作用是为了避免我们定义的接口，或者message出现冲突。

#--go_out详细解读
#想必大家在使用的时候，应该遇到过这些写法：--go_out=paths=import:.、--go_out=paths=source_relative:.，或者--go_out=plugins=grpc:.。
#这样写表达的是啥意思呢？
#所以我们需要知道，--go_out参数是用来指定 protoc-gen-go 插件的工作方式和Go代码的生成位置，而上面的写法正是表明该插件的工作方式。
#--go_out主要的两个参数为plugins 和 paths，分别表示生成Go代码所使用的插件，以及生成的Go代码的位置。--go_out的写法是，参数之间用逗号隔开，最后加上冒号来指定代码的生成位置，比如--go_out=plugins=grpc,paths=import:.。
#paths参数有两个选项，分别是 import 和 source_relative，默认为 import，表示按照生成的Go代码的包的全路径去创建目录层级，source_relative 表示按照 proto源文件的目录层级去创建Go代码的目录层级，如果目录已存在则不用创建。
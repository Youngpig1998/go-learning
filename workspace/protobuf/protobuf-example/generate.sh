#! /bin/bash


# generate cpp or go source code 
protoc -I=. --cpp_out=. addressbook.proto

#编译writer.cpp。第一个writer是编译输出的执行程序名，第二个writer.cpp是源代码
g++ -o writer writer.cpp addressbook.pb.cc -lprotobuf

#编译reader.cpp
g++ -o reader reader.cpp addressbook.pb.cc -lprotobuf



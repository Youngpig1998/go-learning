// Description: 模拟服务端Rpc处理。

package main

import "fmt"

type any interface {}

//function函数集合
// Op 定义的对msg处理的函数集合.
type Op func(msg any) (any, error)

//decode 解码远程消息
func decode(msg any) Op {
	return func(any) (any, error) {
		//decode msg
		fmt.Println("decoding ... ",msg)
		decodeRes := fmt.Sprintf("decode_%v",msg)
		fmt.Println("decode to parameter->", decodeRes)
		return decodeRes, nil
	}
}

// opAction 模拟服务提供方处理业务逻辑
func opAction(parameter any) Op {
	return func(any) (any, error) {
		//decode msg
		fmt.Println("do opAction ...", parameter)
		opRes := fmt.Sprintf("opAction_%v",parameter)
		fmt.Println("after opAction result ->",opRes)

		return opRes,nil
	}
}

// encode 模拟服务提供方处将处理结果打包的过程
func encode(result any) Op {
	return func(any) (any, error) {
		//encode msg
		fmt.Println("encoding ...", result)
		encodedRes := fmt.Sprintf("encode_%v",result)
		fmt.Println("after encoded result ->",encodedRes)

		return encodedRes,nil
	}
}


func TestOp() {
	msg := "[caller_once]"

	//解码
	decodeOp := decode(msg)
	decodeMsg, err := decodeOp(msg)
	if err != nil {
		return
	}

	//执行业务逻辑
	actionOp := opAction(decodeMsg)
	opRes, err := actionOp(decodeMsg)
	if err != nil {
		return
	}

	// 编码
	encodeOp := encode(opRes)
	_, err = encodeOp(opRes)
	if err != nil {
		return
	}
}

func main() {
	TestOp()
}
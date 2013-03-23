package main

import "fmt"

func main(){
	//int型の変数を宣言
	var i int = 1234
	
	//int型からuint32型への変換
	var u uint32 = uint32(i)
	
	//uint32型からfloat32型への変換
	var f float32 = float32(u)
	
	//int型からstring型への変換
	var s string = string(i)
	
	//string型からスライス型への変換
	var b []byte = []byte("abc")
	
	//出力
	fmt.Println(u)
	fmt.Println(f)
	fmt.Println(s)
	fmt.Println(b)
}

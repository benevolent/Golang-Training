package main

import "fmt"

func main(){
	//string型の減数を宣言
	var en string = "golang"
	var ja string = "Go言語"
	
	//文字列の他がさを出力
	fmt.Println(en, " len:",len(en))
	fmt.Println(ja, " len:",len(ja))
}

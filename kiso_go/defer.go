package main

import "fmt"

func main(){
	fmt.Println("start")
	
	//defer関数を遅延実行
	defer delay()
	
	fmt.Println("end")
}

//遅延指定される関数
//deferは呼び出し時に指定するので、
//この関数自体は遅延実行されることを意識していない
func delay(){
	fmt.Println("delayが呼び出されました")
}

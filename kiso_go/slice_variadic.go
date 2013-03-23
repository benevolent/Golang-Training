package main

import "fmt"

func main(){
	s := []string{"a","b","c"}
	
	//スライスをそのまま渡すには...をつける
	test(s...)
	
	//こうしたときと渡される値は一緒
	test("a","b","c")
}

func test(s ...string){
	//スライスの長和と値を出力
	fmt.Println(len(s),s)
}

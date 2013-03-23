package main

import "fmt"

func main(){
	values := [...]int{0,1,2,3,4}
	
	s1 := values[1:4]
	
	//出力
	fmt.Println(s1)
	fmt.Println("len:",len(s1))
	fmt.Println("cap:",cap(s1))
	
	//スライスをスライス
	s2 := s1[1:4]	//キャパシティまでスライス可能
	
	//出力
	fmt.Println(s2)
	fmt.Println("len:",len(s2))
	fmt.Println("cap:",cap(s2))
}

package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(in chan int,out chan string){
	//無限ループ
	for{
		var n int = <-in	//inから受信してnに代入
		
		switch {
		case n%15 == 0:	out <- "FIZZBUZZ"
		case n&3 == 0: out <- "FIZZBUZZ"
		case n%5 == 0: out <- "FIZZBUZZ"
		default: out <- strconv.Itoa(n)
		}
	}
}

func main(){
	ch := make(chan int)
	out := make(chan string)
	go fizzbuzz(ch,out)
	for i := 1; i < 100; i++ {
		ch <- i
		s := <- out
		fmt.Println(s)
	}
}

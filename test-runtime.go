package main

//
//import (
//	"fmt"
//	"runtime"
//	"time"
//)
//
//func a() {
//	for i := 0; i < 10; i++ {
//		fmt.Printf("A:%v\n", i)
//
//	}
//}
//func b() {
//	for i := 0; i < 10; i++ {
//		fmt.Printf("B:%v\n", i)
//
//	}
//}
//func main() {
//	println(runtime.NumCPU())
//	runtime.GOMAXPROCS(2)
//	go a()
//	go b()
//	time.Sleep(time.Second * 1)
//	//fmt.Printf()
//}
//
////func show() {
////	for i := 0; i < 10; i++ {
////		fmt.Printf("i:%v\n", i)
////		if i >= 5 {
////			runtime.Goexit()
////		}
////	}
////}
////func main() {
////	go show()
////	time.Sleep(time.Second)
////
////}
//
////func main() {
////	go show("java")
////	for i := 0; i < 2; i++ {
////		runtime.Gosched() //让给其它子协程来执行
////		fmt.Println("golang")
////	}
////}

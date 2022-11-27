package main

//
//import (
//	"fmt"
//	"time"
//)
//
//var chanInt = make(chan int)
//var chanString = make(chan string)
//
//func main() {
//
//	go func() {
//		chanInt <- 100
//		chanString <- "hello"
//		defer close(chanInt)
//		defer close(chanString)
//	}()
//	for {
//		select {
//		case r := <-chanInt:
//			fmt.Printf("r:%v\n", r)
//		case s := <-chanString:
//			fmt.Printf("s:%v\n", s)
//		default:
//			fmt.Printf("default...\n")
//		}
//		time.Sleep(time.Second)
//	}
//}

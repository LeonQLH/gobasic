package main

//
//import (
//	"fmt"
//	"sync"
//)
//
//var wg sync.WaitGroup
//
//func showMsg(i int) {
//	defer wg.Done()
//	fmt.Printf("i:%v\n", i)
//
//}
//func main() {
//	for i := 0; i < 10; i++ {
//		go showMsg(i)
//		wg.Add(1)
//	}
//	wg.Wait()
//	fmt.Printf("end...")
//}

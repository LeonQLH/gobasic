package main

//
//import "fmt"
//
//type Person struct {
//	name string
//}
//
//type Customer struct {
//	name string
//}
//
//// (per Person) 接收者 receiver
//func (per Person) eat() {
//	fmt.Printf("%v,eat...", per.name)
//}
//func (per Person) sleep() {
//	fmt.Printf("%v,sleep...", per.name)
//}
//func (customer Customer) login(name string, pwd string) bool {
//	fmt.Printf("customer.name:%v\n", customer.name)
//	if name == "tom" && pwd == "123" {
//		return true
//	} else {
//		return false
//	}
//}
//func main() {
//	//per := Person{
//	//	name: "tom",
//	//}
//	//per.eat()
//	//per.sleep()
//	customer := Customer{
//		name: "tome",
//	}
//	b := customer.login("tom", "123")
//	fmt.Printf("b:%v\n", b)
//
//}

package main

//
//import "fmt"
//
//func test() {
//	type Person struct {
//		name string
//	}
//	p1 := Person{
//		name: "tom",
//	}
//	p2 := &Person{
//		name: "tom",
//	}
//	fmt.Printf("p1:%T\n", p1)
//	fmt.Printf("p2:%T\n", p2)
//}
//
//type Person struct {
//	name string
//}
//
//func showPerson1(per Person) {
//	per.name = "...\n"
//	fmt.Printf("name:%v\n", per.name)
//
//}
//func showPerson2(per *Person) {
//	per.name = "tom..."
//}
//func main() {
//	p1 := Person{
//		name: "tom",
//	}
//	p2 := &Person{
//		name: "tom",
//	}
//	showPerson1(p1)
//	fmt.Printf("p1:%v\n", p1)
//	fmt.Printf(".....\n")
//	showPerson2(p2)
//	fmt.Printf("p2:%v\n", *p2)
//}

package main

import "fmt"

/*
*
使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
*/
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	EmployeeID int
	Person
}

func (e Employee) PrintInfo() {
	fmt.Printf("员工ID: %d, 姓名: %s, 年龄: %d\n", e.EmployeeID, e.Name, e.Age)
}

func main() {
	p := Employee{
		EmployeeID: 1,
		Person: Person{
			Name: "小芳",
			Age:  23,
		},
	}
	p.PrintInfo()
}

package task3

import (
	"errors"
	"math/rand"

	"github.com/jmoiron/sqlx"
)

// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
type Employee struct {
	Id         uint    `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func CreateEmployees(db *sqlx.DB) error {
	createSql := `
	CREATE TABLE IF NOT EXISTS employees (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		department VARCHAR(100) NOT NULL,
		salary DECIMAL(10,2) NOT NULL
	);`
	_, err := db.Exec(createSql)
	return err
}

func InitEmployeeData(db *sqlx.DB) error {
	employee := []Employee{
		{Name: "张三", Department: "技术部", Salary: 15000},
		{Name: "李四", Department: "技术部", Salary: 30000},
		{Name: "王二", Department: "技术部", Salary: 20000},
		{Name: "麻子", Department: "人事部", Salary: 10000},
	}
	insert := `INSERT INTO employees(name, department, salary) VALUES (:name, :department, :salary)`
	_, err := db.NamedExec(insert, employee)
	return err
}

// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
func QueryTechnicalDepartment(db *sqlx.DB, department string) ([]Employee, error) {
	if department == "" {
		return nil, errors.New("参数不能为空")
	}
	var employees []Employee
	query := `SELECT * FROM employees where department = ?`
	err := db.Select(&employees, query, department)
	return employees, err
}

// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
func QueryMaxSalary(db *sqlx.DB) (Employee, error) {
	query := `SELECT * FROM employees order by salary desc limit 1`
	var employees Employee
	err := db.Get(&employees, query)
	return employees, err
}

// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
//
//	要求 ：
//	定义一个 Book 结构体，包含与 books 表对应的字段。
//	编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
type Book struct {
	Id        uint    `db:"id"`
	Title     string  `db:"title"`
	Author    string  `db:"author"`
	Price     float64 `db:"price"`
	BookCount int     `db:"bookCount"`
}

func CreateBook(db *sqlx.DB) error {
	createSql := `
	CREATE TABLE IF NOT EXISTS books (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
	    author VARCHAR(255) NOT NULL,
	    price DECIMAL(10,2) NOT NULL
	);`
	_, err := db.Exec(createSql)
	return err
}

func InitBookData(db *sqlx.DB) error {
	book := []Book{
		{Title: "Go编程基础", Author: "Alice", Price: rand.Float64()*30 + 20},
		{Title: "深入浅出Go", Author: "Alice", Price: rand.Float64()*30 + 20},
		{Title: "Go中的数据结构", Author: "Bob", Price: rand.Float64()*30 + 20},
		{Title: "SQL实战指南", Author: "Bob", Price: rand.Float64()*30 + 20},
		{Title: "Go语言初学者", Author: "Charlie", Price: rand.Float64()*30 + 20},
		{Title: "GORM深度剖析", Author: "Charlie", Price: rand.Float64()*30 + 20},
		{Title: "Go设计模式", Author: "David", Price: rand.Float64()*30 + 20},
		{Title: "Go算法解析", Author: "David", Price: rand.Float64()*30 + 20},
		{Title: "Go并发编程", Author: "Eve", Price: rand.Float64()*30 + 20},
		{Title: "Go语言Web开发", Author: "Eve", Price: rand.Float64()*30 + 20},
	}
	inserts := `insert into books(title, author, price) values (:title, :author, :price)`
	_, err := db.NamedExec(inserts, book)
	return err
}

// 查询价格在 20 到 50 之间的书籍。
// 查询特定作者的书籍。
// 获取每个作者的书籍数量和平均价格。
func QueryBook(db *sqlx.DB, authors []string) ([]Book, error) {
	if authors == nil || len(authors) == 0 {
		return nil, errors.New("作者不能为空")
	}
	query := `
	select author,count(*) AS bookCount, avg(price) as price
	from books
	where price between 20 and 50
	  and author in (?)
	group by author`
	query, args, err := sqlx.In(query, authors)
	if err != nil {
		return nil, err
	}
	var books []Book
	errs := db.Select(&books, query, args...)
	return books, errs
}

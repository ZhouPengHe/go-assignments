package main

import (
	"fmt"
	"go-assignments/task3"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func initGormDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"123456789",
		"127.0.0.1",
		"3306",
		"gorm",
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("数据库连接失败：%v", err)
	}
	return db, nil
}

func initSQLxDB() (*sqlx.DB, error) {
	sqlxDb, err := sqlx.Open("mysql", "root:123456789@tcp(127.0.0.1:3306)/gorm")
	if err != nil {
		return nil, fmt.Errorf("sqlx数据库连接失败：%v", err)
	}
	if err := sqlxDb.Ping(); err != nil {
		return nil, fmt.Errorf("sqlx数据库连接检查失败：%v", err)
	}
	return sqlxDb, nil
}

func main() {
	//SQL语句练习 -----------------------------------------------------------------
	db, err := initGormDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()
	// 执行转账操作
	err = task3.TransactionMethod(db)
	if err != nil {
		log.Printf("转账失败：%v", err)
	} else {
		log.Println("转账成功")
	}
	task3.StudentsCRUD(db)

	//Sqlx入门 -----------------------------------------------------------------
	sqlxDb, err := initSQLxDB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlxDb.Close()
	task3.CreateEmployees(sqlxDb)
	task3.InitEmployeeData(sqlxDb)
	department, _ := task3.QueryTechnicalDepartment(sqlxDb, "技术部")
	for _, employee := range department {
		fmt.Printf("技术部[%v]信息：%v\n", employee.Name, employee)
	}
	salary, _ := task3.QueryMaxSalary(sqlxDb)
	fmt.Printf("工资最高的员工信息是[%v], 工资为：[%v]元\n", salary.Name, salary.Salary)

	task3.CreateBook(sqlxDb)
	task3.InitBookData(sqlxDb)
	authors := []string{"Charlie", "Eve"}
	book, err := task3.QueryBook(sqlxDb, authors)
	if err != nil {
		log.Fatalf("查询失败：%v", err)
	}
	for _, b := range book {
		fmt.Printf("作者: %s, 书籍数量: %v, 价格: %.2f\n", b.Author, b.BookCount, b.Price)
	}

	//进阶gorm-----------------------------------------------------------------
	task3.InitGormDB(db)
	info, _ := task3.QueryUserInfo(db, "张三")
	fmt.Println(info)

	num, _ := task3.QueryLargestCommentsNum(db)
	fmt.Println(num)

	task3.DeleteComment(db, 6)
}

package task3

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
type Student struct {
	Id    uint   `gorm:"primaryKey"`
	Name  string `gorm:"type:varchar(100)"`
	Age   uint
	Grade string `gorm:"type:varchar(100)"`
}

func StudentsCRUD(db *gorm.DB) {
	db.AutoMigrate(&Student{})
	CreateStudent(db)
	QueryStudentAge(db)
	UpdateStudentGrade(db)
	DeleteStudentAge(db)
}

// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
func CreateStudent(db *gorm.DB) error {
	stu := Student{Name: "张三", Age: 20, Grade: "三年级"}
	return db.Create(&stu).Error
}

// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
func QueryStudentAge(db *gorm.DB) ([]Student, error) {
	stus := []Student{}
	err := db.Where("age > ?", 18).Find(&stus).Error
	return stus, err
}

// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
func UpdateStudentGrade(db *gorm.DB) error {
	return db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级").Error
}

// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
func DeleteStudentAge(db *gorm.DB) error {
	return db.Where("age < ?", 15).Delete(&Student{}).Error
}

// 假设有两个表：
//
//	accounts 表（包含字段 id 主键， balance 账户余额）和
//	transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
type Account struct {
	Id      uint    `gorm:"primaryKey;comment:账号ID"`
	Balance float64 `gorm:"type:decimal(10,4);not null;comment:账户余额"`
}

type Transaction struct {
	Id            uint    `gorm:"primaryKey;comment:交易Id"`
	FormAccountId uint    `gorm:"not null;comment:转出账户ID"`
	ToAccountId   uint    `gorm:"not null;comment:转入账户ID"`
	Amount        float64 `gorm:"type:decimal(10,4);not null;comment:转账金额"`
}

func InitAccounts(db *gorm.DB) {
	db.AutoMigrate(&Account{}, &Transaction{})
	accs := []Account{
		{Id: 1, Balance: 100},
		{Id: 2, Balance: 50},
	}
	db.Create(&accs)
}

// 要求 ：
//
//	编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
//	在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
//	并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
func TransactionMethod(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		formAccount := Account{}
		if err := tx.First(&formAccount, 1).Error; err != nil {
			return fmt.Errorf("转出账户不存在：%v", err)
		}

		if formAccount.Balance < 100 {
			return errors.New("余额不足")
		}

		formAccount.Balance -= 100
		if err := tx.Save(&formAccount).Error; err != nil {
			return fmt.Errorf("更新转出账户失败: %v", err)
		}

		toAccount := Account{}
		if err := tx.First(&toAccount, 2).Error; err != nil {
			return fmt.Errorf("转入账户不存在：%v", err)
		}
		toAccount.Balance += 100
		if err := tx.Save(&toAccount).Error; err != nil {
			return fmt.Errorf("更新转入账户失败: %v", err)
		}

		transactions := Transaction{FormAccountId: formAccount.Id, ToAccountId: toAccount.Id, Amount: 100}
		if err := tx.Create(&transactions).Error; err != nil {
			return fmt.Errorf("记录转账信息失败: %v", err)
		}
		return nil
	})
}

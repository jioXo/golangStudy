package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	//题目1
	//method1(db)

	//题目2
	method2()
}

/**
/**
题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/

func method1() {
	// 连接数据库
	db := ConnectDB()
	// 自动迁移，创建表
	err := db.AutoMigrate(&Student{})
	if err != nil {
		panic(err)
	}

	// 插入一条新记录
	err = CreateStudent(db, "张三", 20, "三年级")
	err = CreateStudent(db, "李四", 25, "五年级")
	if err != nil {
		panic(err)
	}

	// 查询所有年龄大于18岁的学生信息
	var students []Student
	db.Where("age > ?", 18).Find(&students)
	for _, student := range students {
		println(student.Name, student.Age, student.Grade)
	}

	// 更新姓名为"张三"的学生年级为"四年级"
	db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")

	// 删除年龄小于15岁的学生记录
	db.Where("age < ?", 15).Delete(&Student{})
}

type Student struct {
	ID    int    `gorm:"primaryKey"`
	Name  string `gorm:"column:name"`
	Age   int    `gorm:"column:age"`
	Grade string `gorm:"column:grade"`
}

// CreateStudent 插入一条新记录
func CreateStudent(db *gorm.DB, name string, age int, grade string) error {
	student := Student{Name: name, Age: age, Grade: grade}
	result := db.Create(&student)
	return result.Error
}

// ConnectDB 连接数据库
func ConnectDB() *gorm.DB {
	dsn := "root:admin@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

/*
*
题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，
向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/
func method2() {
	// 连接数据库
	db := ConnectDB()
	// 自动迁移，创建表
	err := db.AutoMigrate(&Acounnt{}, &Transaction{})
	if err != nil {
		panic(err)
	}
	// 假设账户 A 和账户 B 的 ID 分别为 1 和 2
	fromAccountID := 1
	toAccountID := 2
	// 假设账户 A 的余额为 500 元，账户 B 的余额为 300 元
	amount := 100.0
	// 初始化账户 A 和账户 B 的余额
	//db.Create(&Acounnt{ID: fromAccountID, Balance: 500.0})
	//db.Create(&Acounnt{ID: toAccountID, Balance: 300.0})
	// 打印初始余额
	var fromAccount Acounnt
	if err := db.First(&fromAccount, fromAccountID).Error; err != nil {
		panic(err)
	}
	var toAccount Acounnt
	if err := db.First(&toAccount, toAccountID).Error; err != nil {
		panic(err)
	}
	println("初始账户 A 余额:", fromAccount.Balance)
	println("初始账户 B 余额:", toAccount.Balance)

	// 打印转账金额
	println("转账金额:", amount)

	// 执行转账操作
	err = Transfer(db, fromAccountID, toAccountID, amount)
	if err != nil {
		println("转账失败:", err.Error())
	} else {
		println("转账成功")
	}

}

type Acounnt struct {
	ID      int     `gorm:"primaryKey"`
	Balance float64 `gorm:"column:balance"`
}

type Transaction struct {
	ID            int     `gorm:"primaryKey"`
	FromAccountID int     `gorm:"column:from_account_id"`
	ToAccountID   int     `gorm:"column:to_account_id"`
	Amount        float64 `gorm:"column:amount"`
}

func Transfer(db *gorm.DB, fromAccountID int, toAccountID int, amount float64) error {
	// 自动迁移，创建表
	err := db.AutoMigrate(&Acounnt{}, &Transaction{})
	if err != nil {
		panic(err)
	}
	// 开始事务
	tx := db.Begin()

	// 检查账户 A 的余额是否足够
	var fromAccount Acounnt
	if err := tx.First(&fromAccount, fromAccountID).Error; err != nil {
		tx.Rollback()
	}
	if fromAccount.Balance < amount {
		tx.Rollback()
		println("余额不足，无法转账")
		return nil // 余额不足，回滚事务
	}
	// 扣除账户 A 的余额
	fromAccount.Balance -= amount
	if err := tx.Save(&fromAccount).Error; err != nil {
		tx.Rollback()
		return err // 扣除失败，回滚事务
	}
	// 增加账户 B 的余额
	var toAccount Acounnt
	if err := tx.First(&toAccount, toAccountID).Error; err != nil {
		tx.Rollback()
		return err // 账户 B 不存在，回滚事务
	}
	toAccount.Balance += amount
	if err := tx.Save(&toAccount).Error; err != nil {
		tx.Rollback()
		return err // 增加失败，回滚事务
	}
	// 记录转账信息
	transaction := Transaction{
		FromAccountID: fromAccountID,
		ToAccountID:   toAccountID,
		Amount:        amount,
	}
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return err // 记录失败，回滚事务
	}
	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err // 提交失败，回滚事务
	}
	return nil // 转账成功
}

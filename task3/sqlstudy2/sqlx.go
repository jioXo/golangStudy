package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	//method1()
	method2()
}

/*
*
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/
func method1() {
	db := ConnectDB()
	// 创建表（如果不存在）
	// _, err := db.Exec(`CREATE TABLE IF NOT EXISTS employees (
	// 	id INT AUTO_INCREMENT PRIMARY KEY,
	// 	name VARCHAR(100),
	// 	department VARCHAR(100),
	// 	salary DOUBLE
	// )`)
	// if err != nil {
	// 	panic(err)
	// }
	// 插入一些测试数据
	// _, err = db.Exec("INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)", "张三", "技术部", 8000)
	// _, err = db.Exec("INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)", "李四", "市场部", 6000)
	// _, err = db.Exec("INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)", "王五", "技术部", 16000)
	// _, err = db.Exec("INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)", "赵六", "市场部", 3000)

	// 查询所有部门为 "技术部" 的员工信息
	employees, err := GetEmployeesByDepartment(db, "技术部")
	if err != nil {
		panic(err)
	}
	for _, emp := range employees {
		println(emp.ID, emp.Name, emp.Department, emp.Salary)
	}

	// 查询工资最高的员工信息
	highestSalaryEmployee, err := GetHighestSalaryEmployee(db)
	if err != nil {
		panic(err)
	}
	println("工资最高的员工：", highestSalaryEmployee.ID, highestSalaryEmployee.Name, highestSalaryEmployee.Department, highestSalaryEmployee.Salary)
	// 关闭数据库连接
	defer db.Close()
}

func ConnectDB() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/gorm")
	if err != nil {
		panic(err)
	}
	return db
}

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

// 查询指定部门的员工信息
func GetEmployeesByDepartment(db *sqlx.DB, department string) ([]Employee, error) {
	var employees []Employee
	query := "SELECT id, name, department, salary FROM employees WHERE department = ?"
	err := db.Select(&employees, query, department)
	if err != nil {
		return nil, err
	}
	return employees, nil
}

// 查询工资最高的员工信息
func GetHighestSalaryEmployee(db *sqlx.DB) (*Employee, error) {
	var employee Employee
	query := "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1"
	err := db.Get(&employee, query)
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

/*
*
题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/
func method2() {
	db := ConnectDB()

	// 创建表（如果不存在）
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS books (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255),	
		author VARCHAR(100),
		price DOUBLE
	)`)
	if err != nil {
		panic(err)
	}
	// 插入一些测试数据
	_, err = db.Exec("INSERT INTO books (title, author, price) VALUES (?, ?, ?)", "Go语言编程", "张三", 45.0)
	_, err = db.Exec("INSERT INTO books (title, author, price) VALUES (?, ?, ?)", "Python编程", "李四", 55.0)
	_, err = db.Exec("INSERT INTO books (title, author, price) VALUES (?, ?, ?)", "Java编程", "王五", 60.0)
	if err != nil {
		panic(err)
	}
	//查询价格大于50元的书籍
	books, err := GetBooksByPrice(db, 50.0)
	if err != nil {
		panic(err)
	}
	for _, book := range books {
		println(book.ID, book.Title, book.Author, book.Price)
	}
}

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

// 查询价格大于指定值的书籍
func GetBooksByPrice(db *sqlx.DB, minPrice float64) ([]Book, error) {

	var books []Book
	query := "SELECT id, title, author, price FROM books WHERE price > ?"
	err := db.Select(&books, query, minPrice)
	if err != nil {
		return nil, err
	}
	return books, nil
}

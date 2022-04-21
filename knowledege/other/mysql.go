package knowledege

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type user struct {
	id   int
	name string
	age  int
}

//把数据库连接池声明到全局 这样才能够方便的进行操作
var db *sql.DB

func InitDb() (err error) {
	//数据库信息
	dsn := "root:Zodiac100@tcp(127.0.0.1:3306)/gotest"
	//连接数据库
	db, err = sql.Open("mysql", dsn) //不会校验用户名和密码
	if err != nil {                  //dsn格式不正确的时候会报错
		return
	}
	err = db.Ping() //进行校验用户名和密码
	if err != nil {
		return
	}
	//设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	//设置最大空闲连接数
	db.SetMaxIdleConns(2)
	fmt.Println("连接数据库成功")
	queryOne(2)
	insertUser("iux", 20)
	updateRow(18, 2)
	deleteUse(2)
	perpareInsert()
	queryMore(0)
	return
}

//单行查询
func queryOne(id int) {
	//1.写查询单条记录的sql语句
	sqlStr := `select id,name,age from user where id=?;`
	//2.执行 1会替换语句中的?
	rowObj := db.QueryRow(sqlStr, id) //从连接池里拿一个连接出来去数据库查询单条记录
	//3.拿到结果
	var u user
	//使用scan会拿到结果，并且必须要吊用scan，因为scan函数内部会释放连接
	rowObj.Scan(&u.id, &u.name, &u.age)
	fmt.Println(u)
}

//多行查询
func queryMore(id int) {
	sqlStr := "select id,name,age from user where id > ?;"
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	//但一定要关闭连接
	defer rows.Close()
	//循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
			return
		}
		fmt.Println(u)
	}
}

//插入数据
func insertUser(name string, age int) {
	//1.写sql语句
	sqlStr := `insert into user(name,age) values(?,?);`
	//2. exec去执行上面的语句
	ret, err := db.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	//如果是插入数据的操作 能够拿到插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(id)
}

//更新操作
func updateRow(age, id int) {
	sqlStr := `update user set age=? where id>?;`
	ret, err := db.Exec(sqlStr, age, id)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	//获取更新了多少行
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	fmt.Println("更新了", n, "行")
}

//删除数据
func deleteUse(id int) {
	sqlStr := `delete from user where id=?;`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	//获取删除了多少行
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	fmt.Println("删除了", n, "行")
}

//MySql预处理
/*
1.把SQL语句分成两部分，命令部分和数据部分
2.先把命令部分发送给MySql服务端，MySQL服务端进行SQL预处理
3.然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换
4.MySQL服务端执行完整的SQL语句并将结果返回给客户端
优点：
1.优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，
一次编译多次执行，节省后续编译的成本。
2.避免SQL注入的问题。
*/

func perpareInsert() {
	sqlStr := `insert into user(name,age) values(?,?);`
	//将sql语句传给Mysql服务器 进行预编译
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	defer stmt.Close()
	//后续只需要通过stmt进行操作
	m := map[string]int{
		"adq": 12,
		"rvw": 12,
		"rv":  21,
	}
	for k, v := range m {
		stmt.Exec(k, v)
	}
}

//MySql事物操作
//事物就是把多于两个操作的语句进行合并操作
//有三个方法 Begin开始事物 Commit提交事物 Rollback回滚事物
func transactionDemo() {
	//开启事务
	tx, err := db.Begin()
	if err != nil {
		return
	}
	//执行多个操作
	sqlStr1 := `update user set age=age-2 where id=9;`
	_, err = db.Exec(sqlStr1)
	if err != nil {
		//要回滚
		tx.Rollback()
		return
	}
	sqlStr2 := `update user set age=age+2 where id=10;`
	_, err = db.Exec(sqlStr2)
	if err != nil {
		//要回滚
		tx.Rollback()
		return
	}
	//如果两个操作都成功了 就将操作进行提及哦啊
	err = tx.Commit()
	if err != nil {
		//要回滚
		tx.Rollback()
		return
	}
	fmt.Println("事物执行成功")
}

var dbx *sqlx.DB

//使用sqlx简化一些操作
func InitDbX() (err error) {
	//数据库信息
	dsn := "root:Zodiac100@tcp(127.0.0.1:3306)/gotest"
	//连接数据库
	dbx, err = sqlx.Connect("mysql", dsn) //不会校验用户名和密码
	if err != nil {                       //dsn格式不正确的时候会报错
		return
	}
	//Connect省略了ping的过程
	//设置数据库连接池的最大连接数
	dbx.SetMaxOpenConns(10)
	//设置最大空闲连接数
	dbx.SetMaxIdleConns(2)
	fmt.Println("连接数据库成功")
	queryOneX(10)
	queryMoreX(2)
	return
}

type userx struct {
	ID   int
	Name string
	Age  int
}

func queryOneX(id int) {
	sqlStr := `select id,name,age from user where id=?;`
	var u userx
	//可以直接对结构体赋值
	//一下函数为了得知user的结构体有哪些属性，所以内部使用了反射来获取user结构体的属性
	//所以结构体的属性就应该首字母大写，让sqlx这个库见到，因为小写对外不可见
	dbx.Get(&u, sqlStr, id)
	fmt.Println(u)
}
func queryMoreX(id int) {
	sqlStr := `select id,name,age from user where id>?;`
	var u []userx //不需要初始化
	//可以直接对结构体切片赋值
	dbx.Select(&u, sqlStr, id)
	fmt.Println(u)
}

//sql注入 永远不要自己拼接sql语句
//sql预编译

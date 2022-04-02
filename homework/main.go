package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"unicode"
)

func main() {
	InsertContent("./test.txt", "1", 3)
}

//打印9*9乘法表
func homewoke1() {
	for i := 1; i < 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d,", i, j, i*j)
		}
		fmt.Println()
	}
}

//两数之和
func twoSum(target int) (left, right int) {
	a := [...]int{1, 3, 5, 7, 8}
	l := 0
	r := len(a) - 1
	for l < r {
		if a[l]+a[r] == target {
			return l, r
		} else if a[l]+a[r] > target {
			r--
		} else {
			l++
		}
	}
	return 0, 0
}

//slice练习题
func sliceTest() {
	a := make([]int, 5, 10)
	for i := 0; i < len(a); i++ {
		a = append(a, i)
	}
	fmt.Println(a, cap(a)) //[0 0 0 0 0 1 2 3 4 5 6 7 8 9]

	sort.Ints(a[:]) //对切片进行排序

}

//判断字符串汉字的数量
//便利一个带有汉字的字符串的时候 不能直接通过索引便利 因为索引是字节数 一个汉字占三个字节，所以不能
//要先切成slice，再便利
func countHanzi() {
	//如何判断一个字符是汉字
	s := "我是汉字hello"
	for i, v := range s {
		//如何判断一个字符是汉字
		if unicode.Is(unicode.Han, v) {
			fmt.Println(i, v)
		}
	}
}

func pop(slice []int) []int {
	if len(slice) > 0 {
		return slice[1:]
	} else {
		return make([]int, 0)
	}
}

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

//分金币
func dispatchCoin() (left int) {
	left = coins
	rule := map[uint8]int{
		byte('e'): 1,
		byte('E'): 1,
		byte('i'): 2,
		byte('I'): 2,
		byte('o'): 3,
		byte('O'): 3,
		byte('u'): 4,
		byte('U'): 4,
	}
	for _, item := range users {
		distribution[item] = 0
		for _, word := range item {
			if v, ok := rule[uint8(word)]; ok {
				distribution[item] += v
				left -= v
			}
		}
	}
	fmt.Println(distribution, left)
	return
}

//学生管理系统
type Manager struct {
	//students  map[int]Student
	//下面使用引用的方式在map中使用结构体更好 更方便
	//因为如果是上面那种值的写法 你不能通过students[id].name = xxx 直接修改 因为你修改的是一个拷贝
	//而使用引用的写法却可以直接修改 students[id].name = xxx 这样修改的才是原本的值
	students  map[int]*Student
	StudenNum int
}
type Student struct {
	id   int
	name string
}

func (m *Manager) newStudent(id int, name string) Student {
	return Student{
		id,
		name,
	}
}
func (m *Manager) ShowAll() {
	for k, v := range m.students {
		fmt.Printf("%d:%v\n", k, v)
	}
}
func (m *Manager) AddStudent(id int, name string) {
	newStudent := m.newStudent(id, name)
	m.students[id] = &newStudent
	m.StudenNum++
}
func (m *Manager) EditStudent(id int, name string) {
	_, ok := m.students[id]
	if ok == false {
		fmt.Println("查无此人")
	}
	m.students[id].name = name
}
func (m *Manager) DeleteStudent(id int) {
	delete(m.students, id)
	m.StudenNum--
}
func work() {
	manager := Manager{
		students: make(map[int]*Student, 10),
	}
	for {
		fmt.Print(`
欢迎来到学生管理系统，
1. 查看全部学生
2. 添加学生
3. 删除学生
4. 退出
`)
		var chose int
		fmt.Scan(&chose)
		switch chose {
		case 1:
			manager.ShowAll()
		case 2:
			fmt.Println("请输入学生的id和name")
			var id int
			var name string
			fmt.Scanln(&id, &name)
			manager.AddStudent(id, name)
		case 3:
			fmt.Println("请输入学生的id")
			var id int
			fmt.Scanln(&id)
			manager.DeleteStudent(id)
		case 4:
			os.Exit(1)
		case 5:
			fmt.Println("请输入学生的id和name")
			var id int
			var name string
			fmt.Scanln(&id, &name)
			manager.EditStudent(id, name)
		default:
			fmt.Println("请输入正确的数字")
		}
	}
}

//CopyFile拷贝文件
func CopyFile(sourceName, targetName string) (err error) {
	content, err := ioutil.ReadFile(sourceName)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(targetName, content, 0644)
	if err != nil {
		return err
	}
	return err
}

func InsertContent(targetName, content string, line int) (err error) {
	temp, linNum := "", 0
	fileObj, err := os.Open(targetName)
	reader := bufio.NewReader(fileObj)
	for {
		if linNum == line {
			temp = temp + content + "\n"
			linNum++
			continue
		}
		line, err := reader.ReadString('\n') //按行读取
		fmt.Println(line)
		if err == io.EOF {
			temp += line
			linNum++
			break
		}
		if err != nil {
			fmt.Println("读文件过程中出错了")
			return err
		}
		temp += line
		linNum++
	}
	ioutil.WriteFile(targetName+"copy", []byte(temp), 777)
	return
}

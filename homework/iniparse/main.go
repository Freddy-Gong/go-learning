package iniparse

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type ConfIni struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

func Run(path string) {
	conf := ConfIni{}
	m := makeMap(&conf)
	fmt.Println(m)
	readIni(path, m, &conf)
	fmt.Printf("%#v", conf)
}
func readIni(path string, m map[string]string, conf *ConfIni) {
	fileObj, err := os.Open(path)
	if err != nil {
		fmt.Println("文件打不开")
	}
	defer fileObj.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			parse(line, m, conf)
			return
		}
		if err != nil {
			fmt.Printf("read line failed:%v\n", err)
			return
		}
		parse(line, m, conf)
	}
}

func parse(line string, m map[string]string, conf *ConfIni) {
	v := reflect.ValueOf(conf) //如果要通过反射修改结构体里某个字段的值 要通过valueof的结果进行修改
	t := reflect.TypeOf(conf)  //通过Typeof只能获得一些类型信息如 tag type kind
	if strings.Contains(line, "=") {
		option := strings.Split(line, "=")
		key, value := option[0], option[1]

		structKey := m[key]
		field, ok := t.Elem().FieldByName(structKey)
		if ok {
			switch field.Type.Kind() {
			case reflect.String:
				v.Elem().FieldByName(structKey).SetString(value)
			case reflect.Int:
				invalue, _ := strconv.ParseInt(strings.Trim(value, "\n"), 10, 64)
				v.Elem().FieldByName(structKey).SetInt(invalue)
			}
		}
	}
}
func makeMap(conf *ConfIni) map[string]string {
	t := reflect.TypeOf(conf)
	re := make(map[string]string, 10)
	for i := 0; i < t.Elem().NumField(); i++ {
		field := t.Elem().Field(i)
		key := field.Tag.Get("ini")
		value := field.Name
		re[key] = value
	}
	return re
}

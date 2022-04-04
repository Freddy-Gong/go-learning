package knowledege

import (
	"fmt"
	"reflect"
)

//反射是获取运行时变量的值以及类型的方式
//在Go中，任何接口值都是由一个具体类型和具体类型的值两部分组成的
//在反射中 这两个概念分别可由reflec.Type和reflect.Value表示
//所以refelct包提供了reflect.TypeOf和refelct.ValueOf两个函数来获取
//任意对象的Value和Type
//然鹅类型又分为type.Type 和type.Kind
//比如说你可名为person的结构体类型 它的Type就是person kind就是struct
//value也有value.Kind()
func Reflect(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.ValueOf(x))
	//通过reflect设置值
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //但是如果这样直接修改，它其实修改的副本
	}
	//我们需要的是通过指针来修改
	//假设现在v是一个指针 则需要调用Elem方法来进行修改
	// if v.Elem().Kind() == reflect.Int64 {
	// 	v.Elem().SetInt(200)
	// }
	//判断是否为空
	//v.IsNil()   //判断是否为空指针
	v.IsValid() //判断返回值是否有效

	//结构体反射
	//1. StructField类型用来描述结构体的每一个字段的信息
	type Student struct {
		Name  string `json:"name"`
		Score int    `json:"score"`
	}
	stu := Student{
		Name:  "小王子",
		Score: 90,
	}
	t := reflect.TypeOf(stu)
	fmt.Println(t.Name(), t.Kind())
	//通过for循环便利结构体的所有字段信息
	//t.NumField可以返回结构体里面字段的数量
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Println(field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}
	//还可以通过指定字段名获取指定的结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Println(scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}

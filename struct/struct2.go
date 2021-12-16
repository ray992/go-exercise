package main

import "fmt"

type student struct {
	name string
	age  int
	string //匿名字段
	dream []string
}

type fruit struct {
	name string
}

type pear struct {
	fruit //继承
	color string
	Weight int8 //大写公开访问
}

func (s student) updateName(name string)  { //方法
	s.name = name
}

func (s *student) updateAge (age int)  {
	s.age = age
}

func (s *student) setDreams(dreams []string)  { //注意切片及map的复值
	s.dream = make([]string, len(dreams))
	copy(s.dream, dreams)
}

func init()  {
	fmt.Println("开始测试。。。")
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

	ss := student{"mike", 21, "cat", []string{"teacher","doctor"}}
	ss.updateName("jenny")
	fmt.Println(ss)
	ss.updateAge(30)
	fmt.Println(ss)
	var newDreams = []string{"coder", "engineer", "singer"}
	ss.setDreams(newDreams)
	fmt.Println(ss)
	newDreams[0] = "dancer"
	fmt.Println(ss)
}

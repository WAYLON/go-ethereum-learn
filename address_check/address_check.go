package main

type student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	// 错误写法
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	for k, v := range m {
		println(k, "=>", v.Name)
	}

	// 正确
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}
	for k, v := range m {
		println(k, "=>", v.Name)
	}
}

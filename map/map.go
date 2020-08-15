package main

import (
	"fmt"
	"sort"
)

func main() {
	map1 := map[int]string{66: "golang", 22: "python", 88: "c++", 99: "php"}
	for k, v := range map1 {
		fmt.Println(k, "-->", v)

	}

	keys := make([]int, 0, len(map1))
	for k := range map1 {
		keys = append(keys, k)
	}
	fmt.Println("Keys is ", keys)

	sort.Ints(keys)

	fmt.Println("Afetr sorted keys is ", keys)

	for _, key := range keys {
		fmt.Println(key, "-->", map1[key])

	}

	s1 := []string{"pick", "AAA", "aaa", "jjj", "zzz"}
	fmt.Println(s1)
	sort.Strings(s1)
	fmt.Println(s1)

	fmt.Println("============================================")
	m1 := make(map[string]string)
	m1["name"] = "Gjj"
	m1["age"] = "23"
	m1["sex"] = "male"
	m1["job"] = "student"

	m2 := map[string]string{"name": "libai", "age": "33", "sex": "male", "job": "poet"}
	fmt.Println(m1)
	fmt.Println(m2)

	ss1 := make([]map[string]string, 0, 2)
	ss1 = append(ss1, m1)
	ss1 = append(ss1, m2)
	fmt.Println(ss1)
	for i, v := range ss1 {
		//v: m1,m2
		fmt.Printf("第%d个人的信息是:-------------------------\n", i+1)
		fmt.Println("姓名：", v["name"])
		fmt.Println("性别：", v["sex"])
		fmt.Println("职业：", v["job"])
		fmt.Println("年龄：", v["age"])
	}

}

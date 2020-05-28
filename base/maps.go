package main

import "fmt"

// map元素是无序的
func main() {
	m := map[string]string{
		"name":    "zhangsan",
		"course":  "pe",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int)

	var m3 = map[string]int{}

	fmt.Println(m)
	fmt.Println(m2)
	fmt.Println(m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	name := m["name"]
	fmt.Println(name)
}

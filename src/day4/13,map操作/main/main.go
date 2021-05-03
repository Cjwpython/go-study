package main

import (
	"fmt"
)

func testMap() {
	var a map[string]string = map[string]string{
		"key": "value",
	}
	a["abc"] = "cfg"
	fmt.Println(a)
}

func testMap2() {
	a := make(map[string]map[string]string)
	b := make(map[string]string)
	b["213"] = "12312"
	a["345"] = b
	fmt.Println(a)
}

func testMap3() {
	a := make(map[string]map[string]string)
	a["345"] = map[string]string{}
	a["345"]["231"] = "123123"

	vul, ok := a["key1"]
	if ok {
		fmt.Println(vul)
	} else {
		fmt.Println(ok)
	}
	fmt.Println(a)
}

func testMap4() {
	a := make(map[string]map[string]string)
	a["345"] = map[string]string{}
	a["345"]["231"] = "123123"
	for k, v := range a {
		fmt.Println("key:", k, "value", v)
	}
}

func testMap5() {
	var a []map[int]int
	a = make([]map[int]int, 5)
	if a[0] == nil {
		a[0] = make(map[int]int)
		//a[0] = map[int]int{}
	}
	a[0][10] = 10
	fmt.Println(a)

}
func main() {
	//testMap()
	//testMap2()
	//testMap3()
	//testMap4()
	testMap5()
}

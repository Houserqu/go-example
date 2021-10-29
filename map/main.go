package main

import "fmt"

func main() {
	var mapList map[string]int
	mapList = map[string]int{"one": 1, "two": 2}

	fmt.Println(mapList)
	fmt.Println(&mapList)

	setMap(mapList)

	fmt.Println(mapList)

}

func setMap(m map[string]int) {
	fmt.Println(&m)
	m["one"] = 3
}

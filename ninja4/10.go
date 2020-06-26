package main

import "fmt"

func main() {
	m := map[string][]string{"james": []string{"bond", "mf", "womanizer"}, "regularJoe": []string{"regular", "mf", "once a month"}}
	//fmt.Println(m)
	m["arionT"] = []string{"tudi","katya","always"}
	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
	delete(m,"james")
	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
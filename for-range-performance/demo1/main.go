package main

import "fmt"

// func main() {
// 	words := []string{"Go", "语言", "很", "棒"}
// 	// for i, s := range words {
// 	// 	words = append(words, "test")
// 	// 	fmt.Println(i, s)
// 	// }
// 	for i := range words {
// 		fmt.Println(i, words[i])
// 	}
// }

// func main() {
// 	m := map[string]int{
// 		"one":   1,
// 		"two":   2,
// 		"three": 3,
// 	}
// 	for k, v := range m {
// 		delete(m, "two")
// 		m["four"] = 4
// 		fmt.Printf("%v,%v\n", k, v)
// 	}
// }

func main() {
	ch := make(chan string)
	go func() {
		ch <- "Go"
		ch <- "语言"
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}
}

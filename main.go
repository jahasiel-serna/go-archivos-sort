package main

import (
	"fmt"
	"sort"
	"os"
)

func sortSlice(s []string, ascendent bool, fileName string) {
	if ascendent {
		sort.Strings(s)
	} else {
		sort.Sort(sort.Reverse(sort.StringSlice(s)))
	}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range s {
		file.WriteString(v)
		file.WriteString("\n")
	}
	file.Close()
}

func main() {
	var n int
	var temp string
	fmt.Print("n: ")
	fmt.Scan(&n)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Print(i + 1, " ")
		fmt.Scan(&temp)
		s[i] = temp
	}
	
	sortSlice(s, true, "ascendente.txt")
	sortSlice(s, false, "descendente.txt")
}

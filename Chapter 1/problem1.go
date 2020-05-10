package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	textFile, err := os.OpenFile("./numbers.txt", os.O_RDWR, os.ModeAppend)
	check(err)
	defer textFile.Close()

	// for i := 0; i < 10000; i++ {
	// 	rand.Seed(time.Now().UnixNano())
	// 	num := rand.Intn(10000000)
	// 	_, err3 := textFile.WriteString(fmt.Sprintf("%d\n", num))
	// 	check(err3)
	// }

	s := make([]int, 10000)
	scanner := bufio.NewScanner(textFile)
	for scanner.Scan() {
		text := scanner.Text()
		i, err := strconv.Atoi(text)
		check(err)
		s = append(s, i)
	}
	sort.Ints(s)
	fmt.Println(s)
}

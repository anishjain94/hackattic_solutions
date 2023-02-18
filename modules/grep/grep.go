package grep

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func Grep(searched_string string, filename string, c chan []string) {
	var output_lines []string
	var file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755) //create if not exists
	if err != nil {                                                    //if any error (permission, user)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) //scanning the contents
	for scanner.Scan() {              //reading file line by line
		if strings.Contains(scanner.Text(), searched_string) {
			output_lines = append(output_lines, scanner.Text()) //storing if pattern matches.
		}
	}
	c <- output_lines
	// fmt.Println(output_lines, <-c)
}

func main() {
	var files []string
	var matches []string
	var wg sync.WaitGroup

	c := make(chan []string, 3)

	root := "/Users/anishjain/go/src/hackattic_solutions"
	string_to_be_searched := "go"
	_, err1 := os.Stat(root)

	if os.IsNotExist(err1) {
		log.Fatal("File/Directory does not exist.")
	}
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) != ".go" {
			return nil
		}
		files = append(files, path)
		return nil
	}) //getting all the .text files from all the sub folders
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		wg.Add(1) //adding count of goroutines for waitgroup
		go Grep(string_to_be_searched, file, c)
	}
	go func() {
		for i := range c {
			matches = append(matches, i...)
			wg.Done() // decrement the counter as
		}
	}()
	wg.Wait() //wait till counter goes to 0
	println(matches)

}

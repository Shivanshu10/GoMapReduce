package main

import (
	"GoMapReduce/client/clientlib"
	"fmt"
	"log"
	"os"
	"strconv"
)

func printIntro() {
	fmt.Println("Welcome!!!")
	fmt.Println("Please choose one of option")
	fmt.Println("1. Create New Task")
}

func readInput(query string) string {
	fmt.Printf(query + ": ")

	var data string

	_, err := fmt.Scanln(&data)

	if err != nil {
		log.Fatal("Cannot read from Std Input: {}", err)
		os.Exit(1)
	}

	return data
}

func strToInt(str string) int64 {
	data, err := strconv.ParseInt(str, 0, 8)
	if err != nil {
		log.Fatal("Cannot convert to Int: {}", err)
		os.Exit(1)
	}

	return data
}

func AddTaskCmd() {
	num_files := strToInt(readInput("Number Of File"))

	var files []string
	var func_file string
	var n_reduce int64

	var i int64 = 0
	for ; i < num_files; i++ {
		files = append(files, readInput("file name"))
	}

	func_file = readInput("func file name")

	n_reduce = strToInt(readInput("Number of Reducers"))

	if clientlib.AddTask(files, func_file, n_reduce) {
		log.Println("Task Added")
		fmt.Println("Task Added")
	} else {
		log.Println("Failed to Add Task")
		fmt.Println("Failed to Add Task")
	}
}

func main() {
	printIntro()

	option := strToInt(readInput("Option>"))

	if option == 1 {
		AddTaskCmd()
	}
}

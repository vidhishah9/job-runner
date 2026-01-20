package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {

	// Open the file
    file, err := os.Open("tasks.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    // Create a scanner
    scanner := bufio.NewScanner(file)

	//create a new channel called tasks
	tasks := make(chan string)
	//start a goroutine to read the file line by line and send each line to the tasks channel
    go readFileLineByLine(tasks, scanner)
    task := <-tasks
    fmt.Println(task)

}

func readFileLineByLine(tasks chan <-string, scanner *bufio.Scanner) {
    fmt.Println("Reading file line by line...")
	// Read and print lines
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println(line)

        tasks <- line
    }
    // Check for errors
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }

    
}
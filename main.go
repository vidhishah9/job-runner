package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
)

func worker(id int, tasks <-chan string) {
    fmt.Println("Starting worker", id)

    for task := range tasks {
        fmt.Println(task)
    }
}


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

    time.Sleep(10 * time.Millisecond) //wait a bit for queue to load some tasks

    //start 3 worker goroutines to receive tasks from the tasks channel and print them
    for w := 1; w <= 3; w++ {
        go worker(w, tasks)
    }
    time.Sleep(50 * time.Millisecond) //make sure main doesn't exit before goroutines finish
}


func readFileLineByLine(tasks chan <-string, scanner *bufio.Scanner) {
    fmt.Println("Reading file line by line...")
	// Read and print lines
    for scanner.Scan() {
        line := scanner.Text()
        // fmt.Println(line)

        tasks <- line
    }
    // Check for errors
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }

    
}
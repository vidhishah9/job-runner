package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
    "log"
    "net/http"

)

func worker(id int, tasks <-chan string) {

    fmt.Println("Starting worker", id)
    for task := range tasks {
        log.Printf("worker %d got %s", id, task)
        resp, err := http.Get(task)
        if err != nil {
            log.Fatalln(err)
        }
        //We Read the response body on the line below.
        if err != nil {
            log.Fatalln(err)
        }
        //Convert the body to type string
        log.Printf("client: got response!\n")
        log.Printf(resp.Status)
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


    //start 3 worker goroutines to receive tasks from the tasks channel and print them
    for w := 1; w <= 3; w++ {
        go worker(w, tasks)

    }

    //start a goroutine to read the file line by line and send each line to the tasks channel
    readFileLineByLine(tasks, scanner)
    close(tasks)
    time.Sleep(500 * time.Millisecond) //make sure main doesn't exit before goroutines finish
}


func readFileLineByLine(tasks chan <-string, scanner *bufio.Scanner) {
    fmt.Println("Reading file line by line...")
	// Read and print lines
    for scanner.Scan() {
        line := scanner.Text()
        tasks <- line
    }
    // Check for errors
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
    

    
}
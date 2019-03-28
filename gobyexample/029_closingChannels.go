package main

import "fmt"

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

	// infinite loop function
	// pulling jobs and processing them
    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

	// sending j to jobs
    for j := 1 ; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent jobs", j)
    }

    close(jobs)
	fmt.Println("sent all jobs")
	
	// wait here till all jobs are done
    <- done
}
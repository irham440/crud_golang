package utils

import (
	"belajar-go/models"
	"fmt"
	"time"
)

var EmailQueue = make(chan models.EmailJob, 200)

func StartWorker() {
	for i := 1; i <= 100; i++ {
		go SendEmailWorker(i)
	}
}

func SendEmailWorker(id int) {
	for job := range EmailQueue {
		time.Sleep(10 * time.Second)
		fmt.Println("Worker", id, "sent email to:", job.To, "body :", job.Body)
	}
}

func EnqueueEmail(job models.EmailJob) {
	EmailQueue <- job
}

package main

import (
	"context"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
	"worker-pool/workerpool"
)

func main() {
	r := gin.Default()

	dispatch := workerpool.NewDispatcher(10)
	dispatch.Run()
	workerpool.InitJobQueue()

	r.GET("/test", func(c *gin.Context) {
		msg := c.DefaultQuery("msg", "default message")
		job := workerpool.NewJob(func(ctx context.Context) error {
			time.Sleep(time.Second)
			log.Println("I am job, message: ", msg)
			return nil
		})
		workerpool.JobQueue <- job
		c.JSON(http.StatusOK, gin.H{
			"message": msg,
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

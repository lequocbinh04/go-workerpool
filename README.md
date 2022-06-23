# Simple concurrent worker pool in GO

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://go.dev/)

Simple implement concurrent worker pool & queue in go. You can test it by simple command below.

## Clone the project

```
$ git clone https://github.com/lequocbinh04/go-workerpool
$ cd go-workerpool
```
## Build project

```
$ go build
```

##Run
```
$ ./worker-pool
```

Then you can go to [https://localhost:8080/test?msg=Hello worker](https://localhost:8080/test?msg=Hello%20worker) and check log. 

##Apply in your project
You can copy workerpool folder to your project. Init worker pool:
````go
dispatch := workerpool.NewDispatcher(10) // start 10 workers
dispatch.Run()
workerpool.InitJobQueue()
````

and add job to pool:
```go
job := workerpool.NewJob(func(ctx context.Context) error {
	// simple job, replace it :D
    time.Sleep(time.Second)
    log.Println("I am job, message: ", msg)
    return nil
})
workerpool.JobQueue <- job
```


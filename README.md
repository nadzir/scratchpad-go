# scratchpad-go
Playground with go

## Table of contents
- [Setup Dependencies](#Dependencies)
- [Simple Rest API](#Simple-Rest-Api)
- [RabbitMQ](#RabbitMQ)
- [Job Analysis](#Job Analysis)

---

### Dependencies

Using Gin to handle auto-reload
>https://github.com/codegangsta/gin
```
go get github.com/codegangsta/gin
gin run main.go
```

---

### Simple Rest Api

What is it ?

Minimum code to set up rest api using go

Where is the code ?

```
cd ./simpleRestEndpoint
```

How to run ?
```
go run simpleRestEndpoint
curl localhost:3100/simpleEndpoint
```

### Rabbit MQ

What is it ?

The code for the [rabbit mq tutorial](https://www.rabbitmq.com/getstarted.html)


Where is the code ?

```
cd ./rabbitmq
```

### Job Analysis

What is it ?

Experiementing with go project structure, queue and event sourcing.
Trying to analysis the jobs in singapore.

Where is the code ?

```
cd ./job-analysis
```

How to run ?

```
// Start crawling the jobs
cd job-analysis/job-crawler/cmd/main
go run main.go

// Analyse duplicated job description
ob-analysis/job-job-duplicate-checker/cmd/main
go run main.go
```
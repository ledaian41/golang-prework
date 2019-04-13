package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type responseInfo struct {
	status         int
	bytes          int64
	duration       time.Duration
	serverHostname string
	port           string
}

type summaryInfo struct {
	requested int64
	responsed int64
}

type requestInfo struct {
	link    string
	timeout time.Duration
}

func main() {
	fmt.Println("Hi there, I'm An\nHere's my prework")
	requests := flag.Int64("n", 1, "Number of requests to perform")
	concurrency := flag.Int64("c", 1, "Number of multiple requests to make at a time")
	timeout := flag.Int64("t", 1, "Seconds to max. wait for each response")
	timeLimit := flag.Int64("l", 1, "Maximum number of seconds to spend for benchmarking")
	flag.Parse()
	parsedTimeout := parseTimeToSecond(*timeout)
	parsedTimeLimit := parseTimeToSecond(*timeLimit)

	if isValidCommandLine(requests, concurrency) {
		flag.PrintDefaults()
		os.Exit(-1)
	}
	link := flag.Arg(0)

	requestDetails := requestInfo{
		link:    link,
		timeout: parsedTimeout,
	}

	c := make(chan responseInfo)
	summary := summaryInfo{}

	start := time.Now()
	for i := int64(0); i < *concurrency; i++ {
		summary.requested++
		go checkLink(requestDetails, c)
	}

	for response := range c {
		if summary.requested < *requests {
			summary.requested++
			if isExceedTimeLimit(start, parsedTimeLimit) {
				fmt.Println("exceed limit time")
				return
			}
			go checkLink(requestDetails, c)
		}
		summary.responsed++
		fmt.Println(response)
		if summary.responsed == summary.requested {
			break
		}
	}
}

func checkLink(request requestInfo, c chan responseInfo) {
	start := time.Now()
	timeout := time.Duration(request.timeout) * time.Second
	client := &http.Client{
		Timeout: timeout,
	}
	res, err := client.Get(request.link)
	if err != nil {
		panic(err)
	}
	read, _ := io.Copy(ioutil.Discard, res.Body)
	c <- responseInfo{
		status:         res.StatusCode,
		bytes:          read,
		duration:       time.Now().Sub(start),
		serverHostname: res.Header.Get("Server"),
	}
}

func isValidCommandLine(request *int64, concurrency *int64) bool {
	return flag.NArg() == 0 || *request == 0 || *request < *concurrency
}

func parseTimeToSecond(value int64) time.Duration {
	return time.Duration(value) * time.Second
}

func isExceedTimeLimit(start time.Time, limit time.Duration) bool {
	return time.Now().Sub(start) > limit
}

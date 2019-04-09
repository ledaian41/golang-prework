# golang-prework
Coderschool - Go Lang pre-work assignment

### To build:
go build -o [name] main.go

### To run:
./[name] -n [number] -c [number] [URL]

## Required Features

The program must accept the following command line arguments:
* Requests - Number of requests to perform
* Concurrency - Number of multiple requests to make at a time
* URL - The URL for testing.

The program prints usage information if the wrong arguments are provided.
The program performs the specified HTTP requests and prints a summary of the results.
The program’s concurrency is implemented with goroutines.

## Bonus Features

These aren’t described in the video, but a great exercise to get ready for class.

Extend input params with:
* Timeout - Seconds to max. wait for each response
* Timelimit - Maximum number of seconds to spend for benchmarking
* Prints key metrics in summary, for example:
* Server Hostname
* Server Port
* Document Path
* Document Length
* Concurrency Level
* Time taken for tests
* Complete requests
* Failed requests
* Total transferred
* Requests per second
* Time per request
* Transfer rate

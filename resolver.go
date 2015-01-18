package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
    "net"
    "runtime"
    "sync"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func worker(w int, job <-chan string, wg *sync.WaitGroup) {
  // Decreasing internal counter for wait-group as soon as goroutine finishes
  defer wg.Done()
  for j := range job {
    domain := j
    address, err := net.LookupHost(domain)
    if err != nil {
      fmt.Println("Domain:", domain, "did not resovle")
    } else {
      fmt.Println("Domain:", domain, "has address(es):", address)
    }
  }
}


func main() {

  if len(os.Args) != 2 {
    fmt.Println("Error: You must supply an argument with the location of the file that contains Domain Names to resolve - 1 per line")
    os.Exit(1)
  }

  // run with maximum available procs
  runtime.GOMAXPROCS(runtime.NumCPU())

  // open file containing domains
  file, err := os.Open(os.Args[1])
  check(err)
  defer file.Close()

  // setup channel for communication and wait group for thread management
  job := make(chan string)
  wg := new(sync.WaitGroup)

  // start worker routines
  for w := 0; w < runtime.NumCPU(); w++ {
    wg.Add(1)
    go worker(w,job,wg)
  }

  // read from file and send to workers
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    domain := scanner.Text()
    job <- domain
  }
  // close channel so workers know to finish up
  close(job)

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  // wait for all workers to finish up
  wg.Wait()
}
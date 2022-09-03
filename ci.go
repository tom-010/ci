package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// git rev-parse HEAD

func gitHash() string {
  stdout, _, err := run("git", "rev-parse", "HEAD")
  if err != nil {
    fmt.Println(err)
    return ""
  }
  
  return stdout 
}


func run(command string, args ...string) (string, string, error) {
  cmd := exec.Command(command, args...)
  var outb, errb bytes.Buffer
  cmd.Stdout = &outb
  cmd.Stderr = &errb
  err := cmd.Run()
  if err != nil {
    return "", "", err
  }
  return outb.String(), errb.String(), nil
}

func getName() string {
  dt := time.Now()
  res := dt.Format("2006-01-02_15:04:05")
  h := gitHash()
  if h != "" {
    res += "_" + h
  }
  return res
}

func main() {
  fmt.Println(getName())
  stdout, stderr, err := run("./pipeline/01_abc.sh")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("out: %s\nerr: %s\n", stdout, stderr)


  fmt.Println("ci")
  files, err := ioutil.ReadDir("./pipeline/")
  if err != nil {
    log.Fatal(err)
  }
  for _, file := range files {
    fmt.Println(file.Name(), file.IsDir());
    parts := strings.Split(file.Name(), "_")
    if len(parts) > 0 {
      group, err := strconv.Atoi(parts[0])
      if err != nil {
        fmt.Println(err)
      } else {
        fmt.Printf("group: %d\n", group)
      }
    }
  }
  dt := time.Now()
  fmt.Println(dt.Format("2006-02-01_15:04:05"))
}

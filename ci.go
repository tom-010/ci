package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
  cmd := exec.Command("./pipeline/01_abc.sh")
  var outb, errb bytes.Buffer
  cmd.Stdout = &outb
  cmd.Stderr = &errb
  err := cmd.Run()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("out: %s\nerr: %s\n", outb.String(), errb.String())


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
}

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"

	"os"
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
  name := "pipeline.adoc" // getName()
  pipeline := "./pipeline/"
  files, err := filesInFolder(pipeline)
  if err != nil {
    log.Fatal(err)
  }
  report := "# ci\n:toc:\n\n"
  globalOk := true
  for _, file := range files {
    stdout, stderr, err := run(pipeline + file)
    ok := (err == nil)
    globalOk = globalOk && ok
    report += "## " + file
    if !ok {
      report += " FAILED"
    }
    report += "\n"

    stdout = strings.TrimSpace(stdout)
    stderr = strings.TrimSpace(stderr)
    if stdout != "" {
      report += "\n[source]\n----\n"
      report += stdout
      report += "\n----\n\n"
    }
    if stderr != "" {
      report += "\n[source]\n----\n"
      report += stderr
      report += "\n----\n\n"
    }
    err = WriteFile(name, report)
    if err != nil {
      log.Fatal(err)
    }
  }
  
  if !globalOk {
    log.Fatalf("Pipeline failed")
  }
}

func filesInFolder(path string) ([]string, error) {
  res := []string{}
  files, err := ioutil.ReadDir(path)
  if err != nil {
    return res, err
  }

  for _, file := range files {
    if !file.IsDir() {
      res = append(res, file.Name())
    }
  }

  sort.Strings(res)
  return res, nil
}

func WriteFile(path, content string) error {
  return os.WriteFile(path, []byte(content), 0644)
}

func Lookup() {
  stdout, stderr, err := run("./pipeline/01_abc.sh")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("out: %s\nerr: %s\n", stdout, stderr)

  fmt.Print()

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
  fmt.Println(dt.Format("2006-01-02_15:04:05"))
}

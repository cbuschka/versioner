package main

import (
    "fmt"
    "log"
    "os/exec"
    "strings"
)

func main() {
    tags, err := ListTags()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("The output is %v\n", tags)
}

func ListTags() ([]string, error) {
    out, err := exec.Command("git", "tag", "--list").Output()
    if err != nil {
        return nil, err
    }
    outText := string(out)
    tags := strings.Split(outText,"\n")
    return tags, nil
}

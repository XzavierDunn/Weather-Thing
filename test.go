package main

import (
    "os/exec"
    "log"
    "fmt"
)

func main() {
    cmd := exec.Command("python3 .pyText.py")
    stdout, err := cmd.Output()
    if err != nil {
        log.Fatal(err)
        }

    fmt.Println(stdout)
}

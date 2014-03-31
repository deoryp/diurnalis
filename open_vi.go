package main

import (
  "bytes"
  "log"
  "fmt"
  "strings"
  "os"
  "os/exec"
)

func uuidgen() string {
    out, err := exec.Command("uuidgen").Output()
    if err != nil {
        log.Fatal(err)
    }
    return strings.TrimSpace(string(out))
}

func main() {

    var filename bytes.Buffer
    uuid := uuidgen()

    filename.WriteString("/tmp/")
    filename.WriteString(uuid)
    
    //fmt.Println(filename.String())
    
    cmd := exec.Command("/usr/bin/vim", filename.String())
    
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    
    file, err := os.Open(filename.String()) // For read access.
    if err != nil {
    	log.Fatal(err)
    }
    
    data := make([]byte, 2048)
    count, err := file.Read(data)
    if err != nil {
    	log.Fatal(err)
    }
    fmt.Printf("read %d bytes: %q\n", count, data[:count])
    
    os.Remove(filename.String())
    
}

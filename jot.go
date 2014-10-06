package main

import "os"
import "os/exec"
import "fmt"
import "path/filepath"

import "github.com/wgreenberg/jot/io"
import "github.com/wgreenberg/jot/config"

func ensureJotDirExists() {
    if _, err := os.Stat(config.GetJotDir()); os.IsNotExist(err) {
        os.MkdirAll(config.GetJotDir(), 0777)
    }
}

// Opens the editor to a new jot
func CreateNewJot() {
    newJotName := io.UniqueName()
    newJotPath := filepath.Join(config.GetJotDir(), newJotName)
    cmd := exec.Command(config.GetJotEditor(), newJotPath)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        fmt.Println("Error creating the jot:", err)
    }
}

func main() {
    // setup our jot dir if it doesn't exist
    ensureJotDirExists()

    CreateNewJot()
}

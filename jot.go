package main

import "os"
import "os/exec"
import "fmt"
import "path/filepath"

import "github.com/wgreenberg/jot/io"
import "github.com/wgreenberg/jot/config"

type Action int
const (
    ACTION_UNKNOWN = iota
    ACTION_OPEN Action = iota
    ACTION_LOCK Action = iota
    ACTION_UNLOCK Action = iota
    ACTION_GREP Action = iota
    ACTION_LIST Action = iota
)

func ensureJotDirExists() {
    if _, err := os.Stat(config.GetJotDir()); os.IsNotExist(err) {
        os.MkdirAll(config.GetJotDir(), 0777)
    }
}

func GetJotAction() Action {
    var command string
    if len(os.Args) == 1 {
        command = "open"
    } else {
        command = os.Args[1]
    }

    if (command == "open") {
        return ACTION_OPEN
    } else if (command == "grep") {
        return ACTION_GREP
    } else if (command == "lock") {
        return ACTION_LOCK
    } else if (command == "unlock") {
        return ACTION_UNLOCK
    } else if (command == "ls") {
        return ACTION_LIST
    }
    return ACTION_OPEN
}

// Opens the editor to a new jot
func CreateNewJot(name string) {
    if name == "" {
        name = io.UniqueName()
    }
    newJotPath := filepath.Join(config.GetJotDir(), name)
    cmd := exec.Command(config.GetJotEditor(), newJotPath)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        fmt.Printf("Error creating the jot: %s\n", err)
    } else {
        fmt.Printf("Successfully made jot %s\n", name)
    }
}

func main() {
    // setup our jot dir if it doesn't exist
    ensureJotDirExists()
    action := GetJotAction()

    switch {
    case ACTION_OPEN == action:
        if len(os.Args) == 1 {
            CreateNewJot("")
        } else {
            CreateNewJot(os.Args[1])
        }
    case ACTION_LOCK == action:
    case ACTION_UNLOCK == action:
    case ACTION_GREP == action:
        jots := io.ReadAllJots()
        if len(os.Args) == 2 {
            fmt.Printf("Error, grep needs a pattern to search against\n")
            return
        }
        for _, jot := range jots {
            if len(jot.Find(os.Args[2])) > 0 {
                for _, line := range jot.Find(os.Args[2]) {
                    fmt.Printf("%s\t%s\n", jot.Name(), line)
                }
            }
        }
    case ACTION_LIST == action:
        jots := io.ReadAllJots()
        for _, jot := range jots {
            fmt.Printf("%s\t%s\n", jot.Name(), jot.Title())
        }
    case ACTION_UNKNOWN == action:
        fmt.Printf("Error, don't know command: %s\n", os.Args)
    }
}

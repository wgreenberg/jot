package io

import "io/ioutil"
import "path/filepath"
import "crypto/sha1"
import "encoding/hex"
import "os"

import "github.com/wgreenberg/jot/jotlib"
import "github.com/wgreenberg/jot/config"

import "code.google.com/p/go-uuid/uuid"

// Generates a random jot name
func UniqueName() (name string) {
    randomUUID := uuid.New()
    hasher := sha1.New()
    hasher.Write([]byte(randomUUID))
    return hex.EncodeToString(hasher.Sum(nil))[:16]
}

func WriteJot (jot jotlib.Jot) (err bool) {
    jotFileName := filepath.Join(config.GetJotDir(), jot.Name())
    ioerr := ioutil.WriteFile(jotFileName, []byte(jot.Body()), 0777)
    if ioerr != nil {
        return true
    }

    if jot.IsEncrypted {
        lockFileName := filepath.Join(config.GetJotDir(), "." + jot.Name() + ".lock")
        ioerr := ioutil.WriteFile(lockFileName, []byte(jot.LockData()), 0777)
        if ioerr != nil {
            return true
        }
    }

    return false
}

func ReadJot (jotName string) (jot jotlib.Jot, err bool) {
    jotFileName := filepath.Join(config.GetJotDir(), jotName)
    lockFileName := filepath.Join(config.GetJotDir(), "." + jotName + ".lock")
    buffer, ioerr := ioutil.ReadFile(jotFileName)
    if ioerr != nil {
        return jotlib.Jot{}, true
    }

    newJot := jotlib.Jot{}
    newJot.SetName(jotName)
    newJot.SetBody(string(buffer))

    // if the lockfile exists (i.e. the jot is encrypted)
    if _, err := os.Stat(lockFileName); err == nil {
        buffer, ioerr := ioutil.ReadFile(lockFileName)
        if ioerr != nil {
            return newJot, true
        }
        newJot.IsEncrypted = true
        newJot.SetLockData(string(buffer))
    }

    return newJot, false
}

func ReadAllJots() (jots []jotlib.Jot) {
    jotFiles, _ := ioutil.ReadDir(config.GetJotDir())
    for _, jotFile := range jotFiles {
        newJot, err := ReadJot(jotFile.Name())
        if !err {
            jots = append(jots, newJot)
        }
    }
    return jots
}

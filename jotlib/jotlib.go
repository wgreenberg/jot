package jotlib

import "strings"

import "github.com/wgreenberg/jot/jotlib/crypto"

type Jottable interface {
    Body() string
    SetBody(string)
    Name() string
    SetName(string)
    Title() string
    Encrypt(key string) bool
    Decrypt(key string) bool
    Serialize()
}

type Jot struct {
    name string
    lines []string
    nonce [24]byte
    IsEncrypted bool
}

func (j* Jot) SetBody(newBody string) {
    j.lines = strings.Split(newBody, "\n")
}

func (j Jot) Body() string {
    return strings.Join(j.lines, "\n")
}

func (j Jot) Title() string {
    return j.lines[0]
}

func (j* Jot) SetName(newName string) {
    j.name = newName
}

func (j Jot) Name() string {
    return j.name
}

func (j Jot) LockData() string {
    return string(j.nonce[:24])
}

func (j Jot) Find(pattern string) (results []string) {
    for _, line := range j.lines {
        if strings.LastIndex(line, pattern) >= 0 {
            results = append(results, line)
        }
    }
    return results
}

func (j* Jot) SetLockData(data string) {
    var nonce [24]byte

    // get the key as a 32 byte array
    copy(nonce[:], data)
    j.nonce = nonce
}

func (j* Jot) Encrypt(key string) (err bool) {
    if j.IsEncrypted {
        return true
    }

    cryptotext, nonce, err := crypto.EncryptMessage(key, j.Body())
    if !err {
        j.SetBody(cryptotext)
        j.nonce = nonce
        j.IsEncrypted = true
        return false
    }
    return true
}

func (j* Jot) Decrypt(key string) (err bool) {
    plaintext, err := crypto.DecryptMessage(key, j.nonce, j.Body())
    if !err {
        j.SetBody(plaintext)
        j.IsEncrypted = false
        return false
    }
    return true
}

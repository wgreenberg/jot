package jotlib

import "strings"
import "crypto/rand"

import "code.google.com/p/go.crypto/nacl/secretbox"

type Jottable interface {
    Body() string
    SetBody(string)
    Name() string
    SetName(string)
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

func (j* Jot) SetName(newName string) {
    j.name = newName
}

func (j Jot) Name() string {
    return j.name
}

func (j Jot) LockData() string {
    return string(j.nonce[:24])
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

    cryptotext, nonce, err := encrypt(key, j.Body())
    if !err {
        j.SetBody(cryptotext)
        j.nonce = nonce
        j.IsEncrypted = true
        return false
    }
    return true
}

func (j* Jot) Decrypt(key string) (err bool) {
    plaintext, err := decrypt(key, j.nonce, j.Body())
    if !err {
        j.SetBody(plaintext)
        j.IsEncrypted = false
        return false
    }
    return true
}

func encrypt(key string, message string) (out string, nonce [24]byte, err bool) {
    var rawOut []byte
    var rawKey [32]byte

    // get the key as a 32 byte array
    copy(rawKey[:], key)

    // convert message to binary blob
    rawMessage := []byte(message)

    // randomize the nonce
    rand.Reader.Read(nonce[:24])

    rawOut = secretbox.Seal(rawOut[:0], rawMessage, &nonce, &rawKey)

    out = string(rawOut)

    return out, nonce, false
}

func decrypt(key string, nonce [24]byte, message string) (plaintext string, err bool) {
    var rawOut []byte
    var rawKey [32]byte

    // get the key as a 32 byte array
    copy(rawKey[:], key)

    // convert message to binary blob
    rawMessage := []byte(message)

    var ok bool
    rawOut, ok = secretbox.Open(rawOut[:0], rawMessage, &nonce, &rawKey)
    if !ok {
        return "", true
    }

    plaintext = string(rawOut)

    return plaintext, false
}

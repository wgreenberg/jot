package crypto

import "crypto/rand"

import "code.google.com/p/go.crypto/nacl/secretbox"

func EncryptMessage(key string, message string) (out string, nonce [24]byte, err bool) {
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

func DecryptMessage(key string, nonce [24]byte, message string) (plaintext string, err bool) {
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

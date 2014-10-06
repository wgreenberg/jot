package main

import "testing"
import "github.com/wgreenberg/jot/jotlib"

func TestJotName(t *testing.T) {
    testJot := jotlib.Jot{} 
    testName := "Foo"
    if testJot.SetName(testName); testJot.Name() != testName {
        t.Errorf("The name's all wrong. Expected %s, got %s", testName, testJot.Name())
    }
}

func TestJotBody(t *testing.T) {
    testJot := jotlib.Jot{} 
    testBody := `This
    is
    a
    test`
    if testJot.SetBody(testBody); testJot.Body() != testBody {
        t.Errorf("The body's all wrong. Expected %s, got %s", testBody, testJot.Body())
    }
}

func TestJotEncryption(t *testing.T) {
    testJot := jotlib.Jot{}
    originalBody := "Hello"
    key := "foo123456789"
    testJot.SetBody(originalBody)
    if err := testJot.Encrypt(key); err {
        t.Errorf("Encryption failed")
    } 
    
    if testJot.Body() == originalBody {
        t.Errorf("Cryptotext is same as plaintext!!!")
    }

    if err := testJot.Decrypt(key); err {
        t.Errorf("Decryption failed")
    }
    
    if testJot.Body() != originalBody {
        t.Errorf("Decyphered message is different than original message!!! Expected %s, got %s", originalBody, testJot.Body())
    }
}

func TestJotEncryptionIdempotency(t *testing.T) {
    testJot := jotlib.Jot{}
    testJot.SetBody("sup")
    correctKey := "super secret"
    if err := testJot.Encrypt(correctKey); err {
        t.Errorf("Encryption failed")
    }

    if err := testJot.Encrypt("super duper secret"); !err {
        t.Errorf("We shouldn't be able to encrypt an already encrypted file!")
    }

    if err := testJot.Decrypt("definitely not the key"); !err {
        t.Errorf("Wrong key, this should fail")
    }

    if err := testJot.Encrypt("super duper secret"); !err {
        t.Errorf("File is still encrypted, so we still can't encrypt it again!")
    }

    if err := testJot.Decrypt(correctKey); err {
        t.Errorf("Decryption failed")
    }

    if err := testJot.Encrypt("super duper secret"); err {
        t.Errorf("Encrypt away!")
    }
}

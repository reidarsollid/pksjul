package main

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestEncrypt(t *testing.T) {
    text := "text to encrypt"
    expected := "Krwho Fdhvdu"
    actual := Encrypt(text)
    assert.Equal(t, expected , actual )
}

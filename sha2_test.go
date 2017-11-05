package main
import "testing"
import "bytes"
import (
	"crypto/sha256"
)


func TestShort(t *testing.T) {
	sum := sha256.Sum256([]byte("hello"))
	mySum := wikiSha256([]byte("hello"))
	if bytes.Compare(mySum[:], sum[:]) != 0 {
		t.Fatal("invalid sha result")
	}
}

func TestLong(t *testing.T) {
	sum := sha256.Sum256([]byte("my name is bob. this is laura palmer.my name is bob. this is laura palmer."))
	mySum := wikiSha256([]byte("my name is bob. this is laura palmer.my name is bob. this is laura palmer."))
	if bytes.Compare(mySum[:], sum[:]) != 0 {
		t.Fatal("invalid sha result")
	}
}

func TestLong2(t *testing.T) {
	sum := sha256.Sum256([]byte("my name is bob. this is laura palmer.my name is bob. this"))
	mySum := wikiSha256([]byte("my name is bob. this is laura palmer.my name is bob. this"))
	if bytes.Compare(mySum[:], sum[:]) != 0 {
		t.Fatal("invalid sha result")
	}
}
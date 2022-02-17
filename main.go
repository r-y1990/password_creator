package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type PasswordCreater interface {
	Create()
}

type LowerAlphabetType struct {
	UsableStr []string
}

func NewLowerAlphabet() *LowerAlphabetType {
	la := new(LowerAlphabetType)
	la.UsableStr = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	return la
}

type UpperAlphabetType struct {
	UsableStr []string
}

func NewUpperAlphabet() *UpperAlphabetType {
	ua := new(UpperAlphabetType)
	ua.UsableStr = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	return ua
}

type SignType struct {
	UsableStr []string
}

func NewSign() *SignType {
	s := new(SignType)
	s.UsableStr = []string{"!", "#", "$", "%", "&", "'", "(", ")", "-", "^", "@", "[", ";", ":", "]", ",", ".", "/", "=", "~", "|", "`", "{", "+", "*", "}", "<", ">", "?", "_"}
	return s
}

type NumberType struct {
	UsableStr []string
}

func NewNumber() *NumberType {
	n := new(NumberType)
	n.UsableStr = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	return n
}

func main() {
	// TODO 設定ファイルに外だし
	// passwordLength := 12
	// partPasswordLength := 4

	//ps := NewPasswordString()

	n, err := rand.Int(rand.Reader, big.NewInt(4))
	if err != nil {
		return
	}
	fmt.Println(n)
}

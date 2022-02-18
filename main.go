package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

type PasswordCreater interface {
	Create() string
}

type LowerAlphabetType struct {
	UsableStr []string
}

func NewLowerAlphabet() *LowerAlphabetType {
	la := new(LowerAlphabetType)
	la.UsableStr = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	return la
}

func (s LowerAlphabetType) Create() string {
	return PickPasswordStr(s.UsableStr)
}

type UpperAlphabetType struct {
	UsableStr []string
}

func NewUpperAlphabet() *UpperAlphabetType {
	ua := new(UpperAlphabetType)
	ua.UsableStr = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	return ua
}

func (s UpperAlphabetType) Create() string {
	return PickPasswordStr(s.UsableStr)
}

type SignType struct {
	UsableStr []string
}

func NewSign() *SignType {
	s := new(SignType)
	s.UsableStr = []string{"!", "#", "$", "%", "&", "'", "(", ")", "-", "^", "@", "[", ";", ":", "]", ",", ".", "/", "=", "|", "`", "{", "+", "*", "}", "<", ">", "?", "_"}
	return s
}

func (s SignType) Create() string {
	return PickPasswordStr(s.UsableStr)
}

type NumberType struct {
	UsableStr []string
}

func NewNumber() *NumberType {
	n := new(NumberType)
	n.UsableStr = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	return n
}

func (s NumberType) Create() string {
	return PickPasswordStr(s.UsableStr)
}

func PickPasswordStr(s []string) string {
	strLength := int64(len(s))
	n, _ := rand.Int(rand.Reader, big.NewInt(strLength))
	return s[n.Int64()]
}

var (
	length   = flag.Int("l", 12, "Password Length")
	isNoSign = flag.Bool("ns", false, "Password not include sign.")
	isNoNum  = flag.Bool("nn", false, "Password not include number.")
)

func main() {
	flag.Parse()

	// TODO 設定外だし
	passwordLength := length
	creators := []PasswordCreater{NewLowerAlphabet(), NewUpperAlphabet()}

	if !*isNoSign {
		creators = append(creators, NewSign())
	}

	if !*isNoNum {
		creators = append(creators, NewNumber())
	}

	tn := int64(len(creators))

	passwordStr := ""

	for i := 0; i < *passwordLength; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(tn))
		passwordStr = passwordStr + creators[n.Int64()].Create()
	}

	fmt.Println(passwordStr)
	fmt.Println(*length)
	fmt.Println(*isNoNum)
	fmt.Println(*isNoSign)
}

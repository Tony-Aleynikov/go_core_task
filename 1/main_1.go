package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

const salt = "go-2024"

func main() {
	var (
		a int       = 1
		b int       = 02
		c int       = 0x3
		d float64   = 4.0
		e string    = "5"
		f bool      = true
		g complex64 = 6i
	)

	fmt.Println(getType(a))
	fmt.Println(getType(b))
	fmt.Println(getType(c))
	fmt.Println(getType(d))
	fmt.Println(getType(e))
	fmt.Println(getType(f))
	fmt.Println(getType(g))

	s := merge(a, b, c, d, e, f, g)

	r := []rune(s)
	fmt.Println(hash(r))
}

func getType(v any) string {
	return fmt.Sprintf("%T", v)
}

func merge(values ...any) string {
	var b strings.Builder
	for _, v := range values {
		b.WriteString(fmt.Sprintf("%v", v))
	}
	return b.String()
}

func hash(r []rune) [32]byte {
	half := len(r) / 2
	first := []byte(string(r[:half]))
	second := []byte(string(r[half:]))

	salted := make([]byte, 0, len(first)+len(salt)+len(second))
	salted = append(salted, first...)
	salted = append(salted, salt...)
	salted = append(salted, second...)

	return sha256.Sum256(salted)
}

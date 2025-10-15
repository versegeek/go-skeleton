package main

import (
	"fmt"

	"github.com/versegeek/go-skeleton/pkg/version"
)

func main() {
	v := version.Get()
	fmt.Println(v.ToString())
}

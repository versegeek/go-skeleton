package main

import (
	"github.com/versegeek/go-skeleton/internal/server"
	"github.com/versegeek/toolkit/pkg/stack"
)

func main() {
	st := stack.New()
	defer st.MustClose()

	serverProvider := server.New()
	st.MustInit(serverProvider)

	st.MustRun()
}

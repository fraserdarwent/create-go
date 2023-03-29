package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log = Logger()
	router = Router()
	code := m.Run()
	os.Exit(code)
}

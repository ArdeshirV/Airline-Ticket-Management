package main

import (
	"os"
	"testing"

	"github.com/the-go-dragons/final-project/pkg/test"
)

// var a *handlers.Application

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	test.SetupWithData()  // Load fake data
	defer test.Teardown() // Clean fake data
	// a = handlers.NewApplication(dbc)
	return m.Run()
}

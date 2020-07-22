package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestHello(t *testing.T) {
	expected := "Hello World"

	got := TestHello()

	assert.Equal(t, expected, got)
}

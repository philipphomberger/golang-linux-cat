package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCat(t *testing.T) {
	catTest := cat("test.txt")
	assert.Contains(t, catTest, "Hello World\nWhat")

}

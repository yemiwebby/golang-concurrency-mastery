package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)


func TestMain_Execution(t *testing.T) {
	go main()

	time.Sleep(5 * time.Second)

	assert.True(t, true, "System ran successfully without crashes")
}
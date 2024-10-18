package test

import (
	"go_restful_api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializeService(true)

	assert.NotNil(t, err)
	assert.Nil(t, simpleService)
}

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitializeService(false)

	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}

func TestCleanupFunc(t *testing.T) {
	connection, cleanup := simple.InitializeCleanUpFunction("Connection 1")

	assert.NotNil(t, connection)

	cleanup()
}

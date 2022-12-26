package main

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	record := httptest.NewRecorder()
	result := httptest.NewRequest("GET", "/hello", nil)
	helloHandler(record, result)
	body := record.Result().Body
	data, _ := io.ReadAll(body)

	assert.Equal(t, string(data), "hello")

}

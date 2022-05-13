package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	fmt.Println("Running E2E test for health check endpoint")
	resp, err := http.Get("localhost:8080/api/health")
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, 200, resp.StatusCode)

}

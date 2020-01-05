package rest_errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("message", errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, `message`, err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "database error", err.Causes[0])

	errBytes, _ := json.Marshal(err)
	fmt.Println(string(errBytes))
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, `message`, err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, `message`, err.Message)
	assert.EqualValues(t, "not_found", err.Error)
}

func TestNewError(t *testing.T) {
	err:= NewError("esto es un error")
	assert.NotNil(t, err)
	assert.EqualValues(t, "esto es un error", err.Error())
}

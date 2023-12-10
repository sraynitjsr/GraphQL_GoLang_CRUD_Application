package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sraynitjsr/db"
	"github.com/stretchr/testify/assert"
)

type MockDB struct{}

// Mock Implementation
func (m *MockDB) GetNodes() ([]Node, error) {
	return []Node{{ID: 1, Name: "Node1"}, {ID: 2, Name: "Node2"}}, nil
}

// Mock Implementation
func (m *MockDB) GetNodeByID(id int) (*Node, error) {
	if id == 1 {
		return &Node{ID: 1, Name: "Node1"}, nil
	}
	return nil, errors.New("Node not found")
}

func TestGetNodes(t *testing.T) {
	// Create a mock database
	mockDB := &MockDB{}

	// Create a Gin router
	r := gin.Default()

	// Define the route
	r.GET("/nodes", GetNodes(mockDB))

	// Create an HTTP request
	req, err := http.NewRequest("GET", "/nodes", nil)
	assert.NoError(t, err)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Decode the response body
	var nodes []Node
	err = json.Unmarshal(w.Body.Bytes(), &nodes)
	assert.NoError(t, err)

	// Check the response body
	expectedNodes := []Node{{ID: 1, Name: "Node1"}, {ID: 2, Name: "Node2"}}
	assert.Equal(t, expectedNodes, nodes)
}

func TestGetNodeByID(t *testing.T) {
	// Create a mock database
	mockDB := &MockDB{}

	// Create a Gin router
	r := gin.Default()

	// Define the route
	r.GET("/nodes/:id", GetNodeByID(mockDB))

	// Create an HTTP request with a valid ID
	req, err := http.NewRequest("GET", "/nodes/1", nil)
	assert.NoError(t, err)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Decode the response body
	var node Node
	err = json.Unmarshal(w.Body.Bytes(), &node)
	assert.NoError(t, err)

	// Check the response body
	expectedNode := Node{ID: 1, Name: "Node1"}
	assert.Equal(t, expectedNode, node)

	// Create an HTTP request with an invalid ID
	req, err = http.NewRequest("GET", "/nodes/999", nil)
	assert.NoError(t, err)

	// Create a response recorder
	w = httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the response status code for a not found scenario
	assert.Equal(t, http.StatusNotFound, w.Code)
}

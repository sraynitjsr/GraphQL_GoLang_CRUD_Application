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


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

func (m *MockDB) GetNodes() ([]Node, error) {
	return []Node{{ID: 1, Name: "Node1"}, {ID: 2, Name: "Node2"}}, nil
}

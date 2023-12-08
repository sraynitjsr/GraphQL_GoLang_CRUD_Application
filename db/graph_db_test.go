package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/sraynitjsr/model"
)

// MockNeo4jDriver is a mock implementation of the neo4j.Driver interface
type MockNeo4jDriver struct {
	mock.Mock
}

func (m *MockNeo4jDriver) NewSession(config neo4j.SessionConfig) neo4j.Session {
	args := m.Called(config)
	return args.Get(0).(neo4j.Session)
}

func (m *MockNeo4jDriver) Close() error {
	args := m.Called()
	return args.Error(0)
}

func TestGetNodes(t *testing.T) {
	// Create a mock Neo4j driver
	mockDriver := new(MockNeo4jDriver)

	// Create a GraphDatabase with the mock driver
	graphDB := &GraphDatabase{dbDriver: mockDriver}

	// Create a mock session
	mockSession := new(MockNeo4jSession)
	mockDriver.On("NewSession", mock.Anything).Return(mockSession)

	// Set up expectations for the session method calls
	mockSession.On("Run", mock.Anything, mock.Anything).Return(new(MockNeo4jResult), nil)
	mockSession.On("Close").Return(nil)

	// Call the method to be tested
	_, err := graphDB.GetNodes()

	// Assert that expectations were met
	mockDriver.AssertExpectations(t)
	mockSession.AssertExpectations(t)

	// Assert the result of the method
	assert.NoError(t, err)
}

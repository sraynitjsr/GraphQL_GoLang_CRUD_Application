package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/sraynitjsr/model"
)

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

type MockNeo4jSession struct {
	mock.Mock
}

func (m *MockNeo4jSession) Run(cypher string, params map[string]interface{}) (neo4j.Result, error) {
	args := m.Called(cypher, params)
	return args.Get(0).(neo4j.Result), args.Error(1)
}

func (m *MockNeo4jSession) Close() error {
	args := m.Called()
	return args.Error(0)
}

func TestGetNodes(t *testing.T) {
	mockDriver := new(MockNeo4jDriver)
	graphDB := &GraphDatabase{dbDriver: mockDriver}
	mockSession := new(MockNeo4jSession)
	mockDriver.On("NewSession", mock.Anything).Return(mockSession)
	mockSession.On("Run", "MATCH (n:Node) RETURN n", nil).
		Return(new(MockNeo4jResult), nil)
	mockSession.On("Close").Return(nil)

	_, err := graphDB.GetNodes()

	mockDriver.AssertExpectations(t)
	mockSession.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestGetNodeByID(t *testing.T) {
	mockDriver := new(MockNeo4jDriver)
	graphDB := &GraphDatabase{dbDriver: mockDriver}
	mockSession := new(MockNeo4jSession)
	mockDriver.On("NewSession", mock.Anything).Return(mockSession)
	expectedID := 123
	mockSession.On("Run", "MATCH (n:Node) WHERE ID(n) = $id RETURN n", map[string]interface{}{"id": expectedID}).
		Return(new(MockNeo4jResult), nil)
	mockSession.On("Close").Return(nil)

	_, err := graphDB.GetNodeByID(expectedID)

	mockDriver.AssertExpectations(t)
	mockSession.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestCreateNode(t *testing.T) {
	mockDriver := new(MockNeo4jDriver)
	graphDB := &GraphDatabase{dbDriver: mockDriver}
	mockSession := new(MockNeo4jSession)
	mockDriver.On("NewSession", mock.Anything).Return(mockSession)
	mockSession.On("Run", "CREATE (n:Node {name: $name}) RETURN ID(n) as id", map[string]interface{}{"name": "test"}).
		Return(new(MockNeo4jResult), nil)
	mockSession.On("Close").Return(nil)

	_, err := graphDB.CreateNode("test")

	mockDriver.AssertExpectations(t)
	mockSession.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestUpdateNode(t *testing.T) {
	mockDriver := new(MockNeo4jDriver)
	graphDB := &GraphDatabase{dbDriver: mockDriver}
	mockSession := new(MockNeo4jSession)
	mockDriver.On("NewSession", mock.Anything).Return(mockSession)
	mockSession.On("Run", "MATCH (n:Node) WHERE ID(n) = $id SET n.name = $name", map[string]interface{}{"id": 123, "name": "updated"}).
		Return(new(MockNeo4jResult), nil)
	mockSession.On("Close").Return(nil)

	err := graphDB.UpdateNode(123, "updated")

	mockDriver.AssertExpectations(t)
	mockSession.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestDeleteNode(t *testing.T) {
	mockDriver := new(MockNeo4jDriver)
	graphDB := &GraphDatabase{dbDriver: mockDriver}
	mockSession := new(MockNeo4jSession)
	mockDriver.On("NewSession", mock.Anything).Return(mockSession)
	mockSession.On("Run", "MATCH (n:Node) WHERE ID(n) = $id DELETE n", map[string]interface{}{"id": 123}).
		Return(new(MockNeo4jResult), nil)
	mockSession.On("Close").Return(nil)

	err := graphDB.DeleteNode(123)

	mockDriver.AssertExpectations(t)
	mockSession.AssertExpectations(t)
	assert.NoError(t, err)
}

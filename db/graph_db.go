package db

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/sraynitjsr/model"
)

type GraphDatabase struct {
	dbDriver neo4j.Driver
}

var (
	dbURL      = "bolt://localhost:7687"
	dbUsername = "neo4j"
	dbPassword = "my_password"
)

func NewGraphDB() (Database, error) {
	driver, err := neo4j.NewDriver(dbURL, neo4j.BasicAuth(dbUsername, dbPassword, ""))
	if err != nil {
		return nil, fmt.Errorf("failed to create Neo4j driver: %v", err)
	}

	return &GraphDatabase{dbDriver: driver}, nil
}

func (gd *GraphDatabase) GetNodes() ([]model.Node, error) {
	session := gd.dbDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	result, err := session.Run("MATCH (n:Node) RETURN n", nil)
	if err != nil {
		return nil, err
	}

	var nodes []model.Node
	for result.Next() {
		record := result.Record()
		nodeValue, found := record.Get("n")
		if !found {
			continue
		}
		node, ok := nodeValue.(neo4j.Node)
		if !ok {
			continue
		}

		props := node.Props

		nodes = append(nodes, model.Node{
			ID:   int(node.Id),
			Name: props["name"].(string),
		})
	}

	return nodes, nil
}

func (gd *GraphDatabase) GetNodeByID(id int) (*model.Node, error) {
	session := gd.dbDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	result, err := session.Run("MATCH (n:Node) WHERE ID(n) = $id RETURN n", map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	if result.Next() {
		record := result.Record()
		nodeValue, found := record.Get("n")
		if !found {
			return nil, nil
		}
		node, ok := nodeValue.(neo4j.Node)
		if !ok {
			return nil, nil
		}

		props := node.Props

		return &model.Node{
			ID:   int(node.Id),
			Name: props["name"].(string),
		}, nil
	}

	return nil, nil
}

func (gd *GraphDatabase) CreateNode(name string) (int, error) {
	session := gd.dbDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	result, err := session.Run("CREATE (n:Node {name: $name}) RETURN ID(n) as id", map[string]interface{}{"name": name})
	if err != nil {
		return 0, err
	}

	if result.Next() {
		record := result.Record()
		idValue, found := record.Get("id")
		if !found {
			return 0, fmt.Errorf("failed to retrieve node ID")
		}

		id, ok := idValue.(int64)
		if !ok {
			return 0, fmt.Errorf("failed to convert ID to int64")
		}

		return int(id), nil
	}

	return 0, nil
}

func (gd *GraphDatabase) UpdateNode(id int, name string) error {
	session := gd.dbDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.Run("MATCH (n:Node) WHERE ID(n) = $id SET n.name = $name", map[string]interface{}{"id": id, "name": name})
	return err
}

func (gd *GraphDatabase) DeleteNode(id int) error {
	session := gd.dbDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.Run("MATCH (n:Node) WHERE ID(n) = $id DELETE n", map[string]interface{}{"id": id})
	return err
}

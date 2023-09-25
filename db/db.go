package db

import (
	"github.com/sraynitjsr/model"
)

type Database interface {
	GetNodes() ([]model.Node, error)
	GetNodeByID(id int) (*model.Node, error)
	CreateNode(name string) (int, error)
	UpdateNode(id int, name string) error
	DeleteNode(id int) error
}

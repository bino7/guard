package guard

import (
	"github.com/cayleygraph/cayley"
	_ "github.com/cayleygraph/cayley/graph/bolt"
	"log"
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley/graph"
)

type Store cayley.Handle

var (
	store_path = "guard.db"
	cayleyStore *cayley.Handle
	store Store
)

func initStore() error {
	graph.InitQuadStore("bolt", store_path, nil)
	s, err := cayley.NewGraph("bolt", store_path, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	cayleyStore = s
	store = Store(*s)
	return nil
}

type CayleyData interface {
	Quads() []quad.Quad
}

func (s Store)Insert(d CayleyData) error {
	return s.AddQuadSet(d.Quads())
}

func (s Store)Remove(d CayleyData) error {
	tx := graph.NewTransaction()
	for _, quad := range d.Quads() {
		tx.RemoveQuad(quad)
	}
	return s.ApplyTransaction(tx)
}
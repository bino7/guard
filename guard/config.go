package guard

import (
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/quad"
	"fmt"
)

type config struct {
	CheckFrequency int /*unit second*/
}

func (c *config)Quads() []quad.Quad {
	return []quad.Quad{
		quad.Make("config", "checkFrequency", 30, nil),
	}
}

func loadConfig() *config {
	c := new(config)
	path := cayley.StartPath(cayleyStore, quad.StringToValue("config")).Out("checkFrequency")
	it := path.BuildIterator()
	defer it.Close()
	if !it.Next() {
		fmt.Println(store.Insert(c))
	}

	it = cayley.StartPath(cayleyStore, quad.StringToValue("config")).Out("checkFrequency").BuildIterator()
	defer it.Close()
	if it.Next() {
		c.CheckFrequency = store.NameOf(it.Result()).Native().(int)
	} else {
		fmt.Println("no checkFrequency")
	}
	return c
}

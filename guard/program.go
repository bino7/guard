package guard

import (
	"os"
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/quad"
)

type Program struct {
	Name    string
	process map[int]*os.Process
}
func (pro *Program)Quads() []quad.Quad {
	return []quad.Quad{
		quad.Make(pro.Name, "is", "Program", nil),
	}
}

func newProgram(name string)*Program{
	process:=make(map[int]*os.Process)
	pro:= &Program{name,process}
	return pro
}

func (pro *Program)processStart(p *os.Process){
	pro.process[p.Pid]=p
}

func (pro *Program)processExited(p *os.Process){
	delete(pro.process,p.Pid)
}

func loadPrograms()[]*Program{
	pros:=make([]*Program,0)
	it:=cayley.StartPath(cayleyStore, quad.StringToValue("Program")).In("is").BuildIterator()
	defer it.Close()
	if it.Next() {
		name:=store.NameOf(it.Result()).Native().(string)
		pro:=newProgram(name)
		pros=append(pros,pro)
	}
	return pros
}



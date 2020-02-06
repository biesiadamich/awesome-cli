package manager

import (
	"fmt"

    "awesome/internal/package/node"
    "awesome/internal/package/fetcher"
    "awesome/internal/package/parser"
)

type Manager struct {
    Root          *node.Node
    PWD           *node.Node
    History       []Command
    ValidCommands []Command
}

type Command struct {
	Text string
}

func New() Manager {
	return Manager{
		Root:          nil,
		PWD:           nil,
		History:       []Command{},
		ValidCommands: []Command{},
	}
}

func (m *Manager) Execute(command Command) {
	command.Execute(m)
}

func (m *Manager) SetPWD(node *node.Node) {
	if len(node.GetChildren()) == 0 {
		fecthed, err := fetcher.FetchAwsomeRepo(node.GetReadmeURL())

		if err != nil {
			panic(err)
		}

		temp := parser.ParseIndex(fecthed)

		node.SetChildren(temp.GetChildren())
	}

	m.PWD = node
}

func (m *Manager) GetPWD() *node.Node {
	return m.PWD
}

func (m *Manager) Initialize() {
	fecthed, err := fetcher.FetchAwsomeRootRepo()

	if err != nil {
		panic(err)
	}

	root      := parser.ParseIndex(fecthed)
	root.Name  = "Awesome"
	m.Root     = &root
	m.PWD      = m.Root
}

func (c *Command) Execute(m *Manager) {
	switch c.Text {
	case "ls":
		fmt.Println(m.PWD.GetName())
	default:
		fmt.Println("Invalid command.")
	}
}
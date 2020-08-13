package composite

import (
	"fmt"
	"strings"
)

type Node interface {
	Display(depth int)
}

type Directory struct {
	Name     string
	subNodes []Node
}

func (n *Directory) Add(node Node) {
	n.subNodes = append(n.subNodes, node)
}

func (n *Directory) Display(depth int) {
	fmt.Println(strings.Repeat(" ", depth) + n.Name)
	for _, node := range n.subNodes {
		node.Display(depth + 2)
	}
}

type File struct {
	Name string
}

func (n *File) Display(depth int) {
	fmt.Println(strings.Repeat(" ", depth) + n.Name)
}

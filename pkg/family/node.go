package family

import "github.com/GeekchanskiY/pet-project/pkg/humans"

type Node interface {
	GetParents() []Node
	GetChildren() []Node

	AddParent(Node) error
	AddChild(Node) error

	GetHuman() humans.Human
}

type node struct {
	human    humans.Human
	children []Node
	parents  []Node
}

func NewNode(human humans.Human, children, parents []Node) (Node, error) {
	newNode := &node{
		human:    human,
		children: children,
		parents:  parents,
	}

	for _, child := range children {
		err := child.AddParent(newNode)
		if err != nil {
			return nil, err
		}
	}

	for _, parent := range parents {
		err := parent.AddChild(newNode)
		if err != nil {
			return nil, err
		}
	}

	search := searchHumanInTree(newNode, human)
	if search != nil {
		return nil, ErrFamilyCycle
	}

	return newNode, nil
}

func (n *node) GetParents() []Node {
	return n.parents
}

func (n *node) GetChildren() []Node {
	return n.children
}

func (n *node) GetHuman() humans.Human {
	return n.human
}

func (n *node) AddParent(parent Node) error {
	n.parents = append(n.parents, parent)

	return nil
}

func (n *node) AddChild(child Node) error {
	n.children = append(n.children, child)

	return nil
}

func searchHumanInTree(node Node, human humans.Human) Node {
	res := searchHumanInChildren(node, human)
	if res != nil {
		return res
	}

	res = searchHumanInParents(node, human)
	if res != nil {
		return res
	}

	return nil
}

// searchHumanInParents recursively searches all parents until human found in tree
func searchHumanInParents(node Node, human humans.Human) Node {
	for _, parent := range node.GetParents() {
		if parent.GetHuman() == human {
			return node
		}

		res := searchHumanInParents(parent, human)
		if res != nil {
			return res
		}
	}

	return nil
}

// searchHumanInChildren recursively searches all children until human found in tree
func searchHumanInChildren(node Node, human humans.Human) Node {
	for _, child := range node.GetChildren() {
		if child.GetHuman() == human {
			return node
		}

		res := searchHumanInChildren(child, human)
		if res != nil {
			return res
		}
	}

	return nil
}

package family

import "github.com/GeekchanskiY/pet-project/pkg/humans"

type Node interface {
	GetParents() []Node
	GetChildren() []Node
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

func searchHumanInTree(node Node, human humans.Human) Node {
	for _, child := range node.GetChildren() {
		res := searchHumanInChildren(child, human)

		if res != nil {
			return res
		}
	}

	for _, parent := range node.GetParents() {
		res := searchHumanInParents(parent, human)

		if res != nil {
			return res
		}
	}

	return nil
}

// searchHumanInParents recursively searches all parents until human found in tree
func searchHumanInParents(node Node, human humans.Human) Node {
	for _, parent := range node.GetParents() {
		if node.GetHuman() == human {
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
	for _, child := range node.GetParents() {
		if node.GetHuman() == human {
			return node
		}

		res := searchHumanInChildren(child, human)

		if res != nil {
			return res
		}
	}

	return nil
}

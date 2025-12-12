package family

import "github.com/GeekchanskiY/pet-project/pkg/humans"

type Storage interface {
	GetFamily(human humans.Human) Node
	Append(father, mother, children humans.Human) Node
}

type storage struct {
	coreNodes []Node // Family first generation
}

var stg Storage

// NewStorage returns singleton Storage instance
func NewStorage() Storage {
	if stg == nil {
		stg = &storage{
			coreNodes: make([]Node, 0),
		}
	}

	return stg
}

func (s *storage) GetFamily(human humans.Human) Node {
	for _, node := range s.coreNodes {
		res := searchHumanInTree(node, human)
		if res != nil {
			return res
		}
	}

	return nil
}

func (s *storage) Append(father, mother, children humans.Human) Node {
	return nil
}

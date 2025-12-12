package family_test

import (
	"strconv"
	"testing"

	"github.com/GeekchanskiY/pet-project/pkg/family"
	"github.com/GeekchanskiY/pet-project/pkg/humans"
	"github.com/stretchr/testify/require"
)

func Test_NewNode(t *testing.T) {
	h1 := humanGetter(1)
	h2 := humanGetter(2)
	h3 := humanGetter(3)
	h4 := humanGetter(4)

	t.Run("NewNode: success", func(t *testing.T) {
		node, err := family.NewNode(h1, nil, nil)
		require.NoError(t, err)
		require.NotNil(t, node)

		node2, err := family.NewNode(h2, nil, []family.Node{node})
		require.NoError(t, err)
		require.NotNil(t, node2)

		node3, err := family.NewNode(h3, nil, []family.Node{node2})
		require.NoError(t, err)
		require.NotNil(t, node3)

		node4, err := family.NewNode(h4, nil, []family.Node{node3})
		require.NoError(t, err)
		require.NotNil(t, node4)
	})

	t.Run("NewNode: detect cyclic parents", func(t *testing.T) {
		node, err := family.NewNode(h1, nil, nil)
		require.NoError(t, err)
		require.NotNil(t, node)

		nilNode, err := family.NewNode(h1, nil, []family.Node{node})
		require.Error(t, err)
		require.Nil(t, nilNode)

		node2, err := family.NewNode(h2, nil, []family.Node{node})
		require.NoError(t, err)
		require.NotNil(t, node2)

		nilNode, err = family.NewNode(h1, nil, []family.Node{node2})
		require.Error(t, err)
		require.Nil(t, nilNode)

		nilNode, err = family.NewNode(h1, []family.Node{node2}, nil)
		require.Error(t, err)
		require.Nil(t, nilNode)
	})
}

func Test_GetChildren(t *testing.T) {

	t.Run("GetChildren", func(t *testing.T) {

	})
}

func humanGetter(idx int) humans.Human {
	gender := humans.GenderMale
	if idx%2 == 0 {
		gender = humans.GenderFemale
	}

	return humans.New(
		"testHuman"+strconv.Itoa(idx),
		"testHuman"+strconv.Itoa(idx),
		uint8(10+idx),
		gender,
	)
}

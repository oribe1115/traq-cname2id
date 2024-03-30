package cname2id

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/traPtitech/go-traq"
)

func TestNameToID(t *testing.T) {
	rootID := "aaa-111"

	list := []traq.Channel{
		{
			Id:       "aaa-111",
			ParentId: traq.NullableString{},
			Name:     "a",
			Children: []string{
				"bbb-111",
				"bbb-222",
				"bbb-333",
			},
		},
		{
			Id:       "bbb-111",
			ParentId: *traq.NewNullableString(&rootID),
			Name:     "b",
			Children: []string{},
		},
		{
			Id:       "bbb-222",
			ParentId: *traq.NewNullableString(&rootID),
			Name:     "bb",
			Children: []string{},
		},
		{
			Id:       "bbb-333",
			ParentId: *traq.NewNullableString(&rootID),
			Name:     "bbb",
			Children: []string{},
		},
	}

	t.Run("found", func(t *testing.T) {
		got, err := nameToID("#a/bb", list)
		assert.NoError(t, err)
		assert.Equal(t, "bbb-222", got)
	})

	t.Run("not found", func(t *testing.T) {
		_, err := nameToID("#nan/aa/bb", list)
		assert.Error(t, err)
	})
}

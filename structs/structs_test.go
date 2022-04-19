package structs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/whale-clouds/goutil/dump"
	"github.com/whale-clouds/goutil/structs"
)

func TestTryToMap(t *testing.T) {
	mp, err := structs.TryToMap(nil)
	assert.Empty(t, mp)
	assert.NoError(t, err)

	type User struct {
		Name string
		Age  int
		city string
	}

	u := User{
		Name: "inhere",
		Age:  34,
		city: "somewhere",
	}

	mp, err = structs.TryToMap(u)
	assert.NoError(t, err)
	dump.P(mp)
	assert.Contains(t, mp, "Name")
	assert.Contains(t, mp, "Age")
	assert.NotContains(t, mp, "city")

	mp, err = structs.TryToMap(&u)
	assert.NoError(t, err)

	mp = structs.ToMap(&u)
	assert.NoError(t, err)
	dump.P(mp)

	assert.Panics(t, func() {
		structs.MustToMap("abc")
	})
}

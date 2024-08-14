package di

import (
	"testing"

	"gopkg.in/stretchr/testify.v1/assert"
)

type NewInterface interface {
	NewInterface()
}

type NewInterfaceImpl struct {
	value string
}

func (n *NewInterfaceImpl) NewInterface() {
	//TODO implement me
	panic("implement me")
}

var _ NewInterface = (*NewInterfaceImpl)(nil)

func TestBindInterface(t *testing.T) {
	Bind(new(NewInterface), &NewInterfaceImpl{
		value: "test",
	})

	r := Fetch[NewInterface]()

	assert.NotNil(t, r)
	assert.Equal(t, "test", r.(*NewInterfaceImpl).value)
}

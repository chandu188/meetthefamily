package familytree

import "testing"
import "github.com/stretchr/testify/assert"

func TestMember(t *testing.T) {
	name := "Sathish"
	m := NewMember(name, male)
	assert.Nil(t, m.partner)
	assert.Equal(t, name, m.name)
	assert.Equal(t, male, m.gender)
}

func TestCouple(t *testing.T) {
	s := NewMember("Sandhya", female)
	r := NewMember("Ram", male)
	c := NewCouple(s, r)

	ash := NewMember("Ashwik", male)
	c.addChild(ash)

	assert.NotNil(t, s.partner)
	assert.NotNil(t, r.partner)
	assert.Equal(t, s, c.other(r))
	assert.Equal(t, 1, len(s.partner.children))

	assert.NotNil(t, ash.parent)
	assert.Equal(t, s.name, ash.parent.mother.name)
	assert.NotNil(t, r.name, ash.parent.father.name)
}

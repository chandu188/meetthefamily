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
	s.setPartner(r)

	ash := NewMember("Ashwik", male)
	ash.addParents(s, r)
	s.addChild(ash)

	assert.NotNil(t, s.partner)
	assert.NotNil(t, r.partner)

	assert.Equal(t, 1, len(s.partner.children))

	assert.NotNil(t, ash.parents)
	assert.Equal(t, s.name, ash.mother.name)
	assert.NotNil(t, r.name, ash.father.name)
}

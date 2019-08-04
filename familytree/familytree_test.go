package familytree

import "testing"
import "github.com/stretchr/testify/assert"

func TestKingShaw(t *testing.T) {
	f := &familyTree{}
	f.members = make(map[string]*member)
	initializeFamilyTree(f)

	f.addChild("Chitra", "Aria", female)

	relation := "Maternal-Aunt"
	res := f.GetRelationShip("Lavnya", relation)

	assert.Equal(t, "Aria", res)

	relation = "Siblings"

	res = f.GetRelationShip("Aria", relation)
	assert.Equal(t, "Jnki Ahit", res)

}

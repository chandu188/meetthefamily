package familytree

import (
	"sort"
	"strings"
)

type GENDER string
type RESULT string

const (
	male                  GENDER = "Male"
	female                GENDER = "Female"
	personNotFound               = "PERSON_NOT_FOUND"
	childAdditionSucceded        = "CHILD_ADDITION_SUCCEDED"
	childAdditionFailed          = "CHILD_ADDITION_FAILED"
	none                         = "NONE"
	relationNotFound             = "RELATION_NOT_FOUND"
)

// FamilyTree is an interface to query relation ship, add a child to the tree
type FamilyTree interface {
	AddChild(mother string, child string, g string) string
	GetRelationShip(name string, relation string) string
}

type familyTree struct {
	members map[string]*member
	count   int
}

// NewFamilyTree returns an implementation of FamilyTree interface
func NewFamilyTree() FamilyTree {
	f := &familyTree{
		members: make(map[string]*member),
		count:   0,
	}
	initializeFamilyTree(f)
	return f
}

func (f *familyTree) AddChild(mother string, child string, g string) string {
	mo, ok := f.members[mother]
	if !ok {
		return personNotFound
	}
	if mo.gender == male || mo.partner == nil {
		return childAdditionFailed
	}
	m := f.addToFamily(child, GENDER(g))
	mo.partner.addChild(m)
	return childAdditionSucceded
}

func (f *familyTree) GetRelationShip(name string, relation string) string {
	m, ok := f.members[name]
	if !ok {
		return personNotFound
	}

	relFn, ok := relationFnMap[relation]
	if !ok {
		return relationNotFound
	}

	mems := relFn(m)
	if len(mems) == 0 {
		return none
	}
	sort.Sort(relatives(mems))
	res := make([]string, len(mems))
	for i, v := range mems {
		res[i] = v.name
	}
	return strings.Join(res, " ")
}

func (f *familyTree) addToFamily(name string, g GENDER) *member {
	m := NewMember(name, g)
	f.count++
	m.order = f.count
	f.members[name] = m
	return m
}

func (f *familyTree) addChild(mother string, child string, g GENDER) *member {
	mo, ok := f.members[mother]
	if !ok || mo.partner == nil {
		return nil
	}
	if mo.gender == male {
		return nil
	}
	m := f.addToFamily(child, g)
	mo.addChild(m)
	m.addParents(mo, mo.partner)
	return m
}

func initializeFamilyTree(f *familyTree) {
	ks := f.addToFamily("King Shan", male)
	qa := f.addToFamily("Queen Anga", female)
	qa.setPartner(ks)
	ch := f.addChild(qa.name, "Chit", male)
	am := f.addToFamily("Amba", female)
	am.setPartner(ch)

	f.addChild(qa.name, "Ish", male)

	vi := f.addChild(qa.name, "Vish", male)
	lika := f.addToFamily("Lika", female)
	lika.setPartner(vi)

	aras := f.addChild(qa.name, "Aras", male)
	chitra := f.addToFamily("Chitra", female)
	chitra.setPartner(aras)

	satya := f.addChild(qa.name, "Satya", female)
	vyan := f.addToFamily("Vyan", male)
	satya.setPartner(vyan)

	dritha := f.addChild(am.name, "Dritha", female)
	jaya := f.addToFamily("Jaya", male)
	dritha.setPartner(jaya)

	f.addChild(am.name, "Tritha", female)
	f.addChild(am.name, "Vritha", male)

	f.addChild(dritha.name, "Yodhan", male)

	f.addChild(lika.name, "Vila", female)
	f.addChild(lika.name, "Chika", female)

	jnki := f.addChild(chitra.name, "Jnki", female)
	arit := f.addToFamily("Arit", male)
	jnki.setPartner(arit)

	f.addChild(chitra.name, "Ahit", male)

	f.addChild(jnki.name, "Laki", male)
	f.addChild(jnki.name, "Lavnya", female)

	satvy := f.addToFamily("Satvy", female)
	asva := f.addChild(satya.name, "Asva", male)
	satvy.setPartner(asva)

	krpi := f.addToFamily("Krpi", female)
	vyas := f.addChild(satya.name, "Vyas", male)
	krpi.setPartner(vyas)

	f.addChild(satya.name, "Atya", female)

	f.addChild(satvy.name, "Vasa", male)

	f.addChild(krpi.name, "Kriya", male)
	f.addChild(krpi.name, "Krithi", female)

}

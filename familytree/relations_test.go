package familytree

import "testing"
import "github.com/stretchr/testify/assert"

func TestRelation(t *testing.T) {
	f := NewMember("Ram", male)
	m := NewMember("Sandhya", female)
	m.setPartner(f)
	ash := NewMember("Ashwik", male)
	abhi := NewMember("AbhiRam", male)
	m.addChild(ash)
	m.addChild(abhi)
	ash.addParents(m, f)
	abhi.addParents(m, f)

	assert.Nil(t, father(m))
	assert.Nil(t, mother(m))
	assert.Nil(t, father(f))
	assert.Nil(t, mother(f))
	assert.Equal(t, f.name, father(ash)[0].name)
	assert.Equal(t, m.name, mother(abhi)[0].name)
	assert.Equal(t, 2, len(children(m)))
	assert.Equal(t, 2, len(children(f)))
	assert.Equal(t, 0, len(daughters(m)))
	assert.Equal(t, 1, len(brothers(ash)))
}

func setUp() *familyTree {
	f := &familyTree{}
	f.members = make(map[string]*member)

	k := f.addToFamily("Kotilingam", male)
	s := f.addToFamily("Sattamma", female)
	k.setPartner(s)
	f.addChild(s.name, "Shivaiah", male)
	f.addChild(s.name, "Ammai", female)
	f.addChild(s.name, "Uma", female)

	mallesham := f.addToFamily("Mallesham", male)
	anasuya := f.addToFamily("Anasuya", female)
	anasuya.setPartner(mallesham)
	lakshmi := f.addChild(anasuya.name, "Lakshmi", female)
	anjali := f.addChild(anasuya.name, "Anjali", female)
	rajyam := f.addChild(anasuya.name, "Rajyam", female)
	rama := f.addChild(anasuya.name, "Rama", female)
	kishan := f.addChild(anasuya.name, "Kishan", male)

	jayanth := f.addToFamily("Jayanth", male)
	niranjan := f.addChild(s.name, "Niranjan", male)
	laxman := f.addToFamily("Laxman", male)
	omkar := f.addToFamily("Omkar", male)
	pallavi := f.addToFamily("Pallavi", female)

	lakshmi.setPartner(jayanth)
	anjali.setPartner(niranjan)
	rajyam.setPartner(laxman)
	rama.setPartner(omkar)
	pallavi.setPartner(kishan)

	swagath := f.addChild(lakshmi.name, "Swagath", male)
	f.addChild(lakshmi.name, "Swapna", female)
	chinmai := f.addToFamily("Chinmai", female)
	chinmai.setPartner(swagath)
	f.addChild(chinmai.name, "Hamsika", female)

	sandhya := f.addChild(anjali.name, "Sandhya", female)
	ram := f.addToFamily("Ram", male)
	sandhya.setPartner(ram)
	f.addChild(sandhya.name, "Ashwik", male)
	f.addChild(sandhya.name, "Abhiram", male)
	santhosh := f.addChild(anjali.name, "Santhosh", male)
	niharika := f.addToFamily("Niharika", female)
	niharika.setPartner(santhosh)
	f.addChild(anjali.name, "Sathish", male)

	sampath := f.addChild(rajyam.name, "Sampath", male)
	anusha := f.addToFamily("Anusha", female)
	anusha.setPartner(sampath)
	f.addChild(rajyam.name, "Sandeep", male)

	sony := f.addChild(rama.name, "Sony", female)
	sanath := f.addToFamily("Sanath", male)
	sony.setPartner(sanath)
	f.addChild(rama.name, "Sweety", female)

	f.addChild(pallavi.name, "Lucky", female)
	f.addChild(pallavi.name, "Eeshu", male)

	return f
}

func TestParents(t *testing.T) {
	f := setUp()
	tests := []struct {
		name     string
		relation string
		exp      string
	}{
		{
			name:     "Sathish",
			exp:      "Anjali",
			relation: "Mother",
		},
		{
			name:     "Santhosh",
			exp:      "Niranjan",
			relation: "Father",
		},
	}
	for _, tc := range tests {
		assert.Equal(t, tc.exp, f.GetRelationShip(tc.name, tc.relation))
	}
}

func TestChildren(t *testing.T) {
	f := setUp()
	tests := []struct {
		name     string
		relation string
		exp      string
	}{
		{
			name:     "Anjali",
			exp:      "Santhosh Sathish",
			relation: "Son",
		},
		{
			name:     "Niranjan",
			exp:      "Sandhya",
			relation: "Daughter",
		},
	}
	for _, tc := range tests {
		assert.Equal(t, tc.exp, f.GetRelationShip(tc.name, tc.relation))
	}
}

func TestBrotherInLaw(t *testing.T) {
	f := setUp()
	tests := []struct {
		name string
		exp  string
	}{
		{
			name: "Sathish",
			exp:  "Ram",
		},
		// husbands of siblings
		{
			name: "Kishan",
			exp:  "Jayanth Niranjan Laxman Omkar",
		},
		//spouse brother
		{
			name: "Niranjan",
			exp:  "Kishan",
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.exp, f.GetRelationShip(tc.name, "Brother-In-Law"))
	}
}

func TestSisterInLaw(t *testing.T) {
	f := setUp()
	tests := []struct {
		name string
		exp  string
	}{
		// spouse sister
		{
			name: "Sanath",
			exp:  "Sweety",
		},
		{
			name: "Omkar",
			exp:  "Lakshmi Anjali Rajyam",
		},
		// wives of siblings
		{
			name: "Sathish",
			exp:  "Niharika",
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.exp, f.GetRelationShip(tc.name, "Sister-In-Law"))
	}

}

func TestMaternalUncle(t *testing.T) {
	f := setUp()
	tests := []struct {
		name string
		exp  string
	}{
		{
			name: "Sathish",
			exp:  "Kishan",
		},
	}
	for _, tc := range tests {
		assert.Equal(t, tc.exp, f.GetRelationShip(tc.name, "Maternal-Uncle"))
	}
}

func TestPaternalUncle(t *testing.T) {
	f := setUp()
	tests := []struct {
		name string
		exp  string
	}{
		{
			name: "Santhosh",
			exp:  "Shivaiah",
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.exp, f.GetRelationShip(tc.name, "Paternal-Uncle"))
	}
}

func TestMaternalAunt(t *testing.T) {
	f := setUp()
	tests := []struct {
		name string
		exp  string
	}{
		{
			name: "Sathish",
			exp:  "Lakshmi Rajyam Rama",
		},
		{
			name: "Sweety",
			exp:  "Lakshmi Anjali Rajyam",
		},
		{
			name: "Swagath",
			exp:  "Anjali Rajyam Rama",
		},
	}
	for _, tc := range tests {
		assert.Equal(t, tc.exp, f.GetRelationShip(tc.name, "Maternal-Aunt"))
	}
}

func TestPaternalAunt(t *testing.T) {
	f := setUp()
	tests := []struct {
		name string
		exp  string
	}{
		{
			name: "Sathish",
			exp:  "Ammai Uma",
		},
	}
	for _, tc := range tests {
		assert.Equal(t, tc.exp, f.GetRelationShip(tc.name, "Paternal-Aunt"))
	}
}

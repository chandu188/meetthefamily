package familytree

type relatives []*member

func (f relatives) Len() int {
	return len(f)
}
func (f relatives) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
func (f relatives) Less(i, j int) bool {
	return f[i].order < f[j].order
}

type member struct {
	name     string
	gender   GENDER
	partner  *member
	children []*member
	*parents
	order int
}

func (m *member) addChild(c *member) {
	m.children = append(m.children, c)
	if m.partner != nil {
		m.partner.children = append(m.partner.children, c)
	}
}

func (c *member) addParents(m *member, f *member) {
	c.parents = &parents{
		mother: m,
		father: f,
	}
}

func (m *member) setPartner(f *member) {
	m.partner = f
	f.partner = m
}

type parents struct {
	mother *member
	father *member
}

func NewMember(name string, g GENDER) *member {
	return &member{
		name:   name,
		gender: g,
	}
}
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
	name    string
	gender  GENDER
	partner *couple
	parent  *parents
	order   int
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

type couple struct {
	mother   *member
	father   *member
	children []*member
}

func (c *couple) other(m *member) *member {
	if m == c.mother {
		return c.father
	}
	return c.mother
}

func (c *couple) addChild(m *member) {
	c.children = append(c.children, m)
	m.parent = &parents{
		mother: c.mother,
		father: c.father,
	}
}

func NewCouple(m *member, f *member) *couple {
	c := &couple{
		mother:   m,
		father:   f,
		children: make([]*member, 0),
	}
	m.partner = c
	f.partner = c
	return c
}

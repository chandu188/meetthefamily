package familytree

//RelationFn defines a function which gives reltives of the member 'm'
type relationFn func(m *member) []*member

var relationFnMap = map[string]relationFn{
	"Paternal-Uncle": paternalUncles,
	"Maternal-Uncle": maternalUncles,
	"Paternal-Aunt":  paternalAunts,
	"Maternal-Aunt":  maternalAunts,
	"Sister-In-Law":  sisInLaws,
	"Brother-In-Law": brotherInLaws,
	"Son":            sons,
	"Daughter":       daughters,
	"Siblings":       siblings,
	"Mother":         mother,
	"Father":         father,
}

var mother relationFn = func(m *member) []*member {
	if m.parent == nil {
		return nil
	}
	return []*member{m.parent.mother}
}

var father relationFn = func(m *member) []*member {
	if m.parent == nil {
		return nil
	}
	return []*member{m.parent.father}
}

var children relationFn = func(m *member) []*member {
	if m.partner == nil {
		return nil
	}
	return m.partner.children
}

var siblings relationFn = func(m *member) []*member {
	if m.parent == nil {
		return nil
	}
	siblings := make([]*member, 0)
	mother := m.parent.mother
	for _, s := range mother.partner.children {
		if s != m {
			siblings = append(siblings, s)
		}
	}
	return siblings
}

var sons relationFn = func(m *member) []*member {
	return filter(male, children(m))
}

var daughters = func(m *member) []*member {
	return filter(female, children(m))
}

var brothers = func(m *member) []*member {
	return filter(male, siblings(m))
}

var sisters = func(m *member) []*member {
	return filter(female, siblings(m))
}

var spouse = func(m *member) []*member {
	if m.partner == nil {
		return nil
	}
	return []*member{m.partner.other(m)}
}

var grandsons = func(m *member) []*member {
	return apply(sons(m), sons)
}

var grandDaughters = func(m *member) []*member {
	return apply(daughters(m), daughters)
}

//Father siblings children
var cousins = func(m *member) []*member {
	return apply(apply(father(m), siblings), children)
}

var brotherInLaws = func(m *member) []*member {
	sisHusbands := apply(sisters(m), spouse)
	spouseBrothers := apply(spouse(m), brothers)
	return union(sisHusbands, spouseBrothers)
}

var sisInLaws = func(m *member) []*member {
	brothersWives := apply(brothers(m), spouse)
	spouseSisters := apply(spouse(m), sisters)
	return union(brothersWives, spouseSisters)
}

var paternalUncles = func(m *member) []*member {
	return apply(father(m), brothers)
}

var maternalUncles = func(m *member) []*member {
	return apply(mother(m), brothers)
}

var paternalAunts = func(m *member) []*member {
	return apply(father(m), sisters)
}

var maternalAunts = func(m *member) []*member {
	return apply(mother(m), sisters)
}

func union(mems ...[]*member) []*member {
	res := make([]*member, 0)
	for _, m := range mems {
		res = append(res, m...)
	}
	return res
}

func filter(g GENDER, mems []*member) []*member {
	res := make([]*member, 0)
	for _, m := range mems {
		if m.gender == g {
			res = append(res, m)
		}
	}
	return res
}

func apply(mems []*member, fn relationFn) []*member {
	res := make([]*member, 0)
	for _, m := range mems {
		rels := fn(m)
		if len(rels) > 0 {
			res = append(res, rels...)
		}
	}
	return res
}

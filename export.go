package bite

import (
	"github.com/addyreal/bite/internal/rule"
)

type temp struct {
	rules []*rule.Rule
}
type table struct {
	anchor [256][]*rule.Rule
}

func Init() *temp {
	return &temp{}
}

func (x *temp) Add(res any, pattern ...int) {
	if len(pattern) < 1 {
		return
	}

	//temp
	if pattern[0] == -1 {
		panic("beginning with wildcard unimplemented")
	}

	for _, e := range pattern {
		if e < -1 || e > 255 {
			panic("byte value neither -1 or 0-255")
		}
	}
	x.rules = append(x.rules, rule.Make(res, pattern))
}

func (x *temp) Get() *table {
	res := &table{}

	rule.Sort(x.rules)

	for _, r := range x.rules {
		a := r.Anchor()
		res.anchor[a] = append(res.anchor[a], r)
	}

	return res
}

func (t *table) Find(b []byte) (any, bool) {
	if len(b) == 0 {
		return nil, false
	}

	a := b[0]
	for _, r := range t.anchor[a] {
		if r.Matches(b) {
	return r.Get(), true
		}
	}

	return nil, false
}

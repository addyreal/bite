package rule

import "sort"

type Rule struct {
	result  any
	pattern []int
}

func Sort(rs []*Rule) {
	sort.Slice(rs, func(i, j int) bool {
		return !(len(rs[i].pattern) < len(rs[j].pattern))
	})
}

func Make(res any, pattern []int) *Rule {
	return &Rule{result: res, pattern: pattern}
}

func (r *Rule) Get() any {
	return r.result
}

func (r *Rule) Anchor() int {
	return r.pattern[0]
}

func (r *Rule) Matches(b []byte) bool {
	if len(r.pattern) > len(b) {
		return false
	}

	for i, x := range r.pattern {
		if x == -1 {
			continue
		}

		if b[i] != byte(x) {
			return false
		}
	}

	return true
}

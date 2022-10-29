// Content managed by Project Forge, see [projectforge.md] for details.
package uhistory

import "golang.org/x/exp/slices"

type StandupHistories []*StandupHistory

func (s StandupHistories) Get(slug string) *StandupHistory {
	for _, x := range s {
		if x.Slug == slug {
			return x
		}
	}
	return nil
}

func (s StandupHistories) SlugStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range s {
		ret = append(ret, x.Slug)
	}
	return ret
}

func (s StandupHistories) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range s {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (s StandupHistories) Clone() StandupHistories {
	return slices.Clone(s)
}

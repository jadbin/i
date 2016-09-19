package discover

import (
	"sort"

	"github.com/jadbin/i/blog"
)

func GetSeriesList(metalist []*blog.Meta) []*Sticker {
	m := make(map[string]int)
	for _, meta := range metalist {
		if meta.Series != "" {
			m[meta.Series] = m[meta.Series] + 1
		}
	}
	list := []*Sticker{}
	for name, count := range m {
		list = append(list, &Sticker{Name: name, Count: count})
	}
	sort.Sort(StickerList(list))
	return list
}

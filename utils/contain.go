package utils

import (
	"sort"
)

type Contain struct {
	list []string
}

func (c *Contain) Search(item string) bool {
	//----> Sort the string slice.
	sort.Strings(c.list)

	//----> Get the string of interest item.
	i := sort.SearchStrings(c.list, item)

	return c.list[i] == item

}
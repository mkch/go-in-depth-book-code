package main

import (
	"fmt"
	"iter"
	"os"
	"slices"
	"strings"

	"github.com/mkch/iter2"
)

func main() {
	var seq iter.Seq2[*iter2.DirEntry, error] = iter2.WalkDir(
		os.DirFS("testdata"), ".")
	var entries iter.Seq[*iter2.DirEntry] = iter2.Filter(
		iter2.Keys(seq),
		func(entry *iter2.DirEntry) bool {
			name := entry.Entry.Name()
			return len(name) > 1 &&
				strings.HasPrefix(entry.Entry.Name(), ".")
		})
	var names iter.Seq[string] = iter2.Map(
		entries,
		func(entry *iter2.DirEntry) string { return entry.Entry.Name() })
	var nameList []string = slices.Collect(names)
	fmt.Println(nameList)
}

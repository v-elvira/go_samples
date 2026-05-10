package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"
	"testing"
)

// начало решения

// prettify возвращает отформатированное
// строковое представление карты
func prettify(m map[string]int) string {
	keys := slices.Collect(maps.Keys(m))
	if len(m) == 0 {
		return "{}"
	} else if len(m) == 1 {
		return fmt.Sprintf("{ %s: %d }", keys[0], m[keys[0]])
	} else {
		var b strings.Builder
		slices.Sort(keys)
		b.WriteString("{\n")
		for _, k := range keys {
			//b.WriteString(strings.Repeat(" ", 4))
			fmt.Fprintf(&b, "    %s: %d,\n", k, m[k])
		}
		b.WriteRune('}')
		return b.String()
	}
}

// конец решения
func Test(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	const want = "{\n    one: 1,\n    three: 3,\n    two: 2,\n}"
	got := prettify(m)
	if got != want {
		t.Errorf("%v\ngot:\n%v\n\nwant:\n%v", m, got, want)
	}
}

package stringuage

import (
	"fmt"
	s "strings" // give a alias for strings
)

var p = fmt.Println // a shortcut to fmt.Println

func Test_strings() {
	p("Contains", s.Contains("test", "es"))
	p("Count", s.Count("test", "t"))
	p("HasPrefix", s.HasPrefix("test", "st"))
	p("HasSuffix", s.HasSuffix("test", "st"))
	p("Index", s.Index("test", "e"))
	p("Join", s.Join([]string{"a", "b"}, "-"))
	p("Repeat", s.Repeat("ab", 5))
	p("Replace", s.Replace("foo", "o", "e", -1)) // equivalent to ReplaceAll
	p("ReplaceAll", s.ReplaceAll("foo", "o", "e"))
	p("Split", s.Split("a-b-c,d", "-"))
	p("ToLower", s.ToLower("TEsT"))
	p("ToUpper", s.ToUpper("Test"))
}

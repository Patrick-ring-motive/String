package String

import (
	"fmt"
	. "github.com/Patrick-ring-motive/utils"
	"net/http"
	"strings"
	"unicode"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

type String struct {
	Value *string
}

type Strings struct {
	Value []String
}

type StringTypes interface {
	string | *string | String | *String
}

func UnwrapStr[STR StringTypes](s STR) string {
	str := AsInterface(s)
	switch v := str.(type) {
	case string:
		return v
	case *string:
		if v == nil {
			return "nil"
		}
		return *v
	case String:
		if v.Value == nil {
			return "nil"
		}
		return *v.Value
	case *String:
		if v == nil || v.Value == nil {
			return "nil"
		}
		return *v.Value
	default:
		return fmt.Sprint(str)
	}
}

func NewString[STR StringTypes](s STR) String {
	return String{Value: Ptr(UnwrapStr(s))}
}

func NewStrings(ss []string) Strings {
	strs := make([]String, len(ss))
	for i, s := range ss {
		strs[i] = NewString(s)
	}
	return Strings{Value: strs}
}

func OldStrings(strs Strings) []string {
	ss := make([]string, len(strs.Value))
	for i, s := range strs.Value {
		ss[i] = *(s.Value)
	}
	return ss
}

func S(s any) String {
	return NewString(fmt.Sprint(s))
}

func (s String) String() string {
	return *s.Value
}

func (s String) HeaderKey() String {
	return NewString(http.CanonicalHeaderKey(*(s.Value)))
}

func (s String) IncludesAny(substrs ...string) bool {
	for i := 0; i < len(substrs); i++ {
		if s.Contains(substrs[i]) {
			return true
		}
	}
	return false
}

func (s String) Len() int {
	return len(*s.Value)
}

func (s String) Clone() String {
	return NewString(strings.Clone(*s.Value))
}

func (s String) Compare(b string) int {
	return strings.Compare(*s.Value, b)
}

func (s String) Compares(b String) int {
	return strings.Compare(*s.Value, *b.Value)
}

func (s String) Contains(substr string) bool {
	return strings.Contains(*(s.Value), substr)
}

func (s String) ContainsAny(chars string) bool {
	return strings.ContainsAny(*(s.Value), chars)
}

func (s String) ContainsAnyOf(substrs ...string) bool {
	return s.IncludesAny(substrs...)
}

func (s String) ContainsFunc(f func(rune) bool) bool {
	return strings.ContainsFunc(*(s.Value), f)
}

func (s String) ContainsRune(r rune) bool {
	return strings.ContainsRune(*(s.Value), r)
}

func (s String) Count(substr ...string) int {
	sub := ""
	if len(substr) > 0 {
		sub = substr[0]
	}
	return strings.Count(*(s.Value), sub)
}

func (s String) Cut(sep string) (before, after String, found bool) {
	b, a, f := strings.Cut(*(s.Value), sep)
	B := NewString(b)
	A := NewString(a)
	return B, A, f
}

func (s String) Cuts(sep string) [3]String {
	before, after, found := strings.Cut(*(s.Value), sep)
	return [3]String{NewString(before), NewString(after), NewString(S(found))}
}

func (s String) CutPrefix(prefix string) (after String, found bool) {
	a, f := strings.CutPrefix(*(s.Value), prefix)
	A := NewString(a)
	return A, f
}

func (s String) CutsPrefix(prefix string) (after String) {
	a, f := strings.CutPrefix(*(s.Value), prefix)
	AllowUnused(f)
	A := NewString(a)
	return A
}

func (s String) CutSuffix(prefix string) (before String, found bool) {
	b, f := strings.CutSuffix(*(s.Value), prefix)
	B := NewString(b)
	return B, f
}

func (s String) CutsSuffix(prefix string) (before String) {
	b, f := strings.CutSuffix(*(s.Value), prefix)
	AllowUnused(f)
	B := NewString(b)
	return B
}

func (s String) EqualFold(t string) bool {
	return strings.EqualFold(*(s.Value), t)
}

func (s String) Fields() Strings {
	return NewStrings(strings.Fields(*(s.Value)))
}

func (s String) FieldsFunc(f func(rune) bool) Strings {
	return NewStrings(strings.FieldsFunc(*(s.Value), f))
}

func (s String) HasPrefix(prefix string) bool {
	return strings.HasPrefix(*(s.Value), prefix)
}

func (s String) HasSuffix(suffix string) bool {
	return strings.HasSuffix(*(s.Value), suffix)
}

func (s String) Index(substr string) int {
	return strings.Index(*(s.Value), substr)
}

func (s String) IndexAny(chars string) int {
	return strings.IndexAny(*(s.Value), chars)
}

func (s String) IndexAnyOf(substrs ...string) int {
	index := MaxInt
	for i := 0; i < len(substrs); i++ {
		ix := s.Index(substrs[i])
		if ix > -1 && ix < index {
			index = ix
		}
	}
	if index < MaxInt {
		return index
	}
	return -1
}

func (s String) IndexByte(c byte) int {
	return strings.IndexByte(*(s.Value), c)
}

func (s String) IndexFunc(f func(rune) bool) int {
	return strings.IndexFunc(*(s.Value), f)
}
func (s String) IndexRune(r rune) int {
	return strings.IndexRune(*(s.Value), r)
}

func (strs Strings) Join(sep string) String {
	return NewString(strings.Join(OldStrings(strs), sep))
}

func (s String) LastIndex(substr string) int {
	return strings.LastIndex(*(s.Value), substr)
}

func (s String) LastIndexAny(chars string) int {
	return strings.LastIndexAny(*(s.Value), chars)
}

func (s String) LastIndexAnyOf(substrs ...string) int {
	index := -1
	for i := 0; i < len(substrs); i++ {
		if s.Index(substrs[i]) > index {
			return s.Index(substrs[i])
		}
	}
	return index
}

func (s String) LastIndexByte(c byte) int {
	return strings.LastIndexByte(*s.Value, c)
}

func (s String) LastIndexFunc(f func(rune) bool) int {
	return strings.LastIndexFunc(*s.Value, f)
}
func (s String) Map(mapping func(rune) rune) String {
	return NewString(strings.Map(mapping, *s.Value))
}

func (s String) Repeat(count int) String {
	return NewString(strings.Repeat(*s.Value, count))
}

func (s String) Replace(old string, nw string, count ...int) String {
	n := 1
	if len(count) > 0 {
		n = count[0]
	}

	return NewString(strings.Replace(*(s.Value), old, nw, n))
}

func (s String) ReplaceAll(oldnew ...string) String {
	old := ""
	nw := ""
	if len(oldnew) > 0 {
		old = oldnew[0]
	}
	if len(oldnew) > 1 {
		nw = oldnew[1]
	}
	return NewString(strings.ReplaceAll(*s.Value, old, nw))
}

func (s String) Split(seps ...string) Strings {
	sep := ""
	if len(seps) > 0 {
		sep = seps[0]
	}
	return NewStrings(strings.Split(*s.Value, sep))
}

func (s String) SplitAfter(sep string) Strings {
	return NewStrings(strings.SplitAfter(*s.Value, sep))
}

func (s String) SplitAfterN(sep string, n ...int) Strings {
	if len(n) > 0 {
		return NewStrings(strings.SplitAfterN(*s.Value, sep, n[0]))
	}
	return NewStrings(strings.SplitAfter(*s.Value, sep))
}

func (s String) SplitN(sep string, n ...int) Strings {
	if len(n) > 0 {
		return NewStrings(strings.SplitN(*s.Value, sep, n[0]))
	}
	return NewStrings(strings.Split(*s.Value, sep))
}

func (s String) Title() String {
	return NewString(strings.Title(*s.Value))
}

func (s String) ToLower() String {
	return NewString(strings.ToLower(*s.Value))
}

func (s String) ToLowerSpecial(c unicode.SpecialCase) String {
	return NewString(strings.ToLowerSpecial(c, *s.Value))
}

func (s String) ToTitle() String {
	return NewString(strings.ToTitle(*s.Value))
}

func (s String) ToTitleSpecial(c unicode.SpecialCase) String {
	return NewString(strings.ToTitleSpecial(c, *s.Value))
}

func (s String) ToUpper() String {
	return NewString(strings.ToUpper(*s.Value))
}

func (s String) ToUpperSpecial(c unicode.SpecialCase) String {
	return NewString(strings.ToUpperSpecial(c, *s.Value))
}

func (s String) ToValidUTF8(replacement string) String {
	return NewString(strings.ToValidUTF8(*s.Value, replacement))
}

func (s String) Trim(cutset string) String {
	return NewString(strings.Trim(*s.Value, cutset))
}

func (s String) TrimFunc(f func(rune) bool) String {
	return NewString(strings.TrimFunc(*s.Value, f))
}

func (s String) TrimLeft(cutset string) String {
	return NewString(strings.TrimLeft(*s.Value, cutset))
}

func (s String) TrimLeftFunc(f func(rune) bool) String {
	return NewString(strings.TrimLeftFunc(*s.Value, f))
}

func (s String) TrimPrefix(prefix string) String {
	return NewString(strings.TrimPrefix(*s.Value, prefix))
}

func (s String) TrimRight(cutset string) String {
	return NewString(strings.TrimRight(*s.Value, cutset))
}

func (s String) TrimRightFunc(f func(rune) bool) String {
	return NewString(strings.TrimRightFunc(*s.Value, f))
}

func (s String) TrimSpace() String {
	return NewString(strings.TrimSpace(*s.Value))
}

func (s String) TrimSuffix(suffix string) String {
	return NewString(strings.TrimSuffix(*s.Value, suffix))
}

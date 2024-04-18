package String

import (
	"fmt"
	. "github.com/Patrick-ring-motive/utils"
	"io"
	"net/http"
	"strconv"
	"strings"
	"unicode"
)


type Strin [1]string

type Strins struct {
	Value []*Strin
}

type StrinTypes interface {
	string | *string | *Strin |Strin
}

func UnwrapSt[STR StrinTypes](s STR) string {
	str := AsInterface(s)
	switch v := str.(type) {
	case string:
		return v
	case *string:
		if v == nil {
			return "nil"
		}
		return *v
  case Strin:
    return v[0]
	case *Strin:
		if v == nil {
			return "nil"
		}
		return v[0]
	default:
		return fmt.Sprint(str)
	}
}

func NewStrin[STR StrinTypes](s STR) *Strin {
	return &Strin{UnwrapSt(s)}
}

func NewStrins(ss []string) Strins {
	strs := make([]*Strin, len(ss))
	for i, s := range ss {
		strs[i] = NewStrin(s)
	}
	return Strins{Value: strs}
}

func OldStrins(strs Strins) []string {
	ss := make([]string, len(strs.Value))
	for i, s := range strs.Value {
		ss[i] = s[0]
	}
	return ss
}

func (s Strin) Strin() string {
	return s[0]
}

func (s Strin) HeaderKey() *Strin {
	return NewStrin(http.CanonicalHeaderKey(s[0]))
}

func (s Strin) IncludesAny(substrs ...string) bool {
	for i := 0; i < len(substrs); i++ {
		if s.Contains(substrs[i]) {
			return true
		}
	}
	return false
}

func (s Strin) Len() int {
	return len(s[0])
}

func (s Strin) Clone() *Strin {
	return NewStrin(strings.Clone(s[0]))
}

func (s Strin) Compare(b string) int {
	return strings.Compare(s[0], b)
}

func (s Strin) Compares(b *Strin) int {
	return strings.Compare(s[0], b[0])
}

func (s Strin) Contains(substr string) bool {
	return strings.Contains(s[0], substr)
}

func (s Strin) ContainsAny(chars string) bool {
	return strings.ContainsAny(s[0], chars)
}

func (s Strin) ContainsAnyOf(substrs ...string) bool {
	return s.IncludesAny(substrs...)
}

func (s Strin) ContainsFunc(f func(rune) bool) bool {
	return strings.ContainsFunc(s[0], f)
}

func (s Strin) ContainsRune(r rune) bool {
	return strings.ContainsRune(s[0], r)
}

func (s Strin) Count(substr ...string) int {
	sub := ""
	if len(substr) > 0 {
		sub = substr[0]
	}
	return strings.Count(s[0], sub)
}

func (s Strin) Cut(sep string) (before, after *Strin, found bool) {
	b, a, f := strings.Cut(s[0], sep)
	B := NewStrin(b)
	A := NewStrin(a)
	return B, A, f
}

func (s Strin) Cuts(sep string) [3]*Strin {
	before, after, found := strings.Cut(s[0], sep)
	return [3]*Strin{NewStrin(before), NewStrin(after), NewStrin(fmt.Sprint(found))}
}

func (s Strin) CutPrefix(prefix string) (after *Strin, found bool) {
	a, f := strings.CutPrefix(s[0], prefix)
	A := NewStrin(a)
	return A, f
}

func (s Strin) CutsPrefix(prefix string) (after *Strin) {
	a, f := strings.CutPrefix(s[0], prefix)
	AllowUnused(f)
	A := NewStrin(a)
	return A
}

func (s Strin) CutSuffix(prefix string) (before *Strin, found bool) {
	b, f := strings.CutSuffix(s[0], prefix)
	B := NewStrin(b)
	return B, f
}

func (s Strin) CutsSuffix(prefix string) (before *Strin) {
	b, f := strings.CutSuffix(s[0], prefix)
	AllowUnused(f)
	B := NewStrin(b)
	return B
}

func (s Strin) EqualFold(t string) bool {
	return strings.EqualFold(s[0], t)
}

func (s Strin) Fields() Strins {
	return NewStrins(strings.Fields(s[0]))
}

func (s Strin) FieldsFunc(f func(rune) bool) Strins {
	return NewStrins(strings.FieldsFunc(s[0], f))
}

func (s Strin) HasPrefix(prefix string) bool {
	return strings.HasPrefix(s[0], prefix)
}

func (s Strin) HasSuffix(suffix string) bool {
	return strings.HasSuffix(s[0], suffix)
}

func (s Strin) Index(substr string) int {
	return strings.Index(s[0], substr)
}

func (s Strin) IndexAny(chars string) int {
	return strings.IndexAny(s[0], chars)
}

func (s Strin) IndexAnyOf(substrs ...string) int {
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

func (s Strin) IndexByte(c byte) int {
	return strings.IndexByte(s[0], c)
}

func (s Strin) IndexFunc(f func(rune) bool) int {
	return strings.IndexFunc(s[0], f)
}
func (s Strin) IndexRune(r rune) int {
	return strings.IndexRune(s[0], r)
}

func (strs Strins) Join(sep string) *Strin {
	return NewStrin(strings.Join(OldStrins(strs), sep))
}

func (s Strin) LastIndex(substr string) int {
	return strings.LastIndex(s[0], substr)
}

func (s Strin) LastIndexAny(chars string) int {
	return strings.LastIndexAny(s[0], chars)
}

func (s Strin) LastIndexAnyOf(substrs ...string) int {
	index := -1
	for i := 0; i < len(substrs); i++ {
		if s.Index(substrs[i]) > index {
			return s.Index(substrs[i])
		}
	}
	return index
}

func (s Strin) LastIndexByte(c byte) int {
	return strings.LastIndexByte(s[0], c)
}

func (s Strin) LastIndexFunc(f func(rune) bool) int {
	return strings.LastIndexFunc(s[0], f)
}
func (s Strin) Map(mapping func(rune) rune) *Strin {
	return NewStrin(strings.Map(mapping, s[0]))
}

func (s Strin) Repeat(count int) *Strin {
	return NewStrin(strings.Repeat(s[0], count))
}

func (s Strin) Replace(old string, nw string, count ...int) *Strin {
	n := 1
	if len(count) > 0 {
		n = count[0]
	}

	return NewStrin(strings.Replace(s[0], old, nw, n))
}

func (s Strin) ReplaceAll(oldnew ...string) *Strin {
	old := ""
	nw := ""
	if len(oldnew) > 0 {
		old = oldnew[0]
	}
	if len(oldnew) > 1 {
		nw = oldnew[1]
	}
	return NewStrin(strings.ReplaceAll(s[0], old, nw))
}

func (s Strin) Split(seps ...string) Strins {
	sep := ""
	if len(seps) > 0 {
		sep = seps[0]
	}
	return NewStrins(strings.Split(s[0], sep))
}

func (s Strin) SplitAfter(sep string) Strins {
	return NewStrins(strings.SplitAfter(s[0], sep))
}

func (s Strin) SplitAfterN(sep string, n ...int) Strins {
	if len(n) > 0 {
		return NewStrins(strings.SplitAfterN(s[0], sep, n[0]))
	}
	return NewStrins(strings.SplitAfter(s[0], sep))
}

func (s Strin) SplitN(sep string, n ...int) Strins {
	if len(n) > 0 {
		return NewStrins(strings.SplitN(s[0], sep, n[0]))
	}
	return NewStrins(strings.Split(s[0], sep))
}

func (s Strin) Title() *Strin {
	return NewStrin(strings.Title(s[0]))
}

func (s Strin) ToLower() *Strin {
	return NewStrin(strings.ToLower(s[0]))
}

func (s Strin) ToLowerSpecial(c unicode.SpecialCase) *Strin {
	return NewStrin(strings.ToLowerSpecial(c, s[0]))
}

func (s Strin) ToTitle() *Strin {
	return NewStrin(strings.ToTitle(s[0]))
}

func (s Strin) ToTitleSpecial(c unicode.SpecialCase) *Strin {
	return NewStrin(strings.ToTitleSpecial(c, s[0]))
}

func (s Strin) ToUpper() *Strin {
	return NewStrin(strings.ToUpper(s[0]))
}

func (s Strin) ToUpperSpecial(c unicode.SpecialCase) *Strin {
	return NewStrin(strings.ToUpperSpecial(c, s[0]))
}

func (s Strin) ToValidUTF8(replacement string) *Strin {
	return NewStrin(strings.ToValidUTF8(s[0], replacement))
}

func (s Strin) Trim(cutset string) *Strin {
	return NewStrin(strings.Trim(s[0], cutset))
}

func (s Strin) TrimFunc(f func(rune) bool) *Strin {
	return NewStrin(strings.TrimFunc(s[0], f))
}

func (s Strin) TrimLeft(cutset string) *Strin {
	return NewStrin(strings.TrimLeft(s[0], cutset))
}

func (s Strin) TrimLeftFunc(f func(rune) bool) *Strin {
	return NewStrin(strings.TrimLeftFunc(s[0], f))
}

func (s Strin) TrimPrefix(prefix string) *Strin {
	return NewStrin(strings.TrimPrefix(s[0], prefix))
}

func (s Strin) TrimRight(cutset string) *Strin {
	return NewStrin(strings.TrimRight(s[0], cutset))
}

func (s Strin) TrimRightFunc(f func(rune) bool) *Strin {
	return NewStrin(strings.TrimRightFunc(s[0], f))
}

func (s Strin) TrimSpace() *Strin {
	return NewStrin(strings.TrimSpace(s[0]))
}

func (s Strin) TrimSuffix(suffix string) *Strin {
	return NewStrin(strings.TrimSuffix(s[0], suffix))
}

func (s Strin) WriteBuilder(b *strings.Builder) (int, error) {
	return b.WriteString(s[0])
}

func (s Strin) NewReader() *strings.Reader {
	return strings.NewReader(s[0])
}

func (s Strin) Reset(r *strings.Reader) {
	r.Reset(s[0])
}

func (ss Strins) NewReplacer() *strings.Replacer {
	return strings.NewReplacer(OldStrins(ss)...)
}

func (s Strin) Replacer(r *strings.Replacer) *Strin {
	return NewStrin(r.Replace(s[0]))
}

func (s Strin) WriteReplacer(w io.Writer, r *strings.Replacer) (n int, err error) {
	return r.WriteString(w, s[0])
}

func (s Strin) AppendQuote(dst []byte) []byte {
	return strconv.AppendQuote(dst, s[0])
}

func (s Strin) AppendQuoteToASCII(dst []byte) []byte {
	return strconv.AppendQuoteToASCII(dst, s[0])
}

func (s Strin) AppendQuoteToGraphic(dst []byte) []byte {
	return strconv.AppendQuoteToGraphic(dst, s[0])
}

func (s Strin) Atoi() (int, error) {
	return strconv.Atoi(s[0])
}

func (s Strin) CanBackquote() bool {
	return strconv.CanBackquote(s[0])
}

func FormatBoo(b bool) *Strin {
	return NewStrin(strconv.FormatBool(b))
}

func FormatComple(c complex128, fmt byte, prec, bitSize int) *Strin {
	return NewStrin(strconv.FormatComplex(c, fmt, prec, bitSize))
}

func FormatFloa(f float64, fmt byte, prec, bitSize int) *Strin {
	return NewStrin(strconv.FormatFloat(f, fmt, prec, bitSize))
}

func FormatIn(i int64, base int) *Strin {
	return NewStrin(strconv.FormatInt(i, base))
}

func FormatUin(i uint64, base int) *Strin {
	return NewStrin(strconv.FormatUint(i, base))
}

func Ito(i int) *Strin {
	return NewStrin(strconv.Itoa(i))
}

func (s Strin) ParseComplex(bitSize int) (complex128, error) {
	return strconv.ParseComplex(s[0], bitSize)
}

func (s Strin) ParseComplexes(bitSize int) complex128 {
  c,err:= s.ParseComplex(bitSize)
  AllowUnused(err)
  return c
}

func (s Strin) ParseFloat(bitSize int) (float64, error) {
	return strconv.ParseFloat(s[0], bitSize)
}

func (s Strin) ParseFloats(bitSize int) float64 {
  c,err:= s.ParseFloat(bitSize)
  AllowUnused(err)
  return c
}

func (s Strin) ParseInt(base int, bitSize int) (int64, error) {
	return strconv.ParseInt(s[0], base, bitSize)
}

func (s Strin) ParseInts(base int, bitSize int) int64 {
  i,err:= s.ParseInt(base, bitSize)
  AllowUnused(err)
  return i
}

func (s Strin) ParseUint(base int, bitSize int) (uint64, error) {
  return strconv.ParseUint(s[0], base, bitSize)
}

func (s Strin) ParseUints(base int, bitSize int) uint64 {
  i,err:= s.ParseUint(base, bitSize)
  AllowUnused(err)
  return i
}

func (s Strin)Quote() *Strin{
  return NewStrin(strconv.Quote(s[0]))
}

func QuoteRun(r rune) *Strin{
  return NewStrin(strconv.QuoteRune(r))
}

func QuoteRuneToASCI(r rune) *Strin{
  return NewStrin(strconv.QuoteRuneToASCII(r))
}

func QuoteRuneToGraphi(r rune) *Strin{
  return NewStrin(strconv.QuoteRuneToGraphic(r))
}

func (s Strin)QuoteToASCII() *Strin{
  return NewStrin(strconv.QuoteToASCII(s[0]))
}

func (s Strin)QuoteToGraphic() *Strin{
  return NewStrin(strconv.QuoteToGraphic(s[0]))
}

func (s Strin)QuotedPrefix() (*Strin, error){
  str,err:=strconv.QuotedPrefix(s[0])
  return NewStrin(str),err
}

func (s Strin)QuotedPrefixes() *Strin{
  str,err:=s.QuotedPrefix()
  AllowUnused(err)
  return NewStrin(str)
}

func (s Strin)Unquote() (*Strin, error){
  str,err:=strconv.Unquote(s[0])
  return NewStrin(str),err
}

func (s Strin)Unquotes() *Strin{
  str,err:=s.Unquote()
  AllowUnused(err)
  return NewStrin(str)
}

func (s Strin)UnquoteChar(quote byte) (value rune, multibyte bool, tail *Strin, err error){
  value,multibyte,tl,err:=strconv.UnquoteChar(s[0],quote)
  return value,multibyte,NewStrin(tl),err
}

func (s Strin)UnquoteChars(quote byte)*Strin{
  value,multibyte,tl,err:=s.UnquoteChar(quote)
  AllowUnused([]any{value,multibyte,err})
  return NewStrin(tl)
}
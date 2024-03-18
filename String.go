package String

import (
	"fmt"
	"strings"
  "net/http"
  . "github.com/Patrick-ring-motive/utils"
)

type String struct {
	Value *string
}

type Strings interface{
  string|*string|String|*String
}

func UnwrapStr[T Strings](s T) string{
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

func S(s any) String {
	str := fmt.Sprint(s)
	return String{Value: &str}
}



func (s String) HeaderKey() String {
  str := http.CanonicalHeaderKey(*(s.Value))
  return String{Value: &str}
}



func (s String) Contains(substr string) bool {
	return strings.Contains(*(s.Value), substr)
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

func (s String) Replace(old string, nw string, count ...int) String {
	n := 1
	if len(count) > 0 {
		n = count[0]
	}
	str := strings.Replace(*(s.Value), old, nw, n)
	return String{Value: &str}
}

// ReplaceAll replaces all occurrences of old with new.
func (s String) ReplaceAll(oldnew ...string) String {
	old := ""
	nw := ""
	if len(oldnew) > 0 {
		old = oldnew[0]
	}
	if len(oldnew) > 1 {
		nw = oldnew[1]
	}
	 return String{Value:Ptr(strings.ReplaceAll(*s.Value, old, nw))}
}

// ToLower converts all characters to lowercase.
func (s String) ToLower() String {
	str := strings.ToLower(*s.Value)
	return String{Value: &str}
}

// ToUpper converts all characters to uppercase.
func (s String) ToUpper() String {
	str := strings.ToUpper(*s.Value)
	return String{Value: &str}
}

// TrimSpace trims whitespace from both ends of the string.
func (s String) TrimSpace() String {
	str := strings.TrimSpace(*s.Value)
	return String{Value: &str}
}

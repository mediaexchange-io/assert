package assert

import (
	"errors"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"
	"testing"
)

var (
	errorInterface  = reflect.TypeOf((*error)(nil)).Elem()
)

// Matcher hold the current state of the assertion.
type Matcher struct {
	t		*testing.T
	actual	interface{}
	match	bool
}

// With creates a new Matcher with the current test reporter.
func With(t *testing.T) *Matcher {
	m := new(Matcher)
	m.t = t
	return m
}

// That specifies the actual value under test.
func (m *Matcher) That(actual interface{}) *Matcher {
	if m.t == nil {
		panic("Use With(*testing.T) to initialize Matcher")
	}

	m.actual = actual
	return m
}

func (m *Matcher) ThatPanics(actual func()) {
	defer func() {
		if r := recover(); r == nil {
			m.t.Errorf("[%s] Did not panic.", testLine())
			m.match = false
		}
	}()
	m.match = true
	actual()
}

// IsNil verifies the tested valid is `nil`
func (m *Matcher) IsNil() *Matcher {
	if m.match = reflect.TypeOf(m.actual) == nil; !m.match {
		m.t.Errorf("[%s] is not nil", testLine())
	}
	return m
}

// IsNotNil verifies the tested value is not `nil`
func (m *Matcher) IsNotNil() *Matcher {
	if m.match = reflect.TypeOf(m.actual) != nil; !m.match {
		m.t.Errorf("[%s] is nil", testLine())
	}
	return m
}

// IsEmpty matches an empty string.
func (m *Matcher) IsEmpty() *Matcher {
	v := reflect.ValueOf(m.actual)
	if m.match = v.IsValid() && v.Kind() == reflect.String && len(v.String()) == 0; !m.match {
		m.t.Errorf("[%s] is not empty", testLine())
	}
	return m
}

// IsNotEmpty matches a non-empty string.
func (m *Matcher) IsNotEmpty() *Matcher {
	v := reflect.ValueOf(m.actual)
	if m.match = v.IsValid() && v.Kind() == reflect.String && len(v.String()) > 0; !m.match {
		m.t.Errorf("[%s] is empty", testLine())
	}
	return m
}

// IsOk expects the actual value to be nil and the type to be an instance of error.
func (m *Matcher) IsOk() *Matcher {
	t := reflect.TypeOf(m.actual)
	if m.match = t == nil || !t.Implements(errorInterface); !m.match {
		m.t.Errorf("[%s] is not ok", testLine())
	}
	return m
}

// IsEqualTo verifies that the actual value capture in `That()` is equal to the
// expected value.
func (m *Matcher) IsEqualTo(expected interface{}) *Matcher {
	m.match = false
	av := reflect.ValueOf(m.actual)
	ev := reflect.ValueOf(expected)

	// Edge condition: both values are nil. The `IsNil` matcher should be
	// used instead of IsEqualTo(), but we don't want to fail the test over
	// semantics.
	if reflect.TypeOf(m.actual) == nil && reflect.TypeOf(expected) == nil {
		m.match = true
		return m
	}

	// Both values must be valid.
	if av.IsValid() && ev.IsValid() {
		ak, err := basicKind(av)
		if err != nil {
			m.t.Error(err)
			return m
		}

		ek, err := basicKind(ev)
		if err != nil {
			m.t.Error(err)
			return m
		}

		if ak != ek {
			m.t.Errorf("[%s] %s", testLine(), errBadComparison)
			return m
		}

		switch ak {
		case boolKind:
			m.match = av.Bool() == ev.Bool()
		case complexKind:
			m.match = av.Complex() == ev.Complex()
		case floatKind:
			m.match = av.Float() == ev.Float()
		case intKind:
			m.match = av.Int() == ev.Int()
		case stringKind:
			m.match = av.String() == ev.String()
		case uintKind:
			m.match = av.Uint() == ev.Uint()
		default:
			m.t.Error(errBadType)
		}
	}

	if !m.match {
		m.t.Errorf("[%s] expected:<[%s]> but was <[%s]>", testLine(), stringValue(ev), stringValue(av))
	}

	return m
}

// IsGreaterThan matches if the actual value is greater than the expected value.
func (m *Matcher) IsGreaterThan(expected interface{}) *Matcher {
	k, err := typeCheck(m.actual, expected)
	if err != nil {
		m.match = false
		m.t.Error(err)
	} else {
		av := reflect.ValueOf(m.actual)
		ev := reflect.ValueOf(expected)
		switch k {
		case floatKind:
			m.match = av.Float() > ev.Float()
		case intKind:
			m.match = av.Int() > ev.Int()
		case uintKind:
			m.match = av.Uint() > ev.Uint()
		default:
			m.match = false
			m.t.Error(errBadType)
		}

		if !m.match {
			m.t.Errorf("[%s] expected: Greater Than <[%s]> but was <[%s]>", testLine(), stringValue(ev), stringValue(av))
		}
	}

	return m
}

func typeCheck(actual interface{}, expected interface{}) (kind, error) {
	if reflect.TypeOf(actual) == nil {
		return invalidKind, errors.New("Actual value was nil")
	}

	if reflect.TypeOf(expected) == nil {
		return invalidKind, errors.New("Expected value was nil")
	}

	av := reflect.ValueOf(actual)
	ev := reflect.ValueOf(expected)

	ak, err := basicKind(av)
	if err != nil {
		return invalidKind, errors.New("Actual " + err.Error())
	}

	ek, err := basicKind(ev)
	if err != nil {
		return invalidKind, errors.New("Expected " + err.Error())
	}

	if ak != ek {
		return invalidKind, errBadComparison
	}

	return ak, nil
}

// stringValue uses reflection to convert a `reflect.Value` to a string for
// use in error messages.
func stringValue(rv reflect.Value) string {
	switch rv.Kind() {
	case reflect.Bool:
		return strconv.FormatBool(rv.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(rv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(rv.Uint(), 10)
	case reflect.Float32:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 64)
	case reflect.Complex64:
		c := rv.Complex()
		return "(" + strconv.FormatFloat(real(c), 'g', -1, 32) + "," + strconv.FormatFloat(imag(c), 'g', -1, 32) + ")"
	case reflect.Complex128:
		c := rv.Complex()
		return "(" + strconv.FormatFloat(real(c), 'g', -1, 64) + "," + strconv.FormatFloat(imag(c), 'g', -1, 64) + ")"
	case reflect.String:
		return rv.String()
	default:
		// All of the types have been accounted for above, so this should
		// never be reached.
		panic(errBadType)
	}
}

// testLine returns the line the unit test was run from.
func testLine() string {
	lines := strings.Split(string(debug.Stack()), "\n")
	var source int
	for i, s := range lines {
		if strings.HasPrefix(s, "testing.tRunner") {
			source = i - 1
		}
	}

	line := lines[source]

	len := len(line)
	if index := strings.LastIndex(line, " +"); index >= 0 {
		len = index
	}

	if index := strings.LastIndex(line, "/"); index >= 0 {
		line = line[index + 1:len]
	} else if index := strings.LastIndex(line, "\\"); index >= 0 {
		line = line[index + 1:len]
	}

	return line
}

// The following is lifted from https://golang.org/src/text/template/funcs.go
// None of this is available outside of the package, so We're reproducing it.

// Errors returned when comparisons go bad.
var (
	errBadComparisonType    = errors.New("invalid type for comparison")
	errBadComparison        = errors.New("incompatible types for comparison")
	errBadType              = errors.New("invalid type")
)

// These are the basic types, distilled from the variety of more specific types.
type kind int
const (
	invalidKind kind = iota
	boolKind
	complexKind
	intKind
	floatKind
	stringKind
	uintKind
)

// basicKind simplifies the type down to the particular class to which it belongs.
func basicKind(v reflect.Value) (kind, error) {
	switch v.Kind() {
	case reflect.Bool:
		return boolKind, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intKind, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return uintKind, nil
	case reflect.Float32, reflect.Float64:
		return floatKind, nil
	case reflect.Complex64, reflect.Complex128:
		return complexKind, nil
	case reflect.String:
		return stringKind, nil
	}
	return invalidKind, errBadComparisonType
}

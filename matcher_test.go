package assert

import (
	"errors"
	"testing"
)

func TestWith(t *testing.T) {
	assert := With(t)

	if assert == nil {
		t.Error("With returned nil.")
	}
}

func TestMatcher_That(t *testing.T) {
	assert := With(t).That(nil)

	if assert == nil {
		t.Error("That returned nil.")
	}
}

func TestMatcher_That_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("That did not panic")
		}
	}()

	assert := new(Matcher)
	assert.That(nil)
	t.Error("That did not panic.")
}

func TestMatcher_IsNil_WithNil(t *testing.T) {
	assert := With(t).That(nil).IsNil()

	if assert == nil {
		t.Error("IsNil returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsNil matcher failed.")
	}
}

func TestMatcher_IsNil_WithInt(t *testing.T) {
	assert := With(new(testing.T)).That(0).IsNil()

	if assert == nil {
		t.Error("IsNil returned nil")
		return
	}

	if assert.match == true {
		t.Error("IsNil matcher failed.")
	}
}

func TestMatcher_IsNil_WithString(t *testing.T) {
	assert := With(new(testing.T)).That("String").IsNil()

	if assert == nil {
		t.Error("IsNil returned nil")
		return
	}

	if assert.match == true {
		t.Error("IsNil matcher failed.")
	}
}

func TestMatcher_IsNil_WithSlice(t *testing.T) {
	assert := With(new(testing.T)).That(make([]byte, 0)).IsNil()

	if assert == nil {
		t.Error("IsNil returned nil")
		return
	}

	if assert.match == true {
		t.Error("IsNil matcher failed.")
	}
}

func TestMatcher_IsNil_WithObject(t *testing.T) {
	assert := With(new(testing.T)).That(new(Matcher)).IsNil()

	if assert == nil {
		t.Error("IsNil returned nil")
		return
	}

	if assert.match == true {
		t.Error("IsNil matcher failed.")
	}
}

func TestMatcher_IsNotNil_WithNil(t *testing.T) {
	assert := With(new(testing.T)).That(nil).IsNotNil()

	if assert == nil {
		t.Error("IsNotNil returned nil")
		return
	}

	if assert.match == true {
		t.Error("IsNotNil matcher failed.")
	}
}

func TestMatcher_IsNotNil_WithInt(t *testing.T) {
	assert := With(t).That(0).IsNotNil()

	if assert == nil {
		t.Error("IsNotNil returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsNotNil matcher failed.")
	}
}

func TestMatcher_IsNotNil_WithString(t *testing.T) {
	assert := With(t).That("String").IsNotNil()

	if assert == nil {
		t.Error("IsNotNil returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsNotNil matcher failed.")
	}
}

func TestMatcher_IsNotNil_WithSlice(t *testing.T) {
	assert := With(t).That(make([]byte, 0)).IsNotNil()

	if assert == nil {
		t.Error("IsNotNil returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsNotNil matcher failed.")
	}
}

func TestMatcher_IsNotNil_WithObject(t *testing.T) {
	assert := With(t).That(new(Matcher)).IsNotNil()

	if assert == nil {
		t.Error("IsNotNil returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsNotNil matcher failed.")
	}
}

func TestMatcher_IsEmpty_WithEmptyString(t *testing.T) {
	assert := With(t).That("").IsEmpty()

	if assert == nil {
		t.Error("IsEmpty returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsEmpty matcher failed.")
	}
}

func TestMatcher_IsEmpty_WithString(t *testing.T) {
	assert := With(new(testing.T)).That("abc").IsEmpty()

	if assert == nil {
		t.Error("IsEmpty returned nil")
		return
	}

	if assert.match == true {
		t.Error("IsEmpty matcher failed.")
	}
}

func TestMatcher_IsNotEmpty_WithString(t *testing.T) {
	assert := With(t).That("abc").IsNotEmpty()

	if assert == nil {
		t.Error("IsNotEmpty returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsNotEmpty matcher failed.")
	}
}

func TestMatcher_IsNotEmpty_WithEmptyString(t *testing.T) {
	assert := With(new(testing.T)).That("").IsNotEmpty()

	if assert == nil {
		t.Error("IsNotEmpty returned nil")
		return
	}

	if assert.match == true {
		t.Error("IsNotEmpty matcher failed.")
	}
}

func TestMatcher_IsOk_WithError(t *testing.T) {
	err := errors.New("test")
	assert := With(new(testing.T)).That(err).IsOk()

	if assert == nil {
		t.Error("IsOk returned nil")
		return
	}

	if assert.match == true {
		t.Error("IsOk matcher failed.")
	}
}

func TestMatcher_IsOk_WithNil(t *testing.T) {
	assert := With(new(testing.T)).That(nil).IsOk()

	if assert == nil {
		t.Error("IsOk returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsOk matcher failed.")
	}
}

func TestMatcher_Equals_WithNil(t *testing.T) {
	assert := With(t).That(nil).IsEqualTo(nil)

	if assert == nil {
		t.Error("IsEqualTo returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsEqualTo matcher failed")
	}
}


func TestMatcher_Equals_WithBool(t *testing.T) {
	assert := With(t).That(true).IsEqualTo(true)

	if assert == nil {
		t.Error("IsEqualTo returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsEqualTo matcher failed")
	}
}

func TestMatcher_Equals_WithComplex(t *testing.T) {
	assert := With(t).That(complex(1.0, 1.0)).IsEqualTo(complex(1.0, 1.0))

	if assert == nil {
		t.Error("IsEqualTo returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsEqualTo matcher failed")
	}
}

func TestMatcher_Equals_WithFloat(t *testing.T) {
	assert := With(t).That(3.14159).IsEqualTo(3.14159)

	if assert == nil {
		t.Error("IsEqualTo returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsEqualTo matcher failed")
	}
}

func TestMatcher_Equals_WithInt(t *testing.T) {
	assert := With(t).That(-1073741824).IsEqualTo(-1073741824)

	if assert == nil {
		t.Error("IsEqualTo returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsEqualTo matcher failed")
	}
}

func TestMatcher_Equals_WithString(t *testing.T) {
	assert := With(t).That("The quick brown fox jumps over the lazy dog").IsEqualTo("The quick brown fox jumps over the lazy dog")

	if assert == nil {
		t.Error("IsEqualTo returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsEqualTo matcher failed")
	}
}

func TestMatcher_Equals_WithUint(t *testing.T) {
	assert := With(t).That(uint(1073741824)).IsEqualTo(uint(1073741824))

	if assert == nil {
		t.Error("IsEqualTo returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsEqualTo matcher failed")
	}
}

func TestMatcher_Equals_WithDifferentTypes(t *testing.T) {
	assert := With(new(testing.T)).That(true).IsEqualTo(1.0)

	if assert == nil {
		t.Error("IsEqualTo returned nil")
		return
	}

	if assert.match == true {
		t.Error("IsEqualTo matcher failed")
	}
}

func TestMatcher_IsGreaterThan_WithFloat(t *testing.T) {
	assert := With(new(testing.T)).That(3.14159).IsGreaterThan(3.14158)

	if assert == nil {
		t.Error("IsGreaterThan returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsGreaterThan matcher failed")
	}
}

func TestMatcher_IsGreaterThan_WithFloat2(t *testing.T) {
	assert := With(new(testing.T)).That(3.14158).IsGreaterThan(3.14159)

	if assert == nil {
		t.Error("IsGreaterThan returned nil")
		return
	}

	if assert.match == true {
		t.Error("IsGreaterThan matcher failed")
	}
}

func TestMatcher_IsGreaterThan_WithInt(t *testing.T) {
	assert := With(new(testing.T)).That(1073741824).IsGreaterThan(-1073741824)

	if assert == nil {
		t.Error("IsGreaterThan returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsGreaterThan matcher failed")
	}
}

func TestMatcher_IsGreaterThan_WithInt2(t *testing.T) {
	assert := With(new(testing.T)).That(-1073741824).IsGreaterThan(1073741824)

	if assert == nil {
		t.Error("IsGreaterThan returned nil")
		return
	}

	if assert.match == true {
		t.Error("IsGreaterThan matcher failed")
	}
}

func TestMatcher_IsGreaterThan_WithUint(t *testing.T) {
	assert := With(new(testing.T)).That(uint(1073741824)).IsGreaterThan(uint(1073741823))

	if assert == nil {
		t.Error("IsGreaterThan returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsGreaterThan matcher failed")
	}
}

func TestMatcher_IsGreaterThan_WithUint2(t *testing.T) {
	assert := With(new(testing.T)).That(uint(1073741823)).IsGreaterThan(uint(1073741824))

	if assert == nil {
		t.Error("IsGreaterThan returned nil")
		return
	}

	if assert.match == true {
		t.Error("IsGreaterThan matcher failed")
	}
}

func TestMatcher_ThatPanics_WithPanic(t *testing.T) {
	assert := With(t)
	p := func() {
		panic("Panic! at the Disco")
	}

	assert.ThatPanics(p)

	if assert.match == false {
		t.Error("ThatPanics matcher failed.")
	}
}

func TestMatcher_ThatPanics_WithoutPanic(t *testing.T) {
	assert := With(new(testing.T))

	p := func() {
		// Do nothing
	}

	assert.ThatPanics(p)

	if assert.match == true {
		t.Error("ThatPanics matcher failed.")
	}
}

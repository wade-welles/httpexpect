package httpexpect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObjectEmpty(t *testing.T) {
	checker := &mockChecker{}

	value1 := NewObject(checker, nil)

	value1.Empty()
	checker.AssertSuccess(t)
	checker.Reset()

	value1.NotEmpty()
	checker.AssertFailed(t)
	checker.Reset()

	value2 := NewObject(checker, map[string]interface{}{})

	value2.Empty()
	checker.AssertSuccess(t)
	checker.Reset()

	value2.NotEmpty()
	checker.AssertFailed(t)
	checker.Reset()

	value3 := NewObject(checker, map[string]interface{}{"":nil})

	value3.Empty()
	checker.AssertFailed(t)
	checker.Reset()

	value3.NotEmpty()
	checker.AssertSuccess(t)
	checker.Reset()
}

func TestObjectEqualNil(t *testing.T) {
	checker := &mockChecker{}

	value := NewObject(checker, nil)

	assert.Equal(t, map[string]interface{}(nil), value.Raw())

	value.Equal(nil)
	checker.AssertSuccess(t)
	checker.Reset()

	value.NotEqual(nil)
	checker.AssertFailed(t)
	checker.Reset()

	value.Equal(map[string]interface{}{})
	checker.AssertFailed(t)
	checker.Reset()

	value.NotEqual(map[string]interface{}{})
	checker.AssertSuccess(t)
	checker.Reset()

	value.Equal(map[string]interface{}{"":nil})
	checker.AssertFailed(t)
	checker.Reset()

	value.NotEqual(map[string]interface{}{"":nil})
	checker.AssertSuccess(t)
	checker.Reset()
}

func TestObjectEqualEmpty(t *testing.T) {
	checker := &mockChecker{}

	value := NewObject(checker, map[string]interface{}{})

	assert.Equal(t, map[string]interface{}{}, value.Raw())

	value.Equal(nil)
	checker.AssertFailed(t)
	checker.Reset()

	value.NotEqual(nil)
	checker.AssertSuccess(t)
	checker.Reset()

	value.Equal(map[string]interface{}{})
	checker.AssertSuccess(t)
	checker.Reset()

	value.NotEqual(map[string]interface{}{})
	checker.AssertFailed(t)
	checker.Reset()

	value.Equal(map[string]interface{}{"":nil})
	checker.AssertFailed(t)
	checker.Reset()

	value.NotEqual(map[string]interface{}{"":nil})
	checker.AssertSuccess(t)
	checker.Reset()
}

func TestObjectEqual(t *testing.T) {
	checker := &mockChecker{}

	value := NewObject(checker, map[string]interface{}{"foo": 123})

	assert.Equal(t, map[string]interface{}{"foo": 123}, value.Raw())

	value.Equal(nil)
	checker.AssertFailed(t)
	checker.Reset()

	value.NotEqual(nil)
	checker.AssertSuccess(t)
	checker.Reset()

	value.Equal(map[string]interface{}{})
	checker.AssertFailed(t)
	checker.Reset()

	value.NotEqual(map[string]interface{}{})
	checker.AssertSuccess(t)
	checker.Reset()

	value.Equal(map[string]interface{}{"FOO": 123})
	checker.AssertFailed(t)
	checker.Reset()

	value.NotEqual(map[string]interface{}{"FOO": 123})
	checker.AssertSuccess(t)
	checker.Reset()

	value.Equal(map[string]interface{}{"foo": 456})
	checker.AssertFailed(t)
	checker.Reset()

	value.NotEqual(map[string]interface{}{"foo": 456})
	checker.AssertSuccess(t)
	checker.Reset()

	value.Equal(map[string]interface{}{"foo": 123})
	checker.AssertSuccess(t)
	checker.Reset()

	value.NotEqual(map[string]interface{}{"foo": 123})
	checker.AssertFailed(t)
	checker.Reset()
}

func TestObjectContainsKey(t *testing.T) {
	checker := &mockChecker{}

	value := NewObject(checker, map[string]interface{}{"foo": 123, "bar": ""})

	value.ContainsKey("foo")
	checker.AssertSuccess(t)
	checker.Reset()

	value.NotContainsKey("foo")
	checker.AssertFailed(t)
	checker.Reset()

	value.ContainsKey("bar")
	checker.AssertSuccess(t)
	checker.Reset()

	value.NotContainsKey("bar")
	checker.AssertFailed(t)
	checker.Reset()

	value.ContainsKey("BAR")
	checker.AssertFailed(t)
	checker.Reset()

	value.NotContainsKey("BAR")
	checker.AssertSuccess(t)
	checker.Reset()
}

func TestObjectConvertEqual(t *testing.T) {
	type (
		myMap map[string]interface{}
		myInt int
	)

	checker := &mockChecker{}

	value := NewObject(checker, map[string]interface{}{"foo": 123})

	value.Equal(map[string]interface{}{"foo": "123"})
	checker.AssertFailed(t)
	checker.Reset()

	value.NotEqual(map[string]interface{}{"foo": "123"})
	checker.AssertSuccess(t)
	checker.Reset()

	value.Equal(map[string]interface{}{"foo": 123.0})
	checker.AssertSuccess(t)
	checker.Reset()

	value.NotEqual(map[string]interface{}{"foo": 123.0})
	checker.AssertFailed(t)
	checker.Reset()

	value.Equal(map[string]interface{}{"foo": 123})
	checker.AssertSuccess(t)
	checker.Reset()

	value.NotEqual(map[string]interface{}{"foo": 123})
	checker.AssertFailed(t)
	checker.Reset()

	value.Equal(myMap{"foo": myInt(123)})
	checker.AssertSuccess(t)
	checker.Reset()

	value.NotEqual(myMap{"foo": myInt(123)})
	checker.AssertFailed(t)
	checker.Reset()
}

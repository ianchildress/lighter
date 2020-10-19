package lighter

import (
	"os"
	"testing"
)

func TestStringFlag(t *testing.T) {
	testValue := "bar"

	// reset flags
	flags = make(map[string]flag)
	// reset args
	os.Args = os.Args[:1]
	// append args for test
	os.Args = append(os.Args, []string{"--required", testValue}...)

	// test required flag that we include
	foo, err := NewStringFlag("required", "should pass", true)
	if err != nil {
		t.Fatal(err)
	}

	// test flag that is not required and not included
	_, err = NewStringFlag("optional", "should pass", false)
	if err != nil {
		t.Fatal(err)
	}

	if err := Parse(); err != nil {
		t.Fatal(err)
	}

	if foo.Value() != testValue {
		t.Errorf("wanted %v got %v", testValue, foo.Value())
	}
}

func TestFailRequiredStringFlag(t *testing.T) {
	// reset flags
	flags = make(map[string]flag)
	// reset args
	os.Args = os.Args[:1]

	_, err := NewStringFlag("requiredString", "should fail", true)
	if err != nil {
		t.Fatal(err)
	}
	if err := Parse(); err == nil {
		t.Errorf("wanted error got %v", err)
	}
}

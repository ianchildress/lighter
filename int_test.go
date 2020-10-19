package lighter

import (
	"os"
	"strconv"
	"testing"
)

func TestIntFlag(t *testing.T) {
	var testValue int64 = 32

	// reset flags
	flags = make(map[string]flag)
	// reset args
	os.Args = os.Args[:1]
	// append args for test
	os.Args = append(os.Args, []string{"--required", strconv.FormatInt(testValue, 10)}...)

	// test required flag that we include
	foo, err := NewInt64Flag("required", "should pass", true)
	if err != nil {
		t.Fatal(err)
	}

	// test flag that is not required and not included
	_, err = NewInt64Flag("optional", "should pass", false)
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

func TestFailRequiredIntFlag(t *testing.T) {
	// reset flags
	flags = make(map[string]flag)
	// reset args
	os.Args = os.Args[:1]

	_, err := NewStringFlag("requiredInt", "should fail", true)
	if err != nil {
		t.Fatal(err)
	}
	if err := Parse(); err == nil {
		t.Errorf("wanted error got %v", err)
	}
}

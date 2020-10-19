package lighter

import (
	"os"
	"testing"
)

func TestBoolFlag(t *testing.T) {
	// reset flags
	flags = make(map[string]flag)
	// reset args
	os.Args = os.Args[:1]
	// append args for test
	os.Args = append(os.Args, []string{"--required", "true"}...)

	// test required flag that we include
	f, err := NewBoolFlag("required", "should pass", true)
	if err != nil {
		t.Fatal(err)
	}

	// test flag that is not required and not included
	_, err = NewBoolFlag("optional", "should pass", false)
	if err != nil {
		t.Fatal(err)
	}

	if err := Parse(); err != nil {
		t.Fatal(err)
	}

	if f.Value() != true {
		t.Errorf("wanted %v got %v", true, f.Value())
	}
}

func TestFailRequiredBoolFlag(t *testing.T) {
	// reset flags
	flags = make(map[string]flag)
	// reset args
	os.Args = os.Args[:1]

	_, err := NewBoolFlag("requiredBool", "should fail", true)
	if err != nil {
		t.Fatal(err)
	}
	if err := Parse(); err == nil {
		t.Errorf("wanted error got %v", err)
	}
}

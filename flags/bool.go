package flags

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

type BoolFlag struct {
	name        string
	description string
	value       bool
	required    bool
	isSet       bool
	m           sync.Mutex
}

func NewBoolFlag(name, description string, required bool) (*BoolFlag, error) {
	f := &BoolFlag{
		name:        name,
		description: description,
		required:    required,
	}

	// register flag
	if err := registerFlag(f); err != nil {
		return &BoolFlag{}, err
	}

	return f, nil
}

func (f *BoolFlag) Name() string {
	return f.name
}

func (f *BoolFlag) Description() string {
	return f.description
}

func (f *BoolFlag) Value() bool {
	return f.value
}

func (f *BoolFlag) Required() bool {
	return f.required
}

func (f *BoolFlag) IsSet() bool {
	return f.isSet
}

func (f *BoolFlag) parse() error {
	// iterate over arguments looking for this flag
	for i := 0; i < len(os.Args)-1; i++ { // skip last item since we are checking for the flag name not the value
		if os.Args[i] == fmt.Sprintf("--%s", f.name) {
			b, err := strconv.ParseBool(os.Args[i+1])
			if err != nil {
				return fmt.Errorf("bad bool value %s for flag %s", os.Args[i+1], f.name)
			}
			f.value = b
			f.isSet = true
			return nil
		}
	}

	// return error if this flag didn't exist and was required
	if f.required {
		return missingFlagError(f.name)
	}

	return nil
}

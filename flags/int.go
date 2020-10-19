package flags

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

type IntFlag struct {
	name        string
	description string
	value       int64
	required    bool
	isSet       bool
	m           sync.Mutex
}

func NewInt64Flag(name, description string, required bool) (*IntFlag, error) {
	f := &IntFlag{
		name:        name,
		description: description,
		required:    required,
	}

	// register flag
	if err := registerFlag(f); err != nil {
		return &IntFlag{}, err
	}

	return f, nil
}

func (f *IntFlag) Name() string {
	return f.name
}

func (f *IntFlag) Description() string {
	return f.description
}

func (f *IntFlag) Value() int64 {
	return f.value
}

func (f *IntFlag) Required() bool {
	return f.required
}

func (f *IntFlag) IsSet() bool {
	return f.isSet
}

func (f *IntFlag) parse() error {
	// iterate over arguments looking for this flag
	for i := 0; i < len(os.Args)-1; i++ { // skip last item since we are checking for the flag name not the value
		if os.Args[i] == fmt.Sprintf("--%s", f.name) {
			b, err := strconv.ParseInt(os.Args[i+1], 10, 64)
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

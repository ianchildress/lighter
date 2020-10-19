package flags

import (
	"fmt"
	"os"
	"sync"
)

type StringFlag struct {
	name        string
	description string
	value       string
	required    bool
	isSet       bool
	m           sync.Mutex
}

func NewStringFlag(name, description string, required bool) (*StringFlag, error) {
	f := &StringFlag{
		name:        name,
		description: description,
		required:    required,
	}

	// register flag
	if err := registerFlag(f); err != nil {
		return &StringFlag{}, err
	}

	return f, nil
}

func (f *StringFlag) Name() string {
	return f.name
}

func (f *StringFlag) Description() string {
	return f.description
}

func (f *StringFlag) Value() string {
	return f.value
}

func (f *StringFlag) Required() bool {
	return f.required
}

func (f *StringFlag) IsSet() bool {
	return f.isSet
}

func (f *StringFlag) parse() error {
	// iterate over arguments looking for this flag
	for i := 0; i < len(os.Args)-1; i++ { // skip last item since we are checking for the flag name not the value
		if os.Args[i] == fmt.Sprintf("--%s", f.name) {
			f.value = os.Args[i+1]
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

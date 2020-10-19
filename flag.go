package lighter

import (
	"fmt"
	"strings"
	"sync"
)

type flag interface {
	Name() string
	Description() string
	parse() error
}

var flags = make(map[string]flag)
var mutex = new(sync.Mutex)

func registerFlag(f flag) error {
	mutex.Lock()
	defer mutex.Unlock()

	if strings.Contains(f.Name(), " ") {
		return fmt.Errorf("flag %s contains whitespace. invalid character", f.Name())
	}

	if _, ok := flags[f.Name()]; ok {
		return fmt.Errorf("flag %s has already been registered and cannot be registered again", f.Name())
	}

	flags[f.Name()] = f
	return nil
}

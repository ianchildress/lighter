package flags

import "fmt"

func missingFlagError(name string) error {
	return fmt.Errorf("required flag %s is missing", name)
}

func Help() {

}

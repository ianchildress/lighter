package lighter

import "fmt"

func missingFlagError(name string) error {
	return fmt.Errorf("required flag %s is missing", name)
}

func Help() {

}

func HelpWithError(err error) {
	fmt.Println(err)
	Help()
}

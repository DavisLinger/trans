package trans_flag

import "fmt"

type Trans []string

func (t *Trans) String() string {
	return fmt.Sprintf("%v", []string(*t))
}

func (t *Trans) Set(value string) error {
	*t = append(*t, value)
	return nil
}

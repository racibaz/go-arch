package registry

import (
	"fmt"
)

type (
	UnregisteredKey      string
	AlreadyRegisteredKey string
)

func (key UnregisteredKey) Error() string {
	return fmt.Sprintf("nothing has been registered with the key `%s`", string(key))
}

func (key AlreadyRegisteredKey) Error() string {
	return fmt.Sprintf("nothing has been registered with the key `%s`", string(key))
}

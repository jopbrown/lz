package lz

import (
	"errors"
	"fmt"
)

func recoverFromErr(r interface{}) error {
	if err, ok := r.(error); ok {
		return err
	}

	return errors.New(fmt.Sprint(r))
}

package validator

import (
	"errors"
	"reflect"
	"strconv"
)

func castBlogIdToInt(id string) (int, error) {
	bid, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return bid, nil
}

func InvalidId(id string) error {
	bid, err := castBlogIdToInt(id)
	if err != nil {
		return err
	}

	valid := bid == 0 || (reflect.TypeOf(id).Kind() != reflect.Int)

	if !valid {
		return errors.New("Invalid id")
	}

	return nil
}

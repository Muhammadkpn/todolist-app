package petType

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	pkgError "base/pkg/constant/error"
)

type (
	Class struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}

	Type int
)

const (
	Cat Type = iota + 1
	Dog
)

var constants = []Class{
	{"CAT", "Cat"},
	{"DOG", "Dog"},
}

// Code .
func (c Type) Code() string {
	if c < 1 || int(c) > len(constants) {
		return ""
	}
	return constants[c-1].Code
}

// Name .
func (c Type) Name() string {
	if c < 1 || int(c) > len(constants) {
		return ""
	}
	return constants[c-1].Name
}

func (c *Type) Scan(val interface{}) error {
	rawVal, ok := val.(string)
	if !ok {
		return pkgError.ErrUnknownPetType
	}

	index := findIndex(rawVal)

	if index == 0 {
		return pkgError.ErrUnknownPetType
	}

	*c = Type(index)
	return nil
}

func findIndex(code string) int {
	for i, constant := range constants {
		if constant.Code == code {
			return i + 1
		}
	}
	return 0
}

// Value encodes value to the DB
func (c Type) Value() (driver.Value, error) {
	return c.Name(), nil
}

// MarshalJSON presents value to the client
func (c Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name())
}

// UnmarshalJSON parses value from the client, default is B2C
func (c *Type) UnmarshalJSON(val []byte) error {
	var rawVal string
	if err := json.Unmarshal(val, &rawVal); err != nil {
		return err
	}

	index := findIndex(rawVal)
	if index == 0 {
		return fmt.Errorf("invalid type")
	}

	*c = Type(index)
	return nil
}

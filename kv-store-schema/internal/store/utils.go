package store

import (
	"errors"
	"fmt"
)

func getValueType(value any) (Types, error) {
	var valType Types

	switch value.(type) {
	case int:
		{
			valType = Int
		}
	case string:
		{
			valType = String
		}
	case float32:
		{
			valType = Float
		}
	case bool:
		{
			valType = Bool
		}
	default:
		{
			return "", errors.Join(ErrInvalidAttrType, fmt.Errorf("value type: %T", value))
		}
	}

	return valType, nil
}

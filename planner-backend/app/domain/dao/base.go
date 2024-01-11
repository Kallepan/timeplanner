package dao

import (
	"fmt"
	"time"
)

type Base struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func ConvertNullableValueToTime(value []any) (*time.Time, error) {
	/**
	 * Converts a nullable value to a time.Time pointer
	 */
	if value == nil {
		return nil, nil
	}

	convertedValue, ok := value[0].(time.Time)
	if !ok {
		return nil, fmt.Errorf("could not convert value to time.Time")
	}

	return &convertedValue, nil
}

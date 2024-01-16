package dao

import (
	"fmt"
	"testing"
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

func TestConvertInterfaceSliceToStringSlice(t *testing.T) {
	type interfaceToStringSliceTest struct {
		arg1     []interface{}
		expected []string
	}

	var interfaceToStringSliceTests = []interfaceToStringSliceTest{
		{[]interface{}{"a", "b", "c"}, []string{"a", "b", "c"}},
		{[]interface{}{}, []string{}},
		{[]interface{}{"a", 1, "c"}, []string{"a", "1", "c"}},
	}

	for _, test := range interfaceToStringSliceTests {
		actual := convertInterfaceSliceToStringSlice(test.arg1)
		if len(actual) != len(test.expected) {
			t.Errorf("Expected %v, got %v", test.expected, actual)
		}

		for i, v := range actual {
			if v != test.expected[i] {
				t.Errorf("Expected %v, got %v", test.expected, actual)
			}
		}
	}
}

func convertInterfaceSliceToStringSlice(data []interface{}) []string {
	/**
	 * Converts an interface slice to a string slice
	 */
	result := make([]string, len(data))
	for i, v := range data {
		result[i] = v.(string)
	}
	return result
}

package tools
import (
	"fmt"
	"strconv"
	"strings"
)


// Helper function to convert to float
func ToFloat(value interface{}) (float64, error) {
	switch v := value.(type) {
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case int, int8, int16, int32, int64:
		return float64(v.(int)), nil
	case string:
		return strconv.ParseFloat(v, 64)
	default:
		return 0, fmt.Errorf("unsupported type for float conversion: %T", value)
	}
}

// Helper function to convert to int
func ToInt(value interface{}) (int64, error) {
	switch v := value.(type) {
	case int64:
		return v, nil
	case int, int8, int16, int32:
		return int64(v.(int)), nil
	case float64:
		return int64(v), nil
	case float32:
		return int64(v), nil
	case string:
		return strconv.ParseInt(v, 10, 64)
	default:
		return 0, fmt.Errorf("unsupported type for int conversion: %T", value)
	}
}

func ConvertFields(data map[string]interface{}, fieldConfig map[string]interface{}) (map[string]interface{}, error) {
	convertedFields := make(map[string]interface{})

	for field, fieldType := range fieldConfig {
		value, exists := data[field]
		if !exists {
			return nil, fmt.Errorf("field %s is missing in data", field)
		}

		var err error
		switch fieldType {
		case "float16", "float32", "float64":
			convertedFields[field], err = ToFloat(value)
			if err != nil {
				return nil, fmt.Errorf("error converting field %s to %s: %w", field, fieldType, err)
			}
		case "int8", "int16", "int32", "int64":
			convertedFields[field], err = ToInt(value)
			if err != nil {
				return nil, fmt.Errorf("error converting field %s to %s: %w", field, fieldType, err)
			}
		default:
			return nil, fmt.Errorf("unsupported field type %s for field %s", fieldType, field)
		}
	}

	return convertedFields, nil
}

func ToLowerCaseKeyMap(data map[string]interface{}) map[string]interface{} {
	mapKeyLowerCase := make(map[string]interface{})
	for key, value := range data {
		mapKeyLowerCase[strings.ToLower(key)] = value
	}
	return mapKeyLowerCase
}

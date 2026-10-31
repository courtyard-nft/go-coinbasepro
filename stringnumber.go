package coinbasepro

import (
	"encoding/json"
	"errors"
	"strconv"
)

type StringNumber string

func (s *StringNumber) UnmarshalJSON(data []byte) error {
	var v interface{}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch v := v.(type) {
	case float64:
		// Preserve the full precision of the JSON number instead of the
		// six decimal places imposed by fmt.Sprintf("%f").
		*s = StringNumber(strconv.FormatFloat(v, 'f', -1, 64))
	case int:
		*s = StringNumber(strconv.Itoa(v))
	case string:
		*s = StringNumber(v)
	default:
		return errors.New("Not an int or string")
	}

	return nil
}

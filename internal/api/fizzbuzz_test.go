package api

import (
	"net/url"
	"testing"
)

// We can have only one test for config due to the use of environment variables
func TestValidateQueryParams(t *testing.T) {

	tests := []struct {
		name                          string
		int1, int2, limit, str1, str2 string
		err                           bool
	}{
		{"normal", "5", "3", "50", "Fizz", "Buzz", false},
		{"negative", "-5", "-3", "-50", "Fizz", "Buzz", false},
		{"empty_param1", "5", "", "50", "Fizz", "Buzz", true},
		{"empty_param2", "5", "3", "50", "", "Buzz", true},
		{"zero_param", "5", "0", "50", "Fizz", "Buzz", true},
		{"not_int_param", "5", "g", "50", "Fizz", "Buzz", true},
	}
	for _, tt := range tests {
		urlVal := fillValues(tt.int1, tt.int2, tt.limit, tt.str1, tt.str2)
		t.Run(tt.name, func(t *testing.T) {
			_, err := validateQueryParams(urlVal)
			if (err != nil && !tt.err) || (err == nil && tt.err) {
				t.Errorf("test input: %v, validateQueryParams() triggers error = %t", urlVal, tt.err)
			}
		})
	}
}

func fillValues(int1, int2, lim, str1, str2 string) url.Values {
	params := url.Values{}
	params.Add("int1", int1)
	params.Add("int2", int2)
	params.Add("limit", lim)
	params.Add("str1", str1)
	params.Add("str2", str2)
	return params
}

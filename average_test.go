package mathutils

import "testing"

func TestAverage(t *testing.T) {
	tests := []struct {
		name      string
		input     []float64
		expect    float64
		fail_case bool
	}{{
		name:      "success case",
		input:     []float64{1, 2, 3, 4},
		expect:    2.5,
		fail_case: false,
	}, {
		name:      "fail case",
		input:     []float64{1, 2, 3, 4},
		expect:    2.6,
		fail_case: true,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Average(test.input)

			if test.fail_case && result == test.expect {
				t.Errorf("fail")
				return
			}
			if !test.fail_case && result != test.expect {
				t.Errorf("fail")
				return
			}
		})

	}
}

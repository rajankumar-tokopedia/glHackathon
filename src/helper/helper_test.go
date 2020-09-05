package helper

import "testing"

func TestGetRandomFloats(t *testing.T) {

	for _, tt := range []string{"1", "2", "3", "4"} {
		t.Run(tt, func(t *testing.T) {
			if got := GetRandomFloats(); got > 100.00 {
				t.Errorf("GetRandomFloats() = %v", got)
			}
		})
	}
}

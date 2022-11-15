package api

import (
	"reflect"
	"testing"
)

func TestGetJson(t *testing.T) {
	tests := []struct {
		name string
		want Сourse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetJson(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

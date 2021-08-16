package model

import (
	"reflect"
	"testing"
)

func Test_newFieldsFromSliceString(t *testing.T) {
	type args struct {
		fieldsStr []string
	}
	tests := []struct {
		name    string
		args    args
		want    Fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFieldsFromSliceString(tt.args.fieldsStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFieldsFromSliceString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFieldsFromSliceString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

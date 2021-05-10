package dataParser

import (
	"reflect"
	"testing"
)

func Test_arrayToMap(t *testing.T) {
	type args struct {
		fields    []string
		fieldName []string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{"normal input", args{[]string{"a", "b"}, []string{"1", "2"}}, map[string]string{"a": "1", "b": "2"}, false},
		{"field number not equal to header", args{[]string{"a", "b"}, []string{"1", "2", "3"}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := arrayToMap(tt.args.fields, tt.args.fieldName)
			if (err != nil) != tt.wantErr {
				t.Errorf("arrayToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("arrayToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

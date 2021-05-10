package formatConverter

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		args map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    FormatConverter
		wantErr bool
	}{
		{"no avaliable type", args{}, nil, true},
		{"output format is empty", args{map[string]string{"e": ""}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

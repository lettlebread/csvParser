package main

import (
	"reflect"
	"testing"
)

func Test_getFlags(t *testing.T) {
	type args struct {
		flagDict map[string][]string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{"invalid flag dictionary"}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFlags(tt.args.flagDict)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFlags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}

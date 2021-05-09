package dataLoader

import "testing"

func TestFileReader_Exec(t *testing.T) {
	type fields struct {
		filePath string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FileReader{
				filePath: tt.fields.filePath,
			}
			got, err := f.Exec()
			if (err != nil) != tt.wantErr {
				t.Errorf("FileReader.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FileReader.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

package migration

import (
	"reflect"
	"testing"
)

func Test_findFiles(t *testing.T) {
	tests := []struct {
		name      string
		wantFiles []string
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFiles, err := findFiles()
			if (err != nil) != tt.wantErr {
				t.Errorf("findFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFiles, tt.wantFiles) {
				t.Errorf("findFiles() = %v, want %v", gotFiles, tt.wantFiles)
			}
		})
	}
}

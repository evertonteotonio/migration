package migration

import (
	"reflect"
	"testing"
)

func Test_upFiles(t *testing.T) {
	tests := []struct {
		name      string
		wantFiles []string
		wantErr   bool
		path      string
	}{
		{
			name: "list files",
			path: "fixtures",
			wantFiles: []string{
				"fixtures/001_name.up.sql",
				"fixtures/002_b_name.up.sql",
				"fixtures/003_a_name.up.sql",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFiles, err := upFiles(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("upFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFiles, tt.wantFiles) {
				t.Errorf("upFiles() = %v, want %v", gotFiles, tt.wantFiles)
			}
		})
	}
}

func Test_downFiles(t *testing.T) {
	tests := []struct {
		name      string
		wantFiles []string
		wantErr   bool
		path      string
	}{
		{
			name: "list files",
			path: "fixtures",
			wantFiles: []string{
				"fixtures/001_name.down.sql",
				"fixtures/002_b_name.down.sql",
				"fixtures/003_a_name.down.sql",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFiles, err := downFiles(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("downFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFiles, tt.wantFiles) {
				t.Errorf("downFiles() = %v, want %v", gotFiles, tt.wantFiles)
			}
		})
	}
}

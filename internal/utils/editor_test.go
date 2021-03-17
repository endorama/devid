package utils

import "testing"

func Test_isEditorAllowed(t *testing.T) {
	type args struct {
		editorCmd string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"ed (fullpath)", args{"/bin/ed"}, true},
		{"nano (fullpath)", args{"/bin/nano"}, true},
		{"vim (fullpath)", args{"/usr/bin/vim"}, true},
		{"vim", args{"vim"}, true},
		{"nano", args{"nano"}, true},
		{"not allowed", args{"/something/else"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEditorAllowed(tt.args.editorCmd); got != tt.want {
				t.Errorf("isEditorAllowed() = %v, want %v", got, tt.want)
			}
		})
	}
}

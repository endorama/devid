package backup_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/endorama/devid/internal/backup"
)

func TestPerform(t *testing.T) {
	t.Skip("not implemented")
	type args struct {
		b          backup.Task
		passphrase string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := backup.Perform(tt.args.b, tt.args.passphrase); (err != nil) != tt.wantErr {
				t.Errorf("Perform() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPerform_Chdir(t *testing.T) {
	var out bytes.Buffer
	tk, err := backup.NewTask("test", "testdata", &out)
	assert.NoError(t, err)

	want, err := os.Getwd()
	assert.NoError(t, err)

	err = backup.Perform(tk, "test")
	assert.NoError(t, err)

	got, err := os.Getwd()
	assert.NoError(t, err)

	assert.Equal(t, want, got)
}

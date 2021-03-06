package utils_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/endorama/devid/internal/utils"
)

func TestWalk(t *testing.T) {
	files, err := utils.Walk("testdata")
	if err != nil {
		panic(err)
	}

	require.ElementsMatch(t, files, []string{"testdata/file.yaml"}, "do not match")
}

func TestWalkRelative(t *testing.T) {
	files, err := utils.WalkRelative("testdata")
	if err != nil {
		panic(err)
	}

	require.ElementsMatch(t, files, []string{"file.yaml"}, "do not match")
}

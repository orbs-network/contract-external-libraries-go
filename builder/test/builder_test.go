package test

import (
	"github.com/stretchr/testify/require"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func npmCommand(args ...string) *exec.Cmd {
	npm := exec.Command("npm", args...)
	dir, _ := filepath.Abs("../tmp")
	npm.Dir = dir
	npm.Stdout = os.Stdout
	npm.Stderr = os.Stderr

	return npm
}

func TestBuilder(t *testing.T) {
	os.RemoveAll("../tmp")

	t.Run("builder runs properly", func(t *testing.T) {
		build := exec.Command("go", "run", "../main", "-path", "../tmp", "-name", "Hello")
		err := build.Run()
		require.NoError(t, err)
	})

	t.Run("npm install", func(t *testing.T) {
		npm := npmCommand("install")

		err := npm.Run()
		require.NoError(t, err)
	})

	t.Run("npm run gamma:start", func(t *testing.T) {
		npm := npmCommand("run", "gamma:start")
		err := npm.Run()
		require.NoError(t, err)
	})

	t.Run("npm run hello:local", func(t *testing.T) {
		npm := npmCommand("run", "hello:local")
		err := npm.Run()
		require.NoError(t, err)
	})

	t.Run("npm test", func(t *testing.T) {
		npm := npmCommand("test")
		err := npm.Run()
		require.NoError(t, err)
	})

	t.Run("npm run gamma:stop", func(t *testing.T) {
		npm := npmCommand("run", "gamma:stop")
		err := npm.Run()
		require.NoError(t, err)
	})
}
package tasks

import (
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("Giving only task description as input", func(t *testing.T) {
		task := Task{}
		input := "Buy Coca-Cola"

		err := task.New(input)
		assertError(t, err, nil)

		got := task.Description

		if got != input {
			t.Errorf("got %q, want %q", got, input)
		}
	})

	t.Run("Giving an empty input", func(t *testing.T) {
		task := Task{}
		input := ""

		err := task.New(input)
		assertError(t, err, ErrCantAddEmptyTask)

		got := task.Description

		if got != input {
			t.Errorf("got %q, want %q", got, input)
		}
	})

	t.Run("Giving due format", func(t *testing.T) {
		// TODO
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q, want %q", got, want)
	}
}

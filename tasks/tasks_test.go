package tasks

import (
	"testing"
)

func TestNewTask(t *testing.T) {
	test := "Buy eggs"
	aTask := Task{}

	err := aTask.New(test)
	if err != nil {
		t.Errorf("Could not add %q: %q", test, err.Error())
	}
}

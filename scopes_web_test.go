package scopes

import (
	"testing"
)

func Test_create_new_scope(t *testing.T) {
	var c ControllerUseCases = NewControllerUseCases()
	if c != nil {
		t.Errorf("We have a problem")
	}
}

func NewControllerUseCases() ControllerUseCases {
	return nil
}

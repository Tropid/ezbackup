package backup

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	u := Element{"", "", []string{}}
	if u.IsValid() {
		t.Error("Unit is not valid: ", u)
	}

	u = Element{"Test", "", []string{}}
	if !u.IsValid() {
		t.Error("Unit should be valid: ", u)
	}
}

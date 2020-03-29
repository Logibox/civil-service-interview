package v1_test

import (
	"testing"
	csi "github.com/Logibox/civil-service-interview/v1"
)

func TestFindUsers(t *testing.T) {
	users, err := csi.FindUsers("London", "UK", 25000)
	if err != nil {
		t.Errorf("Error not nil: %v", err)
	}
	if len(users) == 0 {
		t.Error("Empty list returned")
	}
}

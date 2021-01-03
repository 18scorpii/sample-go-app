package vals

import (
	"testing"
)

func TestValidDepartmentEnum(t *testing.T) {
	s := "Finance"
	dept := Department(s)
	if err := dept.IsValid(); err != nil {
		t.Errorf("Value [%s] didnot correctly mapped to any enums %v", s, err)
	}
}

func TestFailureOnInValidDepartmentEnum(t *testing.T) {
	s := "Finances"
	dept := Department(s)
	if err := dept.IsValid(); err != nil {
		t.Logf("Value [%s] didnot correctly mapped to any enums %v", s, err)
	} else {
		t.Errorf("Wrong Value [%s] passed mapping as an enum", s)
	}
}

package vals

import "errors"

type Department string

const (
	Finance   Department = "Finance"
	Assembly  Department = "Assembly"
	HR        Department = "HR"
	Sales     Department = "Sales"
	Service   Department = "Service"
	Marketing Department = "Marketing"
)

func (d Department) IsValid() error {
	switch d {
	case Finance, Assembly, HR, Sales, Service, Marketing:
		return nil
	}
	return errors.New("Invalid Department type")
}

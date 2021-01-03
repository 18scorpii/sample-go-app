package vals

import (
	"testing"
)

func TestEmployeesUnmarshalling(t *testing.T) {
	allEmployees, err := LoadEmployeesFromFile("../files/emp.json")
	if err != nil {
		t.Errorf("Error in loading / parsing emp json file due to %v", err)
	} else if len(allEmployees) == 0 {
		t.Errorf("Error in parsing emp json file with result array as %v", allEmployees)
	} else {
		t.Logf("Successfully Parsed Employees are %v \n", allEmployees)
	}
}

func TestEmployeesMarshalling(t *testing.T) {
	allEmployees, err := LoadEmployeesFromFile("../files/emp.json")
	if err == nil && len(allEmployees) > 0 {
		for k, _ := range allEmployees {
			emp := allEmployees[k]
			emp.Version = emp.Version + 1
			allEmployees[k] = emp
		}
		if err = SaveEmployeesToFile("../files/emp.json", &allEmployees); err != nil {
			t.Errorf("Error Writing Employees %v for Marshalling due to %v \n", allEmployees, err)
		} else {
			t.Logf("Successfully Rewritten Employee File")
		}

	} else {
		t.Errorf("Error Loading Employees %v for Marshalling due to %v \n", allEmployees, err)
	}
}

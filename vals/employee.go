package vals

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

type Employee struct {
	Name    string     `json:"name"`
	Id      string     `json:"id"`
	Dept    Department `json:"dept"`
	Version int        `json:"version"`
}

func (e *Employee) IsValid() error {
	if e.Name != "" && e.Id != "" && e.Dept != "" {
		if err := Department(e.Dept).IsValid(); err != nil {
			return errors.New("Not A Valid Departments")
		} else {
			return nil
		}
	} else {
		return errors.New("Invalid Employee Attributes")
	}
}

type Employees struct {
	Employees []Employee `json:"employees"`
}

func (e *Employee) UpdateEmployee(employee Employee) error {
	e.Name = employee.Name
	e.Dept = employee.Dept
	e.Version = e.Version + 1
	if e.Id != employee.Id {
		return errors.New("Id not matching")
	} else {
		return nil
	}
}

/*
Loads Employee from the JSON file : files/emp.json
*/
func LoadEmployeesFromFile(empFilePath string) (map[string]Employee, error) {
	var employees Employees

	empFile, err := os.Open(empFilePath)
	if err != nil {
		log.Fatalf("Error in reading employee json file from the path %v \n", empFilePath)
	}
	defer empFile.Close()
	byteValue, err := ioutil.ReadAll(empFile)
	json.Unmarshal(byteValue, &employees)
	//To load into a generic map interface
	// var result map[string]interface{}
	// json.Unmarshal([]byte(byteValue), &result)
	empMap := make(map[string]Employee)
	for i, v := range employees.Employees {
		id := employees.Employees[i].Id
		empMap[id] = v
	}
	return empMap, err
}

func SaveEmployeesToFile(empFilePath string, empMap *map[string]Employee) error {
	employeeArr := make([]Employee, 0)
	for _, v := range *empMap {
		employeeArr = append(employeeArr, v)
	}
	employees := Employees{employeeArr}
	bytes, err := json.MarshalIndent(employees, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(empFilePath, bytes, 0755)
	if err != nil {
		return err
	} else {
		return nil
	}
}

package models

// Employee struct represents the structure of an employee
//add the bson tag to the ID field to specify that the field should be omitted if it is empty
type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
} // TODO: More details of employee should be included here which are Department, Position, etc.
package models

// Employee struct represents the structure of an employee
//add the bson tag to the ID field to specify that the field should be omitted if it is empty
type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	Department string `json:"department"`
	Position string `json:"position"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
	Duration float64 `json:"duration"` // Duration of employment in months
} // TODO: More details of employee should be included here which are Department, Position, etc.
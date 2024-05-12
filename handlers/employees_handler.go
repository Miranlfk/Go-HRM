package handlers

import (
	"github.com/Miranlfk/golang-experiments/db" // Import the db package
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/Miranlfk/golang-experiments/models"

)
//Retrieve all employees from the database
func GetAllEmployees(c *fiber.Ctx) error {
	query := bson.D{}
	collection := db.Instance.DB.Collection("employee")
	cursor, err := collection.Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	var employees []models.Employee
	if err := cursor.All(c.Context(), &employees); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(employees)
}

//Retrieve an employee from the database using the employee ID
func GetEmployee(c *fiber.Ctx) error {
	collection := db.Instance.DB.Collection("employee")
	id := c.Params("id")
	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	filter := bson.D{{Key: "_id", Value: employeeID}}
	employee := new(models.Employee)
	err = collection.FindOne(c.Context(), filter).Decode(employee)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).SendString("Employee not found")
		}
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(employee)

}

//Add a new employees to the database
func CreateEmployee(c *fiber.Ctx) error {
	collection := db.Instance.DB.Collection("employee")
	employee := new(models.Employee)
	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	employee.ID = ""
	result, err := collection.InsertOne(c.Context(), employee)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	filterRecord := bson.D{{Key: "_id", Value: result.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filterRecord)

	createdEmployee := &models.Employee{}
	createdRecord.Decode(createdEmployee)
	return c.Status(201).JSON(createdEmployee)
}

//Update an employee's details from the database using the employee ID
func UpdateEmployee(c *fiber.Ctx) error {
	collection := db.Instance.DB.Collection("employee")
	id := c.Params("id")
	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	// Retrieve the existing employee document
	existingEmployee := &models.Employee{}
	err = collection.FindOne(c.Context(), bson.D{{Key: "_id", Value: employeeID}}).Decode(existingEmployee)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).SendString("Employee not found")
		}
		return c.Status(500).SendString(err.Error())
	}
	// Parse request body into a new employee object
	updatedEmployee := new(models.Employee)
	if err := c.BodyParser(updatedEmployee); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	// Ensure that name and age cannot be changed by setting the existing name and age to the updated employee
	updatedEmployee.Name = existingEmployee.Name
	updatedEmployee.Age = existingEmployee.Age
	// Update only salary field
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "salary", Value: updatedEmployee.Salary},
			},
		},
	}
	// Perform the update
	err = collection.FindOneAndUpdate(c.Context(), bson.D{{Key: "_id", Value: employeeID}}, update).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).SendString("Employee not found")
		}
		return c.Status(500).SendString(err.Error())
	}
	updatedEmployee.ID = id
	return c.Status(200).JSON(updatedEmployee)
}

//Delete an employee from the database using the employee ID
func DeleteEmployee(c *fiber.Ctx) error {
	collection := db.Instance.DB.Collection("employee")
	id := c.Params("id")
	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	filter := bson.D{{Key: "_id", Value: employeeID}}
	result, err := collection.DeleteOne(c.Context(), &filter)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if result.DeletedCount == 0 {
		return c.Status(404).SendString("Employee not found")
	}
	return c.Status(204).JSON("Record Deleted Successfully")
}


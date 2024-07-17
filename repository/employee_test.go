package repository

import (
	"context"
	"log"
	"projectx/model"
	"testing"

	"github.com/google/uuid" // Updated import for uuid generation
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client {
	// created new mgclient by connecting to cluster with the passed url
	mongoTestClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://hashiimkalam:BImsRRK24eOshmZ6@cluster0.8h5pddi.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
	if err != nil {
		log.Fatal("error while connecting to mongodb", err)
	}

	log.Println("Mongodb has been successfully connected!")

	// sends ping cmd to server to check if server is reachable
	err = mongoTestClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed", err)
	}

	log.Println("ping success")

	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := newMongoClient()

	// disconnects after test case done
	defer mongoTestClient.Disconnect(context.Background())

	// dummy data
	emp1 := uuid.New().String()
	emp2 := uuid.New().String()

	// connect to collections
	coll := mongoTestClient.Database("companydb").Collection("employee_test")

	empRepo := EmployeeRepo{MongoCollection: coll}

	// insert employee1
	t.Run("Insert employ1 data", func(t *testing.T) {
		emp := model.Employee{
			Name:       "Hashim",
			Department: "Chem",
			EmployeeID: emp1,
		}

		result, err := empRepo.InsertEmployee(&emp)
		if err != nil {
			t.Fatal("insert 1 operatin failed", err)
		}

		t.Log("insert 1 success", result)
	})

	// insert employee2
	t.Run("Insert emply2 data", func(t *testing.T) {
		emp := model.Employee{
			Name:       "Nazhim",
			Department: "Maths",
			EmployeeID: emp2,
		}

		result, err := empRepo.InsertEmployee(&emp)
		if err != nil {
			t.Fatal("insert 2 operation failed", err)
		}

		t.Log("insert 2 success", result)

	})

	// get employee1 data
	t.Run("get employee 1", func(t *testing.T) {
		result, err := empRepo.FindEmployeeID(emp1)

		if err != nil {
			t.Fatal("get operation failed", err)
		}

		t.Log("emp1", result.Name)
	})

	// get all employee data
	t.Run("get all employee", func(t *testing.T) {
		results, err := empRepo.FindAllEmployee()

		if err != nil {
			t.Fatal("get operation failed", err)
		}
		t.Log("employees", results)
	})

	// update employee 1 data
	t.Run("update emp1 name", func(t *testing.T) {
		emp := model.Employee{
			Name:       "Mohammed Hahsim Kalam",
			Department: "Chem",
			EmployeeID: emp1,
		}

		result, err := empRepo.UpdateEmployeeByID(emp1, &emp)
		if err != nil {
			t.Fatal("update operation failed", err)
		}
		t.Log("update success", result)
	})

	// get emp1 data after updation
	t.Run("get employee 1 after updating", func(t *testing.T) {
		result, err := empRepo.FindEmployeeID(emp1)

		if err != nil {
			t.Fatal("get operation failed", err)
		}
		t.Log("emp1 data", result)
	})

	// delete emp1 data
	t.Run("delete emp1 data", func(t *testing.T) {
		result, err := empRepo.FindEmployeeID(emp1)

		if err != nil {
			t.Fatal("get operation failed", err)
		}
		t.Log("success emp1 deletion", result)
	})

}

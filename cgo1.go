package main

/*
#include <stdio.h>
#include <stdlib.h>

// Define a C struct
typedef struct {
    int id;
    char name[50];
    float salary;
} Employee;

// Function to create an Employee
Employee* create_employee(int id, const char* name, float salary) {
    Employee* emp = (Employee*)malloc(sizeof(Employee));
    emp->id = id;
    snprintf(emp->name, 50, "%s", name);
    emp->salary = salary;
    return emp;
}

// Function to display Employee details
void display_employee(Employee* emp) {
    if (emp != NULL) {
        printf("Employee ID: %d\n", emp->id);
        printf("Name: %s\n", emp->name);
        printf("Salary: %.2f\n", emp->salary);
    }
}

// Function to free Employee memory
void free_employee(Employee* emp) {
    free(emp);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	// Create a C string for the employee name
	name := C.CString("Alice")
	defer C.free(unsafe.Pointer(name)) // Free the C string after use

	// Call the C function to create an Employee
	emp := C.create_employee(101, name, 75000.50)
	defer C.free_employee(emp) // Ensure the memory is freed

	// Display Employee details using the C function
	C.display_employee(emp)

	// Access the struct fields directly in Go
	fmt.Println("Accessing fields in Go:")
	fmt.Printf("ID: %d\n", emp.id)
	fmt.Printf("Name: %s\n", C.GoString(&emp.name[0]))
	fmt.Printf("Salary: %.2f\n", emp.salary)
}

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
import "unsafe"

func main() {

	var name = C.CString("Jiten")
	defer C.free(unsafe.Pointer(name))

	emp := C.create_employee(101, name, 123.123)
	defer C.free_employee(emp)
	C.display_employee(emp)
}

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

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
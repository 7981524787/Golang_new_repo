#ifndef EMPLOYEE_H
#define EMPLOYEE_H

typedef struct {
    int id;
    char name[50];
    float salary;
} Employee;

Employee* create_employee(int id, const char* name, float salary);
void display_employee(Employee* emp);
void free_employee(Employee* emp);

#endif // EMPLOYEE_H
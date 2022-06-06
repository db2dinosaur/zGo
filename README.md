# zGo
Some tinkering in Go on z/OS. The first three items are from the language introduction on:
    https://go.dev/tour/welcome/

* hello - a "Hello World"
* calc_sqrt - a method to estimate a square root, showing loops
* fibonacci - looping to produce a fibonacci sequence
* goGetDepartments - an HTTP REST call to DB2 for z/OS to get a list of departments from the IVP DEPT table
* goGetEmpByDept - an HTTP REST call to DB2 for z/OS to get a list of employees in or managing a department (cf EMP table). This call takes parms:
    + dept = department ID (see goGetDepartments output)
    + mgr  = manager employee ID (again, see goGetDepartments output)

NEXT - do the goGet... but using certificates/https. This will require us to setup SECPORT in DB2T.
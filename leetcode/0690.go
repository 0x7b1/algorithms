package main

import "fmt"

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func sumImportance(res *int, id int, mapEmployees map[int]int, employees []*Employee) {
    for _, sub := range employees[mapEmployees[id]].Subordinates {
        *res += employees[mapEmployees[sub]].Importance
        sumImportance(res, sub, mapEmployees, employees)
    }
}

func getImportance(employees []*Employee, id int) int {
    mapEmployees := make(map[int]int)

    for i, e := range employees {
        mapEmployees[e.Id] = i
    }

    res := employees[mapEmployees[id]].Importance
    sumImportance(&res, id, mapEmployees, employees)

    return res
}

func main() {
    e1 := &Employee{1, 5, []int{2, 3}}
    e2 := &Employee{2, 3, []int{}}
    e3 := &Employee{3, 3, []int{}}

    fmt.Println(getImportance([]*Employee{e1, e2, e3}, 1))
}

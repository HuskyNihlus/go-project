package main

import (
	"fmt"
)

type param struct {
	On    bool
	Ammo  int
	Power int
}

func (s *param) Shoot() bool {
	if s.On == false {
		return false
	} else if s.Ammo > 0 {
		s.Ammo -= 1
		return true
	}
	return false
}
func (r *param) RideBike() bool {
	if r.On == false {
		return false
	} else if r.Power > 0 {
		r.Power -= 1
		return true

	}
	return false
}
func main() {
	testStruct := new(param)
	fmt.Scan(&testStruct.On, &testStruct.Ammo, &testStruct.Power)
	if testStruct.Ammo > 0 {
		lol := testStruct.Shoot()
		fmt.Printf("On=%t Ammo=%d Shoot()=%t\n", testStruct.On, testStruct.Ammo, lol)
	}
	if testStruct.Power > 0 {
		lol := testStruct.RideBike()
		fmt.Printf("On=%t Power=%d Shoot()=%t\n", testStruct.On, testStruct.Power, lol)
	}
}

//Если значение On == false, то оба метода вернут false.
//Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод возвращает true)
//Метод RideBike работает также, но только зависит от свойства Power
package main

func canBePresident(age, yearsInUkraine int, citizenship string) bool {
	if age < 35 {
		return false
	}

	if yearsInUkraine < 10 {
		return false
	}

	if citizenship != "Ukraine" {
		return false
	}

	return true
}

//func canBePresident(age, yearsInUkraine int, citizenship string) bool {
//	isOldEnough := age >= 35
//	leavesInUkraineEnough := yearsInUkraine >= 10
//	citizenshipUkraine := citizenship == "Ukraine"
//
//	return isOldEnough && leavesInUkraineEnough && citizenshipUkraine
//}

//func canBePresident(age, yearsInUkraine int, citizenship string) bool {
//	return age >= 35 && yearsInUkraine >= 10 && citizenship == "Ukraine"
//}

package milestone3

func CheckAboveAge(age int) bool {
	return age >= 18
}

// grade checking using if else
func CheckGrade(mark int8) string {

	if mark > 90 && mark <= 100 {
		return "Grade A"
	} else if mark > 80 && mark <= 90 {
		return "Grade B"
	} else {
		return "Bad fail"
	}
}

// grade checking using switch
func CheckGradeSwitch(mark int8) string {
	switch {
	case mark > 90:
		return "A"
	case mark > 80:
		return "B"
	default:
		return "F"
	}

	// Can be written lie this too
	/*
		switch mark {
		case 91, 92, 93, 94 ..., 100:
			return "A"
		case 80, 81, 82, 83:
			return "B"
		default:
			return "F"
		}
	*/
}

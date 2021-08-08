package services

func grade(skor int) string {
	if skor <= 40 {
		return "C"
	} else if skor > 40 && skor <= 70 {
		return "B"
	} else {
		return "A"
	}
}

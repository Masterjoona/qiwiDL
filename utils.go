package main

func Contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

/*func CalculateTotalSize(slice []file) int64 {
	var total int64 = 0
	for _, s := range slice {
		total += s.size
	}
	return total
}

func ParseSize(sizeStr string) int64 {
	sizeStr = strings.TrimSpace(sizeStr)
	unit := sizeStr[len(sizeStr)-2:]
	valueStr := sizeStr[:len(sizeStr)-3]
	value, err := strconv.ParseFloat(valueStr, 64)

	if err != nil {
		println("Error parsing size: " + err.Error())
		return 0
	}

	switch strings.ToUpper(unit) {
	case "KB":
		return int64(value * 1024)
	case "MB":
		return int64(value * 1024 * 1024)
	case "GB":
		return int64(value * 1024 * 1024 * 1024)
	default:
		return 0
	}
}*/

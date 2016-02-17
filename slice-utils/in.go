package sliceutils

func StringInSlice(slice []string, value string) bool {

	for _, searchValue := range slice {

		if value == searchValue {
			return true
		}

	}

	return false

}

package helpers

import "fmt"

func CheckStatus(status string) bool {
	fmt.Println(status)
	switch status {
	case "Published":
		return true
	case "Draft":
		return true
	case "Trashed":
		return true
	}

	return false

}

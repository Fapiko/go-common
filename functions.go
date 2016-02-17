package go_common
import "log"

func LogError(err error) error {

	if err != nil {
		log.Println(err)
	}

	return err

}


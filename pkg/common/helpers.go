package common

func HandleError(err error) {
	if err != nil {
		print(err.Error())
	}
}

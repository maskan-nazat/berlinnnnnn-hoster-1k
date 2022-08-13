package utils

func HandleError(err error) {
	if err != nil {
		Log(err.Error(), 4)
		return
	}
}

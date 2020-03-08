package helpers

//ResponseErr func ResponseErr(err error,method string)(map[string]interfaces{},int)
func ResponseErr(err error, method string) (map[string]interface{}, int) {
	var resp interface{}
	resp = ResponseDict(method)

	messageResponse, _ := resp.(map[string]interface{})
	message := messageResponse["body"].(map[string]interface{})["message"].(string)
	errMessage := err.Error()

	bodyResponse := map[string]interface{}{
		"code":    messageResponse["body"].(map[string]interface{})["code"],
		"message": message,
		"err":     errMessage,
	}

	responseCode := int(messageResponse["header"].(map[string]interface{})["code"].(float64))
	return bodyResponse, responseCode
}

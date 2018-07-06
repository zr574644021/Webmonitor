package base


func ResultJson(errcode int) (result map[string]interface{}) {
	result = make(map[string]interface{})
	result["status"] = errcode
	return result
}

package modify

func ConvertToNeedType(rawData map[string]interface{}, alias string) map[string]interface{} {
	return rawData[alias].(map[string]interface{})
}

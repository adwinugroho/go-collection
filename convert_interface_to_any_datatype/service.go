package convertinterfacetoanydatatype

// in example, I return to integer
func convertInterfaceToInt(intfData interface{}) int { // anyting return do you want
	switch v := intfData.(type) {
	case float64:
		var newData int = int(v)
		return newData
	case float32:
		var newData int = int(v)
		return newData
	default:
		return v.(int)
	}
}

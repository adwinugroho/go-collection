package reversearray

func ReverseData(datas []string) []string {
	newDatas := make([]string, 0, len(datas))
	for i := len(datas) - 1; i >= 0; i-- {
		newDatas = append(newDatas, datas[i])
	}
	return newDatas
}

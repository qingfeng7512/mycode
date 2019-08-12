package test


//集合 判断是否重复
func checkRepetition(data []string) []string {
	var newData = make([]string, 0)
	for i := 0; i < len(data); i++ {
		repart := false
		for j := i + 1; j < len(data); j++ {
			if data[i] == data[j] {
				repart = true
				break
			}
		}
		if !repart {
			newData = append(newData, data[i])
		}
	}
	return newData
}
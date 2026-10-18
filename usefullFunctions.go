package llg

func MaxStringLength(source []string) int {
	var num string = ""
	var curSize int = 0
	var maxSize int = 0
	for i:=0;i<len(source);i++{
		num = source[i]
		curSize = len(num)
		if curSize > maxSize { maxSize = curSize }
	}
	return maxSize
}

func NextMultipleInt(source int, multiple int) int {
	var result int = multiple
	for true {
		if source <= result {
			break
		}
		result += multiple
	}
	return result
}

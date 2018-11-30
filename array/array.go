package array

// StringInArray try to find string in array
func StringInArray(s string, arr ...string) bool {
	for i := range arr {
		if arr[i] == s {
			return true
		}
	}
	return false
}

// FillStringArray return a filled string slice by a param
func FillStringArray(param string, l int) []string {
	fArray := make([]string, l)
	for i := 0; i < l; i++ {
		fArray[i] = param
	}
	return fArray
}

// SplitStringSlice split a string array to smaller chunk
func SplitStringSlice(buf []string, chunkSize int) [][]string {
	var chunk []string
	chunks := make([][]string, 0, len(buf)/chunkSize+1)
	for len(buf) >= chunkSize {
		chunk, buf = buf[:chunkSize], buf[chunkSize:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:])
	}
	return chunks
}

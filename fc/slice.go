package fc

func SliceChunk(data []string, chunkSize int) (divided [][]string) {

	l := len(data)

	if chunkSize >= l {
		divided = append(divided, data)
		return
	}
	for i := 0; i < l; i += chunkSize {
		end := i + chunkSize
		if end > l {
			end = l
		}

		divided = append(divided, data[i:end])
	}
	return
}

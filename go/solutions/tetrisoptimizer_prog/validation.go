package main

const (
	blockSize = 20
	dot       = '.'
	hashTag   = '#'
	newLine   = '\n'
)

func isValidBlock(data []byte) bool {
	numDots, numHashtags := 0, 0
	for i, v := range data {
		switch v {
		case dot:
			numDots++
		case hashTag:
			numHashtags++
		case newLine:
			if i%5 != 4 {
				return false
			}
		default:
			return false
		}
	}
	return numDots == 12 && numHashtags == 4
}

func validateFile(data []byte) (ok bool, result []tetrimino) {
	for i := 0; i < len(data); {
		if i+blockSize > len(data) {
			break
		}
		if !isValidBlock(data[i : i+blockSize]) {
			break
		}
		t := blockToTetrimino(data[i : i+blockSize])
		if !t.isValid() {
			break
		}
		result = append(result, t)
		i += blockSize
		if i == len(data) {
			ok = true
			break
		}
		if data[i] != '\n' {
			break
		}
		i++
	}
	return ok, result
}

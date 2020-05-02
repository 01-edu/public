package main

type tetrimino [4][]byte

var validUint16Map = map[uint16]bool{
	0x8888: true, 0xf000: true, // I
	0xcc00: true,                                           // O
	0xe400: true, 0x4c40: true, 0x4e00: true, 0x8c80: true, // T
	0x88c0: true, 0xe800: true, 0xc440: true, 0x2e00: true, // L
	0x44c0: true, 0x8e00: true, 0xc880: true, 0xe200: true, // J
	0x6c00: true, 0x8c40: true, // S
	0xc600: true, 0x4c80: true, // Z
}

func (t tetrimino) isEmptyRow(row int) bool {
	result := true
	for i := 0; i < 4; i++ {
		result = result && t[row][i] == dot
	}
	return result
}

func (t tetrimino) isEmptyColumn(column int) bool {
	result := true
	for i := 0; i < 4; i++ {
		result = result && t[i][column] == dot
	}
	return result
}

func (t tetrimino) toUint16() uint16 {
	var result uint16
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result <<= 1
			if t[i][j] == hashTag {
				result |= 1
			}
		}
	}
	return result
}

func (t tetrimino) isValid() bool {
	return validUint16Map[t.toUint16()]
}

func blockToTetrimino(b []byte) tetrimino {
	in := tetrimino{b[0:4], b[5:9], b[10:14], b[15:19]}
	out := tetrimino{
		{dot, dot, dot, dot},
		{dot, dot, dot, dot},
		{dot, dot, dot, dot},
		{dot, dot, dot, dot},
	}
	y, x := 0, 0
	for in.isEmptyRow(y) {
		y++
	}
	for in.isEmptyColumn(x) {
		x++
	}
	for i := 0; i+y < 4; i++ {
		for j := 0; j+x < 4; j++ {
			out[i][j] = in[i+y][j+x]
		}
	}
	return out
}

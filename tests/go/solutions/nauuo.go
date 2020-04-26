package solutions

func Nauuo(plus, minus, rand int) string {
	if rand == 0 {
		if plus > minus {
			return "+"
		}
		if plus < minus {
			return "-"
		}
		if plus == minus {
			return "0"
		}
	}
	if plus > minus+rand {
		return "+"
	}
	if plus+rand < minus {
		return "-"
	}
	if plus+rand >= minus && plus-rand <= minus {
		return "?"
	}
	if minus+rand >= plus && minus-rand <= plus {
		return "?"
	}
	return "?"
}

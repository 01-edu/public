package solutions

//this function will put a in c; c in d; d in b and b in a
func Enigma(a ***int, b *int, c *******int, d ****int) {
	valc := *******c
	*******c = ***a
	vald := ****d
	****d = valc
	valb := *b
	*b = vald
	***a = valb
}

//Helper function used in the test for checking the function Enigma()
func Decript(a ***int, b *int, c *******int, d ****int) {
	vala := ***a
	***a = *******c
	valb := *b
	*b = vala
	vald := ****d
	****d = valb
	*******c = vald
}

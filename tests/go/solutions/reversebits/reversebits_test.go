package main

import (
	"reflect"
	"testing"

	"github.com/01-edu/z01"

	solutions "../../solutions"
)

func challengeBytes(t *testing.T,
	fn1, fn2 interface{}, args ...interface{}) {
	st1 := z01.Monitor(fn1, args)
	st2 := z01.Monitor(fn2, args)
	if !reflect.DeepEqual(st1.Results, st2.Results) {
		t.Errorf("%s(%08b) == %08b instead of %08b\n",
			z01.NameOfFunc(fn1),
			args[0].(byte),
			st1.Results[0].(byte),
			st2.Results[0].(byte),
		)
	}
}

func TestReverseBits(t *testing.T) {
	args := []byte{byte(0x26), byte(0x27), byte(0x28),
		byte(0x29), byte(0xAA), byte(0xBB)}

	for i := 0; i < 10; i++ {
		n := z01.RandIntBetween(0, 255)
		args = append(args, byte(n))
	}

	for _, v := range args {
		challengeBytes(t, ReverseBits, solutions.ReverseBits, v)
	}
}

package main

func IsAnagram(s, t string) bool {
	alph := make([]int, 26)
	for i := 0; i < len(s); i++ {
		if s[i] < 'a' || s[i] > 'z' {
			continue
		}
		alph[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		if t[i] < 'a' || t[i] > 'z' {
			continue
		}
		alph[t[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if alph[i] != 0 {
			return false
		}
	}
	return true
}

package main

func validWordAbbreviation(word string, abbr string) bool {
	//num 字母间数字大小
	num := 0
	//abbrLen 当前字母对应原串中的位置
	abbrLen := 0
	for i := 0; i < len(abbr); i++ {
		if abbr[i] == '0' && num == 0 {
			return false
		}
		if abbr[i] >= '0' && abbr[i] <= '9' {
			num = num*10 + int(abbr[i]-'0')
		} else {
			abbrLen += num
			num = 0
			if abbrLen >= len(word) || word[abbrLen] != abbr[i] {
				return false
			}
			abbrLen++
		}
	}
	return len(word)-abbrLen == num
}

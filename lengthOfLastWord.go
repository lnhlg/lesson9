package lesson9

import "strings"

//LengthOfLastWord: 最后一个单词的长度
//给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。
//单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
/*parameter
s: 字符串
return: 最后一个单词的长度
*/
func LengthOfLastWord(s string) int {

	s = strings.TrimRight(s, " ")
	idx, length := len(s)-1, 0
	for idx >= 0 && s[idx] != ' ' {

		length++
		idx--
	}

	return length
}
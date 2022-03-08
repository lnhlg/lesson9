package lesson9

import "strings"

//ToLowerCase: 转换成小写字母
//给你一个字符串 s ，将该字符串中的大写字母转换成相同的小写字母，返回新的字符串。
/*parameter
s: 字符串
return: 全部字母为小写的字符串
*/
func ToLowerCase(s string) string {

	return strings.ToLower(s)
}

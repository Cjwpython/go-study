package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
func main() {

	str1 := "hahahcsdcsd"
	front := strings.Index(str1, "a") // 获取首次出现的位置
	fmt.Println(front)

	last := strings.LastIndex(str1, "s") // 获取最后出现的位置
	fmt.Println(last)

	str2 := "hehello world"

	_str := strings.Replace(str2, "he", "wo", 1) // 根据转换次数进行替换  最后一位数字为-1 全部替换
	fmt.Println(_str)

	_str1 := strings.Count(str2, "he") //出现的次数
	fmt.Println(_str1)

	_str2 := strings.Repeat("haha", 3) //指定字符串出现多少次
	fmt.Println(_str2)

	str3 := "HELLO WORLD"
	_str3 := strings.ToLower(str3) // 全部小写
	fmt.Println(_str3)

	str4 := "hello world"
	_str4 := strings.ToUpper(str4) // 全部大写
	fmt.Println(_str4)

	str5 := "wodeshijie"
	_str5 := Ucfirst(str5) // 首字母大写
	fmt.Println(_str5)

	str6 := "  wodeshijie "
	_str6 := strings.TrimSpace(str6) // 去除首尾的空格
	fmt.Println(_str6)

	str7 := "adadbdada"
	_str7 := strings.Trim(str7, "ad") // 去除首尾的指定字符
	fmt.Println(_str7)

	str8 := "adadbdada"
	_str8 := strings.TrimLeft(str8, "ad") // 去除首字符串  贪婪模式去除
	fmt.Println(_str8)

	str9 := "adadbdada"
	_str9 := strings.TrimRight(str9, "ad") // 去除尾字符串
	fmt.Println(_str9)

	str10 := "ada dbd ada"
	_str10 := strings.Fields(str10) // 以空格进行分割
	fmt.Println(_str10)

	str11 := "ada/abs/abb"
	_str11 := strings.Split(str11, "/") // 以空格进行分割
	fmt.Println(_str11)

	str12 := []string{"!23123", "123123123", "123123"}
	_str12 := strings.Join(str12, "$") // 指定拼接
	fmt.Println(_str12)

	str13 := 12
	//_str13 := strconv.Itoa(str13)
	_str13 := fmt.Sprintf("%d",str13)  // 将数字转换成字符串
	fmt.Println(_str13)
	fmt.Printf("%T",_str13)

	str14 := "13"
	_str14 ,_:= strconv.Atoi(str14) // 将字符串转换成数字
	fmt.Println(_str14)
	fmt.Printf("%T",_str14)

}

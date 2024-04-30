package gstr

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"unicode"
)

// ToSnakeCase 驼峰转换蛇形
func ToSnakeCase(str string) string {
	var result strings.Builder
	for i, r := range str {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// SnakeToPascal 蛇形转大驼峰
func SnakeToPascal(input string) string {
	titleCaser := cases.Title(language.English)
	words := strings.Split(input, "_")
	for i, word := range words {
		words[i] = titleCaser.String(word)
	}
	return strings.Join(words, "")
}

// SnakeToCamel 将蛇形字符串（snake_case）转换为小驼峰命名（camelCase）
func SnakeToCamel(s string) string {
	titleCaser := cases.Title(language.English)
	parts := strings.Split(s, "_")
	for i := range parts {
		if i > 0 { // 从第二个单词开始将首字母转大写
			parts[i] = titleCaser.String(parts[i])
		}
	}
	return strings.Join(parts, "")
}

// TrimPrefix 删除字符串中第一个下划线及其之前的部分
func TrimPrefix(s string, i string) string {
	index := strings.Index(s, i)
	if index == -1 {
		return s
	}
	return s[index+1:]
}

// ToPascalCase 将字符串转换为大驼峰（Pascal Case）命名方式
func ToPascalCase(input string) string {
	// 首先将输入字符串转化为单词切片
	var words []string
	l := 0
	for s := input; s != ""; s = s[l:] {
		l = strings.IndexFunc(s[1:], unicode.IsUpper) + 1
		if l <= 0 {
			l = len(s)
		}
		words = append(words, s[:l])
	}

	// 将每个单词首字母大写
	titleCaser := cases.Title(language.English)
	for i, word := range words {
		words[i] = titleCaser.String(strings.ToLower(word))
	}

	// 合并所有单词成一个字符串
	return strings.Join(words, "")
}

// ToCamelCase 将字符串转换为小驼峰（camelCase）命名方式
func ToCamelCase(input string) string {
	// 首先将输入字符串转化为单词切片
	var words []string
	l := 0
	for s := input; s != ""; s = s[l:] {
		l = strings.IndexFunc(s[1:], unicode.IsUpper) + 1
		if l <= 0 {
			l = len(s)
		}
		words = append(words, s[:l])
	}

	// 对单词进行处理，第一个单词全部小写，其余单词首字母大写
	titleCaser := cases.Title(language.English)
	for i, word := range words {
		if i == 0 {
			words[i] = strings.ToLower(word) // 第一个单词全部小写
		} else {
			words[i] = titleCaser.String(strings.ToLower(word)) // 其他单词首字母大写
		}
	}

	// 合并处理后的单词为一个字符串
	return strings.Join(words, "")
}

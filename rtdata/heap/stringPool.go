package heap

import "unicode/utf16"
//字符串池

//Map<go字符串 -> java字符串>
var internedStrings = map[string]*Object{}

//根据go字符串返回java字符串，先从internedStrings里面拿，有的话直接返回
//没有的话转换并放入internedStrings，然后返回
func JString(loader *ClassLoader, goStr string) *Object {
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}

	chars := stringToUtf16(goStr)
	jChars := &Object{loader.LoadClass("[C"), chars} //加载char数组类

	jStr := loader.LoadClass("java/lang/String").NewObject() //新建string对象
	jStr.SetRefVar("value", "[C", jChars)

	internedStrings[goStr] = jStr
	return jStr
}

// java字符串 转 go字符串
func GoString(jStr *Object) string {
	charArr := jStr.GetRefVar("value", "[C") //从java String对象拿value字段，是个char数组
	return utf16ToString(charArr.Chars()) //utf16 转go字符串
}

// utf8 -> utf16
func stringToUtf16(s string) []uint16 {
	runes := []rune(s)         // utf32
	return utf16.Encode(runes) // utf32转utf16s
}

// utf16 -> utf8
func utf16ToString(s []uint16) string {
	runes := utf16.Decode(s) // func Decode(s []uint16) []rune
	return string(runes)
}

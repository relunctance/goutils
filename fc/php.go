package fc

/*
Returns a string with backslashes added before characters that need to be escaped. These characters are:

single quote (')
double quote (")
backslash (\)
*/
func Addslashes(str string) string {
	ret := make([]rune, 0, len(str))
	for _, c := range str {
		switch c {
		case
			'\\',
			'"',
			'\'':
			ret = append(ret, '\\')
		}
		ret = append(ret, c)
	}
	return string(ret)
}

// Un-quotes a quoted string.
func Stripslashes(str string) string {
	l := len(str)
	ret := make([]rune, 0, l)
	for i := 0; i < l; i++ {
		if str[i] == '\\' {
			i++
		}
		ret = append(ret, rune(str[i]))
	}
	return string(ret)
}

package lex

func isInt(b byte) bool {
	return b >= '0' && b <= '9'
}

func isString(b byte) bool {
	return b == '"'
}

func isIdent(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || b == '_'
}

func isIdentTail(b byte) bool {
	return isIdent(b) || isInt(b)
}

func isKeyword(str string) bool {
	for _, keyword := range Keywords {
		if str == keyword {
			return true
		}
	}

	return false
}

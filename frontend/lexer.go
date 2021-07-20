package frontend

type Lexer struct {
	source []byte
	length int
	currentIndex int
	nextIndex int
	tokens []Token
}


func lexerInit(contents string) Lexer {
	return Lexer {
		source: []byte(contents),
		length: len([]byte(contents)),
		currentIndex: 0,
		nextIndex: 1,
		tokens: make([]Token, 0),
	}
}


//func Lex(source string) []Token {}

func (L *Lexer) advance() {
	L.currentIndex += 1
	L.nextIndex += 1
}

func (L Lexer) current() rune {
	return rune(L.source[L.currentIndex])
}

func (L Lexer) peek() rune {
	return rune(L.source[L.nextIndex])
}


func (L *Lexer) lexTitle() {
	// consume opening '|'
	L.advance()
	beginning := L.currentIndex

	for i := 0; i < L.length; i += 1 {
		if L.current() == '|' {
			L.advance()
			break
		} else {
			L.advance()
		}
	}

	L.tokens = append(L.tokens, makeToken(L.source[beginning:L.currentIndex], Title))
}



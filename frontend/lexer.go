package frontend


import (
	"strings"
	//"fmt"
)

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


func Lex(source string) []Token {
	L := lexerInit(source)
	L.lexNoteHeader()
	L.lexItems()

	return L.tokens
}


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

func (L Lexer) prev() rune {
	return rune(L.source[(L.currentIndex - 1)])
}

func (L *Lexer) advanceTill(character rune) {
	for L.currentIndex < L.length {
		if L.current() == character {
			break
		} else {
			L.advance()
		}
	}
}

func (L *Lexer) lexNoteHeader() {
	
	for L.currentIndex < L.length {
		if L.current() == '|' {
			L.advance()
			break
		} else {
			L.advance()
		}
	}

	L.lexNoteClass()
	L.lexNoteTitle()
	L.lexNoteDate()

	// moves the lexer to the line after the header
	for L.currentIndex < L.length {
		if L.current() == '\n' {
			L.advance()
			break
		} else {
			L.advance()
		}
	}
}


func (L *Lexer) lexNoteClass() {
	beginning := L.currentIndex

	// stops once it finds either a colon or dash
	for L.currentIndex < L.length {
		if L.current() == ':' || L.current() == '-' {
			break
		} else {
			L.advance()
		}
	}

	contents := strings.TrimSpace(string(L.source[beginning:L.currentIndex]))
	L.tokens = append(L.tokens, makeToken([]byte(contents), NoteClass))
}


func (L *Lexer) lexNoteTitle() {
	// consume the ':' or '-' that followed the class name
	L.advance()
	beginning := L.currentIndex

	for L.currentIndex < L.length {
		// breaks at the opening parenthesis of the Note's date
		if L.current() == '(' {
			break
		} else {
			L.advance()
		} 
	}

	contents := strings.TrimSpace(string(L.source[beginning:L.currentIndex]))
	L.tokens = append(L.tokens, makeToken([]byte(contents), NoteTitle))
}


func (L *Lexer) lexNoteDate() {
	// consume the opening parenthesis of the date
	L.advance()
	beginning := L.currentIndex

	for L.currentIndex < L.length {
		if L.current() == ')' {
			break
		} else {
			L.advance()
		}
	}

	contents := strings.TrimSpace(string(L.source[beginning:L.currentIndex]))
	L.tokens = append(L.tokens, makeToken([]byte(contents), NoteDate))
}


func (L *Lexer) lexItems() {
	beginning := L.currentIndex

	for L.currentIndex < L.length {
		switch L.current() {
		case '#':
			L.lexPerson()
		case '*':
			L.lexEvent()
		case '[':
			L.lexAttributes()
		case '\n':
			contents := L.source[beginning:L.currentIndex]
			L.tokens = append(L.tokens, makeToken(contents, ItemLine))
			L.advance()
			beginning = L.currentIndex
		default:
			L.advance()
		}
	}
}

// left off here implementing the proper nesting of item types for Person
func (L *Lexer) lexPerson() {
	var end int
	// consume opening "#"
	L.advance()
	beginning := L.currentIndex

	for L.currentIndex < L.length {
		if L.current() == '#' {
			end = L.currentIndex
			break
		} else if L.current() == '(' {
			end = L.currentIndex
			L.lexDate()
			break
		} else {
			L.advance()
		}
	}
	
	contents := strings.TrimSpace(string(L.source[beginning:end]))
	L.tokens = append(L.tokens, makeToken([]byte(contents), Person))
	// consume closing '#'
	L.advance()
}


func (L *Lexer) lexDate() {
	// consume opening parenthesis
	L.advance()
	beginning := L.currentIndex

	for L.currentIndex < L.length {
		if L.current() == ')' {
			break
		} else {
			L.advance()
		}
	}

	contents := strings.TrimSpace(string(L.source[beginning:L.currentIndex]))
	L.tokens = append(L.tokens, makeToken([]byte(contents), Date))
	// consume closing parenthesis
	L.advance()
}


func (L *Lexer) lexEvent() {
	var end int
	// consume opening '*'
	L.advance()
	beginning := L.currentIndex

	for L.currentIndex < L.length {
		if L.current() == '*' {
			end = L.currentIndex
			break
		} else if L.current() == '(' {
			end = L.currentIndex
			L.lexDate()
		} else {
			L.advance()
		}
	}


	contents := strings.TrimSpace(string(L.source[beginning:end]))
	L.tokens = append(L.tokens, makeToken([]byte(contents), Event))
	// consume closing '*'
	L.advance()
}


func (L *Lexer) lexAttributes() {
	// consume opening '['
	L.advance()
	beginning := L.currentIndex

	for L.currentIndex < L.length {
		if L.current() == ']' {
			break
		} else {
			L.advance()
		}
	}

	contents := strings.TrimSpace(string(L.source[beginning:L.currentIndex]))
	L.tokens = append(L.tokens, makeToken([]byte(contents), AttributeList))
	// consume closing ']'
	L.advance()
}
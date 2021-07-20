package frontend


type TokenType int
const (
	Title TokenType = iota
	Name
	Date
	Text
)

type Token struct {
	tokentype TokenType
	content []byte
}

func makeToken(content []byte, tokenType TokenType) Token {
	return Token {
		tokentype: tokenType,
		content: content,
	}
}
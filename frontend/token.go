package frontend

import (
	"fmt"
)

type TokenType int
const (
	NoteTitle TokenType = iota
	NoteClass
	NoteName
	NoteDate
	Item
	ItemText
	ItemLine
	IndentationMark
	Person
	Event
	Date
	EventName
	EventDate
	RawText
	AttributeList
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


func PrintToken(token Token) {
	switch token.tokentype {
	case NoteTitle:
		fmt.Printf("NoteTitle(%s)\n", string(token.content))
	case NoteClass:
		fmt.Printf("NoteClass(%s)\n", string(token.content))
	case NoteDate:
		fmt.Printf("NoteDate(%s)\n", string(token.content))
	case NoteName:
		fmt.Printf("NoteName(%s)\n", string(token.content))
	case ItemLine:
		fmt.Printf("ItemLine(%s)\n", string(token.content))
	case Person:
		fmt.Printf("Person(%s)\n", string(token.content))
	case Event:
		fmt.Printf("Event(%s)\n", string(token.content))
	case EventName:
		fmt.Printf("EventName(%s)\n", string(token.content))
	case EventDate:
		fmt.Printf("EventDate(%s)\n", string(token.content))
	case Date:
		fmt.Printf("Date(%s)\n", string(token.content))
	case AttributeList:
		fmt.Printf("AttributeList(%s)\n", string(token.content))
	default:
		fmt.Printf("UnprintableToken\n")
	}
}

func PrintTokens(tokens []Token) {
	for _, token := range tokens {
		PrintToken(token)
	}
}
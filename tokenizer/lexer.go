package tokenizer

import "fmt"

type TokenType int

const (
	LPAREN  TokenType = 0
	RPAREN  TokenType = 1
	ID      TokenType = 2
	DONE    TokenType = 3
	ERROR   TokenType = 4
	GOOSE   TokenType = 5
	GOSLING TokenType = 6
	HONK		TokenType = 7
)

type Token struct {
	tokenType TokenType
	value     string
}

func scan(word string) Token {
	token := Token{value: "NULL"}
	switch word {
	case "Honk":
		token.tokenType = LPAREN
		return token
	case "honK":
		token.tokenType = RPAREN
		return token
	case "FLY":
		token.tokenType = DONE
		token.value = "Fly away little goose"
		return token
	case "HoNK":
		token.tokenType = GOOSE
		token.value = "Goose"
		return token
	case "HoNk":
		token.tokenType = GOSLING
		token.value = "Gosling"
		return token
	case "HONK":
		token.tokenType = HONK
		token.value = "HONK"
		return token
	default:
		token.tokenType = ID
		token.value = word
		return token
	}
}

func tokenize(words []string) []Token {
	tokens := make([]Token, len(words))

	for i, word := range words {
		tokens[i] = scan(word)
	}

	return tokens
}

func printToken(token Token) {
	switch token.tokenType {
		case LPAREN:
			fmt.Printf("(")
		case RPAREN:
			fmt.Printf(")")
		case DONE:
			fmt.Printf("FLY")
		case GOOSE:
			fmt.Printf("(GOOSE)")
		case GOSLING:
			fmt.Printf("(GOSLING)")
		case HONK:
			fmt.Printf("~HONK~")
		default:
			fmt.Printf("ID(%s)", token.value)
	}
}

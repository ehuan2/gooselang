package tokenizer

import "fmt"

type SExpType int

const (
	ATOM    SExpType = 0 // ie a single thing, in our case it'll only be variables
	LIST    SExpType = 1 // ie the final AST will hold multiple ASTs, functions, applications and hatching
	BADSEXP SExpType = 2
	EMPTY 	SExpType = 3
)

// so now we use our tokens and parse them to expressions we can turn into asts
type SExp interface {
	getType() SExpType
	getAtom() Atom
	getSExpListNode() SExpListNode
}

func (s SExpListNode) getType() SExpType {
	return LIST
}

func (s SExpListNode) getSExpListNode() SExpListNode {
	return s
}

func (s SExpListNode) getAtom() Atom {
	return Atom{}
}

type SExpListNode struct {
	first SExp
	rest  *SExpListNode
}

func (a Atom) getType() SExpType {
	return ATOM
}
func (a Atom) getSExpListNode() SExpListNode {
	return SExpListNode{}
}

func (a Atom) getAtom() Atom {
	return a
}
type Atom struct {
	atom Token
}

func (b BadSExp) getType() SExpType {
	return BADSEXP
}
func (s BadSExp) getSExpListNode() SExpListNode {
	return SExpListNode{}
}
func (s BadSExp) getAtom() Atom {
	return Atom{}
}
type BadSExp struct{}

// empty type for the end of the list
func (b Empty) getType() SExpType {
	return EMPTY
}
func (s Empty) getSExpListNode() SExpListNode {
	return SExpListNode{}
}
func (s Empty) getAtom() Atom {
	return Atom{}
}
type Empty struct{}

func IsBadList(list SExpListNode) bool {
	return list.first.getType() == BADSEXP
}

// given the list of tokens, as well as the position of the next int, parse through the sexp list
func parseSExpList(tokens []Token, position int) (SExpListNode, int) {
	// we keep iterating until we hit the RPAREN
	out := SExpListNode{}
	badList := SExpListNode{first: BadSExp{}, rest: nil}

	length := len(tokens)

	if position >= length {
		return badList, length
	}

	t := tokens[position]

	if t.tokenType == ERROR {
		position++ // increment to next one
		return badList, position
	}

	if t.tokenType == RPAREN {
		return SExpListNode{first: Empty{}}, position + 1
	}

	if t.tokenType == LPAREN {
		first := SExpListNode{}
		first, position = parseSExpList(tokens, position + 1)

		if IsBadList(first) {
			return badList, position
		}

		out.first = first
	}

	if t.tokenType == ID || t.tokenType == DONE {
		out.first = Atom{atom: t}
		position++
	}

	rest := SExpListNode{}
	rest, position = parseSExpList(tokens, position)

	if IsBadList(rest) {
		return badList, position
	}

	out.rest = &rest

	return out, position
}

func ParseSExp(tokens []Token) []SExp {
	// strategy: go through each of the tokens, grouping when possible
	// worst case scenario, all of the sexps are individual, we have to make out the same size of length, we'll deal with cutting it down afterwards

	length := len(tokens)
	sexpCounter := 0 // keep track of how many sexps we make
	out := make([]SExp, length)

	for i := 0; i < length; {
		tokenType := tokens[i].tokenType
		
		// atom cases
		if (tokenType == DONE || tokenType == ID) {
			out[sexpCounter] = Atom{atom: tokens[i]}
			sexpCounter++
			i++
		} else if (tokenType == RPAREN || tokenType == ERROR) { // error cases
			out[sexpCounter] = BadSExp{}
			sexpCounter++
			i++
		} else { // list, it's a left parenthese

			out[sexpCounter], i = parseSExpList(tokens, i + 1)
			sexpCounter++

		}

	}

	return out[0:sexpCounter]
}

func PrintSExp(s SExp) {
	sType := s.getType()

	if sType ==	ATOM {
		fmt.Printf("ATOM(%s)", s.getAtom().atom.value)
	} else if sType == LIST {

		list := s.getSExpListNode()

		if IsBadList(list) {
			fmt.Printf("Bad list...")
			return
		}

		fmt.Printf("LIST(")
		for elem := list.first; list.rest != nil && (list.rest.first.getType() != BADSEXP || list.rest.first.getType() != EMPTY); {
			PrintSExp(elem)
			elem = list.rest.first
			list = *list.rest
		}

		PrintSExp(list.first)
		fmt.Printf(")")
	}
}
package tokenizer

import "fmt"

type SExpType int

const (
	ATOM    SExpType = 0 // ie a single thing, in our case it'll only be variables
	LIST    SExpType = 1 // ie the final AST will hold multiple ASTs, functions, applications and hatching
	BADSEXP SExpType = 2
	EMPTY   SExpType = 3
)

// structs to help implement the multiple functions with more ease
func (s notAtom) getAtom() Atom {
	return Atom{}
}

type notAtom struct{}

func (s notSExpListNode) getSExpListNode() SExpListNode {
	return SExpListNode{}
}

type notSExpListNode struct{}

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

type SExpListNode struct {
	first SExp
	rest  *SExpListNode
	notAtom
}

func (a Atom) getType() SExpType {
	return ATOM
}
func (a Atom) getAtom() Atom {
	return a
}

type Atom struct {
	atom Token
	notSExpListNode
}

func (b BadSExp) getType() SExpType {
	return BADSEXP
}

type BadSExp struct { 
	notAtom
	notSExpListNode
}

// empty type for the end of the list
func (b Empty) getType() SExpType {
	return EMPTY
}

type Empty struct {
	notAtom
	notSExpListNode
}

func isBadList(list SExpListNode) bool {
	return list.first.getType() == BADSEXP
}

func generateBadList() SExpListNode {
	return SExpListNode{first: BadSExp{}, rest: nil}
}

func generateEmptyList() SExpListNode {
	return SExpListNode{first: Empty{}, rest: nil}
}

func length(list SExpListNode) int {
	curLength := 0
	for list.rest != nil && (list.rest.first.getType() != BADSEXP || list.rest.first.getType() != EMPTY) {
		curLength++
		list = *list.rest
	}
	return curLength
}

// given the list of tokens, as well as the position of the next int, parse through the sexp list
func parseSExpList(tokens []Token, position int) (SExpListNode, int) {
	// we keep iterating until we hit the RPAREN
	out := SExpListNode{}
	badList := generateBadList()

	length := len(tokens)

	if position >= length {
		return badList, length
	}

	t := tokens[position]

	if t.tokenType == ERROR || t.tokenType == LPAREN || t.tokenType == GOOSE {
		return badList, position + 1
	}

	if t.tokenType == RPAREN {
		return SExpListNode{first: Empty{}}, position + 1
	}

	if t.tokenType == GOSLING { // ie another function

		// add the gosling and the var name in
		nextSExp := SExpListNode{first: Atom{atom: tokens[position]}}

		if position + 1 >= length {
			return generateBadList(), length
		}

		if tokens[position + 1].tokenType != ID {
			return generateBadList(), position + 1 // move onto the next one, assume current one is broken, but we don't move to end
		}

		nextSExp.rest = &SExpListNode{first: Atom{atom: tokens[position + 1]}}

		if position + 2 >= length {
			return generateBadList(), length
		}

		if tokens[position + 2].tokenType != LPAREN {
			return generateBadList(), length
		}

		var recurList SExpListNode
		recurList, position = parseSExpList(tokens, position + 3)

		if isBadList(recurList) {
			return badList, position
		}

		nextSExp.rest.rest = &recurList // we skip over SExpListNode, that's why we only have two rest's

		// now we need to check if HONK comes after the function, ie is an application
		if position < length {
			if tokens[position].tokenType == HONK {
				out.first, position = honkParsing(tokens, position, nextSExp)
			} else {
				out.first = nextSExp
			}
		} else {
			out.first = nextSExp
		}
	}

	if t.tokenType == ID || t.tokenType == DONE || t.tokenType == HONK {
		out.first = Atom{atom: t}
		position++

		if position < length {
			// check for honk
			if tokens[position].tokenType == HONK {
				out.first, position = honkParsing(tokens, position, Atom{atom: t})
			}
		}

	}

	rest := SExpListNode{}
	rest, position = parseSExpList(tokens, position)

	if isBadList(rest) {
		return badList, position
	}

	out.rest = &rest

	return out, position
}

func honkParsing(tokens []Token, i int, nextSExp SExp) (SExp, int) {
	length := len(tokens)
	honkToken := tokens[i]

	var sexpArg SExp
	sexpArg, i = parseSingleSExp(tokens , i + 1)

	// in case it's a bad sexpression, return bad list
	if sexpArg.getType() == BADSEXP {
		return generateBadList(), length
	}

	// in case it's a list and is malformed, return bad list
	if sexpArg.getType() == LIST {
		if isBadList(sexpArg.getSExpListNode()) {
			return generateBadList(), length
		}
	}

	emptyList := generateEmptyList()
	// now it's good, we add HONK as first, argument as second, function as third
	outList := SExpListNode{first: Atom{atom: honkToken}, rest: &SExpListNode{first: sexpArg, rest: &SExpListNode{first: nextSExp, rest: &emptyList}}}
	return outList, i
}

// turns a stream of tokens and makes the first available sexpression out of them
func parseSingleSExp(tokens []Token, i int) (SExp, int) {

	// prelim. checking to make sure we don't go over
	length := len(tokens)
	if i >= length {
		return BadSExp{}, length
	}

	tokenType := tokens[i].tokenType

	// atom cases
	if tokenType == DONE {

		return Atom{atom: tokens[i]}, i + 1

	} else if tokenType == ID {

		// but we need to check if next is a HONK
		if i + 1 >= length {
			return Atom{atom: tokens[i]}, i + 1
		}

		if tokens[i + 1].tokenType == HONK {
			// in this case, it's a list
			return honkParsing(tokens, i + 1, Atom{atom: tokens[i]})
		}

		return Atom{atom: tokens[i]}, i + 1

	} else if tokenType == RPAREN || tokenType == ERROR || tokenType == LPAREN { // error cases

		return BadSExp{}, i + 1

	} else {

		// must be a goose or a gosling
		// in this case we make them a list, ie we group the tokens up
		// form: Goose var-name SExp that's not a goose, deal in the list parsing
		// form: Gosling var-name Honk ... honK or Gosling var-name ID

		nextSExp := SExpListNode{first: Atom{atom: tokens[i]}}

		if i + 1 >= length {
			return generateBadList(), length
		}

		if tokens[i + 1].tokenType != ID {
			return generateBadList(), i + 1 // move onto the next one, assume current one is broken, but we don't move to end
		}

		// goose-or-gosling <var-name> is put into list
		nextSExp.rest = &SExpListNode{first: Atom{atom: tokens[i + 1]}}

		// here is where it differs depending on goose or gosling
		// for goose we skip over just the var name and goose
		// gosling we skip over the Honk as well
		if tokenType == GOOSE {
			var sexp SExp
			sexp, i = parseSingleSExp(tokens, i + 2)

			// add in body of goose, finish with empty
			rest := SExpListNode{first: sexp, rest: &SExpListNode{first: Empty{}}}
			nextSExp.rest.rest = &rest
		} else {
			var sexpListNode SExpListNode
			if i + 2 >= length {
				return generateBadList(), length
			}

			if tokens[i + 2].tokenType != LPAREN {
				return generateBadList(), length
			}

			sexpListNode, i = parseSExpList(tokens, i + 3)
			if isBadList(sexpListNode) {
				return generateBadList(), i
			}

			nextSExp.rest.rest = &sexpListNode // we skip over SExpListNode, that's why we only have two rest's

			// now we need to check if HONK comes after the function, ie is an application
			if i < length {
				if tokens[i].tokenType == HONK {
					return honkParsing(tokens, i, nextSExp)
				}
			}

		}

		return nextSExp, i

	}
}

// turns the tokens into sexpressions
func parseSExp(tokens []Token) []SExp {
	// strategy: go through each of the tokens, grouping when possible
	// worst case scenario, all of the sexps are individual, we have to make out the same size of length, we'll deal with cutting it down afterwards

	length := len(tokens)
	sexpCounter := 0 // keep track of how many sexps we make
	out := make([]SExp, length)

	for i := 0; i < length; {
		out[sexpCounter], i = parseSingleSExp(tokens, i)
		sexpCounter++
	}

	return out[0:sexpCounter]
}

func printSExp(s SExp) {
	sType := s.getType()

	if sType == ATOM {
		fmt.Printf("ATOM(%s)", s.getAtom().atom.value)
	} else if sType == LIST {

		list := s.getSExpListNode()

		if isBadList(list) {
			fmt.Printf("Bad list...")
			return
		}

		fmt.Printf("LIST(")
		for elem := list.first; list.rest != nil && (list.rest.first.getType() != BADSEXP || list.rest.first.getType() != EMPTY); {
			printSExp(elem)
			elem = list.rest.first
			list = *list.rest
		}

		printSExp(list.first)
		fmt.Printf(")")
	} else if sType == BADSEXP {
		fmt.Printf("Bad Sexp...")
	}
}
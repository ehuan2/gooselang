package tokenizer

type SExpType int

const (
	ATOM    SExpType = 0 // ie a single thing, in our case it'll only be variables
	LIST    SExpType = 1 // ie the final AST will hold multiple ASTs, functions, applications and hatching
	BADSEXP SExpType = 2
)

// so now we use our tokens and parse them to expressions we can turn into asts
type SExp interface {
	getType() SExpType
}

func (s SExpListNode) getType() SExpType {
	return LIST
}

type SExpListNode struct {
	first SExp
	rest  *SExpListNode
}

func (a Atom) getType() SExpType {
	return ATOM
}

type Atom struct {
	atom Token
}

func (b BadSExp) getType() SExpType {
	return BADSEXP
}

type BadSExp struct{}


// given the list of tokens, as well as the position of the next int, parse through the sexp list
func parseSExpList(tokens []Token, position int) (SExpListNode, int) {
	// we keep iterating until we hit the RPAREN
	out := SExpListNode{}

	for position < len(tokens) {
		t := tokens[position]
	}

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
		} else if (tokenType == RPAREN || tokenType == ERROR) { // error cases
			out[sexpCounter] = BadSExp{}
			sexpCounter++

			out[sexpCounter] = BadSExp{}
			sexpCounter++
		} else { // list

			out[sexpCounter], sexpCounter = parseSExpList(tokens, sexpCounter)
			sexpCounter++

		}

	}

	return out[0:sexpCounter]

}

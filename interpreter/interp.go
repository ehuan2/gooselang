package interpreter

import "gooselang/AST"

type bouncerType int

const (
	BOUNCER_INTERP bouncerType = 0
	BOUNCER_CONT   bouncerType = 1
)

type bouncer interface {
	getType() bouncerType
}

func (i interpBouncer) getType() bouncerType {
	return BOUNCER_INTERP
}

func (cont applyContBouncer) getType() bouncerType {
	return BOUNCER_CONT
}

type interpBouncer struct{}
type applyContBouncer struct{}

func trampoline(b bouncer) {

}

func Interp(ast AST.AST) {

}

func applyCont() {

}
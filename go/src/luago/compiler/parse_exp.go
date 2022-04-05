package parser

import ."luago/compiler/ast"
import ."luago/compiler/lexer"
import "luago/number"

func parseExpList(lexer *Lexer) []Exp{
	exps := make([]Exp,0,4)
	exps = append(exps,parseExp(lexer))
	for lexer.LookAhead()==TOKEN_SEP_COMMA{
		lexer.NextToken()
		exps = append(exps,parseExp(lexer))
	}
	return exps
}

func parseExp(lexer *Lexer) Exp{
	
}
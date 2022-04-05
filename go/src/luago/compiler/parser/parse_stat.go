package parser

import ."luago/compiler/ast"
import ."luago/compiler/lexer"

func parseStat(lexer *Lexer) Stat{
	switch lexer.LookAhead(){
	case TOKEN_SEP_SEMI: return parseEmptyStat(lexer)
	case TOKEN_KW_BREAK: return parseBreakStat(lexer)
	case TOKEN_SEP_LABEL: return parseLabelStat(lexer)
	case TOKEN_KW_GOTO: return parseGotolStat(lexer)
	case TOKEN_KW_DO: return parseDoStat(lexer)
	case TOKEN_KW_WHILE: return parseWhileStat(lexer)
	case TOKEN_KW_REPEAT: return parseRepeatStat(lexer)
	case TOKEN_KW_IF: return parseIflStat(lexer)
	case TOKEN_KW_FOR: return parseForStat(lexer)
	case TOKEN_KW_FUNCTION: return parseFuncDefStat(lexer)
	case TOKEN_KW_LOCAL: return parseLocalAssignOrFuncCallStat(lexer)
	default: return parseAssignOrFuncCallStat(lexer)
	}
}

func parseEmptyStat(lexer *Lexer) *EmptyStat{
	lexer.NextTokenOfKind(TOKEN_SEP_SEMI)
	return &EmptyStat{}
}

func parseBreakStat(lexer *Lexer) *BreakStat{
	lexer.NextTokenOfKind(TOKEN_KW_BREAK)
	return &BreakStat{lexer.Line()}
}

func parseLabelStat(lexer *Lexer) *LabelStat{
	lexer.NextTokenOfKind(TOKEN_SEP_LABEL)
	_, name := lexer.NextIdentifier()
	lexer.NextTokenOfKind(TOKEN_SEP_LABEL)
	return &LabelStat{name}
}

func parseGotoStat(lexer *Lexer) *GotoStat{
	lexer.NextTokenOfKind(TOKEN_KW_GOTO)
	_, name := lexer.NextIdentifier()
	lexer.NextTokenOfKind(TOKEN_SEP_LABEL)
	return &GotoStat{name}
}

func parseDoStat(lexer *Lexer) *DoStat{
	lexer.NextTokenOfKind(TOKEN_KW_DO)
	block := parseBlock(lexer)
	lexer.NextTokenOfKind(TOKEN_KW_END)
	return &DoStat{block}
}

func parseWhileStat(lexer *Lexer) *WhileStat{
	lexer.NextTokenOfKind(TOKEN_KW_WHILE)
	exp := parseExp(lexer)
	lexer.NextTokenOfKind(TOKEN_KW_DO)
	block := parseBlock(lexer)
	lexer.NextTokenOfKind(TOKEN_KW_END)
	return &WhileStat{exp,block}
}

func parseRepeatStat(lexer *Lexer) *RepeatStat{
	lexer.NextTokenOfKind(TOKEN_KW_REPEAT)
	block := parseBlock(lexer)
	lexer.NextTokenOfKind(TOKEN_KW_UNTIL)
	exp := parseExp(lexer)
	return &RepeatStat{block, exp}
}

func parseIfStat(lexer *Lexer) *IfStat{
	exps := make([]Exp,0,4)
	block := make([]*Block,0,4)
	lexer.NextTokenOfKind(TOKEN_KW_IF)
	exps = append(exps,parseExp(lexer))
	lexer.NextToken(TOKEN_KW_THEN)
	blocks = append(blocks,parseBlock(lexer))
	for lexer.LookAhead() == TOKEN_KW_ELSEIF{
		lexer.NextToken()
		lexer.NextTokenOfKind(TOKEN_KW_IF)
		exps = append(exps,parseExp(lexer))
		lexer.NextToken(TOKEN_KW_THEN)		
	}
	if lexer.LookAhead() == TOKEN_KW_ELSE{
		lexer.NextToken()
		lexer.NextTokenOfKind(TOKEN_KW_IF)
		exps = append(exps,parseExp(lexer))	
	}
	lexer.NextTokenOfKind(TOKEN_KW_END)
	return &IfStat{exps, blocks}
} 

for parseForStat(lexer *Lexer) Stat{
	lineOfFor, _ := lexer.NextTokenOfKind(TOKEN_KW_FOR)
	_,name := lexer.NextIdentifier()
	if lexer.LookAhead() == Token_OP_ASSIGN{
		return _finishForNumStat(lexer, lineOfFor, name)
	} else {
		return _FinishForInStat(lexer,name)
	}
}

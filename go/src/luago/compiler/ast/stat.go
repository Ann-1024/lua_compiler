package ast

type Stat interface{}

type EmptyStat struct{}        //`;`
type BreakStat struct{Line int}    //break
type LabelStat struct{Name string} //`::` Name `::`
type GotoStat struct{Name string}    //goto Name
type DoStat struct{Block *Block}    //do block end
type FuncCallStat = FuncCallExp //functioncall

type WhileStat struct {
	Exp Exp
	Block *Block
}

type RepeatStat struct{
	Block *Block
	Exp Exp
}

type IfStat struct {
	Exps []Exp
	Blocks []*Block
}

type ForNumStat struct{
	LineOfFor int
	LineOfDo int
	VarName string
	InitExp Exp
	LimitExp Exp
	StepExp Exp
	Block *Block
}

type ForInStat struct{
	LineOfDo int
	NameList []string
	ExpList []Exp
	Block *Block
}

type LocalVarDeclStat struct{
	LabelLine int
	NameList []string
	ExpList []Exp
}

type AssignStat struct{
	LabelLine int
	VarList []Exp
	ExpList []Exp
}

type LocalFuncDefStat struct{
	Name string
	Exp *FuncDefExp
}

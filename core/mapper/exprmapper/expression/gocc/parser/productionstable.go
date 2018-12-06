// Code generated by gocc; DO NOT EDIT.

package parser

import "github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/direction"

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Flogo	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Flogo : Expr	<<  >>`,
		Id:         "Flogo",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Flogo : TernaryExpr	<<  >>`,
		Id:         "Flogo",
		NTType:     1,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr : OrExpr	<<  >>`,
		Id:         "Expr",
		NTType:     2,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `OrExpr : OrExpr "||" AndExpr	<< direction.NewExpression(X[0], X[1], X[2]) >>`,
		Id:         "OrExpr",
		NTType:     3,
		Index:      4,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpression(X[0], X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `OrExpr : AndExpr	<<  >>`,
		Id:         "OrExpr",
		NTType:     3,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AndExpr : AndExpr "&&" ConditionalExpr	<< direction.NewExpression(X[0], X[1], X[2]) >>`,
		Id:         "AndExpr",
		NTType:     4,
		Index:      6,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpression(X[0], X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `AndExpr : ConditionalExpr	<<  >>`,
		Id:         "AndExpr",
		NTType:     4,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ConditionalExpr : ConditionalExpr RelOp AddExpr	<< direction.NewExpression(X[0], X[1], X[2]) >>`,
		Id:         "ConditionalExpr",
		NTType:     5,
		Index:      8,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpression(X[0], X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `ConditionalExpr : AddExpr	<<  >>`,
		Id:         "ConditionalExpr",
		NTType:     5,
		Index:      9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AddExpr : AddExpr AddOp MulExpr	<< direction.NewExpression(X[0], X[1], X[2]) >>`,
		Id:         "AddExpr",
		NTType:     6,
		Index:      10,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpression(X[0], X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `AddExpr : MulExpr	<<  >>`,
		Id:         "AddExpr",
		NTType:     6,
		Index:      11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MulExpr : MulExpr MulOp ParenthesesExpr	<< direction.NewExpression(X[0], X[1], X[2]) >>`,
		Id:         "MulExpr",
		NTType:     7,
		Index:      12,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpression(X[0], X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `MulExpr : ParenthesesExpr	<<  >>`,
		Id:         "MulExpr",
		NTType:     7,
		Index:      13,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ParenthesesExpr : ExprLiteral	<<  >>`,
		Id:         "ParenthesesExpr",
		NTType:     8,
		Index:      14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ParenthesesExpr : "(" Expr ")"	<< direction.NewExpressionField(X[1]) >>`,
		Id:         "ParenthesesExpr",
		NTType:     8,
		Index:      15,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpressionField(X[1])
		},
	},
	ProdTabEntry{
		String: `RelOp : "=="	<<  >>`,
		Id:         "RelOp",
		NTType:     9,
		Index:      16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RelOp : "!="	<<  >>`,
		Id:         "RelOp",
		NTType:     9,
		Index:      17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RelOp : "<"	<<  >>`,
		Id:         "RelOp",
		NTType:     9,
		Index:      18,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RelOp : "<="	<<  >>`,
		Id:         "RelOp",
		NTType:     9,
		Index:      19,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RelOp : ">"	<<  >>`,
		Id:         "RelOp",
		NTType:     9,
		Index:      20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RelOp : ">="	<<  >>`,
		Id:         "RelOp",
		NTType:     9,
		Index:      21,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AddOp : "+"	<<  >>`,
		Id:         "AddOp",
		NTType:     10,
		Index:      22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AddOp : "-"	<<  >>`,
		Id:         "AddOp",
		NTType:     10,
		Index:      23,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MulOp : "*"	<<  >>`,
		Id:         "MulOp",
		NTType:     11,
		Index:      24,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MulOp : "/"	<<  >>`,
		Id:         "MulOp",
		NTType:     11,
		Index:      25,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MulOp : "%!"(MISSING)	<<  >>`,
		Id:         "MulOp",
		NTType:     11,
		Index:      26,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Func : function_name "(" ArgsList ")"	<< direction.NewFunction(X[0], X[2]) >>`,
		Id:         "Func",
		NTType:     12,
		Index:      27,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewFunction(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Func : function_name "()"	<< direction.NewFunction(X[0], "") >>`,
		Id:         "Func",
		NTType:     12,
		Index:      28,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewFunction(X[0], "")
		},
	},
	ProdTabEntry{
		String: `ArgsList : ExprLiteral	<< direction.NewArgument(X[0]) >>`,
		Id:         "ArgsList",
		NTType:     13,
		Index:      29,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewArgument(X[0])
		},
	},
	ProdTabEntry{
		String: `ArgsList : ArgsList "," ArgsList	<< direction.NewArguments(X[0], X[2]) >>`,
		Id:         "ArgsList",
		NTType:     13,
		Index:      30,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewArguments(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `TernaryExpr : Expr "?" Expr ":" Expr	<< direction.NewTernaryExpression(X[0], X[2], X[4]) >>`,
		Id:         "TernaryExpr",
		NTType:     14,
		Index:      31,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewTernaryExpression(X[0], X[2], X[4])
		},
	},
	ProdTabEntry{
		String: `ExprLiteral : Literal	<< direction.NewLiteralExpr(X[0]) >>`,
		Id:         "ExprLiteral",
		NTType:     15,
		Index:      32,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewLiteralExpr(X[0])
		},
	},
	ProdTabEntry{
		String: `ExprLiteral : Func	<<  >>`,
		Id:         "ExprLiteral",
		NTType:     15,
		Index:      33,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Literal : Int	<< direction.NewIntLit(X[0]) >>`,
		Id:         "Literal",
		NTType:     16,
		Index:      34,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewIntLit(X[0])
		},
	},
	ProdTabEntry{
		String: `Literal : Float	<< direction.NewFloatLit(X[0]) >>`,
		Id:         "Literal",
		NTType:     16,
		Index:      35,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewFloatLit(X[0])
		},
	},
	ProdTabEntry{
		String: `Literal : Bool	<< direction.NewBool(X[0]) >>`,
		Id:         "Literal",
		NTType:     16,
		Index:      36,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewBool(X[0])
		},
	},
	ProdTabEntry{
		String: `Literal : DoubleQString	<< direction.NewDoubleQuoteStringLit(X[0]) >>`,
		Id:         "Literal",
		NTType:     16,
		Index:      37,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewDoubleQuoteStringLit(X[0])
		},
	},
	ProdTabEntry{
		String: `Literal : SingleQString	<< direction.NewSingleQuoteStringLit(X[0]) >>`,
		Id:         "Literal",
		NTType:     16,
		Index:      38,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewSingleQuoteStringLit(X[0])
		},
	},
	ProdTabEntry{
		String: `Literal : MappingRef	<< direction.NewMappingRef(X[0]) >>`,
		Id:         "Literal",
		NTType:     16,
		Index:      39,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewMappingRef(X[0])
		},
	},
	ProdTabEntry{
		String: `Literal : Nil	<< direction.NewNilLit(X[0]) >>`,
		Id:         "Literal",
		NTType:     16,
		Index:      40,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewNilLit(X[0])
		},
	},
	ProdTabEntry{
		String: `DoubleQString : doublequotes_string	<<  >>`,
		Id:         "DoubleQString",
		NTType:     17,
		Index:      41,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SingleQString : singlequote_string	<<  >>`,
		Id:         "SingleQString",
		NTType:     18,
		Index:      42,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Int : number	<<  >>`,
		Id:         "Int",
		NTType:     19,
		Index:      43,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MappingRef : argument	<<  >>`,
		Id:         "MappingRef",
		NTType:     20,
		Index:      44,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Bool : "true"	<<  >>`,
		Id:         "Bool",
		NTType:     21,
		Index:      45,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Bool : "false"	<<  >>`,
		Id:         "Bool",
		NTType:     21,
		Index:      46,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Float : float	<<  >>`,
		Id:         "Float",
		NTType:     22,
		Index:      47,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Nil : "nil"	<<  >>`,
		Id:         "Nil",
		NTType:     23,
		Index:      48,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Nil : "null"	<<  >>`,
		Id:         "Nil",
		NTType:     23,
		Index:      49,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
}

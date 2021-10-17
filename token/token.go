package token

import (
	"fmt"
)

type TokenType string

// The location of the token
type TokenLoc struct {
	LineNumber  int // The line at which the token is
	ColumnStart int // The index from the begining of the line at which the first character of the token is
	ColumnEnd   int // The index from the begining of the line at which the last character of the token is
}

type Token struct {
	Type    TokenType // The type of the token
	Literal string    // The literal value of the token
	Loc     *TokenLoc // The location of the token
}

// Supported TokenType(s)
const (
	// Special
	ILLEGAL = "ILLEGAL" // Any unsopported token
	EOF     = "EOF"     // End-Of-File Token
	// Identifiers
	IDENTIFIER = "IDENTIFIER"
	// Numbers
	NUMBER = "NUMBER"
	// Strings
	STRING = "STRING"
	// Mathematical Operators
	PLUS     = "PLUS"     // +
	MINUS    = "MINUS"    // -
	ASTERISK = "ASTERISK" // *
	POWER    = "POWER"    // ^
	MODULUS  = "MODULUS"  // %
	SLASH    = "SLASH"    // /
	// Assignment Mathematical Operators
	EQ_PLUS     = "EQ_PLUS"     // +=
	EQ_MINUS    = "EQ_MINUS"    // -=
	EQ_ASTERISK = "EQ_ASTERISK" // *=
	EQ_POWER    = "EQ_POWER"    // ^=
	EQ_MODULUS  = "EQ_MODULUS"  // %=
	EQ_SLASH    = "EQ_SLASH"    // /=
	// POSTFIX OPERATORS
	POS_PLUS  = "POS_PLUS"  // ++
	POS_MINUS = "POS_MINUS" // --
	// Boolean Operators
	AND    = "AND"    // &&
	OR     = "OR"     // ||
	EQ     = "EQ"     // ==
	NOT_EQ = "NOT_EQ" // !=
	// Spread Operator
	SPREAD = "SPREAD" // ...
	// Keywords
	LET       = "LET"
	VAR       = "VAR"
	CONST     = "CONST"
	UNDEFINED = "UNDEFINED"
	TYPEOF    = "TYPEOF"
	DO        = "DO"
	WHILE     = "WHILE"
	IF        = "IF"
	ELSE      = "ELSE"
	SWITCH    = "SWITCH"
	CASE      = "CASE"
	DEFAULT   = "DEFAULT"
	RETURN    = "RETURN"
	FUNCTION  = "FUNCTION"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	OF        = "OF" // for of
	IN        = "IN" // for in
	// Symbols
	LPAREN     = "LPAREN"     // (
	RPAREN     = "RPAREN"     // )
	LCBRACE    = "LCBRACE"    // {
	RCBRACE    = "RCBRACE"    // }
	DOT        = "DOT"        // .
	COMMA      = "COMMA"      // ,
	Q_MARK     = "Q_MARK"     // ?
	COLON      = "COLON"      // :
	SEMI_COLON = "SEMI_COLON" // ;
	LSBRACE    = "LSBRACE"    // [
	RSBRACE    = "RSBRACE"    // ]
	LESS       = "LESS"       // <
	GREATER    = "GREATER"    // >
	ASSIGN     = "ASSIGN"     // =
	TEMPLATE   = "TEMPLATE"   // $
	NOT        = "NOT"        // !
	NILL       = "NILL"
)

func (tk *Token) Print() {
	if tk.Loc != nil {
		fmt.Printf("#%d [%d:%d] -> <%s:%s>\n", tk.Loc.LineNumber, tk.Loc.ColumnStart, tk.Loc.ColumnEnd, tk.Type, "'"+tk.Literal+"'")
	}
}

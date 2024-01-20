package grammar

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type Program struct {
	Pos lexer.Position

	Lines []*BasicLine `@@*`

	// Map of BASIC labels to tokenized lines
	BasicTable map[int]*BasicLine
	// Map of real file line numbers to tokenized lines
	FileTable map[int]*BasicLine
}

type BasicLine struct {
	Pos lexer.Position

	Label int `@Number`

	Comment    *string      `( @Comment`
	Statements []*Statement `  | ( @@ ( ":" @@ )* ) )`
	EOL        string       `( EOL | EOF )`
}

type Statement struct {
	Pos lexer.Position

	Tokens []*StatementToken `( @@ )+`
}

type StatementToken struct {
	Pos lexer.Position

	BasicToken *string `( @BasicToken`
	Value      *Value  ` | @@ )`
	Trailing   *string `@Trailing?`
}

type Value struct {
	Pos lexer.Position

	Number   *float64 `  @Number`
	Variable *string  `| @Ident`
	String   *string  `| @String`
	// Parentheses-surrounded statements
	Subexpression *Statement `| "(" @@ ")"`
}

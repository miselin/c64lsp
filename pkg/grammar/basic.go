package grammar

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type BasicGrammar struct {
	lexer  *lexer.StatefulDefinition
	parser *participle.Parser[Program]
}

func NewGrammar() BasicGrammar {
	// fix BASIC tokens that have meaning in regex (e.g. '+')
	t := []string{}
	for _, token := range tokens {
		// don't match REM, it's a comment
		if token == "REM" {
			continue
		}
		t = append(t, "(?i)"+regexp.QuoteMeta(token))
	}

	retokens := fmt.Sprintf("(%s)", strings.Join(t, "|"))

	lexer := lexer.MustSimple([]lexer.SimpleRule{
		// C64 BASIC tokens
		{Name: "BasicToken", Pattern: retokens},
		// Comments - "REM ....."
		{Name: "Comment", Pattern: `(?i)rem[^\n]*`},
		// A quoted string
		{Name: "String", Pattern: `"(\\"|[^"])*"`},
		// A number - can be either float or integer
		{Name: "Number", Pattern: `[-+]?(\d*\.)?\d+`},
		// Variable names. $ and % suffixes are optional but meaningful (string, integer vars)
		// C64 BASIC actually only considers the first two characters of a variable name
		// Note also that variables cannot contain BASIC tokens in their name.
		{Name: "Ident", Pattern: `[a-zA-Z]+[%$]?`},
		// Trailing characters for e.g. PRINT
		{Name: "Trailing", Pattern: `[,;]`},
		// end-of-statement (multiple statements on one line)
		{Name: "EOS", Pattern: `[:]`},
		// miscellaneous whitespace characters
		{Name: "Punct", Pattern: `[-[!@#%^&*()+_={}\|:;"'<,>.?/]|]`},
		{Name: "whitespace", Pattern: `[ \t]+`},
		// end-of-line terminators
		{Name: "EOL", Pattern: `[\n\r]+`},
	})

	parser := participle.MustBuild[Program](
		participle.Lexer(lexer),
		participle.CaseInsensitive("Ident"),
		participle.CaseInsensitive("Comment"),
		participle.CaseInsensitive("BasicToken"),
		participle.Unquote("String"),
		participle.UseLookahead(2),
	)

	return BasicGrammar{
		lexer:  lexer,
		parser: parser,
	}
}

func (grammar *BasicGrammar) Parse(filename, code string) (*Program, error) {
	program, err := grammar.parser.ParseString(filename, code)
	if err != nil {
		return nil, err
	}

	program.BasicTable = make(map[int]*BasicLine)
	program.FileTable = make(map[int]*BasicLine)

	for _, cmd := range program.Lines {
		program.BasicTable[cmd.Label] = cmd
		program.FileTable[cmd.Pos.Line] = cmd
	}

	return program, nil
}

func DumpStatement(stmt *Statement, indent int) {
	istr := strings.Repeat(" ", indent)
	fmt.Printf("%sstmt\n", istr)

	for _, tok := range stmt.Tokens {
		if tok.BasicToken != nil {
			fmt.Printf("%stok @%d -> BASIC '%s'", istr, tok.Pos.Column, *tok.BasicToken)
		} else if tok.Value != nil {
			if tok.Value.Number != nil {
				fmt.Printf("%stok @%d -> number %f", istr, tok.Pos.Column, *tok.Value.Number)
			} else if tok.Value.String != nil {
				fmt.Printf("%stok @%d -> string '%s'", istr, tok.Pos.Column, *tok.Value.String)
			} else if tok.Value.Variable != nil {
				fmt.Printf("%stok @%d -> var %s", istr, tok.Pos.Column, *tok.Value.Variable)
			} else if tok.Value.Subexpression != nil {
				DumpStatement(tok.Value.Subexpression, indent+2)
			}
		}

		fmt.Printf("%s\n", istr)
	}
}

func DumpLine(line *BasicLine) {
	fmt.Printf("BASIC line %d\n", line.Label)

	if line.Comment != nil {
		fmt.Printf(" -> comment '%s'\n", *line.Comment)
	}

	for _, stmt := range line.Statements {
		DumpStatement(stmt, 2)
	}
}

func (program *Program) Dump() {
	for _, line := range program.Lines {
		DumpLine(line)
	}
}

func (program *Program) FindTextLine(line int) *BasicLine {
	// lexer lines are 1-based
	r, found := program.FileTable[line+1]
	if !found {
		return nil
	}

	return r
}

func (program *Program) FindBasicLine(line int) *BasicLine {
	r, found := program.BasicTable[line]
	if !found {
		return nil
	}

	return r
}

func (program *Program) FindBasicTokenAt(line, character int) *string {
	l := program.FindTextLine(line)
	if l == nil {
		return nil
	}

	for _, stmt := range l.Statements {
		for _, tok := range stmt.Tokens {
			if tok.BasicToken != nil {
				if tok.Pos.Column <= character && (tok.Pos.Column+len(*tok.BasicToken)) >= character {
					return tok.BasicToken
				}
			}
		}
	}

	return nil
}

package reference

import (
	"errors"
	"fmt"
	"strings"
)

var functionsMap = map[string]*BasicFunction{}
var functionsMapBuilt = false

var FunctionNotFound = errors.New("function does not exist")

func buildFunctionsMap() {
	if functionsMapBuilt {
		return
	}

	functionsMapBuilt = true

	for _, fn := range BasicFunctions {
		functionsMap[fn.Name] = &fn
	}
}

func GetFunctionDocs(fn string) (*BasicFunction, error) {
	buildFunctionsMap()

	ref, ok := functionsMap[fn]
	if !ok {
		return nil, FunctionNotFound
	}

	return ref, nil
}

func (fn *BasicFunction) Markdown() string {
	var sb strings.Builder

	name := strings.ToUpper(fn.Name)

	sb.WriteString(fmt.Sprintf("## %s\n\n", name))
	sb.WriteString(fmt.Sprintf("### TYPE: %s\n", fn.Type))
	sb.WriteString(fmt.Sprintf("### FORMAT: `%s`\n\n\n", fn.Type))
	sb.WriteString(fmt.Sprintf("**Action:** %s\n\n\n", fn.Action))
	sb.WriteString(fmt.Sprintf("**EXAMPLES of %s Function:**\n\n", name))
	sb.WriteString(fn.ExampleMarkdown)

	return sb.String()
}

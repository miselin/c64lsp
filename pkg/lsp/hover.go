package lsp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/miselin/c64lsp/pkg/reference"
	"github.com/sourcegraph/jsonrpc2"
)

func (h *lspHandler) handleTextDocumentHover(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (result any, err error) {
	if req.Params == nil {
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
	}

	var params HoverParams
	if err := json.Unmarshal(*req.Params, &params); err != nil {
		return nil, err
	}

	return h.hover(params.TextDocument.URI, &params)
}

func (h *lspHandler) hover(uri DocumentURI, params *HoverParams) (*Hover, error) {
	parsed, ok := h.parsed[params.TextDocument.URI]
	if !ok {
		return nil, fmt.Errorf("hover but no parsed file")
	}

	token := parsed.FindBasicTokenAt(params.Position.Line, params.Position.Character)
	if token == nil {
		// nothing to show
		return nil, nil
	}

	docs, err := reference.GetFunctionDocs(*token)
	if errors.Is(err, reference.FunctionNotFound) {
		// we just don't know this function
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("blah %v", token)
	}

	// TODO: range
	return &Hover{
		Contents: MarkupContent{Kind: "markdown", Value: docs.Markdown()},
	}, nil
}

package lsp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/sourcegraph/jsonrpc2"
)

func (h *lspHandler) handleTextDocumentDefinition(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (result any, err error) {
	if req.Params == nil {
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
	}

	var params DocumentDefinitionParams
	if err := json.Unmarshal(*req.Params, &params); err != nil {
		return nil, err
	}

	return h.definition(ctx, params.TextDocument.URI, &params)
}

func (h *lspHandler) definition(ctx context.Context, uri DocumentURI, params *DocumentDefinitionParams) ([]Location, error) {
	logger := zerolog.Ctx(ctx)

	logger.Info().Msgf("definition request: %#v", params)

	return nil, fmt.Errorf("definition not implemented")
}

/**
definition request: &lsp.DocumentDefinitionParams{TextDocumentPositionParams:lsp.TextDocumentPositionParams{TextDocument:lsp.TextDocumentIdentifier{URI:"file:///g%3A/Code/c64lsp/examples/maze.bas"}, Position:lsp.Position{Line:9, Character:29}}, WorkDoneProgressParams:lsp.WorkDoneProgressParams{WorkDoneToken:interface {}(nil)}, PartialResultParams:lsp.PartialResultParams{PartialResultToken:interface {}(nil)}}
**/

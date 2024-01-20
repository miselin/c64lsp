package lsp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/sourcegraph/jsonrpc2"
)

func (h *lspHandler) handleTextDocumentCompletion(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (result any, err error) {
	if req.Params == nil {
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
	}

	var params CompletionParams
	if err := json.Unmarshal(*req.Params, &params); err != nil {
		return nil, err
	}

	return h.completion(ctx, params.TextDocument.URI, &params)
}

func (h *lspHandler) completion(ctx context.Context, uri DocumentURI, params *CompletionParams) ([]CompletionItem, error) {
	logger := zerolog.Ctx(ctx)

	logger.Info().Msgf("completion request: %#v", params)

	return nil, fmt.Errorf("completion not implemented")
}

/**
completion request: &lsp.CompletionParams{TextDocumentPositionParams:lsp.TextDocumentPositionParams{TextDocument:lsp.TextDocumentIdentifier{URI:"file:///g%3A/Code/c64lsp/examples/maze.bas"}, Position:lsp.Position{Line:8, Character:4}}, CompletionContext:lsp.CompletionContext{TriggerKind:0, TriggerCharacter:(*string)(nil)}}
**/

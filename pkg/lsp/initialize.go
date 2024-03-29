package lsp

import (
	"context"
	"encoding/json"
	"path/filepath"

	"github.com/sourcegraph/jsonrpc2"
)

func (h *lspHandler) handleInitialize(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (result any, err error) {
	if req.Params == nil {
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
	}

	h.conn = conn

	var params InitializeParams
	if err := json.Unmarshal(*req.Params, &params); err != nil {
		return nil, err
	}

	rootPath, err := fromURI(params.RootURI)
	if err != nil {
		return nil, err
	}
	h.rootPath = filepath.Clean(rootPath)
	h.addFolder(rootPath)

	return InitializeResult{
		Capabilities: ServerCapabilities{
			TextDocumentSync:   TDSKFull,
			DefinitionProvider: true,
			HoverProvider:      true,
			CompletionProvider: &CompletionOptions{
				TriggerCharacters: []string{},
			},
		},
	}, nil
}

package lsp

// Golang structs and definitions for types defined by the Lanaguage Server Protocol
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/

type File struct {
	LanguageID string
	Text       string
	Version    int
}

// DocumentURI specifies the URI for a document.
type DocumentURI string

// InitializeParams defines parameters sent by the client when initializing the language server.
type InitializeParams struct {
	ProcessID    int                `json:"processId,omitempty"`
	RootURI      DocumentURI        `json:"rootUri,omitempty"`
	Capabilities ClientCapabilities `json:"capabilities,omitempty"`
	Trace        string             `json:"trace,omitempty"`
}

// ClientCapabilities outlines the capabilities that a language server client supports.
// Largely ignored by this language server implementation.
type ClientCapabilities struct{}

// InitializeResult is sent by the langauge server when it has finished initializing.
type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities,omitempty"`
}

// TextDocumentSyncKind specifies the kind of synchronization that should be used for a text document.
type TextDocumentSyncKind int

const (
	// TDSKNone specifies that no synchronization should be used.
	TDSKNone TextDocumentSyncKind = iota
	// TDSKFull specifies that the entire document should be sent, on every change.
	TDSKFull
	// TDSKIncremental specifies that only updates to the document should be sent.
	TDSKIncremental
)

// CompletionOptions configures client completion capability.
type CompletionOptions struct {
	ResolveProvider   bool     `json:"resolveProvider,omitempty"`
	TriggerCharacters []string `json:"triggerCharacters"`
}

// ServerCapabilities defines the capabilities of the language server.
type ServerCapabilities struct {
	TextDocumentSync   TextDocumentSyncKind `json:"textDocumentSync,omitempty"`
	CompletionProvider *CompletionOptions   `json:"completionProvider,omitempty"`
	DefinitionProvider bool                 `json:"definitionProvider,omitempty"`
	HoverProvider      bool                 `json:"hoverProvider,omitempty"`
}

// TextDocumentItem is an item to transfer a text document from the client to the server.
type TextDocumentItem struct {
	URI        DocumentURI `json:"uri"`
	LanguageID string      `json:"languageId"`
	Version    int         `json:"version"`
	Text       string      `json:"text"`
}

// VersionedTextDocumentIdentifier identifies a specific version of a text document.
type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version int `json:"version"`
}

// TextDocumentIdentifier identifies a text document.
type TextDocumentIdentifier struct {
	URI DocumentURI `json:"uri"`
}

// DidOpenTextDocumentParams defines parameters sent from the client when a text document is opened.
type DidOpenTextDocumentParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}

// DidCloseTextDocumentParams defines parameters sent from the client when a text document is closed.
type DidCloseTextDocumentParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// TextDocumentContentChangeEvent defines a change to a text document.
type TextDocumentContentChangeEvent struct {
	Range       Range  `json:"range"`
	RangeLength int    `json:"rangeLength"`
	Text        string `json:"text"`
}

// DidChangeTextDocumentParams defines parameters sent from the client when a text document is changed.
type DidChangeTextDocumentParams struct {
	TextDocument   VersionedTextDocumentIdentifier  `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

// DidSaveTextDocumentParams defines parameters sent from the client when a text document is saved.
type DidSaveTextDocumentParams struct {
	Text         *string                `json:"text"`
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// TextDocumentPositionParams defines a position in a text document for use in other parameters.
type TextDocumentPositionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position     Position               `json:"position"`
}

// CompletionParams defines parameters to be sent when requesting completion.
type CompletionParams struct {
	TextDocumentPositionParams
	CompletionContext CompletionContext `json:"contentChanges"`
}

// CompletionContext defines the context for a completion request.
type CompletionContext struct {
	TriggerKind      int     `json:"triggerKind"`
	TriggerCharacter *string `json:"triggerCharacter"`
}

// HoverParams defines parameters to be sent when requesting hover information.
type HoverParams struct {
	TextDocumentPositionParams
}

// Location defines a location in a document.
type Location struct {
	URI   DocumentURI `json:"uri"`
	Range Range       `json:"range"`
}

// Range defines a range of positions.
type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

// Position defines a position by line and character.
type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

// TextEdit defines a text edit to be applied to a document.
type TextEdit struct {
	Range   Range  `json:"range"`
	NewText string `json:"newText"`
}

// CompletionItemKind defines completion kinds for completion items.
type CompletionItemKind int

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionItemKind
const (
	TextCompletion          CompletionItemKind = 1
	MethodCompletion        CompletionItemKind = 2
	FunctionCompletion      CompletionItemKind = 3
	ConstructorCompletion   CompletionItemKind = 4
	FieldCompletion         CompletionItemKind = 5
	VariableCompletion      CompletionItemKind = 6
	ClassCompletion         CompletionItemKind = 7
	InterfaceCompletion     CompletionItemKind = 8
	ModuleCompletion        CompletionItemKind = 9
	PropertyCompletion      CompletionItemKind = 10
	UnitCompletion          CompletionItemKind = 11
	ValueCompletion         CompletionItemKind = 12
	EnumCompletion          CompletionItemKind = 13
	KeywordCompletion       CompletionItemKind = 14
	SnippetCompletion       CompletionItemKind = 15
	ColorCompletion         CompletionItemKind = 16
	FileCompletion          CompletionItemKind = 17
	ReferenceCompletion     CompletionItemKind = 18
	FolderCompletion        CompletionItemKind = 19
	EnumMemberCompletion    CompletionItemKind = 20
	ConstantCompletion      CompletionItemKind = 21
	StructCompletion        CompletionItemKind = 22
	EventCompletion         CompletionItemKind = 23
	OperatorCompletion      CompletionItemKind = 24
	TypeParameterCompletion CompletionItemKind = 25
)

// CompletionItemTag defines tags for completion items, such as deprecation tags.
type CompletionItemTag int

// InsertTextFormat defines the format of a completion.
type InsertTextFormat int

const (
	// PlainTextTextFormat is a plain string to be inserted.
	PlainTextTextFormat InsertTextFormat = 1
	// SnippetTextFormat is a snippet string to be inserted, and can include placeholders.
	SnippetTextFormat InsertTextFormat = 2
)

// CompletionItem is
type CompletionItem struct {
	Label               string              `json:"label"`
	Kind                CompletionItemKind  `json:"kind,omitempty"`
	Tags                []CompletionItemTag `json:"tags,omitempty"`
	Detail              string              `json:"detail,omitempty"`
	Documentation       string              `json:"documentation,omitempty"` // string | MarkupContent
	Deprecated          bool                `json:"deprecated,omitempty"`
	Preselect           bool                `json:"preselect,omitempty"`
	SortText            string              `json:"sortText,omitempty"`
	FilterText          string              `json:"filterText,omitempty"`
	InsertText          string              `json:"insertText,omitempty"`
	InsertTextFormat    InsertTextFormat    `json:"insertTextFormat,omitempty"`
	TextEdit            *TextEdit           `json:"textEdit,omitempty"`
	AdditionalTextEdits []TextEdit          `json:"additionalTextEdits,omitempty"`
	CommitCharacters    []string            `json:"commitCharacters,omitempty"`
	Data                any                 `json:"data,omitempty"`
}

// Hover defines the response to a hover request.
type Hover struct {
	Contents any    `json:"contents"`
	Range    *Range `json:"range"`
}

// MarkedString defines either human-readable or code content, depending on the language.
type MarkedString struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}

// MarkupKind defines the type of content in a MarkupContent.
type MarkupKind string

const (
	// PlainText is plain text.
	PlainText MarkupKind = "plaintext"
	// Markdown is markdown.
	Markdown MarkupKind = "markdown"
)

// MarkupContent defines a content string and its kind.
type MarkupContent struct {
	Kind  MarkupKind `json:"kind"`
	Value string     `json:"value"`
}

// DidChangeConfigurationParams defines parameters sent from the client when the configuration changes.
type DidChangeConfigurationParams struct {
	Settings any `json:"settings"`
}

// DocumentDefinitionParams defines parameters sent from the client when requesting a document definition.
type DocumentDefinitionParams struct {
	TextDocumentPositionParams
}

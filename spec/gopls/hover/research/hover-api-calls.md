# Gopls Hover API Calls Documentation

## Overview

Gopls implements the Language Server Protocol (LSP) `textDocument/hover` request to provide information about code elements when hovering over them in an editor. The hover functionality is implemented in `gopls/internal/golang/hover.go` and is part of the server's request handling system.

## LSP Protocol Details

### Request Parameters
The hover request follows the LSP specification with the method name `textDocument/hover` and parameters:

```json
{
  "textDocument": {
    "uri": "file:///path/to/file.go"
  },
  "position": {
    "line": 10,
    "character": 5
  }
}
```

### Response Format
The response follows the LSP Hover structure:

```json
{
  "contents": {
    "kind": "markdown",
    "value": "Documentation content"
  },
  "range": {
    "start": {"line": 10, "character": 5},
    "end": {"line": 10, "character": 10}
  }
}
```

## Implementation Details

### File Location
The hover implementation is located in:
```
gopls/internal/golang/hover.go
```

### Key Components

1. **Hover Function**: The main hover function processes the request and returns hover information
2. **HoverInfo Structure**: Contains the structured information about the hovered symbol
3. **findExamples() Function**: Helper function (proposed) to extract example code for display in hover

### Data Sources
The hover implementation relies on:
- Abstract Syntax Tree (AST) of the file
- Type information for the file and all dependencies
- Documentation comments
- Struct field size and offset information
- Struct tags

## Features

### Information Displayed
When hovering over code elements, gopls can display:
- Symbol name, kind, and type signature
- Constant values (for constants)
- Abbreviated declarations (for types)
- Documentation comments
- Links to pkg.go.dev documentation
- Struct field size and offset information
- Struct tags
- Package-level constants
- Example code (when implemented)

### Configuration Options

#### hoverKind
Controls the information displayed in hover text:
- `FullDocumentation`: Complete documentation
- `NoDocumentation`: No documentation
- `SingleLine`: Single line only (for editor plugins)
- `SynopsisDocumentation`: Synopsis only

#### linksInHover
Controls presence of documentation links in hover markdown:
- `FullDocumentation`
- `NoDocumentation`
- `SingleLine`
- `SynopsisDocumentation`

#### linkTarget
Base URL for links to Go package documentation:
- `"pkg.go.dev"` (default)
- `"godoc.org"`
- Custom documentation site

## Limitations

### Position-only Requests
The LSP specification only provides a position (not a selection) for hover requests, which means gopls cannot provide information about sub-expressions like `f(x).y` without additional selection data.

### Structured Hover Removal
The experimental `Structured` hover kind that returned JSON was removed in gopls v0.18.0 as it was not intended for human-readable content.

## Integration with Editor Features

### Supported Editors
Gopls hover functionality works with any LSP-compatible editor:
- VS Code (with Go extension)
- Vim (with coc.nvim or similar)
- Emacs (with eglot or lsp-mode)
- Sublime Text (with LSP package)
- Other LSP-compatible editors

### Client Capabilities
Clients can request either plain text or Markdown format for hover content.

## Future Enhancements

### Example Code in Hover
A feature request exists to render example tests as part of hover documentation, which would involve implementing a `findExamples()` function within the hover handler.

## Technical Architecture

### Request Flow
1. Client sends `textDocument/hover` request with position
2. Server package receives and routes request
3. Golang package hover handler processes the request
4. AST and type information are used to gather symbol information
5. HoverInfo structure is populated
6. Response is formatted according to client's requested content type
7. Response is sent back to client

### Dependencies
The hover implementation depends on:
- `golang.org/x/tools/go/packages` for package loading
- `golang.org/x/tools/go/ast` for AST parsing
- `golang.org/x/tools/go/types` for type checking
- Internal gopls caching and session management systems
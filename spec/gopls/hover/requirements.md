## Overview

Implement hover documentation capability in MCP server to retrieve and display documentation for Go symbols using `gopls` integration

## Acceptance Criteria

### Functional Requirements

1. Function symbol hover

Given a Go source file is open in the MCP server
When the user hovers over a function symbol
Then the server retrieves and displays the function's documentation, including its signature and comments.

2. Type symbol hover

Given a Go source file is open in the MCP server
When the user hovers over a type symbol (struct, interface, etc.)
Then the server retrieves and displays the type's documentation, including its definition and comments.

3. Constant and Variable symbol hover

Given a Go source file is open in the MCP server
When the user hovers over a variable symbol
Then the server retrieves and displays the variable's documentation, including its type and comments.

4. Invalid symbol hover

Given a Go source file is open in the MCP server
When the user hovers over an invalid symbol: whitespace, comments, operators, keywords
Then the server responds with an appropriate message indicating that no documentation is available.

5. Unavailable documentation

Given a Go source file is open in the MCP server
When the user hovers over a symbol for which documentation is not available
Then the server responds with an appropriate message indicating that no documentation is available.

6. External package symbol hover

Given a Go source file is open in the MCP server
When the user hovers over a symbol from an external package
Then the server retrieves and displays the documentation for that symbol, including its definition and comments.

### Non-Functional Requirements

1. Error Handling

- gopls startup failure
- invalid file path or position
- timeouts 5s when waiting for gopls response
- malformed requests: invalid position

2. Compatibility

Test with gopls versions 0.20.0 but still support extension to future versions

3. Logging and observability

- Log all hover requests with file path and position
- Log errors with detailed messages
- Support debug mode for verbose logging

## Problems

1. User may change go files during a prompt session
   Solution: Use `textDocument/didChange` notifications to keep gopls in sync with the latest file content.

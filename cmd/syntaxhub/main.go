package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{Name: "lsp-servers"}, nil)

	// Using the generic AddTool automatically populates the input and output
	// schema of the tool.
	type args struct {
		Name string `json:"name" jsonschema:"the person to greet"`
	}

	mcp.AddTool(server, &mcp.Tool{
		Name:        "greet",
		Description: "say hi",
	}, func(_ context.Context, _ *mcp.CallToolRequest, args args) (*mcp.CallToolResult, any, error) {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Hi " + args.Name},
			},
		}, nil, nil
	})

	// Run the server on the stdio transport.
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Printf("Server failed: %v", err)
	}
}

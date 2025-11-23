package gopls_test

import (
	"context"
	"log"
	"strings"
	"testing"

	"github.com/khaitranhq/syntaxhub/test/e2e"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestHoverFunction(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	mcpClient, err := e2e.NewMcpClient(ctx)
	if err != nil {
		t.Fatalf("failed to create MCP client: %v", err)
	}

	params := &mcp.CallToolParams{
		Name:      "greet",
		Arguments: map[string]any{"name": "you"},
	}

	res, err := mcpClient.Session.CallTool(ctx, params)
	if err != nil {
		log.Fatalf("CallTool failed: %v", err)
	}

	messages := strings.Builder{}

	for _, content := range res.Content {
		str, err := content.MarshalJSON()
		if err != nil {
			log.Fatalf("Failed to marshal content: %v", err)
		}

		messages.Write(str)
	}

	log.Printf("CallTool result: %+v", messages.String())
}

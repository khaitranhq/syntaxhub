package gopls_test

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/khaitranhq/syntaxhub/test/e2e"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/samber/lo"
)

//nolint:cyclop
func TestHoverFunction(t *testing.T) {
	t.Parallel()

	goFilePath := filepath.Join("example", "main.go")

	if _, err := os.Stat(goFilePath); os.IsNotExist(err) {
		t.Fatalf("example file not found: %v", err)
	}

	ctx := context.Background()

	mcpClient, err := e2e.NewMcpClient(ctx)
	if err != nil {
		t.Fatalf("failed to create MCP client: %v", err)
	}

	defer func() {
		if err := mcpClient.Close(); err != nil {
			t.Errorf("failed to close MCP client: %v", err)
		}
	}()

	absGoFilePath, err := filepath.Abs(goFilePath)
	if err != nil {
		t.Fatalf("failed to get absolute path: %v", err)
	}

	hoverParams := map[string]any{
		"file_path": absGoFilePath,
		"position": map[string]any{
			"line":      14, // 0-indexed line number (line 3 in file)
			"character": 12, // Position in the "Add" function name
		},
	}

	callToolParams := &mcp.CallToolParams{
		Name:      "hover",
		Arguments: hoverParams,
	}

	result, err := mcpClient.CallTool(ctx, callToolParams)
	if err != nil {
		t.Fatalf("failed to call hover tool: %v", err)
	}

	if result == nil {
		t.Fatal("expected non-nil result from hover tool")
	}

	if len(result.Content) == 0 {
		t.Fatal("expected result to contain content, got empty content")
	}

	hoverTexts := lo.Map(result.Content, func(content mcp.Content, _ int) string {
		textContent, ok := content.(*mcp.TextContent)
		if !ok {
			t.Fatalf("expected TextContent type, got %T", content)
		}

		return textContent.Text
	})

	if len(hoverTexts) == 0 {
		t.Fatal("expected at least one hover text, got none")
	}

	hoverText := strings.Join(hoverTexts, "\n")

	// Verify that the hover text contains documentation
	expectedStrings := []string{
		"Add",               // Function name
		"adds two integers", // Part of the documentation
		"int",               // Type information
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(hoverText, expected) {
			t.Errorf("expected hover text to contain %q, got: %s", expected, hoverText)
		}
	}

	t.Logf("Hover text received: %s", hoverText)
}

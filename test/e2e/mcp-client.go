package e2e

import (
	"context"
	"fmt"
	"os/exec"
	"path"
	"time"

	"github.com/khaitranhq/syntaxhub/internal/utils/filesystem"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const (
	DefaultGoBuildTimeoutSeconds = 10
)

// McpClient is a client for interacting with an MCP server.
type McpClient struct {
	Session *mcp.ClientSession
}

// NewMcpClient creates a new MCP client and connects to the MCP server.
func NewMcpClient(ctx context.Context) (*McpClient, error) {
	client := mcp.NewClient(&mcp.Implementation{
		Name: "mcp-lsp-client",
	}, nil)

	goModRootPath, err := filesystem.GetGoModuleRootPath(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get Go module root path: %w", err)
	}

	buildTargetPath := path.Join(goModRootPath, "build", "syntaxhub")

	timeoutCtx, cancel := context.WithTimeout(ctx, DefaultGoBuildTimeoutSeconds*time.Second)
	defer cancel()

	_, err = exec.CommandContext(
		timeoutCtx,
		"go",
		"build",
		"-o",
		buildTargetPath,
		path.Join(goModRootPath, "cmd/syntaxhub/main.go"),
	).
		Output()
	if err != nil {
		return nil, fmt.Errorf("failed to build MCP server binary: %w", err)
	}

	mcpSession, err := client.Connect(
		ctx,
		&mcp.CommandTransport{Command: exec.CommandContext(ctx, buildTargetPath)},
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MCP server: %w", err)
	}

	return &McpClient{
		Session: mcpSession,
	}, nil
}

// Close closes the MCP client session.
func (c *McpClient) Close() error {
	err := c.Session.Close()
	if err != nil {
		return fmt.Errorf("failed to close MCP client session: %w", err)
	}

	return nil
}

// CallTool calls a tool on the MCP server with the given parameters.
func (c *McpClient) CallTool(
	ctx context.Context,
	params *mcp.CallToolParams,
) (*mcp.CallToolResult, error) {
	callToolResult, err := c.Session.CallTool(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("CallTool failed: %w", err)
	}

	return callToolResult, nil
}

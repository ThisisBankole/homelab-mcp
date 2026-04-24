package main

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func handleJournalLogs(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()
	service, _ := args["service"].(string)
	lines := "50"
	if n, ok := args["lines"].(float64); ok {
		lines = fmt.Sprintf("%.0f", n)
	}
	out, err := runCommand("journalctl", "-u", service, "-n", lines, "--no-pager")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error: %s", err)), nil
	}
	return mcp.NewToolResultText(out), nil
}
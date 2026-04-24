package main

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
)

func runCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s: %w", strings.TrimSpace(string(out)), err)
	}
	return strings.TrimSpace(string(out)), nil
}

func handleDiskUsage(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()
	path := "/"
	if p, ok := args["path"].(string); ok && p != "" {
		path = p
	}
	out, err := runCommand("df", "-h", path)
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error: %s", err)), nil
	}
	return mcp.NewToolResultText(out), nil
}

func handleMemoryUsage(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	out, err := runCommand("free", "-h")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error: %s", err)), nil
	}
	return mcp.NewToolResultText(out), nil
}

func handleUptime(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	out, err := runCommand("uptime")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error: %s", err)), nil
	}
	return mcp.NewToolResultText(out), nil
}
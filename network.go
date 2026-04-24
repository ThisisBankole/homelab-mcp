package main

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func handleOpenPorts(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	out, err := runCommand("ss", "-tlnp")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error: %s", err)), nil
	}
	return mcp.NewToolResultText(out), nil
}

func handlePing(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()
	host, _ := args["host"].(string)
	out, err := runCommand("ping", "-c", "4", host)
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error: %s", err)), nil
	}
	return mcp.NewToolResultText(out), nil
}
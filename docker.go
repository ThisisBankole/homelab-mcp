package main

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func handleListContainers(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()
	all, _ := args["all"].(bool)
	format := "table {{.Names}}\t{{.Status}}\t{{.Ports}}"
	cmdArgs := []string{"ps", "--format", format}
	if all {
		cmdArgs = []string{"ps", "-a", "--format", format}
	}
	out, err := runCommand("docker", cmdArgs...)
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error: %s", err)), nil
	}
	return mcp.NewToolResultText(out), nil
}

func handleContainerLogs(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()
	name, _ := args["name"].(string)
	lines := "50"
	if n, ok := args["lines"].(float64); ok {
		lines = fmt.Sprintf("%.0f", n)
	}
	out, err := runCommand("docker", "logs", "--tail", lines, name)
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error: %s", err)), nil
	}
	return mcp.NewToolResultText(out), nil
}

func handleRestartContainer(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()
	name, _ := args["name"].(string)
	out, err := runCommand("docker", "restart", name)
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error: %s", err)), nil
	}
	return mcp.NewToolResultText(out), nil
}

func handleContainerStats(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	out, err := runCommand("docker", "stats", "--no-stream", "--format",
		"table {{.Name}}\t{{.CPUPerc}}\t{{.MemUsage}}")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error: %s", err)), nil
	}
	return mcp.NewToolResultText(out), nil
}
package main

import (
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer("homelab-mcp", "1.0.0")

	// Docker tools
	s.AddTool(mcp.NewTool("list_containers",
		mcp.WithDescription("List Docker containers"),
		mcp.WithBoolean("all", mcp.Description("Include stopped containers")),
	), handleListContainers)

	s.AddTool(mcp.NewTool("container_logs",
		mcp.WithDescription("Get logs for a container"),
		mcp.WithString("name", mcp.Required(), mcp.Description("Container name")),
		mcp.WithNumber("lines", mcp.Description("Lines to tail, default 50")),
	), handleContainerLogs)

	s.AddTool(mcp.NewTool("restart_container",
		mcp.WithDescription("Restart a Docker container by name"),
		mcp.WithString("name", mcp.Required(), mcp.Description("Container name")),
	), handleRestartContainer)

	s.AddTool(mcp.NewTool("container_stats",
		mcp.WithDescription("Show CPU and memory usage per container"),
	), handleContainerStats)

	// System tools
	s.AddTool(mcp.NewTool("disk_usage",
		mcp.WithDescription("Show disk usage across mounts"),
		mcp.WithString("path", mcp.Description("Specific path, default /")),
	), handleDiskUsage)

	s.AddTool(mcp.NewTool("memory_usage",
		mcp.WithDescription("Show RAM and swap usage"),
	), handleMemoryUsage)

	s.AddTool(mcp.NewTool("system_uptime",
		mcp.WithDescription("Show system uptime and load average"),
	), handleUptime)

	// Network tools
	s.AddTool(mcp.NewTool("open_ports",
		mcp.WithDescription("List open listening ports"),
	), handleOpenPorts)

	s.AddTool(mcp.NewTool("ping_host",
		mcp.WithDescription("Ping a host to check connectivity"),
		mcp.WithString("host", mcp.Required(), mcp.Description("Hostname or IP")),
	), handlePing)

	// Logs
	s.AddTool(mcp.NewTool("journal_logs",
		mcp.WithDescription("Get systemd journal logs for a service"),
		mcp.WithString("service", mcp.Required(), mcp.Description("Service name e.g. cloudflared")),
		mcp.WithNumber("lines", mcp.Description("Lines to return, default 50")),
	), handleJournalLogs)

	httpServer := server.NewStreamableHTTPServer(s)

	log.Println("homelab-mcp listening on :8082")
	log.Fatal(httpServer.Start(":8082"))
}

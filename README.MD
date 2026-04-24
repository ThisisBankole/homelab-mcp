# homelab-mcp

A remote MCP (Model Context Protocol) server that lets Claude control and monitor my homelab via natural language.

## What it does

Connects Claude to my homelab so I can ask things like:

- _"Is Plex running?"_
- _"Show me the last 50 lines of Sonarr logs"_
- _"How much disk space is left?"_
- _"What ports are open?"_

Claude calls the right tool, runs the command on arrakis, and replies in plain English.

## Tools

| Tool | Command | Description |
|---|---|---|
| `list_containers` | `docker ps` | List running/stopped containers |
| `container_logs` | `docker logs` | Get logs for a container |
| `restart_container` | `docker restart` | Restart a container by name |
| `container_stats` | `docker stats` | CPU and memory per container |
| `disk_usage` | `df -h` | Disk usage across mounts |
| `memory_usage` | `free -h` | RAM and swap usage |
| `system_uptime` | `uptime` | Uptime and load average |
| `open_ports` | `ss -tlnp` | Listening ports |
| `ping_host` | `ping` | Check connectivity to a host |
| `journal_logs` | `journalctl` | Systemd service logs |

## Stack

- **Language:** Go
- **Library:** `github.com/mark3labs/mcp-go`
- **Transport:** Streamable HTTP
- **Deployment:** Docker via Portainer


## Local development

```bash
# Run locally
make run

# Build binary
make build

# Deploy to arrakis
make deploy
```

## Deployment

Code is pulled from GitHub by Portainer. To redeploy after a push:

1. Push changes to `main`
2. Portainer → Stacks → `homelab-mcp` → Pull and redeploy

## How it connects to Claude

1. Go to Claude.ai → Settings → Integrations
2. Add MCP server: use your domain
3. Claude fetches the tools list and can call them in any conversation

## Security

- Commands are never passed as raw shell strings — all args go through `exec.Command` directly
- Only the tools defined in Go are accessible — no arbitrary command execution
- Docker socket is mounted read/write (required for container management)
- Exposed only via Cloudflare Tunnel — no direct port exposure to the internet
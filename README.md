# JAWAD.OS — Bloomberg-Inspired Engineering Workstation

<div align="center">

```text
  ██╗ █████╗ ██╗    ██╗ █████╗ ██████╗       ██████╗ ███████╗
  ██║██╔══██╗██║    ██║██╔══██╗██╔══██╗     ██╔═══██╗██╔════╝
  ██║███████║██║ █╗ ██║███████║██║  ██║     ██║   ██║███████╗
████║██╔══██║██║███╗██║██╔══██║██║  ██║     ██║   ██║╚════██║
╚███║██║  ██║╚███╔███╔╝██║  ██║██████╔╝     ╚██████╔╝███████║
 ╚══╝╚═╝  ╚═╝ ╚══╝╚══╝ ╚═╝  ╚═╝╚═════╝       ╚═════╝ ╚══════╝
```

**An interactive terminal developer portfolio engineered with the visual authority of a Bloomberg Terminal.**  
*Built with Go, Charm Bubble Tea, Lip Gloss, Bubbles, and Wish SSH.*

[![Go Version](https://img.shields.io/badge/Go-1.27+-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![Bubble Tea](https://img.shields.io/badge/TUI-Bubble%20Tea-7D56F4?style=flat-square)](https://github.com/charmbracelet/bubbletea)
[![Lip Gloss](https://img.shields.io/badge/Styling-Lip%20Gloss-EE799F?style=flat-square)](https://github.com/charmbracelet/lipgloss)
[![Wish SSH](https://img.shields.io/badge/Transport-Wish%20SSH-50fa7b?style=flat-square)](https://github.com/charmbracelet/wish)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](LICENSE)

</div>

---

## Remote Access (Zero Installation Required)

Visitors can connect directly to JAWAD.OS from any terminal worldwide over SSH:

```bash
ssh jawad.example.com -p 2222
```

*(Connects to an isolated, multi-tenant Bubble Tea runtime)*

---

## 1. Design Identity: Institutional Engineering Workstation

- **Developer:** Jawad Ahmad
- **Role:** Aspiring Full-Stack Software Engineer
- **Location:** Pakistan
- **Education:** Computer Science Student
- **Tagline:** *Building software. Exploring systems. Understanding how things work.*

**JAWAD.OS** combines the visual hierarchy and high information density of a Bloomberg Terminal with the precision of modern developer tools. It rejects neon/matrix hacker clichés in favor of sharp single-line geometry, structured data tables, and an authoritative institutional color system.

### Color System Tokens

| Token | Hex | Role | Usage |
| :--- | :--- | :--- | :--- |
| **Canvas** | `#101214` | Base Background | Terminal background canvas |
| **Surface** | `#181B1E` | Primary Surface | Data panels, technical dossiers |
| **Surface Elevated** | `#202428` | Highlight Surface | Active table rows, modals |
| **Border** | `#343A40` | Structural Grid | Single-line borders (`┌─┐`) and dividers |
| **Primary Text** | `#E6E8E9` | Readability | Core copy, values, labels |
| **Secondary Text** | `#92999F` | Metadata | Timestamps, secondary labels, shortcuts |
| **Accent Amber** | `#F0A52B` | Primary Highlight | Active navigation, section tags, tickers |
| **Success Green** | `#76B88A` | Semantic Status | `ONLINE ●`, completed records |
| **Info Cyan** | `#67B8C8` | Technical Meta | Stack badges, pipeline flows |
| **Warning Orange** | `#D97757` | Semantic Alert | Development status, errors, warnings |

---

## 2. Multi-Panel Workstation Layout

```text
┌────────────────────────────────────────────────────────────────────────────────────────┐
│ JAWAD.OS | ENGINEERING WORKSTATION                           ONLINE ● | 02:00:00 UTC   │
├────────────────────────────────────────────────────────────────────────────────────────┤
│ [01 PROFILE]   02 PROJECTS   03 STACK   04 JOURNEY   05 INTERESTS   06 CONTACT         │
├───────────────────────────────────────────────────────────────────┬────────────────────┤
│ PRIMARY WORKSPACE (70%)                                           │ TELEMETRY (30%)    │
│                                                                   │                    │
│ JAWAD AHMAD                                                       │ LOCATION: PK       │
│ Software Engineer in Progress                                     │ EDU: CS Student    │
│                                                                   │ FOCUS: Full-Stack  │
│ CURRENT FOCUS:                                                    │ STACK: Go, TS, Py  │
│ ▸ Backend Engineering  ▸ System Design  ▸ Distributed Systems    ├────────────────────┤
│                                                                   │ SESSION INFO       │
│ FEATURED PROJECT: WORKFLOW AUTOMATION PLATFORM                    │ User: guest        │
│ Status: IN DEVELOPMENT  |  Engine: Go (DAG + Redis Task Queue)     │ Term: xterm-256    │
│ Concept: Trigger -> Execution -> AI Summary -> Webhook Dispatch   │ Host: local/wish   │
├───────────────────────────────────────────────────────────────────┴────────────────────┤
│ COMMAND > _                                                                            │
├────────────────────────────────────────────────────────────────────────────────────────┤
│ F1/? HELP | ↑↓ NAVIGATE | ENTER SELECT | : COMMAND | ESC BACK | Q EXIT                 │
└────────────────────────────────────────────────────────────────────────────────────────┘
```

- **Wide Terminals (`>= 98` columns):** Activates dual-pane layout (~68% primary workspace / ~32% contextual telemetry and directory).
- **Compact Terminals (`< 98` columns):** Gracefully stacks panels vertically without horizontal clipping or truncation.

---

## 3. Workstation Modules

| Module | Description | Key Interactions |
| :--- | :--- | :--- |
| **Boot Sequence** | Simulated hardware initialization sequence with animated step ticks | Press `ENTER` / `SPACE` to skip |
| **Dashboard** | 70/30 multi-panel overview: Identity, focus quadrant, featured project, and telemetry | `↑`/`↓`, `1`-`6`, `ENTER` |
| **01 PROFILE** | Personnel dossier, engineering philosophy, and active technical directions | Viewport scroll (`↑`/`↓`, `k`/`j`, mouse wheel) |
| **02 PROJECTS** | Technical records ledger (`ID | PROJECT | STATUS | TECHNOLOGIES`) + DAG architecture dossier | `Tab` / `Shift+Tab` to switch projects |
| **03 STACK** | Structured engineering taxonomy (Programming, Backend, Frontend, Data, Concepts) | No arbitrary percentages |
| **04 JOURNEY** | Chronological milestone ledger with audit log styling | Verified timeline |
| **05 INTERESTS** | "Beyond the Terminal": Poetry, Superbikes, Mechanical Engineering, Curiosity | Editorial workstation panels |
| **06 CONTACT** | Communications console table with integrated OSC 52 terminal clipboard copy | `↑`/`↓` select channel, `ENTER` to copy |
| **Command Mode** | Persistent command bar (`COMMAND > _`) with fuzzy typo suggestion | Press `:` anywhere to enter command mode |
| **Help Screen** | Institutional keyboard cheat sheet and command reference | Toggle with `F1` or `?` |
| **Easter Eggs** | `:neofetch`, `:sudo hire jawad`, `:whoami`, `:date`, `:secret` | Workstation modal dialogs |

---

## 4. Featured Project: Workflow Automation Platform

- **Status:** In Development
- **Stack:** Go, React, TypeScript, PostgreSQL, Redis, Docker
- **Pipeline:** `[GitHub Issue Trigger] ──▶ [Workflow DAG Engine] ──▶ [AI Summary] ──▶ [Slack Webhook]`
- **Key Architectural Decisions:**
  - **Go Execution Engine:** Low memory footprint, fast cold starts, and native goroutine concurrency.
  - **DAG Representation:** Directed acyclic graph representation for workflow node execution and dependency resolution.
  - **Redis-Backed Queues:** Decoupled trigger ingestion from worker processing for reliable task delivery.
- **Engineering Challenges:** Idempotent task execution across network failures, dynamic schema parsing for action payloads.
- **Key Learnings:** Distributed queue mechanics, exponential backoff retry strategies, and strict state machine boundaries.

---

## 5. Local Development Guide

### Prerequisites
- Go 1.24+ (`go version`)
- Git

### Commands

```bash
# Clone the repository
git clone https://github.com/lexisync/jawad-os.git
cd jawad-os

# Run locally in terminal
make run

# Run directly to dashboard (skip boot sequence)
make run-skip-boot

# Run test suite & race detector
make test
go test -race ./...

# Run linter
make lint

# Launch local Wish SSH server
make ssh-serve

# Connect to local SSH server
make ssh-connect
```

---

## 6. Docker & Production Deployment

### Docker Compose

```bash
docker compose up -d
```

Connect remotely:
```bash
ssh localhost -p 2222
```

### Native Systemd Linux Service

1. Build production binary:
   ```bash
   CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o jawad cmd/jawad/main.go
   ```
2. Copy binary and config to `/opt/jawad-os/`:
   ```bash
   sudo mkdir -p /opt/jawad-os/configs /opt/jawad-os/.ssh
   sudo cp jawad /opt/jawad-os/
   sudo cp configs/portfolio.yaml /opt/jawad-os/configs/
   ```
3. Enable systemd unit:
   ```bash
   sudo cp deploy/jawad-os.service /etc/systemd/system/
   sudo systemctl daemon-reload
   sudo systemctl enable --now jawad-os
   ```

---

## 7. Keybindings & Command Mode Reference

### Global Navigation
- `↑` / `k`: Scroll up / navigate up
- `↓` / `j`: Scroll down / navigate down
- `PgUp` / `PgDn`: Fast scroll viewport
- `Mouse Wheel`: Scroll active module
- `ENTER`: Select item / copy contact info via OSC 52
- `1` - `6`: Jump directly to module 01–06 from Dashboard
- `Tab` / `Shift+Tab`: Switch project records in Project Explorer
- `ESC`: Return to previous screen / exit command mode / close modal
- `F1` / `?`: Toggle Help reference
- `q` / `Ctrl+C`: Quit or disconnect

### Command Prompt (`:`)
Press `:` anywhere to open command mode:

| Command | Action |
| :--- | :--- |
| `:profile` / `:about` | Navigate to Profile dossier (Screen 01) |
| `:projects` | Navigate to Projects Explorer (Screen 02) |
| `:projects workflow` | Jump to Workflow Platform architecture dossier |
| `:stack` / `:skills` | Open categorized Technology Stack matrix (Screen 03) |
| `:journey` / `:experience`| Open Engineering Journey ledger (Screen 04) |
| `:interests` | Open "Beyond the Terminal" (Screen 05) |
| `:contact` | Open Communications console (Screen 06) |
| `:home` | Return to main Workstation Dashboard |
| `:whoami` | Print active session identity |
| `:date` | Print local and UTC system time |
| `:neofetch` | Display `JAWAD.OS SYSTEM TELEMETRY` modal |
| `:sudo hire jawad` | Run interactive recruiter hiring protocol |
| `:secret` | Read personal developer note from Jawad |
| `:clear` | Clear status message |
| `:exit` / `:quit` | Exit application |

---

## 8. Security Considerations

- **Complete Session Isolation:** Each SSH connection instantiates an independent Bubble Tea model with isolated memory.
- **Unprivileged Execution:** Production Docker container runs as user `jawad` (UID `10001`). Systemd service enables strict Linux sandboxing.
- **No Command Shell Execution:** SSH sessions run within the internal Bubble Tea event loop; no Linux shell (`/bin/sh`) is ever exposed.

---

<div align="center">

Crafted with dedication by **Jawad Ahmad** • Built with the [Charm](https://charm.sh) ecosystem.

</div>

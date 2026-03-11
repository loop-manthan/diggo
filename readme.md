# diggo

A command-line tool written in Go that displays directory sizes in a tree format.

## Features

- Recursively scans directories and calculates cumulative sizes
- Displays results as an indented tree with `├──` / `└──` connectors
- Human-readable size output (B, KB, MB, GB, …)
- Configurable depth limit

## Usage

```
diggo [flags] [path]
```

| Flag | Default | Description |
|------|---------|-------------|
| `-depth` | `-1` (unlimited) | Max depth to display (0 = root only, 1 = root + children, …) |

If no path is given, the current directory (`.`) is used.

### Examples

```bash
# Full tree of the current directory
diggo

# Only show two levels deep
diggo -depth 2

# Scan a specific path
diggo /home/user/projects
```

## Project Structure

```
diggo/
├── main.go           # Entry point — parses flags, runs scan, prints tree
├── model/
│   └── dir.go        # DirInfo struct (Path + Size)
├── scan/
│   └── scan.go       # Recursively walks the filesystem and aggregates sizes
└── format/
    └── format.go     # Tree and list formatters with human-readable sizes
```

## Building

```bash
go build -o diggo .
```

## Requirements

- Go 1.21+

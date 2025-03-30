# Clean Architecture AFWTls

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Architecture](#architecture)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Flags](#flags)
- [Dependencies](#dependencies)
- [Contributors](#contributors)

## Overview

**AFWTls** is a cross-platform CLI application written in Go that replicates the behavior of the Linux `ls` command. It lists directory contents with enhanced capabilities and terminal visuals. This project started as a refactor of the `EDLS` project from EDteam's [Go desde cero 2023](https://ed.team/cursos/go) course and was restructured with Clean Architecture to promote scalability and maintainability.

## Features

- Cross-platform compatibility: Linux, Windows, and macOS.
- Hidden file support.
- Pattern matching (case-insensitive).
- Custom record limits.
- Sort by file size or modification time.
- Reverse sorting order.
- Color-coded and icon-enhanced file display.

## Architecture

AFWTls follows the **Clean Architecture** principles, ensuring clear separation of concerns:

- **Adapters Layer:** Interfaces with the CLI (flag parsing) and UI (printing logic).
- **Domain Layer:** Contains core business entities and constants.
- **Use Case Layer:** Implements business rules such as reading directory contents, filtering, sorting, and file classification.

### Design Patterns Applied

- **Strategy Pattern (manual):** Used in the sorting logic with pluggable strategies for sorting by name, size, or time.
- **Builder-like Composition:** Building file entities with contextual metadata (e.g. user, group, file type).

## Project Structure

```
├── cmd
│   └── afwtls
│       └── cli
│           └── main.go         # CLI entrypoint
├── internal
│   ├── adapters
│   │   ├── flag                # Flag parsing logic
│   │   └── ui                  # Colorful output rendering
│   ├── domain
│   │   ├── constants           # File types and extensions
│   │   └── entities            # File entity definition
│   └── usecase
│       ├── files               # Core business logic (list, sort, classify files)
│       └── flag                # Command execution orchestration
├── go.mod                      # Dependencies
├── go.sum                      # Checksum file
└── README.md                   # This file
```

## Installation

```bash
git clone https://github.com/AndresFWilT/afwtls.git
cd afwtls
go build -o afwtls ./cmd/afwtls/cli
```

## Usage

Run the command inside any directory:

```bash
./afwtls
```

You can also pass a specific directory:

```bash
./afwtls /path/to/dir
```

## Flags

| Flag      | Description                                 |
|-----------|---------------------------------------------|
| `-a`      | Show all files, including hidden ones       |
| `-n int`  | Limit the number of files shown             |
| `-p str`  | Filter files by a pattern (case-insensitive)|
| `-s`      | Sort by file size (ascending)               |
| `-t`      | Sort by last modification time (ascending)  |
| `-r`      | Reverse the sort order                      |

## Dependencies

- [`fatih/color`](https://github.com/fatih/color): Colored terminal output
- [`x/sys`](https://pkg.go.dev/golang.org/x/sys): OS-specific utilities
- [`x/exp`](https://pkg.go.dev/golang.org/x/exp): Experimental Go features

## Contributors

- **Andrés Felipe Wilches Torres**  
  [Contact me!](mailto:andresfwilchestdev@gmail.com)

---

> This project is part of my personal journey in mastering Go and applying Clean Architecture. Pull requests and feedback are welcome!

# MiniDB

A SQLite-inspired embedded relational database written in Go.

MiniDB is an educational database engine built from scratch to explore database internals, storage engines, SQL parsing, indexing, and transaction management.

## Project Status

Currently under development.

## Goals

- [ ] SQL lexer and parser
- [ ] Abstract Syntax Tree (AST)
- [ ] In-memory query execution
- [ ] Persistent database storage
- [ ] Page-based storage engine
- [ ] B+ tree indexing
- [ ] Database catalog
- [ ] CRUD operations
- [ ] Write-ahead logging
- [ ] Transaction recovery
- [ ] Command-line interface

## Requirements

- Go (version specified in `go.mod`)
- Git

## Getting Started

Clone the repository:

```bash
git clone https://github.com/YOUR_USERNAME/minidb.git
cd minidb
```

Run tests:

```bash
make test
```

Build:

```bash
make build
```

Run:

```bash
./bin/minidb --version
```

## Architecture

The database will eventually consist of several major components:

- SQL frontend: lexical analysis and parsing.
- Query engine: planning and execution.
- Catalog: table and schema metadata.
- Storage engine: pages, records, and B+ trees.
- Transaction manager: commit, rollback, and recovery.

Detailed architecture documentation will be added as these components are implemented.

## License

MIT

# MiniDB Architecture

## Overview

MiniDB is an educational, SQLite-inspired embedded relational database written in Go.

The project is designed around a modular architecture consisting of a SQL frontend, query engine, catalog, and persistent storage engine.

The primary goal is to maintain clear boundaries between query processing and storage management.

## 1. System Architecture

The database consists of the following components:

1. Public API
2. SQL frontend
3. Query engine
4. Catalog
5. Storage engine
6. Transaction manager (planned)

### SQL Frontend

Responsible for converting SQL statements into structured representations.

Components:

- Lexer: Converts SQL characters into tokens.
- Parser: Converts tokens into an Abstract Syntax Tree.
- AST: Represents the structure of SQL statements.

The SQL frontend does not access database files or retrieve records.

### Query Engine

Responsible for validating and executing SQL statements.

Components:

- Planner: Determines how queries should be executed.
- Executor: Performs planned database operations.
- Expression evaluator: Evaluates comparisons and logical expressions.

### Catalog

Responsible for managing database metadata.

Stores:

- Table definitions.
- Column definitions.
- Data types.
- Table storage locations.
- Schema information.

Catalog metadata will eventually be persisted in the database file.

### Storage Engine

Responsible for persistent record management.

Components:

- B+ tree: Organizes records and supports key-based searching.
- Page: Defines the binary layout of database pages.
- Record: Handles record serialization and deserialization.
- Pager: Reads, writes, allocates, and caches database pages.

### Transaction Manager

Planned component responsible for:

- Transaction boundaries.
- Commit and rollback.
- Write-ahead logging.
- Crash recovery.
- Database durability.

## 2. Query Execution Pipeline

A SQL statement follows this general execution path:

1. Public API receives SQL.
2. Lexer generates tokens.
3. Parser constructs an AST.
4. Catalog resolves tables and columns.
5. Planner generates an execution plan.
6. Executor performs the plan.
7. Storage engine retrieves or modifies records.
8. Results are returned through the public API.

## 3. Package Organization

```text
cmd/minidb/          CLI application

internal/sql/        SQL lexer, parser, and AST

internal/engine/     Query planning and execution

internal/catalog/    Table and schema metadata

internal/storage/
    pager/            Database file and page management
    page/             Page layouts and encoding
    record/           Record serialization
    btree/            B+ tree implementation

db.go                Public database API
result.go            Public query result types
```

## 4. Architectural Principles

### Separation of Concerns

Each package should have one clearly defined responsibility.

The SQL frontend must not depend on storage implementation details.

The pager must not understand SQL statements.

### Dependency Direction

Higher-level components should depend on small abstractions rather than concrete storage implementations wherever this establishes a useful boundary.

### Minimal Interfaces

Interfaces should be introduced at meaningful component boundaries rather than indiscriminately.

### Incremental Development

Each major component should be independently testable before integration.

### Explicit Binary Storage

Database records and pages will use an explicitly defined binary format rather than generic object serialization.

## 5. Initial Design Decisions

| Decision | Choice |
|---|---|
| Language | Go |
| Database model | Embedded relational |
| Storage | Single binary database file |
| Page size | 4096 bytes |
| Tree structure | B+ tree-inspired |
| SQL parser | Hand-written recursive descent |
| Initial execution model | Single-process |
| Initial concurrency model | Single writer, synchronized access |
| Dependencies | Go standard library |

## 6. Deliberate Simplifications

MiniDB is not intended to reproduce the complete SQLite implementation.

The initial implementation will use:

- A limited SQL grammar.
- A simple query planner.
- A single primary storage tree initially.
- Basic transaction semantics.
- A simplified write-ahead logging implementation.

Additional functionality will be introduced incrementally.

## 7. Testing Strategy

Each component will have independent unit tests.

Integration tests will verify communication between components.

Additional testing will include:

- Database persistence.
- B+ tree invariants.
- Binary page encoding and decoding.
- SQL parsing.
- Crash recovery.
- Concurrent access.
- Fuzz testing.
- Benchmarks.

## 8. Project Objective

The finished project should demonstrate an understanding of database architecture, compiler fundamentals, data structures, binary storage, query execution, transaction management, and Go systems programming.

---
search:
  exclude: true
---

# TIG Engineering & Architectural Standards (Target Conventions)

This document defines the **target engineering conventions, clean architecture standards, and Go best practices** for refactoring and evolving the **TIG** codebase.

---

## 1. Architectural Model & Layer Rules

The project follows **Clean Architecture / Hexagonal Architecture (Ports and Adapters)** combined with **Domain-Driven Design (DDD)**.

```
+-------------------------------------------------------------+
|                     Representation Layer                    |
|      (CLI Commands, Terminal Tables, Flags, Presenters)     |
+-------------------------------------------------------------+
                              |
                              v (depends on)
+-------------------------------------------------------------+
|                      Application Layer                      |
|       (Use Cases, Input/Output DTOs, Repository/Service Ports) |
+-------------------------------------------------------------+
               /                              \
 (depends on) /                                \ (implements)
             v                                  v
+-------------------------+         +-------------------------+
|      Domain Layer       |         |  Infrastructure Layer   |
|   (Entities, Value      |         |  (Git, OpenSSH, JSON,   |
|    Objects, Invariants) |         |   Filesystem Adapters)  |
+-------------------------+         +-------------------------+
```

### Dependency Inversion & Isolation Rules
1. **Inward Dependencies Only**:
   - Outer layers depend on inner layers.
   - `Representation` → `Application` → `Domain`.
   - `Infrastructure` implements interfaces (ports) defined in `Application`.
2. **Domain Purity**:
   - The `domain` package has **zero external dependencies** outside the Go standard library.
   - No CLI references, no JSON tags, no filesystem imports.
3. **Application Decoupling**:
   - Use cases must depend only on abstractions (ports: `repositories`, `services`).
   - Use cases must never import `infrastructure` or `representation` packages.
4. **Single Source of Truth**:
   - All active code lives under `app/`. Duplicate root directories (`/domain`, `/application`, `/infrastructure`, `/representation`) must be eliminated.

---

## 2. Go Idiomatic Naming Conventions

### 2.1. Packages (Standard Library Style)
* **Rule**: Package names must be **all lowercase, single-word, concise, and without punctuation or camelCase**.
* Avoid redundant suffixes like `Entity`, `Service`, `Helper`.
* In Go, the package name qualifies the type: `user.User` is idiomatic; `UserEntity.User` is an anti-pattern.

| Current / Non-idiomatic | Target Standard (Go Best Practice) |
| :--- | :--- |
| `package UserEntity` | `package user` |
| `package RepositoryEntity` | `package repository` |
| `package OrganizationEntity` | `package organization` |
| `package usecases` | `package usecase` |
| `package cmds` | `package cmd` (or `package cli`) |
| `package infrastructureServices` | `package service` |

### 2.2. File Naming (Idiomatic Go)
* **Standard**: File names should be lowercase using `snake_case`.
* Multi-dot suffixes (`.use-case.go`, `.cmd.go`) are from Node.js/NestJS. Go groups responsibility by directory and package, not dotted file extensions.

| Current Pattern | Target Standard |
| :--- | :--- |
| `add-new-user.use-case.go` | `add_new_user.go` (inside `app/application/usecase/`) |
| `checkout-user.cmd.go` | `checkout_user.go` (inside `app/representation/cmd/`) |
| `organization-repository.go` | `organization.go` (inside `app/application/repository/`) |
| `git.service.go` | `git.go` (inside `app/infrastructure/service/`) |
| `system.const.go` | `constants.go` (inside `app/infrastructure/`) |

### 2.3. Structs, Interfaces & Constructors
* **Constructors**: Always use `New<Type>()`:
  - `func NewAddNewUser(repo repository.Organization, git service.Git) *AddNewUser`
  - Domain entities: `user.New(name, email, host string)` or `user.NewFromDraft(draft Draft)` instead of generic `Constructor(...)`.
* **Interface Naming**:
  - One-method interfaces end with `-er` (e.g., `Reader`, `Writer`, `Cloner`).
  - Complex interfaces name their role clearly: `OrganizationRepository` or `GitService`.
* **Accessors / Getters**:
  - Never use `Get` prefixes for getters:
    - ✅ `user.Name()`, `user.Email()`
    - ❌ `user.GetName()`
* **Setters / Mutators**:
  - Use `Set...` or intention-revealing verbs: `user.SetName(...)` or `user.AssignRepository(...)`.

---

## 3. Error Handling Best Practices

### 3.1. Avoid Boolean Status Returns

Collapsing errors into `bool` loses failure context (why did Git fail? permissions, network, or invalid arguments?).

* ❌ **Anti-pattern**:
```go
func (gs *GitService) CreateSSHKey(alias, email string) bool
func (cu *CheckoutUserUseCase) Execute(alias string) (bool, error)
```

* ✅ **Go Best Practice**:
```go
// Infrastructure services return an explicit error
func (s *GitService) CreateSSHKey(alias, email string) error

// Use cases return an error (and data if applicable), not a boolean
func (uc *CheckoutUser) Execute(alias string) error
```

### 3.2. Sentinel & Wrapped Errors

* Define sentinel domain errors using `errors.New(...)`:
```go
var (
    ErrUserNotFound       = errors.New("user not found")
    ErrUserAlreadyExists  = errors.New("user already exists")
    ErrEmptySSHKey        = errors.New("ssh key cannot be empty")
)
```

* Wrap low-level errors to preserve the error stack using `fmt.Errorf`:
```go
if err := s.fs.CreateDirectory(path); err != nil {
    return fmt.Errorf("failed to create directory %q: %w", path, err)
}
```

---

## 4. Layer Implementations & Design Patterns

### 4.1. Domain Layer (`app/domain/`)
* **Entities**: Hold private fields to guard invariants.
* **Validation**: Validation must happen during creation or mutation inside domain constructors/methods, returning domain sentinel errors immediately if data is invalid.
* **Rich Domain Models**: Business behavior belongs on entities (e.g., `org.AddUser(u)` checks for duplicate alias or existing global user rather than doing it in the use case).

### 4.2. Application Layer (`app/application/`)
* **Use Cases (Interactors)**:
  - Each use case represents a single user intent.
  - Receives an `Input` DTO, validates presentation boundaries, invokes Domain entities and Repository ports, and returns either `Output` DTO or `error`.
  - Accept `context.Context` as the first argument in `Execute(ctx context.Context, input Input) error` to support timeouts and cancellations.

### 4.3. Infrastructure Layer (`app/infrastructure/`)
* **OS / Process Execution**:
  - Capture `stderr` when running `git` or `ssh-keygen` so real system error messages are returned in errors.
* **Storage / Persistence**:
  - **Atomic TOML Persistence**: Stores configurations safely in `~/.tig/organization.toml` using atomic file writes (write to temporary `.tmp` file, sync, then rename) to prevent state corruption on crashes.

### 4.4. Representation Layer (`app/representation/`)
* **CLI Command Handlers**:
  - Keep commands thin: parse flags → map to UseCase Input DTO → invoke UseCase → format terminal output.
  - Zero business logic and zero direct disk/git manipulation in CLI command files.

---

## 5. Automated Testing Standards

* **Zero-Mock, AAA Testing Pattern**:
  - All tests follow the Arrange, Act, Assert pattern clearly demarcated with comments: `// Arrange`, `// Act`, `// Assert`.
  - Avoid mock-heavy unit tests that test implementation details instead of behavior.
* **Test Isolation**:
  - Production code inside `app/` contains zero test files.
  - Unit tests live under `test/unit/` testing pure domain logic without side-effects.
  - Integration tests live under `test/integration/` exercising real CLI commands and file operations in isolated temporary directories (`t.TempDir()`).
* **Containerized Execution**:
  - Automated tests execute reproducibly in Docker using `docker compose run --build --rm app-test`.

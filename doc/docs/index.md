# TIG (The Invisible Git)

**TIG** is a command-line tool written in Go that simplifies managing multiple Git identities, OpenSSH keys, and repository configurations on a single workstation without manual SSH configuration editing.

---

## Key Features

- **Multi-Identity Git Management**: Switch seamlessly between personal, work, and client identities.
- **Automated OpenSSH Key Generation & Agent Integration**: Generates Ed25519 / RSA keys and maintains your `~/.ssh/config` file automatically.
- **Smart Origin Transformation**: Automatically transforms repository clone origins (e.g., `git@ssh.bitbucket.org:...` or `git@github.com:...`) to route through user-specific SSH host aliases.
- **Atomic TOML Persistence**: Stores configurations safely in `~/.tig/db.toml` using atomic file writes to prevent corruption.
- **Zero-Mock, AAA Integration & Unit Testing**: Full test coverage isolated in `test/unit` and `test/integration` executed through Docker containers.

---

## Installation

### Prerequisites
Make sure **git** and **openssh-client** are installed:
```bash
# Debian / Ubuntu
sudo apt install -y openssh-client git

# macOS
brew install git openssh
```

### Install with Official Script
```bash
wget https://github.com/sneuder/tig/releases/download/v2.0.0/install.sh
chmod +x install.sh
sudo ./install.sh
```

### Or Install via Go
```bash
go install github.com/sneuder/tig@latest
```

*For complete installation instructions, see the [Installation Guide](installation.md).*

---

## Basic Workflow

1. **Add a user identity**:
   ```bash
   tig user add -a work -n "Jane Doe" -e "jane@company.com" -H b
   ```

2. **Generate SSH keys**:
   ```bash
   tig user checkout -a work
   tig user generate
   ```

3. **Inspect your configuration**:
   ```bash
   tig user config
   ```

4. **Add and clone repositories**:
   ```bash
   tig repository add -a backend -o git@ssh.bitbucket.org:org/backend.git
   tig repository clone
   ```

---

## Documentation Sections

- [Installation Guide](installation.md): All installation methods, requirements, and troubleshooting.
- [Architecture & Conventions](conventions.md): Clean Architecture principles, package layout, DDD patterns, and AAA testing standards.
- [User Commands](commands/user.md): Detailed reference for managing user identities, switching accounts, and generating SSH keys.
- [Repository Commands](commands/repository.md): Detailed reference for registering, listing, and cloning repositories with host alias resolution.

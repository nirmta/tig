# TIG (The Invisible Git)

[![Go Report Card](https://goreportcard.com/badge/github.com/sneuder/tig)](https://goreportcard.com/report/github.com/sneuder/tig)

**TIG** simplifies managing multiple Git identities, OpenSSH keys, and repository configurations on a single workstation without manual SSH configuration editing.

---

## Features

- **Multi-Identity Git Management**: Switch seamlessly between personal, work, and client identities.
- **Automated OpenSSH Key Management**: Automatically generates key pairs and maintains `~/.ssh/config`.
- **Smart Origin Transformation**: Automatically rewrites repository clone origins (e.g. `git@ssh.bitbucket.org:...` or `git@github.com:...`) to match the user's SSH Host alias (`<user-alias>:<path>`).
- **Atomic TOML Persistence**: Configuration stored safely in `~/.tig/db.toml` with zero risk of partial file writes.
- **Cross-Platform Release via GoReleaser**: Pre-built binaries and packages for Linux (`.deb`, `tar.gz`) and macOS (Homebrew tap, `tar.gz`).
- **Clean Architecture & AAA Tests**: Production code separated from zero-mock unit and integration tests.

---

## Installation & Binary Paths

### Prerequisites
Ensure `git` and `openssh-client` are installed:
```bash
# Debian / Ubuntu
sudo apt update && sudo apt install -y openssh-client git

# macOS (Homebrew)
brew install git openssh
```

### macOS Installation (Homebrew)
Install via the official TIG Homebrew tap:
```bash
brew tap nirmta/tig
brew install tig
```
*Installed path:*
- Apple Silicon (M1/M2/M3/M4): `/opt/homebrew/bin/tig`
- Intel Macs: `/usr/local/bin/tig`

### Linux Installation (Debian / Ubuntu `.deb`)
```bash
wget https://github.com/nirmta/tig/releases/latest/download/tig_linux_amd64.deb
sudo dpkg -i tig_linux_amd64.deb
```
*Installed path:* `/usr/bin/tig`

### Manual Download (GoReleaser Archives)
```bash
# Linux
curl -LO "https://github.com/nirmta/tig/releases/latest/download/tig_Linux_$(uname -m).tar.gz"
tar -xzf tig_Linux_*.tar.gz
sudo mv tig /usr/local/bin/

# macOS
curl -LO "https://github.com/nirmta/tig/releases/latest/download/tig_Darwin_$(uname -m).tar.gz"
tar -xzf tig_Darwin_*.tar.gz
sudo mv tig /usr/local/bin/
```
*Target system path:* `/usr/local/bin/tig`

### Shell PATH Setup
If you move the binary to `~/.local/bin` or use `go install`, ensure it is in your `PATH`:
```bash
# For zsh (~/.zshrc)
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.zshrc

# For bash (~/.bashrc)
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
```

---

## CLI Commands

### User Management (`tig user`)

- **Add User Identity**:
  ```bash
  tig user add -a work -n "Jane Doe" -e "jane@company.com" -H b
  # -H flag options: g (GitHub), b (Bitbucket), l (GitLab)
  ```

- **List Users**:
  ```bash
  tig user list
  ```

- **Switch Active User**:
  ```bash
  tig user checkout -a work
  ```

- **Generate SSH Key**:
  ```bash
  tig user generate
  ```

- **View Active Config & Public Key**:
  ```bash
  tig user config
  ```

- **Remove User**:
  ```bash
  tig user remove -a work
  ```

### Repository Management (`tig repository`)

- **Add Repository** (Origin automatically transforms to alias routing):
  ```bash
  tig repository add -a backend -o git@ssh.bitbucket.org:mycompany/backend-api.git
  # Saved in ~/.tig/organization.toml as: work:mycompany/backend-api.git
  ```

- **List Repositories**:
  ```bash
  tig repository list
  ```

- **Batch Clone Repositories**:
  ```bash
  tig repository clone
  # Clones all active user repositories into ./work/<repo-alias>/
  ```

- **Remove Repository**:
  ```bash
  tig repository remove -a backend
  ```

### Backup & Disaster Recovery (`tig backup`, `tig restore`)

- **Backup Configuration**:
  ```bash
  tig backup -o ~/tig-backup.toml
  ```

- **Restore Configuration & Regenerate SSH Setup**:
  ```bash
  tig restore -i ~/tig-backup.toml
  ```

---

## Documentation

Full documentation is available via MkDocs in [`doc/`](doc/):

```bash
docker compose up doc-live
```
Open **http://localhost:8081** to view the live documentation site:
- [Installation Guide](doc/docs/installation.md)
- [Architecture & Conventions Guide](doc/docs/conventions.md)
- [User Commands Reference](doc/docs/commands/user.md)
- [Repository Commands Reference](doc/docs/commands/repository.md)
- [Backup & Restore Reference](doc/docs/commands/backup_restore.md)

---

## Testing

Run unit and integration tests inside Docker:

```bash
docker compose run --build --rm app-test
```

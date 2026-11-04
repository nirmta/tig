# TIG (The Invisible Git)

**TIG** is a command-line tool written in Go that simplifies managing multiple Git identities, OpenSSH keys, and repository configurations on a single workstation without manual SSH configuration editing.

---

## The Problem

Developers frequently work with multiple Git accounts (e.g., personal GitHub, company Bitbucket, client GitLab). Managing this manually is frustrating and error-prone:

- **SSH Key Collisions**: Standard SSH cannot easily distinguish which private key to send when multiple accounts share the same host (`github.com` or `bitbucket.org`).
- **Messy `~/.ssh/config`**: Developers are forced to hand-edit SSH configuration files with complex `Host` aliases and remember custom hostnames for every single clone.
- **Wrong Author Commits**: Accidentally committing proprietary work code with your personal email, or personal side-projects with your company email.
- **Machine Migration Friction**: Moving to a new laptop requires manually tracking down which keys belonged to which account and reconfiguring everything from scratch.

---

## How TIG Solves It

TIG makes multi-account Git management invisible and frictionless:

1. **Automated SSH Configuration & Key Management**:
   Create a profile with `tig user add` and `tig user generate`. TIG automatically generates the SSH keypair, registers it with `ssh-agent`, and maintains your `~/.ssh/config` file without any manual editing.

2. **Smart Origin Transformation**:
   When you register a repository using standard Git URLs (e.g., `git@ssh.bitbucket.org:org/repo.git`), TIG automatically transforms the origin URL to route through the user's dedicated SSH Host alias (`user-alias:org/repo.git`). Git automatically uses the correct private key on clone and push.

3. **Instant Profile Switching**:
   Switching contexts is a single command (`tig user checkout <alias>`), which instantly updates your active global `git config` (`user.name` and `user.email`) to match.

4. **Atomic Storage & Safe Recovery**:
   All profiles and repository relationships are stored in `~/.tig/organization.toml` using atomic writes to prevent corruption. Back up your workstation with `tig backup` and restore everything on a new machine with `tig restore`.

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
wget https://github.com/nirmta/tig/releases/download/v4.0.0-rc.5/install.sh
chmod +x install.sh
sudo ./install.sh
```

### Or Install via Go
```bash
go install github.com/nirmta/tig@latest
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
   tig repository add -a backend -o git@ssh.bitbucket.org:mycompany/backend-api.git
   tig repository clone
   ```

5. **Backup and restore**:
   ```bash
   tig backup -o ~/my-backup.toml
   tig restore -i ~/my-backup.toml
   ```

---

## Documentation Sections

- [Installation Guide](installation.md): All installation methods, requirements, and troubleshooting.
- [User Commands](commands/user.md): Detailed reference for managing user identities, switching accounts, and generating SSH keys.
- [Repository Commands](commands/repository.md): Detailed reference for registering, listing, and cloning repositories with host alias resolution.
- [Backup & Restore](commands/backup_restore.md): Exporting and recovering workstation configurations safely.

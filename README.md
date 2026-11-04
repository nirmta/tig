# TIG (The Invisible Git)

[![Go Report Card](https://goreportcard.com/badge/github.com/sneuder/tig)](https://goreportcard.com/report/github.com/sneuder/tig)
[![Documentation](https://img.shields.io/badge/docs-online-blue)](https://tig-doc-nine.vercel.app/)

**TIG** is a developer CLI tool built to solve a universal pain point: **managing multiple Git identities, OpenSSH keys, and repository configurations on a single machine without constantly breaking your setup.**

📖 **Full Documentation & Command Reference**: [https://tig-doc-nine.vercel.app/](https://tig-doc-nine.vercel.app/)

---

## The Problem

Developers frequently work with multiple Git accounts (e.g., personal GitHub, company Bitbucket, client GitLab). Managing this manually is frustrating and error-prone:
- **SSH Key Collisions**: Standard SSH can't easily distinguish which private key to send when multiple accounts share the same host (`github.com` or `bitbucket.org`).
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

4. **Atomic TOML Storage & Safe Recovery**:
   All profiles and repository relationships are stored in `~/.tig/organization.toml` using atomic writes to prevent corruption. Back up your workstation with `tig backup` and restore everything on a new machine with `tig restore`.

---

## Quick Start

### Installation

```bash
# macOS (Homebrew Tap)
brew tap nirmta/tig
brew install tig

# Linux (Debian / Ubuntu)
wget https://github.com/nirmta/tig/releases/latest/download/tig_linux_amd64.deb
sudo dpkg -i tig_linux_amd64.deb
```

*(Pre-compiled binaries for all architectures and Go installation options are available in the [Installation Guide](https://tig-doc-nine.vercel.app/installation/)).*

### Basic Workflow

```bash
# 1. Create a work identity
tig user add -a work -n "Jane Doe" -e "jane@company.com" -H b

# 2. Switch to it and generate SSH credentials
tig user checkout -a work
tig user generate

# 3. Add a repository (origin automatically adapts to the SSH profile)
tig repository add -a api -o git@ssh.bitbucket.org:mycompany/backend-api.git

# 4. Clone your tracked repositories
tig repository clone
```

---

## Complete Documentation

Visit the official documentation site for complete command references, flags, architecture details, and guides:

👉 **[https://tig-doc-nine.vercel.app/](https://tig-doc-nine.vercel.app/)**

- **[Installation Guide](https://tig-doc-nine.vercel.app/installation/)**: Homebrew, `.deb` packages, manual binary downloads, and shell PATH setup.
- **[User Commands](https://tig-doc-nine.vercel.app/commands/user/)**: Identity registration, profile switching, SSH key generation, and config inspection.
- **[Repository Commands](https://tig-doc-nine.vercel.app/commands/repository/)**: Registering repositories, smart origin transformation, listing, and batch cloning.
- **[Backup & Disaster Recovery](https://tig-doc-nine.vercel.app/commands/backup_restore/)**: Machine migration, configuration export, and atomic restoration.

---

## Development & Testing

To run the full suite of unit and mock-free integration tests:

```bash
docker compose run --build --rm app-test
```

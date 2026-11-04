# Backup & Restore Commands

The `tig backup` and `tig restore` commands allow you to export and recover your entire TIG organization configuration (`~/.tig/organization.toml`) to and from a single backup file.

---

## 1. `tig backup`

Exports the current organization configuration to a TOML file.

### Aliases
`tig b`

### Flags

| Flag | Short | Type | Default | Description |
| :--- | :---: | :---: | :---: | :--- |
| `--output` | `-o` | String | `organization-backup.toml` | Destination backup file path. |

### Examples

```bash
# Default backup to ./organization-backup.toml
tig backup

# Custom destination path
tig backup -o ~/my-tig-backup.toml
```

---

## 2. `tig restore`

Overwrites `~/.tig/organization.toml` with the specified backup file and regenerates your workstation's `~/.ssh/config` for all restored users.

### Flags

| Flag | Short | Type | Required | Description |
| :--- | :---: | :---: | :---: | :--- |
| `--input` | `-i` | String | **Yes** | Path to the backup TOML file. |

### Examples

```bash
# Using flag
tig restore -i ~/my-tig-backup.toml

# Or passing as positional argument
tig restore ~/my-tig-backup.toml
```

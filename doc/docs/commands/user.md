# User Commands (`tig user`)

The `tig user` command group manages workstation identities, git configurations, and associated OpenSSH keys.

---

## 1. `tig user add`

Registers a new user identity in the local TIG database.

### Aliases
`tig user a`

### Flags

| Flag | Short | Type | Required | Description |
| :--- | :---: | :---: | :---: | :--- |
| `--alias` | `-a` | String | **Yes** | Unique alias identifier (e.g., `work`, `personal`). |
| `--name` | `-n` | String | **Yes** | Full user name for `git config user.name`. |
| `--email` | `-e` | String | **Yes** | Email address for `git config user.email`. |
| `--host` | `-H` | String | **Yes** | Repository host code: `g` (GitHub), `b` (Bitbucket), `l` (GitLab). |

### Example

```bash
tig user add -a work -n "Alice Smith" -e "alice@company.com" -H b
```

---

## 2. `tig user list`

Lists all registered user identities and marks the currently active identity.

### Aliases
`tig user l`

### Example

```bash
tig user list
```

#### Output
```text
┌───────┬─────────────┬───────────────────┬───────────────┐
│ ALIAS │  USER NAME  │       EMAIL       │  REPOSITORY   │
├───────┼─────────────┼───────────────────┼───────────────┤
│ work  │ Alice Smith │ alice@company.com │ bitbucket.org │
└───────┴─────────────┴───────────────────┴───────────────┘
```

---

## 3. `tig user checkout`

Switches the active user identity. Updates your global git configuration (`user.name`, `user.email`) to match the selected user.

### Aliases
`tig user c`

### Flags

| Flag | Short | Type | Required | Description |
| :--- | :---: | :---: | :---: | :--- |
| `--alias` | `-a` | String | **Yes** | Alias of the user identity to activate. |

### Example

```bash
tig user checkout -a work
```

---

## 4. `tig user generate`

Generates an OpenSSH key pair for the currently active user, registers the key with `ssh-agent`, and updates `~/.ssh/config` with host routing directives.

### Aliases
`tig user g`

### Example

```bash
tig user generate
```

Generated files:
- `~/.ssh/tig/<alias>/id_rsa`
- `~/.ssh/tig/<alias>/id_rsa.pub`
- Entry added to `~/.ssh/config`:
  ```ssh-config
  Host <alias>
    HostName <host>
    User git
    AddKeysToAgent yes
    IdentityFile ~/.ssh/tig/<alias>/id_rsa
  ```

---

## 5. `tig user config`

Displays the active user's current Git settings and the public SSH key for easy copying to GitHub, GitLab, or Bitbucket.

### Aliases
`tig user s`

### Example

```bash
tig user config
```

---

## 6. `tig user remove`

Deletes a user identity, removes their associated repositories, and cleans up their SSH keys from `~/.ssh/tig/`.

### Aliases
`tig user r`

### Flags

| Flag | Short | Type | Required | Description |
| :--- | :---: | :---: | :---: | :--- |
| `--alias` | `-a` | String | **Yes** | Alias of the user to delete. |

### Example

```bash
tig user remove -a work
```

# Repository Commands (`tig repository`)

The `tig repository` command group manages remote repository tracking and batch cloning attached to the active user identity.

---

## 1. `tig repository add`

Registers a remote repository under the currently active user. Automatically converts standard SCP and SSH URLs into user-alias routed URLs so Git routes connections using the user's specific SSH key.

### Aliases
`tig repository a`

### Flags

| Flag | Short | Type | Required | Description |
| :--- | :---: | :---: | :---: | :--- |
| `--alias` | `-a` | String | **Yes** | Repository alias / folder name identifier. |
| `--origin` | `-o` | String | **Yes** | Remote Git origin URL. |

### Smart Origin Transformation
When you provide a standard Git SSH URL, `tig` rewrites the origin to match the user's SSH Host alias:

| Input Origin | Active User Alias | Saved Transformed Origin |
| :--- | :---: | :--- |
| `git@ssh.bitbucket.org:mobohubbv2/backend.git` | `work` | `work:mobohubbv2/backend.git` |
| `git@github.com:nirmta/tig.git` | `work` | `work:nirmta/tig.git` |
| `git@gitlab.com:org/suborg/repo.git` | `work` | `work:org/suborg/repo.git` |
| `ssh://git@github.com/nirmta/repo.git` | `work` | `work:nirmta/repo.git` |

### Example

```bash
tig repository add -a backend -o git@ssh.bitbucket.org:mobohubbv2/mobohubb-backend-v2.git
```

---

## 2. `tig repository list`

Lists all repositories registered under the currently active user.

### Aliases
`tig repository l`

### Example

```bash
tig repository list
```

#### Output
```text
┌─────────┬───────────────────────────────────────────────┐
│  ALIAS  │                    ORIGIN                     │
├─────────┼───────────────────────────────────────────────┤
│ backend │ work:mobohubbv2/mobohubb-backend-v2.git       │
└─────────┴───────────────────────────────────────────────┘
```

---

## 3. `tig repository clone`

Batch-clones all registered repositories for the currently active user into `./<user-alias>/<repo-alias>` using the transformed alias routing.

### Aliases
`tig repository c`

### Example

```bash
tig repository clone
```

Clones each repository into:
```text
./work/backend/
```

---

## 4. `tig repository remove`

Removes a repository from the active user's tracked repository list.

### Aliases
`tig repository r`

### Flags

| Flag | Short | Type | Required | Description |
| :--- | :---: | :---: | :---: | :--- |
| `--alias` | `-a` | String | **Yes** | Alias of the repository to remove. |

### Example

```bash
tig repository remove -a backend
```

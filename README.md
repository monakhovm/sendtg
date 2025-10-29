# sendtg

`sendtg` is a simple CLI tool to send messages to Telegram via the Telegram Bot API.  
It supports `.env` files, environment variables, and inline token flags, and includes autocompletion scripts for bash, zsh, fish, and PowerShell.

---

## Features

- Send Telegram messages directly from the command line
- Token priority: CLI flag (`--token`/`-t`) → ENV → `.env` file
- Autocompletion support for bash, zsh, fish, and PowerShell
- View utility version via `version` subcommand
- Fully documented CLI in the style of `docker` or `kubectl`

---

## Installation

Clone the repository and build:

```bash
git clone https://github.com/yourusername/sendtg.git
cd sendtg
go build -o sendtg main.go
```

Or install via Go:

```bash
go install github.com/yourusername/sendtg@latest
```

---

## Usage

sendtg [chat_id] [message] [flags]
sendtg [command]


### Arguments

- `chat_id` — Telegram chat ID or user ID (e.g., `123456789`)
- `message` — The message text to send

### Flags

- `-t, --token string` — Telegram Bot Token (highest priority)
- `-h, --help` — Show help for sendtg

### Commands

- `completion` — Generate autocompletion scripts
- `version` — Show sendtg version
- `help` — Help about any command

---

## Token Priority

1. CLI flag `--token` or `-t` (highest priority)  
2. Environment variable `TELEGRAM_BOT_TOKEN`  
3. `.env` file in current directory (fallback)

---

## Examples

### Using `.env` file

```bash
echo 'TELEGRAM_BOT_TOKEN=123456:ABCDEF' > .env
sendtg 123456789 "Hello from .env!"
```

### Using environment variable

```bash
export TELEGRAM_BOT_TOKEN=123456:ABCDEF
sendtg 123456789 "Hello from ENV!"
```

### Using token inline (highest priority)

```bash
sendtg --token=123456:ABCDEF 123456789 "Hello from --token!"
sendtg -t 123456:ABCDEF 123456789 "Hello from -t!"
```

---

## Autocompletion

Generate autocompletion scripts for your shell:

#### Bash

```bash
sendtg completion bash > /etc/bash_completion.d/sendtg
source /etc/bash_completion.d/sendtg
```

#### Zsh

```bash
sendtg completion zsh > "${fpath[1]}/_sendtg"
autoload -U compinit && compinit
source ~/.zshrc
```

#### Fish

```bash
sendtg completion fish | source
```

#### PowerShell

```powershell
sendtg completion powershell | Out-String | Invoke-Expression
```

After setup, pressing `<TAB>` will autocomplete flags and commands:

```bash
sendtg 123456789 "Hi" -<TAB>
# Suggests: -t --token --help -h
```

---

## Version

```bash
sendtg version
# Output: sendtg v1.0.0
```

---

## .env Example

```env
# Telegram Bot Token
TELEGRAM_BOT_TOKEN=123456789:ABCDEF1234567890abcdefGHIJKLMN
```

---

## License

MIT License © 2025.

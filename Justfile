set dotenv-load

# aliases
alias fmt := fumpt
alias c := check

# List all available actions
default:
  just --list

# Prepare git hooks
git_init:
  chmod +x .githooks/prepare-commit-msg
  chmod +x .githooks/pre-commit
  git config core.hooksPath .githooks

# Run linters and formatters
check +FILES=".": (fumpt FILES) (lint FILES)

# Run gofumpt
fumpt +FILES=".":
  gofumpt -w -l -extra {{FILES}}

# Run golangci-lint
lint *FILES:
  golangci-lint run {{FILES}}

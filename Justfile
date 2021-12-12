set dotenv-load

# aliases
alias dev := develop

# List all available actions
default:
  just --list

# Prepare git hooks
git_init:
  chmod +x .githooks/prepare-commit-msg
  chmod +x .githooks/pre-commit
  git config core.hooksPath .githooks

# Prepare development environment
develop: git_init

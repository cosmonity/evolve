#!/usr/bin/env bash

set -e

REPO_ROOT="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/.." &>/dev/null && pwd)"
export REPO_ROOT

lint_module() {
  local root="$1"
  shift
  if [ -f $root ]; then
    cd "$(dirname "$root")"
  else
    cd "$REPO_ROOT/$root"
  fi
  echo "linting $(grep "^module" go.mod) [$(date -Iseconds -u)]"
  golangci-lint run ./... -c "${REPO_ROOT}/.golangci.yml" "$@"
}
export -f lint_module

# if LINT_DIFF env is set, only lint the files in the current commit otherwise lint all files
if [[ -z "${LINT_DIFF:-}" ]]; then
  find "${REPO_ROOT}" -type f -name go.mod -print0 | xargs -0 -I{} bash -c 'lint_module "$@"' _ {} "$@"
else
  if [[ -z $GIT_DIFF ]]; then
    GIT_DIFF=$(git diff --name-only) || true
  fi

  if [[ -z "$GIT_DIFF" ]]; then
    echo "no files to lint"
    exit 0
  fi

  GIT_DIFF=$(echo $GIT_DIFF | tr -d "'" | tr ' ' '\n' | grep '\.go$' | grep -v '\.pb\.go$' | grep -Eo '^[^/]+\/[^/]+' | uniq)

  lint_sdk=false
  for dir in ${GIT_DIFF[@]}; do
    if [[ ! -f "$REPO_ROOT/$dir/go.mod" ]]; then
      lint_sdk=true
    else
      lint_module $dir "$@"
    fi
  done

  if [[ $lint_sdk ]]; then
    cd "$REPO_ROOT"
    echo "linting repo [$(date -Iseconds -u)]"
    golangci-lint run ./... -c "${REPO_ROOT}/.golangci.yml" "$@"
  fi
fi

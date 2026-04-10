#!/usr/bin/env bash
#
# Compare two mcpchecker *eval* JSON files (mcpchecker check output) by running
# `mcpchecker summary ... --output json` on each and comparing token/task fields.
#
# Usage: compare-mcp-eval-baseline.sh <base_eval.json> <new_eval.json>
#
# Exits 0 when the comparison runs successfully. Writes to GITHUB_OUTPUT when set:
#   baseline-changed=true|false
#
# Prints human-readable messages to stderr; the decision is also in GITHUB_OUTPUT / stdout summary line.
#
set -euo pipefail

GOPATH="${GOPATH:-${HOME}/go}"
MCPCHECKER="${GOPATH}/bin/mcpchecker"

BASE_FILE="${1:?Usage: $0 <base_eval.json> <new_eval.json>}"
NEW_FILE="${2:?}"

if [[ ! -f "${MCPCHECKER}" ]]; then
  echo "Error: mcpchecker not found at ${MCPCHECKER}" >&2
  exit 1
fi

if [[ ! -f "${BASE_FILE}" ]]; then
  echo "Error: base file not found: ${BASE_FILE}" >&2
  exit 1
fi
if [[ ! -f "${NEW_FILE}" ]]; then
  echo "Error: new file not found: ${NEW_FILE}" >&2
  exit 1
fi

# Empty or non-array base means first baseline — always update
if ! jq -e 'type == "array" and length > 0' "${BASE_FILE}" >/dev/null 2>&1; then
  echo "No prior eval data in base file (empty or missing); baseline update required." >&2
  if [[ -n "${GITHUB_OUTPUT:-}" ]]; then
    echo "baseline-changed=true" >> "${GITHUB_OUTPUT}"
  fi
  echo "baseline-changed=true"
  exit 0
fi

BASE_SUMMARY="$(mktemp)"
NEW_SUMMARY="$(mktemp)"
trap 'rm -f "${BASE_SUMMARY}" "${NEW_SUMMARY}"' EXIT

if ! "${MCPCHECKER}" summary "${BASE_FILE}" --output json > "${BASE_SUMMARY}"; then
  echo "Error: mcpchecker summary failed for base file" >&2
  exit 1
fi
if ! "${MCPCHECKER}" summary "${NEW_FILE}" --output json > "${NEW_SUMMARY}"; then
  echo "Error: mcpchecker summary failed for new file" >&2
  exit 1
fi

# Compare aggregate fields and per-task tokensEstimated / mcpSchemaTokens (by task name).
CHANGED="$(jq -n \
  --slurpfile base "${BASE_SUMMARY}" \
  --slurpfile new "${NEW_SUMMARY}" \
  '
  ($base[0]) as $b
  | ($new[0]) as $n
  | ($b | {totalTokensEstimate, totalMcpSchemaTokens, tasksPassed, tasksTotal}) as $ba
  | ($n | {totalTokensEstimate, totalMcpSchemaTokens, tasksPassed, tasksTotal}) as $na
  | (($b.tasks // []) | map({name, tokensEstimated, mcpSchemaTokens}) | sort_by(.name)) as $bt
  | (($n.tasks // []) | map({name, tokensEstimated, mcpSchemaTokens}) | sort_by(.name)) as $nt
  | if $ba != $na then true
    elif $bt != $nt then true
    else false
    end
  ' | tr -d '\n')"

if [[ "${CHANGED}" == "true" ]]; then
  echo "MCP eval summary differs from master baseline (totals and/or per-task token estimates)." >&2
  if [[ -n "${GITHUB_OUTPUT:-}" ]]; then
    echo "baseline-changed=true" >> "${GITHUB_OUTPUT}"
  fi
  echo "baseline-changed=true"
else
  echo "MCP eval summary is identical to master baseline (totals and per-task token fields); skipping baseline PR." >&2
  echo "::notice::MCP eval baseline unchanged (mcpchecker summary match); skipping baseline update PR." >&2
  if [[ -n "${GITHUB_OUTPUT:-}" ]]; then
    echo "baseline-changed=false" >> "${GITHUB_OUTPUT}"
  fi
  echo "baseline-changed=false"
fi

exit 0

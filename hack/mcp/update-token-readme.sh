#!/usr/bin/env bash
#
# Updates the token consumption section in ai/mcp/README.md using
# mcpchecker summary --output json. Prefers a fresh run at mcpchecker-results/
# (make mcp-run-eval); otherwise uses the committed baseline at tests/evals/mcpchecker-results/.
#
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOTDIR="${SCRIPT_DIR}/../.."
GOPATH="${GOPATH:-${HOME}/go}"
MCPCHECKER="${GOPATH}/bin/mcpchecker"

README_FILE="${ROOTDIR}/ai/mcp/README.md"
RUN_OUTPUT="${ROOTDIR}/mcpchecker-results/mcpchecker-gemini-eval-out.json"
COMMITTED_BASELINE="${ROOTDIR}/tests/evals/mcpchecker-results/mcpchecker-gemini-eval-out.json"

if [[ -f "${RUN_OUTPUT}" ]]; then
  EVAL_FILE="${RUN_OUTPUT}"
elif [[ -f "${COMMITTED_BASELINE}" ]]; then
  EVAL_FILE="${COMMITTED_BASELINE}"
else
  echo "Error: no evaluation JSON found. Expected ${RUN_OUTPUT} (after make mcp-run-eval) or ${COMMITTED_BASELINE}." >&2
  exit 1
fi

if [[ ! -x "${MCPCHECKER}" ]]; then
  echo "Error: mcpchecker not found at ${MCPCHECKER}" >&2
  exit 1
fi

if [[ ! -f "${README_FILE}" ]]; then
  echo "Error: ${README_FILE} not found" >&2
  exit 1
fi

TMP_SUMMARY="$(mktemp)"
trap 'rm -f "${TMP_SUMMARY}"' EXIT
"${MCPCHECKER}" summary "${EVAL_FILE}" --output json > "${TMP_SUMMARY}"

TASKS_TOTAL=$(jq -r '.tasksTotal' "${TMP_SUMMARY}")
TASKS_PASSED=$(jq -r '.tasksPassed' "${TMP_SUMMARY}")
TOTAL_TOKENS=$(jq -r '.totalTokensEstimate' "${TMP_SUMMARY}")
MCP_SCHEMA_TOKENS=$(jq -r '.totalMcpSchemaTokens' "${TMP_SUMMARY}")
TASK_PASS_RATE=$(jq -r '.taskPassRate' "${TMP_SUMMARY}")
ASSERTION_PASS_RATE=$(jq -r '.assertionPassRate' "${TMP_SUMMARY}")

PASS_PCT=$(awk "BEGIN {printf \"%.0f\", ${TASK_PASS_RATE} * 100}")
ASSERT_PCT=$(awk "BEGIN {printf \"%.0f\", ${ASSERTION_PASS_RATE} * 100}")

MARKDOWN="### Evaluation Summary\n"
MARKDOWN+="\n"
MARKDOWN+="| Metric | Value |\n"
MARKDOWN+="|--------|-------|\n"
MARKDOWN+="| Tasks Passed | ${TASKS_PASSED}/${TASKS_TOTAL} (${PASS_PCT}%) |\n"
MARKDOWN+="| Assertions Pass Rate | ${ASSERT_PCT}% |\n"
MARKDOWN+="| Total Tokens Estimate | ${TOTAL_TOKENS} |\n"
MARKDOWN+="| MCP Schema Tokens | ${MCP_SCHEMA_TOKENS} |\n"
MARKDOWN+="\n"
MARKDOWN+="### Per-Task Breakdown\n"
MARKDOWN+="\n"
MARKDOWN+="| Task | Tokens Estimate | MCP Schema Tokens | Passed |\n"
MARKDOWN+="|------|----------------:|------------------:|--------|\n"

TASK_COUNT=$(jq '.tasks | length' "${TMP_SUMMARY}")
for i in $(seq 0 $((TASK_COUNT - 1))); do
  NAME=$(jq -r ".tasks[${i}].name" "${TMP_SUMMARY}")
  TOKENS=$(jq -r ".tasks[${i}].tokensEstimated" "${TMP_SUMMARY}")
  SCHEMA=$(jq -r ".tasks[${i}].mcpSchemaTokens" "${TMP_SUMMARY}")
  PASSED=$(jq -r ".tasks[${i}].taskPassed" "${TMP_SUMMARY}")
  if [[ "${PASSED}" == "true" ]]; then
    STATUS="✅"
  else
    STATUS="❌"
  fi
  MARKDOWN+="| ${NAME} | ${TOKENS} | ${SCHEMA} | ${STATUS} |\n"
done

START_MARKER="<!-- TOKENS-CONSUMPTION-START -->"
END_MARKER="<!-- TOKENS-CONSUMPTION-END -->"

REPLACEMENT="${START_MARKER}\n\n$(echo -e "${MARKDOWN}")\n${END_MARKER}"

awk -v start="${START_MARKER}" -v end="${END_MARKER}" -v replacement="${REPLACEMENT}" '
  $0 ~ start { print replacement; skip=1; next }
  $0 ~ end { skip=0; next }
  !skip { print }
' "${README_FILE}" > "${README_FILE}.tmp"

mv "${README_FILE}.tmp" "${README_FILE}"

echo "Updated token consumption section in ${README_FILE}"

#!/usr/bin/env sh

set -euo pipefail
set -x

# Note this script is only working with the local testnet in docker.

RPC_ADDR="tcp://localhost:26657"
CONTAINER="lorenzonode0"
BINARY="/lorenzod/lorenzod"
LORENZOD="docker exec -i $CONTAINER $BINARY"

PROPOSAL_FILE="proposal.json"
UPGRADE_PROPOSAL_FILE="upgrade_proposal.json"
TESTNESTS_DIR=$(git rev-parse --show-toplevel)/.testnets

CHAIN_ID=$(${LORENZOD} q block --node "$RPC_ADDR" | jq -r '.block.header.chain_id')

func_latest_block() {
  ${LORENZOD} q block --node "$RPC_ADDR" | jq -r '.block.header.height'
}

func_check_docker_running() {
  if ! docker inspect --format '{{.State.Running}}' lorenzonode0 2>/dev/null | grep -q "true"; then
    echo "Error: lorenzonode0 is not running."
    exit 1
  fi
}

func_upgrade_proposal() {
  echo "Preparing upgrade proposal..."

  if [ -z "${UPGRADE_NAME:-}" ]; then
    echo "Error: UPGRADE_NAME is not set."
    exit 1
  fi

  if [ ! -f "$UPGRADE_PROPOSAL_FILE" ]; then
    echo "Error: $UPGRADE_PROPOSAL_FILE not found in the current directory."
    exit 1
  fi

  cp "$UPGRADE_PROPOSAL_FILE" "$TESTNESTS_DIR/$PROPOSAL_FILE"

  current_height=$(func_latest_block)
  upgrade_height=$((current_height + 50))

  jq --arg name "$UPGRADE_NAME" \
     --arg height "$upgrade_height" \
     '.messages[0].plan.name = $name | .messages[0].plan.height = $height' \
     "$TESTNESTS_DIR/$PROPOSAL_FILE" > "$TESTNESTS_DIR/${PROPOSAL_FILE}.tmp"

  mv "$TESTNESTS_DIR/${PROPOSAL_FILE}.tmp" "$TESTNESTS_DIR/$PROPOSAL_FILE"

  echo "Proposal prepared with name: $UPGRADE_NAME and height: $upgrade_height"
}

func_prepare_proposal() {
  echo "Preparing a proposal..."

  local proposal=$1
  if [ -z "$proposal" ]; then
    echo "Error: proposal path is not set."
    exit 1
  fi

  if [ ! -f "$proposal" ]; then
    echo "Error: proposal file '$proposal' does not exist."
    exit 1
  fi

  cp "$proposal" "$TESTNESTS_DIR/$PROPOSAL_FILE"

  echo "Proposal prepared."
}

func_submit_proposal() {
  echo "Submitting proposal..."

  tx_hash=$(${LORENZOD} tx gov submit-proposal "/data/${PROPOSAL_FILE}" \
    --from node0 \
    --chain-id "${CHAIN_ID}" \
    --gas 400000 \
    --gas-prices 10alrz \
    --home "/data/node0/lorenzod" \
    --node "${RPC_ADDR}" \
    --keyring-backend test \
    --yes \
    --output json | jq -r '.txhash')
  sleep 5

  code=$(${LORENZOD} q tx "${tx_hash}" --node "${RPC_ADDR}" --output json | jq -r '.code')
  if [ "${code}" -ne 0 ]; then
    echo "Error: Failed to submit proposal."
    exit 1
  fi

  echo "Proposal submitted."
}

func_vote_proposal() {
  echo "Voting on proposal..."
  latest_proposal_id=$(${LORENZOD} q gov proposals --count-total --output json --node "${RPC_ADDR}" | jq -r '.pagination.total')

  for i in $(seq 0 3); do
    ${LORENZOD} tx gov vote "${latest_proposal_id}" yes \
      --from "node${i}" \
      --chain-id "${CHAIN_ID}" \
      --gas 400000 \
      --gas-prices 10alrz \
      --keyring-backend test \
      --home "/data/node${i}/lorenzod" \
      --node "${RPC_ADDR}" \
      --yes
    echo "Node${i} voted yes on proposal $latest_proposal_id."
  done

  sleep 5
  echo "Voting completed."
}

func_propose_upgrade() {
  func_upgrade_proposal

  func_submit_proposal

  func_vote_proposal
}

func_propose_general() {
  func_prepare_proposal "$1"

  func_submit_proposal

  func_vote_proposal
}

main() {
    func_check_docker_running

    case "${1:-}" in
        upgrade)
            func_propose_upgrade
            ;;
        general)
            shift
            func_propose_general "$@"
            ;;
        *)
            echo "Usage: $0 {upgrade|general [proposal_file]}"
            exit 1
            ;;
    esac
}

main "$@"
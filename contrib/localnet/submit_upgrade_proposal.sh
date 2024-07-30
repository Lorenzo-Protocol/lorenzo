#!/usr/bin/env sh

set -euo pipefail
set -x

RPC_ADDR="tcp://localhost:26657"
LORENZOD="docker exec -i lorenzonode0 /lorenzod/lorenzod"
PROPOSAL_FILE="upgrade_proposal.json"
TESTNESTS_DIR=$(git rev-parse --show-toplevel)/.testnets
CHAIN_ID=$(${LORENZOD} q block --node "$RPC_ADDR" | jq -r '.block.header.chain_id')

func_latest_block() {
  ${LORENZOD} q block --node "$RPC_ADDR" | jq -r '.block.header.height'
}

func_prepare_proposal() {
  echo "Preparing proposal..."

  if [ ! -f "$PROPOSAL_FILE" ]; then
    echo "Error: $PROPOSAL_FILE not found in the current directory."
    exit 1
  fi

  cp "$PROPOSAL_FILE" "${PROPOSAL_FILE}.tmp"
  mv "${PROPOSAL_FILE}.tmp" "$TESTNESTS_DIR/$PROPOSAL_FILE"

  if [ -z "${UPGRADE_NAME:-}" ]; then
    echo "Error: UPGRADE_NAME is not set."
    exit 1
  fi

  current_height=$(func_latest_block)
  upgrade_height=$((current_height + 50))

  jq --arg name "$UPGRADE_NAME" \
     --arg height "$upgrade_height" \
     '.messages[0].plan.name = $name | .messages[0].plan.height = $height' \
     "$TESTNESTS_DIR/$PROPOSAL_FILE" > "$TESTNESTS_DIR/${PROPOSAL_FILE}.tmp" \
      && mv "$TESTNESTS_DIR/${PROPOSAL_FILE}.tmp" "$TESTNESTS_DIR/$PROPOSAL_FILE"

  echo "Proposal prepared with name: $UPGRADE_NAME and height: $upgrade_height"
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

submit_and_vote() {
  func_prepare_proposal
  func_submit_proposal
  func_vote_proposal
}

submit_and_vote

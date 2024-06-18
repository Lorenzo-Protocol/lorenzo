#!/usr/bin/env sh
set -euo pipefail
set -x

BINARY=/lorenzod/${BINARY:-lorenzod}
ID=${ID:-0}
LOG=${LOG:-lorenzod.log}

if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder."
	exit 1
fi

export LORENZODHOME="/data/node${ID}/lorenzod"

if [ -d "$(dirname "${LORENZODHOME}"/"${LOG}")" ]; then
  "${BINARY}" --home "${LORENZODHOME}" "$@" | tee "${LORENZODHOME}/${LOG}"
else
  "${BINARY}" --home "${LORENZODHOME}" "$@"
fi

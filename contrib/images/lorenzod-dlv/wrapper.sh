#!/usr/bin/env sh
set -euo pipefail
set -x

DEBUG=${DEBUG:-0}
BINARY=/lorenzod/${BINARY:-lorenzod}
ID=${ID:-0}
LOG=${LOG:-lorenzod.log}

if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'simd'"
	exit 1
fi

export LORENZODHOME="/data/node${ID}/lorenzod"

if [ "$DEBUG" -eq 1 ]; then
  dlv --listen=:2345 --continue --headless=true --api-version=2 --accept-multiclient exec "${BINARY}" -- --home "${LORENZODHOME}" "$@"
elif [ "$DEBUG" -eq 1 ] && [ -d "$(dirname "${LORENZODHOME}"/"${LOG}")" ]; then
  dlv --listen=:2345 --continue --headless=true --api-version=2 --accept-multiclient exec "${BINARY}" -- --home "${LORENZODHOME}" "$@" | tee "${LORENZODHOME}/${LOG}"
elif [ -d "$(dirname "${LORENZODHOME}"/"${LOG}")" ]; then
  "${BINARY}" --home "${LORENZODHOME}" "$@" | tee "${LORENZODHOME}/${LOG}"
else
  "${BINARY}" --home "${LORENZODHOME}" "$@"
fi
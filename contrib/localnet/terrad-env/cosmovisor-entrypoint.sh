#!/usr/bin/env sh

BINARY=/terrad/${BINARY:-cosmovisor}
ID=${ID:-0}
LOG=${LOG:-terrad.log}

if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'terrad'"
	exit 1
fi

BINARY_CHECK="$(file "$BINARY" | grep 'ELF 64-bit LSB executable, x86-64')"

if [ -z "${BINARY_CHECK}" ]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi

export TERRADHOME="/terrad/node${ID}/terrad"

if [ -d "$(dirname "${TERRADHOME}"/"${LOG}")" ]; then
    "${BINARY}" run "$@" --home "${TERRADHOME}" | tee "${TERRADHOME}/${LOG}"
else
    "${BINARY}" run "$@" --home "${TERRADHOME}"
fi
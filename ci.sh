#!/bin/bash
set -e

# Check that the code has been formatted correctly.
GOFMT_DIFF=$(gofmt -s -d */*.go cmds webboot)
if [[ -n "${GOFMT_DIFF}" ]]; then
	echo 'Error: Go source code is not formatted:'
	printf '%s\n' "${GOFMT_DIFF}"
	echo 'Run `gofmt -s -w *.go pkg cmds'
	exit 1
fi

#!/bin/bash

APPDIR=/tmp/spectred-temp
SPECTRED_RPC_PORT=29587

rm -rf "${APPDIR}"

spectred --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${SPECTRED_RPC_PORT}" --profile=6061 &
SPECTRED_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${SPECTRED_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $SPECTRED_PID

wait $SPECTRED_PID
SPECTRED_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Spectred exit code: $SPECTRED_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SPECTRED_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1

#!/bin/bash
rm -rf /tmp/spectred-temp

spectred --devnet --appdir=/tmp/spectred-temp --profile=6061 --loglevel=debug &
SPECTRED_PID=$!

sleep 1

rpc-stability --devnet -p commands.json --profile=7000
TEST_EXIT_CODE=$?

kill $SPECTRED_PID

wait $SPECTRED_PID
SPECTRED_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Spectred exit code: $SPECTRED_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SPECTRED_EXIT_CODE -eq 0 ]; then
  echo "rpc-stability test: PASSED"
  exit 0
fi
echo "rpc-stability test: FAILED"
exit 1

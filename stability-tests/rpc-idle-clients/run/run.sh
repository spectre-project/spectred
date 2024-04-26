#!/bin/bash
rm -rf /tmp/spectred-temp

NUM_CLIENTS=128
spectred --devnet --appdir=/tmp/spectred-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
SPECTRED_PID=$!
SPECTRED_KILLED=0
function killSpectredIfNotKilled() {
  if [ $SPECTRED_KILLED -eq 0 ]; then
    kill $SPECTRED_PID
  fi
}
trap "killSpectredIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $SPECTRED_PID

wait $SPECTRED_PID
SPECTRED_EXIT_CODE=$?
SPECTRED_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "Spectred exit code: $SPECTRED_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SPECTRED_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1

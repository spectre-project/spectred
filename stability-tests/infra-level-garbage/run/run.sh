#!/bin/bash
rm -rf /tmp/spectred-temp

spectred --devnet --appdir=/tmp/spectred-temp --profile=6061 &
SPECTRED_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:18611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $SPECTRED_PID

wait $SPECTRED_PID
SPECTRED_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Spectred exit code: $SPECTRED_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SPECTRED_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1

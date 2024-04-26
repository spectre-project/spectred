#!/bin/bash
rm -rf /tmp/spectred-temp

spectred --simnet --appdir=/tmp/spectred-temp --profile=6061 &
SPECTRED_PID=$!

sleep 1

orphans --simnet -alocalhost:18511 -n20 --profile=7000
TEST_EXIT_CODE=$?

kill $SPECTRED_PID

wait $SPECTRED_PID
SPECTRED_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Spectred exit code: $SPECTRED_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SPECTRED_EXIT_CODE -eq 0 ]; then
  echo "orphans test: PASSED"
  exit 0
fi
echo "orphans test: FAILED"
exit 1

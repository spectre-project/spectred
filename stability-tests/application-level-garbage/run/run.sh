#!/bin/bash
rm -rf /tmp/spectred-temp

spectred --devnet --appdir=/tmp/spectred-temp --profile=6061 --loglevel=debug &
SPECTRED_PID=$!
SPECTRED_KILLED=0
function killSpectredIfNotKilled() {
    if [ $SPECTRED_KILLED -eq 0 ]; then
      kill $SPECTRED_PID
    fi
}
trap "killSpectredIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:18611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $SPECTRED_PID

wait $SPECTRED_PID
SPECTRED_KILLED=1
SPECTRED_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Spectred exit code: $SPECTRED_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SPECTRED_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1

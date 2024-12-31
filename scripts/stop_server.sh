#!/bin/bash

# Set the name of the process to kill
PROCESS_NAME="app"

# Attempt to kill the process
sudo killall "$PROCESS_NAME" 2>/dev/null

# Log the result
if [ $? -eq 0 ]; then
  echo "Process '$PROCESS_NAME' stopped successfully."
else
  echo "Process '$PROCESS_NAME' not found or failed to stop."
fi
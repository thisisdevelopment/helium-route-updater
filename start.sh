#!/bin/bash

while true; do
    $@
    echo "$@ stopped, restarting in 5 seconds"
    sleep 5
done

#!/bin/sh
npx react-scripts start 2>&1 | tee -a "/logs/${PROJECT_ID}.log"

#!/bin/sh
npx react-scripts start 2>&1 | tee "/logs/${PROJECT_ID}.log"

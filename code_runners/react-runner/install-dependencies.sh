#!/bin/sh
pnpm install 2>&1 | tee -a "/logs/${PROJECT_ID}.log"
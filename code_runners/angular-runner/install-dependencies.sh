#!/bin/sh
pnpm install --legacy-peer-deps 2>&1 | tee -a "/logs/${PROJECT_ID}.log"
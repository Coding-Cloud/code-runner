#!/bin/sh
pnpm install 2>&1 | tee "/logs/${PROJECT_ID}.log"
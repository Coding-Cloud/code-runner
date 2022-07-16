#!/bin/sh
pnpm config set store-dir=/root/.local/share/pnpm
pnpm install
npx react-scripts start 2>&1 | tee "/logs/${PROJECT_ID}.log"

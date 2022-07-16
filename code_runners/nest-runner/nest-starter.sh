#!/bin/sh
pnpm config set store-dir=/root/.local/share/pnpm
pnpm install
nest start --watch 2>&1 | tee "/logs/${PROJECT_ID}.log"

#!/bin/sh
pnpm config set store-dir=/root/.local/share/pnpm
pnpm install
ng serve --host "${HOST}" --port "${PORT}" --public-host "${PUBLIC_URL}" --disable-host-check --watch --poll 2000 2>&1 | tee "/logs/${PROJECT_ID}.log"

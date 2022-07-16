#!/bin/sh
ng serve --host "${HOST}" --port "${PORT}" --public-host "${PUBLIC_URL}" --disable-host-check --watch --poll 2000 2>&1 | tee "/logs/${PROJECT_ID}.log"

#!/bin/sh
ng serve --host "${HOST}" --port "${PORT}" --public-host "${PUBLIC_URL}" --disable-host-check 2>&1 | tee -a "/logs/${PROJECT_ID}.log"

#!/bin/sh
env-cmd -f .env nest start --watch 2>&1 | tee -a "/logs/${PROJECT_ID}.log"

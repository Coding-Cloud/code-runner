#!/bin/sh
nest start --watch 2>&1 | tee "/logs/${PROJECT_ID}.log"

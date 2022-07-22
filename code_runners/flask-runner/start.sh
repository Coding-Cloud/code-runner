#!/bin/sh
python -m flask run --host=${HOST} --port=${PORT} 2>&1 | tee -a "/logs/${PROJECT_ID}.log"

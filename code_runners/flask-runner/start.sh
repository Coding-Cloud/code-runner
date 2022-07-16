#!/bin/sh
python -m flask run --host=${HOST} --port=${PORT} 2>&1 | tee "/logs/${PROJECT_ID}.log"

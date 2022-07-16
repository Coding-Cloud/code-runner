#!/bin/sh
pip install -r requirements.txt --target=/libs
python -m flask run --host=${HOST} --port=${PORT} 2>&1 | tee "/logs/${PROJECT_ID}.log"

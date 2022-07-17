#!/bin/sh
kill -9 `ps faux | grep npx | awk -F' ' '{print $2}'`

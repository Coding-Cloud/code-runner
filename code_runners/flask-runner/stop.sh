#!/bin/sh
kill -9 `ps faux | grep python | awk -F' ' '{print $2}'`

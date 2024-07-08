#!/bin/bash

st ./tailwindcss-linux-x64 -i ./dist/main.css -o ./dist/tailwind.css --watch &
sleep 1
st ../go/bin/air &
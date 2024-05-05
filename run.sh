#!/bin/sh
air -c ./.air.toml & pnpm exec tailwindcss -i "cmd/views/styles.css" -o "public/styles.css" --watch

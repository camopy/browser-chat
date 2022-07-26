#!/bin/sh
if $DEV; then
  echo 'Runing migrations...'
  go run /browser-chat/cmd/migrate/main.go up > /dev/null 2>&1 &

  echo 'Start application...'
  go run /browser-chat/cmd/app/main.go
else
  echo 'Runing migrations...'
  /browser-chat/bin/migrate up > /dev/null 2>&1 &

  echo 'Start application...'
  /browser-chat/bin/app
fi
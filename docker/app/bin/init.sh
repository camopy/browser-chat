#!/bin/sh
echo 'Runing migrations...'
/browser-chat/bin/migrate up > /dev/null 2>&1 &

echo 'Start application...'
/browser-chat/bin/app
#!/bin/sh
curl -s 'https://api.github.com/users/kigiri' | jq -r '.id'

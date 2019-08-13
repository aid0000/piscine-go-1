#!bin/bash
curl -s 'https://api.github.com/users/MSilva95' | jq -r '.id'
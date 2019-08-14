#!/bin/sh
curl -s 'https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json' | jq '.[51].name'
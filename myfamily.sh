#!/bin/sh
curl -s 'https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json' | jq -r '.[ '${HERO_ID}-1' ].connections.relatives'
curl -s 'https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json' | jq ".[] | select (/id==$hero_id) | grep 'relatives' | cut -d '|'


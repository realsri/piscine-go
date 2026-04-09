curl -s https://01.gritlab.ax/assets/superhero/all.json | jq --arg id "$HERO_ID" '.[] | select(.id == ($id|tonumber)).connections.relatives' | tr -d '"'

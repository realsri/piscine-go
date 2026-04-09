curl -s https://assets.01-edu.org/superhero/all.json | jq -r '.[] | select(.id == 70) | .name'

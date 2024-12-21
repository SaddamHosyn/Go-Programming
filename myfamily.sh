
curl -s "https://01.gritlab.ax/assets/superhero/all.json"| jq --arg HERO_ID "$HERO_ID" '.[] | select(.id == ( $HERO_ID|tonumber)) | .connections.relatives' | sed 's/"//g; s/\n//g'

HERO_ID=1 curl  "https://01.gritlab.ax/assets/superhero/all.json"| jq '.[] | select(.id == "$HERO_ID") | .connections.relatives' | sed 's/"//g; s/\n//g'
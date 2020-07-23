#!/bin/bash
set -e

username=$1
SCORE=$2
AK=$3

id=$(curl "https://api.airtable.com/v0/appxj3f7gzCBEB9JL/tbllhRwtGFi3m1Dmx?fields%5B%5D=username&filterByFormula=%7Busername%7D+%3D+%27$username%27&maxRecords=1" \
  -H "Authorization: Bearer $AK" \
  | jq -c ".records[] | select(.fields.username == \"$username\").id" \
  | sed 's/"//g' \
  | grep .) \
  && {
    echo "$username already exist, id: $id."
  } || {
    echo "not found, create one"
    id=$(curl -X POST https://api.airtable.com/v0/appxj3f7gzCBEB9JL/Board \
      -H "Authorization: Bearer $AK" \
      -H "Content-Type: application/json" \
      --data "{
        \"records\": [
          {
            \"fields\": {
              \"username\": \"$username\"
            }
          }
        ]
      }" \
    | jq .records[].id \
    | sed 's/"//g')
  }

echo "update raw table."
curl -X POST https://api.airtable.com/v0/appxj3f7gzCBEB9JL/Raw \
  -H "Authorization: Bearer $AK" \
  -H "Content-Type: application/json" \
  --data "{
    \"records\": [
      {
        \"fields\": {
          \"TS\": \"$(date "+%Y-%m-%dT%H:%M:%S")\",
          \"Score\": $SCORE,
          \"Board\": [
            \"$id\"
          ]
        }
      }
    ]
  }" > /dev/null



    curl -v -s -S -X POST http://127.0.0.1:9000/alerts/notify \
    --header 'Content-Type: application/json; charset=utf-8' \
    --header 'Accept: application/json' \
    --header 'User-Agent: orion/0.10.0' \
    -d  '{
        "data": [
            {
                "id": "urn:ngsi-ld:Alert:MOD1",
                "type": "Alert",
                "humidity": "40",
                "temperature": "20",
                "condition": "values are too low",
                "dateObserved": "2019-07-16T14:00:00Z",
                "refModule": "urn:ngsi-ld:Alert:MOD1"
            }
        ],
        "subscriptionId": "57458eb60962ef754e7c0998"
    }'


    curl -v -s -S -X POST http://190.119.192.192:8080/sub \
    --header 'Content-Type: application/json; charset=utf-8' \
    --header 'Accept: application/json' \
    --header 'User-Agent: orion/0.10.0' \
    -d  '{
        "data": [
            {
                "id": "urn:ngsi-ld:Alert:MOD1",
                "type": "Alert",
                "humidity": "40",
                "temperature": "20",
                "condition": "values are too low"
                "dateObserved": "2019-07-16T14:00:00Z"
            }
        ],
        "subscriptionId": "57458eb60962ef754e7c0998"
    }'

curl -iX POST \
  --url 'http://localhost:1026/v2/entities' \
  --header 'Content-Type: application/json' \
  --data ' {
      "id":"urn:ngsi-ld:Alert:MOD1", "type":"Alert",
      "name":{"type":"Text", "value":"Alert"},
}'


curl localhost:1026/v2/entities/urn:ngsi-ld:Alert:MOD1/attrs?options=keyValues -s -S -H 'Content-Type: application/json' -X PUT -d '{
      "temperature": '29',
      "co2" : '100',
      "condition": "asdfjkashfklasdjhfklasdjhfklasdjfh"
}'


curl -iX POST \
  --url 'http://localhost:1026/v2/subscriptions' \
  --header 'Content-Type: application/json' \
  --data '{
  "description": "Notify Server ",
  "subject": {
    "entities": [{"idPattern": ".*", "type": "Alert"}]
  },
  "notification": {
    "http": {
      "url": "http://190.119.192.192:8080/sub"
    },
    "attrsFormat" : "keyValues"
  }
}'


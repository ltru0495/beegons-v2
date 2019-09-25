while true
do
  temp=$(shuf -i 35-53 -n 1);
  number=$(shuf -i 1-3113 -n 1);
  da=$(date +'%Y-%m-%dT%H:%M:%S.%6NZ')
  daO="${da}"
  echo "data sent"

    curl http://localhost:1026/v2/entities/urn:ngsi-ld:DataObserved:MOD1/attrs -s -S -H 'Content-Type: application/json' -X PUT -d '{
      "temperature": {
        "value": '$temp',
        "type": "Float"
      },
      "humidity": {
        "value": '$number',
        "type": "Float"
      },
      "dateObserved": {
        "type": "Text",
        "value":"'${daO}'"
      }

    }'
    sleep 1
done

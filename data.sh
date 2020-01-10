COUNTER="1"

while true
do
  temp=$(shuf -i 18-23 -n 1);
  hum=$(shuf -i 70-100 -n 1);
  da=$(date +'%Y-%m-%dT%H:%M:%S.%6NZ')
  daO="${da}"
  echo "data sent"
  echo $temp
  echo $daO



    curl http://localhost:1026/v2/entities/urn:ngsi-ld:DataObserved:MOD1/attrs -s -S -H 'Content-Type: application/json' -X POST -d '{
      "temperature": {
        "value": '$temp',
        "type": "Float"
      },
      "dateObserved": {
        "type": "Text",
        "value":"'${daO}'"
      },
      "humidity": {
        "type": "Float",
        "value": '$hum'
      }

    }'

    sleep 2.5
done

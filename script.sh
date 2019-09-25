while true;
do
temp=$(shuf -i 18-53 -n 1);
number=$(shuf -i 1-3113 -n 1);
da=$(date +'%Y-%m-%dT%H:%M:%S.%6NZ')
daO="${da}"
curl -v -s -S -X POST http://localhost:9001     --header 'Content-Type: application/json; charset=utf-8'     --header 'Accept: application/json'     --header 'User-Agent: orion/0.10.0'     --header "Fiware-Service: demo"     --header "Fiware-ServicePath: /test"     -d  '{
         "data": [
             {
                 "id": "urn:ngsi-ld:DataObserved:MOD1","type": "Module",
                 "temperature": {"type": "Float","value": '$temp',"metadata": {}},
                 "dateObserved": {"type": "Text", "value":"'${daO}'","metadata": {}},
                 "humidity": {"type": "Float","value": 40,"metadata": {}}
             }
         ],
         "subscriptionId": "57458eb60962ef754e7c0998"
     }';      sleep 1; done


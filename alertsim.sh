while true
do
  temp=$(shuf -i 45-53 -n 1);
  hum=$(shuf -i 70-100 -n 1);
  da=$(date +'%Y-%m-%dT%H:%M:%S.%6NZ')
  daO="${da}"
  echo "data sent"
  echo $temp
  echo $daO



  curl -v -s -S -X POST http://localhost:9000/data/notify \
  --header 'Content-Type: application/json; charset=utf-8' \
  --header 'Accept: application/json' \
  -d  '{
   "data": [
      {
       "id": "urn:ngsi-ld:DataObserved:MOD1","type": "DataObserved",
       "pressure": 40,
       "humidity":'$hum',
       "temperature": '$temp',
       "dateObserved": "'${daO}'"
    }
   ],
  "subscriptionId": "57458eb60962ef754e7c0998"
  }'



    sleep 2
done

# HOST=127.0.0.1
# HOSTNAME=beegons
# PORT=9000
# DB_USERNAME=
# DB_PASSWORD=
# DB_HOST=mongo-db
# DB_PORT=27017
# DB_DATABASE=test
# ORION_CB_HOST=orion
# ORION_CB_PORT=1026
# CYGNUS_HOST=cygnus
# CYGNUS_PORT=5050
# CYGNUS_DATABASE=sth_default
# FLINK_HOST=taskmanager
# FLINK_PORT=9001
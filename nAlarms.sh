N_ALARMS=$1

if [ -z "$1" ]
then
	echo "N_ALARMS not specified, set to 60."
	N_ALARMS=60
fi

echo "Número de alarmas: $N_ALARMS"
let COUNTER="1"
let MINUTE=60
DELAY="$(echo $MINUTE / $N_ALARMS | bc -l)"
echo "Envío de alerta cada $DELAY segundo(s)."

while [ $COUNTER -le $N_ALARMS ]
do
	echo "Número de alerta: $COUNTER"

	temp=$(shuf -i 45-53 -n 1);
  	da=$(date +'%Y-%m-%dT%H:%M:%S.%6NZ')
  	daO="${da}"

	curl http://localhost:1026/v2/entities/urn:ngsi-ld:DataObserved:MOD1/attrs -s -S -H 'Content-Type: application/json' -X POST -d '{
		"temperature": {
			"value": '$temp',
			"type": "Float"
		},
		"dateObserved": {
			"type": "Text",
			"value":"'${daO}'"
		}
	}'

	echo "Alerta enviada"
	sleep $DELAY;
	((COUNTER++))
done


echo "Finalizado."

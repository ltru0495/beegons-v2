$(function(){

    function MQTTconnect() {
        var theme = 'air';

        $.globalMessenger({ theme: theme });
        Messenger.options = { theme: theme  };
        mqttUser = "CTIC-SMARTCITY"
        mqttPwd = "YTICTRAMS-CITC"
        console.log("connecting ..."+mqtt_server+" "+mqtt_port);
        var time = new Date();
        mqtt = new Paho.MQTT.Client(mqtt_server, Number(mqtt_port), "clientjs"+time.getTime());
        var options = {
            timeout: 10,
            onSuccess: onConnect,
            userName : mqttUser,
            password: mqttPwd
        };
        mqtt.connect(options);
        mqtt.onMessageArrived = onMessageArrived;
    }

    function onConnect() {
        console.log("Connected");
        mqtt.subscribe("alert/+/+/+");
    }

    function onMessageArrived(message){
        console.log(message.payloadString);

        var obj = JSON.parse(message.payloadString);
        console.log(obj);
        message =  obj.message+" en "+obj.value +"\Hora: "+obj.timestamp    ;

        
        Messenger().post({
          message: message,
          type: "error"
        })
        
        
    }

    MQTTconnect();
    
});
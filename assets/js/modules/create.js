var nSensors = 0;
var sensors = [];

$('#addsensor').on('click', function(e) {
    e.preventDefault();
    appendSensor();
});

$('#create').on('click', function(e) {
    e.preventDefault();
    var mod = {};


    mod.name = "MOD1"
    mod.dataType = "AirQuality"
    mod.mac = "123ADASDA0012";
    mod.state = "OK";
    mod.supportedProtocol = ["HTTP"];
    mod.coordinates =  [-12.000123, -77.12312];


    mod.sensors = JSON.stringify(sensors);
    post("/module/create", mod, "post")

    e.stopPropagation();
});

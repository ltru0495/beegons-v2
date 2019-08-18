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
    mod.protocol = "HTTP";
    mod.latitude = -11.98164
    mod.longitude = -76.99925;

    sensors = [{
        name: "Temperatura",
        model: "SPEC",
        parameter: "temperature",
        dataType: mod.dataType,
        unit: "~"
    }, {
        name: "Humedad",
        model: "SPEC",
        parameter: "humidity",
        dataType: mod.dataType,
        unit: "~"
    }, {
        name: "Presion",
        model: "station",
        parameter: "pressure",
        dataType: mod.dataType,
        unit: "ppm"
    }, {
        name: "O3",
        model: "SPEC",
        parameter: "o3",
        dataType: mod.dataType,
        unit: "ppm"
    }, {
        name: "NO2",
        model: "SPEC",
        parameter: "no2",
        dataType: mod.dataType,
        unit: "ppm"
    }, {
        name: "H2S",
        model: "SPEC",
        parameter: "h2s",
        dataType: mod.dataType,
        unit: "ppm"
    }, {
        name: "CO",
        model: "SPEC",
        parameter: "co",
        dataType: mod.dataType,
        unit: "ppm"
    }, {
        name: "SO2",
        model: "SPEC",
        parameter: "so2",
        dataType: mod.dataType,
        unit: "ppm"
    }, {
        name: "Dirección del Viento",
        model: "station",
        parameter: "windDirection",
        dataType: mod.dataType,
        unit: "°"
    }, {
        name: "Velocidad de Viento",
        model: "station",
        parameter: "windSpeed",
        dataType: mod.dataType,
        unit: "km/h"
    }, {
        name: "UV",
        model: "station",
        parameter: "uv",
        dataType: mod.dataType,
        unit: "~"
    }, {
        name: "UVA",
        model: "station",
        parameter: "uva",
        dataType: mod.dataType,
        unit: "~"
    }, {
        name: "UVB",
        model: "station",
        parameter: "uvb",
        dataType: mod.dataType,
        unit: "~"
    },
    {
        name: "Calidad de Aire",
        model: "aqi",
        parameter: "airQualityIndex",
        dataType: mod.dataType,
        unit: "~"
    }
    ];



    mod.sensors = JSON.stringify(sensors);

    console.log(mod);

    post("/module/create", mod, "post")


    e.stopPropagation();
});

function appendSensor() {
    var time = new Date();
    var id = 'sensor' + nSensors + '_' + time.getTime();

    $('#sensors').append(
        '<div id="' + id + '" class="sensor">' +
        '<input type="text" id="type_' + id + '" class="form-control" data-placement="top" placeholder="Tipo">' +
        '<input type="text" id="name_' + id + '" class="form-control" data-placement="top" placeholder="Nombre">' +
        '<input type="text" id="unit_' + id + '" class="form-control" data-placement="top" placeholder="Unidad">' +
        '<input type="text" id="sensor_' + id + '" class="form-control" data-placement="top" placeholder="Marca">' +
        '</div>');

    $('#remove-sensor' + id).on('click', function(e) {
        e.preventDefault();
        var result = confirm('Desea eliminar el sensor: ?');
        if (result) {
            $('#' + id).remove();
            nSensors--;
        }
    });
}
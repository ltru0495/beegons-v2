var nSensors = 0;
var sensors = [];

$('#addsensor').on('click', function(e) {
    e.preventDefault();
    appendSensor();
});

$('#create').on('click', function(e) {
    e.preventDefault();
    var mod = {};
    sensors = [{
        name: "Temperatura",
        model: "SPEC",
        type: "temperature",
        unit: "~"
    }, {
        name: "Humedad",
        model: "SPEC",
        type: "humidity",
        unit: "~"
    }, {
        name: "Ozono (O3) ",
        model: "SPEC",
        type: "o3",
        unit: "ppm "
    }];

    
    mod.name ="MOD1"
    mod.type ="Module";
    mod.mac = "123ADASDA0012";
    mod.state = "OK";
    mod.protocol = "HTTP";
    mod.location = [-11.98164, -76.99925];

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
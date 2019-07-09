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
        type: "humedad",
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

    mod.sensors = JSON.stringify(sensors);

    console.log(mod);

    post("/module/create", mod, "post")
    // var sensorsGroup = $(".sensor");

    // let sensors = [];
    // let sensor = {};
    // for (let i =0 ; i < sensorsGroup.length; i ++) {

    //     id = sensorsGroup[i].id;
    //     sensor.type = $('#type_'+id).val();
    //     sensor.name = $('#name_'+id).val();
    //     sensor.unit = $('#unit_'+id).val();
    //     sensor.sensor = $('#sensor_'+id).val();


    //     sensors.push(sensor)

    // }

    // console.log(sensors);


    // var formData = {
    //     type: $("#type"),
    //     mac: $("#mac"),
    //     name: $("#name"),
    //     state: $("#state"),
    //     protocol: $("#protocol"),
    //     sensors: JSON.stringify(sensors)
    // };

    // var post_url = url_server + '/module/create';
    // $.ajaxSettings.traditional = true;

    // $.ajax({
    //     url: post_url,
    //     type: 'post',
    //     data: formData,
    //     success: function() {
    //         console.log('SUCCESSSS');
    //     },
    //     error: function(jqXHR, textStatus, errorThrown) {
    //         console.log('ERROR');
    //         console.log('jqXHR:');
    //         console.log(jqXHR);
    //         console.log('textStatus:');
    //         console.log(textStatus);
    //         console.log('errorThrown:');
    //         console.log(errorThrown);

    //     },
    //     complete: function() {
    //         window.location = '/module/register';
    //     }
    // });

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
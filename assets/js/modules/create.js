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

function appendSensor() {
    var time = new Date();
    var id = 'sensor' + nSensors + '_' + time.getTime();

    $('#sensors').append('<div id="'+id+'" class="sensor">'+
            '<div class="row"></div>'+
            '<div class="row">' +
                '<div class="col-md-3"></div>'+
                '<label class="col-md-6">Sensor</label>' +
            '</div>'+
            '<div class="form-group-sm row">' +
                '<label class="col-md-4"></label>'+
                '<label class="col-md-1 control-label text-md-right" for="sensorid_'+id+'">Id</label>'+
                '<div class="col-md-4">' +
                    '<input type="text" id="sensorid_'+id+'" class="form-control" data-placement="top" placeholder="Id">'+
                '</div>'+
            '</div>'+
            '<div class="form-group-sm row">' +
                '<label class="col-md-4"></label>'+
                '<label class="col-md-1 control-label text-md-right" for="type_'+id+'">Tipo</label>'+
                '<div class="col-md-4">' +
                    '<input type="text" id="type_'+id+'" class="form-control" data-placement="top" placeholder="Tipo">'+
                '</div>'+
            '</div>'+
            '<div class="form-group-sm row">' +
                '<label class="col-md-4"></label>'+
                '<label class="col-md-1 control-label text-md-right" for="label'+id+'">Label</label>'+
                '<div class="col-md-4">' +
                    '<input type="text" id="label_'+id+'" class="form-control" data-placement="top" placeholder="Label">'+
                '</div>'+
            '</div>'+
            '<div class="form-group-sm row">' +
                '<label class="col-md-4"></label>'+
                '<label class="col-md-1 control-label text-md-right" for="model_'+id+'">Modelo</label>'+
                '<div class="col-md-4">' +
                    '<input type="text" id="model_'+id+'" class="form-control" data-placement="top" placeholder="Modelo ">'+
                '</div>'+
            '</div>'+
            '<div class="form-group-sm row">' +
                '<label class="col-md-4"></label>'+
                '<label class="col-md-1 control-label text-md-right" for="state_'+id+'">Estado</label>'+
                '<div class="col-md-4">' +
                    '<input type="text" id="state_'+id+'" class="form-control" data-placement="top" placeholder="Estado">'+
                '</div>'+
            '</div>'+

            '<div class="form-group-sm row">' +
                '<label class="col-md-4"></label>'+
                '<label class="col-md-1 control-label text-md-right" for="unit_'+id+'">Unidad</label>'+
                '<div class="col-md-4">' +
                    '<input type="text" id="unit_'+id+'" class="form-control" data-placement="top" placeholder="Unidad">'+
                '</div>'+
            '</div>'+

            '<div class="form-group row">' +
                '<label class="col-md-4"></label>'+
                '<label class="col-md-2 control-label text-md-right">Eliminar</label>'+
                '<button id="remove-sensor'+id+'" class="fa fa-minus-circle"></button>' +
            '</div>'+
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
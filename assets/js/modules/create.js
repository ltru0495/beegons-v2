var nSensors = 0;
var sensors = [];

$('#addsensor').on('click', function(e) {
    e.preventDefault();
    appendSensor();
});


function getModuleFromForm() {
    return { 
        name: $("#name").val(),
        dataType: $("#dataType").val(),
        mac: $("#mac").val(),
        state: $("#state").val(),
        supportedProtocol: $("#protocol").val(),
    };
}

function checkModule(m) {
    if(m.name.split(" ").length> 1) return false;
    if(m.name === "") return false;
    if(m.dataType === "") return false;
    return true;
}


$('#create').on('click', function(e) {
    e.preventDefault();
    // var mod = {};
    // mod.name = "MOD1"
    // mod.dataType = "AirQuality"
    // mod.mac = "123ADASDA0012";
    // mod.state = "OK";
    // mod.supportedProtocol = ["HTTP"];
    // mod.coordinates =  [-12.000123, -77.12312];
    // post("/module/create", mod, "post")


    let mod = getModuleFromForm();
    if(!checkModule(mod)) {
        alert("Ha ocurrido un error");
        return;
    }
    post("/module/create", mod, "post");


    console.log(mod);
    e.stopPropagation();
});

$(function(){

    var map;
    var marker;
    var circle;
    $('#addsensor').on('click', function(e) {
        e.preventDefault();
        appendSensor();
    });

    function drawCircle(map,  marker) {
        if(circle != undefined) {
            circle.setMap(null);
        }

        circle = map.drawCircle({
            radius: 500,
            center: {lat: marker.latitude, lng: marker.longitude},
            strokeColor: '#0000FF',
            strokeOpacity: 0.3,
            fillColor:"#00FFFF",
            fillOpacity:0.2
        });
    }

    function getMarkerFromEvent(event) {
        var lat = event.latLng.lat();
        var lng = event.latLng.lng();
        marker = {latitude: lat ,longitude: lng};
        return marker;
    }

    
    function initGmap(){
        var map = new GMaps({
            el: '#gmap',
            lat: -12.000123,
            lng: -77.12312,
            zoomControl : false,
            panControl : false,
            streetViewControl : false,  
            mapTypeControl: false,
            overviewMapControl: false
        });
        
        $('#gmap').width('100%');
        $('#col-map').height('250px');

        map.on('click', function(event) {
            marker = getMarkerFromEvent(event);
            drawCircle(map, marker);
            map.removeMarkers();
            map.addMarker({
                lat: marker.latitude,
                lng: marker.longitude,
                draggable:true,
                dragend: function(event) {
                    marker = getMarkerFromEvent(event);        
                    drawCircle(map, marker)
                }   
            });
        });
    }

    

    function getModuleFromForm() {
        if (marker == undefined) {
            return {
                name: "" 
            }
        }
        return { 
            name: $("#name").val(),
            dataType: $("#type").val(),
            mac: $("#mac").val(),
            state: $("#state").val(),
            supportedProtocol: $("#protocol").val(),
            coordinates:[marker.latitude, marker.longitude],
        };
    }

    function checkModule(m) {
        if(m.name.split(" ").length> 1) return false;
        if(m.name === "") return false;
        if(m.dataType === "") return false;
        return true;
    }

    function moduleCreated(res) {
        
        if(res.status < 300 && res.status >= 200 ) {
            alert("Módulo creado correctamente");
        } else {
            alert("Error al crear módulo")
        }
    }


    $('#create').on('click', function(e) {
        e.preventDefault();
        // var mod = {};
        // mod.name = "MODTEST3"
        // mod.dataType = "AirQuality"
        // mod.mac = "123ADASDA0012";
        // mod.state = "OK";
        // mod.supportedProtocol = ["HTTP"];
        // mod.coordinates =  [-12.000123, -77.12312];
        // post("/module/create", mod, "post", moduleCreated)

        let mod = getModuleFromForm();
        if(!checkModule(mod)) {
            alert("Ha ocurrido un error");
            return;
        }
        post("/module/create", mod, "post", moduleCreated);


        console.log(mod);
        e.stopPropagation();
    });

    function pageLoad(){
        initGmap();
    }

    pageLoad();
    SingApp.onPageLoad(pageLoad);
});
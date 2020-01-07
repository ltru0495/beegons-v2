$(function(){

    var map;
    var marker;
    var circle;
 
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
            if(circle != undefined) {
                circle.setMap(null);
            }
            map.removeMarkers();
            map.addMarker({
                lat: marker.latitude,
                lng: marker.longitude,
                draggable:true,
                dragend: function(event) {
                    marker = getMarkerFromEvent(event);
                }   
            });
        });
    }

    

    function getParkingSpotFromForm() {
        if (marker == undefined) {
            return {
                name: "" 
            }
        }
        return { 
            name: $("#name").val(),
            type: $("#type").val(),
            mac: $("#mac").val(),
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

        let ps = getParkingSpotFromForm();
        if(!checkModule(ps)) {
            alert("Ha ocurrido un error");
            return;
        }
        post("/parkingspot/create", ps, "post", moduleCreated);


        console.log(ps);
        e.stopPropagation();
    });

    function pageLoad(){
        initGmap();
    }

    pageLoad();
    SingApp.onPageLoad(pageLoad);
});
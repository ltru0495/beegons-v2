$(function() {

    var map;
    var circle;

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

    function initGmap(){
        map = new GMaps({
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

    function getMarkerFromEvent(event) {
        var lat = event.latLng.lat();
        var lng = event.latLng.lng();
        marker = {latitude: lat ,longitude: lng};
        return marker;
    }

    function pageLoad() {
        initGmap();

        $('#moduleid').val("none");
        $("#parameter").val("none")

        $("#moduleid").on("change", function() {
            let selectedModule = $("#moduleid").val();
            dataId = selectedModule.replace("Module", "DataObserved");
            $.getJSON(url_server + `/api/module/${selectedModule}`, function(res) {
                if (res.status == 200) {
                    let mod = res.content;
                    let location = { latitude: mod.location.coordinates[0], longitude: mod.location.coordinates[1]}
                    map.removeMarkers();
                    map.setCenter({lat: location.latitude, lng: location.longitude})
                    map.addMarker({
                        lat: location.latitude,
                        lng: location.longitude,
                        draggable:true,
                        dragend: function(event) {
                            marker = getMarkerFromEvent(event);        
                            drawCircle(map, marker)
                        }   
                    });
                    drawCircle(map, location )
                }
            }).fail(function() {
                
            });
        });
    }

    pageLoad();
    SingApp.onPageLoad(pageLoad);
});
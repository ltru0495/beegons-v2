$(function(){
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

        $.getJSON(url_server + `/api/modules`, function(res) {
            let mods = res.content //Module With Data
            mods.forEach(mod => {
                
                addMarker(map, mod);
            });
        });
    }

    function addMarker(map, mod) {
        
        let marker={ latitude: mod.location.coordinates[0], longitude: mod.location.coordinates[1]};
        map.addMarker({
            lat: marker.latitude,
            lng: marker.longitude,
            click: (e) => {
                $.getJSON(url_server + `/api/module/${mod.id}/realtime`, function(res) {
                    let sensors = res.content;
                    if (sensors == null) return;
                    sensors.forEach(sensor => {
                        console.log(sensor)
                    });
                });
            }
        });
        drawCircle(map, marker)
    }

    function drawCircle(map,  marker) {
        circle = map.drawCircle({
            radius: 500,
            center: {lat: marker.latitude, lng: marker.longitude},
            strokeColor: '#0000FF',
            strokeOpacity: 0.3,
            fillColor:"#00FFFF",
            fillOpacity:0.2
        });
    }

    function pageLoad(){
        initGmap();
    }

    pageLoad();
    SingApp.onPageLoad(pageLoad);
});
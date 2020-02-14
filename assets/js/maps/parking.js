$(function(){


    if (window["WebSocket"]) {
        console.log("Connecting WS...")
        conn = new WebSocket(ws_server+"/ws/ps");

        conn.onclose = function(evt) {
            console.log("WS Connection closed");
        };
        conn.onmessage = function(evt) {
            var messageData = evt.data;
            // console.log(messageData)
            var data = JSON.parse(messageData);
            // console.log(data)
        };
    } else {
        console.log("Your Browser does not support WebSockets")
    }
    function initGmap(){
        var map = new GMaps({
            el: '#gmap',
            lat: -12.000123,
            lng: -77.12312,
            mapTypeId: 'hybrid',
            zoomControl : false,
            panControl : false,
            streetViewControl : false,
            mapTypeControl: false,
            overviewMapControl: false
        });

        $.getJSON(url_server + `/api/parkingspots`, function(res) {
            let spots = res.content
            spots.forEach(spot => {
                console.log(spot)
                if (spot.status) {
                    addMarker(map, spot, "green");
                } else {
                    addMarker(map, spot, "red");
                }
            });
        });
    }

    function addMarker(map, parkingSpot, color) {
        let marker={ latitude: parkingSpot.location.coordinates[0], longitude: parkingSpot.location.coordinates[1]};
        let url = "http://maps.google.com/mapfiles/ms/icons/"+color+"-dot.png";
        map.addMarker({
            lat: marker.latitude,
            lng: marker.longitude,
            click: (e) => {
                $.getJSON(url_server + `/api/parkingspot/${parkingSpot.id}`, function(res) {
                    let ps = res.content;
                    if (ps == null) return;
                    console.log(ps)
                });
            },
            icon: {
                url: url
            }
        });
    }



    function pageLoad(){
        initGmap();
    }

    pageLoad();
    SingApp.onPageLoad(pageLoad);
});

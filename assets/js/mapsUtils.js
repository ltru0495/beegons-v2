function drawCircle(map,  marker) {
    let circle;

    circle = map.drawCircle({
        radius: 500,
        center: {lat: marker.latitude, lng: marker.longitude},
        strokeColor: '#0000FF',
        strokeOpacity: 0.3,
        fillColor:"#00FFFF",
        fillOpacity:0.2
    });
    return circle;
}

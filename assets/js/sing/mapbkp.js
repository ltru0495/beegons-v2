// $(function(){
//     function pageLoad(){
//         var map = new ol.Map({
//         target: 'mapid',
//             layers: [
//                 new ol.layer.Tile({
//                     source: new ol.source.OSM()
//                 })
//             ],
//             view: new ol.View({

//                 // -77.043993, -12.076269
//                 center: ol.proj.fromLonLat([-77.049779, -12.016480]),
//                 zoom: 18
//             })
//         });

//         $.getJSON(url_server+'/api/v1/modules', function(modules) {
//             var polygons;
//             var path= [];
//             for(var i=0; i < modules.length;i ++) {
//                 polygons = modules[i]['polygons'];
//                 path = [];
//                 // console.log(polygons);
//                 for(var k=0; k< polygons.length; k++) {
//                     path.push([polygons[k].latitude, polygons[k].longitude]);

//                 }
//                 // polygon = map.drawPolygon({
//                 //     paths: path,
//                 //     strokeColor: '#BBD8E9',
//                 //     strokeOpacity: 1,
//                 //     strokeWeight: 3,
//                 //     fillColor: '#BBD8E9',
//                 //     fillOpacity: 0.6
//                 // });
//                 // map.addMarker({
//                 //   lat: modules[i].latitude,
//                 //   lng: modules[i].longitude,
//                 //   title: 'Marker #',
//                 //   infoWindow: {
//                 //     content : moduleContent(modules[i])
//                 //   }
//                 // });
//             }
//         });

//     }
//     pageLoad();
//     SingApp.onPageLoad(pageLoad);


// });



// var map;
// var polAdd = [];
// var pathAdd ;
// function moduleContent(module) {
//   var content = '<h4>Módulo '+ module.module_id+'</h4>'+
//     // '<p> Latitud:'+module.latitude+'</p>'+
//     // '<p> Longitud: '+module.longitude+'</p>'
//     '<h5>Sensores:</h5>';
//     for(var  i = 0 ; i < module.sensors.length; i ++) {
//       content += '<p style="font-size: 10px;">'+module.sensors[i].id_sensor+' '+ module.sensors[i].type+'</p>';
//     }
//   return content;
// }
// $(document).ready(function(){
//     map = new GMaps({
//         el: '#gmap',
//         lat: -12.0187326,
//         lng: -77.0508533,
//         zoom: 17.25,
//     });
//     $.getJSON(url_server+'/api/v1/modules', function(modules) {
//         var polygons;
//         var path= [];
//         for(var i=0; i < modules.length;i ++) {
//             polygons = modules[i]['polygons'];
//             path = [];
//             // console.log(polygons);
//             for(var k=0; k< polygons.length; k++) {
//                 path.push([polygons[k].latitude, polygons[k].longitude]);
//             }
//             polygon = map.drawPolygon({
//                 paths: path,
//                 strokeColor: '#BBD8E9',
//                 strokeOpacity: 1,
//                 strokeWeight: 3,
//                 fillColor: '#BBD8E9',
//                 fillOpacity: 0.6
//             });
//             map.addMarker({
//               lat: modules[i].latitude,
//               lng: modules[i].longitude,
//               title: 'Marker #',
//               infoWindow: {
//                 content : moduleContent(modules[i])
//               }
//             });
//         }
//     });
// });
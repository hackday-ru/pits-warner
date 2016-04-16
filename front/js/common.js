/**
 * Created by desiresdesigner on 4/16/16.
 */

$( document ).ready(function() {

    var map = L.map('map');
    var markers = [];

    function moveend() {
        //console.log("lat: " + map.getCenter().lat + "; lng: " + map.getCenter().lng + "; zoom: " + map.getZoom());
        var bounds = map.getBounds();
        var points = grable(bounds._northEast.lat, bounds._northEast.lng, bounds._southWest.lat, bounds._southWest.lng);
        //console.log(points.length);
        for (var i = 0, n = markers.length; i < n; ++i) {
            map.removeLayer(markers[i]);
        }
        markers = [];
        for (var i = 0, n = points.length; i < n; ++i) {
            //console.log(points[i].lat + "; " + points[i].lng);
            var marker = L.marker([Number(points[i].lat), Number(points[i].lng)]).addTo(map);
            markers.push(marker);
        }
    }
    
    map.on('moveend', moveend);

    map.setView([59.89444, 30.26417], 10); //[51.505, -0.09]

    var mapLayer = L.tileLayer('http://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        maxZoom: 18,
        attribution: 'Map data &copy; <a href="http://openstreetmap.org">OpenStreetMap</a> contributors, ' +
        '<a href="http://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, ' +
        'Imagery © <a href="http://mapbox.com">Mapbox</a>',
        id: 'examples.map-i86knfo3'
    }).addTo(map);

    //var searchLayer = L.tileLayer().addTo(map);
    //map.addControl( new L.Control.Search({layer: mapLayer}) );

    map.addControl( new L.Control.Search({
            url: 'https://nominatim.openstreetmap.org/search?format=json&q={s}',
            jsonpParam: 'json_callback',
            propertyName: 'display_name',
            propertyLoc: ['lat','lon'],
            markerLocation: true,
            autoType: false,
            autoCollapse: true,
            minLength: 2,
            zoom:16
    }));

    //mapLayer.on('load', moveend);
    //mapLayer.addTo(map);

    // var searchCtrl = L.control.fuseSearch()
    // searchCtrl.addTo(map);

});
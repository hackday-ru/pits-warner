/**
 * Created by desiresdesigner on 4/16/16.
 */

$( document ).ready(function() {

    var map = L.map('map');
    var markers = [];

    /*var marker_icon = L.icon({
        iconUrl: './img/marker1.png',
        iconSize: [16, 16]
    });*/

    var marker_icon = L.divIcon({className: 'icon'});
    
    function moveend() {
        var bounds = map.getBounds();
        for (var i = 0, n = markers.length; i < n; ++i) {
            map.removeLayer(markers[i]);
        }
        markers = [];

        $.ajax({
            type: 'GET',
            dataType: "json",
            url: 'http://52.58.116.75:8080/pits?lat0=' + bounds._northEast.lat + '&lon0=' + bounds._northEast.lng +
            '&lat1=' + bounds._southWest.lat + '&lon1=' + bounds._southWest.lng,
            success: function(data){
                console.log("success");
                $.each(data, function (key, val) {
                    //console.log(val.lat + "; " + val.lng);
                    markers.push(L.marker([Number(val.lat), Number(val.lng)], {icon: marker_icon}).addTo(map));
                })
            },
            error: function (err) {
                console.log("err");
                console.log(err);
            }
        });
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
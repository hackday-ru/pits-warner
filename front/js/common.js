/**
 * Created by desiresdesigner on 4/16/16.
 */

  

var App = {}

//App.host = 'http://52.58.116.75:8080'
App.host = 'http://localhost:8080'

App.updateMarkers = function(data) {
  for (var i = 0, n = markers.length; i < n; ++i) {
    map.removeLayer(markers[i]);
  }
  $.each(data, function (key, val) {
    addMarker(val)
  })
};

App.addMarker = function(p) {
  markers.push(
    L.marker([+p.Lat, +p.Lnt], {icon: App.markerIcon}).addTo(map)
  );
}

App.getPitsUrl = function() {
  var point = L.getCenter();
  var radius = 10000;
  App.postfix = '/pits?lng=%0.7f&lat=%0.7f&radius=%0.7f'.format(
    point.lat, point.lng, radius
  );
  return App.host + App.postfix;  
}

$( document ).ready(function() {

    var map = L.map('map');
    var markers = [];

    /*var marker_icon = L.icon({
        iconUrl: './img/marker1.png',
        iconSize: [16, 16]
    });*/

    App.markerIcon = L.divIcon({className: 'icon'});
    
    function moveend() {
        var bounds = map.getBounds();
        
        markers = [];

        $.ajax({
            type: 'GET',
            dataType: "json",
            url: App.getPitsUrl(),
            success: function(data){
              App.updateMarkers(data);
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

String.prototype.format = function() {
  var args = Array.prototype.slice.call(arguments);
  return sprintf.apply(sprintf, [this].concat(args));
};
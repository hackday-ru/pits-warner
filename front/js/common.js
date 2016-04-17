/**
 * Yamki, Epam hackday 2016
 */

var App = {}

//App.host = 'http://52.58.116.75:8080'
App.host = 'http://localhost:8080'

App.initPosition = [59.891132851001714, 30.316622257232662];
App.initZoom = 15;

App.updateMarkers = function(items) {
  for (var i = 0, n = App.markers.length; i < n; ++i) {
    App.map.removeLayer(markers[i]);
  }
  $.each(items, function (key, val) {
    App.addMarker(val)
  });
};

App.addMarker = function(p) {
  App.markers.push(
    L.marker([+p.Lat, +p.Lng], {icon: App.markerIcon}).addTo(App.map)
  );
}

App.getPitsUrl = function() {
  
  return 'testData/pits.json'
  
  var point = App.map.getCenter();
  var radius = 1000;
  App.postfix = '/pits?lng=%0.7f&lat=%0.7f&radius=%0.7f'.format(
    point.lat, point.lng, radius
  );
  return App.host + App.postfix;
}

$(function() {

    var map = App.map = L.map('map');
    App.markers = [];

    /*var marker_icon = L.icon({
        iconUrl: './img/marker1.png',
        iconSize: [16, 16]
    });*/

    App.markerIcon = L.divIcon({className: 'icon'});
    
    function moveend() {
      var bounds = map.getBounds();
      $.ajax({
          type: 'GET',
          dataType: "json",
          url: App.getPitsUrl(),
          success: function(data){
            App.updateMarkers(data.Items);
          },
          error: function (err) {
            console.log(err);
          }
      });
    }

    App.map.on('moveend', moveend);
    App.map.setView(App.initPosition, App.initZoom);

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
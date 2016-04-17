/**
 * Created by desiresdesigner on 4/16/16.
 */
var grable = function(northEast_lat, northEast_lng, southWest_lat, southWest_lng) {
    //console.log(northEast_lat + "; " + northEast_lng);
    //console.log(southWest_lat + "; " + southWest_lng);
    var result;
    
    $.ajax({
        type: 'GET',
        url: '52.58.116.75:8080/pits?lat0=' + northEast_lat + '&lon0=' + northEast_lng +
        '&lat1=' + southWest_lat + '&lon1=' + southWest_lng,
        // data: 'northEast_lat=' + northEast_lat + '&northEast_lng=' + northEast_lng +
        //         '&southWest_lat=' + southWest_lat + '&southWest_lng=' + southWest_lng,
        success: function(data){
            return data;
        },
        error: function (err) {
            return err;
        }
    });

    /*result = {
        0 : {
            "lat" : 59.89444,
            "lng" : 30.26417
        },
        1 : {
            "lat" : 59.9458321,
            "lng" : 30.4765999
        },
        3 : {
            "lat" : 59.8845205,
            "lng" : 29.8843764
        },
        4 : {
            "lat" : 60.010483,
            "lng" : 30.6571437
        }
    };*/
    
    //return result;
}
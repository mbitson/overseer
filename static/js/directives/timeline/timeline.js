/**
 * Created by mbitson on 6/6/2015.
 */
angular
    .module('overseerApp.timeline', [])
    .directive( 'osTimeline', function () {
        return {
            restrict:    'E',
            templateUrl: '/static/js/directives/timeline/timeline.html'
        };
    })
    .controller('timelineCtrl', function($scope, $mdDialog, Events) {
        $scope.init = function(){
            $scope.events = [];
            $scope.daysWithEvents = [];
            Events.get().success(function(data){
                $scope.events = data;
                $scope.daysWithEvents = uniqueBy($scope.events, function ( x ) {
                    var d = x.Created.split( /[- T :]/ );
                    return d[ 0 ] + '-' + d[ 1 ] + '-' + d[ 2 ];
                });
            });
        };

        $scope.init();
    });

function uniqueBy( arr, fn ) {
    var unique = {};
    var distinct = [];
    arr.forEach( function ( x ) {
        var key = fn( x );
        if ( !unique[ key ] ) {
            distinct.push( key );
            unique[ key ] = true;
        }
    } );
    return distinct;
}
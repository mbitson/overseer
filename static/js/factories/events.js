/**
 * Created by mbitson on 6/6/2015.
 */
angular.module('overseerApp.eventsFactory', [])
    .factory( 'Events', [ '$http', function ( $http ) {
        return {
            get:         function () {
                return $http.get( '/api/events' );
            },
            show:        function ( id ) {
                return $http.get( '/api/events/' + id );
            },
            create:      function ( data ) {
                return $http( {
                    method:  'POST',
                    url:     '/api/events',
                    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
                    data:    $.param( data )
                } );
            },
            destroy:     function ( id ) {
                return $http.delete( '/api/events/' + id );
            },
            update:      function ( id, data ) {
                return $http( {
                    method:  'PUT',
                    url:     '/api/events/' + id,
                    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
                    data:    $.param( data )
                } );
            }
        }
    } ] );
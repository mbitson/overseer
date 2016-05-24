/**
 * Created by mbitson on 6/6/2015.
 */
angular.module('monitorsFactory', [])
    .factory( 'Monitors', [ '$http', function ( $http ) {
        return {
            get:         function () {
                return $http.get( '/api/monitors' );
            },
            show:        function ( siteId ) {
                return $http.get( '/api/monitors/' + siteId );
            },
            create:      function ( data ) {
                return $http( {
                    method:  'POST',
                    url:     '/api/monitors',
                    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
                    data:    $.param( data )
                } );
            },
            destroy:     function ( id ) {
                return $http.delete( '/api/monitors/' + id );
            },
            update:      function ( id, data ) {
                return $http( {
                    method:  'PUT',
                    url:     '/api/monitors/' + id,
                    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
                    data:    $.param( data )
                } );
            }
        }
    } ] );
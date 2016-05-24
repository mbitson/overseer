/**
 * Created by mbitson on 6/6/2015.
 */
angular.module('overseerApp.sitesFactory', [])
    .factory( 'Sites', [ '$http', function ( $http ) {
        return {
            get:         function () {
                return $http.get( '/api/sites' );
            },
            show:        function ( id ) {
                return $http.get( '/api/sites/' + id );
            },
            create:      function ( data ) {
                return $http( {
                    method:  'POST',
                    url:     '/api/sites',
                    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
                    data:    $.param( data )
                } );
            },
            destroy:     function ( id ) {
                return $http.delete( '/api/sites/' + id );
            },
            update:      function ( id, data ) {
                return $http( {
                    method:  'PUT',
                    url:     '/api/sites/' + id,
                    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
                    data:    $.param( data )
                } );
            }
        }
    } ] );
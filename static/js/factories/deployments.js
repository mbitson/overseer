/**
 * Created by mbitson on 6/6/2015.
 */
angular.module('overseerApp.deploymentsFactory', [])
    .factory( 'Deployments', [ '$http', function ( $http ) {
        return {
            get:         function () {
                return $http.get( '/api/deployments' );
            },
            show:        function ( id ) {
                return $http.get( '/api/deployments/' + id );
            },
            create:      function ( data ) {
                return $http( {
                    method:  'POST',
                    url:     '/api/deployments',
                    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
                    data:    $.param( data )
                } );
            },
            destroy:     function ( id ) {
                return $http.delete( '/api/deployments/' + id );
            },
            update:      function ( id, data ) {
                return $http( {
                    method:  'PUT',
                    url:     '/api/deployments/' + id,
                    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
                    data:    $.param( data )
                } );
            }
        }
    } ] );
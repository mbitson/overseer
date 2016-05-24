/**
 * Created by mbitson on 6/6/2015.
 */
angular.module('overseerApp.applicationTypesFactory', [])
    .factory( 'ApplicationTypes', [ '$http', function ( $http ) {
        return {
            get:         function () {
                return $http.get( '/api/applicationTypes' );
            },
            show:        function ( id ) {
                return $http.get( '/api/applicationTypes/' + id);
            },
            create:      function ( data ) {
                return $http( {
                    method:  'POST',
                    url:     '/api/applicationTypes',
                    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
                    data:    $.param( data )
                } );
            },
            destroy:     function ( id ) {
                return $http.delete( '/api/applicationTypes/' + id );
            },
            update:      function ( id, data ) {
                return $http( {
                    method:  'PUT',
                    url:     '/api/applicationTypes/' + id,
                    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
                    data:    $.param( data )
                } );
            }
        }
    } ] );
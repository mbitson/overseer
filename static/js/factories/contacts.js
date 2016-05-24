/**
 * Created by mbitson on 6/6/2015.
 */
angular.module('overseerApp.contactsFactory', [])
    .factory( 'Contacts', [ '$http', function ( $http ) {
        return {
            get:         function () {
                return $http.get( '/api/contacts' );
            },
            show:        function ( id ) {
                return $http.get( '/api/contacts/' + id );
            },
            create:      function ( data ) {
                return $http( {
                    method:  'POST',
                    url:     '/api/contacts',
                    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
                    data:    $.param( data )
                } );
            },
            destroy:     function ( id ) {
                return $http.delete( '/api/contacts/' + id );
            },
            update:      function ( id, data ) {
                return $http( {
                    method:  'PUT',
                    url:     '/api/contacts/' + id,
                    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
                    data:    $.param( data )
                } );
            }
        }
    } ] );
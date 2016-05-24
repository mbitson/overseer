/**
 * Created by mbitson on 6/6/2015.
 */
angular.module('alertsFactory', [])
    .factory( 'Alerts', [ '$http', function ( $http ) {
        return {
            get:         function () {
                return $http.get( '/api/alerts' );
            },
            show:        function ( siteId ) {
                return $http.get( '/api/alerts/' + siteId );
            },
            recent: function ( siteId ) {
                return $http.get( '/api/alerts/recent/' + siteId );
            },
            create:      function ( data ) {
                return $http( {
                    method:  'POST',
                    url:     '/api/alerts',
                    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
                    data:    $.param( data )
                } );
            },
            destroy:     function ( id ) {
                return $http.delete( '/api/alerts/' + id );
            },
            update:      function ( id, data ) {
                return $http( {
                    method:  'PUT',
                    url:     '/api/alerts/' + id,
                    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
                    data:    $.param( data )
                } );
            }
        }
    } ] );
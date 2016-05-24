/**
 * Created by mbitson on 6/6/2015.
 */
angular.module('overseerApp.monitorRunsFactory', [])
    .factory( 'Runs', [ '$http', function ( $http ) {
        return {
            get:         function () {
                return $http.get( '/api/monitorRuns' );
            },
            show:        function ( runRequest ) {
                return $http.get( '/api/monitorRuns/' +
                runRequest.id + '/' +
                runRequest.date.startDate.toISOString().slice(0, 19).replace('T', '%20') + '/' +
                runRequest.date.endDate.toISOString().slice(0, 19).replace('T', '%20') );
            },
            create:      function ( data ) {
                return $http( {
                    method:  'POST',
                    url:     '/api/monitorRuns',
                    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
                    data:    $.param( data )
                } );
            },
            destroy:     function ( id ) {
                return $http.delete( '/api/monitorRuns/' + id );
            },
            update:      function ( id, data ) {
                return $http( {
                    method:  'PUT',
                    url:     '/api/monitorRuns/' + id,
                    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
                    data:    $.param( data )
                } );
            }
        }
    } ] );
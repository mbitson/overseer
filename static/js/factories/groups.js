/**
 * Created by mbitson on 6/6/2015.
 */
angular.module('overseerApp.groupsFactory', [])
    .factory( 'Groups', [ '$http', function ( $http ) {
        return {
            get:         function () {
                return $http.get( '/api/groups' );
            },
            show:        function ( id ) {
                return $http.get( '/api/groups/' + id );
            },
            create:      function ( data ) {
                return $http( {
                    method:  'POST',
                    url:     '/api/groups',
                    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
                    data:    $.param( data )
                } );
            },
            destroy:     function ( id ) {
                return $http.delete( '/api/groups/' + id );
            },
            update:      function ( id, data ) {
                return $http( {
                    method:  'PUT',
                    url:     '/api/groups/' + id,
                    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
                    data:    $.param( data )
                } );
            },
            saveAllContacts: function(groupsItems){
                var groups = angular.copy(groupsItems);
                angular.forEach(groups, function(group){
                    if(typeof group.Contacts === "object"){
                        group.Contacts = group.Contacts.join( ',' );
                    }
                });
                return $http( {
                    method:  'POST',
                    url:     '/api/groups/saveContacts',
                    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
                    data:    JSON.stringify(groups)
                } );
            }
        }
    } ] );
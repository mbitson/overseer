/**
 * Created by mbitson on 8/5/2015.
 */
angular
    .module('overseerApp.settingsContactsGroupCreateCtrl', [])
    .controller('settingsContactsGroupCreateCtrl', function($scope, Groups, $state, Contacts) {
        $scope.init = function(){
            $scope.getAllContacts();
        };

        $scope.getAllContacts = function () {
            $scope.contacts = [];
            Contacts.get().success( function ( data ) {
                $scope.contacts = data;
            } );
        };

        $scope.add = function(group){
            group.Contacts = group.Contacts.join(',');
            Groups.create( group ).success( function ( data ) {
                if ( data !== null && typeof data === 'object' ) {
                    $state.go( 'settings.contacts.list' );
                } else {
                    alert( data );
                }
            } );
        };

        $scope.init();
    });


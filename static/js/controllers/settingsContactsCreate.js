/**
 * Created by mbitson on 8/5/2015.
 */
angular
    .module('overseerApp.settingsContactsCreateCtrl', [])
    .controller('settingsContactsCreateCtrl', function($scope, Contacts, $state) {
        $scope.init = function(){
        };

        $scope.add = function(contact){
            Contacts.create( contact ).success( function ( data ) {
                if ( data !== null && typeof data === 'object' ) {
                    $state.go( 'settings.contacts.list' );
                } else {
                    alert( data );
                }
            } );
        };

        $scope.init();
    });


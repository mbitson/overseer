/**
 * Created by mbitson on 8/5/2015.
 */
angular
    .module('overseerApp.settingsContactsCtrl', [])
    .controller('settingsContactsCtrl', function($scope, Contacts, Groups) {
		/**
         * Init Function
         */
        $scope.init = function(){
            $scope.getAllContacts();
            $scope.getAllGroups();
        };

		/**
         * Load Functions
         */
        $scope.getAllContacts = function(){
            $scope.contacts = [];
            Contacts.get().success(function(data){
                angular.forEach( data, function ( contact ) {
                    delete contact.User;
                } );
                $scope.contacts = data;
            });
        };
        $scope.getAllGroups = function () {
            $scope.groups = [];
            Groups.get().success( function ( data ) {
                angular.forEach(data, function(group){
                    group.Contacts = group.Contacts.split(',' ).map(Number);
                });
                $scope.groups = data;
                $scope.assignGroupWatchers();
            } );
        };

        /*
         Watchers
         */
        $scope.assignGroupWatchers = function(){
            $scope.$watch( 'groups' , function (newValue, oldValue) {
                if(newValue !== oldValue){
                    Groups.saveAllContacts($scope.groups).success(function(){
                        $scope.success( 'Groups Saved Successfully' );
                    });
                }
            }, true );
        };

        /*
         Messaging user
         */
        $scope.success = function(string){
            toastr.success( string );
        };

        // Fire init!
        $scope.init();
    });


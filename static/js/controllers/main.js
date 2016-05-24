/**
 * Created by mbitson on 6/6/2015.
 */
angular.module('overseerApp.mainCtrl', []).
    controller('mainCtrl',  function($mdSidenav, $scope) {

        $scope.toggleSites = function(){
            $mdSidenav('sites')
                .toggle();
        };

    });
/**
 * Created by mbitson on 8/5/2015.
 */
angular
    .module('overseerApp.settingsCtrl', [])
    .controller('settingsCtrl', function($scope, $state) {
        $scope.init = function(){
            $scope.$state = $state;
        };

        $scope.init();
    });
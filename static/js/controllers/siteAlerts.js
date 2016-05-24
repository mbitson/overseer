/**
 * Created by mbitson on 6/6/2015.
 */
angular.module('overseerApp.siteAlertCtrl', []).
    controller('siteAlertCtrl',  function($scope, Alerts, $stateParams) {
        $scope.init = function(){
            $scope.alerts = [];
            $scope.loadAlerts();
        };

        $scope.loadAlerts = function(){
            Alerts.show($stateParams.siteId).success(function(data){
                $scope.alerts = data;
            });
        };

        $scope.init();
    });
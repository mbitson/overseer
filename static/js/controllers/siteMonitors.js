/**
 * Created by mbitson on 6/6/2015.
 */
angular.module('overseerApp.siteMonitorCtrl', []).
    controller('siteMonitorCtrl',  function($scope, Monitors, $stateParams, $interval) {
        $scope.init = function(){
            $scope.monitors = [];
            $scope.loadMonitors();
        };

        $scope.loadMonitors = function(){
            Monitors.show($stateParams.siteId).success(function(data){
                $scope.monitors = data;
            });
            $interval($scope.updateMonitorTimes, 500);
        };

        $scope.updateMonitorTimes = function(){
            angular.forEach($scope.monitors, function(monitor){
                monitor.percentTo = Math.floor($scope.timeToNextRun(monitor.Checked_date, monitor.Interval));
            });
        };

        $scope.timeToNextRun = function(lastRun, interval){
            var now = new Date().getTime() / 1000;
            var last = new Date(lastRun).getTime() / 1000;
            var since = now - last;
            var difference = interval - since;
            // TODO - Fix timezone inconsistencies.
            difference = difference - (60*60*4);
            return 100 - ((difference / interval) * 100);
        };

        $scope.init();
    });
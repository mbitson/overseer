/**
 * Created by mbitson on 6/6/2015.
 */
angular.module('overseerApp.siteCtrl', ["chart.js"]).
    controller('siteCtrl',  function($scope, $rootScope, $state, $stateParams, Sites, $location, Runs, $filter, Alerts) {
        $scope.$state = $state;
        $scope.init = function(){
            timeRounder = 1000*60*10; // Get time interval to round by
            now = new Date(); // Get now
            start = new Date(now); // Initialize start
            start.setHours(start.getHours()-3); // Compute start
            $scope.runRequest = {
                date:{
                    startDate: new Date( Math.round( start.getTime() / timeRounder ) * timeRounder ), // Set start date to rounded start
                    endDate: new Date( Math.round( now.getTime() / timeRounder ) * timeRounder ) // Set end date to rounded end
                },
                id: $stateParams.siteId
            };
            $scope.loadRuns();
        };

        $scope.loadRuns = function(){
            $scope.site = {};
            $scope.runs = [];
            $scope.alerts = [];
            Alerts.recent( $stateParams.siteId ).success( function ( data ) {
                $scope.alerts = data;
            } );
            Sites.show($stateParams.siteId).success(function(data){
                if(typeof data == 'object'){
                    $scope.site = data;
                }else{
                    $scope.abort();
                }
            });
            Runs.show($scope.runRequest).success(function(data){
                if(typeof data == 'object'){
                    $scope.runs = data;
                    $scope.updateCharts();
                }else{
                    $scope.abort();
                }
            });
        };

        $scope.$on( 'MONITOR.RUN', function ( eventDetails, event ) {
            var run = event.Data[0];
            if( $stateParams.siteId == run.Monitor.Site.Id){
                $scope.runs.push( event.Data[ 0 ] );
                $scope.updateCharts();
            }
        } );

        $scope.$on( 'SITE.UPDATE', function ( eventDetails, event ) {
            var site = event.Data[ 0 ];
            if ( $stateParams.siteId == site.Id ) {
                $scope.site = site;
            }
        } );

        $scope.updateCharts = function(){
            $scope.rebuildMonitoringChart();
            $scope.rebuildAvailabilityChart();
        };

        $scope.rebuildMonitoringChart = function(){
            $scope.MonitoringLabels = [];
            $scope.MonitoringSeries = ['Average Response Time', 'Average DNS Lookup Time', 'Downtime'];
            $scope.MonitoringData = [[], [], []];
            $scope.MonitoringLowest = Number.POSITIVE_INFINITY;
            $scope.MonitoringHighest = Number.NEGATIVE_INFINITY;
            angular.forEach($scope.runs, function (run) {
                if(typeof run.Response_time == 'string'){
                    run.Response_time = run.Response_time.replace( ',', '' );
                }
                if (run.Response_time < $scope.MonitoringLowest) $scope.MonitoringLowest = run.Response_time;
                if (run.Response_time > $scope.MonitoringHighest) $scope.MonitoringHighest = run.Response_time;
            });
            angular.forEach($scope.runs, function(run){
                var timeRun = new Date(run.Time_run);
                $scope.MonitoringLabels.push( (timeRun.getMonth() + 1) + "/" + timeRun.getDate() + " - " + timeRun.getHours() + ":" + timeRun.getMinutes() );
                $scope.MonitoringData[0].push(run.Response_time);
                if(run.Status_code == '200'){
                    $scope.MonitoringData[1].push(0);
                }else{
                    $scope.MonitoringData[1].push($scope.MonitoringHighest*.2);
                }
            });
        };

        $scope.rebuildAvailabilityChart = function(){
            $scope.AvailabilityColors = ['#6fc15f', '#c74c42'];
            $scope.AvailabilityLabels = ["Live", "Down"];
            $scope.AvailabilityData = [];
            var live = 0;
            var down = 0;
            angular.forEach($scope.runs, function(run, key){
                if(run.Status_code == '200'){
                    live++;
                }else{
                    down++;
                }
            });
            $scope.AvailabilityData.push(live);
            $scope.AvailabilityData.push(down);
            $scope.AvailabilityPercent = parseInt((live/($scope.runs.length))*100);
        };

        $scope.delete = function(site){
            Sites.destroy(site.Id).success(function(data){
                if(typeof data != 'undefined' && data == "Success"){
                    $scope.abort();
                }
            });
        };

        $scope.comingSoon = function(){
            alert("Coming soon!");
        };

        $scope.abort = function(){
            $rootScope.reloadSites();
            $location.path('/timeline');
        };

        $scope.init();
    });
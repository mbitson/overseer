angular
    .module('overseerApp.SiteAddCtrl', [])
    .controller('SiteAddCtrl', function($scope, $mdDialog, Sites, ApplicationTypes, Deployments, $location) {
        $scope.init = function(){
            $scope.site = {};
            $scope.applicationTypes = [];
            $scope.cmsIntegration = "off";
            $scope.deployment = {
                type: "FTPS"
            };
            $scope.application = {
                type: 0
            };
            ApplicationTypes.get().success(function(data){
                $scope.applicationTypes = data;
            });
        };
        $scope.selectApplicationType = function(applicationType){
            $scope.application.type = applicationType.Id;
        };
        $scope.add = function(site) {
            Sites.create(site).success(function(data)
            {
                if(data !== null && typeof data === 'object'){
                    $location.path( '/site/'+data.Id+'/dashboard' );
                }else{
                    alert(data);
                }
            });
        };
        $scope.init();
    });
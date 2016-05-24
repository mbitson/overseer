/**
 * Created by mbitson on 6/6/2015.
 */
angular
    .module('overseerApp.siteNavigationCtrl', [])
    .controller('siteNavigationCtrl', function($scope, $mdDialog, Sites, $mdSidenav, $rootScope, $mdToast, $location) {
        $scope.init = function(){
            $scope.sites = [];
            Sites.get().success(function(data){
                $scope.sites = data;
            });
        };

        $scope.$on( 'SITE.UPDATE', function ( eventDetails, event ) {
            var site = event.Data[ 0 ];
            $scope.updateSite(site);
        } );

        $scope.updateSite = function(site){
            for ( var i = 0; i < $scope.sites.length; i++ ) {
                if ( $scope.sites[ i ].Id === site.Id ) {
                    $scope.processToast($scope.sites[i], site);
                    $scope.sites[ i ] = site;
                    return true;
                }
            }
            return false;
        };

        $scope.processToast = function(oldsite, newsite){
            if(oldsite.Status == 1){
                if(newsite.Status == 0){
                    // DISPLAY NOW SITE IS DOWN NOTICE
                    $scope.showSiteActionToast( newsite.Domain+" now down!", newsite.Id);
                }
            }else{
                if ( newsite.Status == 1 ) {
                    // DISPLAY NOW SITE IS UP NOTICE
                    $scope.showSiteActionToast( newsite.Domain + " now up!", newsite.Id );
                }
            }
        };

        $scope.showSiteActionToast = function(message, id){
            var toast = $mdToast.simple()
                .content( message )
                .action( 'View' )
                .position( {
                    bottom: true,
                    top:    false,
                    left:   false,
                    right:  true
                } );
            $mdToast.show( toast ).then( function () {
                $location.path( '/site/' + id + '/dashboard' );
            } );
        };

        $rootScope.reloadSites = $scope.init;

        $scope.toggleSites = function(){
            $mdSidenav('sites')
                .toggle();
        };

        $scope.hasDownSites = function(){
            var flag = false;
            angular.forEach($scope.sites, function(site){
                if(site.Status == 0){
                    flag = true;
                }
            });
            return flag;
        };

        $scope.init();
    });
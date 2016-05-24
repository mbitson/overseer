/**
 * Created by mbitson on 6/6/2015.
 */
var overseer = angular.module('authApp', [
    'ngMaterial',
    'ngAnimate',
    'ngMdIcons'
]);
overseer.config(function ($mdThemingProvider) {
    $mdThemingProvider.definePalette('clear', {
        "50"  : "#FFFFFF",
        "100" : "#FFFFFF",
        "200" : "#FFFFFF",
        "300" : "#FFFFFF",
        "400" : "#FFFFFF",
        "500" : "#FFFFFF",
        "600" : "#cbcaca",
        "700" : "#aeadad",
        "800" : "#919090",
        "900" : "#747474",
        "A100": "#f8f8f8",
        "A200": "#f4f3f3",
        "A400": "#ecebeb",
        "A700": "#aeadad"
    });
    $mdThemingProvider.theme('default')
        .primaryPalette('red')
        .accentPalette('clear');
});
overseer.controller('authCtrl', function($scope){
    $scope.init = function(){
        $scope.username = '';
        $scope.password = '';
    };
    $scope.init();
});
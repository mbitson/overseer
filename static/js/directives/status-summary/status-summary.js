/**
 * Created by mbitson on 6/6/2015.
 */
angular
	.module( 'overseerApp.statusSummary', [] )
	.directive( 'statusSummary', function () {
		return {
			restrict:    'E',
			templateUrl: '/static/js/directives/status-summary/status-summary.html'
		};
	} )
	.controller( 'statusSummaryCtrl', function ( $scope, $mdDialog, Events ) {
		$scope.init = function () {
		};

		$scope.init();
	} );
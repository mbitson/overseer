/**
 * Created by mbitson on 6/6/2015.
 */
var overseer = angular.module( 'overseerApp', [
	'oc.lazyLoad',
	'ngWebSocket',
	'overseerApp.liveFactory',
	'overseerApp.monitorRunsFactory',
	'overseerApp.mainCtrl',
	'overseerApp.sitesFactory',
	'overseerApp.siteNavigationCtrl',
	'ui.router',
	'ngMaterial',
	'ngAnimate',
	'ngMdIcons',
	'chart.js'
] );
overseer.config( function ( $stateProvider, $urlRouterProvider, $mdThemingProvider ) {
	$stateProvider.
		state( 'site', {
			url:         '/site/:siteId',
			templateUrl: '/static/views/sites/view.tmpl.html',
			controller:  'siteCtrl',
			resolve:     {
				loadSites:      function ( $ocLazyLoad ) {
					return $ocLazyLoad.load( [
						{ files: [ '/static/js/controllers/site.js' ] }
					] );
				},
				loadSiteAlerts: function ( $ocLazyLoad ) {
					return $ocLazyLoad.load( [
						{ files: [ '/static/js/controllers/siteAlerts.js' ] },
						{ files: [ '/static/js/factories/alerts.js' ] }
					] );
				}
			}
		} ).
		state( 'site.dashboard', {
			url:         '/dashboard',
			templateUrl: '/static/views/sites/pages/dashboard.tmpl.html',
			resolve:     {
				loadDatePicker: function ( $ocLazyLoad ) {
					return $ocLazyLoad.load( [
						{ files: [ '/static/node_modules/bootstrap/dist/css/bootstrap.min.css' ] }
					] )
				}
			}
		} ).
		state( 'site.monitors', {
			url:         '/monitors',
			templateUrl: '/static/views/sites/pages/monitors.tmpl.html',
			controller:  'siteMonitorCtrl',
			resolve:     {
				loadSiteMonitors: function ( $ocLazyLoad ) {
					return $ocLazyLoad.load( [
						{ files: [ '/static/js/controllers/siteMonitors.js' ] },
						{ files: [ '/static/js/factories/monitors.js' ] }
					] );
				}
			}
		} ).
		state( 'site.alerts', {
			url:         '/alerts',
			templateUrl: '/static/views/sites/pages/alerts.tmpl.html',
			controller:  'siteAlertCtrl',
			resolve:     {
				loadSiteAlerts: function ( $ocLazyLoad ) {
					return $ocLazyLoad.load( [
						{ files: [ '/static/js/controllers/siteAlerts.js' ] },
						{ files: [ '/static/js/factories/alerts.js' ] }
					] );
				}
			}
		} ).
		state( 'create', {
			url:         '/create',
			templateUrl: '/static/views/sites/add.tmpl.html',
			controller:  'SiteAddCtrl',
			resolve:     {
				loadCreate: function ( $ocLazyLoad ) {
					return $ocLazyLoad.load( [
						{
							files: [
								'/static/js/controllers/siteCreate.js',
								'/static/js/factories/applicationTypes.js',
								'/static/js/factories/deployments.js'
							]
						}
					] );
				}
			}
		} ).
		state( 'dashboard', {
			url:         '/dashboard',
			templateUrl: '/static/views/pages/dashboard.html',
			controller:  'dashboardCtrl',
			resolve:     {
				loadDashboard:              function ( $ocLazyLoad ) {
					return $ocLazyLoad.load( [
						{ files: [ '/static/js/controllers/dashboard.js' ] }
					] );
				},
				loadTimeline:               function ( $ocLazyLoad ) {
					return $ocLazyLoad.load( [
						{ files: [ '/static/js/directives/timeline/timeline.js' ] },
						{ files: [ '/static/js/factories/events.js' ] }
					] );
				},
				loadStatusSummaryDirective: function ( $ocLazyLoad ) {
					return $ocLazyLoad.load( [
						{ files: [ '/static/js/directives/status-summary/status-summary.js' ] },
						{ files: [ '/static/js/controllers/siteNavigation.js' ] }
					] )
				}
			}
		} ).
		state( 'settings', {
			url:         '/settings',
			templateUrl: '/static/views/settings/view.html',
			controller:  'settingsCtrl',
			resolve:     {
				loadSettings: function ( $ocLazyLoad ) {
					return $ocLazyLoad.load( [
						{ files: [ '/static/js/controllers/settings.js' ] }
					] );
				}
			}
		} ).
		state( 'settings.account', {
			url:         '/account',
			templateUrl: '/static/views/settings/account.html',
			controller:  'settingsAccountCtrl',
			resolve:     {
				loadSettings: function ( $ocLazyLoad ) {
					return $ocLazyLoad.load( [
						{ files: [ '/static/js/controllers/settingsAccount.js' ] }
					] );
				}
			}
		} ).
		state( 'settings.contacts', {
			url:         '/contacts',
			templateUrl: '/static/views/settings/contacts/view.html'
		} ).
		state( 'settings.contacts.list', {
			url:         '/list',
			templateUrl: '/static/views/settings/contacts/list.html',
			controller:  'settingsContactsCtrl',
			resolve:     {
				loadSettings: function ( $ocLazyLoad ) {
					return $ocLazyLoad.load( [
						{
							files: [
								'/static/js/controllers/settingsContacts.js',
								'/static/js/factories/contacts.js',
								'/static/js/factories/groups.js',
								'/static/node_modules/chosenjs/chosen.min.css',
								'/static/node_modules/chosenjs/chosen.jquery.min.js',
								'/static/node_modules/angular-chosen-localytics/dist/angular-chosen.min.js',
								'/static/node_modules/toastr/build/toastr.min.css',
								'/static/node_modules/toastr/build/toastr.min.js'
							]
						}
					] );
				}
			}
		} ).
		state( 'settings.contacts.create', {
			url:         '/create',
			templateUrl: '/static/views/settings/contacts/create.html',
			controller:  'settingsContactsCreateCtrl',
			resolve:     {
				loadSettings: function ( $ocLazyLoad ) {
					return $ocLazyLoad.load( [
						{ files: [ '/static/js/controllers/settingsContactsCreate.js', '/static/js/factories/contacts.js' ] }
					] );
				}
			}
		} ).
		state( 'settings.contacts.group', {
			url:         '/group',
			templateUrl: '/static/views/settings/contacts/group.html',
			controller:  'settingsContactsGroupCreateCtrl',
			resolve:     {
				loadSettings: function ( $ocLazyLoad ) {
					return $ocLazyLoad.load( [
						{ files: [
							'/static/node_modules/chosenjs/chosen.min.css',
							'/static/node_modules/chosenjs/chosen.jquery.min.js',
							'/static/node_modules/angular-chosen-localytics/dist/angular-chosen.min.js',
							'/static/js/controllers/settingsContactsGroupCreate.js',
							'/static/js/factories/groups.js',
							'/static/js/factories/contacts.js'
						] }
					] );
				}
			}
		} );
	$urlRouterProvider.otherwise( '/dashboard' );
	$mdThemingProvider.definePalette( 'clear', {
		"50":   "#FFFFFF",
		"100":  "#FFFFFF",
		"200":  "#FFFFFF",
		"300":  "#FFFFFF",
		"400":  "#FFFFFF",
		"500":  "#FFFFFF",
		"600":  "#cbcaca",
		"700":  "#aeadad",
		"800":  "#919090",
		"900":  "#747474",
		"A100": "#f8f8f8",
		"A200": "#f4f3f3",
		"A400": "#ecebeb",
		"A700": "#aeadad"
	} );
	$mdThemingProvider.theme( 'default' ).primaryPalette( 'red' ).accentPalette( 'clear' );
	$mdThemingProvider.definePalette( 'dark', {
		"50":                   "#f1f1f1",
		"100":                  "#d4d4d4",
		"200":                  "#b7b7b7",
		"300":                  "#9e9e9e",
		"400":                  "#868686",
		"500":                  "#6e6e6e",
		"600":                  "#606060",
		"700":                  "#535353",
		"800":                  "#454545",
		"900":                  "#373737",
		"A100":                 "#d4d4d4",
		"A200":                 "#b7b7b7",
		"A400":                 "#868686",
		"A700":                 "#535353",
		'contrastDefaultColor': 'light'
	} );
	$mdThemingProvider.definePalette( 'darkish', {
		"50":                   "#fbfbfb",
		"100":                  "#f1f1f1",
		"200":                  "#e9e9e9",
		"300":                  "#e1e1e1",
		"400":                  "#d9d9d9",
		"500":                  "#d2d2d2",
		"600":                  "#b8b8b8",
		"700":                  "#9e9e9e",
		"800":                  "#838383",
		"900":                  "#696969",
		"A100":                 "#f1f1f1",
		"A200":                 "#e9e9e9",
		"A400":                 "#d9d9d9",
		"A700":                 "#9e9e9e",
		'contrastDefaultColor': 'light'
	} );
	$mdThemingProvider.theme( 'darcula' ).primaryPalette( 'dark' ).accentPalette( 'darkish' );
} );
'use strict';
var ws;
angular.module( 'overseerApp.liveFactory', [])
	.run( function ( $websocket, $rootScope ) {
		ws = $websocket( 'ws://go.mbitson.com/api/live/join' );

		ws.onMessage( function ( message ) {
			event = JSON.parse( message.data ) ;
			$rootScope.$broadcast(event.Name, event);
		} );

		// If the WebSocket is closed...
		ws.onClose( function ()
		{
			// Redirect user to login!
			window.location.href="/user/login";
		} );
	} );
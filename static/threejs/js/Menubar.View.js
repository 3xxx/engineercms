import { UIPanel, UIRow } from './libs/ui.js';

function MenubarView( editor ) {

	const strings = editor.strings;

	const container = new UIPanel();
	container.setClass( 'menu' );

	const title = new UIPanel();
	title.setClass( 'title' );
	title.setTextContent( strings.getKey( 'menubar/view' ) );
	container.add( title );

	const options = new UIPanel();
	options.setClass( 'options' );
	container.add( options );

	// Fullscreen

	const option = new UIRow();
	option.setClass( 'option' );
	option.setTextContent( strings.getKey( 'menubar/view/fullscreen' ) );
	option.onClick( function () {

		if ( document.fullscreenElement === null ) {

			document.documentElement.requestFullscreen();

		} else if ( document.exitFullscreen ) {

			document.exitFullscreen();

		}

		// Safari

		if ( document.webkitFullscreenElement === null ) {

			document.documentElement.webkitRequestFullscreen();

		} else if ( document.webkitExitFullscreen ) {

			document.webkitExitFullscreen();

		}

	} );
	options.add( option );

	// VR (Work in progress)

	if ( 'xr' in navigator ) {

		navigator.xr.isSessionSupported( 'immersive-vr' )
			.then( function ( supported ) {

				if ( supported ) {

					const option = new UIRow();
					option.setClass( 'option' );
					option.setTextContent( 'VR' );
					option.onClick( function () {

						editor.signals.toggleVR.dispatch();

					} );
					options.add( option );

				}

			} );

	}

	return container;

}

export { MenubarView };

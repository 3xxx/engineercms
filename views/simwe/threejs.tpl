<!DOCTYPE html>
<html lang="en">

<head>
  <title>three.js editor</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, user-scalable=no, minimum-scale=1.0, maximum-scale=1.0">
  <link rel="apple-touch-icon" href="/static/threejs/images/icon.png">
  <link rel="manifest" href="/static/threejs/manifest.json">
  <link rel="shortcut icon" href="/static/threejs/files/favicon_white.ico" media="(prefers-color-scheme: dark)" />
  <link rel="shortcut icon" href="/static/threejs/files/favicon.ico" media="(prefers-color-scheme: light)" />
</head>

<body>
  <link rel="stylesheet" href="/static/threejs/css/main.css">
  <!-- <script src="https://unpkg.com/@ffmpeg/ffmpeg@0.9.6/dist/ffmpeg.min.js" defer></script> -->
  <script src="/static/threejs/examples/js/libs/draco/draco_encoder.js"></script>
  <link rel="stylesheet" href="/static/threejs/js/libs/codemirror/codemirror.css">
  <link rel="stylesheet" href="/static/threejs/js/libs/codemirror/theme/monokai.css">
  <script src="/static/threejs/js/libs/codemirror/codemirror.js"></script>
  <script src="/static/threejs/js/libs/codemirror/mode/javascript.js"></script>
  <script src="/static/threejs/js/libs/codemirror/mode/glsl.js"></script>
  <script src="/static/threejs/js/libs/esprima.js"></script>
  <script src="/static/threejs/js/libs/jsonlint.js"></script>
  <link rel="stylesheet" href="/static/threejs/js/libs/codemirror/addon/dialog.css">
  <link rel="stylesheet" href="/static/threejs/js/libs/codemirror/addon/show-hint.css">
  <link rel="stylesheet" href="/static/threejs/js/libs/codemirror/addon/tern.css">
  <script src="/static/threejs/js/libs/codemirror/addon/dialog.js"></script>
  <script src="/static/threejs/js/libs/codemirror/addon/show-hint.js"></script>
  <script src="/static/threejs/js/libs/codemirror/addon/tern.js"></script>
  <script src="/static/threejs/js/libs/acorn/acorn.js"></script>
  <script src="/static/threejs/js/libs/acorn/acorn_loose.js"></script>
  <script src="/static/threejs/js/libs/acorn/walk.js"></script>
  <script src="/static/threejs/js/libs/ternjs/polyfill.js"></script>
  <script src="/static/threejs/js/libs/ternjs/signal.js"></script>
  <script src="/static/threejs/js/libs/ternjs/tern.js"></script>
  <script src="/static/threejs/js/libs/ternjs/def.js"></script>
  <script src="/static/threejs/js/libs/ternjs/comment.js"></script>
  <script src="/static/threejs/js/libs/ternjs/infer.js"></script>
  <script src="/static/threejs/js/libs/ternjs/doc_comment.js"></script>
  <script src="/static/threejs/js/libs/tern-threejs/threejs.js"></script>
  <script src="/static/threejs/js/libs/signals.min.js"></script>
  <!-- Import maps polyfill -->
  <!-- Remove this when import maps will be widely supported -->
  <!-- <script async src="https://unpkg.com/es-module-shims@1.3.6/dist/es-module-shims.js"></script> -->
  <script type="importmap">
    {
				"imports": {
					"three": "/static/threejs/build/three.module.js"
				}
			}
		</script>
  <script type="module">
    import * as THREE from 'three';
		import { Editor } from '/static/threejs/js/Editor.js';
		import { Viewport } from '/static/threejs/js/Viewport.js';
		import { Toolbar } from '/static/threejs/js/Toolbar.js';
		import { Script } from '/static/threejs/js/Script.js';
		import { Player } from '/static/threejs/js/Player.js';
		import { Sidebar } from '/static/threejs/js/Sidebar.js';
		import { Menubar } from '/static/threejs/js/Menubar.js';
		import { Resizer } from '/static/threejs/js/Resizer.js';
		import { VRButton } from '/static/threejs/examples/jsm/webxr/VRButton.js';

		import { GLTFLoader } from '/static/threejs/examples/jsm/loaders/GLTFLoader.js';
		import { RGBELoader } from '/static/threejs/examples/jsm/loaders/RGBELoader.js';

		window.URL = window.URL || window.webkitURL;
		window.BlobBuilder = window.BlobBuilder || window.WebKitBlobBuilder || window.MozBlobBuilder;
		Number.prototype.format = function () {
			return this.toString().replace( /(\d)(?=(\d{3})+(?!\d))/g, '$1,' );
		};
		//
		const editor = new Editor();
		window.editor = editor; // Expose editor to Console
		window.THREE = THREE; // Expose THREE to APP Scripts and Console
		window.VRButton = VRButton; // Expose VRButton to APP Scripts
		const viewport = new Viewport( editor );
		document.body.appendChild( viewport.dom );
		const toolbar = new Toolbar( editor );
		document.body.appendChild( toolbar.dom );
		const script = new Script( editor );
		document.body.appendChild( script.dom );
		const player = new Player( editor );
		document.body.appendChild( player.dom );
		const sidebar = new Sidebar( editor );
		document.body.appendChild( sidebar.dom );
		const menubar = new Menubar( editor );
		document.body.appendChild( menubar.dom );
		const resizer = new Resizer( editor );
		document.body.appendChild( resizer.dom );


	// const camera = editor.camera;
	const scene = editor.scene;
	const sizes = {
	  width: window.innerWidth,
	  height: window.innerHeight,
	}
	/**
	 * Light
	 */
	const directionLight = new THREE.DirectionalLight()
	directionLight.castShadow = true
	directionLight.position.set(5, 5, 6)
	directionLight.shadow.camera.near = 1
	directionLight.shadow.camera.far = 20
	directionLight.shadow.camera.top = 10
	directionLight.shadow.camera.right = 10
	directionLight.shadow.camera.bottom = -10
	directionLight.shadow.camera.left = -10

	const directionLightHelper = new THREE.DirectionalLightHelper(directionLight, 2)
	directionLightHelper.visible = false
	scene.add(directionLightHelper)
	
	const directionalLightCameraHelper = new THREE.CameraHelper(directionLight.shadow.camera)
	directionalLightCameraHelper.visible = false
	scene.add(directionalLightCameraHelper)
	
	const ambientLight = new THREE.AmbientLight(new THREE.Color('#ffffff'), 0.3)
	scene.add(ambientLight, directionLight)
	
	scene.add(new THREE.AmbientLight(0x666666));//环境光

	// Renderer
	const renderer = new THREE.WebGLRenderer({
	  // canvas,
	  antialias: true,
	})
	renderer.setSize(sizes.width, sizes.height)
	renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2))
	renderer.shadowMap.enabled = true

	const gltfLoader = new GLTFLoader().setPath( '/static/threejs/' );
	gltfLoader.load('4422.gltf',(gltf) => {
		// gltf.scene.scale.set(0.03, 0.03, 0.03)
		scene.add(gltf.scene)//.children[0]
		render();
		// scene.add(gltf.scene)
	},
	(progress) => {
		console.log(progress)
	},
	(error) => {
		console.log(error)
		},
	)

		//
		editor.storage.init( function () {
			editor.storage.get( function ( state ) {
				if ( isLoadingFromHash ) return;
				if ( state !== undefined ) {
					editor.fromJSON( state );
				}
				const selected = editor.config.getKey( 'selected' );
				if ( selected !== undefined ) {
					editor.selectByUuid( selected );
				}
			} );
			//
			let timeout;
			function saveState() {
				if ( editor.config.getKey( 'autosave' ) === false ) {
					return;
				}
				clearTimeout( timeout );
				timeout = setTimeout( function () {
					editor.signals.savingStarted.dispatch();
					timeout = setTimeout( function () {
						editor.storage.set( editor.toJSON() );
						editor.signals.savingFinished.dispatch();
					}, 100 );
				}, 1000 );
			}
			const signals = editor.signals;
			signals.geometryChanged.add( saveState );
			signals.objectAdded.add( saveState );
			signals.objectChanged.add( saveState );
			signals.objectRemoved.add( saveState );
			signals.materialChanged.add( saveState );
			signals.sceneBackgroundChanged.add( saveState );
			signals.sceneEnvironmentChanged.add( saveState );
			signals.sceneFogChanged.add( saveState );
			signals.sceneGraphChanged.add( saveState );
			signals.scriptChanged.add( saveState );
			signals.historyChanged.add( saveState );
		} );
		//
		document.addEventListener( 'dragover', function ( event ) {
			event.preventDefault();
			event.dataTransfer.dropEffect = 'copy';
		} );
		document.addEventListener( 'drop', function ( event ) {
			event.preventDefault();
			if ( event.dataTransfer.types[ 0 ] === 'text/plain' ) return; // Outliner drop
			if ( event.dataTransfer.items ) {
				// DataTransferItemList supports folders
				editor.loader.loadItemList( event.dataTransfer.items );
			} else {
				editor.loader.loadFiles( event.dataTransfer.files );
			}
		} );
		function onWindowResize() {
			editor.signals.windowResize.dispatch();
		}
		window.addEventListener( 'resize', onWindowResize );
		onWindowResize();
		//
		let isLoadingFromHash = false;
		const hash = window.location.hash;
		if ( hash.slice( 1, 6 ) === 'file=' ) {
			const file = hash.slice( 6 );
			if ( confirm( 'Any unsaved data will be lost. Are you sure?' ) ) {
				const loader = new THREE.FileLoader();
				loader.crossOrigin = '';
				loader.load( file, function ( text ) {
					editor.clear();
					editor.fromJSON( JSON.parse( text ) );
				} );
				isLoadingFromHash = true;
			}
		}
		// ServiceWorker
		if ( 'serviceWorker' in navigator ) {
			try {
				navigator.serviceWorker.register( '/static/threejs/sw.js' );
			} catch ( error ) {
			}
		}

	// Canvas
	// const canvas = document.querySelector('#mainCanvas') as HTMLCanvasElement
	
	// Scene
	// const scene = new THREE.Scene()
	// 这里重要，必须从editor里来
	// const camera = editor.camera;
	// const scene = editor.scene;
	// // Gui
	// // const gui = new dat.GUI()
	
	// // Size
	// const sizes = {
	//   width: window.innerWidth,
	//   height: window.innerHeight,
	// }

	// // Camera
	// const camera = new THREE.PerspectiveCamera(75, sizes.width / sizes.height, 0.1, 100)
	// camera.position.set(4, 4, 12)
	
	// Controls
	// const controls = new OrbitControls(camera, canvas)
	// controls.enableDamping = true
	// controls.zoomSpeed = 0.3
	// controls.target = new THREE.Vector3(0, 3, 0)


	/**
	 * Objects
	 */
	// plane
	// const plane = new THREE.Mesh(
	//   new THREE.PlaneGeometry(15, 15),
	//   new THREE.MeshStandardMaterial({
	//     color: '#607D8B',
	//   })
	// )
	// plane.rotateX(-Math.PI / 2)
	// plane.receiveShadow = true
	// scene.add(plane)
	
	/**
	 * Light
	 */
	// const directionLight = new THREE.DirectionalLight()
	// directionLight.castShadow = true
	// directionLight.position.set(5, 5, 6)
	// directionLight.shadow.camera.near = 1
	// directionLight.shadow.camera.far = 20
	// directionLight.shadow.camera.top = 10
	// directionLight.shadow.camera.right = 10
	// directionLight.shadow.camera.bottom = -10
	// directionLight.shadow.camera.left = -10

	// const directionLightHelper = new THREE.DirectionalLightHelper(directionLight, 2)
	// directionLightHelper.visible = false
	// scene.add(directionLightHelper)
	
	// const directionalLightCameraHelper = new THREE.CameraHelper(directionLight.shadow.camera)
	// directionalLightCameraHelper.visible = false
	// scene.add(directionalLightCameraHelper)
	
	// const ambientLight = new THREE.AmbientLight(new THREE.Color('#ffffff'), 0.3)
	// scene.add(ambientLight, directionLight)
	
	
	// scene.add(new THREE.AmbientLight(0x666666));//环境光


	// // Renderer
	// const renderer = new THREE.WebGLRenderer({
	//   // canvas,
	//   antialias: true,
	// })
	// renderer.setSize(sizes.width, sizes.height)
	// renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2))
	// renderer.shadowMap.enabled = true

	// Animations
	// const tick = () => {
	//   stats.begin()
	//   controls.update()
	
	//   // Render
	//   renderer.render(scene, camera)
	//   stats.end()
	//   requestAnimationFrame(tick)
	// }
	
	// tick()

	// let camera, scene, renderer;
	// init();
	// render();
	// function init() {
	// 	const container = document.createElement( 'div' );
	// 	document.body.appendChild( container );
	// 	camera = new THREE.PerspectiveCamera( 45, window.innerWidth / window.innerHeight, 0.25, 20 );
	// 	camera.position.set( - 1.8, 0.6, 2.7 );
	// 	scene = new THREE.Scene();
	// 	new RGBELoader()
	// 		.setPath( '/static/threejs/' )
	// 		.load( 'royal_esplanade_1k.hdr', function ( texture ) {
	// 			texture.mapping = THREE.EquirectangularReflectionMapping;
	// 			scene.background = texture;
	// 			scene.environment = texture;
	// 			render();
	// 			// model
	// 			const loader = new GLTFLoader().setPath( '/static/threejs/' );
	// 			loader.load( 'bazi.gltf', function ( gltf ) {
	// 				scene.add( gltf.scene );
	// 				render();
	// 			} );
	// 		} );
	// 	renderer = new THREE.WebGLRenderer( { antialias: true } );
	// 	renderer.setPixelRatio( window.devicePixelRatio );
	// 	renderer.setSize( window.innerWidth, window.innerHeight );
	// 	renderer.toneMapping = THREE.ACESFilmicToneMapping;
	// 	renderer.toneMappingExposure = 1;
	// 	renderer.outputEncoding = THREE.sRGBEncoding;
	// 	container.appendChild( renderer.domElement );
	// 	const controls = new OrbitControls( camera, renderer.domElement );
	// 	controls.addEventListener( 'change', render ); // use if there is no animation loop
	// 	controls.minDistance = 2;
	// 	controls.maxDistance = 10;
	// 	controls.target.set( 0, 0, - 0.2 );
	// 	controls.update();
	// 	window.addEventListener( 'resize', onWindowResize );
	// }

	// const gltfLoader = new GLTFLoader().setPath( '/static/threejs/' );
	// gltfLoader.load('scene.gltf',(gltf) => {
	// 	// console.log('success')
	// 	// console.log(gltf)
	// 	// gltf.scene.scale.set(0.03, 0.03, 0.03)
	// 	scene.add(gltf.scene.children[0])
	// 	// console.log(gltf.scene)
	// 	// gltf.scene.traverse( function ( child ) {
	// 	//   if ( child.isMesh ) {
	// 	//     child.material.emissive =  child.material.color;
	// 	//     child.material.emissiveMap = child.material.map ;
	// 	//   }
	// 	// });
		
	// 	render();
	// 	// scene.add(gltf.scene)
	// 	// console.log(scene)
	// },
	// (progress) => {
	// 	// console.log('progress')
	// 	console.log(progress)
	// },
	// (error) => {
	// 	// console.log('error')
	// 	console.log(error)
	// 	},
	// )
	// gltfLoader.load('/static/threejs/scene.gltf', gltf => {
	// scene.add(gltf.scene)
	// })

	// function onWindowResize() {
	// 	camera.aspect = window.innerWidth / window.innerHeight;
	// 	camera.updateProjectionMatrix();
	// 	renderer.setSize( window.innerWidth, window.innerHeight );
	// 	render();
	// }
	//
	function render() {
		renderer.render( scene, camera );
	}

	</script>
</body>

</html>
<!-- 关闭网格：GridHelper——D:\gowork\src\github.com\3xxx\engineercms\static\threejs\js\Viewport.js: -->
<!-- 加载gltf文件例子——D:\gowork\src\github.com\3xxx\engineercms\static\threejs\examples\webgl_loader_gltf.html -->
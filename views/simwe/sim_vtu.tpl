<!DOCTYPE html>
<html lang="en">
	<head>
		<title>three.js webgl - loaders - vtk loader</title>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, user-scalable=no, minimum-scale=1.0, maximum-scale=1.0">
		<link type="text/css" rel="stylesheet" href="/static/sim/main.css">
	</head>

	<body>
		<div id="info">
		<a href="https://threejs.org" target="_blank" rel="noopener">three.js</a> -
		vtk formats loader test<br />
		Legacy vtk model from <a href="http://www.cc.gatech.edu/projects/large_models/" target="_blank" rel="noopener">The GeorgiaTech Lagre Geometric Model Archive</a>
		</div>

		<script type="module">

			import * as THREE from '/static/sim/js/three.module.js';

			import Stats from '/static/sim/jsm/libs/stats.module.js';

			import { TrackballControls } from '/static/sim/jsm/controls/TrackballControls.js';
			import { VTKLoader } from '/static/sim/jsm/loaders/VTKLoader.js';

			let container, stats;

			let camera, controls, scene, renderer;

			init();
			animate();

			function init() {

				camera = new THREE.PerspectiveCamera( 60, window.innerWidth / window.innerHeight, 0.01, 1e10 );
				camera.position.z = 0.2;

				scene = new THREE.Scene();

				scene.add( camera );

				// light

				const hemiLight = new THREE.HemisphereLight( 0xffffff, 0x000000, 1 );
				scene.add( hemiLight );

				const dirLight = new THREE.DirectionalLight( 0xffffff, 0.5 );
				dirLight.position.set( 2, 2, 2 );
				scene.add( dirLight );

				const loader = new VTKLoader();
				loader.load( "/static/sim/models/vtk/bunny.vtk", function ( geometry ) {
					// loader.load( "/static/sim/models/vtk/asc.vtu", function ( geometry ) {
					//ansys输出的rst经过python读取二进制文件，输出非格式化网格数据后用paraview转成格式化网格数据，可以显示
					geometry.center();
					geometry.computeVertexNormals();

					const material = new THREE.MeshLambertMaterial( { color: 0xffffff } );
					const mesh = new THREE.Mesh( geometry, material );
					mesh.position.set( - 0.075, 0.005, 0 );
					mesh.scale.multiplyScalar( 0.0002 );// 调整显示比例
					scene.add( mesh );

				} );

				const loader1 = new VTKLoader();
				loader1.load( '/static/sim/models/vtk/cube_ascii.vtp', function ( geometry ) {

					geometry.computeVertexNormals();
					geometry.center();

					const material = new THREE.MeshLambertMaterial( { color: 0x00ff00 } );
					const mesh = new THREE.Mesh( geometry, material );

					mesh.position.set( - 0.025, 0, 0 );
					mesh.scale.multiplyScalar( 0.01 );


					scene.add( mesh );

				} );

				const loader2 = new VTKLoader();
				loader2.load( '/static/sim/models/vtk/cube_binary.vtp', function ( geometry ) {

					geometry.computeVertexNormals();
					geometry.center();

					const material = new THREE.MeshLambertMaterial( { color: 0x0000ff } );
					const mesh = new THREE.Mesh( geometry, material );

					mesh.position.set( 0.025, 0, 0 );
					mesh.scale.multiplyScalar( 0.01 );


					scene.add( mesh );

				} );

				const loader3 = new VTKLoader();
				loader3.load( '/static/sim/models/vtk/cube_no_compression.vtp', function ( geometry ) {

					geometry.computeVertexNormals();
					geometry.center();

					const material = new THREE.MeshLambertMaterial( { color: 0xff0000 } );
					const mesh = new THREE.Mesh( geometry, material );

					mesh.position.set( 0.075, 0, 0 );
					mesh.scale.multiplyScalar( 0.01 );


					scene.add( mesh );

				} );

				// renderer

				renderer = new THREE.WebGLRenderer();
				renderer.setPixelRatio( window.devicePixelRatio );
				renderer.setSize( window.innerWidth, window.innerHeight );

				container = document.createElement( 'div' );
				document.body.appendChild( container );
				container.appendChild( renderer.domElement );

				// controls

				controls = new TrackballControls( camera, renderer.domElement );
				controls.minDistance = .1;
				controls.maxDistance = 0.5;
				controls.rotateSpeed = 5.0;

				stats = new Stats();
				container.appendChild( stats.dom );

				//

				window.addEventListener( 'resize', onWindowResize );

			}

			function onWindowResize() {

				camera.aspect = window.innerWidth / window.innerHeight;
				camera.updateProjectionMatrix();

				renderer.setSize( window.innerWidth, window.innerHeight );

				controls.handleResize();

			}

			function animate() {

				requestAnimationFrame( animate );

				controls.update();

				renderer.render( scene, camera );

				stats.update();

			}

		</script>

	</body>
</html>

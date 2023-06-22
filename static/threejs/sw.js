// r138

const cacheName = 'threejs-editor';

const assets = [
	'/static/threejs/',

	'/static/threejs/manifest.json',
	'/static/threejs/images/icon.png',

	'/static/threejs/files/favicon.ico',

	'/static/threejs/build/three.module.js',

	'/static/examples/jsm/controls/TransformControls.js',

	'/static/examples/jsm/libs/chevrotain.module.min.js',
	'/static/examples/jsm/libs/fflate.module.js',

	'/static/examples/js/libs/draco/draco_decoder.js',
	'/static/examples/js/libs/draco/draco_decoder.wasm',
	'/static/examples/js/libs/draco/draco_encoder.js',
	'/static/examples/js/libs/draco/draco_wasm_wrapper.js',

	'/static/examples/js/libs/draco/gltf/draco_decoder.js',
	'/static/examples/js/libs/draco/gltf/draco_decoder.wasm',
	'/static/examples/js/libs/draco/gltf/draco_wasm_wrapper.js',

	'/static/examples/jsm/libs/motion-controllers.module.js',

	'/static/examples/jsm/libs/rhino3dm/rhino3dm.wasm',
	'/static/examples/jsm/libs/rhino3dm/rhino3dm.js',

	'/static/examples/jsm/loaders/3DMLoader.js',
	'/static/examples/jsm/loaders/3MFLoader.js',
	'/static/examples/jsm/loaders/AMFLoader.js',
	'/static/examples/jsm/loaders/ColladaLoader.js',
	'/static/examples/jsm/loaders/DRACOLoader.js',
	'/static/examples/jsm/loaders/FBXLoader.js',
	'/static/examples/jsm/loaders/GLTFLoader.js',
	'/static/examples/jsm/loaders/KMZLoader.js',
	'/static/examples/jsm/loaders/IFCLoader.js',
	'/static/examples/jsm/loaders/ifc/web-ifc-api.js',
	'/static/examples/jsm/loaders/ifc/web-ifc.wasm',
	'/static/examples/jsm/loaders/MD2Loader.js',
	'/static/examples/jsm/loaders/OBJLoader.js',
	'/static/examples/jsm/loaders/MTLLoader.js',
	'/static/examples/jsm/loaders/PLYLoader.js',
	'/static/examples/jsm/loaders/RGBELoader.js',
	'/static/examples/jsm/loaders/STLLoader.js',
	'/static/examples/jsm/loaders/SVGLoader.js',
	'/static/examples/jsm/loaders/TGALoader.js',
	'/static/examples/jsm/loaders/TDSLoader.js',
	'/static/examples/jsm/loaders/VOXLoader.js',
	'/static/examples/jsm/loaders/VRMLLoader.js',
	'/static/examples/jsm/loaders/VTKLoader.js',
	'/static/examples/jsm/loaders/XYZLoader.js',

	'/static/examples/jsm/curves/NURBSCurve.js',
	'/static/examples/jsm/curves/NURBSUtils.js',

	'/static/examples/jsm/interactive/HTMLMesh.js',
	'/static/examples/jsm/interactive/InteractiveGroup.js',

	'/static/examples/jsm/environments/RoomEnvironment.js',

	'/static/examples/jsm/exporters/ColladaExporter.js',
	'/static/examples/jsm/exporters/DRACOExporter.js',
	'/static/examples/jsm/exporters/GLTFExporter.js',
	'/static/examples/jsm/exporters/OBJExporter.js',
	'/static/examples/jsm/exporters/PLYExporter.js',
	'/static/examples/jsm/exporters/STLExporter.js',
	'/static/examples/jsm/exporters/USDZExporter.js',

	'/static/examples/jsm/helpers/VertexNormalsHelper.js',

	'/static/examples/jsm/geometries/TeapotGeometry.js',

	'/static/examples/jsm/webxr/VRButton.js',
	'/static/examples/jsm/webxr/XRControllerModelFactory.js',

	'/static/threejs/images/rotate.svg',
	'/static/threejs/images/scale.svg',
	'/static/threejs/images/translate.svg',

	'/static/threejs/js/libs/codemirror/codemirror.css',
	'/static/threejs/js/libs/codemirror/theme/monokai.css',

	'/static/threejs/js/libs/codemirror/codemirror.js',
	'/static/threejs/js/libs/codemirror/mode/javascript.js',
	'/static/threejs/js/libs/codemirror/mode/glsl.js',

	'/static/threejs/js/libs/esprima.js',
	'/static/threejs/js/libs/jsonlint.js',

	'/static/threejs/js/libs/codemirror/addon/dialog.css',
	'/static/threejs/js/libs/codemirror/addon/show-hint.css',
	'/static/threejs/js/libs/codemirror/addon/tern.css',

	'/static/threejs/js/libs/codemirror/addon/dialog.js',
	'/static/threejs/js/libs/codemirror/addon/show-hint.js',
	'/static/threejs/js/libs/codemirror/addon/tern.js',
	'/static/threejs/js/libs/acorn/acorn.js',
	'/static/threejs/js/libs/acorn/acorn_loose.js',
	'/static/threejs/js/libs/acorn/walk.js',
	'/static/threejs/js/libs/ternjs/polyfill.js',
	'/static/threejs/js/libs/ternjs/signal.js',
	'/static/threejs/js/libs/ternjs/tern.js',
	'/static/threejs/js/libs/ternjs/def.js',
	'/static/threejs/js/libs/ternjs/comment.js',
	'/static/threejs/js/libs/ternjs/infer.js',
	'/static/threejs/js/libs/ternjs/doc_comment.js',
	'/static/threejs/js/libs/tern-threejs/threejs.js',

	'/static/threejs/js/libs/signals.min.js',
	'/static/threejs/js/libs/ui.js',
	'/static/threejs/js/libs/ui.three.js',

	'/static/threejs/js/libs/app.js',
	'/static/threejs/js/Player.js',
	'/static/threejs/js/Script.js',

	//

	'/static/threejs/css/main.css',

	'/static/threejs/js/EditorControls.js',
	'/static/threejs/js/Storage.js',

	'/static/threejs/js/Editor.js',
	'/static/threejs/js/Config.js',
	'/static/threejs/js/History.js',
	'/static/threejs/js/Loader.js',
	'/static/threejs/js/LoaderUtils.js',
	'/static/threejs/js/Menubar.js',
	'/static/threejs/js/Menubar.File.js',
	'/static/threejs/js/Menubar.Edit.js',
	'/static/threejs/js/Menubar.Add.js',
	'/static/threejs/js/Menubar.Play.js',
	'/static/threejs/js/Menubar.Examples.js',
	'/static/threejs/js/Menubar.Help.js',
	'/static/threejs/js/Menubar.View.js',
	'/static/threejs/js/Menubar.Status.js',
	'/static/threejs/js/Resizer.js',
	'/static/threejs/js/Sidebar.js',
	'/static/threejs/js/Sidebar.Scene.js',
	'/static/threejs/js/Sidebar.Project.js',
	'/static/threejs/js/Sidebar.Project.Materials.js',
	'/static/threejs/js/Sidebar.Project.Renderer.js',
	'/static/threejs/js/Sidebar.Project.Video.js',
	'/static/threejs/js/Sidebar.Settings.js',
	'/static/threejs/js/Sidebar.Settings.History.js',
	'/static/threejs/js/Sidebar.Settings.Shortcuts.js',
	'/static/threejs/js/Sidebar.Settings.Viewport.js',
	'/static/threejs/js/Sidebar.Properties.js',
	'/static/threejs/js/Sidebar.Object.js',
	'/static/threejs/js/Sidebar.Geometry.js',
	'/static/threejs/js/Sidebar.Geometry.BufferGeometry.js',
	'/static/threejs/js/Sidebar.Geometry.Modifiers.js',
	'/static/threejs/js/Sidebar.Geometry.BoxGeometry.js',
	'/static/threejs/js/Sidebar.Geometry.CircleGeometry.js',
	'/static/threejs/js/Sidebar.Geometry.CylinderGeometry.js',
	'/static/threejs/js/Sidebar.Geometry.DodecahedronGeometry.js',
	'/static/threejs/js/Sidebar.Geometry.ExtrudeGeometry.js',
	'/static/threejs/js/Sidebar.Geometry.IcosahedronGeometry.js',
	'/static/threejs/js/Sidebar.Geometry.LatheGeometry.js',
	'/static/threejs/js/Sidebar.Geometry.OctahedronGeometry.js',
	'/static/threejs/js/Sidebar.Geometry.PlaneGeometry.js',
	'/static/threejs/js/Sidebar.Geometry.RingGeometry.js',
	'/static/threejs/js/Sidebar.Geometry.SphereGeometry.js',
	'/static/threejs/js/Sidebar.Geometry.ShapeGeometry.js',
	'/static/threejs/js/Sidebar.Geometry.TetrahedronGeometry.js',
	'/static/threejs/js/Sidebar.Geometry.TorusGeometry.js',
	'/static/threejs/js/Sidebar.Geometry.TorusKnotGeometry.js',
	'/static/threejs/js/Sidebar.Geometry.TubeGeometry.js',
	'/static/threejs/js/Sidebar.Geometry.TeapotGeometry.js',
	'/static/threejs/js/Sidebar.Material.js',
	'/static/threejs/js/Sidebar.Material.BooleanProperty.js',
	'/static/threejs/js/Sidebar.Material.ColorProperty.js',
	'/static/threejs/js/Sidebar.Material.ConstantProperty.js',
	'/static/threejs/js/Sidebar.Material.MapProperty.js',
	'/static/threejs/js/Sidebar.Material.NumberProperty.js',
	'/static/threejs/js/Sidebar.Material.Program.js',
	'/static/threejs/js/Sidebar.Animation.js',
	'/static/threejs/js/Sidebar.Script.js',
	'/static/threejs/js/Strings.js',
	'/static/threejs/js/Toolbar.js',
	'/static/threejs/js/Viewport.js',
	'/static/threejs/js/Viewport.Camera.js',
	'/static/threejs/js/Viewport.Info.js',
	'/static/threejs/js/Viewport.ViewHelper.js',
	'/static/threejs/js/Viewport.VR.js',

	'/static/threejs/js/Command.js',
	'/static/threejs/js/commands/AddObjectCommand.js',
	'/static/threejs/js/commands/RemoveObjectCommand.js',
	'/static/threejs/js/commands/MoveObjectCommand.js',
	'/static/threejs/js/commands/SetPositionCommand.js',
	'/static/threejs/js/commands/SetRotationCommand.js',
	'/static/threejs/js/commands/SetScaleCommand.js',
	'/static/threejs/js/commands/SetValueCommand.js',
	'/static/threejs/js/commands/SetUuidCommand.js',
	'/static/threejs/js/commands/SetColorCommand.js',
	'/static/threejs/js/commands/SetGeometryCommand.js',
	'/static/threejs/js/commands/SetGeometryValueCommand.js',
	'/static/threejs/js/commands/MultiCmdsCommand.js',
	'/static/threejs/js/commands/AddScriptCommand.js',
	'/static/threejs/js/commands/RemoveScriptCommand.js',
	'/static/threejs/js/commands/SetScriptValueCommand.js',
	'/static/threejs/js/commands/SetMaterialCommand.js',
	'/static/threejs/js/commands/SetMaterialColorCommand.js',
	'/static/threejs/js/commands/SetMaterialMapCommand.js',
	'/static/threejs/js/commands/SetMaterialValueCommand.js',
	'/static/threejs/js/commands/SetMaterialVectorCommand.js',
	'/static/threejs/js/commands/SetSceneCommand.js',
	'/static/threejs/js/commands/Commands.js',

	//

	'/static/examples/arkanoid.app.json',
	'/static/examples/camera.app.json',
	'/static/examples/particles.app.json',
	'/static/examples/pong.app.json',
	'/static/examples/shaders.app.json'

];

self.addEventListener( 'install', async function () {

	const cache = await caches.open( cacheName );

	assets.forEach( function ( asset ) {

		cache.add( asset ).catch( function () {

			console.warn( '[SW] Cound\'t cache:', asset );

		} );

	} );

} );

self.addEventListener( 'fetch', async function ( event ) {

	const request = event.request;
	event.respondWith( networkFirst( request ) );

} );

async function networkFirst( request ) {

	return fetch( request )
		.then( async function ( response ) {

			const cache = await caches.open( cacheName );

			cache.put( request, response.clone() );

			return response;

		} )
		.catch( async function () {

			const cachedResponse = await caches.match( request );

			if ( cachedResponse === undefined ) {

				console.warn( '[SW] Not cached:', request.url );

			}

			return cachedResponse;

		} );

}

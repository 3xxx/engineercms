<!DOCTYPE html>
<html lang="zh">
<head>
    <meta charset="UTF-8">
    <title>GLTF模型加载案例</title>
    <style type="text/css">
        html, body {
            margin: 0;
            height: 100%;
        }

        canvas {
            display: block;
        }

    </style>
</head>
<body onload="draw();">

</body>
<script src="/static/sim/js/three.js"></script>
<script src="/static/sim/js/GLTFLoader.js"></script>
<script src="/static/sim/js/OrbitControls.js"></script>
<script src="/static/sim/js/stats.min.js"></script>
<script src="/static/sim/js/dat.gui.min.js"></script>
<script>
    var renderer, camera, scene, gui, stats, ambientLight, directionalLight, control;

    function initRender() {
        renderer = new THREE.WebGLRenderer({antialias: true});
        renderer.setSize(window.innerWidth, window.innerHeight);
        //告诉渲染器需要阴影效果
        renderer.shadowMap.enabled = true;
        renderer.shadowMap.type = THREE.PCFSoftShadowMap; // 默认的是，没有设置的这个清晰 THREE.PCFShadowMap
        document.body.appendChild(renderer.domElement);
    }

    function initCamera() {
        camera = new THREE.PerspectiveCamera(45, window.innerWidth / window.innerHeight, 0.1, 1000);
        camera.position.set(0, 100, 200);
        camera.lookAt(new THREE.Vector3(0, 0, 0));
    }

    function initScene() {
        scene = new THREE.Scene();
    }

    function initGui() {
        //声明一个保存需求修改的相关数据的对象
        gui = {};

        var datGui = new dat.GUI();
        //将设置属性添加到gui当中，gui.add(对象，属性，最小值，最大值）
    }

    function initLight() {
        ambientLight = new THREE.AmbientLight("#111111");
        scene.add(ambientLight);

        directionalLight = new THREE.DirectionalLight("#ffffff");
        directionalLight.position.set(40, 60, 10);

        directionalLight.shadow.camera.near = 1; //产生阴影的最近距离
        directionalLight.shadow.camera.far = 400; //产生阴影的最远距离
        directionalLight.shadow.camera.left = -50; //产生阴影距离位置的最左边位置
        directionalLight.shadow.camera.right = 50; //最右边
        directionalLight.shadow.camera.top = 50; //最上边
        directionalLight.shadow.camera.bottom = -50; //最下面

        //这两个值决定生成阴影密度 默认512
        directionalLight.shadow.mapSize.height = 1024;
        directionalLight.shadow.mapSize.width = 1024;

        //告诉平行光需要开启阴影投射
        directionalLight.castShadow = true;

        scene.add(directionalLight);
    }

    function initModel() {

        //底部平面
        var planeGeometry = new THREE.PlaneGeometry(100, 100);
        var planeMaterial = new THREE.MeshLambertMaterial({color: 0xaaaaaa, side: THREE.DoubleSide});
        var plane = new THREE.Mesh(planeGeometry, planeMaterial);
        plane.rotation.x = -0.5 * Math.PI;
        plane.position.y = -.1;
        plane.receiveShadow = true; //可以接收阴影
        scene.add(plane);

        //创建gltf加载器
        var loader = new THREE.GLTFLoader();
        loader.load('/static/sim/js/scene.gltf', function (gltf) {
            gltf.scene.scale.set(.1,.1,.1);
            scene.add(gltf.scene);
        });

    }

    function initStats() {
        stats = new Stats();
        document.body.appendChild(stats.dom);
    }

    function initControl() {
        control = new THREE.OrbitControls(camera, renderer.domElement);
    }

    function render() {

        control.update();

        renderer.render(scene, camera);
    }

    function onWindowResize() {

        camera.aspect = window.innerWidth / window.innerHeight;
        camera.updateProjectionMatrix();
        renderer.setSize(window.innerWidth, window.innerHeight);

    }

    function animate() {
        //更新控制器
        render();

        //更新性能插件
        stats.update();

        requestAnimationFrame(animate);
    }

    function draw() {
        initGui();
        initRender();
        initScene();
        initCamera();
        initLight();
        initModel();
        initStats();

        initControl();

        animate();
        window.onresize = onWindowResize;
    }
</script>
</html>
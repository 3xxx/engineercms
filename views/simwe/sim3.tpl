<!-- <!DOCTYPE html>
<html>

<head>
  <meta charset=utf-8>
  <title>我的第一个Three.js案例</title>
  <style>
  body {
    margin: 0;
  }

  canvas {
    width: 100%;
    height: 100%;
    display: block;
  }
  </style>
</head>

<body onload="init()">
  <script src="https://cdn.bootcss.com/three.js/92/three.js"></script>
  <script>

  var renderer, camera, scene, geometry, material, mesh;//声明一些全局变量


  function initRenderer() {//初始化渲染器
    renderer = new THREE.WebGLRenderer(); //实例化渲染器
    renderer.setSize(window.innerWidth, window.innerHeight); //设置宽和高
    document.body.appendChild(renderer.domElement); //添加到dom
  }


  function initScene() {//初始化场景
    scene = new THREE.Scene(); //实例化场景
  }


  function initCamera() {//初始化相机
    camera = new THREE.PerspectiveCamera(45, window.innerWidth / window.innerHeight, 0.1, 200); //实例化相机
    camera.position.set(0, 0, 15);
  }


  function initMesh() {//创建模型
    geometry = new THREE.BoxGeometry(2, 2, 2); //创建几何体
    material = new THREE.MeshNormalMaterial(); //创建材质

    mesh = new THREE.Mesh(geometry, material); //创建网格
    scene.add(mesh); //将网格添加到场景
  }


  function animate() {//运行动画
    requestAnimationFrame(animate); //循环调用函数

    mesh.rotation.x += 0.01; //每帧网格模型的沿x轴旋转0.01弧度
    mesh.rotation.y += 0.02; //每帧网格模型的沿y轴旋转0.02弧度

    renderer.render(scene, camera); //渲染界面
  }


  function init() {//初始化函数，页面加载完成是调用
    initRenderer();
    initScene();
    initCamera();
    initMesh();

    animate();
  }
  </script>
</body>

</html> -->
<!DOCTYPE html>
<html lang="zh">

<head>
  <meta charset="UTF-8">
  <title>JSON模型加载案例</title>
  <style type="text/css">
  html,
  body {
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
<script src="/static/sim/jsm/controls/OrbitControls.js"></script>
<script src="/static/sim/js/libs/stats.min.js"></script>

<script src="/static/sim/js/dat.gui.min.js"></script>

<!-- <script src="/static/sim/js/GLTFLoader.js"></script> -->
<script src="/static/sim/jsm/loaders/VTKLoader.js"></script>
<script>
var renderer, camera, scene, gui, stats, ambientLight, directionalLight, control;

function initRender() {
  renderer = new THREE.WebGLRenderer({ antialias: true });
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
  gui = {
    createScene: function() {
      //首先先删除掉当前场景所含有的立方体
      deleteGroup("group");
      //创建一个新的模型组
      let group = new THREE.Group();
      group.name = "group";
      let geometry = new THREE.BoxGeometry(10, 10, 10);
      for (let i = 0; i < 30; i++) {
        let material = new THREE.MeshLambertMaterial({ color: randomColor() });
        let mesh = new THREE.Mesh(geometry, material);
        //随机位置
        mesh.position.set(THREE.Math.randFloatSpread(200), THREE.Math.randFloatSpread(200), THREE.Math.randFloatSpread(200));
        group.add(mesh);
      }
      scene.add(group);
    },
    exporterScene: function() {
      //首先将场景转成json对象
      let group = scene.getObjectByName("group");
      if (!group) return;
      let obj = group.toJSON();
      //将json对象转成json字符串并存储
      download("file.json", JSON.stringify(obj));
    },
    importerScene: function() {
      //创建一个input来获取json数据
      let input = document.createElement("input");
      input.type = "file";
      input.addEventListener("change", function() {
        let file = input.files[0];
        console.log(file)
        //判断是否是json格式的文件
        if (file.type.indexOf("json") >= 0) {
          //首先先删除掉当前场景所含有的立方体
          deleteGroup("group");

          //读取文件内的内容
          let reader = new FileReader();
          reader.readAsText(file);
          reader.onloadend = function() {
            console.log(this.result)
            //使用three.js的JSONLoader将模型导入到场景
            let loader = new THREE.ObjectLoader();
            let group = loader.parse(JSON.parse(this.result));
            scene.add(group);
          }


        } else {
          //底部平面
          var planeGeometry = new THREE.PlaneGeometry(100, 100);
          var planeMaterial = new THREE.MeshLambertMaterial({ color: 0xaaaaaa, side: THREE.DoubleSide });
          var plane = new THREE.Mesh(planeGeometry, planeMaterial);
          plane.rotation.x = -0.5 * Math.PI;
          plane.position.y = -.1;
          plane.receiveShadow = true; //可以接收阴影
          scene.add(plane);

          //创建gltf加载器
          // var loader = new THREE.GLTFLoader();
          // loader.load('/static/sim/js/scene.gltf', function(gltf) {
          //   gltf.scene.scale.set(.1, .1, .1);
          //   scene.add(gltf.scene);
          // });

          //辅助工具
          var helper = new THREE.AxesHelper(50);
          scene.add(helper);

          var loader = new THREE.VTKLoader();
          loader.load("/static/sim/models/vtk/hexrstpolydata222.vtk", function(geometry) {
            //模型的法向量有问题，更新一下法向量
            geometry.computeFaceNormals();
            geometry.computeVertexNormals();
            //创建纹理
            var mat = new THREE.MeshLambertMaterial({ color: 0xaaaaaa });
            //创建模型
            var group = new THREE.Mesh(geometry, mat);
            group.scale.set(10, 10, 10);
            scene.add(group);
          });

          // 原文链接：https://blog.csdn.net/qq_30100043/article/details/79633022

          // var file    = document.querySelector('input[type=file]').files[0];
          //   var reader  = new FileReader();

          //   reader.addEventListener("load", function () {
          //     preview.src = reader.result;
          //   }, false);

          //   if (file) {
          //     reader.readAsDataURL(file);
          //   }

          //读取文件内的内容
          // let reader = new FileReader();
          // // reader.readAsText(file);
          // // reader.readAsDataUrl(file);
          // reader.readAsDataURL(file);
          // reader.onloadend = function() {
          //    console.log(this.result)
          //   //   //使用three.js的JSONLoader将模型导入到场景
          //   //   // let loader = new THREE.ObjectLoader();
          //   let loader = new THREE.GLTFLoader();
          //   let gltf = loader.load(this.result);
          //   scene.add(gltf);
          // }

        }

      });
      input.click();
    },
    loaderScene: function() {
      //首先先删除掉当前场景所含有的立方体
      deleteGroup("group");

      //使用JSONLoader加载json格式文件
      let loader = new THREE.ObjectLoader();

      loader.load("/static/sim/js/file.json", function(group) {
        scene.add(group);
      });
    }
  };

  var datGui = new dat.GUI();
  //将设置属性添加到gui当中，gui.add(对象，属性，最小值，最大值）
  datGui.add(gui, "createScene").name("添加模型");
  datGui.add(gui, "exporterScene").name("导出模型");
  datGui.add(gui, "importerScene").name("导入模型");
  datGui.add(gui, "loaderScene").name("加载模型");

  gui.loaderScene();
}

//随机颜色
function randomColor() {
  var arrHex = ["0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"],
    strHex = "#",
    index;

  for (var i = 0; i < 6; i++) {
    index = Math.round(Math.random() * 15);
    strHex += arrHex[index];
  }

  return strHex;
}

//保存文件
function download(filename, text) {
  var pom = document.createElement('a');
  pom.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(text));
  pom.setAttribute('download', filename);
  if (document.createEvent) {
    var event = document.createEvent('MouseEvents');
    event.initEvent('click', true, true);
    pom.dispatchEvent(event);
  } else {
    pom.click();
  }
}

//删除group
function deleteGroup(name) {
  let group = scene.getObjectByName(name);
  if (!group) return;
  //删除掉所有的模型组内的mesh
  group.traverse(function(item) {
    if (item instanceof THREE.Mesh) {
      item.geometry.dispose(); //删除几何体
      item.material.dispose(); //删除材质
    }
  });

  scene.remove(group);
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
  initRender();
  initScene();
  initCamera();
  initLight();
  initStats();
  initGui();

  initControl();

  animate();
  window.onresize = onWindowResize;
}
</script>

</html>
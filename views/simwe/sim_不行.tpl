<!-- vtkPolyDataReader -->
<!DOCTYPE html>
<html lang="en" xmlns:th="http:www.thymeleaf.org">
<head>
    <meta charset="UTF-8">
    <link rel="stylesheet" href="https://unpkg.com/element-ui/lib/theme-chalk/index.css">
    <script type="text/javascript" src="https://unpkg.com/vtk.js"></script>
    <title>Vue Test</title>
</head>
<body>
    <div id="app">
        <el-container style="border: 1px solid #eee">
            <el-aside width="200px" style="background-color: rgb(238, 241, 246)">
                <el-menu :default-openeds="['1']" :default-active="$route.path" @select="handleSelect">
                    <el-menu-item index="">
                        <template slot="title"><i class="el-icon-message"></i>首页</template>
                    </el-menu-item>
                    <el-menu-item index="/menu-1-index">
                        <template slot="title"><i class="el-icon-menu"></i>菜单1</template>
                    </el-menu-item>
                    <el-menu-item index="/menu-2-index">
                        <template slot="title"><i class="el-icon-setting"></i>菜单2</template>
                    </el-menu-item>
                </el-menu>
            </el-aside>
            <el-container>
                <router-view></router-view>
            </el-container>
        </el-container>
    </div>
    <script>
      // var fullScreenRenderer = vtk.Rendering.Misc.vtkFullScreenRenderWindow.newInstance();
  // var actor = vtk.Rendering.Core.vtkActor.newInstance();
  // var mapper = vtk.Rendering.Core.vtkMapper.newInstance();
  // var cone = vtk.Filters.Sources.vtkConeSource.newInstance();

  // actor.setMapper(mapper);
  // mapper.setInputConnection(cone.getOutputPort());

  // var renderer = fullScreenRenderer.getRenderer();
  // renderer.addActor(actor);
  // renderer.resetCamera();

  // var renderWindow = fullScreenRenderer.getRenderWindow();
  // renderWindow.render();

  // Load the rendering pieces we want to use (for both WebGL and WebGPU)
  // import 'vtk.js/Sources/  Rendering/Profiles/Volume';

  // import vtkColorTransferFunction from 'vtk.js/Sources/   Rendering/Core/ColorTransferFunction';
  // import vtkFullScreenRenderWindow from 'vtk.js/Sources/   Rendering/Misc/FullScreenRenderWindow';
  // import vtkHttpDataSetReader from 'vtk.js/Sources/   IO/Core/HttpDataSetReader';
  // import vtkPiecewiseFunction from 'vtk.js/Sources/   Common/DataModel/PiecewiseFunction';
  // import vtkVolume from 'vtk.js/Sources/   Rendering/Core/Volume';
  // import vtkVolumeMapper from 'vtk.js/Sources/   Rendering/Core/VolumeMapper';

  // // Force the loading of HttpDataAccessHelper to support gzip decompression
  // import 'vtk.js/Sources/IO/Core/DataAccessHelper/HttpDataAccessHelper';


  // ----------------------------------------------------------------------------
  // Standard rendering code setup
  // ----------------------------------------------------------------------------

  // const fullScreenRenderer = vtkFullScreenRenderWindow.newInstance({
  //   background: [0, 0, 0],
  // });
  // const renderer = fullScreenRenderer.getRenderer();
  // const renderWindow = fullScreenRenderer.getRenderWindow();

  // fullScreenRenderer.addController(controlPanel);

  // ----------------------------------------------------------------------------
  // Example code
  // ----------------------------------------------------------------------------
  // Server is not sending the .gz and with the compress header
  // Need to fetch the true file name and uncompress it locally
  // ----------------------------------------------------------------------------

  // const reader = vtk.IO.Core.vtkHttpDataSetReader.newInstance({ fetchGzip: true });

  // const actor = vtkVolume.newInstance();
  // const mapper = vtkVolumeMapper.newInstance();
  // mapper.setSampleDistance(1.3);
  // actor.setMapper(mapper);

  // create color and opacity transfer functions

  // const ctfun = vtk.Rendering.Core.vtkColorTransferFunction.newInstance();
  // ctfun.addRGBPoint(0, 85 / 255.0, 0, 0);
  // ctfun.addRGBPoint(95, 1.0, 1.0, 1.0);
  // ctfun.addRGBPoint(225, 0.66, 0.66, 0.5);
  // ctfun.addRGBPoint(255, 0.3, 1.0, 0.5);
  // const ofun = vtk.Common.DataModel.vtkPiecewiseFunction.newInstance();
  // ofun.addPoint(0.0, 0.0);
  // ofun.addPoint(255.0, 1.0);

  // actor.getProperty().setRGBTransferFunction(0, ctfun);
  // actor.getProperty().setScalarOpacity(0, ofun);
  // actor.getProperty().setScalarOpacityUnitDistance(0, 3.0);
  // actor.getProperty().setInterpolationTypeToLinear();
  // actor.getProperty().setUseGradientOpacity(0, true);
  // actor.getProperty().setGradientOpacityMinimumValue(0, 2);
  // actor.getProperty().setGradientOpacityMinimumOpacity(0, 0.0);
  // actor.getProperty().setGradientOpacityMaximumValue(0, 20);
  // actor.getProperty().setGradientOpacityMaximumOpacity(0, 1.0);
  // actor.getProperty().setShade(true);
  // actor.getProperty().setAmbient(0.2);
  // actor.getProperty().setDiffuse(0.7);
  // actor.getProperty().setSpecular(0.3);
  // actor.getProperty().setSpecularPower(8.0);

  // mapper.setInputConnection(reader.getOutputPort());


  // reader.setUrl(`/static/sim/models/vtk/bunny.vtk`).then(() => {
  //   reader.loadData().then(() => {
  //     renderer.addVolume(actor);
  //     const interactor = renderWindow.getInteractor();
  //     interactor.setDesiredUpdateRate(15.0);
  //     renderer.resetCamera();
  //     renderer.getActiveCamera().zoom(1.5);
  //     renderer.getActiveCamera().elevation(70);
  //     renderer.resetCamera();
  //     renderWindow.render();
  //   });
  // });



// import 'vtk.js/Sources/favicon';

// Load the rendering pieces we want to use (for both WebGL and WebGPU)
// import 'vtk.js/Sources/Rendering/Profiles/Geometry';
// import 'vtk.js/Sources/Rendering/Profiles/Molecule'; // vtkSphereMapper + vtkStickMapper

// import vtkFullScreenRenderWindow from 'vtk.js/Sources/Rendering/Misc/FullScreenRenderWindow';
// import vtkActor from 'vtk.js/Sources/Rendering/Core/Actor';
// import vtkPDBReader from 'vtk.js/Sources/IO/Misc/PDBReader';
// import vtkSphereMapper from 'vtk.js/Sources/Rendering/Core/SphereMapper';
// import vtkStickMapper from 'vtk.js/Sources/Rendering/Core/StickMapper';
// import vtkMoleculeToRepresentation from 'vtk.js/Sources/Filters/General/MoleculeToRepresentation';

// ----------------------------------------------------------------------------
// Standard rendering code setup
// ----------------------------------------------------------------------------

const fullScreenRenderer = vtk.Rendering.Misc.vtkFullScreenRenderWindow.newInstance({
  background: [0, 0, 0],
});
const renderer = fullScreenRenderer.getRenderer();
const renderWindow = fullScreenRenderer.getRenderWindow();

// ----------------------------------------------------------------------------
// Example code
// ----------------------------------------------------------------------------
// import vtkPolyDataReader from 'vtk.js/Sources/IO/Legacy/PolyDataReader';
const reader = vtk.IO.Legacy.vtkPolyDataReader.newInstance();
// const reader = vtk.IO.Misc.vtkPDBReader.newInstance();
const filter = vtk.Filters.General.vtkMoleculeToRepresentation.newInstance();
const sphereMapper = vtk.Rendering.Core.vtkSphereMapper.newInstance();
const stickMapper = vtk.Rendering.Core.vtkStickMapper.newInstance();
const sphereActor = vtk.Rendering.Core.vtkActor.newInstance();
const stickActor = vtk.Rendering.Core.vtkActor.newInstance();

filter.setInputConnection(reader.getOutputPort());
filter.setHideElements(['H']);

// render sphere
sphereMapper.setInputConnection(filter.getOutputPort(0));
sphereMapper.setScaleArray(filter.getSphereScaleArrayName());
sphereActor.setMapper(sphereMapper);

// render sticks
stickMapper.setInputConnection(filter.getOutputPort(1));
stickMapper.setScaleArray('stickScales');
stickMapper.setOrientationArray('orientation');
stickActor.setMapper(stickMapper);

// reader.setUrl(`${__BASE_PATH__}/data/molecule/pdb/caffeine.pdb`).then(() => {
// reader.setUrl(`/static/sim/models/vtk/2LYZ.pdb`).then(() => {
  reader.setUrl(`/static/sim/models/vtk/hexrstpolydata222.vtk`).then(() => {
    // 不支持ansys的vtk文件转换为格式化网格的文件——连sphere也不现实，这个没有用处
  renderer.resetCamera();
  renderWindow.render();
});

renderer.addActor(sphereActor);
renderer.addActor(stickActor);
renderer.resetCamera();
renderWindow.render();

// -----------------------------------------------------------
// Make some variables global so that you can inspect and
// modify objects in your browser's developer console:
// -----------------------------------------------------------

global.reader = reader;
global.filter = filter;
global.sphereMapper = sphereMapper;
global.stickMapper = stickMapper;
global.sphereActor = sphereActor;
global.stickActor = stickActor;
global.renderer = renderer;
global.renderWindow = renderWindow;

    </script>
    <script src="https://unpkg.com/vue/dist/vue.js"></script>
    <script src="https://unpkg.com/vue-router/dist/vue-router.js "></script>
    <script src="https://unpkg.com/element-ui/lib/index.js"></script>
    <script src="https://unpkg.com/axios/dist/axios.min.js"></script>
    <script src="/static/sim/js/util.js"></script>
    <script src="/static/sim/js/route.js"></script>
    <script src="/static/sim/js/main.js"></script>

</body>
</html>
<!-- vtkHttpDataSetReader 只支持vtp文件夹-->
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
// import 'vtk.js/Sources/favicon';

// Load the rendering pieces we want to use (for both WebGL and WebGPU)
// import 'vtk.js/Sources/Rendering/Profiles/Geometry';

// import vtkFullScreenRenderWindow from 'vtk.js/Sources/Rendering/Misc/FullScreenRenderWindow';
// import vtkActor from 'vtk.js/Sources/Rendering/Core/Actor';
// import vtkCamera from 'vtk.js/Sources/Rendering/Core/Camera';
// import vtkMapper from 'vtk.js/Sources/Rendering/Core/Mapper';
// import vtkHttpDataSetReader from 'vtk.js/Sources/IO/Core/HttpDataSetReader';
// import vtkWindowedSincPolyDataFilter from 'vtk.js/Sources/Filters/General/WindowedSincPolyDataFilter';

// import controlPanel from './controller.html';

// Force DataAccessHelper to have access to various data source
// import 'vtk.js/Sources/IO/Core/DataAccessHelper/HttpDataAccessHelper';

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

const actor = vtk.Rendering.Core.vtkActor.newInstance();
renderer.addActor(actor);

const mapper = vtk.Rendering.Core.vtkMapper.newInstance({ interpolateScalarBeforeMapping: true });
actor.setMapper(mapper);

const cam = vtk.Rendering.Core.vtkCamera.newInstance();
renderer.setActiveCamera(cam);
cam.setFocalPoint(0, 0, 0);
cam.setPosition(0, 0, 10);
cam.setClippingRange(0.1, 50.0);

// Build pipeline
const reader = vtk.IO.Core.vtkHttpDataSetReader.newInstance({ fetchGzip: true });
reader.setUrl(`/static/sim/models/vtk/vtp.vtp`).then(() => {
// reader.setUrl(`/static/sim/models/vtk/cow.vtp`).then(() => {
// 只支持vtp文件夹
  reader.loadData().then(() => {
    renderer.resetCamera();
    renderWindow.render();
  });
});

const smoothFilter = vtk.Filters.General.vtkWindowedSincPolyDataFilter.newInstance({
  nonManifoldSmoothing: 0,
  numberOfIterations: 10,
});
smoothFilter.setInputConnection(reader.getOutputPort());
mapper.setInputConnection(smoothFilter.getOutputPort());

// ----------------------------------------------------------------------------
// UI control handling
// ----------------------------------------------------------------------------

fullScreenRenderer.addController(controlPanel);

// Warp setup
[
  'numberOfIterations',
  'passBand',
  'featureAngle',
  'edgeAngle',
  'nonManifoldSmoothing',
  'featureEdgeSmoothing',
  'boundarySmoothing',
].forEach((propertyName) => {
  document.querySelector(`.${propertyName}`).addEventListener('input', (e) => {
    let value;
    if (Number.isNaN(e.target.valueAsNumber)) {
      value = e.target.checked ? 1 : 0;
    } else {
      value = e.target.valueAsNumber;
    }
    if (propertyName === 'passBand') {
      // This formula maps:
      // 0.0  -> 1.0   (almost no smoothing)
      // 0.25 -> 0.1   (average smoothing)
      // 0.5  -> 0.01  (more smoothing)
      // 1.0  -> 0.001 (very strong smoothing)
      value = 10.0 ** (-4.0 * value);
    }
    smoothFilter.set({ [propertyName]: value });
    renderWindow.render();
    console.log({ [propertyName]: value });
  });
});

// -----------------------------------------------------------

renderer.resetCamera();
renderWindow.render();

// -----------------------------------------------------------
// Make some variables global so that you can inspect and
// modify objects in your browser's developer console:
// -----------------------------------------------------------

global.source = reader;
global.filter = smoothFilter;
global.mapper = mapper;
global.actor = actor;

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
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
  // import 'vtk.js/Sources/favicon';

  // Load the rendering pieces we want to use (for both WebGL and WebGPU)
  // import 'vtk.js/Sources/Rendering/Profiles/Geometry';

  // import vtkFullScreenRenderWindow from 'vtk.js/Sources/Rendering/Misc/FullScreenRenderWindow';
  // import vtkActor from 'vtk.js/Sources/Rendering/Core/Actor';
  // import vtkMapper from 'vtk.js/Sources/Rendering/Core/Mapper';

  // import vtkSTLWriter from 'vtk.js/Sources/IO/Geometry/STLWriter';
  // import vtkSTLReader from 'vtk.js/Sources/IO/Geometry/STLReader';
  // import vtkPolyDataReader from 'vtk.js/Sources/IO/Legacy/PolyDataReader';
  // ----------------------------------------------------------------------------
  // Standard rendering code setup
  // ----------------------------------------------------------------------------

  const fullScreenRenderer = vtk.Rendering.Misc.vtkFullScreenRenderWindow.newInstance();
  const renderer = fullScreenRenderer.getRenderer();
  const renderWindow = fullScreenRenderer.getRenderWindow();

  const reader = vtk.IO.Legacy.vtkPolyDataReader.newInstance();
  const writerReader = vtk.IO.Geometry.vtkSTLReader.newInstance();

  const writer = vtk.IO.Geometry.vtkSTLWriter.newInstance();

  reader
    .setUrl(`/static/sim/damrst.vtp`, { loadData: true })
    // 无法显示paraview转换的ansys的vtk非格式化网格文件
    // .setUrl(`/static/sim/models/vtk/bunny.vtk`, { loadData: true })
    .then(() => {
      writer.setInputData(reader.getOutputData());
      const fileContents = writer.getOutputData();
      // Can also use a static function to write to STL:
      // const fileContents = vtkSTLWriter.writeSTL(reader.getOutputData());

      // Display the resulting STL
      writerReader.parseAsArrayBuffer(fileContents.buffer);
      renderer.resetCamera();
      renderWindow.render();

      // Add a download link for it
      const blob = new Blob([fileContents], { type: 'application/octet-steam' });
      const a = window.document.createElement('a');
      a.href = window.URL.createObjectURL(blob, {
        type: 'application/octet-steam',
      });
      a.download = 'sphere.stl';
      a.text = 'Download';
      a.style.position = 'absolute';
      a.style.left = '50%';
      a.style.bottom = '10px';
      document.body.appendChild(a);
      a.style.background = 'white';
      a.style.padding = '5px';
    });

  const actor = vtk.Rendering.Core.vtkActor.newInstance();
  const mapper = vtk.Rendering.Core.vtkMapper.newInstance();
  actor.setMapper(mapper);

  mapper.setInputConnection(writerReader.getOutputPort());

  renderer.addActor(actor);

  global.writer = writer;
  global.writerReader = writerReader;
  global.mapper = mapper;
  global.actor = actor;
  global.renderer = renderer;
  global.renderWindow = renderWindow;

    </script>
    <script src="https://unpkg.com/vue@2.5.20/dist/vue.js"></script>
    <script src="https://unpkg.com/vue-router@3.5.0/dist/vue-router.js "></script>
    <script src="https://unpkg.com/element-ui/lib/index.js"></script>
    <script src="https://unpkg.com/axios/dist/axios.min.js"></script>
    <script src="/static/sim/js/util.js"></script>
    <script src="/static/sim/js/route.js"></script>
    <script src="/static/sim/js/main.js"></script>

</body>
</html>
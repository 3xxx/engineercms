<!DOCTYPE html>
<!-- saved from url=(0076)https://kitware.github.io/vtk-js/examples/GeometryViewer/GeometryViewer.html -->
<html foxified="">

<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
  <meta name="viewport" content="width=device-width, height=device-height, initial-scale=1, maximum-scale=1, minimum-scale=1, user-scalable=no">
  <link rel="icon" href="https://kitware.github.io/vtk-js/icon/favicon-16x16.png" sizes="16x16" type="image/png">
  <link rel="icon" href="https://kitware.github.io/vtk-js/icon/favicon-32x32.png" sizes="32x32" type="image/png">
  <link rel="icon" href="https://kitware.github.io/vtk-js/icon/favicon-96x96.png" sizes="96x96" type="image/png">
  <link rel="icon" href="https://kitware.github.io/vtk-js/icon/favicon-160x160.png" sizes="160x160" type="image/png">
  <link rel="icon" href="https://kitware.github.io/vtk-js/icon/favicon-196x196.png" sizes="196x196" type="image/png">
  <style type="text/css">
  .FPSMonitor-module_verticalContainer__1oES5 {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: stretch;
  }

  .FPSMonitor-module_horizontalContainer__3dO_q {
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
  }

  .FPSMonitor-module_leftPane__3PHsp {
    flex: none;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: stretch;
  }

  .FPSMonitor-module_rightPane__30Een {
    flex: 1;
    display: grid;
    grid-template-columns: auto auto;
    grid-auto-rows: 1.5em;
    grid-column-gap: 5px;
    grid-row-gap: 2px;
    padding: 10px;
  }

  .FPSMonitor-module_title__3a5vQ {
    flex: 1;
    font-weight: bold;
    padding: 5px 10px 0 10px;
  }

  .FPSMonitor-module_graph__lvtIQ {
    flex: none;
    border: solid 1px black;
    margin: 10px;
    border-radius: 2px;
    overflow: hidden;
  }

  .FPSMonitor-module_label__3saqc {
    font-weight: bold;
    text-transform: capitalize;
    text-align: right;
    align-self: center;
  }

  .FPSMonitor-module_value__2WrfF {
    font-style: italic;
    text-align: center;
    align-self: center;
  }
  </style>
  <style>.GeometryViewer-module-button_fORXA {
  position: absolute;
  right: 5px;
  top: 5px;
  width: 1em;
  z-index: 2;
  cursor: pointer;
}

.GeometryViewer-module-rootController_uBiGG {
  display: flex;
  flex-direction: column;
  position: absolute;
  top: 5px;
  left: 5px;
  right: 5px;
  z-index: 1;
}

.GeometryViewer-module-control_LaXi4 {
  display: flex;
  flex-direction: row;
  flex: 1;
  align-items: center;
}

.GeometryViewer-module-fullScreen_RWlQz {
  position: absolute;
  width: 100vw;
  height: 100vh;
  top: 0;
  left: 0;
  overflow: hidden;
  background: black;
  margin: 0;
  padding: 0;
  z-index: -1;
  display: flex;
  align-items: center;
  justify-content: center;
}

.GeometryViewer-module-fullParentSize_LzcO5 {
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  overflow: hidden;
}

.GeometryViewer-module-bigFileDrop_taNXL {
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  background-color: white;
  background-image: url(https://kitware.github.io/vtk-js/examples/GeometryViewer/3fdac36cbee8ffdf80fa.jpg);
  background-repeat: no-repeat;
  background-position: center;
  background-size: contain;
  border-radius: 10px;
  width: 50px;
  padding: calc(50vh - 2em) calc(50vw - 25px - 2em);
  cursor: pointer;
}

.GeometryViewer-module-selector_lZA7U {
  background: transparent;
  border: none;
  margin: 5px;
  z-index: 1;
  max-width: 100px;
  min-width: 100px;
}

label.GeometryViewer-module-selector_lZA7U {
  font-size: 12px;
  text-overflow: ellipsis;
  overflow: hidden;
}

select:focus {
  outline: none;
  border: none;
}

.GeometryViewer-module-progress_LIQPI {
  flex: none;
  font-size: 50px;
  color: white;
  z-index: 1;
  background: rgba(128,128,128,.5);
  padding: 20px;
  border-radius: 10px;
  -webkit-user-select: none;
     -moz-user-select: none;
          user-select: none;
}

.GeometryViewer-module-dark_paSVP {
  color: white;
}

.GeometryViewer-module-dark_paSVP option {
  color: black;
}

.GeometryViewer-module-light_kemHx {
  color: black;
}

.GeometryViewer-module-light_kemHx option {
  color: white;
}

.GeometryViewer-module-fpsMonitor_yHvkR {
  position: absolute;
  bottom: 10px;
  left: 10px;
  background-color: rgba(255, 255, 255, 0.5);
  border-radius: 5px;
  border: solid 1px gray;
}
</style>
</head>

<body style="margin: 0px; padding: 0px;">
  <div class="content GeometryViewer-module-fullScreen_RWlQz">
    <div>
      <div class="GeometryViewer-module-bigFileDrop_taNXL"><input type="file" multiple="" accept=".dat" style="display: none;"></div>
    </div>
  </div>
  <script defer="defer" src="/static/sim/GeometryViewer2.js"></script>
</body>

</html>
<!-- https://github.com/Kitware/vtk-js/blob/master/Examples/Applications/GeometryViewer/index.md -->
<!-- 页面显示出错，则到上面这个页面，下载保存，然后将页面中的script部分剪切出来，单独放到GeometryViewer2.js文件，再在这个页面中引用即可。 -->
<!-- 替换为accept=".dat" ，并且路径中带有fileURL就能自动加载这个dat文件，出处已经不记得来自哪里了-->
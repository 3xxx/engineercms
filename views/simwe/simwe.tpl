<!DOCTYPE html>

<head>
  <link rel="import" href="/static/sim/Sources/controller.html">
</head>
<html>

<body>
  <!-- <script src="/static/vue.js/vue.js"></script> -->
  <!-- <script type="text/javascript" src="https://unpkg.com/vtk.js"></script> 1.首先引入vtk.js -->
  <script type="text/javascript" src="/static/sim/vtk.js"></script>
  <script type="module">
    import '/static/sim/Sources/favicon.js';
// import {deepEqual} from '/static/sim/deep-equal/index.js';

import { AttributeTypes } from '/static/sim/Sources/Common/DataModel/DataSetAttributes/Constants.js';
import { FieldDataTypes } from '/static/sim/Sources/Common/DataModel/DataSet/Constants.js';
import { Representation } from '/static/sim/Sources/Rendering/Core/Property/Constants.js';

import {
  BehaviorCategory,
  ShapeBehavior,
  TextPosition,
} from '/static/sim/Sources/Widgets/Widgets3D/ShapeWidget/Constants.js';

import { VerticalTextAlignment } from '/static/sim/Sources/Widgets/SVG/SVGLandmarkRepresentation/Constants.js';

import { ViewTypes } from '/static/sim/Sources/Widgets/Core/WidgetManager/Constants.js';

import { vec3 } from '/static/sim/gl-matrix/esm/index.js';

// import * as vtkMath from 'vtk.js/Sources/Common/Core/Math/index.js';

// ----------------------------------------------------------------------------
// Standard rendering code setup
// ----------------------------------------------------------------------------

// 2.之后vtk.js有一个全新的类vtkFullScreenRenderWindow，该类将在浏览器全屏显示窗口。
const fullScreenRenderer = vtk.Rendering.Misc.vtkFullScreenRenderWindow.newInstance({
  background: [0, 0, 0],
});
const renderer = fullScreenRenderer.getRenderer();//3.从vtkFullScreenRenderWindow类获取render，
const renderWindow = fullScreenRenderer.getRenderWindow();//4.再次通过vtkFullScreenRenderWindow获取window
const linemapper = vtk.Rendering.Core.vtkMapper.newInstance();
const lineactor = vtk.Rendering.Core.vtkActor.newInstance();

const render = renderWindow.render;

var content = `
  <table>
    <tr>  
      <td><button id="addLineWidget">直线</button></td>
    </tr>
    <tr>
      <td><button id="addPolylineButton">GrabFocus画连续线</button></td>
    </tr>
    <tr> 
      <td><input type="checkbox">显示点号</td>
    </tr>
    <tr>  
      <td><button id="addDimWidget">尺寸</button></td>
    </tr>
    <tr>  
      <td><button id="removeDimWidget">删除尺寸</button></td>
    <tr>
    <tr>  
      <td><button id="mouseBoxSelector">矩形选择框</button></td>
    <tr>
    <tr>
      <td>
        <select class="widget">
          <option name="rectangleWidget">rectangleWidget</option>
          <option name="ellipseWidget" selected>ellipseWidget</option>
          <option name="circleWidget">circleWidget</option>
        </select>
      </td>
    </tr>
    <tr>
      <td>
        <button class="reset">reset</button>
      </td>
      <td>
        <button id="download">保存</button>
      </td>
    </tr>
  </table>

  <table style="display:none">
    <thead>
      <tr>
        <td>Modifier Key</td>
        <td>Left Button</td>
        <td>Middle Button</td>
        <td>Right Button</td>
        <td>Scroll</td>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>None</td>
        <td>
          <select class='leftButton' style="width: 100%">
            <option value='None'>None</option>
            <option value='Pan'>Pan</option>
            <option value='Zoom'>Zoom</option>
            <option value='Roll'>Roll</option>
            <option value='Rotate' selected>Rotate</option>
            <option value='MultiRotate'>MultiRotate</option>
            <option value='ZoomToMouse'>ZoomToMouse</option>
          </select>
        </td>
        <td>
          <select class='middleButton' style="width: 100%">
            <option value='None'>None</option>
            <option value='Pan' selected>Pan</option>
            <option value='Zoom'>Zoom</option>
            <option value='Roll'>Roll</option>
            <option value='Rotate'>Rotate</option>
            <option value='MultiRotate'>MultiRotate</option>
            <option value='ZoomToMouse'>ZoomToMouse</option>
          </select>
        </td>
        <td>
          <select class='rightButton' style="width: 100%">
            <option value='None'>None</option>
            <option value='Pan'>Pan</option>
            <option value='Zoom' selected>Zoom</option>
            <option value='Roll'>Roll</option>
            <option value='Rotate'>Rotate</option>
            <option value='MultiRotate'>MultiRotate</option>
            <option value='ZoomToMouse'>ZoomToMouse</option>
          </select>
        </td>
        <td>
          <select class='scrollMiddleButton' style="width: 100%">
            <option value='None'>None</option>
            <option value='Zoom' selected >Zoom</option>
            <option value='ZoomToMouse' >ZoomToMouse</option>
          </select>
        </td>
      </tr>
      <tr>
        <td>Shift +</td>
        <td>
          <select class='shiftLeftButton' style="width: 100%">
            <option value='None'>None</option>
            <option value='Pan'>Pan</option>
            <option value='Zoom'>Zoom</option>
            <option value='Roll' selected>Roll</option>
            <option value='Rotate'>Rotate</option>
            <option value='MultiRotate'>MultiRotate</option>
            <option value='ZoomToMouse'>ZoomToMouse</option>
          </select>
        </td>
        <td>
          <select class='shiftMiddleButton' style="width: 100%">
            <option value='None'>None</option>
            <option value='Pan'>Pan</option>
            <option value='Zoom'>Zoom</option>
            <option value='Roll'>Roll</option>
            <option value='Rotate' selected>Rotate</option>
            <option value='MultiRotate'>MultiRotate</option>
            <option value='ZoomToMouse'>ZoomToMouse</option>
          </select>
        </td>
        <td>
          <select class='shiftRightButton' style="width: 100%">
            <option value='None'>None</option>
            <option value='Pan' selected>Pan</option>
            <option value='Zoom'>Zoom</option>
            <option value='Roll'>Roll</option>
            <option value='Rotate'>Rotate</option>
            <option value='MultiRotate'>MultiRotate</option>
            <option value='ZoomToMouse'>ZoomToMouse</option>
          </select>
        </td>
        <td>
          <select class='shiftScrollMiddleButton' style="width: 100%">
            <option value='None' selected >None</option>
            <option value='Zoom'>Zoom</option>
            <option value='ZoomToMouse' >ZoomToMouse</option>
          </select>
        </td>
      </tr>
      <tr>
        <td>Ctrl +</td>
        <td>
          <select class='controlLeftButton' style="width: 100%">
            <option value='None'>None</option>
            <option value='Pan'>Pan</option>
            <option value='Zoom' selected>Zoom</option>
            <option value='Roll'>Roll</option>
            <option value='Rotate'>Rotate</option>
            <option value='MultiRotate'>MultiRotate</option>
            <option value='ZoomToMouse'>ZoomToMouse</option>
          </select>
        </td>
        <td>
          <select class='controlMiddleButton' style="width: 100%">
            <option value='None'>None</option>
            <option value='Pan'>Pan</option>
            <option value='Zoom'>Zoom</option>
            <option value='Roll'>Roll</option>
            <option value='Rotate' selected>Rotate</option>
            <option value='MultiRotate'>MultiRotate</option>
            <option value='ZoomToMouse'>ZoomToMouse</option>
          </select>
        </td>
        <td>
          <select class='controlRightButton' style="width: 100%">
            <option value='None'>None</option>
            <option value='Pan'>Pan</option>
            <option value='Zoom'>Zoom</option>
            <option value='Roll'>Roll</option>
            <option value='Rotate'>Rotate</option>
            <option value='MultiRotate'>MultiRotate</option>
            <option value='ZoomToMouse' selected>ZoomToMouse</option>
          </select>
        </td>
        <td>
          <select class='controlScrollMiddleButton' style="width: 100%">
            <option value='None' selected >None</option>
            <option value='Zoom'>Zoom</option>
            <option value='ZoomToMouse' >ZoomToMouse</option>
          </select>
        </td>
      </tr>
      <tr>
        <td>Alt +</td>
        <td>
          <select class='altLeftButton' style="width: 100%">
            <option value='None'>None</option>
            <option value='Pan'>Pan</option>
            <option value='Zoom' selected>Zoom</option>
            <option value='Roll'>Roll</option>
            <option value='Rotate'>Rotate</option>
            <option value='MultiRotate'>MultiRotate</option>
            <option value='ZoomToMouse'>ZoomToMouse</option>
          </select>
        </td>
        <td>
          <select class='altMiddleButton' style="width: 100%">
            <option value='None'>None</option>
            <option value='Pan'>Pan</option>
            <option value='Zoom'>Zoom</option>
            <option value='Roll'>Roll</option>
            <option value='Rotate' selected>Rotate</option>
            <option value='MultiRotate'>MultiRotate</option>
            <option value='ZoomToMouse'>ZoomToMouse</option>
          </select>
        </td>
        <td>
          <select class='altRightButton' style="width: 100%">
            <option value='None'>None</option>
            <option value='Pan'>Pan</option>
            <option value='Zoom'>Zoom</option>
            <option value='Roll'>Roll</option>
            <option value='Rotate'>Rotate</option>
            <option value='MultiRotate'>MultiRotate</option>
            <option value='ZoomToMouse' selected>ZoomToMouse</option>
          </select>
        </td>
        <td>
          <select class='altScrollMiddleButton' style="width: 100%">
            <option value='None' selected>None</option>
            <option value='Zoom'>Zoom</option>
            <option value='ZoomToMouse'>ZoomToMouse</option>
          </select>
        </td>
      </tr>
      <!-- <tr>
        <td>Scroll</td>
        <td colspan="3">
          <select class='scrollMiddleButton' style="width: 100%">
            <option value='None'>None</option>
            <option value='Zoom'>Zoom</option>
          </select>
        </td>
      </tr> -->
    </tbody>
  </table>
`

fullScreenRenderer.addController(content);

// const representationSelector = document.querySelector('.representations');
// const resolutionChange = document.querySelector('.resolution');
// const vrbutton = document.querySelector('.vrbutton');

// const lineSource = vtk.Filters.Sources.vtkLineSource.newInstance();//5.先导入直线数据源，
// linemapper.setInputConnection(lineSource.getOutputPort());//6.然后将数据传入mapper，
// lineactor.getProperty().setPointSize(20);
// // actor.getProperty().setRepresentation(Representation.POINTS);//直线显示的是点状
// lineactor.getProperty().setRepresentation(Representation.WIREFRAME);//直线显示的是线状

// lineactor.setMapper(linemapper);//7.再将actor与mapper绑定。
// linemapper.setInputConnection(lineSource.getOutputPort());
// renderer.addActor(lineactor);//8.将actor加入到render内，
// renderer.resetCamera();
// renderWindow.render();//9.调用window的render()方法开始渲染。

const cube = vtk.Filters.Sources.vtkCubeSource.newInstance();
const widgetmapper = vtk.Rendering.Core.vtkMapper.newInstance();
const widgetactor = vtk.Rendering.Core.vtkActor.newInstance();
// const linemapper = vtk.Rendering.Core.vtkMapper.newInstance();
// const lineactor = vtk.Rendering.Core.vtkActor.newInstance();

widgetactor.setMapper(widgetmapper);
widgetmapper.setInputConnection(cube.getOutputPort());
widgetactor.getProperty().setOpacity(0.5);

renderer.addActor(widgetactor);

const widgetManager = vtk.Widgets.Core.vtkWidgetManager.newInstance();
widgetManager.setRenderer(renderer);

// -----------------------------------------------------------
// 绘制直线
// -----------------------------------------------------------  

document.querySelector('#addLineWidget').addEventListener('click', () => {
  let currentHandle = null;
  widgetManager.releaseFocus(widget);
  widget = vtk.Widgets.Widgets3D.vtkLineWidget.newInstance();
  currentHandle = widgetManager.addWidget(widget);
  lineWidget = currentHandle;
  observeDistance();
  widgetManager.grabFocus(widget);
});

// ----------------------------------------------------------------------------
// Widget manager  标注尺寸
// ----------------------------------------------------------------------------

const linewidgetmapper =  vtk.Rendering.Core.vtkMapper.newInstance();
const linewidgetactor = vtk.Rendering.Core.vtkActor.newInstance();

linewidgetactor.setMapper(linewidgetmapper);
linewidgetmapper.setInputConnection(cube.getOutputPort());
linewidgetactor.getProperty().setOpacity(0.5);

renderer.addActor(linewidgetactor);

const linewidgetManager = vtk.Widgets.Core.vtkWidgetManager.newInstance();
linewidgetManager.setRenderer(renderer);

let widget = null;

let lineWidget = null;
let selectedWidgetIndex = 0;

let getHandle = {};
let serializedState = {};
renderer.resetCamera();

// -----------------------------------------------------------
// Text Modifiers ------------------------------------------
// -----------------------------------------------------------
// 文字位置
function updateLinePos() {
  const input = document.getElementById('linePos').value;
  const subState = lineWidget.getWidgetState().getPositionOnLine();
  subState.setPosOnLine(input / 100);
  lineWidget.placeText();
  renderWindow.render();
}

// function updateText() {
//   const input = document.getElementById('txtIpt').value;
//   lineWidget.setText(input);
//   renderWindow.render();
// }
// document.querySelector('#txtIpt').addEventListener('keyup', updateText);
// updateText();

function observeDistance() {
  // lineWidget.onInteractionEvent(() => {
  //   document.getElementById('distance').innerHTML = widget
  //     .getDistance()
  //     .toFixed(2)+'!';
  // });
  // 实时显示尺寸
  lineWidget.onInteractionEvent(() => {
    lineWidget.setText(widget.getDistance().toFixed(2))
  });

  // lineWidget.onEndInteractionEvent(() => {
  //   document.getElementById('distance').innerHTML = widget
  //     .getDistance()
  //     .toFixed(2)+'~';
  // });
  // 标注尺寸尝试
  lineWidget.onEndInteractionEvent(() => {
    lineWidget.setText(widget.getDistance().toFixed(2))
  });
}

function updateHandleShape(handleId) {
  const e = document.getElementById(`idh${handleId}`);
  // const shape = e.options[e.selectedIndex].value;
  // 改为箭头
  // const shape = e.options[1].value;
  const shape = "cone";
  const handle = getHandle[handleId];
  if (handle) {
    handle.setShape(shape);
    lineWidget.updateHandleVisibility(handleId - 1);
    lineWidget.getInteractor().render();
    // updateCheckBoxes(handleId, shape);
    observeDistance();
  }
}

function setWidgetColor(currentWidget, color) {
  currentWidget.getWidgetState().getHandle1().setColor(color);
  currentWidget.getWidgetState().getHandle2().setColor(color);
  currentWidget.getWidgetState().getMoveHandle().setColor(color);
}

const inputHandle1 = document.getElementById('idh1');
const inputHandle2 = document.getElementById('idh2');

const checkBoxes = ['visiH1', 'visiH2'].map((id) =>
  document.getElementById(id)
);

const handleCheckBoxInput = (e) => {
  if (lineWidget == null) {
    return;
  }
  if (e.target.id === 'visiH1') {
    lineWidget.getWidgetState().getHandle1().setVisible(e.target.checked);
    lineWidget.updateHandleVisibility(0);
  } else {
    lineWidget.getWidgetState().getHandle2().setVisible(e.target.checked);
    lineWidget.updateHandleVisibility(1);
  }
  lineWidget.getInteractor().render();
  renderWindow.render();
};
// checkBoxes.forEach((checkBox) =>
//   checkBox.addEventListener('input', handleCheckBoxInput)
// );

document.querySelector('#addDimWidget').addEventListener('click', () => {
  let currentHandle = null;
  widgetManager.releaseFocus(widget);
  widget = vtk.Widgets.Widgets3D.vtkLineWidget.newInstance();
  // widget.placeWidget(cube.getOutputData().getBounds());
  currentHandle = widgetManager.addWidget(widget);
  lineWidget = currentHandle;

  getHandle = {
    1: lineWidget.getWidgetState().getHandle1(),
    2: lineWidget.getWidgetState().getHandle2(),
  };

  updateHandleShape(1);
  updateHandleShape(2);

  observeDistance();

  widgetManager.grabFocus(widget);

  currentHandle.onStartInteractionEvent(() => {
    const index = widgetManager.getWidgets().findIndex((cwidget) => {
      if (DeepEqual(currentHandle.getWidgetState(), cwidget.getWidgetState()))
        return 1;
      return 0;
    });
    getHandle = {
      1: currentHandle.getWidgetState().getHandle1(),
      2: currentHandle.getWidgetState().getHandle2(),
    };
    setWidgetColor(widgetManager.getWidgets()[selectedWidgetIndex], 0.5);
    setWidgetColor(widgetManager.getWidgets()[index], 0.2);
    selectedWidgetIndex = index;
    lineWidget = currentHandle;
    document.getElementById('idh1').value =
      getHandle[1].getShape() === '' ? 'sphere' : getHandle[1].getShape();
    document.getElementById('idh2').value =
      getHandle[1].getShape() === '' ? 'sphere' : getHandle[2].getShape();
    document.getElementById('visiH1').checked = lineWidget
      .getWidgetState()
      .getHandle1()
      .getVisible();
    document.getElementById('visiH2').checked = lineWidget
      .getWidgetState()
      .getHandle2()
      .getVisible();
    document.getElementById('txtIpt').value = lineWidget
      .getWidgetState()
      .getText()
      .getText();  
  });
  
});

document.querySelector('#removeDimWidget').addEventListener('click', () => {
  widgetManager.removeWidget(widgetManager.getWidgets()[selectedWidgetIndex]);
  if (widgetManager.getWidgets().length !== 0) {
    selectedWidgetIndex = widgetManager.getWidgets().length - 1;
    setWidgetColor(widgetManager.getWidgets()[selectedWidgetIndex], 0.2);

    serializedState = {
      handle1: widgetManager.getWidgets()[selectedWidgetIndex].getWidgetState().getState().handle1.origin,
      handle2: widgetManager.getWidgets()[selectedWidgetIndex].getWidgetState().getState().handle2.origin,
    }
    console.log(serializedState)
  }
});

  
// ----------------------------------------------------------------------------
// 绘制pl线
// ----------------------------------------------------------------------------
// const cone = vtk.Filters.Sources.vtkConeSource.newInstance();

const polylinewidget = vtk.Widgets.Widgets3D.vtkPolyLineWidget.newInstance();
polylinewidget.placeWidget(cube.getOutputData().getBounds());

widgetManager.addWidget(polylinewidget);

// renderer.resetCamera();
widgetManager.enablePicking();
// widgetManager.grabFocus(polylinewidget);

// -----------------------------------------------------------
// UI control handling
// -----------------------------------------------------------
document.querySelector('#addPolylineButton').addEventListener('click', () => {
  widgetManager.grabFocus(polylinewidget);
});

document
  .querySelector('input[type=checkbox]')
  .addEventListener('change', (ev) => {
    widgetManager.setUseSvgLayer(ev.target.checked);
  });
// ----------------------------------------------------------------------------
// 绘制pl线——END
// ----------------------------------------------------------------------------


// ----------------------------------------------------------------------------
// 绘制矩形、椭圆、圆 UI logic
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// Standard rendering code setup
// ----------------------------------------------------------------------------

// scene
const scene = {};

scene.fullScreenRenderer = fullScreenRenderer//vtk.Rendering.Misc.vtkFullScreenRenderWindow.newInstance({
//   rootContainer: document.body,
//   background: [0.1, 0.1, 0.1],
// });

scene.renderer = scene.fullScreenRenderer.getRenderer();
scene.renderWindow = scene.fullScreenRenderer.getRenderWindow();
scene.openGLRenderWindow =
  scene.fullScreenRenderer.getApiSpecificRenderWindow();
scene.camera = scene.renderer.getActiveCamera();

// setup 2D view_——锁定二维平面，但是不能缩放
// scene.camera.setParallelProjection(true);
// scene.iStyle = vtk.Interaction.Style.vtkInteractorStyleImage.newInstance();
// scene.iStyle.setInteractionMode('IMAGE_SLICING');
// scene.renderWindow.getInteractor().setInteractorStyle(scene.iStyle);
// scene.fullScreenRenderer.addController(controlPanel);

// function setCamera(sliceMode, renderer, data) {
//   const ijk = [0, 0, 0];
//   const position = [0, 0, 0];
//   const focalPoint = [0, 0, 0];
//   const viewUp = sliceMode === 1 ? [0, 0, 1] : [0, 1, 0];
//   data.indexToWorld(ijk, focalPoint);
//   ijk[sliceMode] = 1;
//   data.indexToWorld(ijk, position);
//   renderer.getActiveCamera().set({ focalPoint, position, viewUp });
//   renderer.resetCamera();
// }

// ----------------------------------------------------------------------------
// Widget manager
// ----------------------------------------------------------------------------

scene.widgetManager = vtk.Widgets.Core.vtkWidgetManager.newInstance();
scene.widgetManager.setRenderer(scene.renderer);

// Widgets
const widgets = {};
widgets.rectangleWidget = vtk.Widgets.Widgets3D.vtkRectangleWidget.newInstance({
  // resetAfterPointPlacement: true,
  modifierBehavior: {
    None: {
      [BehaviorCategory.PLACEMENT]:
        ShapeBehavior[BehaviorCategory.PLACEMENT].CLICK_AND_DRAG,
      [BehaviorCategory.POINTS]: ShapeBehavior[BehaviorCategory.POINTS].CORNER_TO_CORNER,
      [BehaviorCategory.RATIO]: ShapeBehavior[BehaviorCategory.RATIO].FREE,
    },
  },
});
widgets.ellipseWidget = vtk.Widgets.Widgets3D.vtkEllipseWidget.newInstance({
  modifierBehavior: {
    None: {
      [BehaviorCategory.PLACEMENT]:
        ShapeBehavior[BehaviorCategory.PLACEMENT].CLICK_AND_DRAG,
      [BehaviorCategory.POINTS]:
        ShapeBehavior[BehaviorCategory.POINTS].CORNER_TO_CORNER,
      [BehaviorCategory.RATIO]: ShapeBehavior[BehaviorCategory.RATIO].FREE,
    },
  },
});
widgets.circleWidget = vtk.Widgets.Widgets3D.vtkEllipseWidget.newInstance({
  modifierBehavior: {
    None: {
      [BehaviorCategory.PLACEMENT]:
        ShapeBehavior[BehaviorCategory.PLACEMENT].CLICK_AND_DRAG,
      [BehaviorCategory.POINTS]: ShapeBehavior[BehaviorCategory.POINTS].RADIUS,
      [BehaviorCategory.RATIO]: ShapeBehavior[BehaviorCategory.RATIO].FREE,
    },
  },
});
// Make a large handle for demo purpose
widgets.circleWidget.getWidgetState().getPoint1Handle().setScale1(20);
widgets.circleWidget
  .getWidgetState()
  .setTextPosition([
    TextPosition.MAX,
    TextPosition.CENTER,
    TextPosition.CENTER,
  ]);

scene.rectangleHandle = scene.widgetManager.addWidget(
  widgets.rectangleWidget,
  ViewTypes.SLICE
);
scene.rectangleHandle.setHandleVisibility(true);
scene.rectangleHandle.setTextProps({
  ...scene.rectangleHandle.getTextProps(),
  'text-anchor': 'middle',
  verticalAlign: VerticalTextAlignment.MIDDLE,
});
widgets.rectangleWidget
  .getWidgetState()
  .setTextPosition([
    TextPosition.CENTER,
    TextPosition.CENTER,
    TextPosition.CENTER,
  ]);

scene.ellipseHandle = scene.widgetManager.addWidget(
  widgets.ellipseWidget,
  ViewTypes.SLICE
);
scene.ellipseHandle.setTextProps({
  ...scene.ellipseHandle.getTextProps(),
  'text-anchor': 'middle',
  verticalAlign: VerticalTextAlignment.MIDDLE,
});

scene.circleHandle = scene.widgetManager.addWidget(
  widgets.circleWidget,
  ViewTypes.SLICE
);
scene.circleHandle.setGlyphResolution(64);

scene.widgetManager.grabFocus(widgets.ellipseWidget);
let activeWidget = 'ellipseWidget';

// ----------------------------------------------------------------------------
// Ready logic
// ----------------------------------------------------------------------------

function ready(scope, picking = false) {
  scope.renderer.resetCamera();
  scope.fullScreenRenderer.resize();
  if (picking) {
    scope.widgetManager.enablePicking();
  } else {
    scope.widgetManager.disablePicking();
  }
}

function readyAll() {
  ready(scene, true);
}

function updateControlPanel(im, ds) {
  const slicingMode = im.getSlicingMode();
  const extent = ds.getExtent();
  document.querySelector('.slice').setAttribute('min', extent[slicingMode * 2]);
  document
    .querySelector('.slice')
    .setAttribute('max', extent[slicingMode * 2 + 1]);
}

function updateWidgetVisibility(widget, slicePos, i, widgetIndex) {
  /* testing if the widget is on the slice and has been placed to modify visibility */
  const widgetVisibility =
    !scene.widgetManager.getWidgets()[widgetIndex].getPoint1() ||
    widget.getWidgetState().getPoint1Handle().getOrigin()[i] === slicePos[i];
  return widget.setVisibility(widgetVisibility);
}

function updateWidgetsVisibility(slicePos, slicingMode) {
  updateWidgetVisibility(widgets.rectangleWidget, slicePos, slicingMode, 0);
  updateWidgetVisibility(widgets.ellipseWidget, slicePos, slicingMode, 1);
  updateWidgetVisibility(widgets.circleWidget, slicePos, slicingMode, 2);
}

// ----------------------------------------------------------------------------
// Load image
// ----------------------------------------------------------------------------

// const image = {
//   imageMapper: vtk.Rendering.Core.vtkImageMapper.newInstance(),
//   actor: vtk.Rendering.Core.vtkImageSlice.newInstance(),
// };

// background image pipeline
// image.actor.setMapper(image.imageMapper);

// const reader = vtk.IO.Core.vtkHttpDataSetReader.newInstance({ fetchGzip: true });
// reader
//   .setUrl(`/data/volume/LIDC2.vti`, { loadData: false })
//   .then(() => {
    // const data = reader.getOutputData();
    // image.data = data;

    // set input data
    // image.imageMapper.setInputData(data);

    // add actors to renderers
    // scene.renderer.addViewProp(image.actor);
    // default slice orientation/mode and camera view
    // const sliceMode = vtk.Rendering.Core.vtkImageMapper.SlicingMode.K;
    // image.imageMapper.setSlicingMode(sliceMode);
    // image.imageMapper.setSlice(0);

    // set 2D camera position
    // setCamera(sliceMode, scene.renderer, image.data);

    // updateControlPanel(image.imageMapper, data);

    scene.rectangleHandle.getRepresentations()[1].setDrawBorder(true);
    scene.rectangleHandle.getRepresentations()[1].setDrawFace(false);
    scene.rectangleHandle.getRepresentations()[1].setOpacity(1);
    scene.circleHandle.getRepresentations()[1].setDrawBorder(true);
    scene.circleHandle.getRepresentations()[1].setDrawFace(false);
    scene.circleHandle.getRepresentations()[1].setOpacity(1);
    scene.ellipseHandle.getRepresentations()[1].setDrawBorder(true);
    scene.ellipseHandle.getRepresentations()[1].setDrawFace(false);
    scene.ellipseHandle.getRepresentations()[1].setOpacity(1);

    // set text display callback
    // scene.ellipseHandle.onInteractionEvent(() => {
    //   const worldBounds = scene.ellipseHandle.getBounds();
    //   const { average, minimum, maximum } = image.data.computeHistogram(
    //     worldBounds,
    //     vtkSphere.isPointIn3DEllipse
    //   );

    //   const text = `average: ${average.toFixed(
    //     0
    //   )} \nmin: ${minimum} \nmax: ${maximum} `;

    //   widgets.ellipseWidget.getWidgetState().getText().setText(text);
    // });

    scene.circleHandle.onInteractionEvent(() => {
      const worldBounds = scene.circleHandle.getBounds();

      const text = `radius: ${(
        vec3.distance(
          [worldBounds[0], worldBounds[2], worldBounds[4]],
          [worldBounds[1], worldBounds[3], worldBounds[5]]
        ) / 2
      ).toFixed(2)}`;
      widgets.circleWidget.getWidgetState().getText().setText(text);
    });

    scene.rectangleHandle.onInteractionEvent(() => {
      const worldBounds = scene.rectangleHandle.getBounds();

      const dx = Math.abs(worldBounds[0] - worldBounds[1]);
      const dy = Math.abs(worldBounds[2] - worldBounds[3]);
      const dz = Math.abs(worldBounds[4] - worldBounds[5]);

      const perimeter = 2 * (dx + dy + dz);
      const area = dx * dy + dy * dz + dz * dx;

      const text = `perimeter: ${perimeter.toFixed(1)}mm\narea: ${area.toFixed(
        1
      )}mm²`;
      widgets.rectangleWidget.getWidgetState().getText().setText(text);
    });

    // const update = () => {
    //   // const slicingMode = image.imageMapper.getSlicingMode() % 3;

    //   if (slicingMode > -1) {
    //     const ijk = [0, 0, 0];
    //     const slicePos = [0, 0, 0];

    //     // position
    //     // ijk[slicingMode] = image.imageMapper.getSlice();
    //     // data.indexToWorld(ijk, slicePos);

    //     widgets.rectangleWidget.getManipulator().setOrigin(slicePos);
    //     widgets.ellipseWidget.getManipulator().setOrigin(slicePos);
    //     widgets.circleWidget.getManipulator().setOrigin(slicePos);

    //     updateWidgetsVisibility(slicePos, slicingMode);

    //     scene.renderWindow.render();

    //     // update UI
    //     document
    //       .querySelector('.slice')
    //       .setAttribute('max', data.getDimensions()[slicingMode] - 1);
    //   }
    // };
    // image.imageMapper.onModified(update);
    // trigger initial update
    // update();

    readyAll();
  // });

// register readyAll to resize event
window.addEventListener('resize', readyAll);
readyAll();



// UI logic
function resetWidgets() {
  scene.rectangleHandle.reset();
  scene.ellipseHandle.reset();
  scene.circleHandle.reset();
  // const slicingMode = image.imageMapper.getSlicingMode() % 3;
  // updateWidgetsVisibility(null, slicingMode);
  scene.widgetManager.grabFocus(widgets[activeWidget]);
}

// document.querySelector('.slice').addEventListener('input', (ev) => {
//   image.imageMapper.setSlice(Number(ev.target.value));
// });

// document.querySelector('.axis').addEventListener('input', (ev) => {
//   const sliceMode = 'IJKXYZ'.indexOf(ev.target.value) % 3;
//   // image.imageMapper.setSlicingMode(sliceMode);

//   // setCamera(sliceMode, scene.renderer, image.data);
//   resetWidgets();
//   scene.renderWindow.render();
// });

document.querySelector('.widget').addEventListener('input', (ev) => {
  // For demo purpose, hide ellipse handles when the widget loses focus
  if (activeWidget === 'ellipseWidget') {
    widgets.ellipseWidget.setHandleVisibility(true);
  }
  scene.widgetManager.grabFocus(widgets[ev.target.value]);
  activeWidget = ev.target.value;
  if (activeWidget === 'ellipseWidget') {
    widgets.ellipseWidget.setHandleVisibility(true);
    scene.ellipseHandle.updateRepresentationForRender();
  }

  // scene.renderWindow.getViews()[0].captureNextImage().then((image_data) => {
  //   console.log(image_data)
  //   //do what you want with your image data
  // })

  // renderWindow.captureImages()[0].then(//这个可以，简单
  //   (image) => { console.log(image); }
  // );

});

document.querySelector('.reset').addEventListener('click', () => {
  resetWidgets();
  scene.renderWindow.render();
});
// ----------------------------------------------------------------------------
// 绘制矩形END
// ----------------------------------------------------------------------------

// -----------------------------------------------------------
// 视图BEGIN
// -----------------------------------------------------------
function majorAxis(vec3, idxA, idxB) {
  const axis = [0, 0, 0];
  const idx = Math.abs(vec3[idxA]) > Math.abs(vec3[idxB]) ? idxA : idxB;
  const value = vec3[idx] > 0 ? 1 : -1;
  axis[idx] = value;
  return axis;
}
const axes = vtk.Rendering.Core.vtkAxesActor.newInstance();
const orientationWidget = vtk.Interaction.Widgets.vtkOrientationMarkerWidget.newInstance({
  actor: axes,
  interactor: renderWindow.getInteractor(),
});
orientationWidget.setEnabled(true);
orientationWidget.setViewportCorner(
  vtk.Interaction.Widgets.vtkOrientationMarkerWidget.Corners.BOTTOM_LEFT
);
orientationWidget.setViewportSize(0.15);
orientationWidget.setMinPixelSize(100);
orientationWidget.setMaxPixelSize(300);
// ----------------------------------------------------------------------------
// Widget manager
// ----------------------------------------------------------------------------

const axeswidgetManager = vtk.Widgets.Core.vtkWidgetManager.newInstance();
axeswidgetManager.setRenderer(orientationWidget.getRenderer());

const axeswidget = vtk.Widgets.Widgets3D.vtkInteractiveOrientationWidget.newInstance();
axeswidget.placeWidget(axes.getBounds());
axeswidget.setBounds(axes.getBounds());
axeswidget.setPlaceFactor(1);
const vw = axeswidgetManager.addWidget(axeswidget);
// Manage user interaction
vw.onOrientationChange(({ up, direction, action, event }) => {
  const focalPoint = scene.camera.getFocalPoint();
  const position = scene.camera.getPosition();
  const viewUp = scene.camera.getViewUp();

  const distance = Math.sqrt(
    vtk.Common.Core.vtkMath.distance2BetweenPoints(position, focalPoint)
  );
  scene.camera.setPosition(
    focalPoint[0] + direction[0] * distance,
    focalPoint[1] + direction[1] * distance,
    focalPoint[2] + direction[2] * distance
  );

  if (direction[0]) {
    scene.camera.setViewUp(majorAxis(viewUp, 1, 2));
  }
  if (direction[1]) {
    scene.camera.setViewUp(majorAxis(viewUp, 0, 2));
  }
  if (direction[2]) {
    scene.camera.setViewUp(majorAxis(viewUp, 0, 1));
  }

  orientationWidget.updateMarkerOrientation();
  axeswidgetManager.enablePicking();
  render();
});

renderer.resetCamera();
axeswidgetManager.enablePicking();
render();
// -----------------------------------------------------------
// 视图END
// -----------------------------------------------------------

// -----------------------------------------------------------
// 鼠标控制BEGIN
// -----------------------------------------------------------
const interactorStyle = vtk.Interaction.Style.vtkInteractorStyleManipulator.newInstance();
fullScreenRenderer.getInteractor().setInteractorStyle(interactorStyle);

const uiComponents = {};
const selectMap = {
  leftButton: { button: 1 },
  middleButton: { button: 2 },
  rightButton: { button: 3 },
  shiftLeftButton: { button: 1, shift: true },
  shiftMiddleButton: { button: 2, shift: true },
  shiftRightButton: { button: 3, shift: true },
  controlLeftButton: { button: 1, control: true },
  controlMiddleButton: { button: 2, control: true },
  controlRightButton: { button: 3, control: true },
  altLeftButton: { button: 1, alt: true },
  altMiddleButton: { button: 2, alt: true },
  altRightButton: { button: 3, alt: true },
  scrollMiddleButton: { scrollEnabled: true, dragEnabled: false },
  shiftScrollMiddleButton: {
    scrollEnabled: true,
    dragEnabled: false,
    shift: true,
  },
  controlScrollMiddleButton: {
    scrollEnabled: true,
    dragEnabled: false,
    control: true,
  },
  altScrollMiddleButton: {
    scrollEnabled: true,
    dragEnabled: false,
    alt: true,
  },
};

const manipulatorFactory = {
  None: null,
  Pan: vtk.Interaction.Manipulators.vtkMouseCameraTrackballPanManipulator,
  Zoom: vtk.Interaction.Manipulators.vtkMouseCameraTrackballZoomManipulator,
  Roll: vtk.Interaction.Manipulators.vtkMouseCameraTrackballRollManipulator,
  Rotate: vtk.Interaction.Manipulators.vtkMouseCameraTrackballRotateManipulator,
  MultiRotate: vtk.Interaction.Manipulators.vtkMouseCameraTrackballMultiRotateManipulator,
  ZoomToMouse: vtk.Interaction.Manipulators.vtkMouseCameraTrackballZoomToMouseManipulator,
};

function reassignManipulators() {
  interactorStyle.removeAllMouseManipulators();
  Object.keys(uiComponents).forEach((keyName) => {
    const klass = manipulatorFactory[uiComponents[keyName].manipName];
    if (klass) {
      const manipulator = klass.newInstance();
      manipulator.setButton(selectMap[keyName].button);
      manipulator.setShift(!!selectMap[keyName].shift);
      manipulator.setControl(!!selectMap[keyName].control);
      manipulator.setAlt(!!selectMap[keyName].alt);
      if (selectMap[keyName].scrollEnabled !== undefined) {
        manipulator.setScrollEnabled(selectMap[keyName].scrollEnabled);
      }
      if (selectMap[keyName].dragEnabled !== undefined) {
        manipulator.setDragEnabled(selectMap[keyName].dragEnabled);
      }
      interactorStyle.addMouseManipulator(manipulator);
    }
  });

  // Always add gesture
  interactorStyle.addGestureManipulator(
    vtk.Interaction.Manipulators.vtkGestureCameraManipulator.newInstance()
  );
}

Object.keys(selectMap).forEach((name) => {
  const elt = document.querySelector(`.${name}`);
  elt.addEventListener('change', (e) => {
    vtkDebugMacro(`Changing action of ${name} to ${e.target.value}`);
    uiComponents[name].manipName = e.target.value;
    reassignManipulators();
  });
  uiComponents[name] = {
    elt,
    manipName: elt.value,
  };
});

// Populate with initial manipulators
reassignManipulators();
// -----------------------------------------------------------
// 鼠标控制END
// -----------------------------------------------------------

// -----------------------------------------------------------
// 矩形选择框BEGIN
// -----------------------------------------------------------
const boxSelector = vtk.Interaction.Manipulators.vtkMouseBoxSelectorManipulator.newInstance({
  button: 1,
});
boxSelector.onBoxSelectChange(({ selection }) => {
  console.log('Apply selection:', selection.join(', '));
});
// boxSelector.onBoxSelectInput(console.log);
// document.querySelector('#removeDimWidget').addEventListener('click', () => {
//   widgetManager.removeWidget(widgetManager.getWidgets()[selectedWidgetIndex]);
//   if (widgetManager.getWidgets().length !== 0) {
//     selectedWidgetIndex = widgetManager.getWidgets().length - 1;
//     setWidgetColor(widgetManager.getWidgets()[selectedWidgetIndex], 0.2);
//   }
// });
document.querySelector('#mouseBoxSelector').addEventListener('click', () => {
  // widgetManager.grabFocus(polylinewidget);
  const iStyle = vtk.Interaction.Style.vtkInteractorStyleManipulator.newInstance();
  iStyle.addMouseManipulator(boxSelector);
  renderWindow.getInteractor().setInteractorStyle(iStyle);
});
// -----------------------------------------------------------
// 矩形选择框END
// -----------------------------------------------------------

// -----------------------------------------------------------
// 文件下载BEGIN
// -----------------------------------------------------------

const writer = vtk.IO.XML.vtkXMLPolyDataWriter.newInstance();
writer.setFormat(vtk.IO.XML.vtkXMLWriter.FormatTypes.BINARY);
// writer.setInputConnection(reader.getOutputPort());
// writer.setInputConnection(cube.getOutputPort());
// writer.setInputConnection(widgetManager.getOutputPort());
// writer.setInputConnection(polylinewidget.getOutputPort());
// writer.setInputConnection(linewidgetmapper.getOutputPort());
// writer.setInputConnection(scene.getOutputPort());
// writer.setInputConnection(widgetmapper.getOutputPort());
// writer.setInputConnection(widgetactor.getOutputPort());
// writer.setInputConnection(renderer.getOutputPort());
// writer.setInputConnection(render.getOutputPort());
// writer.setInputConnection(lineWidget.getOutputPort());
// writer.setInputConnection(widgetManager.getWidgets());


document.querySelector('#download').addEventListener('click', () => {
    const writerReader = vtk.IO.XML.vtkXMLPolyDataReader.newInstance();
  // const fileContents = writer.write(reader.getOutputData());


  if (widgetManager.getWidgets().length !== 0) {
    selectedWidgetIndex = widgetManager.getWidgets().length - 1;
    // setWidgetColor(widgetManager.getWidgets()[selectedWidgetIndex], 0.2);

    // serializedState = {
      // handle1: widgetManager.getWidgets()[selectedWidgetIndex].getWidgetState().getState().handle1.origin,
      // handle2: widgetManager.getWidgets()[selectedWidgetIndex].getWidgetState().getState().handle2.origin,
    // }
    const fileContents = writer.write(cube.getOutputData());
    console.log(cube.getOutputData())
    console.log(widgetManager.getWidgets()[selectedWidgetIndex].getWidgetState().getState())
    // fileContents = writer.write(widgetManager.getWidgets()[selectedWidgetIndex].getWidgetState().getState());
    // Try to read it back.
    const textEncoder = new TextEncoder();
    writerReader.parseAsArrayBuffer(textEncoder.encode(fileContents));
  renderer.resetCamera();
  renderWindow.render();

  const blob = new Blob([fileContents], { type: 'text/plain' });
  const a = window.document.createElement('a');
  a.href = window.URL.createObjectURL(blob, { type: 'text/plain' });
  a.download = 'cow.vtp';
  a.text = 'Download';
  a.style.position = 'absolute';
  a.style.left = '50%';
  a.style.bottom = '10px';
  document.body.appendChild(a);
  a.style.background = 'white';
  a.style.padding = '5px';
  }

// scene.renderWindow.getViews()[0].captureNextImage().then((image_data) => {
//   console.log(image_data)
//  //do what you want with your image data
// })
// renderWindow.captureImages()[0].then(
//   (image) => { console.log(image); }
// );
// scene.renderWindow.render()

  // const writerReader = vtk.IO.XML.vtkXMLPolyDataReader.newInstance();
  // // const fileContents = writer.write(reader.getOutputData());
  // const fileContents = writer.write(cube.getOutputData());
  // console.log(cube.getOutputData())
  // const fileContents = writer.write(cube.getOutputData());
  // // Try to read it back.
  // const textEncoder = new TextEncoder();
  // writerReader.parseAsArrayBuffer(textEncoder.encode(fileContents));

  // renderer.resetCamera();
  // renderWindow.render();

  // const blob = new Blob([fileContents], { type: 'text/plain' });
  // const a = window.document.createElement('a');
  // a.href = window.URL.createObjectURL(blob, { type: 'text/plain' });
  // a.download = 'cow.vtp';
  // a.text = 'Download';
  // a.style.position = 'absolute';
  // a.style.left = '50%';
  // a.style.bottom = '10px';
  // document.body.appendChild(a);
  // a.style.background = 'white';
  // a.style.padding = '5px';
})
// -----------------------------------------------------------
// 文件下载END
// -----------------------------------------------------------

// -----------------------------------------------------------
// Make some variables global so that you can inspect and
// modify objects in your browser's developer console:
// -----------------------------------------------------------
var global={}
// global.source = coneSource;
// global.mapper = mapper;
// global.actor = actor;
// global.lineSource = lineSource;
// global.pipelines = pipelines;
global.renderer = renderer;
global.renderWindow = renderWindow;

</script>
</body>

</html>
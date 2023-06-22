
<!DOCTYPE html>
<head>
  <link rel="import" href="/static/sim/Sources/controller.html">
</head>
<html>
<body>

<script src="/static/vue.js/vue.js"></script>
<!-- 下面使用vtk.js完成一个hello world程序

<script type="text/javascript" src="https://unpkg.com/vtk.js"></script>
<script type="text/javascript">
    const cone = vtk.Filters.Sources.vtkConeSource.newInstance();
    const mapper = vtk.Rendering.Core.vtkMapper.newInstance();
    const actor = vtk.Rendering.Core.vtkActor.newInstance();
    mapper.setInputConnection(cone.getOutputPort());
    actor.setMapper(mapper);
 
    const fullRender = vtk.Rendering.Misc.vtkFullScreenRenderWindow.newInstance();
    const render = fullRender.getRenderer();
    render.addActor(actor);
    render.resetCamera();
 
    const renderWindow = fullRender.getRenderWindow();
    renderWindow.render();
</script> 
首先引入vtk.js，然后与vtk类似，
先导入椎体数据源，
然后将数据传入mapper，
再将actor与mapper绑定。
之后vtk.js有一个全新的类vtkFullScreenRenderWindow，该类将在浏览器全屏显示窗口。
从vtkFullScreenRenderWindow类获取render，
将actor加入到render内，
再次通过vtkFullScreenRenderWindow获取window，
调用window的render()方法开始渲染。 -->

<script type="text/javascript" src="https://unpkg.com/vtk.js"></script><!-- 1.首先引入vtk.js -->
<script type="module"> 
// ************##############
// var link = document.createElement('link');
// link.rel = 'import';
// link.href = '/static/sim/Sources/controller.html';
// document.head.appendChild(link);
// // alert(link.import) // null，因为import的内容是异步载入，需要通过下面的onload来处理
// link.onload = function(){
//   alert(link.import)
// };
// link.onerror = function(){
//  alert(link.import)
// };


import '/static/sim/Sources/favicon.js';
// Load the rendering pieces we want to use (for both WebGL and WebGPU)
// import '/static/sim/Sources/Rendering/Profiles/Geometry.js';
// import vtkActor from '/static/sim/Sources/Rendering/Core/Actor/index.js';
// import vtkCalculator from '/static/sim/Sources/Filters/General/Calculator/index.js';
// import vtkConeSource from '/static/sim/Sources/Filters/Sources/ConeSource/index.js';
// import vtkFullScreenRenderWindow from '/static/sim/Sources/Rendering/Misc/FullScreenRenderWindow/index.js';
// import vtkMapper from '/static/sim/Sources/Rendering/Core/Mapper/index.js';
import { AttributeTypes } from '/static/sim/Sources/Common/DataModel/DataSetAttributes/Constants.js';
import { FieldDataTypes } from '/static/sim/Sources/Common/DataModel/DataSet/Constants.js';
import { Representation } from '/static/sim/Sources/Rendering/Core/Property/Constants.js';
// Force DataAccessHelper to have access to various data source
// import '/static/sim/Sources/IO/Core/DataAccessHelper/HtmlDataAccessHelper.js';
// import '/static/sim/Sources/IO/Core/DataAccessHelper/HttpDataAccessHelper.js';
// import '/static/sim/Sources/IO/Core/DataAccessHelper/JSZipDataAccessHelper.js';
// import controlPanel from '/static/sim/Sources/controller.html';

// import vtkRenderWindowWithControlBar from 'vtk.js/Sources/Rendering/Misc/RenderWindowWithControlBar';
// import vtkSlider from 'vtk.js/Sources/Interaction/UI/Slider';
// import vtkCornerAnnotation from 'vtk.js/Sources/Interaction/UI/CornerAnnotation';


// ----------------------------------------------------------------------------
// Standard rendering code setup
// ----------------------------------------------------------------------------

// 2.之后vtk.js有一个全新的类vtkFullScreenRenderWindow，该类将在浏览器全屏显示窗口。
const fullScreenRenderer = vtk.Rendering.Misc.vtkFullScreenRenderWindow.newInstance({
  background: [0, 0, 0],
});
const renderer = fullScreenRenderer.getRenderer();//3.从vtkFullScreenRenderWindow类获取render，
const renderWindow = fullScreenRenderer.getRenderWindow();//4.再次通过vtkFullScreenRenderWindow获取window

// ----------------------------------------------------------------------------
// Example code
// ----------------------------------------------------------------------------
// create a filter on the fly, sort of cool, this is a random scalars
// filter we create inline, for a simple cone you would not need
// this
// ----------------------------------------------------------------------------

// const coneSource = vtk.Filters.Sources.vtkConeSource.newInstance({ height: 100.0, radius: 50 });
// const filter = vtk.Filters.General.vtkCalculator.newInstance();

// filter.setInputConnection(coneSource.getOutputPort());

// filter.setFormulaSimple(FieldDataTypes.CELL, [], 'random', () => Math.random());

// filter.setFormula({
//   getArrays: (inputDataSets) => ({
//     input: [],
//     output: [
//       {
//         location: FieldDataTypes.CELL,
//         name: 'Random',
//         dataType: 'Float32Array',
//         attribute: AttributeTypes.SCALARS,
//       },
//     ],
//   }),
//   evaluate: (arraysIn, arraysOut) => {
//     const [scalars] = arraysOut.map((d) => d.getData());
//     for (let i = 0; i < scalars.length; i++) {
//       scalars[i] = Math.random();
//     }
//   },
// });

const linemapper = vtk.Rendering.Core.vtkMapper.newInstance();
// mapper.setInputConnection(filter.getOutputPort());

const lineactor = vtk.Rendering.Core.vtkActor.newInstance();
// actor.setMapper(mapper);
// actor.setPosition(0.0, 0.0, -20.0);

// renderer.addActor(actor);
// renderer.resetCamera();


// ----------------------------------------------------------------------------
// Create widget
// ----------------------------------------------------------------------------
// const widget = vtk.Interaction.Widgets.vtkHandleWidget.newInstance();
// widget.setInteractor(renderWindow.getInteractor());
// widget.setEnabled(1);

// renderWindow.render();

// -----------------------------------------------------------
// UI control handling
// -----------------------------------------------------------

// fullScreenRenderer.addController(controlPanel);
  // <tr>
  //   <td>
  //     <button class='vrbutton' style="width: 100%">Send To VR</button>
  //   </td>
  // </tr>
  // <tr>
  //   <td>
  //     <select class='representations' style="width: 100%">
  //       <option value='0'>Points</option>
  //       <option value='1'>Wireframe</option>
  //       <option value='2' selected>Surface</option>
  //     </select>
  //   </td>
  // </tr>
  // <tr>
  //   <td>
  //     <input class='resolution' type='text' min='4' max='80' value='6' />
  //   </td>
  // </tr>

var content = `<table>
  // <tr>
  //   <td>Resolution</td>
  //   <td colspan="3">
  //     <input style="width: 100%" class='resolution' type="text" min="1" max="25" step="1" value="10" />
  //   </td>
  // </tr>
  <tr style="text-align: center;">
    <td>直线两点坐标</td>
    <td>X</td>
    <td>Y</td>
    <td>Z</td>
  </tr>
  <tr>
    <td>Point  1</td>
    <td>
      <input style="width: 50px" class='x1' type="text" min="-1" max="1" step="0.1" value="-1" />
    </td>
    <td>
      <input style="width: 50px" class='y1' type="text" min="-1" max="1" step="0.1" value="0" />
    </td>
    <td>
      <input style="width: 50px" class='z1' type="text" min="-1" max="1" step="0.1" value="0" />
    </td>
  </tr>
  <tr>
    <td>Point  2</td>
    <td>
      <input style="width: 50px" class='x2' type="text" min="-1" max="1" step="0.1" value="1" />
    </td>
    <td>
      <input style="width: 50px" class='y2' type="text" min="-1" max="1" step="0.1" value="0" />
    </td>
    <td>
      <input style="width: 50px" class='z2' type="text" min="-1" max="1" step="0.1" value="0" />
    </td>
  </tr>

  
  <tr style="text-align: center;">
    <td>绘制圆</td>
    <td>X</td>
    <td>Y</td>
    <td>Z</td>
  </tr>
  <tr>
    <td>圆心坐标</td>
    <td>
      <input style="width: 50px" class='center' data-index="0" type="text" min="-1" max="1" step="0.1" value="0" />
    </td>
    <td>
      <input style="width: 50px" class='center' data-index="1" type="text" min="-1" max="1" step="0.1" value="0" />
    </td>
    <td>
      <input style="width: 50px" class='center' data-index="2" type="text" min="-1" max="1" step="0.1" value="0" />
    </td>
  </tr>

  <tr>
    <td>半径</td>
    <td colspan="3">
      <input class='radius' type="text" min="0.5" max="2.0" step="0.1" value="1.0" />
    </td>
  </tr>
  <tr>
    <td>Resolution</td>
    <td colspan="3">
      <input class='resolution' type="text" min="4" max="100" step="1" value="6" />
    </td>
  </tr>
  <tr>
    <td>显示边</td>
    <td colspan="3">
      <input class='lines' type="checkbox" checked />
    </td>
  </tr>
  <tr>
    <td>显示面</td>
    <td colspan="3">
      <input class='face' type="checkbox" checked />
    </td>
  </tr>

  <tr style="text-align: center;">
    <td>方向</td>
    <td>X</td>
    <td>Y</td>
    <td>Z</td>
  </tr>
  <tr>
    <td>方向坐标</td>
    <td>
      <input style="width: 50px" class='direction' data-index="0" type="text" min="-1" max="1" step="0.1" value="0" />
    </td>
    <td>
      <input style="width: 50px" class='direction' data-index="1" type="text" min="-1" max="1" step="0.1" value="0" />
    </td>
    <td>
      <input style="width: 50px" class='direction' data-index="2" type="text" min="-1" max="1" step="0.1" value="0" />
    </td>
  </tr>

  <button id="polybutton">GrabFocus画线</button>
  <input type="checkbox">Show SVG layer显示点号
  <br />
  <button id="addLineWidget">直线</button> 
  <br />
  // <div>
  <button id="addDisWidget">Add widget尺寸</button> 
  <br />
  <button id="removeWidget">Remove widget删除尺寸</button> 
  <br />
  <br />
  <div>Distance: <span id="distance">0</span></div>
  <br />
  // <div>Text:</div>
  // <textarea id="txtIpt" name="name"></textarea>
  // <br />
  // <br />
  // <div>Line position</div>
  // <input
  //   title="Line position"
  //   id="linePos"
  //   type="range"
  //   min="0"
  //   max="100"
  //   step="1"
  //   value="50"
  // />
  // <br />

  // <br />
  // <label for="pet-select">Handle1</label>
  // <br />

  // <select name="handle1" class="HandleSource" id="idh1">
  //   <option value="sphere">Sphere</option>
  //   <option value="cone">Cone</option>
  //   <option value="cube">Cube</option>
  //   <option value="triangle">Triangle</option>
  //   <option value="4pointsArrowHead">4 points arrow head</option>
  //   <option value="6pointsArrowHead">6 points arrow head</option>
  //   <option value="star">Star</option>
  //   <option value="disk">Disk</option>
  //   <option value="circle">Circle</option>
  //   <option value="viewFinder">View Finder</option>
  //   <option value="voidSphere">None</option>
  // </select>
  // <div>Visibility
  //   <input id="visiH1" type="checkbox" />
  //   <br />
  // </div>
  // <br />
  // <label for="pet-select">Handle2</label>
  // <br />
  // <select name="handle2" class="HandleSource" id="idh2">
  //   <option value="sphere">Sphere</option>
  //   <option value="cone">Cone</option>
  //   <option value="cube">Cube</option>
  //   <option value="triangle">Triangle</option>
  //   <option value="4pointsArrowHead">4 points arrow head</option>
  //   <option value="6pointsArrowHead">6 points arrow head</option>
  //   <option value="star">Star</option>
  //   <option value="disk">Disk</option>
  //   <option value="circle">Circle</option>
  //   <option value="viewFinder">View Finder</option>
  //   <option value="voidSphere">None</option>
  // </select>
  // <div>Visibility
  //   <input id="visiH2" type="checkbox" />
  //   <br />
  // </div>
// </div>
  
</table>`

fullScreenRenderer.addController(content);

const representationSelector = document.querySelector('.representations');
const resolutionChange = document.querySelector('.resolution');
const vrbutton = document.querySelector('.vrbutton');

// representationSelector.addEventListener('change', (e) => {
//   const newRepValue = Number(e.target.value);
//   actor.getProperty().setRepresentation(newRepValue);
//   renderWindow.render();
// });

// resolutionChange.addEventListener('input', (e) => {
//   console.log(e.target.value);
//   const resolution = Number(e.target.value);
//   coneSource.setResolution(resolution);
//   renderWindow.render();
// });

// vrbutton.addEventListener('click', (e) => {
//   if (vrbutton.textContent === 'Send To VR') {
//     fullScreenRenderer.getApiSpecificRenderWindow().startXR();
//     vrbutton.textContent = 'Return From VR';
//   } else {
//     fullScreenRenderer.getApiSpecificRenderWindow().stopXR();
//     vrbutton.textContent = 'Send To VR';
//   }
// });

// ------------绘制圆
// const fullScreenRenderer = vtkFullScreenRenderWindow.newInstance({
//   background: [0, 0, 0],
// });
// const renderer = fullScreenRenderer.getRenderer();
// const renderWindow = fullScreenRenderer.getRenderWindow();

const lineSource = vtk.Filters.Sources.vtkLineSource.newInstance();//5.先导入直线数据源，
// lineSource.setLines(true);
linemapper.setInputConnection(lineSource.getOutputPort());//6.然后将数据传入mapper，
lineactor.getProperty().setPointSize(20);
// actor.getProperty().setRepresentation(Representation.POINTS);//直线显示的是点状
lineactor.getProperty().setRepresentation(Representation.WIREFRAME);//直线显示的是线状

lineactor.setMapper(linemapper);//7.再将actor与mapper绑定。
linemapper.setInputConnection(lineSource.getOutputPort());
renderer.addActor(lineactor);//8.将actor加入到render内，
renderer.resetCamera();
renderWindow.render();//9.调用window的render()方法开始渲染。

['resolution'].forEach((propertyName) => {
  document.querySelector(`.${propertyName}`).addEventListener('input', (e) => {
    const value = Number(e.target.value);
    lineSource.set({ [propertyName]: value });
    renderWindow.render();
  });
});
const mapping = 'xyz';
const points = [
  [0, 0, 0],
  [0, 0, 0],
];
['x1', 'y1', 'z1', 'x2', 'y2', 'z2'].forEach((propertyName) => {
  document.querySelector(`.${propertyName}`).addEventListener('input', (e) => {
    const value = Number(e.target.value);
    const pointIdx = Number(propertyName[1]);
    points[pointIdx - 1][mapping.indexOf(propertyName[0])] = value;
    lineSource.set({ [`point${pointIdx}`]: points[pointIdx - 1] });
    renderWindow.render();
  });
});

function createCirclePipeline() {
  const cylinderSource = vtk.Filters.Sources.vtkCircleSource.newInstance();
  const circleactor = vtk.Rendering.Core.vtkActor.newInstance();
  const circlemapper = vtk.Rendering.Core.vtkMapper.newInstance();
  cylinderSource.setLines(true);
  cylinderSource.setFace(true);
  
  circleactor.setMapper(circlemapper);
  circlemapper.setInputConnection(cylinderSource.getOutputPort());
  renderer.addActor(circleactor);
  return { cylinderSource, circlemapper, circleactor };
}

const pipelines = [createCirclePipeline()];
pipelines[0].circleactor.getProperty().setColor(1, 0, 0);
// 设置圆的方向
pipelines[0].cylinderSource.setDirection(0, 0, 1);

renderer.resetCamera();
renderWindow.render();
['radius', 'resolution'].forEach((propertyName) => {
  document.querySelector(`.${propertyName}`).addEventListener('input', (e) => {
    const value = Number(e.target.value);
    pipelines[0].cylinderSource.set({ [propertyName]: value });
    renderWindow.render();
  });
});
['lines', 'face'].forEach((propertyName) => {
  document.querySelector(`.${propertyName}`).addEventListener('input', (e) => {
    pipelines[0].cylinderSource.set({ [propertyName]: e.target.checked });
    renderWindow.render();
  });
});

const centerElems = document.querySelectorAll('.center');
function updateTransformedCircle() {
  const center = [0, 0, 0];
  for (let i = 0; i < 3; i++) {
    center[Number(centerElems[i].dataset.index)] = Number(centerElems[i].value);
  }
  pipelines[0].cylinderSource.set({ center });
  renderWindow.render();
}
for (let i = 0; i < 3; i++) {
  centerElems[i].addEventListener('input', updateTransformedCircle);
}

const directionElems = document.querySelectorAll('.direction');
function updateDirectionformedCircle() {
  const direction = [0, 0, 0];
  for (let i = 0; i < 3; i++) {
    direction[Number(directionElems[i].dataset.index)] = Number(directionElems[i].value);
  }
  pipelines[0].cylinderSource.set({ direction });
  renderWindow.render();
}
for (let i = 0; i < 3; i++) {
  directionElems[i].addEventListener('input', updateDirectionformedCircle);
}


// ----------------------------------------------------------------------------
// Widget manager——绘制pl线
// ----------------------------------------------------------------------------
const cone = vtk.Filters.Sources.vtkConeSource.newInstance();
// const lineSource = vtk.Filters.Sources.vtkLineSource.newInstance();//5.先导入直线数据源

const widgetmapper = vtk.Rendering.Core.vtkMapper.newInstance();
const widgetactor = vtk.Rendering.Core.vtkActor.newInstance();
// const linemapper = vtk.Rendering.Core.vtkMapper.newInstance();
// const lineactor = vtk.Rendering.Core.vtkActor.newInstance();

widgetactor.setMapper(widgetmapper);
widgetmapper.setInputConnection(cone.getOutputPort());
widgetactor.getProperty().setOpacity(0.5);

renderer.addActor(widgetactor);

const widgetManager = vtk.Widgets.Core.vtkWidgetManager.newInstance();
widgetManager.setRenderer(renderer);

const polylinewidget = vtk.Widgets.Widgets3D.vtkPolyLineWidget.newInstance();
polylinewidget.placeWidget(cone.getOutputData().getBounds());

widgetManager.addWidget(polylinewidget);

renderer.resetCamera();
widgetManager.enablePicking();
widgetManager.grabFocus(polylinewidget);

// -----------------------------------------------------------
// UI control handling
// -----------------------------------------------------------

// fullScreenRenderer.addController(controlPanel);

document.querySelector('#polybutton').addEventListener('click', () => {
  widgetManager.grabFocus(polylinewidget);
});
// 是否显示点的标签
document
  .querySelector('input[type=checkbox]')
  .addEventListener('change', (ev) => {
    widgetManager.setUseSvgLayer(ev.target.checked);
  });

// ----------------------------------------------------------------------------
// Widget manager  标注尺寸
// ----------------------------------------------------------------------------
const cube = vtk.Filters.Sources.vtkCubeSource.newInstance();
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

function updateText() {
  const input = document.getElementById('txtIpt').value;
  lineWidget.setText(input);
  renderWindow.render();
}
// document.querySelector('#txtIpt').addEventListener('keyup', updateText);
// // updateText();

function observeDistance() {
  lineWidget.onInteractionEvent(() => {
    document.getElementById('distance').innerHTML = widget
      .getDistance()
      .toFixed(2)+'!';
  });
  // 实时显示尺寸
  lineWidget.onInteractionEvent(() => {
    lineWidget.setText(widget.getDistance().toFixed(2))
  });

  lineWidget.onEndInteractionEvent(() => {
    document.getElementById('distance').innerHTML = widget
      .getDistance()
      .toFixed(2)+'~';
  });
  // 标注尺寸尝试
  lineWidget.onEndInteractionEvent(() => {
    lineWidget.setText(widget.getDistance().toFixed(2))
  });
}

// setDistance();
// document.querySelector('#linePos').addEventListener('input', updateLinePos);
// updateLinePos();

// Handle Sources ------------------------------------------

// function updateCheckBoxes(handleId, shape) {
//   if (shape === 'voidSphere') {
//     document
//       .getElementById(`visiH${handleId}`)
//       .setAttribute('disabled', 'disabled');
//   } else if (
//     shape !== 'voidSphere' &&
//     document.getElementById(`visiH${handleId}`).getAttribute('disabled') ===
//       'disabled'
//   ) {
//     document.getElementById(`visiH${handleId}`).removeAttribute('disabled');
//   }
// }

function updateHandleShape(handleId) {
  const e = document.getElementById(`idh${handleId}`);
  // const shape = e.options[e.selectedIndex].value;
  // 改为箭头
  // const shape = e.options[1].value;
  const shape="cone";
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

// inputHandle1.addEventListener('input', updateHandleShape.bind(null, 1));
// inputHandle2.addEventListener('input', updateHandleShape.bind(null, 2));

// inputHandle1.value =
//   getHandle[1].getShape() === '' ? 'sphere' : getHandle[1].getShape();
// inputHandle2.value =
//   getHandle[2].getShape() === '' ? 'sphere' : getHandle[2].getShape();
// updateCheckBoxes(1, getHandles[1].getShape());
// updateCheckBoxes(2, getHandles[2].getShape());

// document.getElementById(
//   'visiH1'
// ).checked = lineWidget.getWidgetState().getHandle1().getVisible();
// document.getElementById(
//   'visiH2'
// ).checked = lineWidget.getWidgetState().getHandle2().getVisible();

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

document.querySelector('#addDisWidget').addEventListener('click', () => {
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

document.querySelector('#removeWidget').addEventListener('click', () => {
  widgetManager.removeWidget(widgetManager.getWidgets()[selectedWidgetIndex]);
  if (widgetManager.getWidgets().length !== 0) {
    selectedWidgetIndex = widgetManager.getWidgets().length - 1;
    setWidgetColor(widgetManager.getWidgets()[selectedWidgetIndex], 0.2);
  }
});



// ----------------------------------------------------------------------------
// Grid on which to map the glyph
// x: 3, y: 2, z: 1
// ----------------------------------------------------------------------------

// const baseGrid = vtk.Common.DataModel.vtkPolyData.newInstance();
// baseGrid.getPoints().setData(Float32Array.from([
//   0, 0, 0,
//   1, 0, 0,
//   2, 0, 0,
//   0, 1, 0,
//   1, 1, 0,
// ]), 3);

// const anglesArray = new Float32Array(5 * 3);
// const anglesDataArray = vtk.Common.Core.vtkDataArray.newInstance({
//   name: 'angles',
//   values: anglesArray,
//   numberOfComponents: 3,
// });
// baseGrid.getPointData().addArray(anglesDataArray);
// function updateAngles() {
//   const x = Number(1);
//   const y = Number(1);
//   const z = Number(1);
//   for (let i = 0; i < anglesArray.length; i += 3) {
//     anglesArray[i] = (Math.PI * x) / 180;
//     anglesArray[i + 1] = (Math.PI * y) / 180;
//     anglesArray[i + 2] = (Math.PI * z) / 180;
//   }
//   anglesDataArray.modified();
//   baseGrid.modified(); // Should not be needed
//   renderWindow.render();
// }
// updateAngles();

// renderer.addActor(pointActor);
// renderer.addActor(textActor);
// renderer.addActor(planeActor);
// ----------------------------------------------------------------------------
// Example code
// ----------------------------------------------------------------------------

// const boxSelector = vtk.Interaction.Manipulators.vtkMouseBoxSelectorManipulator.newInstance({
//   button: 1,
// });
// boxSelector.onBoxSelectChange(({ selection }) => {
//   console.log('Apply selection:', selection.join(', '));
// });

// boxSelector.onBoxSelectInput(console.log);

// const iStyle = vtk.Interaction.Style.vtkInteractorStyleManipulator.newInstance();
// iStyle.addMouseManipulator(boxSelector);
// renderWindow.getInteractor().setInteractorStyle(iStyle);

// -----------------------------------------------------------
// Make some variables global so that you can inspect and
// modify objects in your browser's developer console:
// -----------------------------------------------------------
var global={}
// global.source = coneSource;
// global.mapper = mapper;
// global.actor = actor;
global.lineSource = lineSource;
// global.pipelines = pipelines;
global.renderer = renderer;
global.renderWindow = renderWindow;



// ***************20220310

// const body = document.querySelector('body');
// const rootContainer = document.createElement('div');
// rootContainer.style.position = 'relative';
// rootContainer.style.width = '500px';
// rootContainer.style.height = '500px';

// body.appendChild(rootContainer);
// body.style.margin = '0';

// // Create render window inside container
// const renderWindow = vtk.Rendering.Misc.vtkRenderWindowWithControlBar.newInstance({
//   controlSize: 25,
// });
// renderWindow.setContainer(rootContainer);

// // Add some content to the renderer
// const coneSource = vtk.Filters.Sources.vtkConeSource.newInstance();
// const mapper = vtk.Rendering.Core.vtkMapper.newInstance();
// const actor = vtk.Rendering.Core.vtkActor.newInstance();

// mapper.setInputConnection(coneSource.getOutputPort());
// actor.setMapper(mapper);

// renderWindow.getRenderer().addActor(actor);
// renderWindow.getRenderer().resetCamera();
// renderWindow.getRenderWindow().render();

// // Set control bar to be red so we can see it + layout setup...
// renderWindow.getControlContainer().style.background = '#eee';
// renderWindow.getControlContainer().style.display = 'flex';
// renderWindow.getControlContainer().style.alignItems = 'stretch';
// renderWindow.getControlContainer().style.justifyContent = 'stretch';
// renderWindow.getControlContainer().innerHTML = `
//   <button alt="Left"   title="Left"   value="left">L</button>
//   <button alt="Top"    title="Top"    value="top">T</button>
//   <button alt="Right"  title="Right"  value="right">R</button>
//   <button alt="Bottom" title="Bottom" value="bottom">B</button>
//   <div class="js-slider"></div>
// `;

// // Add corner annotation
// const cornerAnnotation = vtk.Interaction.UI.vtkCornerAnnotation.newInstance();
// cornerAnnotation.setContainer(renderWindow.getRenderWindowContainer());
// cornerAnnotation.getAnnotationContainer().style.color = 'white';
// cornerAnnotation.updateMetadata(coneSource.get('resolution', 'mtime'));
// cornerAnnotation.updateTemplates({
//   nw(meta) {
//     return `<b>Resolution: </b> ${meta.resolution}`;
//   },
//   se(meta) {
//     return `<span style="font-size: 50%"><i style="color: red;">mtime:</i>${meta.mtime}</span><br/>Annotation Corner`;
//   },
// });

// // Add slider to the control bar
// const sliderContainer = renderWindow
//   .getControlContainer()
//   .querySelector('.js-slider');
// sliderContainer.style.flex = '1';
// sliderContainer.style.position = 'relative';
// sliderContainer.style.minWidth = '25px';
// sliderContainer.style.minHeight = '25px';

// const slider = vtk.Interaction.UI.vtkSlider.newInstance();
// slider.generateValues(6, 60, 55);
// slider.setValue(6);
// slider.setContainer(sliderContainer);
// slider.onValueChange((resolution) => {
//   coneSource.set({ resolution });
//   renderWindow.getRenderWindow().render();
//   cornerAnnotation.updateMetadata(coneSource.get('resolution', 'mtime'));
// });

// function updateSizeAndOrientation() {
//   renderWindow.resize();
//   slider.resize();
//   renderWindow.getControlContainer().style.flexFlow = slider.getOrientation()
//     ? 'row'
//     : 'column';
//   setTimeout(slider.resize, 0);
// }
// updateSizeAndOrientation();

// // Handle UI to change bar location
// function onClick(e) {
//   renderWindow.setControlPosition(e.target.value);
//   updateSizeAndOrientation();
// }

// const list = document.querySelectorAll('button');
// let count = list.length;
// while (count--) {
//   list[count].style.width = '25px';
//   list[count].style.height = '25px';
//   list[count].style.flex = 'none';
//   list[count].addEventListener('click', onClick);
// }

// ************

</script>

</body>
</html>


<!-- // var link = document.querySelector('link[rel="import"]');
// link.import.body.querySelector('table').cloneNode(true);
// 使用导入中的内容
// 导入的内容不在主文档中, 仅仅作为主文档附属存在
// 导入的内容
// var content = document.querySelector('link[rel="import"]').import;
// 使用导入内容的DOM元素
// var test = content.querySelector('.test');
// 使用导入内容的样式表
// var styles = content.querySelector('link[rel="stylesheet"]');
// document.head.appendChild(styles.cloneNode(true));

// var pulse = heart.querySelector('div.pulse');
// console.log(link) -->



<!DOCTYPE html>
<head>
  <link rel="import" href="/static/sim/Sources/controller.html">
</head>
<html>
<body>

<!-- <script src="/static/vue.js/vue.js"></script> -->

<script type="text/javascript" src="https://unpkg.com/vtk.js"></script><!-- 1.首先引入vtk.js -->
<script type="module"> 
import '/static/sim/Sources/favicon.js';

import { AttributeTypes } from '/static/sim/Sources/Common/DataModel/DataSetAttributes/Constants.js';
import { FieldDataTypes } from '/static/sim/Sources/Common/DataModel/DataSet/Constants.js';
import { Representation } from '/static/sim/Sources/Rendering/Core/Property/Constants.js';

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

var content = `<table>
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
</table>`

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
  }
});


// ----------------------------------------------------------------------------
// 绘制pl线
// ----------------------------------------------------------------------------
// const cone = vtk.Filters.Sources.vtkConeSource.newInstance();

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

const polylinewidget = vtk.Widgets.Widgets3D.vtkPolyLineWidget.newInstance();
polylinewidget.placeWidget(cube.getOutputData().getBounds());

widgetManager.addWidget(polylinewidget);

renderer.resetCamera();
widgetManager.enablePicking();
widgetManager.grabFocus(polylinewidget);

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


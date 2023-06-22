import 'vtk.js/Sources/favicon';

// Load the rendering pieces we want to use (for both WebGL and WebGPU)
import 'vtk.js/Sources/Rendering/Profiles/All';

import vtkFullScreenRenderWindow from 'vtk.js/Sources/Rendering/Misc/FullScreenRenderWindow';
import vtkInteractorStyleImage from 'vtk.js/Sources/Interaction/Style/InteractorStyleImage';
import vtkSplineWidget from 'vtk.js/Sources/Widgets/Widgets3D/SplineWidget';
import vtkWidgetManager from 'vtk.js/Sources/Widgets/Core/WidgetManager';

import { splineKind } from 'vtk.js/Sources/Common/DataModel/Spline3D/Constants';

import controlPanel from './controlPanel.html';

// ----------------------------------------------------------------------------
// Standard rendering code setup
// ----------------------------------------------------------------------------

const fullScreenRenderer = vtkFullScreenRenderWindow.newInstance({
  background: [0, 0, 0],
});
const renderer = fullScreenRenderer.getRenderer();
const renderWindow = fullScreenRenderer.getRenderWindow();
const iStyle = vtkInteractorStyleImage.newInstance();
renderWindow.getInteractor().setInteractorStyle(iStyle);

// ----------------------------------------------------------------------------
// Widget manager
// ----------------------------------------------------------------------------

const widgetManager = vtkWidgetManager.newInstance();
widgetManager.setRenderer(renderer);

const widget = vtkSplineWidget.newInstance();

const widgetRepresentation = widgetManager.addWidget(widget);

renderer.resetCamera();

// -----------------------------------------------------------
// UI control handling
// -----------------------------------------------------------

fullScreenRenderer.addController(controlPanel);

const tensionInput = document.querySelector('.tension');
const onTensionChanged = () => {
  widget.getWidgetState().setSplineTension(parseFloat(tensionInput.value));
  renderWindow.render();
};
tensionInput.addEventListener('input', onTensionChanged);
onTensionChanged();

const biasInput = document.querySelector('.bias');
const onBiasChanged = () => {
  widget.getWidgetState().setSplineBias(parseFloat(biasInput.value));
  renderWindow.render();
};
biasInput.addEventListener('input', onBiasChanged);
onBiasChanged();

const continuityInput = document.querySelector('.continuity');
const onContinuityChanged = () => {
  widget
    .getWidgetState()
    .setSplineContinuity(parseFloat(continuityInput.value));
  renderWindow.render();
};
continuityInput.addEventListener('input', onContinuityChanged);
onContinuityChanged();

const splineKindSelector = document.querySelector('.kind');
const onSplineKindSelected = () => {
  const isKochanek = splineKindSelector.selectedIndex === 0;
  tensionInput.disabled = !isKochanek;
  biasInput.disabled = !isKochanek;
  continuityInput.disabled = !isKochanek;
  const kind = isKochanek
    ? splineKind.KOCHANEK_SPLINE
    : splineKind.CARDINAL_SPLINE;
  widget.getWidgetState().setSplineKind(kind);
  renderWindow.render();
};
splineKindSelector.addEventListener('change', onSplineKindSelected);
onSplineKindSelected();

const resolutionInput = document.querySelector('.resolution');
const onResolutionChanged = () => {
  widgetRepresentation.setResolution(resolutionInput.value);
  renderWindow.render();
};
resolutionInput.addEventListener('input', onResolutionChanged);
onResolutionChanged();

const handleSizeInput = document.querySelector('.handleSize');
const onHandleSizeChanged = () => {
  widgetRepresentation.setHandleSizeInPixels(handleSizeInput.value);
  renderWindow.render();
};
handleSizeInput.addEventListener('input', onHandleSizeChanged);
onHandleSizeChanged();

const allowFreehandCheckBox = document.querySelector('.allowFreehand');
const onFreehandEnabledChanged = () => {
  widgetRepresentation.setAllowFreehand(allowFreehandCheckBox.checked);
};
allowFreehandCheckBox.addEventListener('click', onFreehandEnabledChanged);
onFreehandEnabledChanged();

const freehandDistanceInput = document.querySelector('.freehandDistance');
const onFreehandDistanceChanged = () => {
  widgetRepresentation.setFreehandMinDistance(freehandDistanceInput.value);
};
freehandDistanceInput.addEventListener('input', onFreehandDistanceChanged);
onFreehandDistanceChanged();

const placeWidgetButton = document.querySelector('.placeWidget');
placeWidgetButton.addEventListener('click', () => {
  widgetRepresentation.reset();
  widgetManager.grabFocus(widget);
  placeWidgetButton.blur();
});

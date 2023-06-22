import test from 'tape-catch';
import testUtils from 'vtk.js/Sources/Testing/testUtils';

import 'vtk.js/Sources/Rendering/Misc/RenderingAPIs';
import vtkRenderWindow from 'vtk.js/Sources/Rendering/Core/RenderWindow';
import vtkRenderer from 'vtk.js/Sources/Rendering/Core/Renderer';
import vtkCylinderSource from 'vtk.js/Sources/Filters/Sources/CylinderSource';
import vtkActor from 'vtk.js/Sources/Rendering/Core/Actor';
import vtkLight from 'vtk.js/Sources/Rendering/Core/Light';
import vtkMapper from 'vtk.js/Sources/Rendering/Core/Mapper';

import baseline from './testCylinder.png';

test.onlyIfWebGL('Test vtkCylinderSource Rendering', (t) => {
  const gc = testUtils.createGarbageCollector(t);
  t.ok('rendering', 'vtkCylinderSource Rendering');

  // Create some control UI
  const container = document.querySelector('body');
  const renderWindowContainer = gc.registerDOMElement(
    document.createElement('div')
  );
  container.appendChild(renderWindowContainer);

  // create what we will view
  const renderWindow = gc.registerResource(vtkRenderWindow.newInstance());
  const renderer = gc.registerResource(vtkRenderer.newInstance());
  renderWindow.addRenderer(renderer);
  renderer.setBackground(0.32, 0.34, 0.43);

  const light1 = gc.registerResource(vtkLight.newInstance());
  light1.setColor(1, 0, 1);
  light1.setPosition(1, 0, 0);
  const light2 = gc.registerResource(vtkLight.newInstance());
  light2.setColor(1, 1, 0);
  light2.setPosition(0, 1, 0);
  renderer.addLight(light1);
  renderer.addLight(light2);

  const actor = gc.registerResource(vtkActor.newInstance());
  renderer.addActor(actor);

  const mapper = gc.registerResource(vtkMapper.newInstance());
  actor.setMapper(mapper);

  const cylinderSource = gc.registerResource(
    vtkCylinderSource.newInstance({ height: 1.0 })
  );
  cylinderSource.setResolution(40);
  mapper.setInputConnection(cylinderSource.getOutputPort());

  // now create something to view it, in this case webgl
  const glwindow = gc.registerResource(renderWindow.newAPISpecificView());
  glwindow.setContainer(renderWindowContainer);
  renderWindow.addView(glwindow);
  glwindow.setSize(400, 400);

  glwindow.captureNextImage().then((image) => {
    testUtils.compareImages(
      image,
      [baseline],
      'Filters/Sources/CylinderSource/testCylinder',
      t,
      2.5,
      gc.releaseResources
    );
  });
  renderWindow.render();
});

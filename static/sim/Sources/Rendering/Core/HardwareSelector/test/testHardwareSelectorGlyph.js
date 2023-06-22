import test from 'tape-catch';
import testUtils from 'vtk.js/Sources/Testing/testUtils';

import vtkCalculator from 'vtk.js/Sources/Filters/General/Calculator';
import 'vtk.js/Sources/Rendering/Misc/RenderingAPIs';
import vtkRenderWindow from 'vtk.js/Sources/Rendering/Core/RenderWindow';
import vtkRenderer from 'vtk.js/Sources/Rendering/Core/Renderer';
import vtkPlaneSource from 'vtk.js/Sources/Filters/Sources/PlaneSource';
import vtkConeSource from 'vtk.js/Sources/Filters/Sources/ConeSource';
import vtkGlyph3DMapper from 'vtk.js/Sources/Rendering/Core/Glyph3DMapper';
import vtkActor from 'vtk.js/Sources/Rendering/Core/Actor';

import { AttributeTypes } from 'vtk.js/Sources/Common/DataModel/DataSetAttributes/Constants';
import {
  FieldAssociations,
  FieldDataTypes,
} from 'vtk.js/Sources/Common/DataModel/DataSet/Constants';

test('Test HardwareSelectorGlyph', (tapeContext) => {
  const gc = testUtils.createGarbageCollector(tapeContext);
  tapeContext.ok('rendering', 'TestHardwareSelectorGlyph');

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

  const planeSource = vtkPlaneSource.newInstance();
  const simpleFilter = vtkCalculator.newInstance();
  const mapper = vtkGlyph3DMapper.newInstance();
  const actor = vtkActor.newInstance();

  simpleFilter.setFormula({
    getArrays: (inputDataSets) => ({
      input: [{ location: FieldDataTypes.COORDINATE }], // Require point coordinates as input
      output: [
        // Generate two output arrays:
        {
          location: FieldDataTypes.POINT, // This array will be point-data ...
          name: 'pressure', // ... with the given name ...
          dataType: 'Float32Array', // ... of this type ...
          numberOfComponents: 3, // ... with this many components ...
        },
        {
          location: FieldDataTypes.POINT, // This array will be field data ...
          name: 'temperature', // ... with the given name ...
          dataType: 'Float32Array', // ... of this type ...
          attribute: AttributeTypes.SCALARS, // ... and will be marked as the default scalars.
          numberOfComponents: 1, // ... with this many components ...
        },
      ],
    }),
    evaluate: (arraysIn, arraysOut) => {
      // Convert in the input arrays of vtkDataArrays into variables
      // referencing the underlying JavaScript typed-data arrays:
      const [coords] = arraysIn.map((d) => d.getData());
      const [press, temp] = arraysOut.map((d) => d.getData());

      // Since we are passed coords as a 3-component array,
      // loop over all the points and compute the point-data output:
      for (let i = 0, sz = coords.length / 3; i < sz; ++i) {
        press[i * 3] = (coords[3 * i] - 0.5) * (coords[3 * i] - 0.5);
        press[i * 3 + 1] =
          ((coords[3 * i + 1] - 0.5) * (coords[3 * i + 1] - 0.5) + 0.125) * 0.1;
        press[i * 3 + 2] =
          ((coords[3 * i] - 0.5) * (coords[3 * i] - 0.5) +
            (coords[3 * i + 1] - 0.5) * (coords[3 * i + 1] - 0.5) +
            0.125) *
          0.1;
        temp[i] = coords[3 * i + 1] * 0.1;
      }
      // Mark the output vtkDataArray as modified
      arraysOut.forEach((x) => x.modified());
    },
  });

  // The generated 'temperature' array will become the default scalars, so the plane mapper will color by 'temperature':
  simpleFilter.setInputConnection(planeSource.getOutputPort());

  mapper.setInputConnection(simpleFilter.getOutputPort(), 0);

  const coneSource = vtkConeSource.newInstance();
  coneSource.setResolution(12);
  mapper.setInputConnection(coneSource.getOutputPort(), 1);
  mapper.setOrientationArray('pressure');
  mapper.setScalarRange(0.0, 0.1);
  renderer.addActor(actor);
  actor.setMapper(mapper);

  // now create something to view it, in this case webgl
  const glwindow = gc.registerResource(renderWindow.newAPISpecificView());
  glwindow.setContainer(renderWindowContainer);
  renderWindow.addView(glwindow);
  glwindow.setSize(400, 400);
  renderer.resetCamera();
  renderer.getActiveCamera().zoom(1.4);
  renderWindow.render();

  const sel = glwindow.getSelector();
  sel.setFieldAssociation(FieldAssociations.FIELD_ASSOCIATION_POINTS);

  sel.selectAsync(renderer, 200, 200, 250, 300).then((res) => {
    const allGood = res.length === 7 && res[0].getProperties().prop === actor;

    tapeContext.ok(res.length === 7, 'Seven glyphs selected');
    tapeContext.ok(
      res[0].getProperties().compositeID === 71,
      'glyph 71 was the first selected'
    );
    tapeContext.ok(allGood, 'Correct prop was selected');

    gc.releaseResources();
  });
});

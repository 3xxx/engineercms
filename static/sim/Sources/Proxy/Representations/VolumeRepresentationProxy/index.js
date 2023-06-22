import macro from 'vtk.js/Sources/macros';
import vtkBoundingBox from 'vtk.js/Sources/Common/DataModel/BoundingBox';
import vtkImageCropFilter from 'vtk.js/Sources/Filters/General/ImageCropFilter';
import vtkImageMapper from 'vtk.js/Sources/Rendering/Core/ImageMapper';
import vtkImageSlice from 'vtk.js/Sources/Rendering/Core/ImageSlice';
import vtkVolume from 'vtk.js/Sources/Rendering/Core/Volume';
import vtkVolumeMapper from 'vtk.js/Sources/Rendering/Core/VolumeMapper';

import vtkAbstractRepresentationProxy from 'vtk.js/Sources/Proxy/Core/AbstractRepresentationProxy';

// ----------------------------------------------------------------------------

function sum(a, b) {
  return a + b;
}

// ----------------------------------------------------------------------------

function mean(...array) {
  return array.reduce(sum, 0) / array.length;
}

// ----------------------------------------------------------------------------

function updateDomains(dataset, dataArray, model, updateProp) {
  const dataRange = dataArray.getRange();
  const spacing = dataset.getSpacing();
  const bounds = dataset.getBounds();

  const { ijkMode: xIJKAxis } = model.mapperX.getClosestIJKAxis();
  const { ijkMode: yIJKAxis } = model.mapperY.getClosestIJKAxis();
  const { ijkMode: zIJKAxis } = model.mapperZ.getClosestIJKAxis();

  const propToUpdate = {
    xSlice: {
      domain: {
        min: bounds[0],
        max: bounds[1],
        step: spacing[xIJKAxis],
      },
    },
    ySlice: {
      domain: {
        min: bounds[2],
        max: bounds[3],
        step: spacing[yIJKAxis],
      },
    },
    zSlice: {
      domain: {
        min: bounds[4],
        max: bounds[5],
        step: spacing[zIJKAxis],
      },
    },
    windowWidth: {
      domain: {
        min: 0,
        max: dataRange[1] - dataRange[0],
        step: 'any',
      },
    },
    windowLevel: {
      domain: {
        min: dataRange[0],
        max: dataRange[1],
        step: 'any',
      },
    },
  };

  updateProp('xSlice', propToUpdate.xSlice);
  updateProp('ySlice', propToUpdate.ySlice);
  updateProp('zSlice', propToUpdate.zSlice);
  updateProp('windowWidth', propToUpdate.windowWidth);
  updateProp('windowLevel', propToUpdate.windowLevel);

  return {
    xSlice: mean(
      propToUpdate.xSlice.domain.min,
      propToUpdate.xSlice.domain.max
    ),
    ySlice: mean(
      propToUpdate.ySlice.domain.min,
      propToUpdate.ySlice.domain.max
    ),
    zSlice: mean(
      propToUpdate.zSlice.domain.min,
      propToUpdate.zSlice.domain.max
    ),
    windowWidth: propToUpdate.windowWidth.domain.max,
    windowLevel: Math.floor(
      mean(
        propToUpdate.windowLevel.domain.min,
        propToUpdate.windowLevel.domain.max
      )
    ),
  };
}

// ----------------------------------------------------------------------------

function updateConfiguration(dataset, dataArray, { mapper, property }) {
  // Configuration
  // actor.getProperty().setInterpolationTypeToFastLinear();
  property.setInterpolationTypeToLinear();

  const numberOfComponents = dataArray.getNumberOfComponents();
  const scalarOpacityUnitDistance =
    vtkBoundingBox.getDiagonalLength(dataset.getBounds()) /
    Math.max(...dataset.getDimensions());
  for (let component = 0; component < numberOfComponents; component++) {
    // For better looking volume rendering
    // - distance in world coordinates a scalar opacity of 1.0
    property.setScalarOpacityUnitDistance(component, scalarOpacityUnitDistance);

    const dataRange = dataArray.getRange(component);
    // - control how we emphasize surface boundaries
    //  => max should be around the average gradient magnitude for the
    //     volume or maybe average plus one std dev of the gradient magnitude
    //     (adjusted for spacing, this is a world coordinate gradient, not a
    //     pixel gradient)
    //  => max hack: (dataRange[1] - dataRange[0]) * 0.05
    property.setGradientOpacityMinimumValue(component, 0);
    property.setGradientOpacityMaximumValue(
      component,
      (dataRange[1] - dataRange[0]) * 0.05
    );
    // - Use shading based on gradient
    property.setShade(true);
    property.setUseGradientOpacity(component, true);
    // - generic good default
    property.setGradientOpacityMinimumOpacity(component, 0.0);
    property.setGradientOpacityMaximumOpacity(component, 1.0);
  }
  property.setAmbient(0.2);
  property.setDiffuse(0.7);
  property.setSpecular(0.3);
  property.setSpecularPower(8.0);
}

// ----------------------------------------------------------------------------
// vtkVolumeRepresentationProxy methods
// ----------------------------------------------------------------------------

function vtkVolumeRepresentationProxy(publicAPI, model) {
  // Set our className
  model.classHierarchy.push('vtkVolumeRepresentationProxy');

  // Volume
  model.mapper = vtkVolumeMapper.newInstance();
  model.volume = vtkVolume.newInstance();
  model.property = model.volume.getProperty();
  model.cropFilter = vtkImageCropFilter.newInstance();
  model.mapper.setInputConnection(model.cropFilter.getOutputPort());

  model.sourceDependencies.push(model.cropFilter);

  // Slices
  model.mapperX = vtkImageMapper.newInstance({
    slicingMode: vtkImageMapper.SlicingMode.X,
  });
  model.actorX = vtkImageSlice.newInstance({ visibility: false });
  model.propertySlices = model.actorX.getProperty();
  model.mapperY = vtkImageMapper.newInstance({
    slicingMode: vtkImageMapper.SlicingMode.Y,
  });
  model.actorY = vtkImageSlice.newInstance({
    visibility: false,
    property: model.propertySlices,
  });
  model.mapperZ = vtkImageMapper.newInstance({
    slicingMode: vtkImageMapper.SlicingMode.Z,
  });
  model.actorZ = vtkImageSlice.newInstance({
    visibility: false,
    property: model.propertySlices,
  });

  model.mapperX.setInputConnection(model.cropFilter.getOutputPort());
  model.mapperY.setInputConnection(model.cropFilter.getOutputPort());
  model.mapperZ.setInputConnection(model.cropFilter.getOutputPort());
  // model.sourceDependencies.push(model.mapperX);
  // model.sourceDependencies.push(model.mapperY);
  // model.sourceDependencies.push(model.mapperZ);

  // connect rendering pipeline
  model.volume.setMapper(model.mapper);
  model.volumes.push(model.volume);

  // Connect slice pipeline
  model.actorX.setMapper(model.mapperX);
  model.actors.push(model.actorX);
  model.actorY.setMapper(model.mapperY);
  model.actors.push(model.actorY);
  model.actorZ.setMapper(model.mapperZ);
  model.actors.push(model.actorZ);

  function setInputData(inputDataset) {
    const [name, location] = publicAPI.getColorBy();
    publicAPI.rescaleTransferFunctionToDataRange(name, location);

    const lutProxy = publicAPI.getLookupTableProxy(name);
    const pwfProxy = publicAPI.getPiecewiseFunctionProxy(name);

    model.property.setRGBTransferFunction(0, lutProxy.getLookupTable());
    model.property.setScalarOpacity(0, pwfProxy.getPiecewiseFunction());

    updateConfiguration(inputDataset, publicAPI.getDataArray(), model);
    if (model.sampleDistance < 0 || model.sampleDistance > 1) {
      publicAPI.setSampleDistance();
    }
    if (model.edgeGradient < 0 || model.edgeGradient > 1) {
      publicAPI.setEdgeGradient();
    }

    // Update domains
    const state = updateDomains(
      inputDataset,
      publicAPI.getDataArray(),
      model,
      publicAPI.updateProxyProperty
    );
    publicAPI.set(state);

    // Check for 2D volumes
    const numberOfDimensions = inputDataset
      .getDimensions()
      .reduce((number, dimension) => number + (dimension > 1 ? 1 : 0), 0);
    if (numberOfDimensions === 2) {
      publicAPI.setIs2DVolume(true);
    }
  }

  model.sourceDependencies.push({ setInputData });

  // API ----------------------------------------------------------------------

  /**
   *  Choose whether the input volume should be treated as a 2D volume
   *  (no volume rendering).
   */
  publicAPI.setIs2DVolume = (is2D) => {
    model.is2DVolume = is2D;
    if (is2D) {
      if (publicAPI.getVisibility()) {
        publicAPI.setSliceVisibility(true);
      }
      model.volume.setVisibility(false);
    }
  };

  publicAPI.isVisible = () => model.volume.getVisibility();

  publicAPI.setVisibility = (isVisible) => {
    if (isVisible) {
      if (model.is2DVolume) {
        publicAPI.setSliceVisibility(true);
      } else {
        model.volume.setVisibility(true);
      }
    } else {
      // Turn off everything
      model.volume.setVisibility(false);
      publicAPI.setSliceVisibility(false);
    }
  };

  publicAPI.getVisibility = () =>
    model.volume.getVisibility() || publicAPI.getSliceVisibility();

  publicAPI.isVisible = publicAPI.getVisibility;

  publicAPI.setSliceVisibility = (isVisible) => {
    if (isVisible && model.is2DVolume) {
      const normalAxis = publicAPI.getInputDataSet().getDimensions().indexOf(1);
      if (model.actors[normalAxis]) {
        model.actors[normalAxis].setVisibility(true);
      }
    } else {
      model.actors.forEach((actor) => actor.setVisibility(isVisible));
    }
  };

  publicAPI.getSliceVisibility = () =>
    model.actorX.getVisibility() ||
    model.actorY.getVisibility() ||
    model.actorZ.getVisibility();

  publicAPI.setSampleDistance = (distance = 0.4) => {
    if (model.sampleDistance !== distance) {
      model.sampleDistance = distance;
      const sourceDS = publicAPI.getInputDataSet();
      const sampleDistance =
        0.7 *
        Math.sqrt(
          sourceDS
            .getSpacing()
            .map((v) => v * v)
            .reduce((a, b) => a + b, 0)
        );
      model.mapper.setSampleDistance(
        sampleDistance * 2 ** (distance * 3.0 - 1.5)
      );

      publicAPI.modified();
    }
  };

  publicAPI.setEdgeGradient = (edgeGradient = 0.2) => {
    if (model.edgeGradient !== edgeGradient) {
      model.edgeGradient = edgeGradient;
      const dataArray = publicAPI.getDataArray();
      const numberOfComponents = dataArray.getNumberOfComponents();
      if (edgeGradient === 0) {
        for (let component = 0; component < numberOfComponents; component++) {
          model.volume.getProperty().setUseGradientOpacity(component, false);
        }
      } else {
        for (let component = 0; component < numberOfComponents; component++) {
          const dataRange = dataArray.getRange(component);
          model.volume.getProperty().setUseGradientOpacity(component, true);
          const minV = Math.max(0.0, edgeGradient - 0.3) / 0.7;
          if (minV > 0.0) {
            model.volume
              .getProperty()
              .setGradientOpacityMinimumValue(
                component,
                Math.exp(
                  Math.log((dataRange[1] - dataRange[0]) * 0.2) * minV * minV
                )
              );
          } else {
            model.volume
              .getProperty()
              .setGradientOpacityMinimumValue(component, 0.0);
          }
          model.volume
            .getProperty()
            .setGradientOpacityMaximumValue(
              component,
              Math.exp(
                Math.log((dataRange[1] - dataRange[0]) * 1.0) *
                  edgeGradient *
                  edgeGradient
              )
            );
        }
      }
      publicAPI.modified();
    }
  };

  const parentSetColorBy = publicAPI.setColorBy;
  publicAPI.setColorBy = (arrayName, arrayLocation, componentIndex = -1) => {
    parentSetColorBy(arrayName, arrayLocation, componentIndex);
    const lutProxy = publicAPI.getLookupTableProxy(arrayName);
    const pwfProxy = publicAPI.getPiecewiseFunctionProxy(arrayName);
    model.property.setRGBTransferFunction(0, lutProxy.getLookupTable());
    model.property.setScalarOpacity(0, pwfProxy.getPiecewiseFunction());
  };
}

// ----------------------------------------------------------------------------
// Object factory
// ----------------------------------------------------------------------------

const DEFAULT_VALUES = {
  sampleDistance: -1,
  edgeGradient: -1,
  disableSolidColor: true,
  is2DVolume: false,
};

// ----------------------------------------------------------------------------

export function extend(publicAPI, model, initialValues = {}) {
  Object.assign(model, DEFAULT_VALUES, initialValues);

  // Object methods
  vtkAbstractRepresentationProxy.extend(publicAPI, model, initialValues);
  macro.get(publicAPI, model, [
    'sampleDistance',
    'edgeGradient',
    'cropFilter',
    'is2DVolume',
  ]);

  // Object specific methods
  vtkVolumeRepresentationProxy(publicAPI, model);
  macro.proxyPropertyMapping(publicAPI, model, {
    xSlice: { modelKey: 'mapperX', property: 'slice' },
    ySlice: { modelKey: 'mapperY', property: 'slice' },
    zSlice: { modelKey: 'mapperZ', property: 'slice' },
    volumeVisibility: { modelKey: 'volume', property: 'visibility' },
    xSliceVisibility: { modelKey: 'actorX', property: 'visibility' },
    ySliceVisibility: { modelKey: 'actorY', property: 'visibility' },
    zSliceVisibility: { modelKey: 'actorZ', property: 'visibility' },
    windowWidth: { modelKey: 'propertySlices', property: 'colorWindow' },
    windowLevel: { modelKey: 'propertySlices', property: 'colorLevel' },
    useShadow: { modelKey: 'property', property: 'shade' },
    croppingPlanes: { modelKey: 'cropFilter', property: 'croppingPlanes' },
  });
}

// ----------------------------------------------------------------------------

export const newInstance = macro.newInstance(
  extend,
  'vtkVolumeRepresentationProxy'
);

// ----------------------------------------------------------------------------

export default { newInstance, extend, updateConfiguration };

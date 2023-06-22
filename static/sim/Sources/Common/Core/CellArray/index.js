import macro from 'vtk.js/Sources/macros';
import vtkDataArray from 'vtk.js/Sources/Common/Core/DataArray';
import { VtkDataTypes } from 'vtk.js/Sources/Common/Core/DataArray/Constants';

// ----------------------------------------------------------------------------
// Global methods
// ----------------------------------------------------------------------------

function extractCellSizes(cellArray) {
  let currentIdx = 0;
  return cellArray.filter((value, index) => {
    if (index === currentIdx) {
      currentIdx += value + 1;
      return true;
    }
    return false;
  });
}

function getNumberOfCells(cellArray) {
  return extractCellSizes(cellArray).length;
}

// ----------------------------------------------------------------------------
// Static API
// ----------------------------------------------------------------------------

export const STATIC = {
  extractCellSizes,
  getNumberOfCells,
};

// ----------------------------------------------------------------------------
// vtkCellArray methods
// ----------------------------------------------------------------------------

function vtkCellArray(publicAPI, model) {
  // Set our className
  model.classHierarchy.push('vtkCellArray');

  publicAPI.getNumberOfCells = (recompute) => {
    if (model.numberOfCells !== undefined && !recompute) {
      return model.numberOfCells;
    }

    model.cellSizes = extractCellSizes(model.values);
    model.numberOfCells = model.cellSizes.length;
    return model.numberOfCells;
  };

  publicAPI.getCellSizes = (recompute) => {
    if (model.cellSizes !== undefined && !recompute) {
      return model.cellSizes;
    }

    model.cellSizes = extractCellSizes(model.values);
    return model.cellSizes;
  };

  const superSetData = publicAPI.setData;
  publicAPI.setData = (typedArray) => {
    superSetData(typedArray, 1);
    model.numberOfCells = undefined;
    model.cellSizes = undefined;
  };

  /**
   * Returns the point indexes at the given location as a subarray.
   */
  publicAPI.getCell = (loc) => {
    let cellLoc = loc;
    const numberOfPoints = model.values[cellLoc++];
    return model.values.subarray(cellLoc, cellLoc + numberOfPoints);
  };
}

// ----------------------------------------------------------------------------
// Object factory
// ----------------------------------------------------------------------------

function defaultValues(initialValues) {
  return {
    empty: true,
    numberOfComponents: 1,
    dataType: VtkDataTypes.UNSIGNED_INT,
    ...initialValues,
  };
}

// ----------------------------------------------------------------------------

export function extend(publicAPI, model, initialValues = {}) {
  vtkDataArray.extend(publicAPI, model, defaultValues(initialValues));
  vtkCellArray(publicAPI, model);
}

// ----------------------------------------------------------------------------

export const newInstance = macro.newInstance(extend, 'vtkCellArray');

// ----------------------------------------------------------------------------

export default { newInstance, extend, ...STATIC };

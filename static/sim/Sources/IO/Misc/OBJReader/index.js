import macro from 'vtk.js/Sources/macros';
import DataAccessHelper from 'vtk.js/Sources/IO/Core/DataAccessHelper';
import vtkDataArray from 'vtk.js/Sources/Common/Core/DataArray';
import vtkPolyData from 'vtk.js/Sources/Common/DataModel/PolyData';

// Enable several sources for DataAccessHelper
import 'vtk.js/Sources/IO/Core/DataAccessHelper/LiteHttpDataAccessHelper'; // Just need HTTP
// import 'vtk.js/Sources/IO/Core/DataAccessHelper/HttpDataAccessHelper'; // HTTP + gz
// import 'vtk.js/Sources/IO/Core/DataAccessHelper/HtmlDataAccessHelper'; // html + base64 + zip
// import 'vtk.js/Sources/IO/Core/DataAccessHelper/JSZipDataAccessHelper'; // zip

// ----------------------------------------------------------------------------

const data = {};

// ----------------------------------------------------------------------------

function pushVector(src, srcOffset, dst, vectorSize) {
  for (let i = 0; i < vectorSize; i++) {
    dst.push(src[srcOffset + i]);
  }
}

// ----------------------------------------------------------------------------

function begin(splitMode) {
  data.splitOn = splitMode;
  data.pieces = [];
  data.v = [];
  data.vt = [];
  data.vn = [];
  data.f = [[]];
  data.size = 0;
}

// ----------------------------------------------------------------------------

function faceMap(str) {
  const idxs = str.split('/').map((i) => Number(i));
  const vertexIdx = idxs[0] - 1;
  const textCoordIdx = idxs[1] ? idxs[1] - 1 : vertexIdx;
  const vertexNormal = idxs[2] ? idxs[2] - 1 : vertexIdx;
  return [vertexIdx, textCoordIdx, vertexNormal];
}

// ----------------------------------------------------------------------------

function parseLine(line) {
  if (line[0] === '#') {
    return;
  }
  const tokens = line.split(/[ \t]+/);
  if (tokens[0] === data.splitOn) {
    tokens.shift();
    data.pieces.push(tokens.join(' ').trim());
    data.f.push([]);
    data.size++;
  } else if (tokens[0] === 'v') {
    data.v.push(Number(tokens[1]));
    data.v.push(Number(tokens[2]));
    data.v.push(Number(tokens[3]));
  } else if (tokens[0] === 'vt') {
    data.vt.push(Number(tokens[1]));
    data.vt.push(Number(tokens[2]));
  } else if (tokens[0] === 'vn') {
    data.vn.push(Number(tokens[1]));
    data.vn.push(Number(tokens[2]));
    data.vn.push(Number(tokens[3]));
  } else if (tokens[0] === 'f') {
    // Handle triangles for now
    if (data.size === 0) {
      data.size++;
    }
    const cells = data.f[data.size - 1];
    tokens.shift();
    const faces = tokens.filter((s) => s.length > 0 && s !== '\r');
    const size = faces.length;
    cells.push(size);
    for (let i = 0; i < size; i++) {
      cells.push(faceMap(faces[i]));
    }
  }
}

// ----------------------------------------------------------------------------

function end(model) {
  const hasTcoords = !!data.vt.length;
  const hasNormals = !!data.vn.length;
  if (model.splitMode) {
    model.numberOfOutputs = data.size;
    for (let idx = 0; idx < data.size; idx++) {
      const ctMapping = {};
      const polydata = vtkPolyData.newInstance({ name: data.pieces[idx] });
      const pts = [];
      const tc = [];
      const normals = [];
      const polys = [];

      const polyIn = data.f[idx];
      const nbElems = polyIn.length;
      let offset = 0;
      while (offset < nbElems) {
        const cellSize = polyIn[offset];
        polys.push(cellSize);
        for (let pIdx = 0; pIdx < cellSize; pIdx++) {
          const [vIdx, tcIdx, nIdx] = polyIn[offset + pIdx + 1];
          const key = `${vIdx}/${tcIdx}/${nIdx}`;
          if (ctMapping[key] === undefined) {
            ctMapping[key] = pts.length / 3;
            pushVector(data.v, vIdx * 3, pts, 3);
            if (hasTcoords) {
              pushVector(data.vt, tcIdx * 2, tc, 2);
            }
            if (hasNormals) {
              pushVector(data.vn, nIdx * 3, normals, 3);
            }
          }
          polys.push(ctMapping[key]);
        }
        offset += cellSize + 1;
      }

      polydata.getPoints().setData(Float32Array.from(pts), 3);
      polydata.getPolys().setData(Uint32Array.from(polys));

      if (hasTcoords) {
        const tcoords = vtkDataArray.newInstance({
          numberOfComponents: 2,
          values: Float32Array.from(tc),
          name: 'TextureCoordinates',
        });
        polydata.getPointData().setTCoords(tcoords);
      }

      if (hasNormals) {
        const normalsArray = vtkDataArray.newInstance({
          numberOfComponents: 3,
          values: Float32Array.from(normals),
          name: 'Normals',
        });
        polydata.getPointData().setNormals(normalsArray);
      }

      // register in output
      model.output[idx] = polydata;
    }
  } else {
    model.numberOfOutputs = 1;
    const polydata = vtkPolyData.newInstance();
    polydata.getPoints().setData(Float32Array.from(data.v), 3);
    if (hasTcoords && data.v.length / 3 === data.vt.length / 2) {
      const tcoords = vtkDataArray.newInstance({
        numberOfComponents: 2,
        values: Float32Array.from(data.vt),
        name: 'TextureCoordinates',
      });
      polydata.getPointData().setTCoords(tcoords);
    }
    if (hasNormals && data.v.length === data.vn.length) {
      const normalsArray = vtkDataArray.newInstance({
        numberOfComponents: 3,
        values: Float32Array.from(data.vn),
        name: 'Normals',
      });
      polydata.getPointData().setNormals(normalsArray);
    }

    const polys = [];
    const polyIn = data.f[0];
    const nbElems = polyIn.length;
    let offset = 0;
    while (offset < nbElems) {
      const cellSize = polyIn[offset];
      polys.push(cellSize);
      for (let pIdx = 0; pIdx < cellSize; pIdx++) {
        const [vIdx] = polyIn[offset + pIdx + 1];
        polys.push(vIdx);
      }
      offset += cellSize + 1;
    }
    polydata.getPolys().setData(Uint32Array.from(polys));
    model.output[0] = polydata;
  }
}

// ----------------------------------------------------------------------------
// vtkOBJReader methods
// ----------------------------------------------------------------------------

function vtkOBJReader(publicAPI, model) {
  // Set our className
  model.classHierarchy.push('vtkOBJReader');

  // Create default dataAccessHelper if not available
  if (!model.dataAccessHelper) {
    model.dataAccessHelper = DataAccessHelper.get('http');
  }

  // Internal method to fetch Array
  function fetchData(url, option = {}) {
    return model.dataAccessHelper.fetchText(publicAPI, url, option);
  }

  // Set DataSet url
  publicAPI.setUrl = (url, option = {}) => {
    if (url.indexOf('.obj') === -1 && !option.fullpath) {
      model.baseURL = url;
      model.url = `${url}/index.obj`;
    } else {
      model.url = url;

      // Remove the file in the URL
      const path = url.split('/');
      path.pop();
      model.baseURL = path.join('/');
    }

    // Fetch metadata
    return publicAPI.loadData(option);
  };

  // Fetch the actual data arrays
  publicAPI.loadData = (option = {}) =>
    fetchData(model.url, option).then(publicAPI.parseAsText);

  publicAPI.parseAsText = (content) => {
    if (!content) {
      return true;
    }
    if (content !== model.parseData) {
      publicAPI.modified();
    }
    model.parseData = content;
    model.numberOfOutputs = 0;
    begin(model.splitMode);
    content.split('\n').forEach(parseLine);
    end(model);
    return true;
  };

  publicAPI.requestData = (inData, outData) => {
    publicAPI.parseAsText(model.parseData);
  };

  // return Busy state
  publicAPI.isBusy = () => !!model.requestCount;

  publicAPI.getNumberOfOutputPorts = () => model.numberOfOutputs;
}

// ----------------------------------------------------------------------------
// Object factory
// ----------------------------------------------------------------------------

const DEFAULT_VALUES = {
  numberOfOutputs: 1,
  requestCount: 0,
  splitMode: null,
  // baseURL: null,
  // dataAccessHelper: null,
  // url: null,
};

// ----------------------------------------------------------------------------

export function extend(publicAPI, model, initialValues = {}) {
  Object.assign(model, DEFAULT_VALUES, initialValues);

  // Build VTK API
  macro.obj(publicAPI, model);
  macro.get(publicAPI, model, ['url', 'baseURL']);
  macro.setGet(publicAPI, model, ['dataAccessHelper', 'splitMode']);
  macro.algo(publicAPI, model, 0, 1);
  macro.event(publicAPI, model, 'busy');

  // Object methods
  vtkOBJReader(publicAPI, model);
}

// ----------------------------------------------------------------------------

export const newInstance = macro.newInstance(extend, 'vtkOBJReader');

// ----------------------------------------------------------------------------

export default { newInstance, extend };

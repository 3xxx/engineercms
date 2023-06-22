import pako from 'pako';

import macro from 'vtk.js/Sources/macros';
import Base64 from 'vtk.js/Sources/Common/Core/Base64';
import Endian from 'vtk.js/Sources/Common/Core/Endian';
import { DataTypeByteSize } from 'vtk.js/Sources/Common/Core/DataArray/Constants';
import { registerType } from 'vtk.js/Sources/IO/Core/DataAccessHelper';

const { vtkErrorMacro, vtkDebugMacro } = macro;

let requestCount = 0;

function getContent(url) {
  const el = document.querySelector(`.webResource[data-url="${url}"]`);
  return el ? el.innerHTML : null;
}

function getElement(url) {
  return document.querySelector(`.webResource[data-url="${url}"]`);
}

function removeLeadingSlash(str) {
  return str[0] === '/' ? str.substr(1) : str;
}

function fetchText(instance = {}, url, options = {}) {
  return new Promise((resolve, reject) => {
    const txt = getContent(url);
    if (txt === null) {
      reject(new Error(`No such text ${url}`));
    } else {
      resolve(txt);
    }
  });
}

function fetchJSON(instance = {}, url, options = {}) {
  return new Promise((resolve, reject) => {
    const txt = getContent(removeLeadingSlash(url));
    if (txt === null) {
      reject(new Error(`No such JSON ${url}`));
    } else {
      resolve(JSON.parse(txt));
    }
  });
}

function fetchArray(instance = {}, baseURL, array, options = {}) {
  return new Promise((resolve, reject) => {
    const url = removeLeadingSlash(
      [
        baseURL,
        array.ref.basepath,
        options.compression ? `${array.ref.id}.gz` : array.ref.id,
      ].join('/')
    );

    const txt = getContent(url);
    if (txt === null) {
      reject(new Error(`No such array ${url}`));
    } else {
      if (array.dataType === 'string') {
        let bText = atob(txt);
        if (options.compression) {
          bText = pako.inflate(bText, { to: 'string' });
        }
        array.values = JSON.parse(bText);
      } else {
        const uint8array = new Uint8Array(Base64.toArrayBuffer(txt));

        array.buffer = new ArrayBuffer(uint8array.length);

        // copy uint8array to buffer
        const view = new Uint8Array(array.buffer);
        view.set(uint8array);

        if (options.compression) {
          if (array.dataType === 'string' || array.dataType === 'JSON') {
            array.buffer = pako.inflate(new Uint8Array(array.buffer), {
              to: 'string',
            });
          } else {
            array.buffer = pako.inflate(new Uint8Array(array.buffer)).buffer;
          }
        }

        if (array.ref.encode === 'JSON') {
          array.values = JSON.parse(array.buffer);
        } else {
          if (Endian.ENDIANNESS !== array.ref.encode && Endian.ENDIANNESS) {
            // Need to swap bytes
            vtkDebugMacro(`Swap bytes of ${array.name}`);
            Endian.swapBytes(array.buffer, DataTypeByteSize[array.dataType]);
          }

          array.values = macro.newTypedArray(array.dataType, array.buffer);
        }

        if (array.values.length !== array.size) {
          vtkErrorMacro(
            `Error in FetchArray: ${array.name} does not have the proper array size. Got ${array.values.length}, instead of ${array.size}`
          );
        }
      }

      // Done with the ref and work
      delete array.ref;
      if (--requestCount === 0 && instance.invokeBusy) {
        instance.invokeBusy(false);
      }
      if (instance.modified) {
        instance.modified();
      }

      resolve(array);
    }
  });
}

// ----------------------------------------------------------------------------

function fetchImage(instance = {}, url, options = {}) {
  return new Promise((resolve, reject) => {
    const img = getElement(url);
    if (img) {
      resolve(img);
    } else {
      reject(new Error(`No such image ${url}`));
    }
  });
}

// ----------------------------------------------------------------------------

const HtmlDataAccessHelper = {
  fetchJSON,
  fetchText,
  fetchArray,
  fetchImage,
};

registerType('html', (options) => HtmlDataAccessHelper);

// Export fetch methods
export default HtmlDataAccessHelper;

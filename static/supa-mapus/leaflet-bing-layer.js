(function e(t,n,r){function s(o,u){if(!n[o]){if(!t[o]){var a=typeof require=="function"&&require;if(!u&&a)return a(o,!0);if(i)return i(o,!0);var f=new Error("Cannot find module '"+o+"'");throw f.code="MODULE_NOT_FOUND",f}var l=n[o]={exports:{}};t[o][0].call(l.exports,function(e){var n=t[o][1][e];return s(n?n:e)},l,l.exports,e,t,n,r)}return n[o].exports}var i=typeof require=="function"&&require;for(var o=0;o<r.length;o++)s(r[o]);return s})({1:[function(require,module,exports){
(function (global){
var L = (typeof window !== "undefined" ? window['L'] : typeof global !== "undefined" ? global['L'] : null)
var fetchJsonp = require('fetch-jsonp')
var bboxIntersect = require('bbox-intersect')

/**
 * Converts tile xyz coordinates to Quadkey
 * @param {Number} x
 * @param {Number} y
 * @param {Number} z
 * @return {Number} Quadkey
 */
function toQuadKey (x, y, z) {
  var index = ''
  for (var i = z; i > 0; i--) {
    var b = 0
    var mask = 1 << (i - 1)
    if ((x & mask) !== 0) b++
    if ((y & mask) !== 0) b += 2
    index += b.toString()
  }
  return index
}

/**
 * Converts Leaflet BBoxString to Bing BBox
 * @param {String} bboxString 'southwest_lng,southwest_lat,northeast_lng,northeast_lat'
 * @return {Array} [south_lat, west_lng, north_lat, east_lng]
 */
function toBingBBox (bboxString) {
  var bbox = bboxString.split(',')
  return [bbox[1], bbox[0], bbox[3], bbox[2]]
}

var VALID_IMAGERY_SETS = [
  'Aerial',
  'AerialWithLabels',
  'AerialWithLabelsOnDemand',
  'Road',
  'RoadOnDemand',
  'CanvasLight',
  'CanvasDark',
  'CanvasGray',
  'OrdnanceSurvey'
]

var DYNAMIC_IMAGERY_SETS = [
  'AerialWithLabelsOnDemand',
  'RoadOnDemand'
]

/**
 * Create a new Bing Maps layer.
 * @param {string|object} options Either a [Bing Maps Key](https://msdn.microsoft.com/en-us/library/ff428642.aspx) or an options object
 * @param {string} options.BingMapsKey A valid Bing Maps Key (required)
 * @param {string} [options.imagerySet=Aerial] Type of imagery, see https://msdn.microsoft.com/en-us/library/ff701716.aspx
 * @param {string} [options.culture='en-US'] Language for labels, see https://msdn.microsoft.com/en-us/library/hh441729.aspx
 * @return {L.TileLayer} A Leaflet TileLayer to add to your map
 *
 * Create a basic map
 * @example
 * var map = L.map('map').setView([51.505, -0.09], 13)
 * L.TileLayer.Bing(MyBingMapsKey).addTo(map)
 */
L.TileLayer.Bing = L.TileLayer.extend({
  options: {
    bingMapsKey: null, // Required
    imagerySet: 'Aerial',
    culture: 'en-US',
    minZoom: 1,
    maxZoom: 23,
    minNativeZoom: 1,
    maxNativeZoom: 19
  },

  statics: {
    METADATA_URL: 'https://dev.virtualearth.net/REST/v1/Imagery/Metadata/{imagerySet}?key={bingMapsKey}&include=ImageryProviders&uriScheme=https&c={culture}',
    POINT_METADATA_URL: 'https://dev.virtualearth.net/REST/v1/Imagery/Metadata/{imagerySet}/{lat},{lng}?zl={z}&key={bingMapsKey}&uriScheme=https&c={culture}'
  },

  initialize: function (options) {
    if (typeof options === 'string') {
      options = { bingMapsKey: options }
    }
    if (options && options.BingMapsKey) {
      options.bingMapsKey = options.BingMapsKey
      console.warn('use options.bingMapsKey instead of options.BingMapsKey')
    }
    if (!options || !options.bingMapsKey) {
      throw new Error('Must supply options.BingMapsKey')
    }
    options = L.setOptions(this, options)
    if (VALID_IMAGERY_SETS.indexOf(options.imagerySet) < 0) {
      throw new Error("'" + options.imagerySet + "' is an invalid imagerySet, see https://github.com/digidem/leaflet-bing-layer#parameters")
    }
    if (options && options.style && DYNAMIC_IMAGERY_SETS.indexOf(options.imagerySet) < 0) {
      console.warn('Dynamic styles will only work with these imagerySet choices: ' + DYNAMIC_IMAGERY_SETS.join(', '))
    }

    var metaDataUrl = L.Util.template(L.TileLayer.Bing.METADATA_URL, {
      bingMapsKey: this.options.bingMapsKey,
      imagerySet: this.options.imagerySet,
      culture: this.options.culture
    })

    this._imageryProviders = []
    this._attributions = []

    // Keep a reference to the promise so we can use it later
    this._fetch = fetchJsonp(metaDataUrl, {jsonpCallback: 'jsonp'})
      .then(function (response) {
        return response.json()
      })
      .then(this._metaDataOnLoad.bind(this))
      .catch(console.error.bind(console))

    // for https://github.com/Leaflet/Leaflet/issues/137
    if (!L.Browser.android) {
      this.on('tileunload', this._onTileRemove)
    }
  },

  createTile: function (coords, done) {
    var tile = document.createElement('img')

    L.DomEvent.on(tile, 'load', L.bind(this._tileOnLoad, this, done, tile))
    L.DomEvent.on(tile, 'error', L.bind(this._tileOnError, this, done, tile))

    if (this.options.crossOrigin) {
      tile.crossOrigin = ''
    }

    /*
     Alt tag is set to empty string to keep screen readers from reading URL and for compliance reasons
     http://www.w3.org/TR/WCAG20-TECHS/H67
    */
    tile.alt = ''

    // Don't create closure if we don't have to
    if (this._url) {
      tile.src = this.getTileUrl(coords)
    } else {
      this._fetch.then(function () {
        tile.src = this.getTileUrl(coords)
      }.bind(this)).catch(function (e) {
        console.error(e)
        done(e)
      })
    }

    return tile
  },

  getTileUrl: function (coords) {
    var quadkey = toQuadKey(coords.x, coords.y, coords.z)
    var url = L.Util.template(this._url, {
      quadkey: quadkey,
      subdomain: this._getSubdomain(coords),
      culture: this.options.culture
    })
    if (typeof this.options.style === 'string') {
      url += '&st=' + this.options.style
    }
    return url
  },

  // Update the attribution control every time the map is moved
  onAdd: function (map) {
    map.on('moveend', this._updateAttribution, this)
    L.TileLayer.prototype.onAdd.call(this, map)
    this._attributions.forEach(function (attribution) {
      map.attributionControl.addAttribution(attribution)
    })
  },

  // Clean up events and remove attributions from attribution control
  onRemove: function (map) {
    map.off('moveend', this._updateAttribution, this)
    this._attributions.forEach(function (attribution) {
      map.attributionControl.removeAttribution(attribution)
    })
    L.TileLayer.prototype.onRemove.call(this, map)
  },

  /**
   * Get the [Bing Imagery metadata](https://msdn.microsoft.com/en-us/library/ff701712.aspx)
   * for a specific [`LatLng`](http://leafletjs.com/reference.html#latlng)
   * and zoom level. If either `latlng` or `zoom` is omitted and the layer is attached
   * to a map, the map center and current map zoom are used.
   * @param {L.LatLng} latlng
   * @param {Number} zoom
   * @return {Promise} Resolves to the JSON metadata
   */
  getMetaData: function (latlng, zoom) {
    if (!this._map && (!latlng || !zoom)) {
      return Promise.reject(new Error('If layer is not attached to map, you must provide LatLng and zoom'))
    }
    latlng = latlng || this._map.getCenter()
    zoom = zoom || this._map.getZoom()
    var PointMetaDataUrl = L.Util.template(L.TileLayer.Bing.POINT_METADATA_URL, {
      bingMapsKey: this.options.bingMapsKey,
      imagerySet: this.options.imagerySet,
      z: zoom,
      lat: latlng.lat,
      lng: latlng.lng
    })
    return fetchJsonp(PointMetaDataUrl, {jsonpCallback: 'jsonp'})
      .then(function (response) {
        return response.json()
      })
      .catch(console.error.bind(console))
  },

  _metaDataOnLoad: function (metaData) {
    if (metaData.statusCode !== 200) {
      throw new Error('Bing Imagery Metadata error: \n' + JSON.stringify(metaData, null, '  '))
    }
    var resource = metaData.resourceSets[0].resources[0]
    this._url = resource.imageUrl
    this._imageryProviders = resource.imageryProviders || []
    this.options.subdomains = resource.imageUrlSubdomains
    this._updateAttribution()
    return Promise.resolve()
  },

  /**
   * Update the attribution control of the map with the provider attributions
   * within the current map bounds
   */
  _updateAttribution: function () {
    var map = this._map
    if (!map || !map.attributionControl) return
    var zoom = map.getZoom()
    var bbox = toBingBBox(map.getBounds().toBBoxString())
    this._fetch.then(function () {
      var newAttributions = this._getAttributions(bbox, zoom)
      var prevAttributions = this._attributions
      // Add any new provider attributions in the current area to the attribution control
      newAttributions.forEach(function (attr) {
        if (prevAttributions.indexOf(attr) > -1) return
        map.attributionControl.addAttribution(attr)
      })
      // Remove any attributions that are no longer in the current area from the attribution control
      prevAttributions.filter(function (attr) {
        if (newAttributions.indexOf(attr) > -1) return
        map.attributionControl.removeAttribution(attr)
      })
      this._attributions = newAttributions
    }.bind(this))
  },

  /**
   * Returns an array of attributions for given bbox and zoom
   * @private
   * @param {Array} bbox [west, south, east, north]
   * @param {Number} zoom
   * @return {Array} Array of attribution strings for each provider
   */
  _getAttributions: function (bbox, zoom) {
    return this._imageryProviders.reduce(function (attributions, provider) {
      for (var i = 0; i < provider.coverageAreas.length; i++) {
        if (bboxIntersect(bbox, provider.coverageAreas[i].bbox) &&
          zoom >= provider.coverageAreas[i].zoomMin &&
          zoom <= provider.coverageAreas[i].zoomMax) {
          attributions.push(provider.attribution)
          return attributions
        }
      }
      return attributions
    }, [])
  }
})

L.tileLayer.bing = function (options) {
  return new L.TileLayer.Bing(options)
}

module.exports = L.TileLayer.Bing

}).call(this,typeof global !== "undefined" ? global : typeof self !== "undefined" ? self : typeof window !== "undefined" ? window : {})
},{"bbox-intersect":2,"fetch-jsonp":3}],2:[function(require,module,exports){
module.exports = function(bbox1, bbox2){
  if(!(
      bbox1[0] > bbox2[2] ||
      bbox1[2] < bbox2[0] ||
      bbox1[3] < bbox2[1] ||
      bbox1[1] > bbox2[3]
    )){
      return true;
  } else {
    return false;
  }
}
},{}],3:[function(require,module,exports){
(function (global, factory) {
  if (typeof define === 'function' && define.amd) {
    define(['exports', 'module'], factory);
  } else if (typeof exports !== 'undefined' && typeof module !== 'undefined') {
    factory(exports, module);
  } else {
    var mod = {
      exports: {}
    };
    factory(mod.exports, mod);
    global.fetchJsonp = mod.exports;
  }
})(this, function (exports, module) {
  'use strict';

  var defaultOptions = {
    timeout: 5000,
    jsonpCallback: 'callback',
    jsonpCallbackFunction: null
  };

  function generateCallbackFunction() {
    return 'jsonp_' + Date.now() + '_' + Math.ceil(Math.random() * 100000);
  }

  // Known issue: Will throw 'Uncaught ReferenceError: callback_*** is not defined' error if request timeout
  function clearFunction(functionName) {
    // IE8 throws an exception when you try to delete a property on window
    // http://stackoverflow.com/a/1824228/751089
    try {
      delete window[functionName];
    } catch (e) {
      window[functionName] = undefined;
    }
  }

  function removeScript(scriptId) {
    var script = document.getElementById(scriptId);
    document.getElementsByTagName('head')[0].removeChild(script);
  }

  var fetchJsonp = function fetchJsonp(url) {
    var options = arguments[1] === undefined ? {} : arguments[1];

    var timeout = options.timeout != null ? options.timeout : defaultOptions.timeout;
    var jsonpCallback = options.jsonpCallback != null ? options.jsonpCallback : defaultOptions.jsonpCallback;

    var timeoutId = undefined;

    return new Promise(function (resolve, reject) {
      var callbackFunction = options.jsonpCallbackFunction || generateCallbackFunction();

      window[callbackFunction] = function (response) {
        resolve({
          ok: true,
          // keep consistent with fetch API
          json: function json() {
            return Promise.resolve(response);
          }
        });

        if (timeoutId) clearTimeout(timeoutId);

        removeScript(jsonpCallback + '_' + callbackFunction);

        clearFunction(callbackFunction);
      };

      // Check if the user set their own params, and if not add a ? to start a list of params
      url += url.indexOf('?') === -1 ? '?' : '&';

      var jsonpScript = document.createElement('script');
      jsonpScript.setAttribute('src', url + jsonpCallback + '=' + callbackFunction);
      jsonpScript.id = jsonpCallback + '_' + callbackFunction;
      document.getElementsByTagName('head')[0].appendChild(jsonpScript);

      timeoutId = setTimeout(function () {
        reject(new Error('JSONP request to ' + url + ' timed out'));

        clearFunction(callbackFunction);
        removeScript(jsonpCallback + '_' + callbackFunction);
      }, timeout);
    });
  };

  // export as global function
  /*
  let local;
  if (typeof global !== 'undefined') {
    local = global;
  } else if (typeof self !== 'undefined') {
    local = self;
  } else {
    try {
      local = Function('return this')();
    } catch (e) {
      throw new Error('polyfill failed because global object is unavailable in this environment');
    }
  }
  
  local.fetchJsonp = fetchJsonp;
  */

  module.exports = fetchJsonp;
});
},{}]},{},[1]);

(function(factory, window) {
  // define an AMD module that relies on 'leaflet'
  if (typeof define === "function" && define.amd) {
    define(["leaflet"], factory);

    // define a Common JS module that relies on 'leaflet'
  } else if (typeof exports === "object") {
    module.exports = factory(require("leaflet"));
  }

  // attach your plugin to the global 'L' variable
  if (typeof window !== "undefined" && window.L) {
    factory(L);
  }
})(function(L) {
  L.Measure = {
    linearMeasurement: "Distance measurement",
    areaMeasurement: "Area measurement",
    start: "Start",
    meter: "m",
    kilometer: "km",
    squareMeter: "m²",
    squareKilometers: "km²",
  };

  L.Control.Measure = L.Control.extend({
    options: {
      position: "topleft",
      title: "Measurement",
      collapsed: true,
      color: "#FF0080",
    },
    initialize: function(options) {
      L.Util.setOptions(this, options);
    },
    onAdd: function(map) {
      this._map = map;
      this._container || this._initLayout();
      return this._container;
    },
    _buildContainer: function() {
      this._container = L.DomUtil.create("div", "leaflet-control-measure leaflet-bar leaflet-control");

      this._contents = L.DomUtil.create("div", "leaflet-measure-contents", this._container);

      this._link = L.DomUtil.create("a", "leaflet-measure-toggle", this._container);
      this._link.title = this.options.title || "Measurement";
      this._link.href = "#";

      if (this.options.title) {
        var title = L.DomUtil.create("h3", "", this._contents);
        title.innerText = this.options.title;
      }

      this._buildItems();
    },
    _buildItems: function() {
      var ele_ul = L.DomUtil.create("ul", "leaflet-measure-actions", this._contents);
      var ele_li = L.DomUtil.create("li", "leaflet-measure-action", ele_ul);
      var ele_link_line = L.DomUtil.create("a", "start", ele_li);
      ele_link_line.innerText = L.Measure.linearMeasurement;
      ele_link_line.href = "#";
      L.DomEvent.disableClickPropagation(ele_link_line);
      L.DomEvent.on(ele_link_line, "click", this._enableMeasureLine, this);

      ele_li = L.DomUtil.create("li", "leaflet-measure-action", ele_ul);
      var ele_link_area = L.DomUtil.create("a", "leaflet-measure-action start", ele_li);
      ele_link_area.innerText = L.Measure.areaMeasurement;
      ele_link_area.href = "#";
      L.DomEvent.disableClickPropagation(ele_link_area);
      L.DomEvent.on(ele_link_area, "click", this._enableMeasureArea, this);
    },
    _initLayout: function() {
      this._buildContainer();
      L.DomEvent.disableClickPropagation(this._container);
      L.DomEvent.disableScrollPropagation(this._container);
      if (this.options.collapsed) {
        L.DomEvent.on(
          this._container, {
            mouseenter: this._expand,
            mouseleave: this._collapse,
          },
          this
        );
      } else {
        this._expand();
      }
    },
    _enableMeasureLine: function(ev) {
      L.DomEvent.stopPropagation(ev);
      this._measureHandler = new L.MeasureAction(this._map, {
        model: "distance",
        color: this.options.color,
      });
      this._measureHandler.enable();
    },
    _enableMeasureArea: function(ev) {
      L.DomEvent.stopPropagation(ev);
      this._measureHandler = new L.MeasureAction(this._map, {
        model: "area",
        color: this.options.color,
      });
      this._measureHandler.enable();
    },
    _expand: function() {
      this._link.style.display = "none";
      L.DomUtil.addClass(this._container, "leaflet-measure-expanded");
      return this;
    },
    _collapse: function() {
      this._link.style.display = "block";
      L.DomUtil.removeClass(this._container, "leaflet-measure-expanded");
      return this;
    },
  });

  L.control.measure = L.control.Measure = function(options) {
    return new L.Control.Measure(options);
  };

  L.MeasureLable = L.Layer.extend({
    options: {
      offset: new L.Point(0, 30),
      latlng: null,
      content: "",
      className: "",
    },
    initialize: function(options) {
      L.Util.setOptions(this, options);
    },
    onAdd: function(map) {
      this._map = map;
      this._container || this._initLayout();
      map._panes.popupPane.appendChild(this._container);
      map.on("viewreset", this._updatePosition, this);
      if (L.Browser.any3d) {
        map.on("zoomanim", this._zoomAnimation, this);
      }
      this._update();
    },
    addTo: function(map) {
      map.addLayer(this);
      return this;
    },
    onRemove: function(map) {
      map._panes.popupPane.removeChild(this._container);
      map.off({
          viewreset: this._updatePosition,
          zoomanim: this._zoomAnimation,
        },
        this
      );
      this._map = null;
    },
    setLatLng: function(latlng) {
      this.options.latlng = L.latLng(latlng);
      this._updatePosition();
      return this;
    },
    setContent: function(content) {
      this.options.content = content;
      this._updateContent();
      return this;
    },
    _initLayout: function() {
      this._container = L.DomUtil.create("div", this.options.className);
      this._contentNode = L.DomUtil.create("div", "content", this._container);
    },
    _update: function() {
      this._map && (this._updateContent(), this._updatePosition());
    },
    _updateContent: function() {
      if (this.options.content) {
        if (typeof this.options.content == "string") {
          this._contentNode.innerHTML = this.options.content;
        } else {
          this._contentNode.innerHTML = "";
          this._contentNode.appendChild(this.options.content);
        }
      }
    },
    _updatePosition: function() {
      var point = this._map.latLngToLayerPoint(this.options.latlng),
        is3D = L.Browser.any3d,
        offset = this.options.offset;
      is3D && L.DomUtil.setPosition(this._container, point);
      this._containerBottom = -offset.y - (is3D ? 0 : point.y);
      this._containerLeft = offset.x + (is3D ? 0 : point.x);
      this._container.style.bottom = this._containerBottom + "px";
      this._container.style.left = this._containerLeft + "px";
    },
    _zoomAnimation: function(a) {
      a = this._map._latLngToNewLayerPoint(this.options.latlng, a.zoom, a.center);
      L.DomUtil.setPosition(this._container, a);
    },
    enableClose: function() {
      this._closeButton = L.DomUtil.create("span", "close", this._container);
      this._closeButton.innerHTML =
        '<svg class="icon" viewBox="0 0 40 40"><path stroke="#FF0000" stroke-width="3" d="M 10,10 L 30,30 M 30,10 L 10,30" /></svg>';
      return this._closeButton;
    },
  });

  L.MeasureAction = L.Handler.extend({
    options: {
      color: "#FF0080",
      model: "distance", // area or distance
    },

    initialize: function(map, options) {
      this._map = map;
      this._map._measureHandler = this;
      L.Util.setOptions(this, options);
    },
    setModel: function(model) {
      this.options.model = model;
      return this;
    },
    addHooks: function() {
      this._activeMeasure();
    },
    removeHooks: function() {},
    _activeMeasure: function() {
      this._map._measureHandler._measurementStarted && this._map._measureHandler._finishMeasure();
      this._measurementStarted ? this._finishMeasure() : this._enableMeasure();
    },
    _onMouseClick: function(event) {
      var latlng = event.latlng || this._map.mouseEventToLatLng(event);
      if (this._lastPoint && latlng.equals(this._lastPoint)) {
        return;
      }
      if (this._trail.points.length > 0) {
        var points = this._trail.points;
        points.push(latlng);
        var length = points.length;
        this._totalDistance += this._getDistance(points[length - 2], points[length - 1]);
        this._addMeasurePoint(latlng);
        this._addMarker(latlng);
        if (this.options.model !== "area") {
          this._addLable(latlng, this._getDistanceString(this._totalDistance), "leaflet-measure-lable");
        }
      } else {
        this._totalDistance = 0;
        this._addMeasurePoint(latlng);
        this._addMarker(latlng);
        if (this.options.model !== "area") {
          this._addLable(latlng, L.Measure.start, "leaflet-measure-lable");
        }
        this._trail.points.push(latlng);
      }
      this._lastPoint = latlng;
      this._startMove = false;
    },
    _onMouseMove: function(event) {
      var latlng = event.latlng;
      if (this._trail.points.length > 0) {
        if (this._startMove) {
          this._directPath.setLatLngs(this._trail.points.concat(latlng));
        } else {
          this._directPath.setLatLngs([latlng]);
          this._startMove = true;
        }
      }
    },
    _enableMeasure: function() {
      this._trail = {
        overlays: [],
        points: [],
      };
      var map = this._map;
      L.DomUtil.addClass(map._container, "leaflet-measure-map");
      map.contextMenu && map.contextMenu.disable();
      this._measurementStarted = true;
      map.on("click", this._onMouseClick, this);
      map.on("dblclick contextmenu", this._finishMeasure, this);
      map.doubleClickZoom.disable();
      map.on("mousemove", this._onMouseMove, this);
    },
    _disableMeasure: function() {
      var map = this._map;
      L.DomUtil.removeClass(map.getContainer(), "leaflet-measure-map");
      map.contextMenu && map.contextMenu.enable();
      map.off("click", this._onMouseClick, this);
      map.off("dblclick contextmenu", this._finishMeasure, this);
      map.off("mousemove", this._onMouseMove, this);
      map.doubleClickZoom.enable();
      this._measurementStarted = this._startMove = false;
      this.disable();
    },
    _finishMeasure: function(event) {
      if (this._trail.points.length > 0) {
        if (this._trail.points.length > 1) {
          if (!event || event.type === "contextmenu") {
            this._directPath.setLatLngs(this._trail.points);
          }
          if (this.options.model === "area") {
            this._addLable(
              this._lastPoint,
              this._getAreaString(this._trail.points),
              "leaflet-measure-lable",
              true
            );
          } else {
            this._addLable(
              this._lastPoint,
              this._getDistanceString(this._totalDistance),
              "leaflet-measure-lable",
              true
            );
          }
          this._directPath && this._map.removeLayer(this._directPath);
        } else {
          this._clearOverlay.call(this);
        }
      }
      this._disableMeasure();
    },
    _resetDirectPath: function(latlng) {
      if (!this._directPath) {
        if (this.options.model === "area") {
          this._directPath = new L.Polygon([latlng], {
            weight: 2,
            color: this.options.color,
            dashArray: [5, 5],
            fillOpacity: 0,
            interactive: false,
          });
        } else {
          this._directPath = new L.Polyline([latlng], {
            weight: 2,
            color: this.options.color,
            dashArray: [5, 5],
            interactive: false,
          });
        }
        this._map.addLayer(this._directPath);
        this._trail.overlays.push(this._directPath);
      } else {
        this._directPath.addLatLng(latlng);
      }
    },
    _addMeasurePoint: function(latlng) {
      if (!this._measurePath) {
        if (this.options.model === "area") {
          this._measurePath = new L.Polygon([latlng], {
            weight: 2,
            color: this.options.color,
            fillColor: this.options.color,
            fillOpacity: 0.5,
            interactive: false,
          });
        } else {
          this._measurePath = new L.Polyline([latlng], {
            weight: 2,
            color: this.options.color,
            interactive: false,
          });
        }
        this._map.addLayer(this._measurePath);
        this._trail.overlays.push(this._measurePath);
      } else {
        this._measurePath.addLatLng(latlng);
      }
      this._resetDirectPath(latlng);
    },
    _addMarker: function(latLng) {
      var marker = new L.CircleMarker(latLng, {
        color: this.options.color,
        opacity: 1,
        weight: 1,
        fillColor: "#FFFFFF",
        fill: true,
        fillOpacity: 1,
        radius: 3,
        interactive: false,
      });
      this._map.addLayer(marker);
      this._trail.overlays.push(marker);
    },
    _addLable: function(latlng, content, className, ended) {
      var lable = new L.MeasureLable({
        latlng: latlng,
        content: content,
        className: className,
      });
      this._map.addLayer(lable);
      this._trail.overlays.push(lable);
      if (ended) {
        var closeButton = lable.enableClose();
        L.DomEvent.on(closeButton, "click", this._clearOverlay, this);
      }
    },
    _clearOverlay: function() {
      var i = 0,
        overlays = this._trail.overlays,
        length;
      for (length = overlays.length; i < length; i++) {
        this._map.removeLayer(overlays[i]);
      }
    },
    toRadians: function(deg) {
      return deg * (Math.PI / 180);
    },
    square: function(x) {
      return Math.pow(x, 2);
    },
    _getDistanceString: function(distance) {
      return distance < 1e3 ?
        this._numberFormat(distance, 0) + " " + L.Measure.meter :
        this._numberFormat(distance / 1e3, 2) + " " + L.Measure.kilometer;
    },

    _getDistance: function(latlng1, latlng2) {
      var earthRadius = 6378137; // radius of the earth in meter
      var lat1 = this.toRadians(latlng1.lat);
      var lat2 = this.toRadians(latlng2.lat);
      var lat_dif = lat2 - lat1;
      var lng_dif = this.toRadians(latlng2.lng - latlng1.lng);
      var a =
        this.square(Math.sin(lat_dif / 2)) +
        Math.cos(lat1) * Math.cos(lat2) * this.square(Math.sin(lng_dif / 2));
      return 2 * earthRadius * Math.asin(Math.sqrt(a));
    },
    _getAreaString: function(points) {
      var a = Math.round(this._getArea(points));
      return a < 1e6 ?
        this._numberFormat(a, 0) + " " + L.Measure.squareMeter :
        this._numberFormat(a / 1e6, 2) + " " + L.Measure.squareKilometers;
    },
    _getArea: function(points) {
      var earthRadius = 6378137;
      var area = 0;
      var len = points.length;
      var x1 = points[len - 1].lng;
      var y1 = points[len - 1].lat;
      for (var i = 0; i < len; i++) {
        var x2 = points[i].lng;
        var y2 = points[i].lat;
        area += this.toRadians(x2 - x1) * (2 + Math.sin(this.toRadians(y1)) + Math.sin(this.toRadians(y2)));
        x1 = x2;
        y1 = y2;
      }
      return Math.abs((area * earthRadius * earthRadius) / 2.0);
    },
    _numberFormat: function(number, decimals = 2) {
      var thousandsSep = ",";
      var sign = number < 0 ? "-" : "";
      var num = Math.abs(+number || 0);
      var intPart = parseInt(num.toFixed(decimals), 10) + "";
      var j = intPart.length > 3 ? intPart.length % 3 : 0;

      return [
        sign,
        j ? intPart.substr(0, j) + thousandsSep : "",
        intPart.substr(j).replace(/(\d{3})(?=\d)/g, "$1" + thousandsSep),
        decimals ?
        "." +
        Math.abs(num - intPart)
        .toFixed(decimals)
        .slice(2) :
        "",
      ].join("");
    },
  });

  L.measureAction = function(map, options) {
    return new L.MeasureAction(map, options);
  };
}, window);
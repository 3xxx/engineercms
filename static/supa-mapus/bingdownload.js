/**
 * Created by 3xxx on 2023/01/18.
 * 提供了微软bing地图框选按级下载
 */
;
(function(global) {
  "use strict";

  var BingDownload = function(el) {
    return new BingDownload.prototype.init(el)
  };

  // var tileZoom = 12;
  // var rootTileDir = "tiles_cache";

  // var lat_min = 22.99885; //右下角
  // var lat_max = 23.16056; //左上角
  // var lon_min = 113.29101; //左上角
  // var lon_max = 113.64257; //右下角
  // longitude是经度y，latitude是纬度x
  // x
  // x
  // x
  // x
  // o y y y y y y y y

  // const lat = 23.125178;
  // const lon = 113.280637;
  // var latlng = {};
  // latlng.lat = 23.125178;
  // latlng.lng = 113.280637;

  // var BING_KEY = 'AhJFIXaWTfryahNRyQFR6gKLDqmxB34-FGH6VB5VYas4YJWgO1_bQT6WLWYWQD0g';

  // var METADATA_URL = 'https://dev.virtualearth.net/REST/v1/Imagery/Metadata/{imagerySet}?key={bingMapsKey}&include=ImageryProviders&uriScheme=https&c={culture}';
  // var POINT_METADATA_URL = 'https://dev.virtualearth.net/REST/v1/Imagery/Metadata/{imagerySet}/{lat},{lng}?zl={z}&key={bingMapsKey}&uriScheme=https&c={culture}';

  // var metaDataUrl = L.Util.template(METADATA_URL, {
  //   bingMapsKey: BING_KEY,
  //   imagerySet: 'Aerial',
  //   culture: 'en-US'
  // });
  // console.log(metaDataUrl)

  // var _url = "";
  // var _imageryProviders = "";
  // var subdomains = "";
  // var url = {};

  // var defaultOptions = {
  //   timeout: 5000,
  //   jsonpCallback: 'callback',
  //   jsonpCallbackFunction: null
  // };

  const EarthRadius = 6378137;
  const MinLatitude = -85.05112878;
  const MaxLatitude = 85.05112878;
  const MinLongitude = -180;
  const MaxLongitude = 180;
  // var urlList = new Array(); //文件列表

  // var coords = LatLongToPixelXY(latlng.lat, latlng.lng, 12);
  // // console.log(coords)
  // // x: 3336
  // // y:1777
  // var tilexy = PixelXYToTileXY(coords.X, coords.Y);
  // // console.log(tilexy)
  // // var quadkey = TileXYToQuadKey(tilexy.X,tilexy.Y,12)
  // var quadkey = toQuadKey(tilexy.X, tilexy.Y, 12);
  // // console.log(quadkey)

  // // # 左上为原点
  // var tilePixelMax = LatLongToPixelXY(lat_max, lon_max, tileZoom);
  // var tilePixelMin = LatLongToPixelXY(lat_min, lon_min, tileZoom);

  BingDownload.prototype = {
    //输出二维数组元素的值
    // for (i = 0; i < aeg.length; i++) {
    //   for (j = 0; j < aeg[i].length; j++) {
    //     document.write(aeg[i][j]);
    //     document.write('　');
    //   }
    //   document.write('<br>');
    init: function(el) {
      this.el = typeof el === "string" ? document.querySelector(el) : el;
    },
    // _metaDataOnLoad: function(metaData) {
    //   if (metaData.statusCode !== 200) {
    //     throw new Error('Bing Imagery Metadata error: \n' + JSON.stringify(metaData, null, '  '))
    //   }
    //   var resource = metaData.resourceSets[0].resources[0]
    //   url._url = resource.imageUrl
    //   url._imageryProviders = resource.imageryProviders || []
    //   url.subdomains = resource.imageUrlSubdomains
    //   // this._updateAttribution()
    //   return url
    //   // return Promise.resolve()
    // },

    // downloadMap: async function(tilePixelMin,tilePixelMax,url) {
    //   var i = 0
    //   for (var x = tilePixelMin.X; x <= tilePixelMax.X; x = x + 256) {
    //     urlList[i] = new Array();
    //     var j = 0
    //     for (var y = tilePixelMax.Y; y <= tilePixelMin.Y; y = y + 256) {
    //       // coords = LatLongToPixelXY(latlng.lat, latlng.lng, 12)
    //       var tilexy = this.PixelXYToTileXY(x, y)
    //       var quadkey = this.toQuadKey(tilexy.X, tilexy.Y, 12)
    //       // Promise.race([_fetch]).then(function(data) {
    //       // console.log(data._url); //https://ecn.{subdomain}.tiles.virtualearth.net/tiles/a{quadkey}.jpeg?g=13239
    //       // console.log(data.subdomains[parseInt(Math.random() * 4)]);
    //       var tempurl = url._url.replace("{subdomain}", url.subdomains[parseInt(Math.random() * 4)])
    //       // console.log(url)
    //       tempurl = tempurl.replace("{quadkey}", quadkey)
    //       // https://ecn.t0.tiles.virtualearth.net/tiles/a132122220130.jpeg?g=13239
    //       // savePictureZip(tempurl, "a" + quadkey + ".jpeg")
    //       // console.log(x)
    //       // console.log(y)

    //       await this.savePictureZip(tempurl, x + " " + y + ".jpeg")

    //       // urlList.push(x + " " + y + ".jpeg")
    //       // urlList.push(tempurl)
    //       // urlList[i][j] = tempurl;
    //       urlList[i][j] = x + "-" + y + ".jpeg"
    //       // })
    //       // console.log(new Date())
    //       j++
    //       // console.log(j)
    //     }
    //     i++
    //     // console.log(i)
    //     // mergeImgs(urlList).then(base64 => {
    //     //   const imgDom = document.createElement('img')
    //     //   imgDom.src = base64
    //     //   document.body.appendChild(imgDom)
    //     // })
    //   }
    // },

    // savePictureZip: async function(url, name) {
    //   // # 强制睡一会，防止bing服务器限制
    //   var sleepTime = Math.random() * 3
    //   // console.log(sleepTime * 1000)
    //   return new Promise(function(resolve, reject) {
    //     setTimeout(function() {
    //       if (true) {
    //         console.log('请求中...')
    //         console.log(url, name, console.log(new Date()))
    //         fetch(url)
    //           .then(res => res.blob())
    //           .then(blob => {
    //             var a = document.createElement("a");
    //             a.href = URL.createObjectURL(blob);
    //             a.download = name;
    //             a.click();
    //           })
    //         resolve('resolve return')
    //       } else {
    //         reject('reject return')
    //       }
    //     }, sleepTime * 1000)
    //   })

    //   // try {
    //   //   var xhr = new XMLHttpRequest();
    //   //   xhr.open('get', url, true);
    //   //   xhr.responseType = 'blob'; //二进制对象，binary
    //   //   xhr.onload = function() {
    //   //     // fileCount++;
    //   //     // $("#waitInfo").html("正在下载：" + url);
    //   //     var blob = xhr.response; //response 属性返回响应的正文，取决于 responseType 属性。
    //   //     if (blob.type.indexOf("image") > -1) {
    //   //       var a = document.createElement("a");
    //   //       var file = new Blob([blob], {
    //   //         type: "image/jpeg"
    //   //       });
    //   //       a.href = URL.createObjectURL(blob);
    //   //       a.download = name;
    //   //       a.click();
    //   //     }
    //   //   };
    //   // } catch (e) {}
    // },

    // handleDownLoadFile: function(response, type, fileName) {
    //   let blob = new Blob([response], {
    //     type: type + ';charset=utf-8'
    //   });
    //   let src = window.URL.createObjectURL(blob);
    //   if (src) {
    //     let link = document.createElement('a');
    //     link.style.display = 'none';
    //     link.href = src;
    //     link.setAttribute('download', fileName);
    //     document.body.appendChild(link);
    //     link.click();
    //     document.body.removeChild(link); //下载完成移除元素
    //     window.URL.revokeObjectURL(src); //释放掉blob对象
    //   }
    // },

    /**
     * 合并多张图片，返回新的图片
     * @param {Array} list 图片url数组
     * @param {Number} cwith 画布宽度 默认256
     * @param {Number} cheight 画布高度 默认256
     */
    // mergeImgs: function(list, cwith = 256, cheight = 256) {
    //   return new Promise((resolve, reject) => {
    //     const baseList = []

    //     // 创建 canvas 节点并初始化
    //     const canvas = document.createElement('canvas')
    //     canvas.width = cwith
    //     canvas.height = cheight * list.length
    //     const context = canvas.getContext('2d')

    //     list.map((item, index) => {
    //       const img = new Image()
    //       // 跨域crossOrigin属性必须在src之前
    //       img.crossOrigin = 'Anonymous'
    //       img.src = item //"C:/Users/Administrator/Downloads/" + 

    //       img.onload = () => {
    //         context.drawImage(img, 0, cheight * index, cwith, cheight)
    //         const base64 = canvas.toDataURL('image/jpeg')
    //         baseList.push(base64)
    //         if (baseList[list.length - 1]) {
    //           console.log(baseList)
    //           // 返回新的图片
    //           resolve(baseList[list.length - 1])
    //         }
    //       }
    //       var sleepTime = Math.random() * 30
    //       sleep(sleepTime * 1000)
    //     })
    //   })
    // },

    /// <summary>  
    /// Clips a number to the specified minimum and maximum values.  
    /// </summary>  
    /// <param name="n">The number to clip.</param>  
    /// <param name="minValue">Minimum allowable value.</param>  
    /// <param name="maxValue">Maximum allowable value.</param>  
    /// <returns>The clipped value.</returns>  
    Clip: function(n, minValue, maxValue) {
      return Math.min(Math.max(n, minValue), maxValue);
    },

    /// <summary>  
    /// Determines the map width and height (in pixels) at a specified level  
    /// of detail.  
    /// </summary>  
    /// <param name="levelOfDetail">Level of detail, from 1 (lowest detail)  
    /// to 23 (highest detail).</param>  
    /// <returns>The map width and height in pixels.</returns>  
    MapSize: function(levelOfDetail) {
      return 256 << levelOfDetail;
    },

    /// <summary>  
    /// Determines the ground resolution (in meters per pixel) at a specified  
    /// latitude and level of detail.  
    /// </summary>  
    /// <param name="latitude">Latitude (in degrees) at which to measure the  
    /// ground resolution.</param>  
    /// <param name="levelOfDetail">Level of detail, from 1 (lowest detail)  
    /// to 23 (highest detail).</param>  
    /// <returns>The ground resolution, in meters per pixel.</returns>  
    GroundResolution: function(latitude, levelOfDetail) {
      latitude = Clip(latitude, MinLatitude, MaxLatitude);
      return Math.Cos(latitude * Math.PI / 180) * 2 * Math.PI * EarthRadius / MapSize(levelOfDetail);
    },

    /// <summary>  
    /// Determines the map scale at a specified latitude, level of detail,  
    /// and screen resolution.  
    /// </summary>  
    /// <param name="latitude">Latitude (in degrees) at which to measure the  
    /// map scale.</param>  
    /// <param name="levelOfDetail">Level of detail, from 1 (lowest detail)  
    /// to 23 (highest detail).</param>  
    /// <param name="screenDpi">Resolution of the screen, in dots per inch.</param>  
    /// <returns>The map scale, expressed as the denominator N of the ratio 1 : N.</returns>  
    MapScale: function(latitude, levelOfDetail, screenDpi) {
      return GroundResolution(latitude, levelOfDetail) * screenDpi / 0.0254;
    },

    /// <summary>  
    /// Converts a point from latitude/longitude WGS-84 coordinates (in degrees)  
    /// into pixel XY coordinates at a specified level of detail.  
    /// </summary>  
    /// <param name="latitude">Latitude of the point, in degrees.</param>  
    /// <param name="longitude">Longitude of the point, in degrees.</param>  
    /// <param name="levelOfDetail">Level of detail, from 1 (lowest detail)  
    /// to 23 (highest detail).</param>  
    /// <param name="pixelX">Output parameter receiving the X coordinate in pixels.</param>  
    /// <param name="pixelY">Output parameter receiving the Y coordinate in pixels.</param>  
    LatLongToPixelXY: function(latitude, longitude, levelOfDetail) {
      latitude = this.Clip(latitude, MinLatitude, MaxLatitude);
      longitude = this.Clip(longitude, MinLongitude, MaxLongitude);

      var x = (longitude + 180) / 360;
      var sinLatitude = Math.sin(latitude * Math.PI / 180);
      var y = 0.5 - Math.log((1 + sinLatitude) / (1 - sinLatitude)) / (4 * Math.PI);

      var mapSize = this.MapSize(levelOfDetail);
      var pixel = {}
      pixel.X = this.Clip(x * mapSize + 0.5, 0, mapSize - 1);
      pixel.Y = this.Clip(y * mapSize + 0.5, 0, mapSize - 1);

      return pixel
    },

    /// <summary>  
    /// Converts a pixel from pixel XY coordinates at a specified level of detail  
    /// into latitude/longitude WGS-84 coordinates (in degrees).  
    /// </summary>  
    /// <param name="pixelX">X coordinate of the point, in pixels.</param>  
    /// <param name="pixelY">Y coordinates of the point, in pixels.</param>  
    /// <param name="levelOfDetail">Level of detail, from 1 (lowest detail)  
    /// to 23 (highest detail).</param>  
    /// <param name="latitude">Output parameter receiving the latitude in degrees.</param>  
    /// <param name="longitude">Output parameter receiving the longitude in degrees.</param>  
    PixelXYToLatLong: function(pixelX, pixelY, levelOfDetail) {
      var mapSize = MapSize(levelOfDetail);
      var x = (Clip(pixelX, 0, mapSize - 1) / mapSize) - 0.5;
      var y = 0.5 - (Clip(pixelY, 0, mapSize - 1) / mapSize);

      var latitude = 90 - 360 * Math.Atan(Math.Exp(-y * 2 * Math.PI)) / Math.PI;
      var longitude = 360 * x;
      return latitude + "," + longitude
    },

    /// <summary>  
    /// Converts pixel XY coordinates into tile XY coordinates of the tile containing  
    /// the specified pixel.  
    /// </summary>  
    /// <param name="pixelX">Pixel X coordinate.</param>  
    /// <param name="pixelY">Pixel Y coordinate.</param>  
    /// <param name="tileX">Output parameter receiving the tile X coordinate.</param>  
    /// <param name="tileY">Output parameter receiving the tile Y coordinate.</param>  
    PixelXYToTileXY: function(pixelX, pixelY) {
      var tilexy = {}
      tilexy.X = pixelX / 256;
      tilexy.Y = pixelY / 256;
      return tilexy
    },

    /// <summary>  
    /// Converts tile XY coordinates into pixel XY coordinates of the upper-left pixel  
    /// of the specified tile.  
    /// </summary>  
    /// <param name="tileX">Tile X coordinate.</param>  
    /// <param name="tileY">Tile Y coordinate.</param>  
    /// <param name="pixelX">Output parameter receiving the pixel X coordinate.</param>  
    /// <param name="pixelY">Output parameter receiving the pixel Y coordinate.</param>  
    TileXYToPixelXY: function(tileX, tileY) {
      pixelX = tileX * 256;
      pixelY = tileY * 256;
      return pixelX + "," + pixelY
    },

    /// <summary>  
    /// Converts tile XY coordinates into a QuadKey at a specified level of detail.  
    /// </summary>  
    /// <param name="tileX">Tile X coordinate.</param>  
    /// <param name="tileY">Tile Y coordinate.</param>  
    /// <param name="levelOfDetail">Level of detail, from 1 (lowest detail)  
    /// to 23 (highest detail).</param>  
    /// <returns>A string containing the QuadKey.</returns>  
    // function TileXYToQuadKey(tileX, tileY, levelOfDetail) {
    //   var quadKey = new StringBuilder();
    //   for (i = levelOfDetail; i > 0; i--) {
    //     var digit = '0';
    //     var mask = 1 << (i - 1);
    //     if ((tileX & mask) != 0) {
    //       digit++;
    //     }
    //     if ((tileY & mask) != 0) {
    //       digit++;
    //       digit++;
    //     }
    //     quadKey.append(digit);
    //   }
    //   return quadKey.toString();
    // }

    // StringBuilder: function() {
    //   // console.log(this);
    //   //StringBuilder
    //   //    _stringArray: Array[2]
    //   //    __proto__: StringBuilder
    //   this._stringArray = new Array();

    // },

    // StringBuilder.prototype.append = function(str) {
    //   this._stringArray.push(str);
    // },
    // StringBuilder.prototype.toString = function(joinGap) {
    //   return this._stringArray.join(joinGap);
    // },
    // var stringBuilder = new StringBuilder();
    // stringBuilder.append("hi");
    // stringBuilder.append("haojie");
    // console.log(stringBuilder.toString(" ")); // hi haojie


    /// <summary>  
    /// Converts a QuadKey into tile XY coordinates.  
    /// </summary>  
    /// <param name="quadKey">QuadKey of the tile.</param>  
    /// <param name="tileX">Output parameter receiving the tile X coordinate.</param>  
    /// <param name="tileY">Output parameter receiving the tile Y coordinate.</param>  
    /// <param name="levelOfDetail">Output parameter receiving the level of detail.</param>  
    QuadKeyToTileXY: function(quadKey) {
      var tileX = tileY = 0;
      levelOfDetail = quadKey.Length;
      for (i = levelOfDetail; i > 0; i--) {
        mask = 1 << (i - 1);
        switch (quadKey[levelOfDetail - i]) {
          case '0':
            break;
          case '1':
            tileX |= mask;
            break;
          case '2':
            tileY |= mask;
            break;
          case '3':
            tileX |= mask;
            tileY |= mask;
            break;
          default:
            throw new ArgumentException("Invalid QuadKey digit sequence.");
        }
      }
      return tileX + "," + tileY + "," + levelOfDetail
    },
    // *********************

    // generateCallbackFunction: function() {
    //   return 'jsonp_' + Date.now() + '_' + Math.ceil(Math.random() * 100000);
    // },

    // Known issue: Will throw 'Uncaught ReferenceError: callback_*** is not defined' error if request timeout
    // clearFunction: function(functionName) {
    //   // IE8 throws an exception when you try to delete a property on window
    //   // http://stackoverflow.com/a/1824228/751089
    //   try {
    //     delete window[functionName];
    //   } catch (e) {
    //     window[functionName] = undefined;
    //   }
    // },

    // removeScript: function(scriptId) {
    //   var script = document.getElementById(scriptId);
    //   document.getElementsByTagName('head')[0].removeChild(script);
    // },

    // fetchJsonp: function(url) {
    //   var options = arguments[1] === undefined ? {} : arguments[1];
    //   var timeout = options.timeout != null ? options.timeout : defaultOptions.timeout;;
    //   var jsonpCallback = options.jsonpCallback != null ? options.jsonpCallback : defaultOptions.jsonpCallback;
    //   var timeoutId = undefined;
    //   return new Promise(function(resolve, reject) {
    //     var callbackFunction = options.jsonpCallbackFunction || generateCallbackFunction();

    //     window[callbackFunction] = function(response) {
    //       resolve({
    //         ok: true,
    //         // keep consistent with fetch API
    //         json: function json() {
    //           return Promise.resolve(response);
    //         }
    //       });
    //       if (timeoutId) clearTimeout(timeoutId);
    //       removeScript(jsonpCallback + '_' + callbackFunction);
    //       clearFunction(callbackFunction);
    //     };
    //     // Check if the user set their own params, and if not add a ? to start a list of params
    //     url += url.indexOf('?') === -1 ? '?' : '&';
    //     var jsonpScript = document.createElement('script');
    //     jsonpScript.setAttribute('src', url + jsonpCallback + '=' + callbackFunction);
    //     jsonpScript.id = jsonpCallback + '_' + callbackFunction;
    //     document.getElementsByTagName('head')[0].appendChild(jsonpScript);
    //     timeoutId = setTimeout(function() {
    //       reject(new Error('JSONP request to ' + url + ' timed out'));
    //       clearFunction(callbackFunction);
    //       removeScript(jsonpCallback + '_' + callbackFunction);
    //     }, timeout);
    //   });
    // },

    toQuadKey: function(x, y, z) {
      var index = ''
      for (var i = z; i > 0; i--) {
        var b = 0
        var mask = 1 << (i - 1)
        if ((x & mask) !== 0) b++
        if ((y & mask) !== 0) b += 2
        index += b.toString()
      }
      return index
    },

    // _fetch = fetchJsonp(metaDataUrl, { jsonpCallback: 'jsonp' })
    //   .then(function(response) {
    //     return response.json()
    //   })
    //   .then(this._metaDataOnLoad.bind(this))
    //   .catch(console.error.bind(console));

    // $("#downloadMap").click(function(e) {
    //   downloadMap()
    // });

    // $("#merge_btn").click(function(e) {
    //   mergeImgs(urlList).then(base64 => {
    //     const imgDom = document.createElement('img')
    //     imgDom.src = base64
    //     document.body.appendChild(imgDom)
    //   })
    // });


    // $("#rctanglexport_btn").click(function() {
    //   //绘制矩形
    //   map.pm.enableDraw("Rectangle", {
    //     finishOn: "dblclick",
    //     allowSelfIntersection: false,
    //     tooltips: false
    //   });
    // });

    // map.on('pm:create', e => {
    //   //console.log(e);
    //   geojsonLayer = e.layer;
    //   geojsonLayer.addTo(map);

    //   var northEast = e.layer.getBounds()._northEast;
    //   var southWest = e.layer.getBounds()._southWest;
    //   // console.log(northEast)
    //   // console.log(southWest)
    //   tilePixelMax = LatLongToPixelXY(northEast.lat, northEast.lng, tileZoom)
    //   tilePixelMin = LatLongToPixelXY(southWest.lat, southWest.lng, tileZoom)

    //   var promise = new Promise(function(resolve, reject) {
    //     resolve(downloadMap())
    //   }).finally(function() {
    //     // 下载urlList文件
    //     var a = document.createElement("a");
    //     var file = new Blob([JSON.stringify(urlList)], {
    //       type: "application/json"
    //     });
    //     a.href = URL.createObjectURL(file);
    //     a.download = "urlList";
    //     a.click();
    //   });
    // });


    // function sleep(time) {
    //   return new Promise(resolve => setTimeout(resolve, time));
    // };

    // function sleep(time) {
    //   return new Promise(function(resolve) {
    //     setTimeout(resolve, time);
    //   });
    // }

    // quadkey = toQuadKey(3336, 1777, 12)
    // console.log(quadkey)
  };

  BingDownload.prototype.init.prototype = BingDownload.prototype
  // script标签引入插件后全局下挂载一个_$的方法
  global._$ = BingDownload;
})(this || window);

// 使用演示：
// <!-- 页面元素 -->
// <div id="app">hello world</div>
// 为元素设置背景：
// _$('#app').setBg('#ff0')
// 为元素设置背景并改变宽高：
// _$('#app').setBg('#ff0').setHeight('100px').setWidth('200px')

// window.g = ghostwu;

// var oBg = g.BgColor.Bg({
//   'activeClass': 'current',
//   'ele': 'table'
// });
// oBg.setBgColor();
// oBg.hover();

/**
 * 节点配置属性方式配置参数：专业的做法是配置到，每一个需要初始化为插件的那个html标签的属性上面(这个属性的value必须是标准的json字符串)，
 * 这样做的好处是，当存在多个需要初始化为插件的标签时，可以配置各自想要的参数
 * 例如：<div data-setting="{'param1':1,'param2':2}"></div>
 * Created by TonyZeng on 2016/7/6.
 */
/**前面的这个分号，用于避免其他插件没有分号闭合导致在压缩代码后导致语法的错误。*/
// ;
// (function($) {
//   /**这里相当于一个构造函数，当new Demo(demo)时，就会传递参数过来
//    * 不过，我们没有采用new的方式来初始化插件，我们定义了一个init方法来初始化，在init方法中，我们用的其实也就是这个
//    * 构造函数来初始化，init里面用each来循环初始化，将div节点对象（this）指针作为参数传入构造函数的参数列表
//    */
//   var Demo = function(demo) {
//     var self = this;
//     //保存节点对象，（获取自定义配置）getSetting方法会用它来获取节点里面的属性，比如自定义配置属性
//     this.demo = demo;

//     //默认配置
//     this.setting = { "maskDivId": "xxx", "showDivId": "xxx", "closeButtonId": "xxx" };
//     //用自定义配置merge默认配置
//     $.extend(this.setting, this.getSetting());

//     //事件绑定
//     $(".demo").bind("click", function(e) {
//       self.showdiv();
//       //阻止事件冒泡出现(冒泡事件)
//       e.stopPropagation();
//     });
//     $("#" + this.setting.closeButtonId).bind("click", function(e) {
//       self.hidediv();
//       //阻止事件冒泡出现(冒泡事件)
//       e.stopPropagation();
//     });

//   };
//   Demo.prototype = {
//     //获取自定义配置
//     getSetting: function() {
//       var setting = this.demo.attr("data-setting");
//       //如果这个节点属性存在，并且不为空的话就转成json对象传出去
//       if (setting && setting != null) {
//         return $.parseJSON(setting);
//       } else {
//         return {};
//       }
//     },
//     showdiv: function() {
//       document.getElementById(this.setting.showDivId).style.display = "block";
//       document.getElementById(this.setting.maskDivId).style.display = "block";
//     },
//     hidediv: function() {
//       document.getElementById(this.setting.showDivId).style.display = 'none';
//       document.getElementById(this.setting.maskDivId).style.display = 'none';
//     }
//   };
//   //（当页面有多个地方需要初始化成插件时）为了避免用new的方式来初始化插件的繁琐，我们用一个init方法来专门处理初始化
//   Demo.init = function(list) {
//     //这个this指向Demo插件对象
//     var _this_ = this;
//     //用each循环初始化传入的元素为插件
//     list.each(function() {
//       //这个this指向each循环里面的元素对象
//       new _this_($(this));
//     });
//   };
//   //全局注册
//   window["Demo"] = Demo;
// })(jQuery);
// $(function() {
//   Demo.init($(".demo"));
// });

// (function() {
//   function $() {
//     //代码
//     alert("test");
//   }

//   function getElementByClassName() {}
//   //将dtle命名空间注册到winow对象上
//   window["dtle"] = {}
//   //将$函数注册到dtle命名空间下
//   window["dtle"]["$"] = $
//   //将getElementByClassName函数注册到dtle命名空间下
//   window["dtle"]["getElementByClassName"] = getElementByClassName
// })();

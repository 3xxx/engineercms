<!DOCTYPE html>
<html>

<head>
  <meta content="text/html; charset=utf-8" http-equiv="Content-Type">
  <title>flv.js demo</title>
  <style>
    .mainContainer {
      display: block;
      width: 1024px;
      margin-left: auto;
      margin-right: auto;
    }
 
    .urlInput {
      display: block;
      width: 100%;
      margin-left: auto;
      margin-right: auto;
      margin-top: 8px;
      margin-bottom: 8px;
    }
 
    .centeredVideo {
      display: block;
      width: 100%;
      height: 576px;
      margin-left: auto;
      margin-right: auto;
      margin-bottom: auto;
    }
 
    .controls {
      display: block;
      width: 100%;
      text-align: left;
      margin-left: auto;
      margin-right: auto;
    }
    </style>
</head>

<body>
  <p class="mainContainer">
    <input name="urlinput" class="urlInput" type="text" value="/static/download/VID20190517113037.mp4" />
    <div class="options">
      <input type="checkbox" id="isLive" onchange="saveSettings()" checked />
      <label for="isLive">isLive</label>
      <input type="checkbox" id="withCredentials" onchange="saveSettings()" />
      <label for="withCredentials">withCredentials</label>
      <input type="checkbox" id="hasAudio" onchange="saveSettings()" checked />
      <label for="hasAudio">hasAudio</label>
      <input type="checkbox" id="hasVideo" onchange="saveSettings()" checked />
      <label for="hasVideo">hasVideo</label>
    </div>
    <video name="videoElement" class="centeredVideo" controls autoplay width="1024" height="576">
      Your browser is too old which doesn't support HTML5 video.
    </video>
    <br>
    <p class="controls">
      <button onclick="flv_load()">Load</button>
      <button onclick="flv_start()">Start</button>
      <button onclick="flv_pause()">Pause</button>
      <button onclick="flv_destroy()">Destroy</button>
      <input style="width:100px" type="text" name="seekpoint" />
      <button onclick="flv_seekto()">SeekTo</button>
    </p>
  </p>
  <script src="/static/js/flv.min.js"></script>
  <script>
  // var player = document.getElementById('videoElement');
  // if (flvjs.isSupported()) {
  //   var flvPlayer = flvjs.createPlayer({
  //     type: 'flv',
  //     url: '你的视频.flv'
  //   });
  //   flvPlayer.attachMediaElement(videoElement);
  //   flvPlayer.load(); //加载
  // }

  // function flv_start() {
  //   player.play();
  // }

  // function flv_pause() {
  //   player.pause();
  // }

  // function flv_destroy() {
  //   player.pause();
  //   player.unload();
  //   player.detachMediaElement();
  //   player.destroy();
  //   player = null;
  // }

  // function flv_seekto() {
  //   player.currentTime = parseFloat(document.getElementsByName('seekpoint')[0].value);
  // }
  </script>
  <script>
  // var player;
  var player = document.getElementById('videoElement');

  function flv_load() {
    console.log('isSupported: ' + flvjs.isSupported());
    var urlinput = document.getElementsByName('urlinput')[0];
    var xhr = new XMLHttpRequest();
    xhr.open('GET', urlinput.value, true);
    xhr.onload = function(e) {
      var element = document.getElementsByName('videoElement')[0];
      if (typeof player !== "undefined") {
        if (player != null) {
          player.unload();
          player.detachMediaElement();
          player.destroy();
          player = null;
        }
      }

      player = flvjs.createPlayer({
        type: 'mp4',
        url: urlinput.value
      });
      player.attachMediaElement(element);
      player.load();
    }
    xhr.send();
  }

  function flv_start() {
    player.play();
  }

  function flv_pause() {
    player.pause();
  }

  function flv_destroy() {
    player.pause();
    player.unload();
    player.detachMediaElement();
    player.destroy();
    player = null;
  }

  function flv_seekto() {
    var input = document.getElementsByName('seekpoint')[0];
    player.currentTime = parseFloat(input.value);
  }

  function getUrlParam(key, defaultValue) {
    var pageUrl = window.location.search.substring(1);
    var pairs = pageUrl.split('&');
    for (var i = 0; i < pairs.length; i++) {
      var keyAndValue = pairs[i].split('=');
      if (keyAndValue[0] === key) {
        return keyAndValue[1];
      }
    }
    return defaultValue;
  }

  var urlInputBox = document.getElementsByName('urlinput')[0];
  var url = decodeURIComponent(getUrlParam('src', urlInputBox.value));
  urlInputBox.value = url;

  document.addEventListener('DOMContentLoaded', function() {
    flv_load();
  });
  </script>
</body>

</html>
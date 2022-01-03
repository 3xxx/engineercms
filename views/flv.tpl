<!DOCTYPE html>
<html>

<head>
  <meta content="text/html; charset=utf-8" http-equiv="Content-Type">
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <title>flv.js demo</title>
  <!-- <link rel="stylesheet" type="text/css" href="demo.css" /> -->
  <style type="text/css">
  .mainContainer {
    display: block;
    width: 100%;
    margin-left: auto;
    margin-right: auto;
  }

  @media screen and (min-width: 1152px) {
    .mainContainer {
      display: block;
      width: 1152px;
      margin-left: auto;
      margin-right: auto;
    }
  }

  .video-container {
    position: relative;
    margin-top: 8px;
  }

  .video-container:before {
    display: block;
    content: "";
    width: 100%;
    padding-bottom: 56.25%;
  }

  .video-container>div {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
  }

  .video-container video {
    width: 100%;
    height: 100%;
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
    height: 100%;
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
    margin-top: 8px;
    margin-bottom: 10px;
  }

  .logcatBox {
    border-color: #CCCCCC;
    font-size: 11px;
    font-family: Menlo, Consolas, monospace;
    display: block;
    width: 100%;
    text-align: left;
    margin-left: auto;
    margin-right: auto;
  }

  .url-input,
  .options {
    font-size: 13px;
  }

  .url-input {
    display: flex;
  }

  .url-input label {
    flex: initial;
  }

  .url-input input {
    flex: auto;
    margin-left: 8px;
  }

  .url-input button {
    flex: initial;
    margin-left: 8px;
  }

  .options {
    margin-top: 5px;
  }

  .hidden {
    display: none;
  }
  </style>
</head>

<body>
  <div class="mainContainer">
    <div>
      <div id="streamURL">
        <div class="url-input">
          <label for="sURL">Stream URL:</label>
          <input id="sURL" type="text" value="{{.Mp4Link}}"/>
          <!-- "/static/download/VID20190517113037.mp4"  -->
          <button onclick="switch_mds()">Switch to MediaDataSource</button>
        </div>
        <div class="options">
          <input type="checkbox" id="isLive" onchange="saveSettings()" />
          <label for="isLive">isLive</label>
          <input type="checkbox" id="withCredentials" onchange="saveSettings()" />
          <label for="withCredentials">withCredentials</label>
          <input type="checkbox" id="hasAudio" onchange="saveSettings()" checked />
          <label for="hasAudio">hasAudio</label>
          <input type="checkbox" id="hasVideo" onchange="saveSettings()" checked />
          <label for="hasVideo">hasVideo</label>
        </div>
      </div>
      <div id="mediaSourceURL" class="hidden">
        <div class="url-input">
          <label for="msURL">MediaDataSource JsonURL:</label>
          <input id="msURL" type="text" value="/static/download/flv.json" />
          <button onclick="switch_url()">Switch to URL</button>
        </div>
      </div>
    </div>
    <div class="video-container">
      <div>
        <video name="videoElement" class="centeredVideo" controls autoplay>
          Your browser is too old which doesn't support HTML5 video.
        </video>
      </div>
    </div>
    <div class="controls">
      <button onclick="flv_load()">Load</button>
      <button onclick="flv_start()">Start</button>
      <button onclick="flv_pause()">Pause</button>
      <button onclick="flv_destroy()">Destroy</button>
      <input style="width:100px" type="text" name="seekpoint" />
      <button onclick="flv_seekto()">SeekTo</button>
    </div>
    <!-- <textarea name="logcatbox" class="logcatBox" rows="10" readonly></textarea> -->
  </div>
  <script src="/static/js/flv.min.js"></script>
  <script>
    $(function() {
      $('#sURL').val('{{.Mp4Link}}');
      // document.getElementById('name1').value='值';
    })

  var checkBoxFields = ['isLive', 'withCredentials', 'hasAudio', 'hasVideo'];
  var streamURL, mediaSourceURL;

  function flv_load() {
    console.log('isSupported: ' + flvjs.isSupported());
    if (mediaSourceURL.className === '') {
      var url = document.getElementById('msURL').value;

      var xhr = new XMLHttpRequest();
      xhr.open('GET', url, true);
      xhr.onload = function(e) {
        var mediaDataSource = JSON.parse(xhr.response);
        flv_load_mds(mediaDataSource);
      }
      xhr.send();
    } else {
      var i;
      var mediaDataSource = {
        type: 'mp4'
      };
      for (i = 0; i < checkBoxFields.length; i++) {
        var field = checkBoxFields[i];
        /** @type {HTMLInputElement} */
        var checkbox = document.getElementById(field);
        mediaDataSource[field] = checkbox.checked;
      }
      mediaDataSource['url'] = document.getElementById('sURL').value;
      console.log('MediaDataSource', mediaDataSource);
      flv_load_mds(mediaDataSource);
    }
  }

  function flv_load_mds(mediaDataSource) {
    var element = document.getElementsByName('videoElement')[0];
    if (typeof player !== "undefined") {
      if (player != null) {
        player.unload();
        player.detachMediaElement();
        player.destroy();
        player = null;
      }
    }
    player = flvjs.createPlayer(mediaDataSource, {
      enableWorker: false,
      lazyLoadMaxDuration: 3 * 60,
      seekType: 'range',
    });
    player.attachMediaElement(element);
    player.load();
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

  function switch_url() {
    streamURL.className = '';
    mediaSourceURL.className = 'hidden';
    saveSettings();
  }

  function switch_mds() {
    streamURL.className = 'hidden';
    mediaSourceURL.className = '';
    saveSettings();
  }

  function ls_get(key, def) {
    try {
      var ret = localStorage.getItem('flvjs_demo.' + key);
      if (ret === null) {
        ret = def;
      }
      return ret;
    } catch (e) {}
    return def;
  }

  function ls_set(key, value) {
    try {
      localStorage.setItem('flvjs_demo.' + key, value);
    } catch (e) {}
  }

  function saveSettings() {
    if (mediaSourceURL.className === '') {
      ls_set('inputMode', 'MediaDataSource');
    } else {
      ls_set('inputMode', 'StreamURL');
    }
    var i;
    for (i = 0; i < checkBoxFields.length; i++) {
      var field = checkBoxFields[i];
      /** @type {HTMLInputElement} */
      var checkbox = document.getElementById(field);
      ls_set(field, checkbox.checked ? '1' : '0');
    }
    var msURL = document.getElementById('msURL');
    var sURL = document.getElementById('sURL');
    ls_set('msURL', msURL.value);
    ls_set('sURL', sURL.value);
    console.log('save');
  }

  function loadSettings() {
    var i;
    for (i = 0; i < checkBoxFields.length; i++) {
      var field = checkBoxFields[i];
      /** @type {HTMLInputElement} */
      var checkbox = document.getElementById(field);
      var c = ls_get(field, checkbox.checked ? '1' : '0');
      checkbox.checked = c === '1' ? true : false;
    }

    var msURL = document.getElementById('msURL');
    var sURL = document.getElementById('sURL');
    msURL.value = ls_get('msURL', msURL.value);
    sURL.value = ls_get('sURL', sURL.value);
    if (ls_get('inputMode', 'StreamURL') === 'StreamURL') {
      switch_url();
    } else {
      switch_mds();
    }
  }

  function showVersion() {
    var version = flvjs.version;
    document.title = document.title + " (v" + version + ")";
  }

  var logcatbox = document.getElementsByName('logcatbox')[0];
  flvjs.LoggingControl.addLogListener(function(type, str) {
    logcatbox.value = logcatbox.value + str + '\n';
    logcatbox.scrollTop = logcatbox.scrollHeight;
  });

  document.addEventListener('DOMContentLoaded', function() {
    streamURL = document.getElementById('streamURL');
    mediaSourceURL = document.getElementById('mediaSourceURL');
    loadSettings();
    showVersion();
    flv_load();
  });
  </script>
</body>

</html>
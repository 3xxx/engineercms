<!DOCTYPE html>
<html style="height: 100%;">

<head>
  <title>{{.Doc.FileName}}</title>
  <!-- 收藏用logo图标 -->
  <link rel="bookmark" type="image/x-icon" href="/static/img/only.ico" />
  <!-- 网站显示页logo图标 -->
  <link rel="shortcut icon" href="/static/img/only.ico">
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0">
  <meta name="renderer" content="webkit">
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css" />
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script type="text/javascript" src="/static/js/codehandle.js"></script>
  <style type="text/css">
  .floating-button2 {
    display: block;
    width: 58px;
    height: 58px;
    border-radius: 50%;
    background: #f44336;
    color: #fff;
    margin: 0 auto;
    text-align: center;
    float: right;
    /* background-color: #fff; */
    position: fixed;
    top: 40px;
    /*bottom: 20px;*/
    right: 250px;
    /* border: 0 solid #fff; */
    /* border-radius: 500px; */
    box-shadow: 4px 1px 1px #ccc;
    opacity: 0.6;
    z-index: 999;
    display: -webkit-box;
    display: -ms-flexbox;
    display: -webkit-flex;
    display: flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    -webkit-align-items: center;
    align-items: center;
    -webkit-box-pack: center;
    -ms-flex-pack: center;
    -webkit-justify-content: center;
    justify-content: center;
  }
  </style>
</head>

<body style="height: 100%; margin: 0;">
  <!-- 保留<button type="button" class="floating-button2" id="getSignature">获取签名</button> -->

  <div id="placeholder" style="height: 100%"></div>
  <script type="text/javascript" src="{{.Onlyofficeapi_url}}/web-apps/apps/api/documents/api.js"></script>
  <script type="text/javascript">
  var onAppReady = function() {
    console.log("ONLYOFFICE Document Editor is ready");
  };
  var onCollaborativeChanges = function() {
    console.log("The document changed by collaborative user");
  };
  var onDocumentReady = function() {
    console.log("Document is loaded");
  };
  var onDocumentStateChange = function(event) {
    if (event.data) {
      console.log("The document changed");
      console.log(event.data)
    } else {
      console.log("Changes are collected on document editing service");
    }
  };

  var onDownloadAs = function (event) {
    console.log("ONLYOFFICE Document Editor create file: " + event.data);
    // docEditor.downloadAs("doc");
  };


  $("#downloadas").click(function(event) {
    docEditor.downloadAs();
  });

  /**
   * 创建并下载文件
   * @param {String} fileName 文件名
   * @param {String} content 文件内容
   */
  function createAndDownloadFile(fileName, content) {
    var aTag = document.createElement('a');
    var blob = new Blob([content]);
    aTag.download = fileName;
    aTag.href = URL.createObjectURL(blob);
    aTag.click();
    URL.revokeObjectURL(blob);
    console.log(content)
  }

  window.addEventListener('message', function(e) {
    // 将字符串转换为 JavaScript 对象
    str = JSON.parse(e.data)
    if (str.event == "onDownloadAs"){
      var aTag = document.createElement('a');
      // var blob = new Blob([str.data]);——这个是内容，保留
      // aTag.download = "1.doc";
      // aTag.href = URL.createObjectURL(blob);
      aTag.href = str.data
      aTag.click();
    }
  }, false)

  $("#insertImage").click(function(event) {
    console.log("ONLYOFFICE Document Editor insertImage: " + event.data);
    docEditor.insertImage({
      "fileType": "png",
      "url": "{{.Engineercmsapi_url}}/attachment/20190728测试上传文件名修改/2020January/1580363537940306800_small.png"
    });
  })

  var onRequestInsertImage = function(event) {
    console.log("ONLYOFFICE Document Editor insertImage" + event.data);
    docEditor.insertImage({
      "fileType": "png",
      "url": "{{.Engineercmsapi_url}}/attachment/20190728测试上传文件名修改/2020January/1580363537940306800_small.png"
    });
  };

  var onError = function(event) {
    console.log("ONLYOFFICE Document Editor reports an error: code " + event.data.errorCode + ", description " + event.data.errorDescription);
  };
  var onOutdatedVersion = function() {
    location.reload(true);
  };
  var onRequestEditRights = function() {
    console.log("ONLYOFFICE Document Editor requests editing rights");
    var he = location.href.replace("view", "edit");
    location.href = he;
  };

  //历史版本保留1个月。比如Unix时间戳（Unix timestamp）expires=1524547423
  var onRequestHistory = function() {
  };

  var onRequestHistoryClose = function() {
    document.location.reload();
  };

  var onRequestHistoryData = function(event) {
  };

  const config = {
    "document": {
      "fileType": "{{.fileType}}",
      "key": "{{.Key}}", //"Khirz6zTPdfd7"
      "title": "{{.Doc.FileName}}",
      "url":"http:\/\/192.168.100.37/attachment/onlyoffice/测试v3.docx"
      // "url": "{{.Engineercmsapi_url}}/attachment/onlyoffice/{{.Doc.FileName}}?hotqinsessionid={{.Sessionid}}",
    },
    "documentType": "{{.documentType}}",
    "editorConfig": {
      "callbackUrl": "{{.Engineercmsapi_url}}/url-to-callback?id={{.Doc.Id}}",
      "customization": {
        "uiTheme": "theme-dark",
        "unit": "cm",
        "wordHeadingsColor": "#00ff00",
        "zoom": 100,
      },
      "user": {
        "id": {{.Uid }},
        "name": "{{.Username}}"
      },
      "lang": "zh-CN", //"en-US",
      "mode": {{.Mode }}, //"view",//edit
      "region": "zh-CN",
    },
    "height": "100%",
    "type": {{.Type }}, //"desktop",embedded,mobile访问文档的平台类型 网页嵌入
    "width": "100%"
  }
  // window.docEditor = new DocsAPI.DocEditor("placeholder", config);

  // const config = {
  //   "document": {
  // };
  $(function () {
    const configJsonStr = JSON.stringify(config);
    $.ajax({
      type: "POST",
      url: "/v1/onlyoffice/jwtencode",
      contentType: "application/json",
      // data: JSON.stringify({
      //   "jsonStr":configJsonStr
      // }),
      data: configJsonStr,
      dataType: "json",
      success: function (data) {
        console.log("成功")
        console.log(data)
        if(data.jwt){
          config.token = data.jwt;
        }  
        var docEditor = new DocsAPI.DocEditor("placeholder", config);
      },
      error: function (err) {
        console.error(err);
      }
    })
  })
  </script>
</body>

</html>
<!-- {
"key":"1521953170330601700",
"status":2,

"url":"{{.Engineercmsapi_url}}00:9000/cache/files/1521953170330601700_4540/outpu
t.docx/output.docx?md5=eSwnrSSumTeMuh59IoXhCQ==&expires=1524547423&disposition=a
ttachment&ooname=output.docx",

"changesurl":"{{.Engineercmsapi_url}}00:9000/cache/fil
es/1521953170330601700_4540/changes.zip/changes.zip?md5=w6DItkSwyBJkuHDl_CiZZQ==
&expires=1524547423&disposition=attachment&ooname=output.zip",

"history":{
  "serverVersion":"5.0.7",
  "changes":[{
    "created":"2018-03-25 05:23:25",
    "user":{"id":"127.0 .0.1","name":"127.0.0.1"}
    }]
},

"users":["127.0.0.1"],
"actions":[{"type":0,"userid":"127.0.0.1"}],
"lastsave":"2018-03-25T05:23:30.342Z",
"notmodified":false
} -->
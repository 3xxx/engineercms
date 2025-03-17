<!DOCTYPE html>
<html style="height: 100%;">

<head>
  <title>{{.Doc.FileName}}</title>
  <!-- fei-OnlyOffice -->
  <link rel="bookmark" type="image/x-icon" href="/static/img/only.ico" />
  <link rel="shortcut icon" href="/static/img/only.ico">
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0">
  <meta name="renderer" content="webkit">
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
</head>

<body style="height: 100%; margin: 0;">
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
    } else {
      console.log("Changes are collected on document editing service");
    }
  };
  var onDownloadAs = function(event) {
    console.log("ONLYOFFICE Document Editor create file: " + event.data);
  };
  var onError = function(event) {
    console.log("ONLYOFFICE Document Editor reports an error: code " + event.data.errorCode + ", description " + event.data.errorDescription);
  };
  var onOutdatedVersion = function() {
    location.reload(true);
  };
  var onRequestEditRights = function() {
    // console.log("ONLYOFFICE Document Editor requests editing rights");
    // document.location.reload();
    var he = location.href.replace("view", "edit");
    location.href = he;
  };

  var Url = "{{.Engineercmsapi_url}}/{{.FilePath}}?hotqinsessionid={{.Sessionid}}";
  // var Url2="http://192.168.99.1/"+Url.replace(/\u0026/,"");
  // alert(Url2);
  window.docEditor = new DocsAPI.DocEditor("placeholder", {
    "events": {
      "onAppReady": onAppReady,
      "onCollaborativeChanges": onCollaborativeChanges,
      "onDocumentReady": onDocumentReady,
      "onDocumentStateChange": onDocumentStateChange,
      "onDownloadAs": onDownloadAs,
      "onError": onError,
    },

  const config = {
    "document": {
      "fileType": "{{.fileType}}",
      "key": "{{.Key}}", //"Khirz6zTPdfd7"
      "title": "{{.Doc.FileName}}",
      "url": "{{.Engineercmsapi_url}}/{{.FilePath}}?hotqinsessionid={{.Sessionid}}",
    },
    "documentType": "{{.documentType}}",
    "editorConfig": {
      "callbackUrl": "{{.Engineercmsapi_url}}/officeviewcallback?id={{.AttachId}}",
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

"url":"http://192.168.99.100:9000/cache/files/1521953170330601700_4540/outpu
t.docx/output.docx?md5=eSwnrSSumTeMuh59IoXhCQ==&expires=1524547423&disposition=a
ttachment&ooname=output.docx",

"changesurl":"http://192.168.99.100:9000/cache/fil
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
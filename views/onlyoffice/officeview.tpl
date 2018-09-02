<!DOCTYPE html>
<html style="height: 100%;">
	<head>
	  <title>{{.Doc.FileName}}</title>
<!-- fei-OnlyOffice -->
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0">
	<meta name="renderer" content="webkit">
	<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
	
	</head>

	<body style="height: 100%; margin: 0;">
		<div id="placeholder" style="height: 100%"></div>
    <script type="text/javascript" src="http://192.168.99.100:9000/web-apps/apps/api/documents/api.js"></script>

    <script type="text/javascript">

      var onAppReady = function() {
          console.log("ONLYOFFICE DocumenEditor is ready");
      }; 
      var onCollaborativeChanges = function () {
          console.log("The document changed by collaborative user");
      };
      var onDocumentReady = function() {
          console.log("Document is loaded");
      };
      var onDocumentStateChange = function (event) {
          if (event.data) {
              console.log("The document changed");
          } else {
              console.log("Changes are collected on document editing service");
          }
      };
      var onDownloadAs = function (event) {
          console.log("ONLYOFFICE Document Editor create file: " + event.data);
      };
      var onError = function (event) {
          console.log("ONLYOFFICE Document Editor reports an error: code " + event.data.errorCode + ", description " + event.data.errorDescription);
      };
      var onOutdatedVersion = function () {
          location.reload(true);
      };
      var onRequestEditRights = function () {
          // console.log("ONLYOFFICE Document Editor requests editing rights");
          // document.location.reload();
          var he=location.href.replace("view","edit");
          location.href=he;
      };

      var Url="http://192.168.99.1/"+{{.FilePath}};
      // var Url2="http://192.168.99.1/"+Url.replace(/\u0026/,"");
      // alert(Url2);
    	window.docEditor = new DocsAPI.DocEditor("placeholder",
      	{
        "events": {
          "onAppReady": onAppReady,
          "onCollaborativeChanges": onCollaborativeChanges,
          "onDocumentReady": onDocumentReady,
          "onDocumentStateChange": onDocumentStateChange,
          "onDownloadAs": onDownloadAs,
          "onError": onError,
        },

      	"document": {
          "fileType": "{{.fileType}}",
          "key": "{{.Key}}",//"Khirz6zTPdfd7"
          "title": "{{.Doc.FileName}}",
          "url": Url,

          "permissions": {
          	"comment": {{.Comment}},//true,
          	"download": {{.Download}},//true,
          	"edit": {{.Edit}},
          	"print": {{.Print}},//true,
          	"review": {{.Review}}//true
        	},
        },
        "documentType": "{{.documentType}}",
        "editorConfig": {
          "callbackUrl": "http://192.168.99.1/officeviewcallback",
          "user": {
            "id": {{.Uid}},
            "name": "{{.Username}}"
          },
					"customization": {
            "chat": true,
            "commentAuthorOnly": false,
            "compactToolbar": false,
            "customer": {
              "address": "116# Tianshou Road,Guangzhou China",
              "info": "QQ504284",
              "logo": "http://192.168.99.1/static/img/user.jpg",//logo-big.png
              "mail": "xc-qin@163.com",
              "name": "Qin Xiao Chuan",
              "www": "github.com/3xxx"
            },
            "feedback": {
              "url": "http://192.168.99.1/onlyoffice",
              "visible": true
            },
            "forcesave": true,
            "goback": {
              "text": "Go to Documents",
              "url": "http://192.168.99.1/onlyoffice"
            },
            "logo": {
              "image": "http://192.168.99.1/static/img/oo.png",//logo.png
              "imageEmbedded": "http://192.168.99.1/static/img/oo.png",//logo_em.png
              "url": "http://192.168.99.1/onlyoffice"
            },
            "showReviewChanges": false,
            "zoom": 100,
        	},
        	"embedded": {
            "embedUrl": "https://example.com/embedded?doc=exampledocument1.docx",
            "fullscreenUrl": "https://example.com/embedded?doc=exampledocument1.docx#fullscreen",
            "saveUrl": "https://example.com/download?doc=exampledocument1.docx",
            "shareUrl": "https://example.com/view?doc=exampledocument1.docx",
            "toolbarDocked": "top"
        	},
          "lang": "zh-CN",//"en-US",
					"mode": {{.Mode}},//"view",//edit
					"recent": [
            {
              "folder": "Example Files",
              "title": "exampledocument1.docx",
              "url": "https://example.com/exampledocument1.docx"
            },
            {
              "folder": "Example Files",
              "title": "exampledocument2.docx",
              "url": "https://example.com/exampledocument2.docx"
            }
        	]
        },
        "height": "100%",
        // "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.t-IDcSemACt8x4iTMCda8Yhe3iZaWbvV5XKSTbuAn0M",
    		"type": {{.Type}},//"desktop",embedded,mobile访问文档的平台类型 网页嵌入
        "width": "100%"
      });
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
		"user":{"id":"127.0	.0.1","name":"127.0.0.1"}
		}]
},

"users":["127.0.0.1"],
"actions":[{"type":0,"userid":"127.0.0.1"}],
"lastsave":"2018-03-25T05:23:30.342Z",
"notmodified":false
} -->

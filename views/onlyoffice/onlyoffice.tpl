<!DOCTYPE html>
<html style="height: 100%;">
	<head>
	  <title>fei-OnlyOffice</title>

	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0">
	<meta name="renderer" content="webkit">
	<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
	
	</head>

	<body style="height: 100%; margin: 0;">
		<div id="placeholder" style="height: 100%"></div>
    <script type="text/javascript" src="http://192.168.99.100:9000/web-apps/apps/api/documents/api.js"></script>


    <script type="text/javascript">
    	//历史版本保留1个月。比如Unix时间戳（Unix timestamp）expires=1524547423
      var onRequestHistory = function() {

      	docEditor.refreshHistory({
        "currentVersion": {{.currentversion}},
        "history":{{.onlyhistory}},
      //   "history": [
      //     {
      //     	"changes": changes,//[{{.changes1}}], //the changes from the history object returned after saving the document
      //       "created": "2018-03-9 10:15:55",
      //       "key": "1522427166608304100",//1521951775531484800这里影响历史版本切换
      //       "serverVersion": "5.07", //the serverVersion from the history object returned after saving the document
      //   	  "user": {
      //   	    "id": "9",
      //   	    "name": "qin.xc"
      //   	  },
      //   	  "version": 1
      //   	},
      //   	{
      // 		  "changes": changes,
      // 		  "created": "2018-03-10 14:11:35",
      // 		  "key": "1522465759378671300",//
      // 		  "user": {
      // 		      "id": "9",
      // 		      "name": "qin.xc"
      // 		  },
      // 		  "version": 2
      // 		},
      //   	{
      // 		  "changes": changes,
      // 		  "created": "2018-03-11 14:11:35",
      // 		  "key": "1522470906000209200",//当前版本
      // 		  "user": {
      // 		      "id": "9",
      // 		      "name": "qin.xc"
      // 		  },
      // 		  "version": 3
      // 		},
      		
      // 		{
      // 		  "changes": changes,
      // 		  "created": "2018-03-11 14:11:35",
      // 		  "key": "1522475922103673500",//当前版本
      // 		  "user": {
      // 		      "id": "9",
      // 		      "name": "qin.xc"
      // 		  },
      // 		  "version": 4
      // 		},
    		// ]
  			});
			};

			var onRequestHistoryClose = function() {
  		  document.location.reload();
			};

			var onRequestHistoryData = function(event) {
    		var version = event.data;
        var history={{.onlyhistory}};
        var fileUrl="";
        var changeUrl="";
        var key="";
        var previousKey="";
        var previousUrl="";
    		for(var i=history.length-1;i>=0;i--){
    			if (version==history[i].version){
            changeUrl=history[i].changesurl;
            fileUrl=history[i].fileurl;
            key=history[i].key;
            if(i>0){
              previousKey=history[i-1].key;
              previousUrl=history[i-1].fileurl;
            }else{
              previousKey=key;
              previousUrl=fileUrl;
            }
            break;
					}
    		}
				var changeUrl2=changeUrl.replace(/\u0026/,"&");
				// alert(changeUrl2);
    		docEditor.setHistoryData({
    			//下面这里存变化的位置
      		// "changesUrl":"http://192.168.99.100:9000/cache/files/1522475922103673500_7157/changes.zip/changes.zip?md5=syFUueSXdnCWe60Iym001g==&expires=1525068326&disposition=attachment&ooname=output.zip",//string1, //the changesUrl from the JSON object returned after saving the document
      		"changesUrl":changeUrl2,
      		"key": key,
      		"previous": {
      		  "key": previousKey,//这里不影响版本切换。与上个版本对比
      		  "url": previousUrl//http://192.168.99.100:9000/cache/files/1521953170330601700_4540/output.docx/output.docx?md5=eSwnrSSumTeMuh59IoXhCQ==&expires=1524547423&disposition=attachment&ooname=output.docx这里影响版本
      		},
      		"url": fileUrl,
      		"version": version
    		})
			};

    	// alert({{.Doc.FileName}});
    	window.docEditor = new DocsAPI.DocEditor("placeholder",
      	{
        "events": {
          "onRequestHistory": onRequestHistory,
          "onRequestHistoryClose": onRequestHistoryClose,
          "onRequestHistoryData": onRequestHistoryData,
        },

      	"document": {
          "fileType": "{{.fileType}}",
          "key": "{{.Key}}",//"Khirz6zTPdfd7"
          "title": "{{.Doc.FileName}}",
          "url": "http://192.168.99.1/attachment/onlyoffice/{{.Doc.FileName}}",
          "info": {
            "author": "John Smith",
            "created": "2010-07-07 3:46 PM",
            "folder": "Example Files",
            "sharingSettings": [
              {
                "permissions": "Full Access",
                "user": "John Smith"
              },
              {
                "permissions": "Read Only",
                "user": "Kate Cage"
              },
              {
                "permissions": "Deny Access",
                "user": "Albet Tlanp"
              },
               
            ]
          },
          "permissions": {
          	"comment": true,
          	"download": true,
          	"edit": {{.Edit}},
          	"print": true,
          	"review": {{.Review}}//true
        	},
        },
        "documentType": "{{.documentType}}",
        "editorConfig": {
          "callbackUrl": "http://192.168.99.1/url-to-callback?id={{.Doc.Id}}",
          
        	"createUrl": "https://example.com/url-to-create-document/",
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
              "logo": "http://192.168.99.1/static/test/user.jpg",//logo-big.png
              "mail": "xc-qin@163.com",
              "name": "Qin Xiao Chuan",
              "www": "github.com/3xxx"
            },
            "feedback": {
              "url": "http://192.168.99.1/onlyoffice",
              "visible": true
            },
            "forcesave": false,
            "goback": {
              "text": "Go to Documents",
              "url": "http://192.168.99.1/onlyoffice"
            },
            "logo": {
              "image": "http://192.168.99.1/static/test/user.jpg",//logo.png
              "imageEmbedded": "http://192.168.99.1/static/test/user.jpg",//logo_em.png
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
    		"type": {{.Type}},//"desktop",//desktop//embedded//mobile访问文档的平台类型 网页嵌入
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

<!DOCTYPE html>
<html style="height: 100%;">
	<head>
	  <title>fei-OnlyOffice</title>
	</head>
	<body style="height: 100%; margin: 0;">
		<div id="placeholder" style="height: 100%"></div>
    <script type="text/javascript" src="http://192.168.99.100:9000/web-apps/apps/api/documents/api.js"></script>
    <script type="text/javascript">
      var onRequestHistory = function() {
        // var changes=[{
        //     "created":"2018-03-10 14:22:15",
        //     "user":{"id":"8","name":"qin8.xc"}
        // }];
        // alert(changes[0].created);
      	docEditor.refreshHistory({
        "currentVersion": 2,
        "history": [
          {
          "changes": [{{.changes1}}], //the changes from the history object returned after saving the document
            "created": "2018-03-9 10:15:55",
            "key": "1520696086733383100",
            "serverVersion": "{{.serverVersion1}}", //the serverVersion from the history object returned after saving the document
        		  "user": {
        		    "id": "9",
        		    "name": "qin.xc"
        		  },
        		  "version": 1
        		},
      			{
      			  "changes": [{{.changes2}}],//"[{{.changes2}}]",
      			  "created": "2018-03-10 14:11:35",
      			  "key": "1520696086733383100",
      			  "user": {
      			      "id": "8",
      			      "name": "qin8.xc"
      			  },
      			  "version": 2
      			},
    		]
  			});
			};

			var onRequestHistoryClose = function() {
  		  document.location.reload();
			};
			var onRequestHistoryData = function(event) {
    // 指定允许其他域名访问    
    // header('Access-Control-Allow-Origin:*');    
    // 响应类型    
    // header('Access-Control-Allow-Methods:POST');    
    // 响应头设置    
    // header('Access-Control-Allow-Headers:x-requested-with,content-type');
    // response.addHeader("Access-Control-Allow-Origin", "*");
    // HttpServletResponse response = ServletActionContext.getResponse();
    var version = event.data;
    docEditor.setHistoryData({

      "changesUrl": "http://192.168.99.1/attachment/onlyoffice/changes.zip", //the changesUrl from the JSON object returned after saving the document
      "key": "1520696086733383100",//只有这个id起作用
      "previous": {
        "key": "1520696086733383100",
        "url": "http://192.168.99.1/attachment/onlyoffice/to-the-previous-version-of-the-document.docx"
      },
      "url": "http://192.168.99.1/attachment/onlyoffice/url-to-example-document.docx",
      "version":2 //version
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
          	"edit": true,
          	"print": true,
          	"review": true
        	},
        	"title": "{{.Doc.FileName}}",
        	"url": "https://example.com/url-to-example-document.docx"
        },
        "documentType": "{{.documentType}}",
        "editorConfig": {
          "callbackUrl": "http://192.168.99.1/url-to-callback?id={{.Doc.Id}}",
          
        	"createUrl": "https://example.com/url-to-create-document/",
          "user": {
            "id": "{{.Uid}}",
            "name": "{{.Uname}}"
          },


					"customization": {
            "chat": true,
            "commentAuthorOnly": false,
            "compactToolbar": false,
            "customer": {
              "address": "My City, 123a-45",
              "info": "Some additional information",
              "logo": "http://192.168.99.1/static/test/user.jpg",//logo-big.png
              "mail": "john@example.com",
              "name": "John Smith and Co.",
              "www": "example.com"
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
            "zoom": 100
        	},
        	"embedded": {
            "embedUrl": "https://example.com/embedded?doc=exampledocument1.docx",
            "fullscreenUrl": "https://example.com/embedded?doc=exampledocument1.docx#fullscreen",
            "saveUrl": "https://example.com/download?doc=exampledocument1.docx",
            "shareUrl": "https://example.com/view?doc=exampledocument1.docx",
            "toolbarDocked": "top"
        	},

          "lang": "zh-CN",//"en-US",

					// "mode": "view",//edit

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
    		"type": "desktop",//desktop//embedded//mobile访问文档的平台类型 网页嵌入
        "width": "100%"
      });


   	</script>
	</body>
</html>
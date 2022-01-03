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
  <!-- <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script> -->
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
  <!-- 保留<button type="button" data-name="insertImage" id="insertImage" class="btn btn-default">
      <i class="fa fa-download">插入图片</i>
    </button> -->
  <!-- <button type="button" data-name="downloadas" id="downloadas" class="btn btn-default">
    <i class="fa fa-download">下载文件</i>
  </button> -->

  <!-- 保留<button type="button" class="floating-button2" id="getSignature">获取签名</button> -->

  <div id="placeholder" style="height: 100%"></div>
  <script type="text/javascript" src="{{.Onlyofficeapi_url}}/web-apps/apps/api/documents/api.js"></script>
  <!-- http://111.230.181.182:8080/web-apps/apps/api/documents/api.js -->
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
      // docEditor.downloadAs();
    } else {
      console.log("Changes are collected on document editing service");
      //
    }
  };
  // var onDownloadAs = function(event) {
  //   console.log("ONLYOFFICE Document Editor create file: " + event.data);
  //   window.top.postMessage(event.data);
  //   // createAndDownloadFile("test.docx", event.data)
  //   docEditor.downloadAs("doc");
  // };

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
      // aTag.remove();
      // URL.revokeObjectURL(blob);
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
    // document.location.reload();
    var he = location.href.replace("view", "edit");
    location.href = he;
  };

  //历史版本保留1个月。比如Unix时间戳（Unix timestamp）expires=1524547423
  var onRequestHistory = function() {
    docEditor.refreshHistory({
      "currentVersion": {{.currentversion }},
      "history": {{.onlyhistory }},
      //   "history": [
      //     {
      //      "changes": changes,//[{{.changes1}}], //the changes from the history object returned after saving the document
      //       "created": "2018-03-9 10:15:55",
      //       "key": "1522427166608304100",//1521951775531484800这里影响历史版本切换
      //       "serverVersion": "5.07", //the serverVersion from the history object returned after saving the document
      //      "user": {
      //        "id": "9",
      //        "name": "qin.xc"
      //      },
      //      "version": 1
      //    },
      //    {
      //      "changes": changes,
      //      "created": "2018-03-10 14:11:35",
      //      "key": "1522465759378671300",//
      //      "user": {
      //          "id": "9",
      //          "name": "qin.xc"
      //      },
      //      "version": 2
      //    },
      //    {
      //      "changes": changes,
      //      "created": "2018-03-11 14:11:35",
      //      "key": "1522470906000209200",//当前版本
      //      "user": {
      //          "id": "9",
      //          "name": "qin.xc"
      //      },
      //      "version": 3
      //    },
      //    {
      //      "changes": changes,
      //      "created": "2018-03-11 14:11:35",
      //      "key": "1522475922103673500",//当前版本
      //      "user": {
      //          "id": "9",
      //          "name": "qin.xc"
      //      },
      //      "version": 4
      //    },
      // ]
    });
  };

  var onRequestHistoryClose = function() {
    document.location.reload();
  };

  var onRequestHistoryData = function(event) {
    var version = event.data;
    var history = {{.onlyhistory }};
    var fileUrl = "";
    var changeUrl = "";
    var key = "";
    var previousKey = "";
    var previousUrl = "";
    for (var i = history.length - 1; i >= 0; i--) {
      if (version == history[i].version) {
        changeUrl = history[i].changesurl;
        fileUrl = history[i].fileurl;
        key = history[i].key;
        if (i > 0) {
          previousKey = history[i - 1].key;
          previousUrl = history[i - 1].fileurl;
        } else {
          previousKey = key;
          previousUrl = fileUrl;
        }
        break;
      }
    }
    var changeUrl2 = "{{.Engineercmsapi_url}}" + changeUrl.replace(/\u0026/, "&");
    // alert(changeUrl2);
    var previousurl = "{{.Engineercmsapi_url}}" + previousUrl
    // alert(previousurl);
    var fileurl = "{{.Engineercmsapi_url}}" + fileUrl
    // alert(fileurl);
    docEditor.setHistoryData({
      //下面这里存变化的位置
      // "changesUrl":"{{.Engineercmsapi_url}}/cache/files/1522475922103673500_7157/changes.zip/changes.zip?md5=syFUueSXdnCWe60Iym001g==&expires=1525068326&disposition=attachment&ooname=output.zip",//string1, //the changesUrl from the JSON object returned after saving the document
      "changesUrl": changeUrl2,
      "key": key,
      "previous": {
        "key": previousKey, //这里不影响版本切换。与上个版本对比
        "url": previousurl //previousUrl//{{.Engineercmsapi_url}}/cache/files/1521953170330601700_4540/output.docx/output.docx?md5=eSwnrSSumTeMuh59IoXhCQ==&expires=1524547423&disposition=attachment&ooname=output.docx这里影响版本
      },
      "url": fileurl, //fileUrl,
      "version": version
    })
  };

  // alert({{.Doc.FileName}});
  window.docEditor = new DocsAPI.DocEditor("placeholder", {
    "events": {
      "onAppReady": onAppReady,
      "onCollaborativeChanges": onCollaborativeChanges,
      "onDocumentReady": onDocumentReady,
      "onDocumentStateChange": onDocumentStateChange,
      "onDownloadAs": onDownloadAs,
      "onError": onError,
      "onRequestEditRights": onRequestEditRights,
      "onRequestHistory": onRequestHistory,
      "onRequestHistoryClose": onRequestHistoryClose,
      "onRequestHistoryData": onRequestHistoryData,
      "onRequestInsertImage": onRequestInsertImage,
    },

    "document": {
      "fileType": "{{.fileType}}",
      "key": "{{.Key}}", //"Khirz6zTPdfd7"
      "title": "{{.Doc.FileName}}",
      "url": "{{.Engineercmsapi_url}}/attachment/onlyoffice/{{.Doc.FileName}}?hotqinsessionid={{.Sessionid}}",
      "info": {
        "author": "John Smith",
        "created": "2010-07-07 3:46 PM",
        "folder": "Example Files",
        "sharingSettings": [{
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
        "comment": {{.Comment }}, //true,
        "download": {{.Download }}, //true,
        "edit": {{.Edit }},
        "print": {{.Print }}, //true,
        "review": {{.Review }} //true
      },
    },
    "documentType": "{{.documentType}}",
    "editorConfig": {
      "callbackUrl": "{{.Engineercmsapi_url}}/url-to-callback?id={{.Doc.Id}}",
      "createUrl": "https://example.com/url-to-create-document/",
      "user": {
        "id": {{.Uid }},
        "name": "{{.Username}}"
      },
      "customization": {
        "chat": true,
        "commentAuthorOnly": false,
        "compactToolbar": false,
        "customer": {
          "address": "116# Tianshou Road,Guangzhou China",
          "info": "QQ504284",
          "logo": "{{.Engineercmsapi_url}}/static/img/user.jpg", //logo-big.png
          "mail": "xc-qin@163.com",
          "name": "Qin Xiao Chuan",
          "www": "github.com/3xxx"
        },
        "feedback": {
          "url": "{{.Engineercmsapi_url}}/onlyoffice",
          "visible": true
        },
        "forcesave": false,
        "goback": {
          "text": "Go to Documents",
          "url": "{{.Engineercmsapi_url}}/onlyoffice"
        },
        "logo": {
          "image": "{{.Engineercmsapi_url}}/static/img/oo.png", //logo.png
          "imageEmbedded": "{{.Engineercmsapi_url}}/static/img/oo.png", //logo_em.png
          "url": "{{.Engineercmsapi_url}}/onlyoffice"
        },
        "showReviewChanges": false,
        "zoom": 100,
        "macros": true,
        "macrosMode": "warn",
        "mentionShare": true,
        "plugins": true,
        "reviewDisplay": "original",
        "showReviewChanges": false,
        "spellcheck": false,
        "toolbarHideFileName": false,
        "toolbarNoTabs": false,
        "trackChanges": false,
        "unit": "cm",
        "zoom": 100
      },
      "embedded": {
        "embedUrl": "{{.Onlyofficeapi_url}}/embedded?doc={{.Doc.FileName}}",
        "fullscreenUrl": "{{.Onlyofficeapi_url}}/embedded?doc={{.Doc.FileName}}#fullscreen",
        "saveUrl": "{{.Onlyofficeapi_url}}/download?doc={{.Doc.FileName}}",
        "shareUrl": "{{.Onlyofficeapi_url}}/view?doc={{.Doc.FileName}}",
        "toolbarDocked": "top"
      },
      "lang": "zh-CN", //"en-US",
      "mode": {{.Mode }}, //"view",//edit
      "recent": [{
          "folder": "Example Files",
          "title": "exampledocument1.docx",
          "url": "https://example.com/exampledocument1.docx"
        },
        {
          "folder": "Example Files",
          "title": "exampledocument2.docx",
          "url": "https://example.com/exampledocument2.docx"
        }
      ],
      "plugins": {
        // "autostart": [
        //   "asc.{07FD8DFA-DFE0-4089-AL24-0730933CC80A}",
        // ],
        "pluginsData": [
          "{{.Onlyofficeapi_url}}/sdkjs-plugins/photoeditor/config.json"
        ]
      },
    },
    "height": "100%",
    // "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.t-IDcSemACt8x4iTMCda8Yhe3iZaWbvV5XKSTbuAn0M",
    "type": {{.Type }}, //"desktop",embedded,mobile访问文档的平台类型 网页嵌入
    "width": "100%"
  });

  var ret_data = ""

  function myAjax(json) {
    var code = 0;
    //请求加上时间戳，防止浏览器缓存
    $.ajax({
      /* type: "get",
            cache:false,   */
      // url: "http://127.0.0.1:28810/" + "'" + codeHandler.encode(JSON.stringify(json), 'base64') + "'" + "time=" + Date.parse(new Date()),
      async: false,
      url: "{{.Engineercmsapi_url}}/v1/pdfcpu/test/" + "'" + codeHandler.encode(JSON.stringify(json), 'base64') + "'" + "time=" + Date.parse(new Date()),
      success: function(result) {
        // if (!result.ret.startWith("00")) {
        //   code = -1;
        //   alert("Windows服务返回错误,ACTION:" + json.action);
        // } else {
        alert("解析图片数据！");
        // var ret_data = ""
        //ret_data = window.atob(result.bmpdata);
        ret_data = result.bmpdata;
        // $("#fin").attr('src', "data:image/bmp;base64," + ret_data);

        // pdf合并签名，成功后reload
        // $.ajax({
        //   cache: false,
        //   type: 'POST',
        //   url: "http://127.0.0.1/v1/pdfcpu/addwatermarks/"+{{.DocId}},
        //   async: false,
        //   dataType: 'json',
        //   data: { image: ret_data },
        //   success: function (data) {
        //     alert("签名成功!");
        //     window.location.reload(true);
        //   }
        // });
        // }
      },
      error: function(result) {
        code = -1;
        alert("Windows服务调用异常,ACTION:" + json.action);
        console.log(result);
        // var ret_data = "";
        ret_data = window.atob(result.retustring);
        // $("#fin").attr('src', "data:image/bmp;base64," + ret_data);
      }
    });
    return code;
  }

  function getpdf(json) {
    var code = 0;
    //请求加上时间戳，防止浏览器缓存
    $.ajax({
      // url: "http://127.0.0.1:28810/" + "'" + codeHandler.encode(JSON.stringify(json), 'base64') + "'" + "time=" + Date.parse(new Date()),
      success: function(result) {
        if (!result.ret.startWith("00")) {
          code = -1;
          alert("Windows服务返回错误,ACTION:" + json.action);
        } else {
          var ret_data = ""
          ret_data = result.retCode;
          alert("成功生成!");
          //$("#fin").attr('src',"data:image/bmp;base64," + ret_data);
        }
      },
      error: function(result) {
        code = -1;
        alert("Windows服务调用异常,ACTION:" + json.action);
        console.log(result);
        var ret_data = "";
        ret_data = window.atob(result.retustring);
        $("#fin").attr('src', "data:image/bmp;base64," + ret_data);
        //attr('src',''+encodedata)
      }
    });
    return code;
  }

  function popSig() {
    var json;
    try {
      //开始签名
      json = { "action": "01" };
      myAjax(json);
    } catch (e) {
      alert("数据采集失败");
      console.log(e);
    }
  }

  function showsig() {
    //获取图片
    var ret_data = "";
    var json = { "action": "01" };
    $.ajax({
      url: "http://127.0.0.1:28810/" + "'" + window.btoa(JSON.stringify(json)) + "'" + "time=" + Date.parse(new Date()),
      success: function(result) {
        console.log(result);
        if (!result.ret.startWith("00")) {
          alert("获取图片失败");
        } else {
          ret_data = window.atob(result.retustring);
          $("#fin").attr('src', "data:image/bmp;base64," + ret_data);
        }
      },
      error: function(result) {
        alert("Windows服务发生错误");
      }
    });
  }

  function popFP() {
    var json;
    try {
      //弹出采集界面
      json = { "action": "06" };
      if (myAjax(json) != 0) return;
      //设置采集图片大小
      // showfp();
    } catch (e) {
      // alert("数据采集失败");
      console.log(e);
    }
  }

  function showfp() {
    // 获取图片
    var ret_data = "";
    var json = { "action": "02", "variable1": "01" };
    $.ajax({
      url: "http://localhost:28810/" + "'" + window.btoa(JSON.stringify(json)) + "'" + "time=" + Date.parse(new Date()),
      success: function(result) {
        console.log(result);
        if (!result.ret.startWith("00")) {
          alert("获取图片失败");
        } else {
          ret_data = window.atob(result.retustring);
          $("#fin").attr('src', "data:image/bmp;base64," + ret_data);
        }
      },
      error: function(result) {
        alert("Windows服务发生错误");
      }
    });
  }

  function popFPSig() {
    //获取图片
    var json;
    try {
      json = { "action": "03" };
      myAjax(json);
    } catch (e) {
      alert("数据采集失败");
      console.log(e);
    }
  }

  // 获取签名和指纹
  // $("#getSignature").click(function() {
  //   var json;
  //     try {
  //       //开始签名
  //       json = { "action": "01" };
  //       myAjax(json);
  //     } catch (e) {
  //       alert("数据采集失败");
  //       console.log(e);
  //     }
  // })
  $("#getSignature").click(function() {
    // document.getElementById('pageNumber').addEventListener('click', onNextPage);
    // alert("当前页码："+document.getElementById('pageNumber').value);
    // alert("总页码："+document.getElementById('numPages').innerHTML);
    // var pageNumber=document.getElementById('pageNumber').value;
    // var numPages=document.getElementById('numPages').innerHTML;
    // 从页面的左下角（br=botom right）为原点计算偏移量，笛卡尔坐标系
    var offsetdx = "-30"
    var offsetdy = "40"
    //缩放比例
    var scale = "1.0"
    // return;
    popSig();
    // popFP();
    // popFPSig();
    alert(ret_data)
    // $.ajax({
    //   async: true,
    //   url: "{{.Engineercmsapi_url}}/v1/pdfcpu/test/" + "'" + codeHandler.encode(JSON.stringify(json), 'base64') + "'" + "time=" + Date.parse(new Date()),
    //   success: function(result) {
    //       alert("解析图片数据！");
    //       ret_data = result.bmpdata;
    //   },
    //   error: function(result) {
    //     code = -1;
    //     alert("Windows服务调用异常,ACTION:" + json.action);
    //     console.log(result);
    //     ret_data = window.atob(result.retustring);
    //   }
    // });
    docEditor.insertImage({
      "fileType": "png",
      "url": "data:image/bmp;base64," + ret_data
      // "url": "{{.Engineercmsapi_url}}/attachment/onlyoffice/signature/2020March/1584852817365874800.png"
    });
    // pdf合并签名，成功后reload
    // $.ajax({
    //     cache: false,
    //     type: 'POST',
    //     url: "http://127.0.0.1/v1/pdfcpu/addwatermarks/"+{{.DocId}},
    //     async: false,
    //     dataType: 'json',
    //     data: { image: ret_data, pageNumber: pageNumber, numPages: numPages, offsetdx:offsetdx, offsetdy:offsetdy, scale:scale },
    //     success: function (data) {
    //       alert("签名+手印 成功!");
    //       window.location.reload(true);
    //     },
    //     error: function(data) {
    //       alert("签名失败！");
    //       console.log(result);
    //     }
    //   });
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
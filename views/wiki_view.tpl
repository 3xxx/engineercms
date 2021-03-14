<!DOCTYPE html>
{{template "header"}}
<title>Wiki详细</title>
  <meta http-equiv="Content-Type" content="text/html;charset=utf-8"/>
  <link rel="stylesheet" href="/static/froala/css/froala_editor.css">

  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  <!-- <link rel="stylesheet" href="/static/froala/css/froala_style.css"> -->
  <link rel="stylesheet" href="/static/froala/css/plugins/code_view.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/draggable.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/colors.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/emoticons.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/image_manager.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/image.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/line_breaker.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/table.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/char_counter.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/video.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/fullscreen.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/file.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/quick_insert.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/help.css">
  <link rel="stylesheet" href="/static/froala/css/third_party/spell_checker.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/special_characters.css">
  <link rel="stylesheet" href="/static/froala/js/codemirror.min.css">
  <link rel="stylesheet" href="/static/froala/css/themes/red.css">

<link rel="stylesheet" href="/static/froala/css/froala_style.css">

  <style type="text/css">
    /*body {
          text-align: center;
      }*/
      /*div#editor {
          width: 81%;
          margin: auto;
          text-align: left;
      }
      .ss {
        background-color: red;
      }*/
  </style>
</head>

<!-- <div class="navbar navbar-default navbar-static-top"> -->
  <div class="container-fill">{{template "navbar" .}}</div>
  
<body>
<div class="col-md-8 col-md-offset-2">
  <h2>
    {{.Wiki.Title}}
    <!--下面这个.Tid是Wiki.go的view里直接传过来的 target="_blank"-->
    <a href="/wiki/modify?tid={{.Tid}}" class="btn btn-default">修改wiki</a>
  </h2>
  <div class="fr-view">
    {{str2html .Wiki.Content}}
  </div>
</div>

<div id="editor" class="col-md-8 col-md-offset-2">
  {{range .Replies}}
    <h4><font size="4" color="#DC143C">
      {{.Name}}</font>
      <small>{{dateformat .Created "2006-01-02"}}</small>
      {{if $.isLogin}}
      <a href="/reply/delete?tid={{$.tid}}&rid={{.Id}}">删除</a>
      {{end}}
    </h4>
    <div class="content">
    {{str2html .Content}}
    </div>
  {{end}}

  <h4>本文回复</h4>
  <!-- <form method="post" action="/reply/addwiki?op=c"> -->
    <!-- <input type="hidden" name="tid" value="{{.Wiki.Id}}"> -->
    <!-- <div class="form-group"> -->
      <label>显示昵称：</label>
      <input type="text" class="form-control" id="nickname" name="nickname" value="{{.Username}}">
    <!-- </div> -->
      <label>内容：</label>
        <div id='edit' style="margin-top: 10px;">
        </div>  
        <br> 
    <button class="btn btn-primary" onclick="return checkInput();" style="float:right">提交回复</button>
    <br>
  <!-- </form> -->
  <br>
  <br>
</div>

<!-- <script type="text/javascript" src="/static/froala/js/jquery.min.1.11.0.js"></script> -->
  <script type="text/javascript" src="/static/froala/js/codemirror.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/xml.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/froala_editor.min.js" ></script>
  <script type="text/javascript" src="/static/froala/js/plugins/align.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/char_counter.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/code_beautifier.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/code_view.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/colors.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/draggable.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/emoticons.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/entities.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/file.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/font_size.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/font_family.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/fullscreen.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/image.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/image_manager.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/line_breaker.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/inline_style.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/link.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/lists.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/paragraph_format.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/paragraph_style.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/quick_insert.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/quote.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/table.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/save.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/url.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/video.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/help.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/print.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/third_party/spell_checker.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/special_characters.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/word_paste.min.js"></script>
  <script src="/static/froala/js/languages/zh_cn.js"></script>
  <script>
    // $(function(){
    //   $('#edit').froalaEditor()
    // });
$(function (){
  //超大屏幕
  var toolbarButtons   = ['fullscreen', 'bold', 'italic', 'underline', 'strikeThrough', 'subscript', 'superscript', 'fontFamily', 'fontSize', '|', 'color', 'emoticons', 'inlineStyle', 'paragraphStyle', '|', 'paragraphFormat', 'align', 'formatOL', 'formatUL', 'outdent', 'indent', 'quote', 'insertHR', '-', 'insertLink', 'insertImage', 'insertVideo', 'insertFile', 'insertTable', 'undo', 'redo', 'clearFormatting', 'selectAll', 'html'];
  //大屏幕
  var toolbarButtonsMD = ['fullscreen', 'bold', 'italic', 'underline', 'fontFamily', 'fontSize', 'color', 'paragraphStyle', 'paragraphFormat', 'align', 'formatOL', 'formatUL', 'outdent', 'indent', 'quote', 'insertHR', 'insertLink', 'insertImage', 'insertVideo', 'insertFile', 'insertTable', 'undo', 'redo', 'clearFormatting'];
  //小屏幕
  var toolbarButtonsSM = ['fullscreen', 'bold', 'italic', 'underline', 'fontFamily', 'fontSize', 'insertLink', 'insertImage', 'insertTable', 'undo', 'redo'];
  //手机
  var toolbarButtonsXS = ['bold', 'italic', 'fontFamily', 'fontSize', 'undo', 'redo'];
  var pid = $('#pid').val();
  //编辑器初始化并赋值 
  $('#edit').froalaEditor({
      placeholderText: '请输入内容',
      charCounterCount       : true,//默认
      // charCounterMax         : -1,//默认
      saveInterval            : 0,//不自动保存，默认10000
      // theme                    : "red",
      height                   : "300px",
      toolbarBottom           : false,//默认
      toolbarButtonsMD        : toolbarButtonsMD,
      toolbarButtonsSM        : toolbarButtonsSM,
      toolbarButtonsXS        : toolbarButtonsXS,
      toolbarInline           : false,//true选中设置样式,默认false
      imageUploadMethod       : 'POST',
      heightMin: 450,
      charCounterMax: 3000,
      // imageUploadURL: "uploadImgEditor",
      imageParams: { postId: "123" },
      params: {
            acl: '01',
            AWSAccessKeyId: '02',
            policy: '03',
            signature: '04',
          },
      autosave: true,
      autosaveInterval: 2500,
      saveURL: 'hander/FroalaHandler.ashx',
      saveParams: { postId: '1'},
      spellcheck: false,
      imageUploadURL: '/uploadwikiimg',//上传到本地服务器
      imageUploadParams: {pid: '{{.Id}}'},
      imageDeleteURL: 'lib/delete_image.php',//删除图片
      imagesLoadURL: 'lib/load_images.php',//管理图片
      videoUploadURL:'/uploadwikiimg',
      videoUploadParams: {pid: '{{.Id}}'},
      fileUploadURL: '/uploadwikiimg',
      fileUploadParams: {pid: '{{.Id}}'},
      enter: $.FroalaEditor.ENTER_BR,
      language: 'zh_cn',
      // toolbarButtons: ['bold', 'italic', 'underline', 'paragraphFormat', 'align','color','fontSize','insertImage','insertTable','undo', 'redo']
      });
    })
  </script>
<script>
 function checkInput(){
    //是下面这段代码出了问题，等下修改
      var name=document.getElementById("nickname");
      var html = $('div#edit').froalaEditor('html.get');
      if (name.value.length==0){
        alert("请输入昵称");
        return false;
      }
      if (html==""){
        alert("请输入回复内容");
        return false;
      }
      var tid = {{.Wiki.Id}};
      $.ajax({
        type:"post",
        url:"/reply/addwiki",
        data: {tid:tid,content:html},
        success:function(data,status){
          alert("添加成功！");
          window.location.href="/wiki/view/"+data;//刷新页面
        },
      }); 
    }
</script>
</body>
</html>
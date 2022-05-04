<!DOCTYPE html>
{{template "header"}}
  <title>添加wiki</title>
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
  <meta http-equiv="Cache-Control" content="no-siteapp">
  <!-- <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script> -->
  <!-- <script type="text/javascript" src="/static/js/bootstrap.min.js"></script> -->
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/> -->
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>

  <link rel="stylesheet" href="/static/froala2.8.4/css/froala_editor.css">
  <link rel="stylesheet" href="/static/froala2.8.4/css/froala_style.css">
  <link rel="stylesheet" href="/static/froala2.8.4/css/plugins/code_view.css">
  <link rel="stylesheet" href="/static/froala2.8.4/css/plugins/draggable.css">
  <link rel="stylesheet" href="/static/froala2.8.4/css/plugins/colors.css">
  <link rel="stylesheet" href="/static/froala2.8.4/css/plugins/emoticons.css">
  <link rel="stylesheet" href="/static/froala2.8.4/css/plugins/image_manager.css">
  <link rel="stylesheet" href="/static/froala2.8.4/css/plugins/image.css">
  <link rel="stylesheet" href="/static/froala2.8.4/css/plugins/line_breaker.css">
  <link rel="stylesheet" href="/static/froala2.8.4/css/plugins/table.css">
  <link rel="stylesheet" href="/static/froala2.8.4/css/plugins/char_counter.css">
  <link rel="stylesheet" href="/static/froala2.8.4/css/plugins/video.css">
  <link rel="stylesheet" href="/static/froala2.8.4/css/plugins/fullscreen.css">
  <link rel="stylesheet" href="/static/froala2.8.4/css/plugins/file.css">
  <link rel="stylesheet" href="/static/froala2.8.4/css/plugins/quick_insert.css">
  <link rel="stylesheet" href="/static/froala2.8.4/css/plugins/help.css">
  <!-- <link rel="stylesheet" href="/static/froala2.8.4/css/third_party/spell_checker.css"> -->
  <link rel="stylesheet" href="/static/froala2.8.4/css/plugins/special_characters.css">
  <link rel="stylesheet" href="/static/froala2.8.4/js/codemirror.min.css">
  <link rel="stylesheet" href="/static/froala2.8.4/css/themes/red.css">

<style type="text/css">
  img{max-width:100%}
  body {  
    overflow-y: scroll !important;  
  }  
  .view {  
   word-break: break-all;  
  }  
  .vote_area {  
   display: block;  
  }  
  .vote_iframe {  
   background-color: transparent;  
   border: 0 none;  
   height: 100%;  
  }  
  #content_imagescale{display:none !important;} /*去除点击图片后出现的拉伸边框*/
    video{width:100%; height:100%; object-fit: fill}
    #main .article_next_prev span {
          width: 80px;
          margin: 0px;
          padding: 0px;
          padding-left: 30px;
  }

    i#delete{color:#DC143C;}
  </style>
</head>

<!-- <div class="navbar navbar-default navbar-static-top"> -->
  <div class="container-fill">{{template "navbar" .}}</div>

<body>
<div id="editor" class="col-md-8 col-md-offset-2">
  <h2>添加文章</h2>
  <!-- <form> -->
    <div class="form-group">
    <label>文章名称:</label>
    <input type="text" class="form-control" id="title" placeholder="Enter Name" name="title"></div>
    <label>图文正文:</label>
    <textarea id='edit' style="margin-top: 30px;" placeholder="Type some text">
    </textarea>
    <br />
    <!-- <button type="submit" class="btn btn-primary" onclick="return checkInput();"> 添  加 </button> -->
    <!--必须加return才能不跳转-->
    <button type="button" class="btn btn-primary" style="float:right" onclick="addwiki()">保存</button>
      <br />
  <!-- </form> -->
  <br />
</div>
<br />

<!-- <script type="text/javascript" src="/static/froala2.8.4/js/jquery.min.1.11.0.js"></script> -->
  <script type="text/javascript" src="/static/froala2.8.4/js/codemirror.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/xml.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/froala_editor.min.js" ></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/align.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/char_counter.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/code_beautifier.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/code_view.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/colors.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/draggable.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/emoticons.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/entities.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/file.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/font_size.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/font_family.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/fullscreen.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/image.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/image_manager.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/line_breaker.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/inline_style.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/link.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/lists.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/paragraph_format.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/paragraph_style.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/quick_insert.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/quote.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/table.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/save.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/url.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/video.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/help.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/print.min.js"></script>
  <!-- <script type="text/javascript" src="/static/froala2.8.4/js/third_party/spell_checker.min.js"></script> -->
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/special_characters.min.js"></script>
  <script type="text/javascript" src="/static/froala2.8.4/js/plugins/word_paste.min.js"></script>
  <script src="/static/froala2.8.4/js/languages/zh_cn.js"></script>
  <script>
    // $(function(){
    //   $('#edit').froalaEditor()
    // });
$(function (){
  //超大屏幕
  var toolbarButtons = ['fullscreen', 'bold', 'italic', 'underline', 'strikeThrough', 'subscript', 'superscript', 'fontFamily', 'fontSize', '|', 'color', 'emoticons', 'inlineStyle', 'paragraphStyle', '|', 'paragraphFormat', 'align', 'formatOL', 'formatUL', 'outdent', 'indent', 'quote', 'insertHR', '-', 'insertLink', 'insertImage', 'insertVideo', 'insertFile', 'insertTable', 'undo', 'redo', 'clearFormatting', 'selectAll', 'html'];
  //大屏幕
  var toolbarButtonsMD = ['fullscreen', 'bold', 'italic', 'underline', 'fontFamily', 'fontSize', 'color', 'paragraphStyle', 'paragraphFormat', 'align', 'formatOL', 'formatUL', 'outdent', 'indent', 'quote', 'insertHR', 'insertLink', 'insertImage', 'insertVideo', 'insertFile', 'insertTable', 'undo', 'redo', 'clearFormatting'];
  //小屏幕
  var toolbarButtonsSM = ['fullscreen', 'bold', 'italic', 'underline', 'fontFamily', 'fontSize', 'insertLink', 'insertImage', 'insertTable', 'undo', 'redo'];
  //手机
  var toolbarButtonsXS = ['bold', 'italic', 'fontFamily', 'fontSize', 'undo', 'redo'];
  var pid = $('#pid').val();
  //编辑器初始化并赋值 
  $('#edit')
  // .on('froalaEditor.initialized', function (e, editor) {
  //       $('#edit').parents('form').on('submit', function () {
  //         // console.log($('#edit').val());
  //         var title = $('#title').val();
  //         $.ajax({
  //           type:"post",
  //           url:"/wiki/addwiki",
  //           data: {title:title,content:$('#edit').val()},
  //           success:function(data,status){
  //             alert("添加成功！");
  //             window.location.href="/wiki";//刷新页面
  //           },
  //         });
  //         // return false;
  //       })
  //     })
  .froalaEditor({
      // enter: $.FroalaEditor.ENTER_P,
      placeholderText: '请输入内容',
      charCounterCount       : true,//默认
      // charCounterMax         : -1,//默认
      saveInterval            : 0,//不自动保存，默认10000
      // theme                    : "red",
      height                   : "550px",
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
      // imageUploadParams: {pid: '{{.product.ProjectId}}'},
      imageDeleteURL: 'lib/delete_image.php',//删除图片
      imagesLoadURL: 'lib/load_images.php',//管理图片
      videoUploadURL:'/uploadwikiimg',
      // videoUploadParams: {pid: '{{.Id}}'},
      fileUploadURL: '/uploadwikiimg',
      // fileUploadParams: {pid: '{{.Id}}'},
      enter: $.FroalaEditor.ENTER_BR,
      language: 'zh_cn',
      // toolbarButtons: ['bold', 'italic', 'underline', 'paragraphFormat', 'align','color','fontSize','insertImage','insertTable','undo', 'redo']
      });
    })
  </script>
<script type="text/javascript">
   function addwiki(){
    //是下面这段代码出了问题，等下修改
      var name=document.getElementById("title");
      if (name.value.length==0){
        alert("请输入文章名称");
        return false;
      }else{
        var title = $('#title').val();
        $.ajax({
          type:"post",
          url:"/wiki/addwiki",
          data: {title:title,content:$('#edit').val()},
          success:function(data,status){
            alert("添加成功！");
            window.location.href="/wiki/view/"+data;//刷新页面
          },
        });
      }
    }
</script>
</body>
</html> 
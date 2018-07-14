<!DOCTYPE html>
<title>{{.product.Title}}</title>
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
  <meta http-equiv="Cache-Control" content="no-siteapp">
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>

  <link rel="stylesheet" href="/static/froala/css/froala_editor.css">
  <link rel="stylesheet" href="/static/froala/css/froala_style.css">
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
</style>
</head>
<body>
<div id="editor" class="col-xs-12 col-sm-12 col-lg-8 col-md-8 col-md-offset-2">
  <h3>编辑文章</h3>
  <!-- <form> -->
    <!-- <div class="form-group">
      <label>编号:</label>
      <input type="text" class="form-control" readonly="readonly" placeholder="Enter account" value="{{.product.ProjectId}}">
    </div> -->
    <div class="form-group">
      <label>标题:</label>
      <input type="text" class="form-control" readonly="readonly"  placeholder="Enter account" value="{{.product.Title}}"></div>
    <div class="form-group">
      <label>副标题:</label>
      <input id="subtext" name="subtext" type="text" class="form-control"  placeholder="Enter Subtext" value="{{.article.Subtext}}"></div>
    <!-- <div class="form-group"> -->
      <!-- <label>关键字:</label>
      <input type="text" class="form-control" readonly="readonly" value="{{.product.Label}}"></div>
      <div class="form-group"> -->
      <!-- <label>设计:</label>
      <input type="text" class="form-control" readonly="readonly"  placeholder="Enter account" value="{{.product.Principal}}"></div> -->
      <label>文章内容:</label>
        <textarea id='edit' style="margin-top: 30px;" placeholder="Type some text">
          {{str2html .article.Content}}
        </textarea>
        <br />
      <!-- <button type="submit" class="btn btn-primary" style="float:right">提交修改</button> -->
      <button type="button" class="btn btn-primary" style="float:right" onclick="updatearticle()">保存</button>
      <br />
  <!-- </form> -->
  <br />
</div>
<br />
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
  //         var articleid = {{.article.Id}};
  //         var subtext = $('#subtext').val();
  //         $.ajax({
  //           type:"post",
  //           url:"/project/product/updatearticle",
  //           data: {aid:articleid,subtext:subtext,content:$('#edit').val()},
  //           success:function(data,status){
  //             alert("修改成功！");
  //             // window.location.reload();//刷新页面
  //             window.location.href="/project/product/article/"+{{.article.Id}};
  //           },
  //         });
  //         // return false;
  //       })
      // })
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
      imageUploadURL: '/uploadimg',//上传到本地服务器
      imageUploadParams: {pid: '{{.product.ProjectId}}'},
      imageDeleteURL: 'lib/delete_image.php',//删除图片
      imagesLoadURL: 'lib/load_images.php',//管理图片
      videoUploadURL:'/uploadvideo',
      videoUploadParams: {pid: '{{.product.ProjectId}}'},
      fileUploadURL: '/uploadimg',
      fileUploadParams: {pid: '{{.product.ProjectId}}'},
      enter: $.FroalaEditor.ENTER_BR,
      language: 'zh_cn',
      // toolbarButtons: ['bold', 'italic', 'underline', 'paragraphFormat', 'align','color','fontSize','insertImage','insertTable','undo', 'redo']
      });
    })
  </script>

<script type="text/javascript">
  $(function(){//这个和$(document).ready(function()等价
    value={{.product.Label}};
    array=value.split(",");
    var labelarray = new Array(); 
    for (i=0;i<array.length;i++){
      var th1="&nbsp;<a href='/project/product/article/keysearch?keyword="+ array[i]+"'><span class='btn btn-info btn-xs'>" +array[i] + "</span></a>";
      $("small#publish").after(th1);
    }
  })

  //编辑文章
  function updatearticle(){
    var articleid = {{.article.Id}};
    var subtext = $('#subtext1').val();
    var html = $('#edit').val();
      $.ajax({
        type:"post",
        url:"/project/product/updatearticle",
        data: {aid:articleid,subtext:subtext,content:html},//父级id
        success:function(data,status){
          alert("修改“"+data+"”成功！(status:"+status+".)");
          window.location.href="/project/product/article/"+{{.article.Id}};
        },
      });
  }
</script>
</body>
</html>

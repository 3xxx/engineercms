<!DOCTYPE html>
<title>{{.product.Title}}</title>
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<meta http-equiv="Cache-Control" content="no-siteapp">
<script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css" />
<link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />

<!-- <link rel="stylesheet" href="/static/froala/css/froala_editor.css">
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
<link rel="stylesheet" href="/static/froala/css/plugins/special_characters.css"> -->
<link rel="stylesheet" href="/static/froala/css/codemirror.min.css">
<!-- <link rel="stylesheet" href="/static/froala/css/themes/red.css"> -->

  <link rel="stylesheet" href="/static/froala/css/froala_editor.min.css">
  <link rel="stylesheet" href="/static/froala/css/froala_style.min.css">
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

<style type="text/css">
img {
  max-width: 100%
}

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

#content_imagescale {
  display: none !important;
}

/*去除点击图片后出现的拉伸边框*/
video {
  width: 100%;
  height: 100%;
  object-fit: fill
}

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
    <h3>编辑模型指标</h3>
    <!-- <form> -->
    <div class="form-group">
      <label>编号:</label>
      <input type="text" class="form-control" readonly="readonly" placeholder="Enter Title" value="{{.FreeCADModel.Number}}{{.FreeCADModel.TitleB}}{{.FreeCADModel.Version}}">
    </div>
    <div class="form-group">
      <label>指标:</label>
      <input id="indicators" name="indicators" type="text" class="form-control" placeholder="Enter indicators" value="{{.FreeCADModel.Indicators}}"></div>
    <!-- <div class="form-group">
      <label>描述：</label>
      <input id="description" name="description" type="text" class="form-control" placeholder="Enter description" value="{{.FreeCADModel.Description}}"></div> -->
    <div class="form-group">
      <label>描述：</label>
      <textarea class="form-control" id="description" rows="3" placeholder="Enter description">{{.FreeCADModel.Description}}</textarea></div>

    <!-- <div class="form-group"> -->
    <!-- <label>关键字:</label>
      <input type="text" class="form-control" readonly="readonly" value="{{.product.Label}}"></div>
      <div class="form-group"> -->

    <label>模型描述:</label>
    <div id="editor">
      <div id='edit' style="margin-top: 30px;">
        {{str2html .FreeCADModel.Description}}
      </div>
    </div>

    <br />
    <!-- <button type="submit" class="btn btn-primary" style="float:right">提交修改</button> -->
    <button type="button" class="btn btn-primary" style="float:right" onclick="updateFreeCAD()">更新update</button>
    <br />
    <!-- </form> -->
    <br />
  </div>
  <br />

  <script type="text/javascript" src="/static/froala/js/froala_editor.min.js"></script>
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
  <!-- <script type="text/javascript" src="/static/froala/js/third_party/spell_checker.min.js"></script> -->
  <script type="text/javascript" src="/static/froala/js/plugins/special_characters.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/word_paste.min.js"></script>
  <script src="/static/froala/js/languages/zh_cn.js"></script>
  <script>

  $(function() {
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
    // $('#edit').froalaEditor({
    new FroalaEditor("#edit", {
      // enter: $.FroalaEditor.ENTER_P,
      placeholderText: '请不要在这里输入内容',
      charCounterCount: true, //默认
      // charCounterMax         : -1,//默认
      saveInterval: 0, //不自动保存，默认10000
      // theme                    : "red",
      height: "550px",
      toolbarBottom: false, //默认
      toolbarButtonsMD: toolbarButtonsMD,
      toolbarButtonsSM: toolbarButtonsSM,
      toolbarButtonsXS: toolbarButtonsXS,
      toolbarInline: false, //true选中设置样式,默认false
      imageUploadMethod: 'POST',
      heightMin: 450,
      charCounterMax: 3000,
      // imageUploadURL: "uploadImgEditor",
      imageParams: { postId: '{{.FreeCADModel.ID}}' },
      params: {
        acl: '01',
        AWSAccessKeyId: '02',
        policy: '03',
        signature: '04',
      },
      autosave: true,
      autosaveInterval: 2500,
      saveURL: 'hander/FroalaHandler.ashx',
      saveParams: { postId: '1' },
      spellcheck: false,
      imageUploadURL: '/uploadimg', //上传到本地服务器
      imageUploadParams: { pid: '{{.FreeCADModel.ID}}' },
      imageDeleteURL: 'lib/delete_image.php', //删除图片
      imagesLoadURL: 'lib/load_images.php', //管理图片
      videoUploadURL: '/uploadvideo',
      videoUploadParams: { pid: '{{.FreeCADModel.ID}}' },
      fileUploadURL: '/uploadimg',
      fileUploadParams: { pid: '{{.FreeCADModel.ID}}' },
      // enter: $.FroalaEditor.ENTER_BR,
      language: 'zh_cn',
      // toolbarButtons: ['bold', 'italic', 'underline', 'paragraphFormat', 'align','color','fontSize','insertImage','insertTable','undo', 'redo']
    });
  })
  </script>
  <script type="text/javascript">

  //编辑文章
  function updateFreeCAD() {
    var freecadmodelid = {{.FreeCADModel.ID }};
    var indicators = $('#indicators').val();
    var description = $('#description').val();
    // var html = $('#edit').val();
    let editor = new FroalaEditor('#edit', {}, function () {
    });
    var html =editor.html.get()
    $.ajax({
      type: "post",
      url: "/v1/freecad/updatefreecad",
      data: { id: freecadmodelid, indicators: indicators, description: description }, //父级id
      success: function(data, status) {
        alert("修改“" + data.data + "”成功！(status:" + status + ".)");
        // window.location.href = "/project/product/article/" + {{.FreeCADModel.ID }};
      },
    });
  }
  </script>
</body>

</html>
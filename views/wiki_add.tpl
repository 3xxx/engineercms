<!DOCTYPE html>
{{template "header"}}
<title>添加wiki</title>
<meta http-equiv="Content-Type" content="text/html;charset=utf-8"/>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.min.js"> </script>
    <!--建议手动加在语言，避免在ie下有时因为加载语言失败导致编辑器加载失败-->
    <!--这里加载的语言文件会覆盖你在配置项目里添加的语言类型，比如你在配置项目里配置的是英文，这里加载的中文，那最后就是中文-->
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>
<!--引入CSS-->
<link rel="stylesheet" type="text/css" href="/static/fex-team-webuploader/css/webuploader.css">
<!-- <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script> -->
<!--引入JS-->
<script type="text/javascript" src="/static/fex-team-webuploader/dist/webuploader.js"></script>
  <script src="/static/js/bootstrap-treeview.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-treeview.css"/>
  <style>
    i#delete{color:#DC143C;}
  </style>
</head>

<!-- <div class="navbar navbar-default navbar-static-top"> -->
  <div class="container-fill">{{template "navbar" .}}</div>

<body>
<div class="col-md-8 col-md-offset-2">
  <h2>添加文章</h2>
  <form method="post" action="/wiki/addwiki" enctype="multipart/form-data">
    <div class="form-group">
    <label>文章名称:</label>
    <input type="text" class="form-control" id="title" placeholder="Enter Name" name="title"></div>
<label>图文正文:</label>
<div>
    <script id="container" type="text/plain" style="height:300px;"></script><!-- width:1024px; -->
</div>
<br />
  <button type="submit" class="btn btn-primary" onclick="return checkInput();"> 添  加 </button>
  <!--必须加return才能不跳转-->
</form>

<br />
<br />
</div>


<script type="text/javascript">
   function checkInput(){
    //是下面这段代码出了问题，等下修改
      var name=document.getElementById("title");
      if (name.value.length==0){
        alert("请输入文章名称");
        return false;
      }
    }

 //实例化编辑器
    //议使用工厂方法getEditor创建和引用编辑器实例，如果在某个闭包下引用该编辑器，直接调用UE.getEditor('editor')就能拿到相关的实例
    // var ue = UE.getEditor('container');
    var ue = UE.getEditor('container', {
    // toolbars: [
    //     ['fullscreen', 'source', 'undo', 'redo', 'bold']
    // ],
    // toolbars: [[
    //         'fullscreen', 'source', '|', 'undo', 'redo', '|',
    //         'bold', 'italic', 'underline', 'fontborder', 'strikethrough', 'superscript', 'subscript', 'removeformat', 'formatmatch', 'autotypeset', 'blockquote', 'pasteplain', '|', 'forecolor', 'backcolor', 'insertorderedlist', 'insertunorderedlist', 'selectall', 'cleardoc', '|',
    //         'rowspacingtop', 'rowspacingbottom', 'lineheight', '|',
    //         'customstyle', 'paragraph', 'fontfamily', 'fontsize', '|',
    //         'directionalityltr', 'directionalityrtl', 'indent', '|',
    //         'justifyleft', 'justifycenter', 'justifyright', 'justifyjustify', '|', 'touppercase', 'tolowercase', '|',
    //         'link', 'unlink', 'anchor', '|', 'imagenone', 'imageleft', 'imageright', 'imagecenter', '|',
    //         'emotion', 'map', 'gmap', 'insertframe', 'insertcode', 'webapp', 'pagebreak', 'template', 'background', '|',
    //         'horizontal', 'date', 'time', 'spechars', 'wordimage', '|',
    //         'inserttable', 'deletetable', 'insertparagraphbeforetable', 'insertrow', 'deleterow', 'insertcol', 'deletecol', 'mergecells', 'mergeright', 'mergedown', 'splittocells', 'splittorows', 'splittocols', 'charts', '|',
    //         'print', 'preview', 'searchreplace', 'help', 'drafts'
    //     ]],
    autoHeightEnabled: true,
    // autoFloatEnabled: true
});
/* 2.传入参数表,添加到已有参数表里 通过携带参数，实现不同的页面使用不同controllers*/
    ue.ready(function () {
    ue.addListener('focus', function () {//startUpload start-upload startUpload beforeExecCommand是在插入图片之前触发
    var title = $('#title').val();
    var html = ue.getContent();
    ue.execCommand('serverparam', {
        "key":"wiki",
        "title":title,
        'content':html, 
    });
});
});

</script>
    </body>
</html> 
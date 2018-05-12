<!DOCTYPE html>
{{template "header"}}
<title>Wiki详细</title>
    <meta http-equiv="Content-Type" content="text/html;charset=utf-8"/>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.min.js"> </script>
    <!--建议手动加在语言，避免在ie下有时因为加载语言失败导致编辑器加载失败-->
    <!--这里加载的语言文件会覆盖你在配置项目里添加的语言类型，比如你在配置项目里配置的是英文，这里加载的中文，那最后就是中文-->
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>
    <script src="/static/ueditor/ueditor.parse.min.js"></script>
    <!-- <link href="/static/ueditor/third-party/video-js/video1-js.min.css" rel="stylesheet"> -->
    <!-- <script src="/static/ueditor/third-party/video-js/video.min.js"></script> -->
  <style type="text/css">
    img{max-width:100%}
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
    <div class="content">
  {{str2html .Wiki.Content}}
  </div>
<!--   <tr>
    成果附件：
  <ol>
    {{range .Attachment}}
    <th><li> <a href={{.Route}}  target="_blank">{{.FileName}}</a> 文件大小：{{.FileSize}}KB<a href="/attachment/delete?aid={{.Id}}&tid={{$.Tid}}">删除</a></li></th>
    {{end}}
  </ol>
  </tr> -->
</div>

<div class="col-md-8 col-md-offset-2">
  <!-- {{$tid:=.Wiki.Id}} -->
  <!-- {{$isLogin:=.IsLogin}} -->
{{range .Replies}}
  <h4><font size="4" color="#DC143C">
    {{.Name}}</font>
    <small>{{dateformat .Created "2006-01-02"}}</small>
    {{if $isLogin}}
    <a href="/reply/delete?tid={{$tid}}&rid={{.Id}}">删除</a>
    {{end}}
  </h4>
  <!-- {{.Content}} -->
  <div class="content">
  {{str2html .Content}}
  </div>
{{end}}
  <h4>本文回复</h4>
  <form method="post" action="/reply/addwiki?op=c">
    <input type="hidden" name="tid" value="{{.Wiki.Id}}">
    <div class="form-group">
      <label>显示昵称：</label>
      <input type="text" class="form-control" id="nickname" name="nickname" value="{{.Username}}"></div>
    <div class="form-group">
      <label>内容：</label>
      <div>
        <script id="container" type="text/plain" style="height:300px;width: 100%"></script><!-- width:1024px; -->
      </div>
      <!-- <textarea name="content" id="" cols="30" rows="10" class="form-control"></textarea> -->
    </div>
    <button class="btn btn-primary" onclick="return checkInput();" style="float:right">提交回复</button>
    <br>
  </form>
  <br>
  <br>
 <div>
 
</div> 
</div>


<script>
 function checkInput(){
    //是下面这段代码出了问题，等下修改
      var name=document.getElementById("nickname");
      var html = ue.getContent();
      if (name.value.length==0){
        alert("请输入昵称");
        return false;
      }
      if (html==""){
        alert("请输入回复内容");
        return false;
      }
    }

 //实例化编辑器
    //议使用工厂方法getEditor创建和引用编辑器实例，如果在某个闭包下引用该编辑器，直接调用UE.getEditor('editor')就能拿到相关的实例
    // var ue = UE.getEditor('container');
    var ue = UE.getEditor('container', {
    autoHeightEnabled: true,
    autoFloatEnabled: true
    });  
        
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

/* 2.传入参数表,添加到已有参数表里 通过携带参数，实现不同的页面使用不同controllers*/
//startUpload start-upload startUpload beforeExecCommand是在插入图片之前触发

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

        // 语法
        // uParse(selector,[option])
        /*
         selector支持
         id,class,tagName
         */
        /*
         目前支持的参数
         option:
         highlightJsUrl 代码高亮相关js的路径 如果展示有代码高亮，必须给定该属性
         highlightCssUrl 代码高亮相关css的路径 如果展示有代码高亮，必须给定该属性
         liiconpath 自定义列表样式的图标路径，可以不给定，默认'http://bs.baidu.com/listicon/',
         listDefaultPaddingLeft : 自定义列表样式的左边宽度 默认'20',
         customRule 可以传入你自己的处理规则函数，函数第一个参数是容器节点
         */

        uParse('.content',{
            rootPath : '/static/ueditor/'
        })
    </script>
</body>
</html>
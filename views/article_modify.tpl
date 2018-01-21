<!DOCTYPE html>
<title>{{.product.Title}}</title>
     <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta http-equiv="Content-Type" content="text/html;charset=utf-8"/>
    <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
    <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
    <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.min.js"> </script>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>
    <script src="/static/ueditor/ueditor.parse.min.js"></script>
    
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <meta http-equiv="Cache-Control" content="no-siteapp">
    <title>{{.product.Title}} - 珠三角设代</title>
    <!-- <link rel="Stylesheet" type="text/css" href="/static/css/csdn_style.css"> -->
    <!-- <link rel="stylesheet" href="/static/css/csdn_blog_detail.min.css"> -->
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
<div class="col-xs-12 col-sm-12 col-lg-8 col-md-8 col-md-offset-2">
  <h3>编辑文章</h3>
  <form  method="post" action="/project/product/updatearticle" enctype="multipart/form-data">
    <input type="hidden" id="pid" name="pid" value="{{.product.ProjectId}}">
    <input type="hidden" id="aid" name="aid" value="{{.article.Id}}">
    <div class="form-group">
      <label>编号:</label>
      <input type="text" class="form-control" readonly="readonly" placeholder="Enter account" value="{{.product.ProjectId}}"></div>
    <div class="form-group">
      <label>标题:</label>
      <input type="text" class="form-control" readonly="readonly"  placeholder="Enter account" value="{{.product.Title}}"></div>
      <div class="form-group">
      <label>副标题:</label>
      <input id="subtext" name="subtext" type="text" class="form-control"  placeholder="Enter Subtext" value="{{.article.Subtext}}"></div>
      <div class="form-group">
      <label>关键字:</label>
      <input type="text" class="form-control" readonly="readonly" value="{{.product.Label}}"></div>
      <div class="form-group">
      <label>设计:</label>
      <input type="text" class="form-control" readonly="readonly"  placeholder="Enter account" value="{{.product.Principal}}"></div>

      <label>文章内容:</label>
      <div id="content">
          <script id="article_container" type="text/plain" style="height:500px;width: 100%"></script>
      </div>
  <br />
  <button type="submit" class="btn btn-primary" style="float:right">提交修改</button>
  <br />  
 </form>
 <br />
  </div>

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
    //实例化编辑器
    //议使用工厂方法getEditor创建和引用编辑器实例，如果在某个闭包下引用该编辑器，直接调用UE.getEditor('editor')就能拿到相关的实例
    // var ue = UE.getEditor('container');
    var ue = UE.getEditor('article_container', {
      autoHeightEnabled: true,
      // autoFloatEnabled: true
    });
    ue.ready(function () {
        ue.addListener('focus', function () {//startUpload start-upload startUpload beforeExecCommand是在插入图片之前触发
            var pid = $('#pid').val();//productid
            // var html = ue.getContent();
            ue.execCommand('serverparam', {
              "pid":pid 
            });
        });
    });
    uParse('.content',{
      rootPath : '/static/ueditor/'
    })


    /* 2.传入参数表,添加到已有参数表里 通过携带参数，实现不同的页面使用不同controllers*/
    // ue.ready(function () {
    // ue.addListener('focus', function () {
    //   var tid = $('#tid').val();    
    //   var title = $('#title').val();
    //   var html = ue.getContent();
    //   ue.execCommand('serverparam', {
    //       "key":"wiki",
    //       "tid":tid,
    //       "title":title,
    //       'content':html, 
    //       });
    //   });
    // });

    $(function(){
        // var content =$('#content').val();
        //判断ueditor 编辑器是否创建成功
        ue.addListener("ready", function () {
        // editor准备好之后才可以使用
        ue.setContent({{str2html .article.Content}});
        });
    });

  //编辑文章
  function updatearticle(){
    // var radio =$("input[type='radio']:checked").val();
    var articleid = {{.article.Id}};//$('#cid').val();//文章id
    // var prodcode = $('#prodcode1').val();
    // var prodname = $('#prodname1').val();
    var subtext = $('#subtext1').val();
    // var prodprincipal = $('#prodprincipal1').val();
    // var prodlabel = $('#prodlabel1').val();
    var html = ue.getContent();
    // $('#myModal').on('hide.bs.modal', function () {  
    if (prodname&&prodcode){  
      $.ajax({
        type:"post",
        url:"/project/product/updatearticle",
        data: {pid:articleid,subtext:subtext,content:html},//父级id
        success:function(data,status){
          alert("修改“"+data+"”成功！(status:"+status+".)");
          window.location.reload();//刷新页面
        },
      });
    }else{
      alert("请填写编号和名称！");
      return;
    }
  }

</script>

</body>
</html>

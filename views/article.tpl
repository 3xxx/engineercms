<!DOCTYPE html>
<title>CMS系统</title>
<meta name="viewport" content="width=device-width, initial-scale=1">
    <meta http-equiv="Content-Type" content="text/html;charset=utf-8"/>
    <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
    <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
    <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.js"> </script>
    <!--建议手动加在语言，避免在ie下有时因为加载语言失败导致编辑器加载失败-->
    <!--这里加载的语言文件会覆盖你在配置项目里添加的语言类型，比如你在配置项目里配置的是英文，这里加载的中文，那最后就是中文-->
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>
    <script src="/static/ueditor/ueditor.parse.js"></script>
<style type="text/css">
img{max-width:100%}
/*nav {
  background: #4682B4;
}
h3 {line-height: 150%;
  color: #DC143C;#000000
  color:#FFFFFF;
  background: #4682B4;#fff
  margin: 0;
  padding: 0;
  font-family: Georgia, Palatino, serif;
  }
  h3 a {color: inherit;}h3标签里面的a标签继承父级颜色

   h4{
      color: #DC143C;
  font-family: georgia, palatino, "New Century Schoolbook",
  times, serif;
  font-weight: normal;
  /*font-size: 2em;
  margin-top: 1em;
  margin-bottom: 0;
  }*/

</style>
</head>

<body>
<div class="col-lg-8 col-md-offset-2">
<div class="page-header">
  <h2>{{.product.Code}}-{{.product.Title}} <small>{{.article.Subtext}}</small></h2>
  <label>作者：{{.product.Principal}}</label>
  <small id="publish">发表于：{{dateformat .product.Created "2006-01-02"}}</small>
  <!-- <a href='/project/search/'><span class="label label-info">标签</span></a>  -->
  <button type="button" class="btn-group btn-group-sm" id="updatearticle">编辑</button>
  <button type="button" class="btn-group btn-group-sm" id="deletearticle">删除</button>
</div>

  <div class="content">
    {{str2html .article.Content}}
  </div>
</div>
<script type="text/javascript">
  // function label(){//http://caibaojian.com/frame-adjust-content-height.html
  //   value={{.product.Label}};
  //   array=value.split(",");
  //   var labelarray = new Array(); 
  //   for (i=0;i<array.length;i++){
  //     // labelarray[i]="<a href='/project/search/"+array[i]+"'>" + array[i] + "</a>";
  //     var th1="&nbsp;<a href='/project/product/article/keysearch?keyword="+ array[i]+"'><span class='label label-info'>" +array[i] + "</span></a>";
  //     $("small#publish").after(th1);
  //   }
  // }

  $(function(){//这个和$(document).ready(function()等价
    value={{.product.Label}};
    array=value.split(",");
    var labelarray = new Array(); 
    for (i=0;i<array.length;i++){
      // labelarray[i]="<a href='/project/search/"+array[i]+"'>" + array[i] + "</a>";
      var th1="&nbsp;<a href='/project/product/article/keysearch?keyword="+ array[i]+"'><span class='label label-info'>" +array[i] + "</span></a>";
      $("small#publish").after(th1);
    }
    // return labelarray.join(",");
    // $("input#cid").remove();
  })
    //实例化编辑器
    //议使用工厂方法getEditor创建和引用编辑器实例，如果在某个闭包下引用该编辑器，直接调用UE.getEditor('editor')就能拿到相关的实例
    // var ue = UE.getEditor('container');
    var ue = UE.getEditor('container', {
      autoHeightEnabled: true,
      autoFloatEnabled: true
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

    //****编辑文章
    $("#updatearticle").click(function() {
      if ({{.role}}!=1){
        alert("权限不够！");
        return;
      }
      $("input#cid").remove();
      var th1="<input id='cid' type='hidden' name='cid' value='" +{{.article.Id}}+"'/>"
      $(".modal-body").append(th1);//这里是否要换名字$("p").remove();
      $("#prodcode1").val({{.product.Code}});
      $("#prodname1").val({{.product.Title}});
      $('#subtext1').val({{.article.Subtext}})
      $("#prodlabel1").val({{.product.Label}});
      $("#prodprincipal1").val({{.product.Principal}});

      $('#modalArticle').modal({
        show:true,
        backdrop:'static'
        });  
    })

    $(function(){
        var content =$('#content').val();
        //判断ueditor 编辑器是否创建成功
        ue.addListener("ready", function () {
        // editor准备好之后才可以使用
        ue.setContent({{str2html .article.Content}});
        });
    });

  //添加文章
  function updatearticle(){
    // var radio =$("input[type='radio']:checked").val();
    var articleid = $('#cid').val();
    var prodcode = $('#prodcode1').val();
    var prodname = $('#prodname1').val();
    var subtext = $('#subtext1').val();
    var prodprincipal = $('#prodprincipal1').val();
    var prodlabel = $('#prodlabel1').val();
    var html = ue.getContent();
    // $('#myModal').on('hide.bs.modal', function () {  
    if (prodname&&prodcode){  
      $.ajax({
        type:"post",
        url:"/project/product/updatearticle",
        data: {pid:articleid,code:prodcode,title:prodname,subtext:subtext,label:prodlabel,content:html,principal:prodprincipal},//父级id
        success:function(data,status){
          alert("修改“"+data+"”成功！(status:"+status+".)");
          $('#modalArticle').modal('hide');
          window.location.reload();//刷新页面
        },
      });
    }else{
      alert("请填写编号和名称！");
      return;
    }
  }

    // 删除项目
    $("#deletearticle").click(function() {
      if ({{.role}}!=1){
        alert("权限不够！");
        return;
      }
      // if(confirm("确定删除吗？第一次提示！")){
      // }else{
      //   return false;
      // }
      // if(confirm("确定删除项目吗？第二次提示！")){
      // }else{
      //   return false;
      // }
      if(confirm("确定删除项目吗？一旦删除将无法恢复！")){
        $.ajax({
          type:"post",
          url:"/project/product/deletearticle",
          data: {pid:{{.article.Id}}},
          success:function(data,status){
            alert("删除“"+data+"”成功！(status:"+status+".)");
            //关闭标签
            window.close();
          }
        });
      }  
    })
</script>
  <!-- 编辑文章 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalArticle">
      <div class="modal-dialog" style="width: 80%">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">添加文章</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">编号</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="prodcode1" name="prodcode1"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">标题</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodname1" name="prodname1"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">副标题</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="subtext1" name="subtext1"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">关键字</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodlabel1" name="prodlabel2"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">设计</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodprincipal1" name="prodprincipal2"></div>
              </div>
            </div>

            <!-- <label>成果内容:</label>
            <div id="content">
                <script id="editor" type="text/plain" style="height:500px;"></script>
            </div> -->

            <label>文章正文:</label>
              <div id="content">
                <script id="container" type="text/plain" style="height:200px;"></script><!-- width:1024px; -->
              </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
            <button type="button" class="btn btn-primary" onclick="updatearticle()">修改</button>
          </div>
        </div>
      </div>
    </div>
  </div>

</body>
</html>
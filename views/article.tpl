<!DOCTYPE html>
<title>CMS系统</title>
<meta name="viewport" content="width=device-width, initial-scale=1">
    <meta http-equiv="Content-Type" content="text/html;charset=utf-8"/>
    <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
    <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
    <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.js"> </script>

    <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>
    <script src="/static/ueditor/ueditor.parse.js"></script>
<style type="text/css">
  img{max-width:100%}


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
    ue.ready(function () {
        ue.addListener('focus', function () {//startUpload start-upload startUpload beforeExecCommand是在插入图片之前触发
            var pid = $('#pid').val();
            // var html = ue.getContent();
            ue.execCommand('serverparam', {
              "pid":pid 
            });
        });
    });
    uParse('.content',{
      rootPath : '/static/ueditor/'
    })

    //****编辑文章
    $("#updatearticle").click(function() {
      if ({{.role}}!=1){
        alert("权限不够！");
        return;
      }
      $("input#cid").remove();//文章id
      $("input#pid").remove();//项目目录id
      var th1="<input id='cid' type='hidden' name='cid' value='" +{{.article.Id}}+"'/>"
      var th2="<input id='pid' type='hidden' name='pid' value='" +{{.product.ProjectId}}+"'/>"
      $(".modal-body").append(th1);
      $(".modal-body").append(th2);
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
    var articleid = $('#cid').val();//文章id
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

    // 删除文章
    $("#deletearticle").click(function() {
      if ({{.role}}!=1){
        alert("权限不够！");
        return;
      }
      if(confirm("确定删除吗？一旦删除将无法恢复！")){
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
            <h3 class="modal-title">编辑文章</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">编号</label>
                <div class="col-sm-7">
                  <input type="text" readonly="readonly" class="form-control" id="prodcode1" name="prodcode1"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">标题</label>
                <div class="col-sm-7">
                  <input type="text" readonly="readonly" class="form-control" id="prodname1" name="prodname1"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">副标题</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="subtext1" name="subtext1"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">关键字</label>
                <div class="col-sm-7">
                  <input type="text" readonly="readonly" class="form-control" id="prodlabel1" name="prodlabel2"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">设计</label>
                <div class="col-sm-7">
                  <input type="text" readonly="readonly" class="form-control" id="prodprincipal1" name="prodprincipal2"></div>
              </div>
            </div>
            <label>文章正文:</label>
              <div id="content">
                <script id="container" type="text/plain" style="height:200px;width: 100%"></script><!-- width:1024px; -->
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
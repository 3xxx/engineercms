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

    <!-- <link href="/static/css/csdnarticlestyle.css" rel="stylesheet">
    <script src="/static/js/html5shiv.min.js"></script>
    <link rel="Stylesheet" type="text/css" href="/static/css/csdn_style.css">
    <link rel="stylesheet" href="/static/css/csdn_blog_detail.min.css"> -->

    <!-- <link href="/static/css/jquery.tocify.css" rel="stylesheet"> -->
    <!-- <link href="/static/css/prettify.css" type="text/css" rel="stylesheet"/> -->
    <!-- <script src="/static/js/skin2017.js"></script> -->
    <!-- <link rel="stylesheet" href="/static/css/common.css"> -->
    <!-- <link rel="stylesheet" href="/static/css/main.css"> -->
    <!-- 显示表格线 -->
    <!-- [if lt IE 9]-->
    
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <meta http-equiv="Cache-Control" content="no-siteapp">
    <title>{{.product.Title}} - 珠三角设代</title>
    <link rel="Stylesheet" type="text/css" href="/static/css/csdn_style.css">
    <link rel="stylesheet" href="/static/css/csdn_blog_detail.min.css">

<style type="text/css">
  img{max-width:100%}
  video{width:100%; height:100%; object-fit: fill}
  #main .article_next_prev span {
        width: 80px;
        margin: 0px;
        padding: 0px;
        padding-left: 30px;
        /*border: solid blue 10px;*/
        /*-webkit-box-sizing: border-box;*/
        /*-moz-box-sizing: border-box;*/
        /*box-sizing: border-box;*/
}
/*#main .article_next_prev a {


}
#main .article_next_prev .prev_article {

}
#main .article_next_prev .prev_article span {

}
#main .article_next_prev .next_article span {

}
#main .article_next_prev li:hover span {

}
#main .article_next_prev li:hover a {

}*/
        
  /*box-sizing: border-box;*/
  /*webkit-box-sizing: border-box;*/
</style>
</head>
<body>
  <div id="container">
    <div id="body">
      <div id="main">
        <div class="main">
          <div id="article_details" class="details">
            <div class="article_title">
              <span class="ico ico_type_Original"></span>
              <h1>
                <span class="link_title">
                  <!-- <a href="http://blog.csdn.net/hotqin888/article/details/78822774"> -->
                    {{.product.Code}}-{{.product.Title}}
                    <small>{{.article.Subtext}}</small>
                  <!-- </a> -->
                </span>
              </h1>
            </div>
            <div class="article_bar clearfix">
              <label>&nbsp;作者：{{.product.Principal}}</label>
              <small>发表于：{{dateformat .product.Created "2006-01-02 15:04:05"}}</small>
              <small id="publish">阅读{{.article.Views}}次</small>
              <!-- <button type="button" class="btn btn-warning btn-xs"  id="updatearticle">编辑</button> -->
              <a href="/project/product/modifyarticle/{{.article.Id}}" class="btn btn-warning btn-xs">修改</a>
              <button type="button" class="btn btn-danger btn-xs" id="deletearticle">删除</button>
            </div>  
            <div class="article_manage clearfix">
              <!-- <div class="article_l">
                <span class="link_categories">
                  标签：
                  <a href="http://www.csdn.net/tag/meritms" target="_blank" onclick="">meritms</a>
                  <a href="http://www.csdn.net/tag/%e5%b7%a5%e4%bd%9c%e6%b5%81%e5%bc%95%e6%93%8e" target="_blank" onclick="">工作流引擎</a>
                </span>
              </div> -->
              <!-- <div class="article_r">
                <span class="link_postdate">2017-12-16 22:52</span>
                <span class="link_view" title="阅读次数">71人阅读</span>
                <span class="link_comments" title="评论次数">
                  <a href="http://blog.csdn.net/hotqin888/article/details/78822774#comments" onclick="">评论</a>
                  (0)
                </span>
                <span class="link_collect tracking-ad" data-mod="popu_171">
                  <a href="javascript:void(0);" onclick="" title="收藏" target="_blank">收藏</a>
                </span>
                <span class="link_report">
                  <a href="http://blog.csdn.net/hotqin888/article/details/78822774#report" onclick="" title="举报">举报</a>
                </span>
              </div> -->
            </div> 

            <div class="bog_copyright">
              <p class="copyright_p">版权声明：本文为作者原创文章。</p>
            </div>
            <!-- <div id="article_content" class="article_content tracking-ad" data-mod="popu_307" data-dsm="post">
              <p>
                <a target="_blank" href="https://github.com/js-ojus/flow">
                  <span style="font-size:18px">https://github.com/js-ojus/flow</span>
                </a>
              </p>
            </div> -->
            <div id="article_content" class="article_content csdn-tracking-statistics" data-mod="popu_307" data-dsm="post">{{str2html .article.Content}}</div>

            <ul class="article_next_prev">
              {{if .Pre}}
              <li class="prev_article">
                <span>上一篇</span>
                <a href="/project/product/article/{{.PreArticleId}}" onclick="">{{.PreArticleTitle}}</a>
              </li>
              {{end}}
              {{if .Next}}
              <li class="next_article">
                <span>下一篇</span>
                <a href="/project/product/article/{{.NextArticleId}}" onclick="">{{.NextArticleTitle}}</a>
              </li>
              {{end}}
            </ul>
            <div style="clear:both; height:10px;"></div>
          </div>
          <div class="comment_class"></div>
        </div>
      </div>
    </div>
  </div>

<script type="text/javascript">
  $(function(){//这个和$(document).ready(function()等价
    value={{.product.Label}};
    array=value.split(",");
    var labelarray = new Array(); 
    for (i=0;i<array.length;i++){
      // labelarray[i]="<a href='/project/search/"+array[i]+"'>" + array[i] + "</a>";
      var th1="&nbsp;<a href='/project/product/article/keysearch?keyword="+ array[i]+"'><span class='btn btn-info btn-xs'>" +array[i] + "</span></a>";//label label-info
      $("small#publish").after(th1);
    }
    // return labelarray.join(",");
    // $("input#cid").remove();
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
      // if ({{.role}}!=1){
      //   alert("权限不够！");
      //   return;
      // }
      if ({{.product.Uid}}==={{.Uid}}||{{.RoleUpdate}}=="true"){
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
            // var widHeight = $(window).height();
            // var artHeight = $('.article_content').height();
            //     $('.article_content').height(widHeight*2-285).css({'overflow':'hidden'});
            //     var article_show = true;
            //         if(!article_show){
            //             $('.article_content').height(widHeight*2-285).css({'overflow':'hidden'});
            //             $('.readall_box').show().removeClass('readall_box_nobg');
            //             article_show = true;
            //         }else{
            //             $('.article_content').height("").css({'overflow':'hidden'});
            //             $('.readall_box').show().addClass('readall_box_nobg');
            //             $('.readall_box').hide().addClass('readall_box_nobg');
            //             article_show = false;
            //         };
      $('#modalArticle').modal({
        show:true,
        backdrop:'static'
      });
    }else{
        alert("权限不够！"+{{.Uid}});
        return;
      }
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
      if ({{.product.Uid}}==={{.Uid}}||{{.RoleUpdate}}=="true"){

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
      }else{
        alert("权限不够！"+{{.Uid}});
        return;
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
                <script id="article_container" type="text/plain" style="height:200px;width: 100%"></script><!-- width:1024px; -->
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

    <script src="/static/js/jquery-ui-1.9.1.custom.min.js"></script>
    <script src="/static/js/jquery.tocify.min.js"></script>
    <script src="/static/js/prettify.js"></script>
    <script>
        $(function() {

            var toc = $("#toc").tocify({
              selectors: "h2,h3,h4,h5"
            }).data("toc-tocify");

            prettyPrint();
            $(".optionName").popover({ trigger: "hover" });
        });
    </script>

</body>
</html>

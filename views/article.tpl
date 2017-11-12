<!DOCTYPE html>
<title>CMS系统</title>
<meta name="viewport" content="width=device-width, initial-scale=1">
    <meta http-equiv="Content-Type" content="text/html;charset=utf-8"/>
    <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
    <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
    <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
    <link rel="stylesheet" type="text/css" href="/static/css/csdn-style.css"/>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.js"> </script>

    <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>
    <script src="/static/ueditor/ueditor.parse.js"></script>
    <link href="/static/css/jquery.tocify.css" rel="stylesheet">
    <link href="/static/css/prettify.css" type="text/css" rel="stylesheet"/>
    <!-- <script src="/static/js/skin2017.js"></script> -->
<style type="text/css">
  img{max-width:100%}
</style>

</head>

<body>
  <!-- <div class="col-sm-2 hidden-xs"></div> -->
  <!-- <div class="col-sm-6 col-xs-12"> -->
    <!-- <div class="row" style="margin:0;"> -->
      <!-- <div class="col-sm-12 col-xs-12"> -->
  <div class="page-header">    
  <div class="col-lg-2 col-sm-2 col-xs-2">
    <div id="toc">
      </div><!--/.well -->
  </div>
  <div class="col-lg-8 col-sm-8 col-xs-12 col-md-offset-1"><!--   -->
  <!-- <div class="page-header"> -->

   <div class="container clearfix">
    <main>
      <article>
        <h1 class="csdn_top">{{.product.Code}}-{{.product.Title}} <small>{{.article.Subtext}}</small></h1>
        <div class="article_bar clearfix">
          <label>作者：{{.product.Principal}}</label>
          <small>发表于：{{dateformat .product.Created "2006-01-02 15:04:05"}}</small>
          <small id="publish">阅读4241次</small>
          <!-- <a href='/project/search/'><span class="label label-info">标签</span></a>  -->
          <button type="button" class="btn btn-warning btn-xs"  id="updatearticle">编辑</button>
          <button type="button" class="btn btn-danger btn-xs" id="deletearticle">删除</button>
        </div>


        <div id="article_content" class="article_content csdn-tracking-statistics" data-mod="popu_307" data-dsm="post">
        <!-- <div class="content"> style="height: 1425px; overflow: hidden;" -->
          {{str2html .article.Content}}
        </div>
        <!-- </div> -->
        <!-- </div> -->
      </article>

      <div class="readall_box csdn-tracking-statistics" data-mod="popu_376" >
            <a class="btn btn-large btn-gray-fred read_more_btn" target="_self">阅读全文</a>
      </div>
      <div class="article_copyright">
            版权声明：本文为博主原创文章，未经博主允许不得转载。
            <span class="r_ico"><i class="icon iconfont icon-jubao"></i><span class="txt" id="reportBtn">举报</span></span>
      </div>
      <ul class="article_tags clearfix csdn-tracking-statistics" data-mod="popu_377">
            <li class="tit">标签：</li>
            
            <li><a href="http://so.csdn.net/so/search/s.do?q=MeritMS&amp;t=blog" target="_blank">MeritMS</a> <span>/</span></li>
            
            <li><a href="http://so.csdn.net/so/search/s.do?q=%E8%BF%9B%E5%BA%A6&amp;t=blog" target="_blank">进度</a> <span>/</span></li>
            
            <li><a href="http://so.csdn.net/so/search/s.do?q=%E5%9B%BE%E6%A0%87select&amp;t=blog" target="_blank">图标select</a> <span style="display: none;">/</span></li>
            
      </ul>
    </main>
      
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



<script type="text/javascript">
    if($(".article_collect li").length==1){$(".article_collect").hide();}
    if($(".article_tags li").length==1){$(".article_tags").hide();}
    // $(".edit a").attr("href","http://write.blog.csdn.net/postedit/"+fileName);
    $.each($(".edu_li a"),function(){$(this).attr("href",$(this).attr("href").replace("blog7","blog9"))});
    // new CNick('#uid').showNickname();
if($("#fan").html()=="")
{
$("#fan").html(0);  
}
</script>
<script type="text/javascript">
   
var tatolList = []; //评论列表总数
$(function(){

    var oBlog = {
        // left menu
        leftMenu:function(){
            var commentList = true;
            var comment_child_list_show = true;
            $('#comment_list').on('click','.rec_number',function(){
                var com_child_listshow = $(this).parents('.comment_li_box').find('.child_comment').attr('data-listshow');
                var arrow = $(this).find(".iconfont");
                if(com_child_listshow==='false'){
                    arrow.addClass("icon-shangjiantou").removeClass("iconfont-xiajiantou");
                    $(this).parents('.comment_li_box').find('.child_comment').css({'height':'auto'});
                    $(this).parents('.comment_li_box').find('.child_comment').attr({'data-listshow':'true'});
                }else{
                    arrow.addClass("iconfont-xiajiantou").removeClass("icon-shangjiantou");
                    $(this).parents('.comment_li_box').find('.child_comment').css({'height':'0'});
                    $(this).parents('.comment_li_box').find('.child_comment').attr({'data-listshow':'false'});
                }
            });
            $('.more_comment_btn').on('click',showAllCom);
            function showAllCom(){
                var allComHeight = $('.comment_li_insaidbox').height();
                if(!commentList){
                    $('.comment_li_outbox').css({'height':'auto'});
                    commentList = true;
                }else{
                    $('.comment_li_outbox').css({'height':'auto'});
                    commentList = false;
                }
            }
        },
        //code popup
        codePop:function(){
            // code popup
            
            var lang_listCode = '<div id="lang_list" code="code">\
                                    <a href="#html" style="width:95px;" class="long_name" target="_self">HTML/XML</a>\
                                    <a href="#objc" style="width:95px;" class="long_name" target="_self">objective-c</a>\
                                    <a href="#delphi" style="width:58px;" class="zhong_name" target="_self">Delphi</a>\
                                    <a href="#ruby" class="zhong_name" target="_self">Ruby</a>\
                                    <a href="#php" target="_self">PHP</a>\
                                    <a href="#csharp" class="duan_name" target="_self">C#</a>\
                                    <a style=" border-right: none;" href="#cpp" class="duan_name" target="_self">C++</a>\
                                    <a style=" border-bottom:none;width:95px;" href="#javascript" class="long_name" target="_self">JavaScript</a>\
                                    <a style=" border-bottom:none;width:95px;" href="#vb" class="long_name" target="_self">Visual Basic</a>\
                                    <a style=" border-bottom:none;" href="#python" class="zhong_name" target="_self">Python</a>\
                                    <a style=" border-bottom:none;" href="#java" class="zhong_name" target="_self">Java</a>\
                                    <a style="border-bottom:none;" href="#css" class="duan_name" target="_self">CSS</a>\
                                    <a style="border-bottom:none;" href="#sql" class="duan_name" target="_self">SQL</a>\
                                    <a style="border:none; " href="#plain" class="duan_name" target="_self">其它</a>\
                                    <span class="arrb"></span>\
                                </div>';
          
            setTimeout(function(){
              if(!($('#lang_list')[0])){
                $('.bot_bar').append(lang_listCode)
              }
            },1000);
            var comment_language = true;
            $('#ubbtools').on('click',function(){
              if(!($('#lang_list')[0])){
                $('.bot_bar').append(lang_listCode)
              }
                if(!comment_language){
                    $('#lang_list').hide();
                    comment_language = true;
                }else{
                    $('#lang_list').show();
                    comment_language = false;
                }
            });
        },
        //readMore
        readMore:function(){
            var widHeight = $(window).height();
            var artHeight = $('.article_content').height();
            // alert(widHeight);
            // alert(artHeight);
            // if(username === currentUserName ){
            //   $('.readall_box').hide().addClass('readall_box_nobg');
            //    return false;
            // }
            if(artHeight>widHeight){
                $('.article_content').height(widHeight*2-285).css({'overflow':'hidden'});
                var article_show = true;
                $('.read_more_btn').on('click',function(){
                    if(!article_show){
                        $('.article_content').height(widHeight*2-285).css({'overflow':'hidden'});
                        $('.readall_box').show().removeClass('readall_box_nobg');
                        article_show = true;
                    }else{
                        $('.article_content').height("").css({'overflow':'hidden'});
                        $('.readall_box').show().addClass('readall_box_nobg');
                        $('.readall_box').hide().addClass('readall_box_nobg');
                        article_show = false;
                    }
                });
            }else{
              // alert("readall");
                article_show = true;
                $('.article_content').removeClass('article_Hide');
                $('.readall_box').hide().addClass('readall_box_nobg');

            }
        },
        //left menu show and hide
        // leftMenuShow:function(){
        //     // left menu
        //     var leftmenu_show = true;
        //     var leftMenuShow = function(){
        //         var winWdith = $(window).width();
        //         if(!leftmenu_show){
        //             $('.menu_con').css({'overflow':'visible'}).animate({'height':0});
        //             $('.list_father').parents('.menu_con').removeClass('smallShow');
        //             leftmenu_show = true;
        //         }else{
        //             /*if(winWdith<1530){
        //              }else{
        //              $('.menu_con').css({'overflow':'visible'}).animate({'height':'220px'});
        //              }*/
        //             $('.menu_con').css({'overflow':'visible'}).animate({'height':0});
        //             $('.list_father').parents('.menu_con').addClass('smallShow');
        //             leftmenu_show = false;
        //         }
        //     }
        //     $('.left_menu_btn').on('click',function(event){
        //         leftMenuShow();
        //         event.stopPropagation();
        //     });
        //     $('.menu_con').on('click',function(event){
        //         event.stopPropagation();
        //     });
        //     var FixedLeftMenu = $('main').offset().left-80;
        //     $('.left_fixed').css({'left':FixedLeftMenu});
        //     $(document).on('click',function(){
        //         $('.menu_con').css({'overflow':'hidden'}).animate({'height':0});
        //         $('.list_father').parents('.menu_con').removeClass('smallShow');
        //         leftmenu_show = true;
        //     });
        // },
        // left_menu scroll top and down
        topDown:function(){
            var blogDir = $("#csdnBlogDir");
            var h = parseInt(blogDir.height()),
                scrolls = 1;
            $(".left_scroll_down").click(function(event){
                var menuH = parseInt($(".first_li").height());
                if(scrolls < (menuH-h)){
                    scrolls += 24;
                    blogDir.scrollTop(scrolls);
                }else{
                    return;
                }
                event.stopPropagation();
            });
            $(".left_scroll_top").click(function(event){
                if(scrolls >= 2){
                  scrolls -= 24;
                  blogDir.scrollTop(scrolls);
                }
                event.stopPropagation();
            });
        },
        //like
        // likeFn:function(){
        //     var likeBtn = $(".btn-like"),
        //         $likeNum = $('.inf_number_box dl:eq(2) dd');
        //     if (currentUserName) {
        //         $.get(blog_address + "/article/IsDigg?ArticleId=" + fileName, function (data) {
        //             if(data.status==0){
        //                 likeBtn.removeClass('liked');
        //             }else{
        //                 likeBtn.addClass('liked');
        //             }
        //             likeBtn.parents(".right_bar").find(".btn-like .txt").text(data.digg);
        //         });
        //     }else{
        //         likeBtn.removeClass('liked');
        //     }
        //     $(".btn-like").on('click',function(){
        //         var curl = encodeURIComponent(location.href);
        //         if (currentUserName) {
        //             $.get(blog_address + "/article/digg?ArticleId=" + fileName, function (data) {
        //                 if(data.status==0){
        //                     $(".btn-like").removeClass('liked');
        //                       $likeNum.text(parseInt($likeNum.text())-1);
        //                 }else{
        //                     $(".btn-like").addClass('liked');
        //                     $likeNum.text(parseInt($likeNum.text())+1)
        //                 }
        //                 $(".btn-like").parents(".right_bar").find('.btn-like .txt').text(data.digg);
        //             });
        //         }else{
        //             var logintip =  '<div class="login_tip_bg"></div>'+
        //                     '<div class="login_tip_box clearfix">'+
        //                     '<span class="close">关闭</span>'+
        //                     '请先<a href="https://passport.csdn.net/account/login?from=' + curl + '">登录</a> 或 ' +
        //                     '<a href="http://passport.csdn.net/account/register?from=' + curl + '">注册</a>;'+
        //                     '</div>';
        //             $('body').append(logintip);
        //             $(".btn-like").removeClass('liked');
        //             $(document).on('click','.login_tip_box .close',function(){
        //                 $('.login_tip_bg').remove();
        //                 $('.login_tip_box').remove();
        //             })
        //         }
        //     });
        // },
        // right fixed baidu_box
        // fixedWrap:function(){
        //   var fixWrap = $(".fixRight"),
        //                  fixTop = fixWrap.offset().top + 276, // 距顶部的距离加上后出现的广告的高度
        //                  subFirst = fixWrap.find(".edu_li dd:first"),
        //                  subSec = fixWrap.find(".edu_li dt:first"),
        //                  winH = $(window).height(),
        //                  rightH = $(".fixRight").height() + 276,
        //                  leftPos = $('aside').offset().left,
        //                  wrapH = fixWrap.height();
                         
        //                 //  $('.bottomRcom').css({'margin-top':'150px'});
        //              $(window).scroll(function(){
        //                  var scrollTop = $(document).scrollTop(), // 谷歌识别 body  火狐只识别html document 都识别
        //                      baiduBoxLeft = fixWrap.offset().left;
        //                  if(scrollTop-1300 >=0){
        //                    if($('.bottomRcom').css('opacity') == '0' && !$('.bottomRcom').is(":animated")){
        //                       $('.bottomRcom').animate({opacity:1},500);
        //                       $('.bottomRcom').css({'position':'fixed','top':'20px','z-index':'1','left':leftPos,'pointer-events':'all'});
        //                    }
        //                  }else{
        //                    if($('.bottomRcom').css('opacity') == '1' && !$('.bottomRcom').is(":animated")){
        //                       $('.bottomRcom').animate({opacity:0},300);
        //                       $('.bottomRcom').css({'pointer-events':'all','z-index':'-5',});
        //                    }
        //                     //  $('.bottomRcom').animate({opacity:0},1000); 
        //                     //  $('.bottomRcom').css({'width':'300px','opacity':'1','height':'256px','position':'relative','z-index':'999','top':'0','left':'auto','pointer-events':'none'});  
        //                  }
        //              });
        // },
        // Upper right corner show NickName
        // showNickName:function(){
        //   new CNick('.alrLogin dd a:first').showNickname();
        // },
        //去掉tag 斜线
        deleteTagline:function(){
          $('.article_tags li:last-child > span').css({'display':'none'})
        },
        hotcolumnlist:function(){
          var $columnListDOM = $('.column-list'),
              columnListDOMOldHeight = $columnListDOM.height() ,
              $unfoldBtnDOM = $('.unfold-btn');
          if(columnListDOMOldHeight > 276) $columnListDOM.height(276).css({'overflow-y':'hidden'})
          else $unfoldBtnDOM.fadeOut();
          $(document).on('click','.unfold-btn' ,function(){
            $columnListDOM.animate({height:columnListDOMOldHeight},500,function(){
              $unfoldBtnDOM.animate({height:0},500,function(){$unfoldBtnDOM.css({'display':'none'})});
            })
          })
        }
    };
    oBlog.leftMenu();
    oBlog.codePop();
    oBlog.readMore();
    // oBlog.leftMenuShow();
    oBlog.topDown();
    // oBlog.likeFn();
    // oBlog.fixedWrap();
    // oBlog.showNickName();
    oBlog.deleteTagline();
    oBlog.hotcolumnlist();
    //$(window).resize(oBlog.fixedWrap);

    // top nav more popup
    $('.dropdown-toggle').dropdown();
});
        //  超出显示省略号
function overflow_hide(dom,height,className){
  dom.map(function(){
    return $(this).height()>height ?this :null;
  }).addClass('overflow-hide '+ className);
}
</script>
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
/* commend show*/
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
            if(username === currentUserName ){
              $('.readall_box').hide().addClass('readall_box_nobg');
               return false;
            }
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
                article_show = true;
                $('.article_content').removeClass('article_Hide');
                $('.readall_box').hide().addClass('readall_box_nobg');
            }
        },
        //left menu show and hide
        leftMenuShow:function(){
            // left menu
            var leftmenu_show = true;
            var leftMenuShow = function(){
                var winWdith = $(window).width();
                if(!leftmenu_show){
                    $('.menu_con').css({'overflow':'visible'}).animate({'height':0});
                    $('.list_father').parents('.menu_con').removeClass('smallShow');
                    leftmenu_show = true;
                }else{
                    /*if(winWdith<1530){
                     }else{
                     $('.menu_con').css({'overflow':'visible'}).animate({'height':'220px'});
                     }*/
                    $('.menu_con').css({'overflow':'visible'}).animate({'height':0});
                    $('.list_father').parents('.menu_con').addClass('smallShow');
                    leftmenu_show = false;
                }
            }
            $('.left_menu_btn').on('click',function(event){
                leftMenuShow();
                event.stopPropagation();
            });
            $('.menu_con').on('click',function(event){
                event.stopPropagation();
            });
            var FixedLeftMenu = $('main').offset().left-80;
            $('.left_fixed').css({'left':FixedLeftMenu});
            $(document).on('click',function(){
                $('.menu_con').css({'overflow':'hidden'}).animate({'height':0});
                $('.list_father').parents('.menu_con').removeClass('smallShow');
                leftmenu_show = true;
            });
        },
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
        likeFn:function(){
            var likeBtn = $(".btn-like"),
                $likeNum = $('.inf_number_box dl:eq(2) dd');
            if (currentUserName) {
                $.get(blog_address + "/article/IsDigg?ArticleId=" + fileName, function (data) {
                    if(data.status==0){
                        likeBtn.removeClass('liked');
                    }else{
                        likeBtn.addClass('liked');
                    }
                    likeBtn.parents(".right_bar").find(".btn-like .txt").text(data.digg);
                });
            }else{
                likeBtn.removeClass('liked');
            }
            $(".btn-like").on('click',function(){
                var curl = encodeURIComponent(location.href);
                if (currentUserName) {
                    $.get(blog_address + "/article/digg?ArticleId=" + fileName, function (data) {
                        if(data.status==0){
                            $(".btn-like").removeClass('liked');
                              $likeNum.text(parseInt($likeNum.text())-1);
                        }else{
                            $(".btn-like").addClass('liked');
                            $likeNum.text(parseInt($likeNum.text())+1)
                        }
                        $(".btn-like").parents(".right_bar").find('.btn-like .txt').text(data.digg);
                    });
                }else{
                    var logintip =  '<div class="login_tip_bg"></div>'+
                            '<div class="login_tip_box clearfix">'+
                            '<span class="close">关闭</span>'+
                            '请先<a href="https://passport.csdn.net/account/login?from=' + curl + '">登录</a> 或 ' +
                            '<a href="http://passport.csdn.net/account/register?from=' + curl + '">注册</a>;'+
                            '</div>';
                    $('body').append(logintip);
                    $(".btn-like").removeClass('liked');
                    $(document).on('click','.login_tip_box .close',function(){
                        $('.login_tip_bg').remove();
                        $('.login_tip_box').remove();
                    })
                }
            });
        },
        // right fixed baidu_box
        fixedWrap:function(){
          var fixWrap = $(".fixRight"),
                         fixTop = fixWrap.offset().top + 276, // 距顶部的距离加上后出现的广告的高度
                         subFirst = fixWrap.find(".edu_li dd:first"),
                         subSec = fixWrap.find(".edu_li dt:first"),
                         winH = $(window).height(),
                         rightH = $(".fixRight").height() + 276,
                         leftPos = $('aside').offset().left,
                         wrapH = fixWrap.height();
                         
                        //  $('.bottomRcom').css({'margin-top':'150px'});
                     $(window).scroll(function(){
                         var scrollTop = $(document).scrollTop(), // 谷歌识别 body  火狐只识别html document 都识别
                             baiduBoxLeft = fixWrap.offset().left;
                         if(scrollTop-1300 >=0){
                           if($('.bottomRcom').css('opacity') == '0' && !$('.bottomRcom').is(":animated")){
                              $('.bottomRcom').animate({opacity:1},500);
                              $('.bottomRcom').css({'position':'fixed','top':'20px','z-index':'1','left':leftPos,'pointer-events':'all'});
                           }
                         }else{
                           if($('.bottomRcom').css('opacity') == '1' && !$('.bottomRcom').is(":animated")){
                              $('.bottomRcom').animate({opacity:0},300);
                              $('.bottomRcom').css({'pointer-events':'all','z-index':'-5',});
                           }
                            //  $('.bottomRcom').animate({opacity:0},1000); 
                            //  $('.bottomRcom').css({'width':'300px','opacity':'1','height':'256px','position':'relative','z-index':'999','top':'0','left':'auto','pointer-events':'none'});  
                         }
                     });
        },
        // Upper right corner show NickName
        showNickName:function(){
          new CNick('.alrLogin dd a:first').showNickname();
        },
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
    oBlog.leftMenuShow();
    oBlog.topDown();
    oBlog.likeFn();
    oBlog.fixedWrap();
    oBlog.showNickName();
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
// 外部blog方法，多是监听事件
var outoBlog = {
  leftMenuAutoMove:function(){  // 左侧菜单位置自动调整
    var WindowWidth = $(window).width(),
        removing = 80;
    if(WindowWidth < 1366){
      removing = 50;
    }else {
      removing = 80;
    }
    if(WindowWidth < 1280){
      $('.left_fixed').css({'display':'none'})
    }else{
      $('.left_fixed').css({'display':'block'})
    }
    var FixedLeftMenu = $('main').offset().left-removing;
    $('.left_fixed').css({'left':FixedLeftMenu});
  },
  rightAutoMove:function(){ // 右侧广告位置自动调整
    var fixWrap = $(".fixRight"),
        fixleft = $('aside').offset().left;
    if(fixWrap.css('position') === 'fixed'){
        fixWrap.css('left',fixleft)
    }
  },
  MoreComment:function(){
    var $moreComment = $('.more_comment');
    var $comment_bar = $('#comment_bar');
    if($comment_bar.html() === ''){
      $moreComment.css('display','none');
    }else{
      $moreComment.css('display','block');
    }
  },
  returnTop:function(){
    var windowH = $(window).height(),
        scrollTop = $(document).scrollTop(),
        $returnTop = $('.returnTop');
    if(scrollTop>windowH){
      $returnTop.fadeIn(1000);
    }else{
      $returnTop.fadeOut(1000);
    }
  }
}
// 监听事件整合
var oBlogResize =function(){
    outoBlog.leftMenuAutoMove();
    outoBlog.rightAutoMove();
}
var oBlogScroll = function(){
  outoBlog.returnTop()
}
// window resize start
$(window).resize(oBlogResize);
// window scroll start
$(window).scroll(oBlogScroll);
// popu comment submit btn
$('#commentform input[type="submit"]').on('click',function(){
     $(".comment_area_btn").click();
});
//relate artical list
$(function(){
    var loadanAnimation = "<div class='Recommend_loadanAnimation'><img src='http://static.blog.csdn.net/Skin/skin3-template/images/feedLoading.gif'></div>" 
    //根据页数读取数据
    var allpage = 1;
    var noMoreData = false;
    function getRelateListData() {
      var noMoreRecommend = '<div class="clearfix nomore_box" style="text-align:center">没有更多内容了，<a href="http://blog.csdn.net/">返回首页</a></div>';
        allpage += 1;
        if(allpage>10){
          if(noMoreData)return false;
          noMoreData = true;
          $('.recommend_list').append(noMoreRecommend);
            return;
        }
        // $.get("http://blog.csdn.net/csdn/svc/GetRelatedArticles?pageindex="+allpage+"&articleId="+fileName, function (data) {
        // 
        // });
        
        if(!$('.Recommend_loadanAnimation')[0]){
          $('.recommend_list').append(loadanAnimation);
          overflow_hide($('.downloadElement .summary h2 a'),28 ,'overflow-hide-title-height')
        }
        $.ajax({
          url: "http://blog.csdn.net/csdn/svc/GetRelatedArticles?pageindex="+allpage+"&articleId="+fileName,
          beforeSend:function(){
      
          },
          success:function(data){
            $('.Recommend_loadanAnimation').remove();
            if (data.length > 0) {
                $('.recommend_list').append(data);
            }else{
                noMoreData = true;
                $('.recommend_list').append(noMoreRecommend);
                return
            }
          }
        })      
    }
    //初始化加载第一页数据
    getRelateListData();

    var winH = $(window).height(); //页面可视区域高度 

    var scrollHandler = function () {
        if(!noMoreData){
            var pageH = $(document.body).height();
            var scrollT = $(window).scrollTop(); //滚动条top
            var aa = (pageH - winH - scrollT) / winH;
            if (aa < 0.02) {
                getRelateListData();
            }
        }else{
            return
        }
    }
    //函数节流
    var throttleV2 = function(fn, delay, mustRunDelay){
        var timer = null;
        var t_start;
        return function(){
            var context = this, args = arguments, t_curr = +new Date();
            clearTimeout(timer);
            if(!t_start){
                t_start = t_curr;
            }
            if(t_curr - t_start >= mustRunDelay){
                fn.apply(context, args);
                t_start = t_curr;
            }
            else {
                timer = setTimeout(function(){
                    fn.apply(context, args);
                }, delay);
            }
        };
     };
    //定义鼠标滚动事件
    $(window).scroll(throttleV2(scrollHandler,200,400));
});

function loginbox(anchor) {
    document.domain = "csdn.net";
    var $logpop = $("#pop_win");
    var iframeUrlParm ="";
    if (anchor != null && anchor != undefined && anchor != "")
    {
        iframeUrlParm = "?anchor=" + anchor;
    }
    $logpop.html('<iframe src="https://passport.csdn.net/account/loginbox?service=http://static.blog.csdn.net/callback.htm' + iframeUrlParm + '" frameborder="0" height="600" width="400" scrolling="no"></iframe>');

    $('#popup_mask').css({
        opacity: 0.5,
        width: $(document).width() + 'px',
        height: $(document).height() + 'px'
    });
    $('#popup_mask').css("display", "block");

    $logpop.css({
        top: ($(window).height() - $logpop.height()) / 2 + $(window
   ).scrollTop() + 'px',
        left: ($(window).width() - $logpop.width()) / 2
    });

    setTimeout(function () {
        $logpop.show();
        $logpop.css({
            opacity: 1
        });
    }, 200);

    $('#popup_mask').unbind("click");
    $('#popup_mask').bind("click", function () {
        $('#popup_mask').hide();
        var $clopop = $("#pop_win");
        $("#common_ask_div_sc").css("display", "none");
        $clopop.css({
            opacity: 0
        });
        setTimeout(function () {
            $clopop.hide();
        }, 350);
        return false;
    });
}

//comment begin
var list = []; //评论列表

var editorId = "#comment_content";
var verifycodeId = "#img_verifycode";
var listId = "#comment_list";
var listOutBoxClass = ".comment_li_outbox";


$(document).ready(init_comment);

(function ($) {
    $.fn.extend({
        selection: function () {
            var selectedValue = '';
            var me = this[0];
            if (document.selection) {
                var range = document.selection.createRange();
                selectedValue = range.text;
            }
            else if (typeof (me.selectionStart) == 'number') {
                var start = me.selectionStart;
                var end = me.selectionEnd;
                if (start != end) {
                    selectedValue = me.value.substring(start, end);
                }
            }
            return $.trim(selectedValue);
        }, parseHtml: function (val) {
            var me = this[0];
            var value = $(this).val();
            if (document.selection) {
                var range = document.selection.createRange();
                if (range.text) {
                    range.text = val;
                } else {
                    $(this).val(value + val);
                }
            } else if (typeof (me.selectionStart) == 'number') {
                var start = me.selectionStart;
                var end = me.selectionEnd;
                var startVal = value.substring(0, start);
                var endVal = value.substring(end);
                $(this).val(startVal + val + endVal);
            }
            else
                $(this).val(value + val);
            me.selectionStart = me.selectionEnd = $(this).val().length;
            me.focus();
        }
    });
})(jQuery);

function init_comment() {
    load_comment_form();

    editor = $(editorId);

    var editor_inter = null;
    if (editor.length > 0) {    
        //$("#lang_list").append('<a class="long_name" href="#html">HTML/XML</a><a class="long_name" href="#objc">objective-c</a><a class="zhong_name" href="#delphi">Delphi</a><a  class="zhong_name" href="#ruby">Ruby</a><a href="#php">PHP</a><a class="duan_name" href="#csharp">C#</a><a style=" border-right: none;"  class="duan_name" href="#cpp">C++</a><a style=" border-bottom:none;"class="long_name" href="#javascript">JavaScript</a><a style=" border-bottom:none;" class="long_name" href="#vb">Visual Basic</a><a style=" border-bottom:none;" class="zhong_name" href="#python">Python</a><a style=" border-bottom:none;" class="zhong_name" href="#java">Java</a><a style="border-bottom:none;" class="duan_name" href="#css">CSS</a><a style="border-bottom:none;" class="duan_name" href="#sql">SQL</a><a style="border:none;"  class="duan_name" href="#plain">其它</a>');
        editor.focus(function () {
            editor_inter = setInterval(function () {
                commentTip("还能输入" + (1000 - editor.val().length) + "个字符");
            }, 200);
        }).blur(function () {
            if (editor_inter) clearInterval(editor_inter);
        });
    }

    //加载列表
    loadList(1,true);

}
function noComments() {
    $(listOutBoxClass).html('');
}
function loadList(pageIndex, isSub) {
    if (commentscount == 0) {
        outoBlog.MoreComment();
        noComments();
        return;
    }else{
      if(!($(listId)[0]))$(listOutBoxClass).html('<div id="comment_list"></div>');
        pageIndex = parseInt(pageIndex) || 1;
        var url = location.href.toString().split('/');
        var cmtUrl = "http://"+url[2]+"/"+url[3]+"/comment/list/" + fileName + "?page=" + pageIndex +"&size=3";
        if (isSub) cmtUrl += "&_" + Math.random();
        $.get(cmtUrl, function (json) {
            var data = (typeof json == 'object') ? json : eval("(" + json + ")");
            var RecordCount = data.page.RecordCount;
            if (isSub) list = data.list;
            else list = list.concat(data.list)
            if(RecordCount>0){
              if(RecordCount == list.length){
                $("#comment_bar").css("display","none");
              } else if(RecordCount < list.length){
                $("#comment_bar").css("display","none");
              } else{
                  $("#comment_bar").css("display","block");
              }
            }
            var listHtml = '';
            //listHtml+=' <h3 class="com_list_t">共有4条评论</h3>';
            var Url= "http://"+url[2]+"/"+url[3]+"/comment/list/" + fileName + "?page=" + pageIndex +"&size="+RecordCount+"";
            $.get(Url,function(json){
                var tatoldata = (typeof json == 'object') ? json : eval("(" + json + ")");
                if(!(tatoldata.list.length === 0)){
                  tatolList = tatoldata.list;
                }
            })
            //构造主题
            var topics = getTopics(list);
            if(topics.length < 3) {
              topics = loadComment(topics,2);
            }
            var total = data.total > 0 ? data.total : topics.length;
            //组装HTM
            for (var i = 0; i < total; i++) {
                var comment = topics[i];
                var layer = totalFloor - i;
                if(layer >0){
                  listHtml += getItemHtml(comment, layer);
                }

            }
            listHtml += '<div class="clear"></div>';
            //输出列表
            $(listId).html(listHtml);
            //高亮评论中的代码段
            dp.SyntaxHighlighter.HighlightAll('code2');
            //展示昵称
            new CNick(listId + ' a.username').showNickname();

            //分页处理
            /*if (data.page.PageIndex >= data.page.PageCount) {
             $("#comment_bar").hide();
             } else {
             $("#comment_bar").html('<a id="load_comments" class="btn btn-large btn-gray-fred more_comment_btn" data-mod="popu_385" page="' + (pageIndex + 1) + '">更多评论</a>');
             }*/
            $("#comment_bar").html('<a id="load_comments" class="btn btn-large more_comment_btn" page="' + (pageIndex + 1) + '">查看 <span>'+RecordCount +'</span> 条热评<i class="icon iconfont iconfont-xiajiantou"></i></a>');
            //添加按钮事件
            setBtnEvent();
            outoBlog.MoreComment();
            //scrollInit();
        });
    }
    
}
function loadComment(topics,page){
  var url = location.href.toString().split('/');
  var cmtUrl = "http://"+url[2]+"/"+url[3]+"/comment/list/" + fileName + "?page=" + page +"&size=3";
  var list=[]
  
  $.ajax({
    async: false,//设置为同步
    url:cmtUrl,
    success: function(json) {
      var data = (typeof json == 'object') ? json : eval("(" + json + ")");
          list = data.list;
      var topicsList = getTopics(list)
      for (var i = 0; i < topicsList.length; i++) {
        if(topics.length > 2) break;
        topics.push(topicsList[i]);
      }
    }
  })
  if (list.length > 0) {
    if(topics.length < 3){ 
      loadComment(topics,page+1);
    }
  }
  return topics;
}
//获取评论主题
function getTopics(list) {
    var topics = [];
    for (var i = 0; i < list.length; i++) {
        var t = list[i];
        if (t.ParentId == 0) {
            t.Replies = getReplies(t, list);
            topics.push(t);
        }
    }
    return topics;
}
//获取评论回复
function getReplies(item, list) {
    var replies = [];
    for (var i = 0; i < list.length; i++) {
        var r = list[i];
        if (r.ParentId == item.CommentId) {
            r.Replies = getReplies(r, list);
            replies.push(r);
        }
    }
    return replies;
}
//获取评论的HTML
function getItemHtml(comment, index, floor, deep) {    
    var html = ' <div class="comment_li_box clearfix">';
    html += '         <dl class="comment_list clearfix" id="comment_item_' + comment.CommentId + '">';
    html += '           <dt>';
    html += '               <a href="/' + comment.UserName + '"><img src="' + comment.Userface + '" alt="' + comment.UserName + '"></a></dt>';
    html += '           <dd>';
    html += '             <ul class="com_r clearfix">';
    html += '               <li class="top clearfix">';
    html += '                 <h4><a href="/' + comment.UserName + '">' + comment.UserName + '</a></h4>';
    html += '                 <span class="time">' + comment.PostTime + '</span>';
    html += '                 <span class="floor_num" floor=' + (floor || index) + '>' + (comment.ParentId > 0 ? " " : index + '楼') + '</span>';
    html += '               </li>';
    html += '               <li class="mid clearfix">';
    html += '                 <div class="comment_p">' + replaceUBBToHTML(comment) + '</div>';
    html += '               </li>';
    html += '               <li class="bot clearfix">';
    html += '                   <div>';
    html += '                     <a href="#reply" class="com_reply btn btn-noborder reply_btn" title="回复" commentid="' + comment.CommentId + '" floor="' + index + '">回复</a>';
//    html += '                     <a href="#quote" class="com_reply" title="引用" commentid="' + comment.CommentId + '" floor="' + index + '">引用</a>';
//    html += '                     <a href="#report" class="com_reply" title="举报" commentid="' + comment.CommentId + '" floor="' + index + '">举报</a>';
    if(!comment.Replies.length==0){
    html += '                     <button class="btn btn-noborder rec_number">' + comment.Replies.length + '条回复<i class="icon iconfont iconfont-xiajiantou"></i></button>';
    }
    if (currentUserName != "" && comment.UserName != "" && comment.UserName.toLowerCase() == currentUserName.toLowerCase()) {
        html += '                 <a href="#delete" class="btn btn-noborder com_reply" commentid="' + comment.CommentId + '" floor="' + floor + '">删除</a>';
    } else if (currentUserName != "" && currentUserName.toLowerCase() == username.toLowerCase()) {
        html += '                 <a href="#delete" class="com_reply" commentid="' + comment.CommentId + '" floor="' + floor + '">删除</a>';
    }
    html += '                   </div>';
    html += '                </li>';
    html += '              </ul>';
    html += '            </dd>';
    html += '          </dl>';

    if (comment.Replies != null) {
            html += '<div class="child_comment" data-listshow="false">';
            html += '   <div class="autoHeight clearfix">';
        for (var j = 0; j < comment.Replies.length; j++) {
            html += getChildItemHtml(comment.Replies[j], j + 1, index, deep + 1);
        }
            html += '   </div>';
            html += '</div>';
    }
    html += '       </div>';
    return html;
}

function getChildItemHtml(comment, index, floor, deep) {

    //var html = '<div class="child_comment" style="height: auto;">';
    //html += '   <div class="autoHeight clearfix">';
    var html = '    <dl class="comment_list clearfix" id="comment_item_' + comment.CommentId + '">';
    html += '      <dt>';
    html += '          <a href="/' + comment.UserName + '">';
    html += '              <img src="' + comment.Userface + '" alt="' + comment.UserName + '">';
    html += '          </a>';
    html += '      </dt>';
    html += '      <dd>';
    html += '        <ul class="com_r clearfix">';
    html += '          <li class="top clearfix">';
    html += '            <h4><a href="/' + comment.UserName + '">' + comment.UserName + '</a></h4>';
    html += '            <span class="time">' + comment.PostTime + '</span>';
    html += '            <span class="floor_num" floor=' + (floor || index) + '>' + '</span>';
    html += '          </li>';
    html += '          <li class="mid clearfix">';
    html += '            <div class="comment_p">' + replaceUBBToHTML(comment) + '</div>';
    html += '          </li>';
    html += '          <li class="mid clearfix">';   
    html += '               <label>';
//    html += '                 <a href="#reply" class="com_reply" commentid="' + comment.CommentId + '" floor="' + floor + '">回复</a>';
//    html += '                 <a href="#quote" class="com_reply" commentid="' + comment.CommentId + '" floor="' + floor + '">引用</a>';
//    html += '                 <a href="#report" class="com_reply" commentid="' + comment.CommentId + '" floor="' + floor + '">举报</a>';
    if (currentUserName != "" && comment.UserName != "" && comment.UserName.toLowerCase() == currentUserName.toLowerCase()) {
        html += '<a href="#delete" class="com_reply btn btn-noborder" commentid="' + comment.CommentId + '" floor="' + floor + '">删除</a>';
    }
    else if (currentUserName != "" && currentUserName.toLowerCase() == username.toLowerCase()) {
        html += '<a href="#delete" class="com_reply btn btn-noborder" commentid="' + comment.CommentId + '" floor="' + floor + '">删除</a>';
    }
    html += '               </label>';

    html += '                </li>';
    html += '              </ul>';
    html += '            </dd>';
    html += '          </dl>';
//  html += '        </div>';
//  html += '      </div>';
    if (comment.Replies != null) {
            html += '<div class="child_comment" data-listshow="false">';
            html += '   <div class="autoHeight clearfix">';
        for (var j = 0; j < comment.Replies.length; j++) {
            html += getChildItemHtml(comment.Replies[j], j + 1, index, deep + 1);
        }
            html += '   </div>';
            html += '</div>';
    }
    return html;
}

//获取评论对象
function getComment(commentId, list) {
    for (var i = 0; i < list.length; i++) {
        var comment = list[i];
        if (comment.CommentId == commentId)
            return comment;
    }
    return null;
}
function paging(start,end,dom){ // 生成按钮
  var num = Math.floor(totalFloor / 5) + (totalFloor % 5 == 0 ? 0 : 1),
      initNum =  num > 5 ? 5 : num;
  var start = start || 0,
      end = end ? end : initNum ,
      dom = dom || '.page_btn_event[page=1]';
  if(totalFloor < 5)return false;
  
  var html = '<div class="pagebox"><div class="page_btn page_header" page="1"><i class="icon iconfont icon-shouye"></i></div><div class="page_btn page_back"><i class="icon iconfont icon-zuojiantou"></i></div>';
  
  for (var i = start; i < end; i++) {
    html += '<div class="page_btn page_btn_event" page='+(i+1)+'>'+(i+1)+'</div>';
  }
  
  html+='<div class="page_btn page_go"><i class="icon iconfont icon-youjiantou"></i></div><div class="page_btn page_footer" page="'+num+'"><i class="icon iconfont icon-weiye"></i></div></div>';
  $('#comment_bar').html(html);
  $(dom).addClass('page_activ');
}
function RedrawnPaging(dom){ // 重绘分页按钮判断
  var num = Math.floor(totalFloor / 5) + (totalFloor % 5 == 0 ? 0 : 1);
  var nextPage = parseInt($(dom).next().attr('page')), // 后一个按钮的页数
      prevPage = parseInt($(dom).prev().attr('page')), // 前一个按钮的页数
      pageNum = parseInt($(dom).attr('page')); // 当前按钮页数
  if(!Boolean(nextPage) && !(nextPage === num)){
    var nextpageNum = (pageNum+2)<num?pageNum+2:num
    var prevpageNum = pageNum === num? pageNum-5:pageNum-3;
    if(pageNum+1 === num && num > 5) prevpageNum = pageNum-4;
    else if(pageNum+1 === num && num < 6) prevpageNum = 0;
    else if(pageNum === num && num < 6) prevpageNum = 0;
    else if(pageNum === num && num > 5) prevpageNum = pageNum-5;
    paging(prevpageNum,nextpageNum,dom);
  }else if(!Boolean(prevPage) && !(prevPage === 0)){
    var prevpageNum = (pageNum-3)>0?pageNum-3:0;
    var nextpageNum = (pageNum+2)<num?pageNum+2:5;
    if(pageNum === 1 && num > 6) nextpageNum = 5;
    else if(pageNum === 1 && num < 5) nextpageNum = num;
    else if(pageNum === 2 && num < 6) nextpageNum = num -pageNum;
    else if(pageNum === 2 && num > 5) nextpageNum = pageNum+3;
    paging(prevpageNum,nextpageNum,dom);
  }
  $('.page_btn_event[page='+pageNum+']').addClass('page_activ')
}
function showcomment(page){ // 重绘评论内容
  //构造主题
  var topics = getTopics(tatolList);
  var listHtml = '';
  // var total = data.total > 0 ? data.total : topics.length;
  //组装HTM
  for (var i = (page-1)*5; i < (page*5); i++) {
      var comment = topics[i];
      var layer = totalFloor - i;
      if(comment instanceof Object) listHtml += getItemHtml(comment, layer);
  }
  listHtml += '<div class="clear"></div>';
  //输出列表
  $(listId).html(listHtml);
  //高亮评论中的代码段
  dp.SyntaxHighlighter.HighlightAll('code2');
  //展示昵称
  new CNick(listId + ' a.username').showNickname();
}
function removeClass(dom,classs){ // 删除选址难得class
  $.each($(dom),function(key, value){
    $(value).removeClass(classs);
  })
}
function page_btn_event(){ // 分页按钮事件
  
  // 页数
  $(document).on('click','.page_btn_event',function(e){
    removeClass('.page_btn_event','page_activ')
    $(this).addClass('page_activ');
    RedrawnPaging(this)
    var pageIndex = parseInt($(this).attr('page'));
    showcomment(pageIndex)
  })
  // 前进后退
  $(document).on('click','.page_back , .page_go',function(){
    var pageNum = parseInt($('.page_activ').attr('page'));
    var num = Math.floor(totalFloor / 5) + (totalFloor % 5 == 0 ? 0 : 1);
    if($(this).attr('class').indexOf('page_back') != -1){
      if(pageNum < 2) return false;
      showcomment(pageNum-1);
      RedrawnPaging($('.page_activ'))
      removeClass('.page_btn_event','page_activ')
      $('.page_btn_event[page='+ (pageNum-1) +']').addClass('page_activ');
    }else if($(this).attr('class').indexOf('page_go') != -1){
      if(pageNum+1 > num) return false;
      showcomment(pageNum+1)
      RedrawnPaging($('.page_activ'))
      removeClass('.page_btn_event','page_activ')
      $('.page_btn_event[page='+ (pageNum+1) +']').addClass('page_activ');
    }
  })
  // 首尾
  $(document).on('click','.page_header,.page_footer',function(){
    var pageNum = parseInt($('.page_activ').attr('page'));
    var num = Math.floor(totalFloor / 5) + (totalFloor % 5 == 0 ? 0 : 1);
    if($(this).attr('class').indexOf('page_header') != -1){
      showcomment(1)
      if(num > 5){
        paging(0,5)
      }else {
        paging(0,num)
      }
      removeClass('.page_btn_event','page_activ')
      $('.page_btn_event[page=1]').addClass('page_activ');
    }else if($(this).attr('class').indexOf('page_footer') != -1){
      showcomment(num)
      if(num > 5){
        paging(num-5,num)
      }else{
        paging(0,num)
      }
      removeClass('.page_btn_event','page_activ')
      $('.page_btn_event[page='+num+']').addClass('page_activ');
    }
  })
}
page_btn_event() 
function deleteCommentList(Id){
  for (var i = 0; i < tatolList.length; i++) {
    if (tatolList[i].CommentId === parseInt(Id)) {
      tatolList.splice(i, 1);
    }
  }
  if(totalFloor >5){
    showcomment(1)
    paging()
  }else{
    showcomment(1)
    if(totalFloor < 6) $('.more_comment').fadeOut(300)
  }
}
function setBtnEvent() {
    $("#load_comments").click(function () {
        // var page = $(this).attr("page");
        // loadList(page);
        if(totalFloor >5){
          showcomment(1)
          paging()
        }else{
          // loadList(2)
          showcomment(1)
          $('.more_comment').fadeOut(300)
        }
    });

    //评论按钮点击
    $(document).on('click',".com_reply",function(){
        var action = $(this).attr("href");

        action = action.substring(action.lastIndexOf('#'));

        var commentId = $(this).attr("commentid");
        var lists = tatolList != [] ? tatolList : list;
        switch (action) {
            case "#reply":
                if (currentUserName) {                   
                    replyComment(commentId, lists);
                    setEditorFocus();
                }
                return true;
            case "#quote":
                if (currentUserName) {                   
                    quoteComment(commentId, list);
                    setEditorFocus();
                }
                return true;
            case "#report":             
                reportComment(commentId, $(this));
                break;
            case "#delete":
                deleteComment(commentId);
                break;
            default:
                return true;
        }
        return false;
    });    
}
/*使评论框获得焦点*/
function setEditorFocus() {
    var val = editor.val();
    editor.val('');
    editor.focus();
    editor.val(val);
}
//引用评论
function quoteComment(commentId, list) {
    var comment = getComment(commentId, list);
    var content = comment.Content;
    if (comment.Content.length > 50) {
        content = comment.Content.substring(0, 50) + "...";
    }
    editor.val("[quote=" + (comment.UserName == null ? "游客" : comment.UserName) + "]" + content + "[/quote]\r\n");
}
//回复评论
function replyComment(commentId, list) {
    var comment = getComment(commentId, list);
    editor.val('[reply]' + comment.UserName + "[/reply]\r\n");
    $("#comment_replyId").val(commentId);
}
//举报评论
function reportComment(commentId, e) {
    report(commentId, 3, e);
}
//删除评论
function deleteComment(commentId) {
    if (!confirm("你确定要删除这篇评论吗？")) return;
    var delUrl = blog_address + "/comment/delete?commentid=" + commentId + "&filename=" + fileName;
    $.get(delUrl, function (data) {
        if (data.result == 1) {
            --totalFloor;
            $("#comment_item_" + commentId).parent().hide().remove();
            // deleteCommentList(commentId)
        } else {
            alert("你没有删除该评论的权限！");
        }
    });
}
//替换评论的UBB代码
function replaceUBBToHTML(comment) {
    var content = $.trim(comment.Content);

    var re = /\[code=([\w#\.]+)\]([\s\S]*?)\[\/code\]/ig;

    var codelist = [];
    while ((mc = re.exec(content)) != null) {
        codelist.push(mc[0]);
        content = content.replace(mc[0], "--code--");
    }
    content = replaceQuote(content);
    //content = content.replace(/\[e(\d+)\]/g, "<img src=\"" + static_host + "/images/emotions/e$1.gif\"\/>");
    content = content.replace(/\[reply]([\s\S]*?)\[\/reply\][\r\n]{0,2}/gi, "回复$1：");
    content = content.replace(/\[url=([^\]]+)]([\s\S]*?)\[\/url\]/gi, '<a href="$1" target="_blank">$2</a>');
    content = content.replace(/\[img(=([^\]]+))?]([\s\S]*?)\[\/img\]/gi, '<img src="$3" style="max-width:400px;max-height:200px;" border="0" title="$2" />');
    //content = content.replace(/\[(\/?)(b|i|u|p)\]/ig, "<$1$2>");
    content = content.replace(/\r?\n/ig, "<br />");

    if (codelist.length > 0) {
        var re1 = /--code--/ig;
        var i = 0;
        while ((mc = re1.exec(content)) != null) {
            content = content.replace(mc[0], codelist[i]);
            i++;
        }
    }
    content = content.replace(/\[code=([\w#\.]+)\]([\s\S]*?)\[\/code\]/ig, function (m0, m1, m2) {
        if ($.trim(m2) == "") return '';
        return '<pre name="code2" class="' + m1 + '">' + m2 + '</pre>';
    });
    return content;
}
//替换评论的引用
function replaceQuote(str) {
    var m = /\[quote=([^\]]+)]([\s\S]*)\[\/quote\]/gi.exec(str);
    if (m) {
        return str.replace(m[0], '<fieldset><legend>引用“' + m[1] + '”的评论：</legend>' + replaceQuote(m[2]) + '</fieldset>');
    } else {
        return str;
    }
}

function load_comment_form() {
    $("#commentbox").hide();
    var un = getcookie("UserName").toLowerCase();
    if (islock) {
        $("#commentsbmitarear").html("<div class='notice'>该文章已被禁止评论！</div>");
    } else if (currentUserName || (un != null&&un!=""&&un!=undefined)) {
        $("#commentbox").show();
        $(".publish_btn").click(function () {
            $("#commentform").submit();
        });       
    } else {
        var curl = encodeURIComponent(location.href);
        $("#commentsbmitarear").html('<div class="guest_link"><span class="log_ico"><i class="icon iconfont icon-yonghu"></i></span><span class="txt">目前您尚未登录，请 ' +
        //'<a href="javascript:void(0);" onclick="javascript:csdn.showLogin(function (dat) {js_logined(dat.data.userName);});">[登录]</a>或' +
        /*'<a href="javascript:void(0);" onclick="javascript:loginbox();">登录</a> 或 ' +
        '<a href="http://passport.csdn.net/account/register?from=' + curl + '">注册</a> 后进行评论</span></div>');*/
        '<a href="https://passport.csdn.net/account/login?from=' + curl + '">登录</a> 或 ' +
        '<a href="http://passport.csdn.net/account/register?from=' + curl + '">注册</a> 后进行评论</span></div>');
    }
    ubb_event(); 
}

function getcookie(name) {
    var cookie_start = document.cookie.indexOf(name);
    var cookie_end = document.cookie.indexOf(";", cookie_start);
    return cookie_start == -1 ? '' : unescape(document.cookie.substring(cookie_start + name.length + 1, (cookie_end > cookie_start ? cookie_end : document.cookie.length)));
}

var c_doing = false;
function subform(e) {
    var isComment = false;
    if (c_doing) return false;
    var content = $.trim($(editorId).val());
    if (content == "") {
        commentTip("评论内容没有填写!");
        return false;
    } else if (content.length > 1000) {
        commentTip("评论内容太长了，不能超过1000个字符！");
        return false;
    }
    var commentId = $("#commentId").val();
    commentTip("正在发表评论...");
    var beginTime = new Date();
    $(editorId).attr("disabled", true);
    $("button[type=submit]", e).attr("disabled", true);
    if($("#comment_replyId").val() === '') isComment = true
    c_doing = true;
    $.ajax({
        type: "POST",
        url: $(e).attr("action"),
        data: {
            "commentid": commentId,
            "content": content,
            "replyId": $("#comment_replyId").val(),
            "boleattohome": $("#boleattohome").val()
        },
        success: function (data) {
            c_doing = false;
            commentTip(data.content);
            if (data.result) {
                var rcommentid=$("#comment_replyId").val();
                $(editorId).val('');
                $("#comment_replyId,#comment_verifycode").val('');

                commentscount++;
                if(isComment) totalFloor = totalFloor+1;
                loadList(1, true);
                isComment = false;
                $(editorId).attr("disabled", false);
                $("button[type=submit]", e).attr("disabled", false);

                commentTip("发表成功！评论耗时:" + (new Date() - beginTime) + "毫秒");

                if (rcommentid!=undefined && rcommentid != "")
                {
                    $("html,body").animate({ scrollTop: $("#comment_item_" + rcommentid).offset().top }, 1000);
                }
                
            }
        }
    });
    return false;
}

//操作提示
var _c_t;
function commentTip(message) {
    $("#tip_comment").html(message).show();
    clearTimeout(_c_t);
    _c_t = setTimeout(function () {
        $("#tip_comment").hide();
    }, 10000);
}

function ubb_event() {
    //ubb按钮事件
    $(document).on('click',"#lang_list",function(e){
        editor = $(editorId);
        var selectedValue = editor.selection();
        editor.focus();
        var code = $(this).attr("code");
        var targetDOM =e.target
        switch (code) {
            case "code":
                var lang_list = $("#lang_list");
                editor.val("[code=" + $.trim(targetDOM.href.split('#')[1]) + "]\n" + selectedValue + "\n[/code]");                        
                break;
            default:
                editor.val("[" + code + "]" + selectedValue + "[/" + code + "]");
                break;
        }
        return false;
    });

   
    /*editor = $(editorId);
    $("#lang_list").children().each(function () {       
        var selectedValue = editor.selection();
        //editor.focus();

        $(this).unbind("click").click(function () {
            editor.val("[code=" + $.trim(this.href.split('#')[1]) + "]\n" + selectedValue + "\n[/code]");
            setTimeout(function () { $("#lang_list").hide(); }, 200);
        });
    });*/
}
//comment end

//create left directory begin
$(document).ready(function () {
    var hs = $('#article_content').find('h1,h2,h3');
    if(hs.length <= 2){
        $('#blog_artical_directory').hide();
        $('.left_menu .menu_con').hide();
    }
    var h = parseInt($("#csdnBlogDir").height());
    var wrapH = parseInt($(".first_li").height())
    $(".left_menu_btn").on("click",function(){
      if (!$('.first_li')[0]) {
        csdnBlogDirectory(hs);//创建目录
      }

        //top and down show or hide
        var $list_father = $('.list_father'),
            btnWrap = $(".arr_box");
        if(wrapH <= h){
            btnWrap.hide();
            $list_father.height(wrapH+10)
        }else{
            btnWrap.show();
        }
    });
});

function csdnBlogDirectory(hs) {
    // var hs = $('#article_content').find('h1,h2,h3');
    if (hs.length < 2) return;
    var s = '';
    s += '<ol class="first_li">';
    var old_h = 0, ol_cnt = 0;
    var h1Num = 0 , h2Num = 0 , h3Num =0; //记录标题数量
    for (var i = 0; i < hs.length; i++) {
        var h = parseInt(hs[i].tagName.substr(1), 10); // h标签后的数字1，2，3，4，5，6
        if (!old_h) old_h = h;
        if (h > old_h) { s += '<ol class="second_li">'; ol_cnt++; }
        else if (h < old_h && ol_cnt > 0) { s += '</ol>'; ol_cnt--; }
        if (h == 1) {
            while (ol_cnt > 0) { s += '</ol>'; ol_cnt--; }
        }
        old_h = h;
        var tit = hs.eq(i).text().replace(/^[\d.、\s]+/g, '');
        tit = tit.replace(/[^a-zA-Z0-9_\-\s\u4e00-\u9fa5]+/g, '');
        if (tit.length < 100) {
          switch(h){
            case 1:
              ++h1Num,h2Num =0,h3Num=0;
              
              s += '<li><a href="#t' + i + '">'+ h1Num + '. '+ tit + '</a></li>';
              hs.eq(i).html('<a name="t' + i + '"></a>' + hs.eq(i).html());
              break;
              case 2:
                ++h2Num,h3Num=0;
                s += '<li><a href="#t' + i + '">'+h1Num + '-'+ h2Num + '. '+ tit + '</a></li>';
                hs.eq(i).html('<a name="t' + i + '"></a>' + hs.eq(i).html());
                break;
                case 3:
                  ++h3Num
                  s += '<li><a href="#t' + i + '">'+h1Num + '-'+ h2Num + '-'+h3Num +'. '+ tit + '</a></li>';
                  hs.eq(i).html('<a name="t' + i + '"></a>' + hs.eq(i).html());
                  break;
          }
            // s += '<li><a href="#t' + i + '">' + tit + '</a></li>';
            // hs.eq(i).html('<a name="t' + i + '"></a>' + hs.eq(i).html());
        }
    }
    while (ol_cnt > 0) {
        s += '</ol>';
        ol_cnt--;
    }

    $('#csdnBlogDir').html(s);
}
function openct(e) {
    if (e.innerHTML == '[+]') {
        $(e).attr('title', '收起').html('[-]').parent().next().show();
    } else {
        $(e).attr('title', '展开').html('[+]').parent().next().hide();
    }
    e.blur();
    return false;
}
//create left directory end

/*SyntaxHighlighter/shCore-src.js*/
var a, dp = { sh: { Toolbar: {}, Utils: {}, RegexLib: {}, Brushes: {}, Strings: { AboutDialog: '<html><head><title>About...</title></head><body class="dp-about"><table cellspacing="0"><tr><td class="copy"><p class="title">dp.SyntaxHighlighter</div><div class="para">Version: {V}</p><p><a href="http://www.dreamprojections.com/syntaxhighlighter/?ref=about" target="_blank">http://www.dreamprojections.com/syntaxhighlighter</a></p>&copy;2004-2007 Alex Gorbatchev.</td></tr><tr><td class="footer"><input type="button" class="close" value="OK" onClick="window.close()"/></td></tr></table></body></html>' }, ClipboardSwf: 'http://static.blog.csdn.net/scripts/ZeroClipboard/ZeroClipboard.swf', Version: "1.5.1" } };
dp.SyntaxHighlighter = dp.sh;
dp.sh.Toolbar.Commands = {
    ExpandSource: {
        label: "+ expand source", check: function (b) {
            return b.collapse
        }, func: function (b, c) {
            b.parentNode.removeChild(b);
            c.div.className = c.div.className.replace("collapsed", "");
        }
    }, ViewSource: {
        label: "view plain", func: function (b, c) {
            b = dp.sh.Utils.FixForBlogger(c.originalCode).replace(/</g, "&lt;");
            c = window.open("", "_blank", "width=750, height=400, location=0, resizable=1, menubar=0, scrollbars=0");
            c.document.write('<textarea style="width:99%;height:99%">' + b + "</textarea>");
            c.document.close();
        }
    }, CopyToClipboard: {
        label: "copy", check: function () {
            return window.clipboardData != null || dp.sh.ClipboardSwf != null
        }, func: function (b, c) {
            b = dp.sh.Utils.FixForBlogger(c.originalCode).replace(/&lt;/g, "<").replace(/&gt;/g, ">").replace(/&amp;/g, "&");
            if (window.clipboardData)
                window.clipboardData.setData("text", b);
            else if (dp.sh.ClipboardSwf != null) {
                var d = c.flashCopier;
                if (d == null) {
                    d = document.createElement("div");
                    c.flashCopier = d;
                    c.div.appendChild(d);
                }
                d.innerHTML = '<embed src="' + dp.sh.ClipboardSwf + '" FlashVars="clipboard=' + encodeURIComponent(b) + '" width="0" height="0" type="application/x-shockwave-flash"></embed>';
            }
            alert("The code is in your clipboard now");
        }
    }, PrintSource: {
        label: "print", func: function (b, c) {
            b = document.createElement("IFRAME");
            var d = null;
            b.style.cssText = "position:absolute;width:0px;height:0px;left:-500px;top:-500px;";
            document.body.appendChild(b);
            d = b.contentWindow.document;
            dp.sh.Utils.CopyStyles(d, window.document);
            d.write('<div class="' + c.div.className.replace("collapsed", "") + ' printing">' + c.div.innerHTML + "</div>");
            d.close();
            b.contentWindow.focus();
            b.contentWindow.print();
            alert("Printing...");
            document.body.removeChild(b);
        }
    }, About: {
        label: "?", func: function () {
            var b = window.open("", "_blank", "dialog,width=300,height=150,scrollbars=0"), c = b.document;
            dp.sh.Utils.CopyStyles(c, window.document);
            c.write(dp.sh.Strings.AboutDialog.replace("{V}", dp.sh.Version));
            c.close();
            b.focus()
        }
    }
};
dp.sh.Toolbar.Create = function (b) {
    var _code = b.source.className;
    var c = document.createElement("DIV");
    c.className = "tools";
    c.innerHTML = '<b>[' + _code + ']</b> ';
    for (var d in dp.sh.Toolbar.Commands) {
        var f = dp.sh.Toolbar.Commands[d];
        //   f.check != null && !f.check(b) || (c.innerHTML += '<a href="#" class="' + d + '" title="' + f.label + '" onclick="dp.sh.Toolbar.Command(\'' + d + "',this);return false;\">" + f.label + "</a>")
        if (f.label == "print") {
            f.check != null && !f.check(b) || (c.innerHTML += '<span class="tracking-ad" data-mod="popu_169"> <a href="#" class="' + d + '" title="' + f.label + '" onclick="dp.sh.Toolbar.Command(\'' + d + "',this);return false;\">" + f.label + "</a></span>")
        }
        else if (f.label == "copy") {
            f.check != null && !f.check(b) || (c.innerHTML += '<span class="tracking-ad" data-mod="popu_168"> <a href="#" class="' + d + '" title="' + f.label + '" onclick="dp.sh.Toolbar.Command(\'' + d + "',this);return false;\">" + f.label + "</a></span>")
        }
        else {
            f.check != null && !f.check(b) || (c.innerHTML += '<a href="#" class="' + d + '" title="' + f.label + '" onclick="dp.sh.Toolbar.Command(\'' + d + "',this);return false;\">" + f.label + "</a>")
        }
    }
    return c;
};
dp.sh.Toolbar.Command = function (b, c) {
    for (var d = c; d != null && d.className.indexOf("dp-highlighter") == -1;)
        d = d.parentNode;
    d != null && dp.sh.Toolbar.Commands[b].func(c, d.highlighter)
};
dp.sh.Utils.CopyStyles = function (b, c) {
    c = c.getElementsByTagName("link");
    for (var d = 0; d < c.length; d++)
        c[d].rel.toLowerCase() == "stylesheet" && b.write('<link type="text/css" rel="stylesheet" href="' + c[d].href + '"></link>')
};
dp.sh.Utils.FixForBlogger = function (b) {
    return dp.sh.isBloggerMode == true ? b.replace(/<br\s*\/?>|&lt;br\s*\/?&gt;/gi, "\n") : b
};
dp.sh.RegexLib = { MultiLineCComments: new RegExp("/\\*[\\s\\S]*?\\*/", "gm"), SingleLineCComments: new RegExp("//.*$", "gm"), SingleLinePerlComments: new RegExp("#.*$", "gm"), DoubleQuotedString: new RegExp('"(?:\\.|(\\\\\\")|[^\\""\\n])*"', "g"), SingleQuotedString: new RegExp("'(?:\\.|(\\\\\\')|[^\\''\\n])*'", "g") };
dp.sh.Match = function (b, c, d) {
    this.value = b;
    this.index = c;
    this.length = b.length;
    this.css = d
};
dp.sh.Highlighter = function () {
    this.noGutter = false;
    this.addControls = true;
    this.collapse = false;
    this.tabsToSpaces = true;
    this.wrapColumn = 80;
    this.showColumns = true
};
dp.sh.Highlighter.SortCallback = function (b, c) {
    if (b.index < c.index)
        return -1;
    else if (b.index > c.index)
        return 1;
    else if (b.length < c.length)
        return -1;
    else if (b.length > c.length)
        return 1;
    return 0
};
a = dp.sh.Highlighter.prototype;
a.CreateElement = function (b) {
    b = document.createElement(b);
    b.highlighter = this;
    return b
};
a.GetMatches = function (b, c) {
    for (var d = null; (d = b.exec(this.code)) != null;)
        this.matches[this.matches.length] = new dp.sh.Match(d[0], d.index, c)
};
a.AddBit = function (b, c) {
    if (!(b == null || b.length == 0)) {
        var d = this.CreateElement("SPAN");
        b = b.replace(/ /g, "&nbsp;");
        b = b.replace(/</g, "&lt;");
        b = b.replace(/(\r?\n)|(\[BR\])/gm, "&nbsp;<br>");
        if (c != null)
            if (/br/gi.test(b)) {
                b = b.split("&nbsp;<br>");
                for (var f = 0; f < b.length; f++) {
                    d = this.CreateElement("SPAN");
                    d.className = c;
                    d.innerHTML = b[f];
                    this.div.appendChild(d);
                    f + 1 < b.length && this.div.appendChild(this.CreateElement("BR"))
                }
            } else {
                d.className = c;
                d.innerHTML = b;
                this.div.appendChild(d)
            }
        else {
            d.innerHTML = b;
            this.div.appendChild(d)
        }
    }
};
a.IsInside = function (b) {
    if (b == null || b.length == 0)
        return false;
    for (var c = 0; c < this.matches.length; c++) {
        var d = this.matches[c];
        if (d != null)
            if (b.index > d.index && b.index < d.index + d.length)
                return true
    }
    return false
};
a.ProcessRegexList = function () {
    for (var b = 0; b < this.regexList.length; b++)
        this.GetMatches(this.regexList[b].regex, this.regexList[b].css)
};
a.ProcessSmartTabs = function (b) {
    function c(h, e, l) {
        var m = h.substr(0, e);
        h = h.substr(e + 1, h.length);
        e = "";
        for (var i = 0; i < l; i++)
            e += " ";
        return m + e + h
    }
    function d(h, e) {
        if (h.indexOf(p) == -1)
            return h;
        for (var l = 0; (l = h.indexOf(p)) != -1;)
            h = c(h, l, e - l % e);
        return h
    }
    b = b.split("\n");
    for (var f = "", p = "\t", q = 0; q < b.length; q++)
        f += d(b[q], 4) + "\n";
    return f
};
a.SwitchToList = function () {
    var b = this.div.innerHTML.replace(/<(br)\/?>/gi, "\n").split("\n");
    this.addControls == true && this.bar.appendChild(dp.sh.Toolbar.Create(this));
    if (this.showColumns) {
        for (var c = this.CreateElement("div"), d = this.CreateElement("div"), f = 1; f <= 150;)
            if (f % 10 == 0) {
                c.innerHTML += f;
                f += (f + "").length
            } else {
                c.innerHTML += "&middot;";
                f++
            }
        d.className = "columns";
        d.appendChild(c);
        this.bar.appendChild(d)
    }
    f = 0;
    for (c = this.firstLine; f < b.length - 1; f++, c++) {
        d = this.CreateElement("LI");
        var p = this.CreateElement("SPAN");
        d.className = f % 2 == 0 ? "alt" : "";
        p.innerHTML = b[f] + "&nbsp;";
        d.appendChild(p);
        this.ol.appendChild(d)
    }
    this.div.innerHTML = "";
};
a.Highlight = function (b) {
    function c(e) {
        return e.replace(/^\s*(.*?)[\s\n]*$/g, "$1")
    }
    function d(e) {
        return e.replace(/\n*$/, "").replace(/^\n*/, "")
    }
    function f(e) {
        e = dp.sh.Utils.FixForBlogger(e).split("\n");
        for (var l = new RegExp("^\\s*", "g"), m = 1E3, i = 0; i < e.length && m > 0; i++)
            if (c(e[i]).length != 0) {
                var g = l.exec(e[i]);
                if (g != null && g.length > 0)
                    m = Math.min(g[0].length, m)
            }
        if (m > 0)
            for (i = 0; i < e.length; i++)
                e[i] = e[i].substr(m);
        return e.join("\n");
    }
    function p(e, l, m) {
        return e.substr(l, m - l)
    }
    var q = 0;
    if (b == null)
        b = "";
    this.originalCode = b;
    this.code = d(f(b));
    this.div = this.CreateElement("DIV");
    this.bar = this.CreateElement("DIV");
    this.ol = this.CreateElement("OL");
    this.matches = [];
    this.div.className = "dp-highlighter";
    this.div.highlighter = this;
    this.bar.className = "bar";
    this.ol.start = this.firstLine;
    if (this.CssClass != null)
        this.ol.className = this.CssClass;
    if (this.collapse)
        this.div.className += " collapsed";
    if (this.noGutter)
        this.div.className += " nogutter";
    if (this.tabsToSpaces == true)
        this.code = this.ProcessSmartTabs(this.code);
    this.ProcessRegexList();
    if (this.matches.length == 0)
        this.AddBit(this.code, null);
    else {
        this.matches = this.matches.sort(dp.sh.Highlighter.SortCallback);
        for (b = 0; b < this.matches.length; b++)
            if (this.IsInside(this.matches[b]))
                this.matches[b] = null;
        for (b = 0; b < this.matches.length; b++) {
            var h = this.matches[b];
            if (!(h == null || h.length == 0)) {
                this.AddBit(p(this.code, q, h.index), null);
                this.AddBit(h.value, h.css);
                q = h.index + h.length
            }
        }
        this.AddBit(this.code.substr(q), null)
    }
    this.SwitchToList();
    this.div.appendChild(this.bar);
    this.div.appendChild(this.ol)
};
a.GetKeywords = function (b) {
    return "\\b" + b.replace(/ /g, "\\b|\\b") + "\\b"
};
dp.sh.BloggerMode = function () {
    dp.sh.isBloggerMode = true
};
dp.sh.HighlightAll = function (b, c, d, f, p, q) {
    function h() {
        for (var k = arguments, j = 0; j < k.length; j++)
            if (k[j] != null) {
                if (typeof k[j] == "string" && k[j] != "")
                    return k[j] + "";
                if (typeof k[j] == "object" && k[j].value != "")
                    return k[j].value + ""
            }
        return null
    }
    function e(k, j) {
        for (var o = 0; o < j.length; o++)
            if (j[o] == k)
                return true;
        return false
    }
    function l(k, j, o) {
        k = new RegExp("^" + k + "\\[(\\w+)\\]$", "gi");
        for (var s = null, u = 0; u < j.length; u++)
            if ((s = k.exec(j[u])) != null)
                return s[1];
        return o
    }
    function m(k, j, o) {
        o = document.getElementsByTagName(o);
        for (var s = 0; s < o.length; s++) {
            o[s].getAttribute("name") == j && k.push(o[s]);
        }
    }
    var i = [], g = null, v = {};
    m(i, b, "pre");
    m(i, b, "textarea");
    if (i.length != 0) {
        for (var n in dp.sh.Brushes) {
            g = dp.sh.Brushes[n].Aliases;
            if (g != null)
                for (b = 0; b < g.length; b++)
                    v[g[b]] = n
        }
        for (b = 0; b < i.length; b++) {
            n = i[b];
            var r = h(n.attributes["class"], n.className, n.attributes.language, n.language);
            g = "";
            if (r != null) {
                r = r.split(":");
                g = r[0].toLowerCase();
                if (v[g] != null) {
                    g = new dp.sh.Brushes[v[g]];
                    n.style.display = "none";
                    g.noGutter = c == null ? e("nogutter", r) : !c;
                    g.addControls = d == null ? !e("nocontrols", r) : d;
                    g.collapse = f == null ? e("collapse", r) : f;
                    g.showColumns = q == null ? e("showcolumns", r) : q;
                    var w = document.getElementsByTagName("head")[0];
                    if (g.Style && w) {
                        var t = document.createElement("style");
                        t.setAttribute("type", "text/css");
                        if (t.styleSheet)
                            t.styleSheet.cssText = g.Style;
                        else {
                            var x = document.createTextNode(g.Style);
                            t.appendChild(x)
                        }
                        w.appendChild(t);
                    }
                    g.firstLine = p == null ? parseInt(l("firstline", r, 1)) : p;
                    g.source = n;
                    g.Highlight(n.innerHTML);
                    g.div.className += " bg_" + n.className;
                    n.parentNode.insertBefore(g.div, n);
                }
            }
        }
    }
};
dp.sh.Brushes.Xml = function () {
    this.CssClass = "dp-xml";
    this.Style = ""
};
dp.sh.Brushes.Xml.prototype = new dp.sh.Highlighter;
dp.sh.Brushes.Xml.Aliases = ["xml", "xhtml", "xslt", "html", "xhtml"];
dp.sh.Brushes.Xml.prototype.ProcessRegexList = function () {
    function c(d, e) {
        d[d.length] = e
    }
    var a = null, b = null;
    this.GetMatches(new RegExp("(&lt;|<)\\!\\[[\\w\\s]*?\\[(.|\\s)*?\\]\\](&gt;|>)", "gm"), "cdata");
    this.GetMatches(new RegExp("(&lt;|<)!--\\s*.*?\\s*--(&gt;|>)", "gm"), "comments");
    for (b = new RegExp("([:\\w-.]+)\\s*=\\s*(\".*?\"|'.*?'|\\w+)*|(\\w+)", "gm") ; (a = b.exec(this.code)) != null;)
        if (a[1] != null) {
            c(this.matches, new dp.sh.Match(a[1], a.index, "attribute"));
            a[2] != undefined && c(this.matches, new dp.sh.Match(a[2], a.index + a[1].length + a[0].substr(a[1].length).indexOf(a[2]), "attribute-value"))
        }
    this.GetMatches(new RegExp("(&lt;|<)/*\\?*(?!\\!)|/*\\?*(&gt;|>)", "gm"), "tag");
    for (b = new RegExp("(?:&lt;|<)/*\\?*\\s*([:\\w-.]+)", "gm") ; (a = b.exec(this.code)) != null;)
        c(this.matches, new dp.sh.Match(a[1], a.index + a[0].indexOf(a[1]), "tag-name"))
};
dp.sh.Brushes.Vb = function () {
    this.regexList = [{ regex: new RegExp("'.*$", "gm"), css: "comment" }, { regex: dp.sh.RegexLib.DoubleQuotedString, css: "string" }, { regex: new RegExp("^\\s*#.*", "gm"), css: "preprocessor" }, { regex: new RegExp(this.GetKeywords("AddHandler AddressOf AndAlso Alias And Ansi As Assembly Auto Boolean ByRef Byte ByVal Call Case Catch CBool CByte CChar CDate CDec CDbl Char CInt Class CLng CObj Const CShort CSng CStr CType Date Decimal Declare Default Delegate Dim DirectCast Do Double Each Else ElseIf End Enum Erase Error Event Exit False Finally For Friend Function Get GetType GoSub GoTo Handles If Implements Imports In Inherits Integer Interface Is Let Lib Like Long Loop Me Mod Module MustInherit MustOverride MyBase MyClass Namespace New Next Not Nothing NotInheritable NotOverridable Object On Option Optional Or OrElse Overloads Overridable Overrides ParamArray Preserve Private Property Protected Public RaiseEvent ReadOnly ReDim REM RemoveHandler Resume Return Select Set Shadows Shared Short Single Static Step Stop String Structure Sub SyncLock Then Throw To True Try TypeOf Unicode Until Variant When While With WithEvents WriteOnly Xor"), "gm"), css: "keyword" }];
    this.CssClass = "dp-vb"
};
dp.sh.Brushes.Vb.prototype = new dp.sh.Highlighter;
dp.sh.Brushes.Vb.Aliases = ["vb", "vb.net"];
dp.sh.Brushes.Sql = function () {
    this.regexList = [{ regex: new RegExp("--(.*)$", "gm"), css: "comment" }, { regex: dp.sh.RegexLib.DoubleQuotedString, css: "string" }, { regex: dp.sh.RegexLib.SingleQuotedString, css: "string" }, { regex: new RegExp(this.GetKeywords("abs avg case cast coalesce convert count current_timestamp current_user day isnull left lower month nullif replace right session_user space substring sum system_user upper user year"), "gmi"), css: "func" }, { regex: new RegExp(this.GetKeywords("all and any between cross in join like not null or outer some"), "gmi"), css: "op" }, { regex: new RegExp(this.GetKeywords("absolute action add after alter as asc at authorization begin bigint binary bit by cascade char character check checkpoint close collate column commit committed connect connection constraint contains continue create cube current current_date current_time cursor database date deallocate dec decimal declare default delete desc distinct double drop dynamic else end end-exec escape except exec execute false fetch first float for force foreign forward free from full function global goto grant group grouping having hour ignore index inner insensitive insert instead int integer intersect into is isolation key last level load local max min minute modify move name national nchar next no numeric of off on only open option order out output partial password precision prepare primary prior privileges procedure public read real references relative repeatable restrict return returns revoke rollback rollup rows rule schema scroll second section select sequence serializable set size smallint static statistics table temp temporary then time timestamp to top transaction translation trigger true truncate uncommitted union unique update values varchar varying view when where with work"), "gmi"), css: "keyword" }];
    this.CssClass = "dp-sql";
    this.Style = ""
};
dp.sh.Brushes.Sql.prototype = new dp.sh.Highlighter;
dp.sh.Brushes.Sql.Aliases = ["sql"];
dp.sh.Brushes.Ruby = function () {
    this.regexList = [{ regex: dp.sh.RegexLib.SingleLinePerlComments, css: "comment" }, { regex: dp.sh.RegexLib.DoubleQuotedString, css: "string" }, { regex: dp.sh.RegexLib.SingleQuotedString, css: "string" }, { regex: new RegExp(":[a-z][A-Za-z0-9_]*", "g"), css: "symbol" }, { regex: new RegExp("(\\$|@@|@)\\w+", "g"), css: "variable" }, { regex: new RegExp(this.GetKeywords("alias and BEGIN begin break case class def define_method defined do each else elsif END end ensure false for if in module new next nil not or raise redo rescue retry return self super then throw true undef unless until when while yield"), "gm"), css: "keyword" }, { regex: new RegExp(this.GetKeywords("Array Bignum Binding Class Continuation Dir Exception FalseClass File::Stat File Fixnum Fload Hash Integer IO MatchData Method Module NilClass Numeric Object Proc Range Regexp String Struct::TMS Symbol ThreadGroup Thread Time TrueClass"), "gm"), css: "builtin" }];
    this.CssClass = "dp-rb";
    this.Style = ""
};
dp.sh.Brushes.Ruby.prototype = new dp.sh.Highlighter;
dp.sh.Brushes.Ruby.Aliases = ["ruby", "rails", "ror"];
dp.sh.Brushes.Python = function () {
    this.regexList = [{ regex: dp.sh.RegexLib.SingleLinePerlComments, css: "comment" }, { regex: new RegExp("^\\s*@\\w+", "gm"), css: "decorator" }, { regex: new RegExp("(['\"]{3})([^\\1])*?\\1", "gm"), css: "comment" }, { regex: new RegExp('"(?!")(?:\\.|\\\\\\"|[^\\""\\n\\r])*"', "gm"), css: "string" }, { regex: new RegExp("'(?!')*(?:\\.|(\\\\\\')|[^\\''\\n\\r])*'", "gm"), css: "string" }, { regex: new RegExp("\\b\\d+\\.?\\w*", "g"), css: "number" }, { regex: new RegExp(this.GetKeywords("and assert break class continue def del elif else except exec finally for from global if import in is lambda not or pass print raise return try yield while"), "gm"), css: "keyword" }, { regex: new RegExp(this.GetKeywords("None True False self cls class_"), "gm"), css: "special" }];
    this.CssClass = "dp-py";
    this.Style = ""
};
dp.sh.Brushes.Python.prototype = new dp.sh.Highlighter;
dp.sh.Brushes.Python.Aliases = ["py", "python"];
dp.sh.Brushes.Plain = function () {
    this.regexList = []
};
dp.sh.Brushes.Plain.prototype = new dp.sh.Highlighter;
dp.sh.Brushes.Plain.Aliases = ["plain", "text", "txt"];
dp.sh.Brushes.Php = function () {
    this.regexList = [{ regex: dp.sh.RegexLib.SingleLineCComments, css: "comment" }, { regex: dp.sh.RegexLib.MultiLineCComments, css: "comment" }, { regex: dp.sh.RegexLib.DoubleQuotedString, css: "string" }, { regex: dp.sh.RegexLib.SingleQuotedString, css: "string" }, { regex: new RegExp("\\$\\w+", "g"), css: "vars" }, { regex: new RegExp(this.GetKeywords("abs acos acosh addcslashes addslashes array_change_key_case array_chunk array_combine array_count_values array_diff array_diff_assoc array_diff_key array_diff_uassoc array_diff_ukey array_fill array_filter array_flip array_intersect array_intersect_assoc array_intersect_key array_intersect_uassoc array_intersect_ukey array_key_exists array_keys array_map array_merge array_merge_recursive array_multisort array_pad array_pop array_product array_push array_rand array_reduce array_reverse array_search array_shift array_slice array_splice array_sum array_udiff array_udiff_assoc array_udiff_uassoc array_uintersect array_uintersect_assoc array_uintersect_uassoc array_unique array_unshift array_values array_walk array_walk_recursive atan atan2 atanh base64_decode base64_encode base_convert basename bcadd bccomp bcdiv bcmod bcmul bindec bindtextdomain bzclose bzcompress bzdecompress bzerrno bzerror bzerrstr bzflush bzopen bzread bzwrite ceil chdir checkdate checkdnsrr chgrp chmod chop chown chr chroot chunk_split class_exists closedir closelog copy cos cosh count count_chars date decbin dechex decoct deg2rad delete ebcdic2ascii echo empty end ereg ereg_replace eregi eregi_replace error_log error_reporting escapeshellarg escapeshellcmd eval exec exit exp explode extension_loaded feof fflush fgetc fgetcsv fgets fgetss file_exists file_get_contents file_put_contents fileatime filectime filegroup fileinode filemtime fileowner fileperms filesize filetype floatval flock floor flush fmod fnmatch fopen fpassthru fprintf fputcsv fputs fread fscanf fseek fsockopen fstat ftell ftok getallheaders getcwd getdate getenv gethostbyaddr gethostbyname gethostbynamel getimagesize getlastmod getmxrr getmygid getmyinode getmypid getmyuid getopt getprotobyname getprotobynumber getrandmax getrusage getservbyname getservbyport gettext gettimeofday gettype glob gmdate gmmktime ini_alter ini_get ini_get_all ini_restore ini_set interface_exists intval ip2long is_a is_array is_bool is_callable is_dir is_double is_executable is_file is_finite is_float is_infinite is_int is_integer is_link is_long is_nan is_null is_numeric is_object is_readable is_real is_resource is_scalar is_soap_fault is_string is_subclass_of is_uploaded_file is_writable is_writeable mkdir mktime nl2br parse_ini_file parse_str parse_url passthru pathinfo readlink realpath rewind rewinddir rmdir round str_ireplace str_pad str_repeat str_replace str_rot13 str_shuffle str_split str_word_count strcasecmp strchr strcmp strcoll strcspn strftime strip_tags stripcslashes stripos stripslashes stristr strlen strnatcasecmp strnatcmp strncasecmp strncmp strpbrk strpos strptime strrchr strrev strripos strrpos strspn strstr strtok strtolower strtotime strtoupper strtr strval substr substr_compare"), "gmi"), css: "func" }, { regex: new RegExp(this.GetKeywords("and or xor __FILE__ __LINE__ array as break case cfunction class const continue declare default die do else elseif enddeclare endfor endforeach endif endswitch endwhile extends for foreach function include include_once global if new old_function return static switch use require require_once var while __FUNCTION__ __CLASS__ __METHOD__ abstract interface public implements extends private protected throw"), "gm"), css: "keyword" }];
    this.CssClass = "dp-c"
};
dp.sh.Brushes.Php.prototype = new dp.sh.Highlighter;
dp.sh.Brushes.Php.Aliases = ["php"];
dp.sh.Brushes.JScript = function () {
    this.regexList = [{ regex: dp.sh.RegexLib.SingleLineCComments, css: "comment" }, { regex: dp.sh.RegexLib.MultiLineCComments, css: "comment" }, { regex: dp.sh.RegexLib.DoubleQuotedString, css: "string" }, { regex: dp.sh.RegexLib.SingleQuotedString, css: "string" }, { regex: new RegExp("^\\s*#.*", "gm"), css: "preprocessor" }, { regex: new RegExp(this.GetKeywords("abstract boolean break byte case catch char class const continue debugger default delete do double else enum export extends false final finally float for function goto if implements import in instanceof int interface long native new null package private protected public return short static super switch synchronized this throw throws transient true try typeof var void volatile while with"), "gm"), css: "keyword" }];
    this.CssClass = "dp-c"
};
dp.sh.Brushes.JScript.prototype = new dp.sh.Highlighter;
dp.sh.Brushes.JScript.Aliases = ["js", "jscript", "javascript"];
dp.sh.Brushes.Java = function () {
    this.regexList = [{ regex: dp.sh.RegexLib.SingleLineCComments, css: "comment" }, { regex: dp.sh.RegexLib.MultiLineCComments, css: "comment" }, { regex: dp.sh.RegexLib.DoubleQuotedString, css: "string" }, { regex: dp.sh.RegexLib.SingleQuotedString, css: "string" }, { regex: new RegExp("\\b([\\d]+(\\.[\\d]+)?|0x[a-f0-9]+)\\b", "gi"), css: "number" }, { regex: new RegExp("(?!\\@interface\\b)\\@[\\$\\w]+\\b", "g"), css: "annotation" }, { regex: new RegExp("\\@interface\\b", "g"), css: "keyword" }, { regex: new RegExp(this.GetKeywords("abstract assert boolean break byte case catch char class const continue default do double else enum extends false final finally float for goto if implements import instanceof int interface long native new null package private protected public return short static strictfp super switch synchronized this throw throws true transient try void volatile while"), "gm"), css: "keyword" }];
    this.CssClass = "dp-j";
    this.Style = ""
};
dp.sh.Brushes.Java.prototype = new dp.sh.Highlighter;
dp.sh.Brushes.Java.Aliases = ["java"];
dp.sh.Brushes.Delphi = function () {
    this.regexList = [{ regex: new RegExp("\\(\\*[\\s\\S]*?\\*\\)", "gm"), css: "comment" }, { regex: new RegExp("{(?!\\$)[\\s\\S]*?}", "gm"), css: "comment" }, { regex: dp.sh.RegexLib.SingleLineCComments, css: "comment" }, { regex: dp.sh.RegexLib.SingleQuotedString, css: "string" }, { regex: new RegExp("\\{\\$[a-zA-Z]+ .+\\}", "g"), css: "directive" }, { regex: new RegExp("\\b[\\d\\.]+\\b", "g"), css: "number" }, { regex: new RegExp("\\$[a-zA-Z0-9]+\\b", "g"), css: "number" }, { regex: new RegExp(this.GetKeywords("abs addr and ansichar ansistring array as asm begin boolean byte cardinal case char class comp const constructor currency destructor div do double downto else end except exports extended false file finalization finally for function goto if implementation in inherited int64 initialization integer interface is label library longint longword mod nil not object of on or packed pansichar pansistring pchar pcurrency pdatetime pextended pint64 pointer private procedure program property pshortstring pstring pvariant pwidechar pwidestring protected public published raise real real48 record repeat set shl shortint shortstring shr single smallint string then threadvar to true try type unit until uses val var varirnt while widechar widestring with word write writeln xor"), "gm"), css: "keyword" }];
    this.CssClass = "dp-delphi";
    this.Style = ""
};
dp.sh.Brushes.Delphi.prototype = new dp.sh.Highlighter;
dp.sh.Brushes.Delphi.Aliases = ["delphi", "pascal"];
dp.sh.Brushes.CSS = function () {
    this.regexList = [{ regex: dp.sh.RegexLib.MultiLineCComments, css: "comment" }, { regex: dp.sh.RegexLib.DoubleQuotedString, css: "string" }, { regex: dp.sh.RegexLib.SingleQuotedString, css: "string" }, { regex: new RegExp("\\#[a-zA-Z0-9]{3,6}", "g"), css: "value" }, { regex: new RegExp("(-?\\d+)(.\\d+)?(px|em|pt|:|%|)", "g"), css: "value" }, { regex: new RegExp("!important", "g"), css: "important" }, { regex: new RegExp(this.GetKeywordsCSS("ascent azimuth background-attachment background-color background-image background-position background-repeat background baseline bbox border-collapse border-color border-spacing border-style border-top border-right border-bottom border-left border-top-color border-right-color border-bottom-color border-left-color border-top-style border-right-style border-bottom-style border-left-style border-top-width border-right-width border-bottom-width border-left-width border-width border cap-height caption-side centerline clear clip color content counter-increment counter-reset cue-after cue-before cue cursor definition-src descent direction display elevation empty-cells float font-size-adjust font-family font-size font-stretch font-style font-variant font-weight font height letter-spacing line-height list-style-image list-style-position list-style-type list-style margin-top margin-right margin-bottom margin-left margin marker-offset marks mathline max-height max-width min-height min-width orphans outline-color outline-style outline-width outline overflow padding-top padding-right padding-bottom padding-left padding page page-break-after page-break-before page-break-inside pause pause-after pause-before pitch pitch-range play-during position quotes richness size slope src speak-header speak-numeral speak-punctuation speak speech-rate stemh stemv stress table-layout text-align text-decoration text-indent text-shadow text-transform unicode-bidi unicode-range units-per-em vertical-align visibility voice-family volume white-space widows width widths word-spacing x-height z-index"), "gm"), css: "keyword" }, { regex: new RegExp(this.GetValuesCSS("above absolute all always aqua armenian attr aural auto avoid baseline behind below bidi-override black blink block blue bold bolder both bottom braille capitalize caption center center-left center-right circle close-quote code collapse compact condensed continuous counter counters crop cross crosshair cursive dashed decimal decimal-leading-zero default digits disc dotted double embed embossed e-resize expanded extra-condensed extra-expanded fantasy far-left far-right fast faster fixed format fuchsia gray green groove handheld hebrew help hidden hide high higher icon inline-table inline inset inside invert italic justify landscape large larger left-side left leftwards level lighter lime line-through list-item local loud lower-alpha lowercase lower-greek lower-latin lower-roman lower low ltr marker maroon medium message-box middle mix move narrower navy ne-resize no-close-quote none no-open-quote no-repeat normal nowrap n-resize nw-resize oblique olive once open-quote outset outside overline pointer portrait pre print projection purple red relative repeat repeat-x repeat-y rgb ridge right right-side rightwards rtl run-in screen scroll semi-condensed semi-expanded separate se-resize show silent silver slower slow small small-caps small-caption smaller soft solid speech spell-out square s-resize static status-bar sub super sw-resize table-caption table-cell table-column table-column-group table-footer-group table-header-group table-row table-row-group teal text-bottom text-top thick thin top transparent tty tv ultra-condensed ultra-expanded underline upper-alpha uppercase upper-latin upper-roman url visible wait white wider w-resize x-fast x-high x-large x-loud x-low x-slow x-small x-soft xx-large xx-small yellow"), "g"), css: "value" }, { regex: new RegExp(this.GetValuesCSS("[mM]onospace [tT]ahoma [vV]erdana [aA]rial [hH]elvetica [sS]ans-serif [sS]erif"), "g"), css: "value" }];
    this.CssClass = "dp-css";
    this.Style = ""
};
dp.sh.Highlighter.prototype.GetKeywordsCSS = function (a) {
    return "\\b([a-z_]|)" + a.replace(/ /g, "(?=:)\\b|\\b([a-z_\\*]|\\*|)") + "(?=:)\\b"
};
dp.sh.Highlighter.prototype.GetValuesCSS = function (a) {
    return "\\b" + a.replace(/ /g, "(?!-)(?!:)\\b|\\b()") + ":\\b"
};
dp.sh.Brushes.CSS.prototype = new dp.sh.Highlighter;
dp.sh.Brushes.CSS.Aliases = ["css"];
dp.sh.Brushes.CSharp = function () {
    this.regexList = [{ regex: dp.sh.RegexLib.SingleLineCComments, css: "comment" }, { regex: dp.sh.RegexLib.MultiLineCComments, css: "comment" }, { regex: dp.sh.RegexLib.DoubleQuotedString, css: "string" }, { regex: dp.sh.RegexLib.SingleQuotedString, css: "string" }, { regex: new RegExp("^\\s*#.*", "gm"), css: "preprocessor" }, { regex: new RegExp(this.GetKeywords("abstract as base bool break byte case catch char checked class const continue decimal default delegate do double else enum event explicit extern false finally fixed float for foreach get goto if implicit in int interface internal is lock long namespace new null object operator out override params private protected public readonly ref return sbyte sealed set short sizeof stackalloc static string struct switch this throw true try typeof uint ulong unchecked unsafe ushort using virtual void while"), "gm"), css: "keyword" }];
    this.CssClass = "dp-c";
    this.Style = ""
};
dp.sh.Brushes.CSharp.prototype = new dp.sh.Highlighter;
dp.sh.Brushes.CSharp.Aliases = ["c#", "c-sharp", "csharp"];
dp.sh.Brushes.Cpp = function () {
    this.regexList = [{ regex: dp.sh.RegexLib.SingleLineCComments, css: "comment" }, { regex: dp.sh.RegexLib.MultiLineCComments, css: "comment" }, { regex: dp.sh.RegexLib.DoubleQuotedString, css: "string" }, { regex: dp.sh.RegexLib.SingleQuotedString, css: "string" }, { regex: new RegExp("^ *#.*", "gm"), css: "preprocessor" }, { regex: new RegExp(this.GetKeywords("ATOM BOOL BOOLEAN BYTE CHAR COLORREF DWORD DWORDLONG DWORD_PTR DWORD32 DWORD64 FLOAT HACCEL HALF_PTR HANDLE HBITMAP HBRUSH HCOLORSPACE HCONV HCONVLIST HCURSOR HDC HDDEDATA HDESK HDROP HDWP HENHMETAFILE HFILE HFONT HGDIOBJ HGLOBAL HHOOK HICON HINSTANCE HKEY HKL HLOCAL HMENU HMETAFILE HMODULE HMONITOR HPALETTE HPEN HRESULT HRGN HRSRC HSZ HWINSTA HWND INT INT_PTR INT32 INT64 LANGID LCID LCTYPE LGRPID LONG LONGLONG LONG_PTR LONG32 LONG64 LPARAM LPBOOL LPBYTE LPCOLORREF LPCSTR LPCTSTR LPCVOID LPCWSTR LPDWORD LPHANDLE LPINT LPLONG LPSTR LPTSTR LPVOID LPWORD LPWSTR LRESULT PBOOL PBOOLEAN PBYTE PCHAR PCSTR PCTSTR PCWSTR PDWORDLONG PDWORD_PTR PDWORD32 PDWORD64 PFLOAT PHALF_PTR PHANDLE PHKEY PINT PINT_PTR PINT32 PINT64 PLCID PLONG PLONGLONG PLONG_PTR PLONG32 PLONG64 POINTER_32 POINTER_64 PSHORT PSIZE_T PSSIZE_T PSTR PTBYTE PTCHAR PTSTR PUCHAR PUHALF_PTR PUINT PUINT_PTR PUINT32 PUINT64 PULONG PULONGLONG PULONG_PTR PULONG32 PULONG64 PUSHORT PVOID PWCHAR PWORD PWSTR SC_HANDLE SC_LOCK SERVICE_STATUS_HANDLE SHORT SIZE_T SSIZE_T TBYTE TCHAR UCHAR UHALF_PTR UINT UINT_PTR UINT32 UINT64 ULONG ULONGLONG ULONG_PTR ULONG32 ULONG64 USHORT USN VOID WCHAR WORD WPARAM WPARAM WPARAM char bool short int __int32 __int64 __int8 __int16 long float double __wchar_t clock_t _complex _dev_t _diskfree_t div_t ldiv_t _exception _EXCEPTION_POINTERS FILE _finddata_t _finddatai64_t _wfinddata_t _wfinddatai64_t __finddata64_t __wfinddata64_t _FPIEEE_RECORD fpos_t _HEAPINFO _HFILE lconv intptr_t jmp_buf mbstate_t _off_t _onexit_t _PNH ptrdiff_t _purecall_handler sig_atomic_t size_t _stat __stat64 _stati64 terminate_function time_t __time64_t _timeb __timeb64 tm uintptr_t _utimbuf va_list wchar_t wctrans_t wctype_t wint_t signed"), "gm"), css: "datatypes" }, { regex: new RegExp(this.GetKeywords("break case catch class const __finally __exception __try const_cast continue private public protected __declspec default delete deprecated dllexport dllimport do dynamic_cast else enum explicit extern if for friend goto inline mutable naked namespace new noinline noreturn nothrow register reinterpret_cast return selectany sizeof static static_cast struct switch template this thread throw true false try typedef typeid typename union using uuid virtual void volatile whcar_t while"), "gm"), css: "keyword" }];
    this.CssClass = "dp-cpp";
    this.Style = ""
};
dp.sh.Brushes.Cpp.prototype = new dp.sh.Highlighter;
dp.sh.Brushes.Cpp.Aliases = ["cpp", "c", "c++"];
dp.sh.Brushes.Objc = function () {
    var datatypes = 'ATOM BOOL BOOLEAN BYTE CHAR COLORREF DWORD DWORDLONG DWORD_PTR ' +
        'DWORD32 DWORD64 FLOAT HACCEL HALF_PTR HANDLE HBITMAP HBRUSH ' +
        'HCOLORSPACE HCONV HCONVLIST HCURSOR HDC HDDEDATA HDESK HDROP HDWP ' +
        'HENHMETAFILE HFILE HFONT HGDIOBJ HGLOBAL HHOOK HICON HINSTANCE HKEY ' +
        'HKL HLOCAL HMENU HMETAFILE HMODULE HMONITOR HPALETTE HPEN HRESULT ' +
        'HRGN HRSRC HSZ HWINSTA HWND INT INT_PTR INT32 INT64 LANGID LCID LCTYPE ' +
        'LGRPID LONG LONGLONG LONG_PTR LONG32 LONG64 LPARAM LPBOOL LPBYTE LPCOLORREF ' +
        'LPCSTR LPCTSTR LPCVOID LPCWSTR LPDWORD LPHANDLE LPINT LPLONG LPSTR LPTSTR ' +
        'LPVOID LPWORD LPWSTR LRESULT PBOOL PBOOLEAN PBYTE PCHAR PCSTR PCTSTR PCWSTR ' +
        'PDWORDLONG PDWORD_PTR PDWORD32 PDWORD64 PFLOAT PHALF_PTR PHANDLE PHKEY PINT ' +
        'PINT_PTR PINT32 PINT64 PLCID PLONG PLONGLONG PLONG_PTR PLONG32 PLONG64 POINTER_32 ' +
        'POINTER_64 PSHORT PSIZE_T PSSIZE_T PSTR PTBYTE PTCHAR PTSTR PUCHAR PUHALF_PTR ' +
        'PUINT PUINT_PTR PUINT32 PUINT64 PULONG PULONGLONG PULONG_PTR PULONG32 PULONG64 ' +
        'PUSHORT PVOID PWCHAR PWORD PWSTR SC_HANDLE SC_LOCK SERVICE_STATUS_HANDLE SHORT ' +
        'SIZE_T SSIZE_T TBYTE TCHAR UCHAR UHALF_PTR UINT UINT_PTR UINT32 UINT64 ULONG ' +
        'ULONGLONG ULONG_PTR ULONG32 ULONG64 USHORT USN VOID WCHAR WORD WPARAM WPARAM WPARAM ' +
        'char bool short int __int32 __int64 __int8 __int16 long float double __wchar_t ' +
        'clock_t _complex _dev_t _diskfree_t div_t ldiv_t _exception _EXCEPTION_POINTERS ' +
        'FILE _finddata_t _finddatai64_t _wfinddata_t _wfinddatai64_t __finddata64_t ' +
        '__wfinddata64_t _FPIEEE_RECORD fpos_t _HEAPINFO _HFILE lconv intptr_t id ' +
        'jmp_buf mbstate_t _off_t _onexit_t _PNH ptrdiff_t _purecall_handler ' +
        'sig_atomic_t size_t _stat __stat64 _stati64 terminate_function ' +
        'time_t __time64_t _timeb __timeb64 tm uintptr_t _utimbuf ' +
        'va_list wchar_t wctrans_t wctype_t wint_t signed';
    var keywords = 'break case catch class copy const __finally __exception __try ' +
        'const_cast continue private public protected __declspec ' +
        'default delete deprecated dllexport dllimport do dynamic_cast ' +
        'else enum explicit extern if for friend getter goto inline ' +
        'mutable naked namespace new nil NO noinline nonatomic noreturn nothrow NULL ' +
        'readonly readwrite register reinterpret_cast retain strong return SEL selectany self ' +
        'setter sizeof static static_cast struct super switch template ' +
        'thread throw true false try typedef typeid typename union ' +
        'using uuid virtual void volatile whcar_t while YES';
    //顺序很重要
    this.regexList = [
        { regex: new RegExp(this.GetKeywords(datatypes), 'gm'), css: 'keyword' },  // primitive data types
        { regex: new RegExp(this.GetKeywords(keywords), 'gm'), css: 'keyword' },  // keywords
        { regex: new RegExp('@\\w+\\b', 'g'), css: 'keyword' },  // @-keywords
        { regex: new RegExp('[: ]nil', 'g'), css: 'keyword' },  // nil-workaround
        { regex: new RegExp('\\.\\w+', 'g'), css: 'xcodeconstants' },  // accessors
        { regex: new RegExp(' \\w+(?=[:\\]])', 'g'), css: 'vars' },  // messages
        { regex: dp.sh.RegexLib.SingleLineCComments, css: 'comment' },  // comments
        { regex: dp.sh.RegexLib.MultiLineCComments, css: 'comment' },  // comments
        { regex: dp.sh.RegexLib.DoubleQuotedString, css: 'string' },  // strings
        { regex: dp.sh.RegexLib.SingleQuotedString, css: 'string' },  // strings
        { regex: new RegExp('@"[^"]*"', 'gm'), css: 'string' },  // strings
        { regex: new RegExp('\\d', 'gm'), css: 'xcodenumber' },  // numeric values
        { regex: new RegExp('^ *#.*', 'gm'), css: 'xcodepreprocessor' },  // preprocessor
        { regex: new RegExp('\\w+(?= \\*)', 'g'), css: 'keyword' } // object types - variable declaration
        /*{ regex: new RegExp('\\b[A-Z]\\w+\\b(?=[ ,;])', 'gm'),  css: 'xcodekeyword' },*/ // object types - protocol
    ];
    this.CssClass = 'dp-objc';
    this.Style = '.dp-objc .vars { color: #d00; }';
}
dp.sh.Brushes.Objc.prototype = new dp.sh.Highlighter;
dp.sh.Brushes.Objc.Aliases = ['objc'];

/*ZeroClipboard/ZeroClipboard.js*/
var ZeroClipboard = {
    version: "1.0.7", clients: {}, moviePath: 'http://static.blog.csdn.net/scripts/ZeroClipboard/ZeroClipboard.swf', nextId: 1, $: function (thingy) {
        if (typeof (thingy) == 'string')
            thingy = document.getElementById(thingy);
        if (true || !thingy.addClass) {
            thingy.hide = function () {
                this.style.display = 'none';
            };
            thingy.show = function () {
                this.style.display = 'block';
            };
            thingy.addClass = function (name) {
                this.removeClass(name);
                this.className += ' ' + name;
            };
            thingy.removeClass = function (name) {
                var classes = this.className.split(/\s+/);
                var idx = -1;
                for (var k = 0; k < classes.length; k++) {
                    if (classes[k] == name) {
                        idx = k;
                        k = classes.length;
                    }
                }
                if (idx > -1) {
                    classes.splice(idx, 1);
                    this.className = classes.join(' ');
                }
                return this;
            };
            thingy.hasClass = function (name) {
                return !!this.className.match(new RegExp("\\s*" + name + "\\s*"));
            };
        }
        return thingy;
    }, setMoviePath: function (path) {
        this.moviePath = path;
    }, dispatch: function (id, eventName, args) {
        var client = this.clients[id];
        if (client) {
            client.receiveEvent(eventName, args);
        }
    }, register: function (id, client) {
        this.clients[id] = client;
    }, getDOMObjectPosition: function (obj, stopObj) {
        var info = { left: 0, top: 0, width: obj.width ? obj.width : obj.offsetWidth, height: obj.height ? obj.height : obj.offsetHeight };
        while (obj && (obj != stopObj)) {
            info.left += obj.offsetLeft;
            info.top += obj.offsetTop;
            obj = obj.offsetParent;
        }
        return info;
    }, Client: function (elem) {
        this.handlers = {};
        this.id = ZeroClipboard.nextId++;
        this.movieId = 'ZeroClipboardMovie_' + this.id;
        ZeroClipboard.register(this.id, this);
        if (elem)
            this.glue(elem);
    }
};
ZeroClipboard.Client.prototype = {
    id: 0, ready: false, movie: null, clipText: '', handCursorEnabled: true, cssEffects: true, handlers: null, glue: function (elem, appendElem, stylesToAdd) {
        this.domElement = ZeroClipboard.$(elem);
        var zIndex = 99;
        if (this.domElement.style.zIndex) {
            zIndex = parseInt(this.domElement.style.zIndex, 10) + 1;
        }
        if (typeof (appendElem) == 'string') {
            appendElem = ZeroClipboard.$(appendElem);
        }
        else if (typeof (appendElem) == 'undefined') {
            appendElem = document.getElementsByTagName('body')[0];
        }
        var box = ZeroClipboard.getDOMObjectPosition(this.domElement, appendElem);
        this.div = document.createElement('div');
        var style = this.div.style;
        style.position = 'absolute';
        style.left = '' + box.left + 'px';
        style.top = '' + box.top + 'px';
        style.width = '' + box.width + 'px';
        style.height = '' + box.height + 'px';
        style.zIndex = zIndex;
        if (typeof (stylesToAdd) == 'object') {
            for (addedStyle in stylesToAdd) {
                style[addedStyle] = stylesToAdd[addedStyle];
            }
        }
        appendElem.appendChild(this.div);
        this.div.innerHTML = this.getHTML(box.width, box.height);
    }, getHTML: function (width, height) {
        var html = '';
        var flashvars = 'id=' + this.id + '&width=' + width + '&height=' + height;
        if (navigator.userAgent.match(/MSIE/)) {
            var protocol = location.href.match(/^https/i) ? 'https://' : 'http://';
            html += '<object classid="clsid:d27cdb6e-ae6d-11cf-96b8-444553540000" codebase="' + protocol + 'download.macromedia.com/pub/shockwave/cabs/flash/swflash.cab#version=9,0,0,0" width="' + width + '" height="' + height + '" id="' + this.movieId + '" align="middle"><param name="allowScriptAccess" value="always" /><param name="allowFullScreen" value="false" /><param name="movie" value="' + ZeroClipboard.moviePath + '" /><param name="loop" value="false" /><param name="menu" value="false" /><param name="quality" value="best" /><param name="bgcolor" value="#ffffff" /><param name="flashvars" value="' + flashvars + '"/><param name="wmode" value="transparent"/></object>';
        }
        else {
            html += '<embed id="' + this.movieId + '" src="' + ZeroClipboard.moviePath + '" loop="false" menu="false" quality="best" bgcolor="#ffffff" width="' + width + '" height="' + height + '" name="' + this.movieId + '" align="middle" allowScriptAccess="always" allowFullScreen="false" type="application/x-shockwave-flash" pluginspage="http://www.macromedia.com/go/getflashplayer" flashvars="' + flashvars + '" wmode="transparent" />';
        }
        return html;
    }, hide: function () {
        if (this.div) {
            this.div.style.left = '-2000px';
        }
    }, show: function () {
        this.reposition();
    }, destroy: function () {
        if (this.domElement && this.div) {
            this.hide();
            this.div.innerHTML = '';
            var body = document.getElementsByTagName('body')[0];
            try {
                body.removeChild(this.div);
            } catch (e) {
                ;
            }
            this.domElement = null;
            this.div = null;
        }
    }, reposition: function (elem) {
        if (elem) {
            this.domElement = ZeroClipboard.$(elem);
            if (!this.domElement)
                this.hide();
        }
        if (this.domElement && this.div) {
            var box = ZeroClipboard.getDOMObjectPosition(this.domElement);
            var style = this.div.style;
            style.left = '' + box.left + 'px';
            style.top = '' + box.top + 'px';
        }
    }, setText: function (newText) {
        this.clipText = newText;
        if (this.ready)
            this.movie.setText(newText);
    }, addEventListener: function (eventName, func) {
        eventName = eventName.toString().toLowerCase().replace(/^on/, '');
        if (!this.handlers[eventName])
            this.handlers[eventName] = [];
        this.handlers[eventName].push(func);
    }, setHandCursor: function (enabled) {
        this.handCursorEnabled = enabled;
        if (this.ready)
            this.movie.setHandCursor(enabled);
    }, setCSSEffects: function (enabled) {
        this.cssEffects = !!enabled;
    }, receiveEvent: function (eventName, args) {
        eventName = eventName.toString().toLowerCase().replace(/^on/, '');
        switch (eventName) {
            case 'load':
                this.movie = document.getElementById(this.movieId);
                if (!this.movie) {
                    var self = this;
                    setTimeout(function () {
                        self.receiveEvent('load', null);
                    }, 1);
                    return;
                }
                if (!this.ready && navigator.userAgent.match(/Firefox/) && navigator.userAgent.match(/Windows/)) {
                    var self = this;
                    setTimeout(function () {
                        self.receiveEvent('load', null);
                    }, 100);
                    this.ready = true;
                    return;
                }
                this.ready = true;
                this.movie.setText(this.clipText);
                this.movie.setHandCursor(this.handCursorEnabled);
                break;
            case 'mouseover':
                if (this.domElement && this.cssEffects) {
                    this.domElement.addClass('hover');
                    if (this.recoverActive)
                        this.domElement.addClass('active');
                }
                break;
            case 'mouseout':
                if (this.domElement && this.cssEffects) {
                    this.recoverActive = false;
                    if (this.domElement.hasClass('active')) {
                        this.domElement.removeClass('active');
                        this.recoverActive = true;
                    }
                    this.domElement.removeClass('hover');
                }
                break;
            case 'mousedown':
                if (this.domElement && this.cssEffects) {
                    this.domElement.addClass('active');
                }
                break;
            case 'mouseup':
                if (this.domElement && this.cssEffects) {
                    this.domElement.removeClass('active');
                    this.recoverActive = false;
                }
                break;
        }
        if (this.handlers[eventName]) {
            for (var idx = 0, len = this.handlers[eventName].length; idx < len; idx++) {
                var func = this.handlers[eventName][idx];
                if (typeof (func) == 'function') {
                    func(this, args);
                }
                else if ((typeof (func) == 'object') && (func.length == 2)) {
                    func[0][func[1]](this, args);
                }
                else if (typeof (func) == 'string') {
                    window[func](this, args);
                }
            }
        }
    }
};


function report(id, t, e) {
    $.createMask();
    var url = blog_address + "/common/report?id=" + id + "&t=" + t;
    if (location.href.indexOf("dev") > -1)
    {
        var u = blog_address.split('/')[3];
        url = "http://dev.blog.csdn.net:5391/"+u+"/common/report?id=" + id + "&t=" + t;
    }
    
    if (t == 3) {
        var floor = e.attr('floor');
        url += "&floor=" + floor;
    }
    var top = (document.documentElement.clientHeight - 400) / 2 + (document.documentElement.scrollTop || document.body.scrollTop);
    var left = (document.documentElement.clientWidth - 400) / 2;
    $("#report_dialog").load(url).css({ "top": top + "px", "left": left }).show();
}


jQuery.createMask = function () {
    var height = document.documentElement.clientHeight;
    var width = document.documentElement.clientWidth;
    var bodyHeight = $("body").height();
    if (bodyHeight > height) {
        height = bodyHeight;
    }
    var mask = {};
    if ($("#mask_div").length == 0) {
        $("body").append('<div id="mask_div" style="position:absolute;top:0;left:0;filter:alpha(opacity=20);-moz-opacity:0.2;opacity:.2;"></div>')
    }
    mask = $("#mask_div");
    mask.css({ "width": width, "height": height, "background": "#000" });
};
jQuery.removeMask = function () {
    $("#mask_div").remove();
};
Array.prototype.contain = function (val) {
    for (var i = 0; i < this.length; i++) {
        if (this[i] === val) {
            return true;
        }
    }
    return false;
}


function deleteArticle(articleId) {
    if (!confirm("你确定要删除这篇文章吗？"))
        return;
    $.get(blog_address + "/article/delete?articleId=" + articleId, function (txt) {
        var data = eval("(" + txt + ")");
        if (data.result) {
            alert("文章已删除！");
            location.reload();
        } else {
            if (data.content)
                alert(data.content);
            else
                alert("无法删除，请到后台删除！");
        }
    });
}
function openWindow(url, title) {
    var _t = title || encodeURI(document.title);
    var _u = url || encodeURIComponent(document.location);
    var f = function () {
        var left = (screen.width - 600) / 2;
        var top = (screen.height - 450) / 2;
        if (!window.open(url, '', 'toolbar=0,resizable=1,scrollbars=yes,status=1,width=600,height=400'))
            location.href = url;
    }
    if (/Firefox/.test(navigator.userAgent))
        setTimeout(f, 0);
    else
        f();
}

function GetCategoryArticles(id, username, type, aid) {
    var topid = "top_" + id;
    
    if (type == "top") {
        var objtop = $("#" + topid + " li");
        if (objtop.length > 0) {
            return;
        }
    }

   

    var url = "/" + username + "/svc/GetCategoryArticleList?id=" + id + "&type=" + type;
    //url="http://dev.blog.csdn.net:5391"+url;
    $.get(url, function (res) {

        if (type == "top") {
            var objtop = $("#" + topid);
            objtop.html("");
            $(res).each(function (i) {
                var obj = res[i];
                if (aid != obj.articleid) {
                    var articleurl = "http://blog.csdn.net/" + username + "/article/details/" + obj.articleid;
                    var aritcleid = "top_aritcle_" + obj.articleid + Math.random().toString().replace("0.");
                    objtop.append("<li class=\"tracking-ad\" data-mod=\"popu_140\"><em>•</em><a href='" + articleurl + "'  id='" + aritcleid + "' target=\"_blank\"></a></li> ");
                    $("#" + aritcleid).text(obj.title);
                    $("#" + aritcleid).attr("title", obj.title);
                }
            });

            var count = $(objtop.parent().parent().find("em")[0]).text().replace("（", "").replace("）", "");
            if (parseInt(count) > 5) {
                var moreurl = objtop.parent().find(".subItem_t a").attr("href");
                objtop.append("<li style=\"padding-left: 300px;\"><a href='" + moreurl + "' target=\"_blank\">更多</a></li>");
            }

        }
        else if (type == "foot") {

            $(".my_article_t_cur").removeClass("my_article_t_cur");   
            $("#samecate"+id).addClass("my_article_t_cur");

            var objfootleft = $(".my_list.fl");
            var objfootright = $(".my_list.fr");

            objfootleft.html("");
            objfootright.html("");

            var j = 0;

            $.each(res, function (i) {
                var obj = res[i];
                if (j < 11) {
                    if (aid != obj.articleid) {
                        var articleurl = "http://blog.csdn.net/" + username + "/article/details/" + obj.articleid;
                        var aritcleid = "foot_aritcle_" + obj.articleid + Math.random().toString().replace("0.");

                        var html = "<li><a href='" + articleurl + "'  id='" + aritcleid + "' target=\"_blank\"></a><label><span>" + obj.posttime + "</span><i class=\"fa fa-eye\"></i><em>" + obj.viewcount + "</em></label></li> ";
                        if (j % 2 == 1) {
                            objfootright.append(html);
                        }
                        else {
                            objfootleft.append(html);
                        }

                        j++;

                        $("#" + aritcleid).text(obj.title);
                        $("#" + aritcleid).attr("title", obj.title);


                        $(".my_article").show();
                    }
                }
            });

            var count = $(".my_article_t_l.my_article_t_cur em").text().replace("（", "").replace("）", "");
            if (parseInt(count) > 10) {                
                var moreurl=$(".my_article_t_l.my_article_t_cur").attr("moreurl");               
                if (moreurl != "") {
                    //objfootright.append("<li style=\"padding-left: 200px;\"><a href='" + moreurl + "' target=\"_blank\">更多</a></li>");
                    $(".my_more").remove();
                    $(".my_article_c_c").append('<a href=' + moreurl + ' class="my_more">更多文章</a>');
                }
            }
        }
    });
}
//but-comment-topicon
$('.but-comment-topicon').on('click',function(){
    var commentListPosition = $('#comment_form').offset().top;
    $(window).scrollTop(commentListPosition);
});
// 行号，兼容markdown与html
$(function(){
    var oBlog_line = {
        markdown_line:function(){
          $(".markdown_views pre").addClass("prettyprint");       
            // prettyPrint();
            $('pre.prettyprint code').each(function () {
                var lines = $(this).text().split('\n').length;
                var $numbering = $('<ul/>').addClass('pre-numbering').hide();
                $(this).addClass('has-numbering').parent().append($numbering);
                for (i = 1; i <= lines; i++) {
                    $numbering.append($('<li/>').text(i));
                };
                $numbering.fadeIn(1700);
            });
            $('.pre-numbering li').css("color","#999");             
            setTimeout(function(){
                $(".math").each(function(index,value){$(this).find("span").last().css("color","#fff"); })
            });
            setTimeout(function () {
                $(".toc a[target='_blank']").attr("target", "");
            }, 500);

        },
        html_line:function(){
          /*文章中的插入代码*/
            $(".article_content pre").each(function () {
                var $this = $(this);
                if ($this.attr("class").indexOf("brush:") != -1) {
                    var lang = $this.attr("class").split(';')[0].split(':')[1];
                    $this.attr('name', 'code');
                    $this.attr('class', lang);
                }
                if ($this.attr("class")) {
                    $this.attr('name', 'code');
                }
            });
            $('.article_content textarea[name=code]').each(function () {
                var $this = $(this);
                if ($this.attr("class").indexOf(":") != -1) {
                    $this.attr("class", $this.attr("class").split(':')[0]);
                }
            });
            dp.SyntaxHighlighter.HighlightAll('code');
            //修复旧文章中的高亮
            $('.highlighter').addClass('dp-highlighter');

            /*如果浏览器不支持访问剪切板，就用flash实现*/
            if (!window.clipboardData) {
                setTimeout(setCopyBtn, 1000);
            }
        /*使用flash实现复制到剪切板*/
        function setCopyBtn() {
            $('.CopyToClipboard').each(function () {

                var clip = new ZeroClipboard.Client();
                clip.setHandCursor(true);
                clip.addEventListener('load', function (client) { });
                clip.addEventListener('mouseOver', function (client) {
                    var _c_ode = client.movie.parentNode.parentNode.parentNode.parentNode.nextSibling.innerHTML;
                    _c_ode = _c_ode.replace(/&lt;/g, "<").replace(/&gt;/g, ">").replace(/&amp;/g, "&");
                    client.setText(_c_ode);
                });
                clip.addEventListener('complete', function (client, text) {
                    alert("代码已经复制到你的剪贴板。");
                });
                clip.glue(this, this.parentNode);
            });
          }
        }
    }
    var $markdown_views = $('.markdown_views')[0];
    if(!$markdown_views){
      oBlog_line.html_line();
    }else{
      oBlog_line.markdown_line();
    }
})
/*
* @file 博客的代码片段保存到code的代码片
* @author liwz
* @create Time 2016-5-19
* */

$(function()
{
    (function(){
        var oBlog_code = {},
            $target,
            i = 0,b = false,
            title = $.trim($(".list_c_t a").text() || $(".link_title a").text()),
            code,lang,username,description,tags,
            newVal ='';
        username = currentUserName;
        //获取页面中默认的tag
        oBlog_code.getTag = function (){
            var pageTags = $('.link_categories a');
            var len = pageTags.length,oPageTag='';
            len = len <= 5 ? len : 5;
            pageTags.each(function(){
                oPageTag += '<span class=\"label blog_tag\"><span>' + $(this).text() +
                        '</span><a title=\"Removing tag\" href=\"javascript:;\">x</a></span>';
                newVal +=  $(this).text() + ',';
            });
            i=len;
            return oPageTag;
        };
        /**
         * 页面开始加载时把收藏代码片的按钮和弹层预先加载到页面中,上报
         * */

        oBlog_code.preLoad = function(){
            var _this = this;
            var oWrap = $(".dp-highlighter,.prettyprint"),
                    oSnippetsBtn = '<div class="save_code tracking-ad" data-mod="popu_249"><a href="javascript:;"><img src="http://static.blog.csdn.net/images/save_snippets.png"/></a></div>',
                    winHref= window.location.href,
                    oTagHtml;
            oWrap.each(function(){
                $(this).append(oSnippetsBtn);
            });
            oTagHtml = _this.getTag();
            $('body').append("<div id='mask_code'></div>");
            var oHtml = '<div class="gist_edit" ><div class="save_snippets clearfix">'+
                    '<div class="tit"><h3>保存代码片</h3><span>整理和分享保存的代码片，请访问<a href="https://code.csdn.net/snippets_manage" target="_blank">代码笔记</a></span></div>'+
                    '<div class="con_form"><ul class="gist_edit_list clearfix"><li><span class="red">*</span><span class="txt">标题</span>'+
                    '<input id="form_title" class="form-input" placeholder="" type="text"></li><li><span class="red">*</span><span class="txt">描述</span><textarea id="form-textarea" class="form-textarea" placeholder=""></textarea></li><li><span class="red">&nbsp;</span><span class="txt">标签</span>'+
                    '<div id="divSearchTags">';
            oHtml += oTagHtml;
            oHtml +='<input id="insertTag" class="insertTag" placeholder="请输入标签，按Enter生成(最多5项)" type="text" value="" name="insertTag"  maxlength="21" style="color: rgb(51, 51, 51);">'+
                    '<input id="OrganTag" class="OrganTag" type="hidden" name="OrganTag" value='+newVal+'>'+
                    '<input id="OldOrganTag" class="OldOrganTag" type="hidden" name="OldOrganTag" value=""><input type="hidden" name="txtSearchTags"></div>'+

                    '</li></ul></div><div class="bottom-bar"><a href="javascript:;" class="btn-submit btn-cancel">取消</a><span class="tracking-ad"  data-mod="popu_250"><a class="btn-submit btn-confirm"  href="javascript:;">确定</a></span></div></div></div>';

            $('body').append(oHtml);
            $("#form_title").attr("placeholder",decodeURIComponent(title));
            $("#form-textarea").attr("placeholder",decodeURIComponent(title) + '： ' + winHref );
        }
        oBlog_code.preLoad();


         //鼠标移上去显示保存到我的代码片按钮
        (function()
        {
            $(document).delegate('.dp-highlighter,.prettyprint','mouseenter',function()
            {
                if($(this).height() < 80){
                    //$(this).addClass("pad_bot");
                    $(this).find(".save_code img").attr("src",'http://static.blog.csdn.net/images/save_snippets_01.png');
                }
                $(this).find('.save_code').show();
                return false;
            });
            $(document).delegate('.dp-highlighter,.prettyprint','mouseleave',function()
            {
                $(this).find('.save_code').hide();
                //$(this).removeClass("pad_bot");
                return false;
            });
        })();
        oBlog_code.getLang = function (str,htmlEdit){
            var _this = this;
            if(str != undefined){
                str = str.split(' ')[0];
                str = str.substring(9,str.length);
                //lang = preLan;
                _this.handle(str);
            }else{
                htmlEdit = encodeURIComponent(htmlEdit.substring(1,htmlEdit.length-1));
                _this.handle(htmlEdit);
            }
        };
        //处理字符串成code可识别的语言类型
        oBlog_code.handle = function (str){
            var langArr = {
                'cpp':'c++',
                'csharp':'c#',
                'obj-c':'objective-c',
                'objc':'objective-c'
            };
            for(var key in langArr){
                if(str.toLowerCase() == key ){
                    if(str != langArr[str]){
                        str = langArr[str];
                    }
                }
            }
            lang = encodeURIComponent(str);
        };
        //点击保存按钮弹出弹层
        oBlog_code.showPop = function()
        {
            var _this = this;

            $(document).delegate('.save_code','click',function(ev)
            {
                var event = ev || event;
                title = $.trim($(".list_c_t a").text() || $(".link_title a").text());
                $target = $(event.target || event.srcElement);
                code = $(this).parents('.prettyprint').find('code').text() || $(this).parents('.dp-highlighter').next('pre').text() || $(this).parents('.dp-highlighter').siblings('textarea[name="code"]').text();
                code = encodeURIComponent(code);
                if(code == ''){
                    alert("无法获取到code代码！");
                    return;
                }
                var preLan = $(this).parents('.prettyprint').find('code').attr('class');
                var htmlEdit = $(this).parents('.dp-highlighter').find('.bar').find('strong').text() || $(this).parents('.dp-highlighter').find('.bar').find('b').text();
                //获取语言类型
                _this.getLang(preLan,htmlEdit);

                if(currentUserName)
                {
                    $('#mask_code').show();
                    $('.gist_edit').show();

                    var offsetTop = $(window).scrollTop() + $(window).height()/2;
                    $(".gist_edit").css({
                        "top":offsetTop
                    });
                }
                else
                {
                    window.location.href = "https://passport.csdn.net/";
                }
            })
        };
        oBlog_code.showPop();

        oBlog_code.hidePop = function(obj)
        {
            var _this = this;
            $(document).delegate('.btn-cancel','click',function()
            {
                $("#form-textarea").val('');
                $("#form_title").val('');
                _this.resetTag(obj);
            });

            //点击确定保存到code的代码段
            $(document).delegate('.btn-confirm','click',function()
            {
                description = encodeURIComponent($("#form-textarea").val() || $("#form-textarea").attr('placeholder'));
                title = encodeURIComponent($("#form_title").val() || title);
                tags = encodeURIComponent($("#OrganTag").val());

                $.ajax({
                    type: "post",
                    /*dataType:"jsonp",
                    jsonp:'callback',*/
                    url: "http://blog.csdn.net/"+currentUserName+"/svc/addpostcode",
                    data: 'code='+code+'&lang='+lang + '&username='+username + '&title='+title + '&description='+description + '&tags='+tags,
                    success: function(msg){
                        if(msg)
                        {
                            alert("保存成功！");
                            _this.resetTag(obj);
                        }
                    },
                    error:function()
                    {
                        alert("保存失败");
                    }
                });
            });
        };
        oBlog_code.hidePop('#divSearchTags');

        //把弹层隐藏
        oBlog_code.resetTag = function ()
        {
            $('#mask_code').hide();
            $('.gist_edit').hide();
        };

        oBlog_code.enterTag = function (Wrap,insert){
            var _this = this;
            var event = arguments.callee.caller.arguments[0] || window.event; //消除浏览器差异
            if (event.keyCode == 13 || event.keyCode == 188) {
                _this.addTag(Wrap,insert);
            }
        };
        oBlog_code.addTag = function (Wrap,insert)
        {
            var insertval = $.trim($(insert).val());
            if (insertval != "")
            {
                //排重
                $(''+Wrap+'>span>span').each(
                        function ()
                        {
                            var spanval =$.trim($(this).html());
                            if (insertval == spanval)
                            {
                                b = true;
                                //值相同，把flg的值设为true，再调用test()方法时弹出“此项已被选择”
                                $(insert).off('keydown');   //先把keydown事件关闭了，再按空格的时候让test()方法生效
                                if(i < 5)    //只有在i<5的时候才让“此项已选择弹出”,如果i>5时也输入一个和前面重复的标签的话，这句话就不再弹出
                                {
                                    test(true);
                                }
                                return;
                            }
                        }
                );
                //如果标签个数大于5，就把b置为true,即不让再添加
                if(i>=5)
                {
                    b = true;
                    test(false);
                }
                if(b==false)    //只有b=false的时候才添加标签
                {

                    $(insert).before('<span class=\"label blog_tag\"><span>' + insertval +
                            '</span><a title=\"Removing tag\" href=\"javascript:;\">x</a></span>');
                    newVal +=  insertval + ',';
                    i++;
                }
                if(b == true)
                {
                    b = false;      //b=true有两种情况，（1）是和前面标签重复了  （2）已经大于5项，如果是第一种情况，把b置，那后面就要再把b置为false，让后面的还可以再添加
                    $(insert).val('');
                }
            }
            $(insert).val('');   //添加成为标签了，再把输入框中的值给清空
            $('#OrganTag').val(newVal);

            //输入框的时候，如果有相同的就延时300ms弹出警告
            function test(flg)
            {
                if(flg)
                {
                    setTimeout(function(){
                        alert("此项已被选择");
                    },300);
                }
                else{
                    setTimeout(function(){
                        alert("已选择5项");
                    },300);
                }
            }
        };
        /**
         * 删除标签
         * */
        oBlog_code.DelTag = function(obj) {
            $(obj).parent("span").remove();
            var delVal = $(obj).siblings("span").text();
            newVal = newVal.substring(0,newVal.length-1);
            if(newVal.indexOf(delVal) != -1){
                var newValArr = newVal.split(',');
                for(var j =0;j<newValArr.length; j++){
                    if(delVal == newValArr[j]){
                        newValArr.splice(j,1);
                    }
                }
                newVal = newValArr.join(',') + ',';
                $('#OrganTag').val(newVal);
            }
            i--;
        };

        oBlog_code.Tag = function(Wrap,insert){
            var _this = this;
            $(document).delegate($(insert),'keydown',function()
            {
                _this.enterTag(Wrap,insert);
            });
            $(document).delegate(''+Wrap+' span a','click',function()
            {
                _this.DelTag(this);
            });
        };
        oBlog_code.Tag('#divSearchTags','#insertTag');

    }).call(this);
});
$(function () { // 新背景详情页 博主信息区码云逻辑判断
  $.get('http://blog.csdn.net/'+username+'/article/GetMyYunStatus',function(data){
    var code = data.code; //code 1为开通码云，2为没有开通
    var openUserGitee = '<a href="'+data.url+'?utm_source=csdn_blog" target="_blank">'+data.projects_count+'</a>';
    var giteeIndex = 'https://gitee.com?utm_source=csdn_blog';
    var giteeUserSignUp = ' https://gitee.com/signup?utm_source=csdn_blog'
    if(username === currentUserName){ // 博主本身打开详情页
      if(code===1){ // 博主开通码云
        $('aside .inf_number_box dl:nth-child(4) dd').html(openUserGitee);
      }else{
        $('aside .inf_number_box dl:nth-child(4) dd').html('<a class="thinFont" href='+giteeUserSignUp+' target="_blank">去开通</a>');
      }
    }else{
      if(code===1){
        $('aside .inf_number_box dl:nth-child(4) dd').html(openUserGitee);
      }else{
        $('aside .inf_number_box dl:nth-child(4) dd').html('<a class="thinFont" href='+giteeIndex+' target="_blank">未开通</a>');
      }
    }
    
  })
});


$(function () {
    var all_nav_hid = false;
    $('.all_nav').on('mouseenter',function(){
        $('.nav_com').css({'height':'90px','border':'1px solid #E3E3E3','box-shadow':'0 2px 2px rgba(110,110,110,.1)'});
        all_nav_hid = true;
    });
    $('.nav_com').on('mouseleave',function(){
        $('.nav_com').css({'height':'45px','border':'1px solid rgba(0,0,0,0)','box-shadow':'0 0 0 rgba(110,110,110,0)'})
        all_nav_hid = false;
    });
    var navTxt = $('.nav_com').find('li.active a').text();
    if($('.nav_com li').hasClass('active second_bar')){
        $('.all_nav span.txt').text(navTxt);
        $('.all_nav').css({'border-bottom':'4px solid #ca0c16'});
        $('.nav_com').find('li.active').hide();
    }else{
        $('.all_nav span.txt').text('全部');
        $('.all_nav').css({'border-bottom':'4px solid rgba(0,0,0,0)'});
    }

});
$(function(){
  function DOMScroll(num){
      $('html,body').animate({scrollTop: num}, 500);
  }
  function reachComment(){ // 到达评论
    var $CommentBox = $('.comment_box'),
        CommentBoxTop = $CommentBox.offset().top
        DOMScroll(CommentBoxTop);
  }
  $(document).on('click','.btn-pinglun',reachComment);
  $(document).on('click','.returnTop',function(){$('html,body').animate({scrollTop: 0}, 500);});
  
  // focus event
  function contentFocus(){
    $(this).animate({height:84},500);
    $('.comment_area').animate({height:180},500);
    $('.bot_bar').animate({opacity:1},500);
  }
  // blur event
  function contentBlur(e){
    var $obj = $(e.srcElement || e.target),
        val = $('.comment_content').val();
    if (!($($obj).is(".comment_box,.comment_box *"))){
      if(val == ''){
        $('.comment_content').animate({height:40},500);
        $('.comment_area').animate({height:60},500);
        $('.bot_bar').animate({opacity:0},500);
      }
    }
  }
  $(document).on('focus','.comment_content',contentFocus)
  $(document).on('click',contentBlur)
  
  // 右侧推荐栏切换
  function newColumnTabSwitch(e){
    var $obj = $(e.srcElement || e.target),
        $recommend_btn = $('.recommend_btn'),
        $new_btn = $('.new_btn')
        $recommendbox = $('.recommend'),
        $newbox = $('.new_column .new');
    if ($($obj).is(".recommend_btn,.recommend_btn *")){
      $newbox.fadeOut(300,function(){
        $recommendbox.fadeIn(300);
      });
      $recommend_btn.css('border-bottom', '2px solid #f00');
      $new_btn.css('border-bottom', '2px solid #e4ebf4');
    }else if($($obj).is(".new_btn,.new_btn *")){
      $recommendbox.fadeOut(300,function(){
        $newbox.fadeIn(300);
      });
      $new_btn.css('border-bottom', '2px solid #f00');
      $recommend_btn.css('border-bottom', '2px solid #e4ebf4');
    }
  }
  $(document).on('click','.new_column',newColumnTabSwitch)
  

  overflow_hide($('.column-list .title'),44,'overflow-hide-host-height')
  overflow_hide($('.hotArticle-list a'),40 ,'overflow-hide-hotArticle-height')
  overflow_hide($('.imgAndText .right-text .title a'),48 ,'overflow-hide-right-text-height')
  overflow_hide($('.text .content a'),48 ,'overflow-hide-text-height')
  
})

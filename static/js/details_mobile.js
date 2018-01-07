/*
*created by liwz 2015/11/5
*
*/

$(function(){
    //搜索展开隐藏功能
    // var searchFn = (function()
    // {
    //       var oSearch = $("#search_J"),
    //             oSearchC = $("#search_c_J"),
    //             oClose = $(".icon_close");

    //       oSearch.on("click",function()
    //       {
    //           oSearchC.slideToggle();
    //       });
    //       oClose.on("click",function()
    //       {
    //         oSearchC.slideUp();
    //       });
    // })();
// var readMore = (function(){
//     var widHeight = $(window).height();
//     var artHeight = $('.article_c').height();
//     if(artHeight>(widHeight*2)){
//         $('.article_c').height(widHeight*2-285).css({'overflow':'hidden'});
//         var article_show = true;
//         $('.read_more_btn').on('click',bindRead_more);
//     }else{
//         article_show = true;
//         $('.article_c').removeClass('article_Hide');
//         $('.readall_box').hide().addClass('readall_box_nobg');
//     }
//     function bindRead_more(){
//         if(!article_show){
//             $('.article_c').height(widHeight*2).css({'overflow':'hidden'});
//             $('.readall_box').show().removeClass('readall_box_nobg');
//             article_show = true;
//         }else{
//             $('.article_c').height("").css({'overflow':'hidden'});
//             $('.readall_box').show().addClass('readall_box_nobg');
//             $('.readall_box').hide().addClass('readall_box_nobg');
//             article_show = false;
//         }
//     }
// })()
    //点击按钮左侧菜单弹出及隐藏
    // var leftNavFn = (function()
    // {
    //     var oMenuBtn = $("#menu_J"),
    //           oLeftNav = $(".leftNav"),
    //           oMask = $("#mask"),
    //           winH = $(window).height();
    //     oMenuBtn.on('click',function()
    //     {
    //         if(!oLeftNav.is(":animated"))
    //         {
    //             oLeftNav.animate({'left':0},500);
    //             //var oMask = $('<div id="mask"></div>');
    //             oMask.show();
    //             $('body').css({
    //                 'height':winH,
    //                 'overflow-y':'hidden'
    //                 });         //屏蔽滚动条
    //       }
    //         return false;        //阻止默认事件及冒泡
    //     });
    //     oMask.click(function(ev)
    //     {
    //           /*if($(ev.target).attr('id') == 'mask')
    //           {*/
    //               if(!oLeftNav.is(":animated"))
    //               {
    //                   oLeftNav.animate({'left':-40 + "rem"},500);
    //                   oMask.hide();
    //                   $('body').css({
    //                       'height':'auto',
    //                       'overflow-y':'auto'
    //                       });      //打开滚动条
    //               }
    //           //}
    //           return false;
    //     })
    // })();

    //点击关闭底部广告
    // var home_AD = (function()
    // {
    //     var oClose = $(".ad_close");
    //     oClose.on('click',function()
    //     {
    //         $(this).parents(".ad_box").hide();
    //     })
    // })();

    //点击全部博文展示筛选内容
    // var blogList = (function()
    // {
    //     var oTotalBlog = $(".total_blog");
    //     oTotalBlog.on('click',function()
    //     {
    //         $(this).next(".filter_category").toggle();
    //     })
    // })();

    //返回顶部
    var backTop = (function()
    {
            var $backToTopFile = $('.backToTop ')
                    .click(function()
                    {
                        $("body,html").animate({scrollTop:0},500);
                    });

            var backToFun = function ()
            {
                var scrollTop = $(document).scrollTop();
                var winh = $(window).height();
                (scrollTop > 0) ? $backToTopFile.show() : $backToTopFile.hide();
                //IE6下的定位
                if (!window.XMLHttpRequest) {
                    $backToTopFile.css("top", scrollTop + winh - 86);
                }
            };
            $(window).bind('scroll',backToFun);
            backToFun();
    })();

    //左侧菜单导航增加点击时的样式
    // var leftStyle = (function()
    // {
    //     var aLeftA = $(".nav_list a");
    //     aLeftA.on('click',function()
    //     {
    //         $(this).addClass("left_cur");
    //     })
    // })();

    // var preCode = (function()
    // {
        //$("pre").css({'width':'100%!important','overflow':'hidden'});
    //     var aPre = document.getElementsByTagName('pre');
    //     for(var i=0; i<aPre.length; i++)
    //     {
    //         aPre[i].style.width = '100%';
    //         aPre[i].style.overflow = "hidden";
    //     }
    // })();
    // function isWeixinBrowser(){
    //   if( typeof WeixinJSBridge !== "undefined" ) {
    //     return true;
    //   }
    // }
// (function(){
  // $('#menu_J').on('click',function(){
  //   if (document.referrer === '' && isWeixinBrowser()) {
  //     // 没有来源页面信息的时候，改成首页URL地址
  //     window.location="//home.blog.csdn.net";
  //   }else{
  //     console.log('ok');
  //      history.pushState({page: 1}, "title 1", "/");
  //     // window.location=document.referrer;
  //     setTimeout('history.back()',3000)
  //     // history.back();
  //     // history.go(1)
  //   }
  // 
  // })
  // 
  // $(function(){  
  //   // pushHistory();  
  //   window.addEventListener("popstate", function(e) {  
  //     // alert("我监听到了浏览器的返回按钮事件啦");//根据自己的需求实现自己的功能  
  //     console.log(document.referrer !== window.location.href);
  //     if (document.referrer !== window.location.href) {
  //       window.location="//home.blog.csdn.net";
  //     }
  // }, false);  
  //   function pushHistory() {  
  //       var state = {  
  //           title: "title",  
  //           url: "/"  
  //       };  
  //       window.history.pushState(state, "title", "http://m.blog.csdn.net");  
  //   }  
  // 
  // });  
// })()
});
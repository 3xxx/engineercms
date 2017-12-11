/*
*created by liwz 2015/11/5
*
*/

$(function(){
    //搜索展开隐藏功能
    var searchFn = (function()
    {
          var oSearch = $("#search_J"),
                oSearchC = $("#search_c_J"),
                oClose = $(".icon_close");

          oSearch.on("click",function()
          {
              oSearchC.slideToggle();
          });
          oClose.on("click",function()
          {
            oSearchC.slideUp();
          });
    })();

    //点击按钮左侧菜单弹出及隐藏
    var leftNavFn = (function()
    {
        var oMenuBtn = $("#menu_J"),
              oLeftNav = $(".leftNav"),
              oMask = $("#mask"),
              winH = $(window).height();
        oMenuBtn.on('click',function()
        {
            if(!oLeftNav.is(":animated"))
            {
                oLeftNav.animate({'left':0},500);
                //var oMask = $('<div id="mask"></div>');
                oMask.show();
                $('body').css({
                    'height':winH,
                    'overflow-y':'hidden'
                    });         //屏蔽滚动条
          }
            return false;        //阻止默认事件及冒泡
        });
        oMask.click(function(ev)
        {
              /*if($(ev.target).attr('id') == 'mask')
              {*/
                  if(!oLeftNav.is(":animated"))
                  {
                      oLeftNav.animate({'left':-40 + "rem"},500);
                      oMask.hide();
                      $('body').css({
                          'height':'auto',
                          'overflow-y':'auto'
                          });      //打开滚动条
                  }
              //}
              return false;
        })
    })();

    //点击关闭底部广告
    var home_AD = (function()
    {
        var oClose = $(".ad_close");
        oClose.on('click',function()
        {
            $(this).parents(".ad_box").hide();
        })
    })();

    //点击全部博文展示筛选内容
    var blogList = (function()
    {
        var oTotalBlog = $(".total_blog");
        oTotalBlog.on('click',function()
        {
            $(this).next(".filter_category").toggle();
        })
    })();

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
    var leftStyle = (function()
    {
        var aLeftA = $(".nav_list a");
        aLeftA.on('click',function()
        {
            $(this).addClass("left_cur");
        })
    })();

    var preCode = (function()
    {
        //$("pre").css({'width':'100%!important','overflow':'hidden'});
        var aPre = document.getElementsByTagName('pre');
        for(var i=0; i<aPre.length; i++)
        {
            aPre[i].style.width = '100%';
            aPre[i].style.overflow = "hidden";
        }
    })();

});
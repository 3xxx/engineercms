<html lang="zh-CN">

<head>
  <meta name="viewport" content="width=device-width,minimum-scale=1.0,maximum-scale=1.0,user-scalable=no">
  <meta http-equiv="X-UA-Compatible" content="IE=Edge,chrome=1">
  <meta name="renderer" content="webkit">
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <title>工程视频</title>
  <meta name="keywords" content="工程视频">
  <meta name="description" content="简约大气的工程视频">
  <meta name="author" content="order by">
  <link rel="shortcut icon" href="/static/photo/favicon.ico" type="image/x-icon" />
  <link rel="stylesheet" type="text/css" href="/static/photo/bootstrap.min.css">
  <!-- <link rel="stylesheet" type="text/css" href="/static/photo/font-awesome.min.css"> -->
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
  <link rel="stylesheet" type="text/css" href="/static/photo/shadowbox.css">
  <link rel="stylesheet" type="text/css" href="/static/photo/style.css">
  <script src="/static/photo/jquery-2.1.0.min.js"></script>
  <script src="/static/photo/bootstrap.min.js"></script>
  <script src="/static/photo/pandaTab.js"></script>
  <script src="/static/photo/shadowbox.js"></script>
  <script src="/static/photo/jquery.qrcode.min.js"></script>
  <script src="/static/photo/masonry.pkgd.min.js"></script>
  <script src="/static/photo/pandaPicshow.js"></script>
  <style>
    .hidden {
      visibility:hidden;
    }
  </style>
</head>

<body class="y2017 m07 d08 h16 home page pageid-62 page-author-panda page-template page-template-page-all-php bg-gray">
  <div class="backgroundImg" style="background-image:url(/static/photo/1-194601-0-4B0.jpeg)"></div>
  <div id="header">
    <div class="categoryTab-full hidden-xs">
      <div class="container">
        <div class="row">
          <div class="col-sm-10 col-sm-push-1 categoryTab-content no-padding">
            <div class="categoryTab" pandatab="" tabheight="45px" transitiontime="500" rebound="true" navigatorwidth="30px" style="position: relative;">
              <div class="menu-menu-1-container">
                <ul id="menu-menu-1" class="menu" style="margin: 0px; padding: 0px; display: block; white-space: nowrap; word-spacing: -0.43em; overflow: hidden;">
                  <li class="anchor" style="height: 45px; padding: 0px; margin-right: 0px; margin-left: 0px; position: absolute; top: 0px; list-style: none; transition-timing-function: cubic-bezier(0.3, 1.8, 0.4, 1); left: 0px; transition-duration: 500ms; width: 0px;">
                    <div></div>
                  </li>
                  <li style="display: inline-block; height: 45px; margin: 0px; position: relative; z-index: 1; list-style: none; vertical-align: top; word-spacing: normal;"><a href="/" title="网站首页">首页</a></li>
                  <li class="" style="display: inline-block; height: 45px; margin: 0px; position: relative; z-index: 1; list-style: none; vertical-align: top; word-spacing: normal;"><a href="/v1/wx/photo"> 相册 </a></li>
                </ul>
              </div>
            </div>
          </div>
          <div class="col-sm-1 col-sm-push-1 search"><i class="fa fa-search" aria-hidden="true"></i>
            <form method="post" name="searchform" class="sidebar-searchform" id="searchform" action="/plus/search.php">
              <div class="main sidebar-search">
                <input class="searchInput" type="text" id="s" name="q" value="" placeholder="输入关键字">
                <input class="searchIcon" type="submit" value="" id="searchsubmit">
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
    <div class="mobile visible-xs-block">
      <!-- header/mobile -->
      <div class="container">
        <div class="col-xs-2 menu no-padding">
          <div class="menu-icon">
            <div class="bread bread1"></div>
            <div class="bread bread2"></div>
            <div class="bread bread3"></div>
          </div>
          <div class="menu-list">
            <div class="list">
              <div class="menu-menu-1-container">
                <ul id="menu-menu-2" class="menu">
                  <li><a href="/" title="网站首页">首页</a></li>
                  <li class=""><a href="/v1/wx/photo"> 相册 </a></li>
                </ul>
              </div>
            </div>
          </div>
        </div>
        <div class="col-xs-10 mobileSearch"><i class="fa fa-search" aria-hidden="true"></i>
          <form method="post" name="searchform" class="sidebar-searchform" id="searchform" action="/plus/search.php">
            <div class="main sidebar-search">
              <input class="searchInput" type="text" id="s" name="q" value="" placeholder="输入关键字">
              <input class="searchIcon" type="submit" value="" id="searchsubmit">
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
  <!-- #header -->
  <div id="wrapper" class="hfeed">
    <div class="topTitle">
      <h1>工程视频</h1>
      <p class="description">珍贵记忆 - 为工程视频而生</p>
    </div>
    <div class="category container">
      {{range .MonthVideoList}}
      <h2>{{.YearMonth}}</h2>
      <div class="row">
          {{range .VideoLists}}
        <div class="col-xs-6 col-sm-4 col-lg-3 waterfallGrid">
          <a href="/v1/wx/videodetail/{{.Id}}" target="_blank">
            <div class="post-photo album">
              <div class="img">
                <div style="background-image: url(&quot;{{.CoverUrl}}&quot;); height: 249px;" ratio="1" realratio="1"></div>
              </div>
              <h3 class="title">{{.Content}}</h3>
            </div>
          </a>
        </div>
          {{end}}
      </div>
      {{end}}
    </div>
    <div id="loadMore" class="hidden">
      加载数据……
    </div>
  </div>
  <div id="footer">
    <div class="col-lg-12">
      <br>
      <hr />
      <div class="about-link-wrap">
        <a href="/v1/wx/legalnotices" class="ali-about-link" target="_blank">法律声明</a>
        <a href="/v1/wx/cookiespolicy" class="ali-about-link" target="_blank">Cookies政策</a>
        <a href="/v1/wx/integrity" class="ali-about-link" target="_blank">廉正举报</a>
        <a href="/v1/wx/security" class="ali-about-link" target="_blank">安全举报</a>
        <a href="/v1/wx/contact" class="ali-about-link" target="_blank">联系我们</a>
        <a href="https://github.com/3xxx/engineercms" class="ali-about-link" target="_blank"><i class="fa fa-github">源码仓库</i></a>
      </div>
      <div class="about-link-wrap">
        <a target="_blank" href="https://beian.miit.gov.cn/"><img src="/static/img/beiantubiaobianhao.png" class="ali-report-img" alt="粤ICP备2024338992号" loading="lazy"><span class="ali-report-link-text">备案号：粤ICP备2024338992号</span></a>
      </div>
    </div>
  </div>
  <!-- <div id="footer" class="container"><span class="text">Copyright © 2016~2024 工程视频 <a href="https://zsj.itdos.net/" target="_blank">水务设计</a> <a href="https://beian.miit.gov.cn/" target="_blank" rel="nofollow">粤ICP备XXXXXXXX号</a> <a href="/sitemap.xml" target="_blank">XML地图</a> <a href="https://zsj.itdos.net/" target="_blank">工程视频</a></span></div> -->
  <script type="text/javascript" src="/static/photo/theme.js"></script>
  <script type="text/javascript" src="/static/photo/wp-embed.min.js?ver=4.8"></script>
  <script type="text/javascript">
  var i = 2;
  var flag = false;//防止多次调用下拉触发事件
  $(function() {
    var loadData = function(ii) {
      $.getJSON("/v1/wx/getmonthvideodata?page=" + ii, function(data) {
        // console.log(data)
        
        $.each(data, function(i, videolist) {
          console.log(videolist.VideoLists)
          if (videolist.VideoLists.length == 0) {
            $("#loadMore").removeClass('hidden').text('已加载全部数据！');
          }
          // array = tl["start"].split("-");
          // console.log(photolist)
          // console.log(photolist.YearMonth)
          // var _timeline_head_p_small_i_ = $("<i class='glyphicon glyphicon-time'></i>");

          $(".category").append("<h2>"+videolist.YearMonth+"</h2>")
          var _timeline_invert_=""

          $.each(videolist.VideoLists, function(i, videolists) {
            _timeline_invert_ = _timeline_invert_+'<div class="col-xs-6 col-sm-4 col-lg-3 waterfallGrid"><a href="/v1/wx/videodetail/'+videolists.Id+'" target="_blank"><div class="post-photo album"><div class="img"><div style="background-image: url(&quot;'+videolists.CoverUrl+'&quot;); height: 249px;" ratio="1" realratio="1"></div></div><h3 class="title">'+videolists.Content+'</h3> </div></a></div>'
          });

          $(".category").append('<div class="row">'+_timeline_invert_+'</div>')

          // if ($(window).height() >= document.documentElement.scrollHeight) {
          //没有出现滚动条,继续加载下一页
          // loadData(i);
          // i = i + 1;
          // }
        });
      })
    }

    var tcScroll = function() {
      $(window).on('scroll', function() {
        // var scrollTop = document.documentElement.scrollTop
        var scrollTop = $(this).scrollTop();
        var clientHeight = document.documentElement.clientHeight
        var bodyHeight = document.body.clientHeight
        // console.log(scrollTop,clientHeight,bodyHeight)//1706 2205 499

        if (clientHeight-scrollTop-bodyHeight <= 100&& flag === false) {
        // if (scrollTop + windowHeight == scrollHeight) {
            //此处是滚动条到底部时候触发的事件，在这里写要加载的数据，或者是拉动滚动条的操作
          console.log("触底了!!!!");
          //此处是滚动条到底部时候触发的事件，在这里写要加载的数据，或者是拉动滚动条的操作
          loadData(i);
          i = i + 1;
          setTimeout(()=>{
            flag = false;
          },500);
        }
      })
    }
    tcScroll();
  });

  //页面滚到底部异步加载下一页数据
  $(window).scroll(function() {
    //已经滚动到上面的页面高度
    var scrollTop = parseFloat($(this).scrollTop()),
      //页面高度
      scrollHeight = $(document).height(),
      //浏览器窗口高度
      windowHeight = parseFloat($(this).height()),
      totalHeight = scrollTop + windowHeight;
    //此处是滚动条到底部时候触发的事件，在这里写要加载的数据，或者是拉动滚动条的操作
    // if (scrollTop + windowHeight >= scrollHeight - 0.7) {
    if (scrollTop >= scrollHeight) {
      console.log("加载数据……");
      $.ajax({
        type: 'get',
        url: '/v1/wx/getmonthphotodata?page='+i,
        beforeSend: function(XMLHttpRequest) {
          console.log("beforeSend……");
          $("#loadMore").removeClass('hidden').text('正在加载数据...');
        },
        success: function(data) {
          if (data.length == 0) {
            $("#loadMore").removeClass('hidden').text('已加载全部数据！');
          }
          for (var i = 0, length = data.length; i < length; i++) {
            $("#infoList").append("<li><a href='public_content.html?id=" + data[i].Id + "&f=" + first + "&s=" + second + "'><span>" + data[i].Title + "</span><span class='time'>" + data[i].PublishTime + "</span></a></li>");
          }
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
          $("#loadMore").removeClass('hidden').text('数据加载失败，请重试！');
        }
      });

      i++
    }
  });
  </script>
</body>

</html>
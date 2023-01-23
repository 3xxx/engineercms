<!DOCTYPE html>
<html lang="zh-CN" >
<head>
<title>当日照片</title>
<meta name="keywords" content="" />
<meta name="description" content="" />
<meta http-equiv="X-UA-Compatible" content="IE=Edge,chrome=1" >
<meta name="viewport" content="width=device-width,minimum-scale=1.0,maximum-scale=1.0,user-scalable=no"/>
<meta name="renderer" content="webkit">
<meta http-equiv="Content-Type"content="text/html; charset=utf-8"/>
<link rel="shortcut icon" href="/static/photo/favicon.ico" type="image/x-icon" />
<link rel="stylesheet" type="text/css" href="/static/photo/bootstrap.min.css">
<link rel="stylesheet" type="text/css" href="/static/photo/font-awesome.min.css">
<link rel="stylesheet" type="text/css" href="/static/photo/shadowbox.css">
<link rel="stylesheet" type="text/css" href="/static/photo/style.css" />
<script src="/static/photo/jquery-2.1.0.min.js"></script>
<script src="/static/photo/bootstrap.min.js"></script>
<script src="/static/photo/pandaTab.js"></script>
<script src="/static/photo/shadowbox.js"></script>
<script src="/static/photo/jquery.qrcode.min.js"></script>
<script src="/static/photo/masonry.pkgd.min.js"></script>
<script src="/static/photo/pandaPicshow.js"></script>
</head>
<body class="y2017 m07 d09 h01 single postid-55 s-y2017 s-m06 s-d29 s-h22 s-author-panda bg-gray">
<div class="backgroundImg" style="background-image:url(/static/photo/1-194601-0-4B0.jpeg)"></div>

<div id="header">
  <div class="categoryTab-full hidden-xs">
    <div class="container">
      <div class="row">
        <div class="col-sm-10 col-sm-push-1 categoryTab-content no-padding">
          <div class="categoryTab" pandaTab tabHeight='45px' transitionTime='500' rebound='true' navigatorWidth="30px">
            <div class="menu-menu-1-container">
              <ul id="menu-menu-1" class="menu">
                <li ><a href="/"  title="网站首页">首页</a></li>
                <li class=""><a href="/article/"> 文章插图 </a></li>
                <li class=""><a href="/diary/"> 日志插图 </a></li>
              </ul>
            </div>
          </div>
        </div>
        <div class="col-sm-1 col-sm-push-1 search"><i class="fa fa-search" aria-hidden="true"></i>
          <form method="post" name="searchform" class="sidebar-searchform" id="searchform" action="/plus/search.php">
            <div class="main sidebar-search">
              <input class="searchInput" type="text" id="s" name="q" value="" placeholder="输入关键字">
              <input class="searchIcon" type="submit" value="&#xf002" id="searchsubmit">
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
  <div class="mobile visible-xs-block"><!-- header/mobile -->
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
                <li ><a href="/" title="网站首页">首页</a></li>
                <li class=""><a href="/article/"> 文章插图 </a></li>
                <li class=""><a href="/diary/"> 日志插图 </a></li>
              </ul>
            </div>
          </div>
        </div>
      </div>
      <div class="col-xs-10 mobileSearch"><i class="fa fa-search" aria-hidden="true"></i>
        <form method="post" name="searchform" class="sidebar-searchform" id="searchform" action="/plus/search.php">
          <div class="main sidebar-search">
            <input class="searchInput" type="text" id="s" name="q" value="" placeholder="输入关键字">
            <input class="searchIcon" type="submit" value="&#xf002" id="searchsubmit">
          </div>
        </form>
      </div>
    </div>
  </div>
</div>
<div id="wrapper" class="hfeed">
  <div class="topTitle">
    <h1>工程图片</h1>
  </div>
  <div class="category container">
    <div class="row">
      <div class="col-lg-12">
        <div class="post-content">
          <article>
            <div pandaPicshow='enable'>
              <div class="showBox cover" ratio="1.7778"></div>
              <div class="thumbnails">
                {{range .PhotoList}}
                <a fullsize="{{.Url}}" thumbnail="{{.ThumbnailUrl}}" description="{{.YearMonthDay}}"></a>
                {{end}}
              </div>
            </div>
            <div class="col-xs-12 model sxy">
              <ul class="row">
                <div class="post-previous twofifth"> 上一篇： <br><a href='/v1/wx/photodetail?keywords={{.prev}}' rel='prev'>{{.prev}}</a>  </div>
                <div class="post-next twofifth"> 下一篇： <br><a href='/v1/wx/photodetail?keywords={{.next}}' rel='next'>{{.next}}</a>  </div>
              </ul>
            </div>
          </article>
        </div>
        <!-- <div class="xiangguan">
          <div class="">
            <div class="row">
              <div class="col-xs-6 col-sm-4 col-lg-3 waterfallGrid"><a href="/a/luoli/16.html" target="_blank">
                <div class="post-photo album">
                  <div class="img">
                    <div style="background-image:url(/static/photo/1-194601-0-4B0.jpeg);" ratio="1" realratio="1"></div>
                  </div>
                  <h3 class="title">photo拍</h3>
                </div>
                </a></div>
                <div class="col-xs-6 col-sm-4 col-lg-3 waterfallGrid"><a href="/a/luoli/15.html" target="_blank">
                <div class="post-photo album">
                  <div class="img">
                    <div style="background-image:url(/static/photo/1-192G9-0-1218.jpeg);" ratio="1" realratio="1"></div>
                  </div>
                  <h3 class="title">美拍</h3>
                </div>
                </a></div>
                <div class="col-xs-6 col-sm-4 col-lg-3 waterfallGrid"><a href="/a/luoli/14.html" target="_blank">
                <div class="post-photo album">
                  <div class="img">
                    <div style="background-image:url(/static/photo/1-192521-0-4T4.jpeg);" ratio="1" realratio="1"></div>
                  </div>
                  <h3 class="title">套图</h3>
                </div>
                </a></div>
            </div>
          </div>
        </div> -->
      </div>
    </div>
  </div>
  <div class="mainPageNav"></div>
</div>
<div id="footer" class="container">
  <span class="text">Copyright &copy; 2022-2023 工程相册
    <a href=https://zsj.itdos.net/ target='_blank'>水务设计</a>
    <a href="https://beian.miit.gov.cn/" target="_blank" rel="nofollow">粤ICP备XXXXXXXX号</a>
    <a href="/sitemap.xml" target="_blank">XML地图</a>
    <a href="https://zsj.itdos.net/" target="_blank">网站模板</a>
  </span>
</div>
<script type="text/javascript" src="/static/photo/theme.js"></script> 
<script type='text/javascript' src='/static/photo/wp-embed.min.js?ver=4.8'></script>
</body>
</html>
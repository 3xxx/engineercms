<!-- 不可用article替换！！！ -->
<!DOCTYPE html>
<title>{{.product.Title}} - 珠三角设代</title>
<html style="font-size: 25.875px;">

<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=Edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no">
  <meta name="apple-mobile-web-app-status-bar-style" content="yes">
  <!-- <meta name="shenma-site-verification" content="bf32a58a3df1c3452b75fbe96c7feb42_1497604697"> -->
  <!-- <script async="" src="/static//gio.js"></script><script src="/static//hm.js"></script> -->
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css" />
  <script type="text/javascript" src="/static/js/details_mobile.js"></script>
  <!--link( rel="stylesheet" href="http://c.csdnimg.cn/public/common/toolbar/css/index.css" )-->
  <!-- <link rel="stylesheet" href="/static//avatar.css"> -->
  <link rel="stylesheet" href="/static/css/common.css">
  <!-- [if IE 7]-->
  <!--link( rel="stylesheet" href="assets/css/font-awesome-ie7.min.css" )-->
  <!-- [endif]-->
  <link rel="stylesheet" href="/static/css/main.css">
  <!-- [if lt IE 9]-->
  <script src="/static/js/html5shiv.min.js"></script>
  <!-- [endif]-->
  <link rel="Stylesheet" type="text/css" href="/static/css/csdn_style.css">
  <link rel="stylesheet" href="/static/css/csdn_blog_detail.min.css">
  <link rel="stylesheet" href="/static/froala/css/froala_style.css">
  <style type="text/css">
  img {
    max-width: 100%
  }

  video {
    width: 100%;
    height: 100%;
    object-fit: fill
  }

  /*.blog_main .article_next_prev span {
        width: 80px;
        margin: 0px;
        padding: 0px;
        padding-left: 30px;
      }*/
  </style>
  <!-- <style type="text/css">
  img{max-width:100%}
  video{width:100%; height:100%; object-fit: fill}
  #main .article_next_prev span {
    width: 80px;
    margin: 0px;
    padding: 0px;
    padding-left: 30px;
  }
</style> -->
</head>

<body>
  <div class="blog_main">
    <div class="blog_top_wrap">
      <div class="blog_top">
        <a href=""><i id="menu_J" class="iconfont icon_l"></i>
        </a>
        <form method="get" action="http://m.blog.csdn.net/search/index" id="searchform"><i id="search_J" class="iconfont icon_r"></i>
          <div id="search_c_J" class="search">
            <input type="text" placeholder="请输入" id="search" name="keyword" page="1" value="">
            <i class="iconfont icon_search"></i>
            <i class="iconfont icon_close"></i>
          </div>
        </form>
        <script type="text/javascript" src="/static/js/search.js"></script>
        <a href="">
          <h2 class="blog_top_t">珠三角水资源配置</h2>
        </a>
      </div>
    </div>
    <div class="main_list">
      <div class="blog_article">
        <div class="blog_article_t">
          <div class="article_l">
            <img src="/static/img/go.jpg" alt="img">
            <a href="" class="article_user">{{.product.Principal}}</a>
          </div>
          <div class="article_r"></div>
        </div>
        <div class="blog_article_c clearfix">
          <h3 class="article_t">{{.product.Title}}</h3>
          <p class="date"> <i>发表于</i> <em>{{dateformat .product.Created "2006-01-02 15:04:05"}} &nbsp;</em> <i>{{.article.Views}}</i>
            <span>人阅读</span>
          </p>
          <p class="category_p">
            标签：
          </p>
          <div class="article_c">
            <div id="article_content" class="article_content csdn-tracking-statistics" data-mod="popu_307" data-dsm="post"></div>
            <div class="fr-view">
              {{str2html .article.Content}}
            </div>
            <ul class="article_next_prev">
              {{if .Pre}}
              <li class="prev_article">
                <span>上一篇：</span>
                <a href="/project/product/article/{{.PreArticleId}}" onclick="">{{.PreArticleTitle}}</a>
              </li>
              {{end}}
              {{if .Next}}
              <li class="next_article">
                <span>下一篇：</span>
                <a href="/project/product/article/{{.NextArticleId}}" onclick="">{{.NextArticleTitle}}</a>
              </li>
              {{end}}
            </ul>
            <div style="clear:both; height:10px;"></div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <script>
  $(function() {
    $('#share_btn').click(function() {
      $('.popup_cover').stop().show();
      $('.sharePopup_box').stop().slideDown();
    });
    $('.sharePopup_cancel').click(function() {
      $('.popup_cover').stop().hide();
      $('.sharePopup_box').stop().slideUp();
    });
  })
  </script>
  <div class="backToTop" style="z-index: 9999999; display: none;">
    <img src="/static/img/iconfont-fudongxiangshang.png" alt="img"></div>
  <div class="blog_footer">©2016-2018, EngineerCMS, All Rights Reserved</div>
  <script type="text/javascript">
  $(function() { //这个和$(document).ready(function()等价
    value = { {.product.Label } };
    array = value.split(",");
    var labelarray = new Array();
    for (i = 0; i < array.length; i++) {
      var th1 = "&nbsp;<span><a href='/project/product/article/keysearch?keyword=" + array[i] + "'>" + array[i] + "</a></span>"; //label label-info
      $("p.category_p").append(th1); //after(th1);
    }
    // return labelarray.join(",");
    // $("input#cid").remove();
  })
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
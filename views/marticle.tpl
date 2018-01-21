
<!DOCTYPE html>
  <title>{{.product.Title}} - 珠三角设代</title>
    <html style="font-size: 25.875px;"><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
      <meta http-equiv="X-UA-Compatible" content="IE=Edge">
      <meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no">
      <meta name="apple-mobile-web-app-status-bar-style" content="yes">    
      <!-- <meta name="shenma-site-verification" content="bf32a58a3df1c3452b75fbe96c7feb42_1497604697"> -->
      <!-- <script async="" src="/static//gio.js"></script><script src="/static//hm.js"></script> -->
      <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
      <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
      <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
      <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
      <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.min.js"> </script>
      <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>
      <script src="/static/ueditor/ueditor.parse.min.js"></script>
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
      
  <!-- <script src="/static//dm.js"></script> -->
  <!-- GrowingIO Analytics code version 2.1 -->
  <!-- Copyright 2015-2017 GrowingIO, Inc. More info available at http://www.growingio.com -->
  <!-- End GrowingIO Analytics code version: 2.1 -->
<style type="text/css">
  img{max-width:100%}
  video{width:100%; height:100%; object-fit: fill}
  /*.blog_main .article_next_prev span {
        width: 80px;
        margin: 0px;
        padding: 0px;
        padding-left: 30px;
      }*/
</style>
</head>
<body>
<!-- <link href="/static//jiathis_share.css" rel="stylesheet" type="text/css"> -->
<!-- <iframe frameborder="0" style="position: absolute; display: none; opacity: 0;" src="/static//saved_resource.html"></iframe> -->
<!-- <div class="jiathis_style" style="position: absolute; z-index: 1000000000; display: none; top: 50%; left: 50%; overflow: auto;"></div><div class="jiathis_style" style="position: absolute; z-index: 1000000000; display: none; overflow: auto;"></div> -->
<!-- <iframe frameborder="0" src="/static//jiathis_utility.html" style="display: none;"></iframe> -->
<!-- <div id="MathJax_Message" style="display: none;"></div> -->

<!-- <script language="JavaScript" type="text/javascript" src="/static//jquery.cookie.js"></script> -->
<!-- <script type="text/javascript" src="/static//main.js"></script> -->

<!-- <script type="text/javascript" src="/static//MathJax.js"></script> -->
<!-- <link rel="stylesheet" href="/static//markdown_views.css"> -->

<!-- <script type="text/javascript" src="/static//digg.js"></script> -->

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
    <!-- <div id="_oee02xdda9"></div>
  -->
  <div class="blog_article">
    <div class="blog_article_t">
      <div class="article_l">
        <!-- <a href="http://m.blog.csdn.net/hotqin888">  
        -->
        <img src="/static/img/go.jpg" alt="img">  
        <!-- </a>  
        -->
        <a href="" class="article_user">{{.product.Principal}}</a>
      </div>
      <div class="article_r"></div>
    </div>
    <div class="blog_article_c clearfix">
      <h3 class="article_t">{{.product.Title}}</h3>
      <p class="date"> <i>发表于</i> <em>{{dateformat .product.Created "2006-01-02 15:04:05"}}  &nbsp;</em>  <i>{{.article.Views}}</i>
        <span>人阅读</span>
      </p>
      <p class="category_p">
        标签：
        <!-- 分类：
              <span>beego</span>
      <span>golang</span>
      <span>MeritMS</span>
      <span>校审</span>
      -->
      </p>

    <div class="article_c">
      <div id="article_content" class="article_content csdn-tracking-statistics" data-mod="popu_307" data-dsm="post">{{str2html .article.Content}}</div>

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

    <!-- <div class="readall_box readall_box_nobg" style="display: none;">  
    <div class="read_more_mask"></div>
    <div class="read_more_btn">阅读全文</div>
  </div>
  -->
  </div>
  <!-- <div class="blog_handle">  
  -->
  <!-- <span id="digg" style="cursor:pointer">  
  <i class="iconfont"></i> <em>0</em>
  </span>
  -->
  <!-- <span id="bury" style="cursor:pointer">  
  <i class="iconfont"></i>
  <em>0</em>
  </span>
  -->
  <!-- </div>  
  -->
  <!-- <div class="prev_next"></div>
  -->
  </div>

  <!-- <div id="_su0y0s4emtc"></div>
  -->
  <!-- <div style="display: none" id="tags">设计</div>
  -->
  <!-- <script type="text/javascript" src="/static//main-1.0.0.js"></script>
  -->
  <!-- <script type="text/javascript" src="/static//highcode.js"></script>
  -->
  <!-- <script type="text/javascript" src="/static//comment.js"></script>
  -->
  </div>

  <!-- <div class="leftNav">  
  <ul class="nav_list">
  <li>
  <a href="http://m.blog.csdn.net/home/index">
  <i>•</i>
  <span>首页</span>
  <img src="/static/img/iconfont-youjiantou.png" alt="img" class="arrow_r"></a>
  </li>
  <li>
  <a href="http://m.blog.csdn.net/Column/Column?Channel=mobile&amp;Type=hot">
  <i>•</i>
  <span>移动开发</span>
  <img src="/static/img/iconfont-youjiantou.png" alt="img" class="arrow_r"></a>
  </li>
  <li>
  <a href="http://m.blog.csdn.net/Column/Column?Channel=enterprise&amp;Type=hot">
  <i>•</i>
  <span>架构</span>
  <img src="/static/img/iconfont-youjiantou.png" alt="img" class="arrow_r"></a>
  </li>
  <li>
  <a href="http://m.blog.csdn.net/Column/Column?Channel=cloud&amp;Type=hot">
  <i>•</i>
  <span>云计算/大数据</span>
  <img src="/static/img/iconfont-youjiantou.png" alt="img" class="arrow_r"></a>
  </li>
  <li>
  <a href="http://m.blog.csdn.net/Column/Column?Channel=www&amp;Type=hot">
  <i>•</i>
  <span>互联网</span>
  <img src="/static/img/iconfont-youjiantou.png" alt="img" class="arrow_r"></a>
  </li>
  <li>
  <a href="http://m.blog.csdn.net/Column/Column?Channel=system&amp;Type=hot">
  <i>•</i>
  <span>运维</span>
  <img src="/static/img/iconfont-youjiantou.png" alt="img" class="arrow_r"></a>
  </li>
  <li>
  <a href="http://m.blog.csdn.net/Column/Column?Channel=database&amp;Type=hot">
  <i>•</i>
  <span>数据库</span>
  <img src="/static/img/iconfont-youjiantou.png" alt="img" class="arrow_r"></a>
  </li>
  <li>
  <a href="http://m.blog.csdn.net/Column/Column?Channel=web&amp;Type=hot">
  <i>•</i>
  <span>前端</span>
  <img src="/static/img/iconfont-youjiantou.png" alt="img" class="arrow_r"></a>
  </li>
  <li>
  <a href="http://m.blog.csdn.net/Column/Column?Channel=code&amp;Type=hot">
  <i>•</i>
  <span>编程语言</span>
  <img src="/static/img/iconfont-youjiantou.png" alt="img" class="arrow_r"></a>
  </li>
  <li>
  <a href="http://m.blog.csdn.net/Column/Column?Channel=software&amp;Type=hot">
  <i>•</i>
  <span>研发管理</span>
  <img src="/static/img/iconfont-youjiantou.png" alt="img" class="arrow_r"></a>
  </li>
  <li>
  <a href="http://m.blog.csdn.net/Column/Column?Channel=other&amp;Type=hot">
  <i>•</i>
  <span>综合</span>
  <img src="/static/img/iconfont-youjiantou.png" alt="img" class="arrow_r"></a>
  </li>
  </ul>
  </div>
  -->
  <!-- <div id="mask"></div>
  -->
  </div>

    <!-- <div class="backToTop" style="display: none;"><img src="/static//iconfont-fudongxiangshang.png" alt="img"></div> -->
    <!--分享弹窗：-->
    <!-- <div class="popup_cover"></div> -->
    <!-- <div class="sharePopup_box"> -->
      <!-- <div class="sharePopup"> -->
        <!--JiaThis Button BEGIN-->    
        <!-- <div class="jiathis_style_32x32">
          <a class="jiathis_button_weixin" title="分享到微信">
            <span class="jiathis_txt jiathis_separator jtico jtico_weixin"> <b><i class="iconfont"></i></b>  <em>微信分享</em>
            </span>
          </a>
          <a class="jiathis_button_tsina" title="分享到微博">
            <span class="jiathis_txt jiathis_separator jtico jtico_tsina"> <b><i class="iconfont"></i></b>  <em>新浪微博</em>
            </span>
          </a>
          <a class="jiathis_button_cqq" title="分享到QQ好友">
            <span class="jiathis_txt jiathis_separator jtico jtico_cqq">
              <b>
                <i class="iconfont"></i>
              </b>
              <em>QQ好友</em>
            </span>
          </a>
          <a class="jiathis_button_qzone" title="分享到QQ空间">
            <span class="jiathis_txt jiathis_separator jtico jtico_qzone">
              <b>
                <i class="iconfont"></i>
              </b>
              <em>QQ空间</em>
            </span>
          </a>
        </div> -->
        <!-- <script type="text/javascript" src="/static//jia.js" charset="utf-8"></script> -->
        <!-- <script type="text/javascript" src="/static//plugin.client.js" charset="utf-8"></script> -->
        <!--JiaThis Button END-->
      <!-- </div> -->
      <!-- <div class="sharePopup_cancel">取 消</div> -->
    <!-- </div> -->
    <script>
      $(function(){
        $('#share_btn').click(function(){
          $('.popup_cover').stop().show();
          $('.sharePopup_box').stop().slideDown();
        });
        $('.sharePopup_cancel').click(function(){
          $('.popup_cover').stop().hide();
          $('.sharePopup_box').stop().slideUp();
        });
      })
    </script>

   <div class="backToTop" style="z-index: 9999999; display: none;">
     <img src="/static/img/iconfont-fudongxiangshang.png" alt="img"></div>
   <div class="blog_footer">©2016-2018, EngineerCMS, All Rights Reserved</div>
<!--     <script type="text/javascript" src="/static//fontSize.js"></script>
<script type="text/javascript" charset="utf-8" src="/static//tracking-1.0.3.js"></script> -->

<!-- <div style="position:fixed; top:0; left:0; overflow:hidden;"><input style="position:absolute; left:-300px;" type="text" value="" id="focus_retriever" readonly="true"></div> -->
<script type="text/javascript">
  $(function(){//这个和$(document).ready(function()等价
    value={{.product.Label}};
    array=value.split(",");
    var labelarray = new Array(); 
    for (i=0;i<array.length;i++){
      // labelarray[i]="<a href='/project/search/"+array[i]+"'>" + array[i] + "</a>";
      // var th1="&nbsp;分类：<a href='/project/product/article/keysearch?keyword="+ array[i]+"'><span class='btn btn-info btn-xs'>" +array[i] + "</span></a>";//label label-info
     var th1="&nbsp;<span><a href='/project/product/article/keysearch?keyword="+ array[i]+"'>" +array[i] + "</a></span>";//label label-info
      $("p.category_p").append(th1);//after(th1);
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


    $(function(){
        // var content =$('#content').val();
        //判断ueditor 编辑器是否创建成功
        ue.addListener("ready", function () {
        // editor准备好之后才可以使用
        ue.setContent({{str2html .article.Content}});
        });
    });

</script>

</body>
</html>

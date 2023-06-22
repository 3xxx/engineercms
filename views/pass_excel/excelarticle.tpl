<!DOCTYPE html>
<!-- https://froala.com/wysiwyg-editor/docs/overview/ -->
<!-- Step 4. Display Content -->
<!-- To preserve the look of the edited HTML outside of the rich text editor you have to include the following CSS files. -->

<!-- CSS rules for styling the element inside the editor such as p, h1, h2, etc. -->
<!-- <link href="../css/froala_style.min.css" rel="stylesheet" type="text/css" />
Also, you should make sure that you put the edited content inside an element that has the class fr-view. -->

<head>
  <!-- 下面这行导致移动端字体不变小 -->
  <meta name="HandheldFriendly" content="True">
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
  <link rel="stylesheet" href="/static/wordpress/header.css" type="text/css">
  <link rel="stylesheet" href="/static/wordpress/style(1).css" type="text/css" media="screen">
  <link rel="stylesheet" type="text/css" href="/static/wordpress/style.css" media="screen">
  <link rel="stylesheet" href="/static/wordpress/noticons.css">
  <link rel="stylesheet" href="/static/wordpress/blog.css" type="text/css" media="screen">
  <link rel="stylesheet" href="/static/wordpress/blog-sidebar.css" type="text/css" media="screen">

  <link rel="stylesheet" href="/static/froala/css/froala_style.css">
  <!-- <link rel="stylesheet" href="/static/froala/css/plugins/image.css"> -->

  <meta name="description" content="With Recurring Payments, your supporters become your sustainers, and a reliable income stream frees you to push your creative boundaries.">
  <title>{{.article.Title}}</title>
  <!-- <link rel="alternate" type="application/rss+xml" title="The WordPress.com Blog » A New Way to Earn Money on WordPress.com Comments Feed" href="https://en.blog.wordpress.com/2019/11/12/recurring-payments/feed/"> -->
  <script src="/static/wordpress/quant.js" async="" type="text/javascript"></script>
  <script src="/static/wordpress/rules-p-3Ma3jHaQMB_bS.js" async=""></script>
  <script src="/static/wordpress/wp-emoji-release.min.js" type="text/javascript" defer=""></script>
  <script src="/static/wordpress/wp-emoji-release.min.js" type="text/javascript" defer=""></script>
  <!-- <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" /> -->
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <style type="text/css">
  img.wp-smiley,
  img.emoji {
    display: inline !important;
    border: none !important;
    box-shadow: none !important;
    height: 1em !important;
    width: 1em !important;
    margin: 0 .07em !important;
    vertical-align: -0.1em !important;
    background: none !important;
    padding: 0 !important;
  }

  .floating-button {
    display: block;
    width: 58px;
    height: 58px;
    border-radius: 50%;
    background: #31b0d5;
    color: #fff;
    margin: 0 auto;
    text-align: center;
    float: right;
    /* background-color: #fff; */
    position: fixed;
    /* top: 20px; */
    bottom: 80px;
    right: 20px;
    /* border: 0 solid #fff; */
    /* border-radius: 500px; */
    box-shadow: 4px 1px 1px #ccc;
    opacity: 0.6;

    display: -webkit-box;
    display: -ms-flexbox;
    display: -webkit-flex;
    display: flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    -webkit-align-items: center;
    align-items: center;
    -webkit-box-pack: center;
    -ms-flex-pack: center;
    -webkit-justify-content: center;
    justify-content: center;
    /*color: #fff;
        position: relative;
        right: 16px;
        bottom: 88px;
        width: 56px;
        height: 56px;
        border-radius: 50%;
        z-index: 1500;
        overflow: hidden;
        -webkit-transition-duration: 300ms;
        transition-duration: 300ms;
        box-shadow: 0 10px 20px rgba(0, 0, 0, .19), 0 6px 6px rgba(0, 0, 0, .23);
        display: -webkit-box;
        display: -ms-flexbox;
        display: -webkit-flex;
        display: flex;
        -webkit-box-align: center;
        -ms-flex-align: center;
        -webkit-align-items: center;
        align-items: center;
        -webkit-box-pack: center;
        -ms-flex-pack: center;
        -webkit-justify-content: center;
        justify-content: center;
        background-color: #2196f3;*/
  }

  #editor:link,
  #editor:visited,
  h2 a:visited {
    color: #fff;
  }

  .floating-button2 {
    display: block;
    width: 58px;
    height: 58px;
    border-radius: 50%;
    background: #f44336;
    color: #fff;
    margin: 0 auto;
    text-align: center;
    float: right;
    /* background-color: #fff; */
    position: fixed;
    /* top: 20px; */
    bottom: 20px;
    right: 20px;
    /* border: 0 solid #fff; */
    /* border-radius: 500px; */
    box-shadow: 4px 1px 1px #ccc;
    opacity: 0.6;

    display: -webkit-box;
    display: -ms-flexbox;
    display: -webkit-flex;
    display: flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    -webkit-align-items: center;
    align-items: center;
    -webkit-box-pack: center;
    -ms-flex-pack: center;
    -webkit-justify-content: center;
    justify-content: center;
  }

  #deletearticle:hover,
  input[type="submit"]:hover,
  .button:hover,
  .blog-subscribe input[type=submit]:hover {
    background: #f44336;
    border-color: #a8bece;
    color: #2e4453;
  }

  /*.demo {
    width: 400px;
    height: 300px;
    margin: 50px auto;
  }*/

  img {
    -webkit-filter: drop-shadow(10px 10px 10px rgba(0, 0, 0, .5));
    /*考虑浏览器兼容性：兼容 Chrome, Safari, Opera */
    filter: drop-shadow(10px 10px 10px rgba(0, 0, 0, .5));
  }

  .next_article {
    list-style: none;
  }

  .prev_article {
    list-style: none;
  }

  #post-41687 .article_next_prev span {
    width: 51px;
    height: 26px;
    line-height: 26px;
    /*display: inline-block;*/
    color: #fff;
    padding-left: 27px;
    margin-right: 7px;
    background: #999;
  }

  #post-41687 .article_next_prev a {
    text-decoration: none;
    color: #333;
  }

  #post-41687 .article_next_prev .prev_article {
    margin-bottom: 3px;
  }

  #post-41687 .article_next_prev .prev_article span {
    background-image: url("/static/img/skin-type-icon.png");
    background-repeat: no-repeat;
    background-position: 8px -111px;
    margin-right: 7px;
  }

  #post-41687 .article_next_prev .next_article span {
    background-image: url("/static/img/skin-type-icon.png");
    background-repeat: no-repeat;
    background-position: 8px -140px;
    margin-right: 7px;
  }

  #post-41687 .article_next_prev li:hover span {
    background-color: #3d84b0;
  }

  #post-41687 .article_next_prev li:hover a {
    color: #3d84b0;
  }

  #post-41687 .article_next_prev li:hover span {
    background-color: #c88326;
  }

  #post-41687 .article_next_prev li:hover a {
    color: #c88326;
  }

  #post-41687 .article_next_prev #btnDigg {
    background: #f90;
  }

  #post-41687 .article_next_prev #btnBury {
    background: #ff7900;
  }
  </style>
  <style id="wp-block-library-inline-css">
  .has-text-align-justify {
    text-align: justify;
  }
  </style>
  <style id="jetpack-global-styles-frontend-style-inline-css">
  :root {
    --font-headings: unset;
    --font-base: unset;
    --font-headings-default: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen-Sans, Ubuntu, Cantarell, "Helvetica Neue", sans-serif;
    --font-base-default: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen-Sans, Ubuntu, Cantarell, "Helvetica Neue", sans-serif;
  }
  </style>
  <link rel="stylesheet" type="text/css" id="gravatar-card-css" href="/static/wordpress/hovercard.min.css">
  <link rel="stylesheet" type="text/css" id="gravatar-card-services-css" href="/static/wordpress/services.min.css">
  <!-- <script type="text/javascript" src="/static/wordpress/PEJHFPIHPJC2PD3IMTCWTT"></script> -->
  <!-- <script async="true" type="text/javascript" src="/static/wordpress/WV6A5O5PBJBIBDYGZHVBM5"></script> -->
</head>

<body class="wpcomblog" style="padding-top: 40px; position: relative;">
  <header class="toolbar wpcom-masterbar wpcom-header" id="yeeId_3">
    <div class="wpcom-navigation site-navigation wpcom-nav" role="navigation">
      <h1 class="wpcom-title"><a class="wpcom-logo" href="https://zsj.itdos.com/"><span>EngineerCMS.com</span></a></h1>
      <nav class="wpcom-pages">
        <a class="menu-toggle">Menu</a>
        <ul class="pages-menu">
          <li class="menu-features">
            <a href="https://zsj.itdos.com/index" title="Features">首页</a>
          </li>
          <li class="menu-themes">
            <a href="https://zsj.itdos.com/project" title="WordPress Themes for Blogs at WordPress.com">项目</a>
          </li>
          <li class="menu-plans">
            <a href="https://zsj.itdos.com/projectgant" title="Plans">进度</a>
          </li>
          <li class="menu-news">
            <a href="https://zsj.itdos.com/article" title="Blog">文章</a>
          </li>
          <li class="menu-support">
            <a href="https://zsj.itdos.com/standard" title="Support">规范</a>
          </li>
          <li class="menu-signup">
            <a href="https://zsj.itdos.com/register" title="Sign Up">注册</a>
          </li>
          <li class="menu-login">
            <a href="https://zsj.itdos.com/login" title="Log In" class="login-link">登录</a>
          </li>
        </ul>
      </nav>
    </div>
  </header>
  <a id="editor" {{if ne true .RoleUpdate}} style="display:none" {{end}} href="/v1/excel/editorexcelarticle/{{.article.ID}}" class="floating-button">修改</a>
  <button {{if ne true .RoleUpdate}} style="display:none" {{end}} type="button" class="floating-button2" id="deletearticle">删除</button>
  <div id="wrapper" class="en-wrapper">
    <div id="content" class="widecolumn single">
      <style type="text/css">
      #post-focus.has-header-img {
        background-image: url("/static/wordpress/anthony-delanoix-hzgs56ze49s-unsplash.jpg?w=2000");
      }
      </style>
      <div id="blog-header" class="single has-header-img">
        <div class="inner">
          <h2 id="blog-title"><a href="https://zsj.itdos.com/">The EngineerCMS Blog</a></h2>
          <p id="blog-tagline"><a href="https://zsj.itdos.com/">The EngineerCMS Blog</a></p>
        </div>
      </div>
      <div id="post-focus" class="has-header-img post-41687 single">
        <div id="post-focus-alignment">
        </div>
      </div>
      <div class="post" id="post-41687">
        <div class="post-title-top">
          <h2 class="post-title">{{.article.Title}}</h2>
          <div class="entry">
            <p>{{.article.Subtext}}</p>
          </div>
        </div>
        <div class="post-meta-top">
          <div id="wpcom-comments-date">
            <span class="date" id="datetime"><svg class="icon icon-clock" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" viewBox="0 0 24 24" style="enable-background:new 0 0 24 24;" xml:space="preserve">
                <g id="time">
                  <path d="M12,4c4.411,0,8,3.589,8,8s-3.589,8-8,8s-8-3.589-8-8S7.589,4,12,4 M12,2C6.477,2,2,6.477,2,12s4.477,10,10,10 s10-4.477,10-10S17.523,2,12,2L12,2z M15.8,15.4L13,11.667V7h-2v5.333l3.2,4.266L15.8,15.4z"></path>
                </g>
                <g id="Layer_1"></g>
              </svg></span>
          </div>
          <img alt="" src="{{.avatar}}" class="avatar avatar-64 grav-hashed grav-hijack" height="64" width="64" id="grav-c88872d732506446e8b44b677a6a20e0-0">
          <p class="post-author-name"></p>
        </div>
        <div class="entrytext">

          <div class="fr-view">
            {{str2html .article.Content}}
          </div>

        </div>
        <div id="entry-comments">
          <h3 id="comments" class="commentsheader">26 Comments</h3>
          <p class="nocomments"><span>Comments are closed.</span></p>
          <div style="clear: both; margin-top: 15px;"></div>
        </div>
      </div>
    </div>
  </div>
</body>
<script type="text/javascript">

  // $(document).ready(function () {
// $("#datetime").append("我的第一个jQuery代码!");
// });
  $(document).ready(function() { 
   var time2=moment({{.article.CreatedAt}}, 'YYYY-MM-DD h:mm:ss a').add(8, 'hours').format('YYYY-MM-DD h:mm:ss a');
    $("#datetime").append(time2);
  })

// 删除文章
$("#deletearticle").click(function() {
  if ({{.RoleDelete }} == "true") {
    if (confirm("确定删除吗？一旦删除将无法恢复！")) {
      $.ajax({
        type: "post",
        url: "/project/product/deletearticle",
        data: { pid: {{.article.ID }} },
        success: function(data, status) {
          alert("删除“" + data + "”成功！(status:" + status + ".)");
          //关闭标签
          window.close();
        }
      });
    }
  } else {
    alert("权限不够！" + {{.Uid }});
    return;
  }
})
</script>

</html>
<!-- CSS中URL路径

项目中，为a标签添加背景，老是没效果 （VS2013中相关文件的位置：CSS文件位于/Content中，图片位于/images中）

为a标签添加背景的CSS代码为：background:  url(images/sort.png) no-repeat 0 3px;

一直没有效果，后来才发现是绝对路径，相对路径搞错了。
url(images/sort.png)表示的是当前文件夹下的images文件夹下的sort.png图片，因为CSS文件位于/Content文件夹，所以系统就去/Content/images下寻找sort.png图片，而此时图片位于/images中，当然就找不到了。

修复方法：

1. 绝对路径: url(/images/sort.png)  表示到根目录下寻找images文件夹里面的sort.png图片

2. 相对路径：url(../images/sort.png)  表示从当前目录返回到上一层目录，即/目录，然后再查找/目录下的images目录里面的sort.png图片

以下是百度的内容：

"." 代表当前所在目录，相对路径。如:<a href="./abc">文本</a>或<img src="./abc" />；
".." 代表上一层目录，相对路径。如:<a href="../abc">文本</a>或<img src="../abc"/>；
"../../" 代表的是上一层目录的上一层目录，相对路径。 如:<img src="../../abc" />；
"/" 代表根目录,绝对路径。 -->
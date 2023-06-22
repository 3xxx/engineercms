<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <title>EngineerCMS</title>
  <meta name="viewport" content="width=device-width, initial-scale=1  maximum-scale=1">
  <meta name="mobile-web-app-capable" content="yes">
  <meta name="apple-mobile-web-app-capable" content="yes">
  <meta name="apple-touch-fullscreen" content="yes">
  <meta name="HandheldFriendly" content="True">
  <link rel="stylesheet" href="/static/voxes/css/materialize.css">
  <link rel="stylesheet" href="/static/voxes/font-awesome/css/font-awesome.min.css">
  <link rel="stylesheet" href="/static/voxes/css/normalize.css">
  <link rel="stylesheet" href="/static/voxes/css/owl.carousel.css">
  <link rel="stylesheet" href="/static/voxes/css/owl.theme.css">
  <link rel="stylesheet" href="/static/voxes/css/owl.transitions.css">
  <link rel="stylesheet" href="/static/voxes/css/fakeLoader.css">
  <link rel="stylesheet" href="/static/voxes/css/magnific-popup.css">
  <link rel="stylesheet" href="/static/voxes/css/style.css">
  <link rel="shortcut icon" href="/static/voxes/img/favicon.png">
  <!-- <link href='https://fonts.loli.net/icon?family=Material+Icons' rel='stylesheet'> -->
  <style type="text/css">
  .filtr-item{ border:0px solid #000; height:165px;overflow:hidden}
  .filtr-item img{max-width: 300px;_width:expression(this.width < 300 ? "300px" : this.width);}

  html {
    /*font-family: GillSans, Calibri, Trebuchet, sans-serif;*/
    font-family: "Consolas", "Microsoft Yahei", Arial, monospace;
  }

  /*  @font-face {
    font-family: 'Material Icons';
    font-style: normal;
    font-weight: 400;
    src: url(https://fonts.gstatic.com/s/materialicons/v43/flUhRq6tzZclQEJ-Vdg-IuiaDsNc.woff2) format('woff2');
  }*/

  p {
    margin: 5px 0;
    line-height: 25px;
  }

  /*  body {
    margin: 0;
    padding: 0;
    font-size: 16px;
    font-family: 'Montserrat', sans-serif;
    background: #fff;
    color: #999;
  }*/

  /*  .material-icons {
    font-family: 'Material Icons';
    font-weight: normal;
    font-style: normal;
    font-size: 24px;
    line-height: 1;
    letter-spacing: normal;
    text-transform: none;
    display: inline-block;
    white-space: nowrap;
    word-wrap: normal;
    direction: ltr;
    -webkit-font-feature-settings: 'liga';
    -webkit-font-smoothing: antialiased;
  }*/

  .header-search {
    display: block;
    text-align: center;
    background: rgba(0, 0, 0, 0.6);
    opacity: 0;
    visibility: hidden;
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    -webkit-transition: all .3s ease-in-out;
    transition: all .3s ease-in-out;
    z-index: 999;
    cursor: pointer
  }

  .search-show .header-search {
    opacity: 0.95;
    visibility: visible
  }

  .search-form {
    width: 100%;
    position: absolute;
    top: 50%;
    -webkit-transform: translateY(-50%);
    -ms-transform: translateY(-50%);
    transform: translateY(-50%)
  }

  .search-form label {
    color: #fff
  }

  .search-form .input-search {
    background-color: #fff;
    color: #4b4b4b;
    height: auto;
    width: 100%;
    font-size: 18px;
    line-height: 1;
    border: 0;
    margin: 0 auto;
    padding: 20px 50px 20px 30px;
    width: 100%
  }

  .search-form .input-search:focus {
    outline-style: none
  }

  .search-group {
    position: relative;
    max-width: 500px;
    margin: 0 auto;
    width: 80%
  }

  .search-submit {
    position: absolute;
    right: 0;
    top: 3px;
    background: transparent;
    border: 0;
    font-size: 24px;
    bottom: 0;
    padding: 0;
    right: 15px;
    height: 100%;
    color: #ff4157
  }

  .search-submit:hover {
    color: #ff4157
  }

  .bg-dark .search-submit:hover,
  .bg-dark-alt .search-submit:hover {
    color: #ff4157
  }

  .search-close {
    padding: 30px 0;
    display: inline-block;
    color: #fff;
    font-size: 30px
  }

  .search-close:hover {
    color: rgba(255, 255, 255, 0.7)
  }

  body.search-show {
    overflow: hidden
  }

  @media(min-width:768px) {
    .search-form .input-search {
      padding: 30px
    }

    .search-submit {
      right: 25px
    }

    .search-group {
      max-width: 800px;
      width: 90%
    }
  }

  @media(min-width:992px) {
    .search-mobile {
      display: none
    }
  }
  </style>
</head>

<body>
  <!-- navbar top -->
  <div class="navbar-top">
    <div class="side-nav-panel-left"><a href="javascript:void(0)" data-activates="slide-out-left" class="side-nav-left"><i class="fa fa-bars"></i></a></div><!-- site brand -->
    <div class="site-brand"><a href="/index">
        <h1>E</h1>
      </a></div><!-- end site brand -->
    <div class="side-nav-panel-right"><a href="" class="side-nav-right"><i class="fa fa-user">{{.Username}}</i></a></div>
  </div><!-- end navbar top -->
  <!-- side nav left-->
  <div class="side-nav-panel-left">
    <ul id="slide-out-left" class="side-nav side-nav-panel collapsible">
      <li class="li-top"><a href="/index"><i class="fa fa-home"></i>Home</a></li>
      <li><a href="/project"><i class="fa fa-product-hunt"></i>Project</a></li>
      <li><a href="/standard"><i class="fa fa-scribd"></i>Standard</a></li>
      <li><a href="/v1/wx/photoswipe"><i class="fa fa-photo"></i>Photo Album</a></li>
      <li><a href="/v1/flv/flvlist"><i class="fa fa-video-camera"></i>Video</a></li>
      <!-- <li><a href="pricing.html"><i class="fa fa-dollar"></i>公告</a></li> -->
      <!-- <li><a href="error404.html"><i class="fa fa-hourglass-half"></i>404</a></li> -->
      <li><a href="/mindoc"><i class="fa fa-book"></i>Book</a></li>
      <!-- <li><a href="about-us.html"><i class="fa fa-user"></i>About Us</a></li> -->
      <!-- <li><a href="contact.html"><i class="fa fa-envelope-o"></i>Contact Us</a></li> -->
      <li><a href="/login"><i class="fa fa-sign-in"></i>Login</a></li>
      <li><a href="" onclick="return logout();"><i class="fa fa-sign-out"></i>Logout</a></li>
      <li><a href="/regist"><i class="fa fa-user-plus"></i>Register</a></li>
      <li><a class="search-trigger"><i class="fa fa-search"></i>Search</a></li>
    </ul>
  </div><!-- end side nav left-->
  <div class="header-search" onkeypress="getKey();">
    <div class="search-form">
      <div class="search-group"><input type="text" class="input-search" placeholder="Search ..." id="keyword" ><button class="search-submit" onclick="searchAllProjectsProduct()"><i class="fa fa-search"></i></button></div>
    </div>
    <!-- <form>
      <div class="input-field">
        <input id="search" type="search" required>
        <label class="label-icon" for="search"><i class="material-icons">search</i></label>
        <i class="material-icons">close</i>
      </div>
    </form> -->
  </div><!-- header-search -->
  <!-- slider -->
  <div class="slider">
    <ul class="slides">
      <li><img src="/static/voxes/img/slide1.jpg" alt="">
        <div class="caption slider-content center-align">
          <div class="container">
            <h2>Welcome to EngineerCMS</h2>
            <h4>项目资料管理.</h4><a href="" class="button-default">Learn More</a>
          </div>
        </div>
      </li>
      <li><img src="/static/voxes/img/slide2.jpg" alt="">
        <div class="caption slider-content center-align">
          <div class="container">
            <h2>A Modern Manage System</h2>
            <h4>随时随地收集资料.</h4><a href="" class="button-default">Learn More</a>
          </div>
        </div>
      </li>
      <li><img src="/static/voxes/img/slide3.jpg" alt="">
        <div class="caption slider-content center-align">
          <div class="container">
            <h2>A Fresh Collection System</h2>
            <h4>相册、视频、文字、规范查询、考勤…….</h4><a href="" class="button-default">Learn More</a>
          </div>
        </div>
      </li>
    </ul>
  </div><!-- end slider -->
  <!-- we are -->
  <div class="section we-are">
    <div class="container">
      <div class="section-head">
        <h4>设计项目知识管理</h4>
        <p>基于golang语言（beego框架）。engineercms是一款基于web的知识管理系统。支持提取码分享文件，onlyoffice实时文档协作，直接在线编辑dwg文件、office文档，在线利用mindoc创作你的书籍。通用的业务流程设置。手机端配套小程序</p>
      </div>
    </div>
  </div><!-- and we are -->
  <!-- features -->
  <div class="section features bg-second">
    <div class="container">
      <div class="row">
        <div class="col s12 m9 l6">
          <div class="card">
            <div class="card-image waves-effect waves-block waves-light">
              <img class="activator" src="https://zsj.itdos.net/attachment/SL1205%E7%8F%A0%E6%B1%9F%E4%B8%89%E8%A7%92%E6%B4%B2%E6%B0%B4%E8%B5%84%E6%BA%90%E9%85%8D%E7%BD%AE%E5%B7%A5%E7%A8%8B/%E8%AE%BE%E8%AE%A1%E5%8D%95%E4%BD%8D/%E8%AE%BE%E4%BB%A3%E6%97%A5%E8%AE%B0/2021February/1614340511137317300_small.jpg">
            </div>
            <div class="card-content">
              <span class="card-title activator grey-text text-darken-4">设代日记<i class="fa fa-ellipsis-v right"></i></span>
              <p><a href="https://zsj.itdos.net/project/product/article/1010">查看详情</a></p>
            </div>
            <div class="card-reveal">
              <span class="card-title grey-text text-darken-4">设代日记<i class="material-icons right">close</i></span>
              <p>LG02#粤海2、3号盾构机穿越大金山复杂地质施工及刀具磨损技术研讨会</p>
            </div>
          </div>
        </div>
        <div class="col s12 m9 l6">
          <div class="card">
            <div class="card-image waves-effect waves-block waves-light">
              <img class="activator" src="https://zsj.itdos.net/attachment/SL1205%E7%8F%A0%E6%B1%9F%E4%B8%89%E8%A7%92%E6%B4%B2%E6%B0%B4%E8%B5%84%E6%BA%90%E9%85%8D%E7%BD%AE%E5%B7%A5%E7%A8%8B/%E8%AE%BE%E8%AE%A1%E5%8D%95%E4%BD%8D/%E8%AE%BE%E4%BB%A3%E6%97%A5%E8%AE%B0/2021February/1614155842759316400_small.jpg">
            </div>
            <div class="card-content">
              <span class="card-title activator grey-text text-darken-4">设代日记<i class="fa fa-ellipsis-v right"></i></span>
              <p><a href="https://zsj.itdos.net/project/product/article/1008">查看详情</a></p>
            </div>
            <div class="card-reveal">
              <span class="card-title grey-text text-darken-4">设代日记<i class="material-icons right">close</i></span>
              <p>地质分院院长廖品忠指导设代工作</p>
            </div>
          </div>
        </div>
      </div>
      <!--       <div class="row">
        <div class="col s12 m9 l6">
          <div class="card">
            <div class="card-image waves-effect waves-block waves-light">
              <img class="activator" src="/static/voxes/img/about.jpg">
            </div>
            <div class="card-content">
              <span class="card-title activator grey-text text-darken-4">卡片标题<i class="fa fa-ellipsis-v right"></i></span>
              <p><a href="javascript:void(0)">这是一个链接</a></p>
            </div>
            <div class="card-reveal">
              <span class="card-title grey-text text-darken-4">卡片标题<i class="material-icons right">close</i></span>
              <p>单击后显示的产品的详细信息。</p>
            </div>
          </div>
        </div>
        <div class="col s12 m9 l6">
          <div class="card">
            <div class="card-image waves-effect waves-block waves-light">
              <img class="activator" src="/static/voxes/img/about.jpg">
            </div>
            <div class="card-content">
              <span class="card-title activator grey-text text-darken-4">卡片标题<i class="fa fa-ellipsis-v right"></i></span>
              <p><a href="javascript:void(0)">这是一个链接</a></p>
            </div>
            <div class="card-reveal">
              <span class="card-title grey-text text-darken-4">卡片标题<i class="material-icons right">close</i></span>
              <p>单击后显示的产品的详细信息。</p>
            </div>
          </div>
        </div> -->
    </div>
  </div>
  </div><!-- end features -->
  <!-- home portfolio -->
  <div class="section">
    <div class="container">
      <div class="section-head">
        <h4>工程相册</h4>
        <div class="underline"></div>
      </div>
      <div class="gallery">
        <ul class="simplefilter">
          <li class="active" data-filter="all">
            <h5>All</h5>
          </li>
          <li data-filter="1">
            <h5>工作井</h5>
          </li>
          <li data-filter="2">
            <h5>盾构</h5>
          </li>
          <li data-filter="3">
            <h5>泵站</h5>
          </li>
        </ul>
        <div class="row">
          <div class="filtr-container">
            <div class="col s12 m9 l6 filtr-item col-pd" data-category="1, 3"><a href="https://zsj.itdos.net/attachment/SL1205%E7%8F%A0%E6%B1%9F%E4%B8%89%E8%A7%92%E6%B4%B2%E6%B0%B4%E8%B5%84%E6%BA%90%E9%85%8D%E7%BD%AE%E5%B7%A5%E7%A8%8B/%E8%AE%BE%E8%AE%A1%E5%8D%95%E4%BD%8D/%E8%AE%BE%E4%BB%A3%E6%97%A5%E8%AE%B0/2021February/1612834685149969400_small.jpg" class="image-popup"><img class="responsive-img z-depth-1" src="https://zsj.itdos.net/attachment/SL1205%E7%8F%A0%E6%B1%9F%E4%B8%89%E8%A7%92%E6%B4%B2%E6%B0%B4%E8%B5%84%E6%BA%90%E9%85%8D%E7%BD%AE%E5%B7%A5%E7%A8%8B/%E8%AE%BE%E8%AE%A1%E5%8D%95%E4%BD%8D/%E8%AE%BE%E4%BB%A3%E6%97%A5%E8%AE%B0/2021February/1612834685149969400_small.jpg" alt="sample image"></a></div>
            <div class="col ss12 m9 l6 filtr-item col-pd" data-category="2, 3"><a href="https://zsj.itdos.net/attachment/SL1205%E7%8F%A0%E6%B1%9F%E4%B8%89%E8%A7%92%E6%B4%B2%E6%B0%B4%E8%B5%84%E6%BA%90%E9%85%8D%E7%BD%AE%E5%B7%A5%E7%A8%8B/%E8%AE%BE%E8%AE%A1%E5%8D%95%E4%BD%8D/%E8%AE%BE%E4%BB%A3%E6%97%A5%E8%AE%B0/2021January/1611932658884366400_small.jpg" class="image-popup"><img class="responsive-img z-depth-1" src="https://zsj.itdos.net/attachment/SL1205%E7%8F%A0%E6%B1%9F%E4%B8%89%E8%A7%92%E6%B4%B2%E6%B0%B4%E8%B5%84%E6%BA%90%E9%85%8D%E7%BD%AE%E5%B7%A5%E7%A8%8B/%E8%AE%BE%E8%AE%A1%E5%8D%95%E4%BD%8D/%E8%AE%BE%E4%BB%A3%E6%97%A5%E8%AE%B0/2021January/1611932658884366400_small.jpg" alt="sample image"></a></div>
            <div class="col s12 m9 l6 filtr-item col-pd" data-category="3, 3"><a href="https://zsj.itdos.net/attachment/SL1205%E7%8F%A0%E6%B1%9F%E4%B8%89%E8%A7%92%E6%B4%B2%E6%B0%B4%E8%B5%84%E6%BA%90%E9%85%8D%E7%BD%AE%E5%B7%A5%E7%A8%8B/%E8%AE%BE%E8%AE%A1%E5%8D%95%E4%BD%8D/%E8%AE%BE%E4%BB%A3%E6%97%A5%E8%AE%B0/2021January/1611717392212454200_small.jpg" class="image-popup"><img class="responsive-img z-depth-1" src="https://zsj.itdos.net/attachment/SL1205%E7%8F%A0%E6%B1%9F%E4%B8%89%E8%A7%92%E6%B4%B2%E6%B0%B4%E8%B5%84%E6%BA%90%E9%85%8D%E7%BD%AE%E5%B7%A5%E7%A8%8B/%E8%AE%BE%E8%AE%A1%E5%8D%95%E4%BD%8D/%E8%AE%BE%E4%BB%A3%E6%97%A5%E8%AE%B0/2021January/1611717392212454200_small.jpg" alt="sample image"></a></div>
            <div class="col s12 m9 l6 filtr-item col-pd" data-category="1, 2"><a href="https://zsj.itdos.net/attachment/SL1205%E7%8F%A0%E6%B1%9F%E4%B8%89%E8%A7%92%E6%B4%B2%E6%B0%B4%E8%B5%84%E6%BA%90%E9%85%8D%E7%BD%AE%E5%B7%A5%E7%A8%8B/%E8%AE%BE%E8%AE%A1%E5%8D%95%E4%BD%8D/%E8%AE%BE%E4%BB%A3%E6%97%A5%E8%AE%B0/2021January/1610553863132352900_small.jpg" class="image-popup"><img class="responsive-img z-depth-1" src="https://zsj.itdos.net/attachment/SL1205%E7%8F%A0%E6%B1%9F%E4%B8%89%E8%A7%92%E6%B4%B2%E6%B0%B4%E8%B5%84%E6%BA%90%E9%85%8D%E7%BD%AE%E5%B7%A5%E7%A8%8B/%E8%AE%BE%E8%AE%A1%E5%8D%95%E4%BD%8D/%E8%AE%BE%E4%BB%A3%E6%97%A5%E8%AE%B0/2021January/1610553863132352900_small.jpg" alt="sample image"></a></div>
            <div class="col s12 m9 l6 filtr-item col-pd" data-category="2, 1"><a href="https://zsj.itdos.net/attachment/SL1205%E7%8F%A0%E6%B1%9F%E4%B8%89%E8%A7%92%E6%B4%B2%E6%B0%B4%E8%B5%84%E6%BA%90%E9%85%8D%E7%BD%AE%E5%B7%A5%E7%A8%8B/%E8%AE%BE%E8%AE%A1%E5%8D%95%E4%BD%8D/%E8%AE%BE%E4%BB%A3%E6%97%A5%E8%AE%B0/2021January/1610553863852084600_small.jpg" class="image-popup"><img class="responsive-img z-depth-1" src="https://zsj.itdos.net/attachment/SL1205%E7%8F%A0%E6%B1%9F%E4%B8%89%E8%A7%92%E6%B4%B2%E6%B0%B4%E8%B5%84%E6%BA%90%E9%85%8D%E7%BD%AE%E5%B7%A5%E7%A8%8B/%E8%AE%BE%E8%AE%A1%E5%8D%95%E4%BD%8D/%E8%AE%BE%E4%BB%A3%E6%97%A5%E8%AE%B0/2021January/1610553863852084600_small.jpg" alt="sample image"></a></div>
            <div class="col s12 m9 l6 filtr-item col-pd" data-category="1, 1"><a href="https://zsj.itdos.net/attachment/SL1205%E7%8F%A0%E6%B1%9F%E4%B8%89%E8%A7%92%E6%B4%B2%E6%B0%B4%E8%B5%84%E6%BA%90%E9%85%8D%E7%BD%AE%E5%B7%A5%E7%A8%8B/%E8%AE%BE%E8%AE%A1%E5%8D%95%E4%BD%8D/%E8%AE%BE%E4%BB%A3%E6%97%A5%E8%AE%B0/2021January/1610437910191014300_small.jpg" class="image-popup"><img class="responsive-img z-depth-1" src="https://zsj.itdos.net/attachment/SL1205%E7%8F%A0%E6%B1%9F%E4%B8%89%E8%A7%92%E6%B4%B2%E6%B0%B4%E8%B5%84%E6%BA%90%E9%85%8D%E7%BD%AE%E5%B7%A5%E7%A8%8B/%E8%AE%BE%E8%AE%A1%E5%8D%95%E4%BD%8D/%E8%AE%BE%E4%BB%A3%E6%97%A5%E8%AE%B0/2021January/1610437910191014300_small.jpg" alt="sample image"></a></div>
          </div>
        </div>
      </div>
    </div>
  </div><!-- end home portfolio -->
  <!-- review -->
  <div class="section review-users bg-second">
    <div class="container">
      <div id="owl-review-users">
        <div class="item">
          <!-- <img src="img/testimonial1.jpg" alt=""> -->
          <video class="responsive-video" controls style="height: 400px">
            <source src="/static/voxes/video/movie1.mp4" type="video/mp4">
          </video>
          <h6>进库溢流井</h6>
          <p>公明水库进库闸溢流井</p>
          <ul>
            <li><i class="fa fa-star"></i></li>
            <li><i class="fa fa-star"></i></li>
            <li><i class="fa fa-star"></i></li>
            <li><i class="fa fa-star"></i></li>
            <li><i class="fa fa-star"></i></li>
          </ul>
        </div>
        <div class="item">
          <!-- <img src="img/testimonial2.jpg" alt=""> -->
          <video class="responsive-video" controls style="height: 400px">
            <source src="/static/voxes/video/movie2.mp4" type="video/mp4">
          </video>
          <h6>工作井</h6>
          <p>井壁洒水养护</p>
          <ul>
            <li><i class="fa fa-star"></i></li>
            <li><i class="fa fa-star"></i></li>
            <li><i class="fa fa-star"></i></li>
            <li><i class="fa fa-star"></i></li>
            <li><i class="fa fa-star"></i></li>
          </ul>
        </div>
        <div class="item">
          <!-- <img src="img/testimonial3.jpg" alt=""> -->
          <video class="responsive-video" controls style="height: 400px">
            <source src="/static/voxes/video/movie3.mp4" type="video/mp4">
          </video>
          <h6>Roman</h6>
          <p>视频</p>
          <ul>
            <li><i class="fa fa-star"></i></li>
            <li><i class="fa fa-star"></i></li>
            <li><i class="fa fa-star"></i></li>
            <li><i class="fa fa-star"></i></li>
            <li><i class="fa fa-star"></i></li>
          </ul>
        </div>
      </div>
    </div>
  </div><!-- end review -->
  <!-- loader -->
  <!-- <div id="fakeLoader"></div> -->
  <!-- end loader -->
  <!-- footer -->
  <div class="footer">
    <div class="container">
      <div class="about-us-foot">
        <h6>EngineerCMS</h6>
        <p>手机端、小程序、桌面端.</p>
      </div>
      <div class="social-media"><a href=""><i class="fa fa-facebook"></i></a><a href=""><i class="fa fa-twitter"></i></a><a href=""><i class="fa fa-google"></i></a><a href=""><i class="fa fa-linkedin"></i></a><a href=""><i class="fa fa-instagram"></i></a></div>
      <div class="copyright"><span>©All Right By <a href="https://zsj.itdos.net/">engineercms</a></span></div>
    </div>
  </div><!-- end footer -->
  <!-- scripts -->
  <script src="/static/voxes/js/jquery.min.js"></script>
  <script src="/static/voxes/js/materialize.min.js"></script>
  <script src="/static/voxes/js/owl.carousel.min.js"></script>
  <script src="/static/voxes/js/jquery.filterizr.js"></script>
  <script src="/static/voxes/js/jquery.magnific-popup.min.js"></script>
  <script src="/static/voxes/js/portfolio.js"></script>
  <!-- <script src="/static/voxes/js/contact-form.js"></script> -->
  <script src="/static/voxes/js/fakeLoader.min.js"></script>
  <script src="/static/voxes/js/main.js"></script>
  <script type="text/javascript">
  var searcharea = $('.header-search'),
    searchTrigger = $('.search-trigger'),
    siteBody = $('body');

  $('.search-trigger').on('click', function(e) {
    e.preventDefault();
    siteBody.addClass('search-show');
  });

  searcharea.on('click', function(e) {
    if (!$(e.target).is('.input-search')) {
      if (siteBody.hasClass('search-show')) {
        siteBody.removeClass('search-show');
      }
    }
  });

  $('.navbar-toggle').on('click', function(e) {
    var $self = $(this),
      _self_toggle = ($self.data('menu-toggle'));
    $self.toggleClass('active');
    if ($main_navbar_classic.exists()) {
      $('#' + _self_toggle).slideToggle().toggleClass(_open_menu);
    } else {
      $('#' + _self_toggle).parent().toggleClass(_open_menu);
    }
    e.preventDefault();
  });

  //登出功能
  function logout() {
    $.ajax({
      type: 'get',
      url: '/logout',
      data: {},
      success: function(result) {
        if (result.islogin) {
          // $("#status").html("登出成功");
          // alert("登出成功");
          // console.log("1");
          Materialize.toast('登出成功', 3000, 'rounded')
          // window.location.reload();
        } else {
          // $("#status").html("登出失败");
          // alert("登出失败")
          Materialize.toast('登出失败', 3000, 'rounded')
          // console.log("2");
        }
      }
    })
  };
  // 检索所有项目
  function searchAllProjectsProduct(){
    window.location.href="/v1/wx/searchproduct?keyword="+$("#keyword").val()
  }
  function getKey() {
    if (event.keyCode == 13) {
      window.location.href="/v1/wx/searchproduct?keyword="+$("#keyword").val()
    }
  }
  </script>
</body>

</html>
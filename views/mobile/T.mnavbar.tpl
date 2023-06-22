{{define "mnavbar"}}
<nav class="navbar navbar-default navbar-static-top">
<!-- <!DOCTYPE html> -->
<!-- <html lang="en"> -->

<!-- <head> -->
  <!-- <meta charset="UTF-8"> -->
  <!-- <title>EngineerCMS</title> -->
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
  <link href='https://fonts.loli.net/icon?family=Material+Icons' rel='stylesheet'>
  <style type="text/css">
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
    <div class="site-brand"><a href="home.html">
        <h1>E</h1>
      </a></div><!-- end site brand -->
    <div class="side-nav-panel-right"><a href="" class="side-nav-right"><i class="fa fa-envelope"></i></a></div>
  </div><!-- end navbar top -->
  <!-- side nav left-->
  <div class="side-nav-panel-left">
    <ul id="slide-out-left" class="side-nav side-nav-panel collapsible">
      <li class="li-top"><a href="/index"><i class="fa fa-home"></i>Home</a></li>
      <li><a href="/project"><i class="fa fa-product-hunt"></i>Project</a></li>
      <li><a href="/standard"><i class="fa fa-scribd"></i>Standard</a></li>
      <li><a href="/v1/wx/photoswipe"><i class="fa fa-photo"></i>相册</a></li>
      <li><a href="/v1/flv/flvlist"><i class="fa fa-video-camera"></i>视频</a></li>
      <!-- <li><a href="pricing.html"><i class="fa fa-dollar"></i>公告</a></li> -->
      <!-- <li><a href="error404.html"><i class="fa fa-hourglass-half"></i>404</a></li> -->
      <li><a href="/mindoc"><i class="fa fa-book"></i>Book</a></li>
      <!-- <li><a href="about-us.html"><i class="fa fa-user"></i>About Us</a></li> -->
      <!-- <li><a href="contact.html"><i class="fa fa-envelope-o"></i>Contact Us</a></li> -->
      <li><a href="/login"><i class="fa fa-sign-in"></i>Login</a></li>
      <li><a href="/regist"><i class="fa fa-user-plus"></i>Register</a></li>
      <li><a href="" class="search-trigger"><i class="fa fa-search"></i>Search</a></li>
    </ul>
  </div><!-- end side nav left-->
  <div class="header-search">
    <form role="search" method="POST" class="search-form" action="#">
      <div class="search-group"><input type="text" class="input-search" placeholder="Search ..."><button class="search-submit" type="submit"><i class="fa fa-search"></i></button></div>
    </form>
  </div><!-- . header-search -->
</nav>
{{end}}


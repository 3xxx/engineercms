<!DOCTYPE html>
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">

  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>

<script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap.min.js"></script>


<meta name="description" content="EngineerCMS - 土木工程师知识管理系统。让项目与团队管理的更加容易一点！">
<meta name="keywords" content="EngineerCMS，EngineerCMS系统，EngineerCMS办公系统，EngineerCMS项目管理，项目管理软件，CMS">
<meta name="author" content="Lock">
<title>EngineerCMS</title>
<link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>

<!--[if lt IE 9]>
      <script src="js/html5shiv.min.js"></script>
      <script src="js/respond.min.js"></script>
<![endif]-->
<link href="/static/css/engineercms.css" rel="stylesheet">
<link rel="apple-touch-icon" href="/static/img/apple-touch-icon.png">
<link rel="icon" href="/static/img/favicon.ico">
    <script>
        var _hmt = _hmt || [];
        (function() {
         var hm = document.createElement("script");
         hm.src = "//hm.baidu.com/hm.js?8f10bedad0b5bb0597750838087f6793";
         var s = document.getElementsByTagName("script")[0];
         s.parentNode.insertBefore(hm, s);
         })();
        </script>
</head>
  <div class="container-fill">
        {{template "navbar" .}}
  </div>
<body>

<div id="myCarousel" class="carousel slide" data-ride="carousel">
  <ol class="carousel-indicators">
    <li data-target="#myCarousel" data-slide-to="0" class="active"></li>
    <li data-target="#myCarousel" data-slide-to="1"></li>
    <li data-target="#myCarousel" data-slide-to="2"></li>
  </ol>
  <div class="carousel-inner" role="listbox">
    <div class="item active"><img class="first-slide" src="/static/img/index_01.jpg">
      <div class="container">
        <div class="carousel-caption">
          <h1>工程师知识管理系统</h1>
          <p style="margin-bottom:40px;">运行于个人电脑的微服务，标准化管理个人资料，轻松发布，方便知识的继承。</p>
          <p><a class="btn btn-lg btn-primary"  href="/static/download/EngineerCMS-v1-win.txt" role="button">下载EngineerCMS</a>&nbsp;<a class="btn btn-lg btn-primary" href="http://112.74.42.44/project" target="_blank" role="button">现在去体验</a></p>
        </div>
      </div>
    </div>
    <div class="item"> <img class="first-slide" src="/static/img/index_02.jpg">
      <div class="container">
        <div class="carousel-caption">
          <h1>MeritMS团队成果和价值管理</h1>
          <p style="margin-bottom:40px;">采用GoLang语言和Bootstrap框架，开源、安全、先进、简洁、高速。</p>
          <p><a class="btn btn-lg btn-primary" href="/static/download/MeritMS-v1-win.txt" role="button">下载MeritMS</a>&nbsp;<a class="btn btn-lg btn-primary" href="http://112.74.42.44:8080/" target="_blank" role="button">现在去体验</a></p>
        </div>
      </div>
    </div>
    <div class="item"> <img class="first-slide" src="/static/img/index_03.jpg">
      <div class="container">
        <div class="carousel-caption">
          <h1>水利供水管线设计工具</h1>
          <p style="margin-bottom:40px;">采用Excel结合AutoCAD，开源、简洁、高速。</p>
          <p><a class="btn btn-lg btn-primary" href="/static/download/HydroWS-v1-win.txt" role="button">下载HydroWS</a>&nbsp;<a class="btn btn-lg btn-primary" href="#contact" role="button">联系技术解决问题</a></p>
        </div>
      </div>
    </div>
  </div>
  <a class="left carousel-control" href="#myCarousel" role="button" data-slide="prev"> <span class="glyphicon glyphicon-chevron-left" aria-hidden="true"></span> <span class="sr-only">Previous</span> </a> <a class="right carousel-control" href="#myCarousel" role="button" data-slide="next"> <span class="glyphicon glyphicon-chevron-right" aria-hidden="true"></span> <span class="sr-only">Next</span> </a> </div>
<div class="container marketing"> <a name="introduction"></a>
  <div class="row featurette">
    <div class="col-md-7">
      <h2 class="featurette-heading">界面优雅</h2>
      <p class="lead">界面优雅，简洁大方。更加适合用户操作习惯</p>
    </div>
    <div class="col-md-5"> <img class="featurette-image img-responsive center-block" src="/static/img/c1.png"> </div>
  </div>
  <hr class="featurette-divider">
  <div class="row featurette">
    <div class="col-md-7 col-md-push-5">
      <h2 class="featurette-heading">知识管理</h2>
      <p class="lead">系统结合了档案管理系统形式、SharePoint‘点-平台’理念、ProjectWise协同设计需求。</p>
    </div>
    <div class="col-md-5 col-md-pull-7"> <img class="featurette-image img-responsive center-block" src="/static/img/c2.png"> </div>
  </div>
  <hr class="featurette-divider">
  <div class="row featurette">
    <div class="col-md-7">
      <h2 class="featurette-heading">项目管理</h2>
      <p class="lead">提供精简的项目管理功能，减少冗余的流程，更加扁平化。</p>
    </div>
    <div class="col-md-5"> <img class="featurette-image img-responsive center-block" src="/static/img/c3.png"> </div>
  </div>
  <hr class="featurette-divider">
  <div class="row featurette">
    <div class="col-md-7 col-md-push-5">
      <h2 class="featurette-heading">简便的成果和价值管理</h2>
      <p class="lead">方便管理员工设计成果、校审流程、绩效考核、价值评测...</p>
    </div>
    <div class="col-md-5 col-md-pull-7"> <img class="featurette-image img-responsive center-block" src="/static/img/c4.png"> </div>
  </div>
  <hr class="featurette-divider">
  <div class="row"> <a name="contact"></a>
    <div class="col-lg-4"> <img class="img-circle" src="/static/img/qin.png" width="140" height="140">
      <h2>Qin</h2>
      <p>同时关注工程师个人的知识和团队的成果，幸好上线了EngineerCMS系统，不仅操作简单易学，界面清晰、操作流畅 </p>
    </div>
    <div class="col-lg-4"> <img class="img-circle" src="/static/img/fancy.png" width="140" height="140">
      <h2>Fancy</h2>
      <p>以前同时管理好几个项目，又要关注团队建设，常常让我分身乏术，自从安装了EngineerCMS系统和MeritMS系统，随时掌握项目进度、团队绩效，团队水平提高很快。</p>
    </div>
    <div class="col-lg-4"> <img class="img-circle" src="/static/img/ys.png" width="140" height="140">
      <h2>贴心服务</h2>
      <p>贴心管家服务，解决问题，传递快乐！</p>
      <h3>QQ：504284</h3>
	  <h3>技术解答：<a target="_blank" href="http://wpa.qq.com/msgrd?v=3&uin=504284&site=qq&menu=yes"><img border="0" src="http://wpa.qq.com/pa?p=2:504284:52" alt="EngineerCMS在线解答" title="EngineerCMS在线解答"/></a></h3>
    </div>
  </div>
  <hr class="featurette-divider">
  <a name="deploy"></a>
  <div class="row">
    <div class="col-lg-4">
      <h2>独立部署</h2>
      <p style="text-align:left;line-height:26px; background-color: #d2dce5;padding: 15px;"> 独立部署即为在后端运行程序，让程序跑在后台。<br>
        <br/>
        <b>linux</b> <br/>
        在 linux 下面部署，我们可以利用 nohup 命令，把应用部署在后端，如下所示： <br/>
        nohup ./EngineerCMS & <br/>
        这样你的应用就跑在了 Linux 系统的守护进程 <br>
        <br/>
        <b>Windows</b> <br/>
        在 Windows系 统中，设置开机自动，后台运行，有如下几种方式： <br/>
        1.制作 bat 文件，放在“启动”里面<br/>
        2.制作成服务 <br>
        <br>
      </p>
    </div>
    <div class="col-lg-4">
      <h2>Apache部署</h2>
      <p style="text-align:left;line-height:26px; background-color: #d2dce5;padding: 15px;"> &lt;VirtualHost *:80&gt;<br/>
        &nbsp;&nbsp;ServerAdmin EngineerCMS@test.com<br/>
        &nbsp;&nbsp; ServerName EngineerCMS.me<br/>
        &nbsp;&nbsp;ProxyRequests Off<br/>
        &nbsp;&nbsp;&lt;Proxy *&gt;<br/>
        &nbsp;&nbsp;&nbsp;&nbsp;Order deny,allow<br/>
        &nbsp;&nbsp;&nbsp;&nbsp;Allow from all<br/>
        &nbsp;&nbsp;&lt;/Proxy&gt;<br/>
        &nbsp;&nbsp; ProxyPass / http://127.0.0.1:8088/<br/>
        &nbsp;&nbsp; ProxyPassReverse / http://127.0.0.1:8088/<br/>
        &lt;/VirtualHost&gt;<br>
        <br>
        <br>
        <br>
      </p>
    </div>
    <div class="col-lg-4">
      <h2>Nginx部署</h2>
      <p style="text-align:left;line-height:26px; background-color: #d2dce5;padding: 15px;"> server {<br/>
        listen       80;<br/>
        &nbsp;&nbsp;server_name  .EngineerCMS.me;<br/>
        &nbsp;&nbsp;charset utf-8; <br/>
        &nbsp;&nbsp;location / {<br/>
        &nbsp;&nbsp;&nbsp;&nbsp;try_files /_not_exists_ @backend;<br/>
        &nbsp;&nbsp;}<br/>
        &nbsp;&nbsp;location @backend {<br/>
        &nbsp;&nbsp;&nbsp;proxy_set_header X-Forwarded-For $remote_addr;<br/>
        &nbsp;&nbsp;&nbsp;proxy_set_header Host            $http_host;<br/>
        &nbsp;&nbsp;&nbsp;proxy_pass http://127.0.0.1:8088;<br/>
        &nbsp;&nbsp;}<br/>
        }<br/>
      </p>
    </div>
  </div>
  <hr class="featurette-divider">
  <footer>
    <p class="pull-right"><a href="javascript:void(0)">Back to top</a></p>
    <!-- <p>&copy; 2017 Lock &middot; <a href="http://www.milu365.com" target="_blank">迷路365</a> &middot; <a href="http://mianshi.milu365.cn" target="_blank">面试小助手</a> </p> -->
    <h4>Copyright © 2016-2018 EngineerCMS</h4>
    <p>
      网站由 <i class="user icon"></i>
      <a target="_blank" href="https://github.com/3xxx">@3xxx</a>
      建设，并由
      <a target="_blank" href="http://golang.org">golang</a>
      和
      <a target="_blank" href="http://beego.me">beego</a>
      提供动力。
    </p>

    <p>
      请给 <i class="glyphicon glyphicon-envelope"></i>
      <a class="email" href="mailto:504284@qq.com">我</a>
      发送反馈信息或提交
      <i class="tasks icon"></i>
      <a target="_blank" href="https://github.com/3xxx/engineercms/issues">网站问题</a>
      。
    </p>
  </footer>
</div>
<!-- <script src="/static/js/bootstrap.min.js"></script> -->
</body>
</html>

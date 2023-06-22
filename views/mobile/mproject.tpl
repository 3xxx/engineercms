<!-- 具体一个项目的侧栏，右侧为project_products.tpl,显示任意侧栏下的成果 -->
<!DOCTYPE html>

<head>
  <title>项目详细-EngiCMS</title>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css" />
  <!-- <link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap-theme.min.css"> -->
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script src="/static/js/bootstrap-treeview.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-treeview.css" />
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
  
  <link rel="stylesheet" href="/static/froala/css/froala_editor.css">
  <link rel="stylesheet" href="/static/froala/css/froala_style.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/code_view.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/draggable.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/colors.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/emoticons.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/image_manager.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/image.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/line_breaker.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/table.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/char_counter.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/video.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/fullscreen.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/file.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/quick_insert.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/help.css">
  <!-- 上传文件 -->
  <script src="/static/dropzone/dropzone.js"></script>
  <!-- <link rel="stylesheet" href="/static/dropzone/dropzone.css"> -->
  <script>
    Dropzone.autoDiscover = false;
  </script>
  <!-- <link rel="stylesheet" href="/static/dropzone/style.css"> -->
  <style type="text/css">
  button {margin: 10px;}
  .navbar-default {
    background-color: #e74c3c;
    border-color: #c0392b;
  }

  .navbar-default .navbar-brand {
    color: #ecf0f1;
  }

  .navbar-default .navbar-brand:hover,
  .navbar-default .navbar-brand:focus {
    color: #ffbbbc;
  }

  .navbar-default .navbar-text {
    color: #ecf0f1;
  }

  .navbar-default .navbar-nav>li>a {
    color: #ecf0f1;
  }

  .navbar-default .navbar-nav>li>a:hover,
  .navbar-default .navbar-nav>li>a:focus {
    color: #ffbbbc;
  }

  .navbar-default .navbar-nav>.active>a,
  .navbar-default .navbar-nav>.active>a:hover,
  .navbar-default .navbar-nav>.active>a:focus {
    color: #ffbbbc;
    background-color: #c0392b;
  }

  .navbar-default .navbar-nav>.open>a,
  .navbar-default .navbar-nav>.open>a:hover,
  .navbar-default .navbar-nav>.open>a:focus {
    color: #ffbbbc;
    background-color: #c0392b;
  }

  .navbar-default .navbar-toggle {
    border-color: #c0392b;
  }

  .navbar-default .navbar-toggle:hover,
  .navbar-default .navbar-toggle:focus {
    background-color: #c0392b;
  }

  .navbar-default .navbar-toggle .icon-bar {
    background-color: #ecf0f1;
  }

  .navbar-default .navbar-collapse,
  .navbar-default .navbar-form {
    border-color: #ecf0f1;
  }

  .navbar-default .navbar-link {
    color: #ecf0f1;
  }

  .navbar-default .navbar-link:hover {
    color: #ffbbbc;
  }

  @media (max-width: 767px) {
    .navbar-default .navbar-nav .open .dropdown-menu>li>a {
      color: #ecf0f1;
    }

    .navbar-default .navbar-nav .open .dropdown-menu>li>a:hover,
    .navbar-default .navbar-nav .open .dropdown-menu>li>a:focus {
      color: #ffbbbc;
    }

    .navbar-default .navbar-nav .open .dropdown-menu>.active>a,
    .navbar-default .navbar-nav .open .dropdown-menu>.active>a:hover,
    .navbar-default .navbar-nav .open .dropdown-menu>.active>a:focus {
      color: #ffbbbc;
      background-color: #c0392b;
    }
  }

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

  .badge {
    display: inline-block !important;
    min-width: 10px !important;
    padding: 3px 7px !important;
    font-size: 12px !important;
    font-weight: 700 !important;
    line-height: 1 !important;
    color: #fff !important;
    text-align: center !important;
    white-space: nowrap !important;
    vertical-align: baseline !important;
    background-color: #777 !important;
    border-radius: 10px !important;
  }

  .section {
    padding: 35px 0;
    padding-top: 35px;
    padding-right: 15px;
    padding-bottom: 0px;
    padding-left: 15px;
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

  /*dropz*/
/*  html,
  body {
    height: 100%;
  }*/

  #actions {
    margin: 2em 0;
  }

  /* Mimic table appearance */
  div.table {
    display: table;
  }

  div.table .file-row {
    display: table-row;
  }

  div.table .file-row>div {
    display: table-cell;
    vertical-align: top;
    border-top: 1px solid #ddd;
    padding: 8px;
  }

  div.table .file-row:nth-child(odd) {
    background: #f9f9f9;
  }

  /* The total progress gets shown by event listeners */
  #total-progress {
    opacity: 0;
    transition: opacity 0.3s linear;
  }

  /* Hide the progress bar when finished */
  #previews .file-row.dz-success .progress {
    opacity: 0;
    transition: opacity 0.3s linear;
  }

  /* Hide the delete button initially */
  #previews .file-row .delete {
    display: none;
  }

  /* Hide the start and cancel buttons and show the delete button */
  #previews .file-row.dz-success .start,
  #previews .file-row.dz-success .cancel {
    display: none;
  }

  #previews .file-row.dz-success .delete {
    display: block;
  }
  </style>
</head>

<body style="height: auto; overflow-y: auto;">
  <!-- navbar top -->
  <div class="navbar-top">
    <div class="side-nav-panel-left"><a href="javascript:void(0)" data-activates="slide-out-left" class="side-nav-left"><i class="fa fa-bars"></i></a></div><!-- site brand -->
    <div class="site-brand"><a href="/index">
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
    <!-- <form>
      <div class="input-field">
        <input id="search" type="search" required>
        <label class="label-icon" for="search"><i class="material-icons">search</i></label>
        <i class="material-icons">close</i>
      </div>
    </form> -->
  </div><!-- . header-search -->
  <div class="section col-xs-12 col-sm-6 col-md-3 col-lg-3">
    <div id="tree" class="section"></div>
  </div>
  <div class="col-xs-12 col-sm-6 col-md-9 col-lg-9">
    <div class="breadcrumbs">
      <ol class="breadcrumb" split="&gt;">
        <li>
          <a href="javascript:gototree({{.Category.Id}})"> <i class="fa fa-home" aria-hidden="true"></i>
            项目编号：{{.Category.Code}}
          </a>
        </li>
      </ol>
    </div>
    <iframe src="/project/{{.Id}}/{{.Id}}" name='iframepage' id="iframepage" frameborder="0" width="100%" scrolling="no" marginheight="0" marginwidth="0" onload="this.height=800"></iframe>
  </div>
  <div class="fixed-action-btn">
    <a class="btn-floating btn-large red">
      <i class="large material-icons">add</i>
    </a>
    <ul>
      <li><a class="btn-floating red" onclick="addarticle()"><i class="material-icons">library_books</i></a></li>
      <!-- <li><a class="btn-floating yellow darken-1" onclick="addarticle()"><i class="material-icons">format_quote</i></a></li> -->
      <!-- <li><a class="btn-floating green"><i class="material-icons">publish</i></a></li> -->
      <li><a class="btn-floating blue" onclick="addattachment()"><i class="material-icons">attach_file</i></a></li>
    </ul>
  </div>
  <!-- 批量上传 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalTable" style="width: 100%">
      <div class="modal-dialog" id="modalDialog">
        <div class="modal-content">
          <div class="modal-header" style="background-color: #8bc34a">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">批量添加成果</h3>
            <label>**请选择标准命名的电子文件上传**</label>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group" style="width: 100%;">
                <label class="col-xs-12 control-label">关键字</label>
                <div class="col-xs-12">
                  <input type="tel" class="form-control" id="prodlabel" name="prodlabel" placeholder="以英文,号分割"></div>
              </div>
              <div class="form-group" style="width: 100%;">
                <label class="col-xs-12 control-label">设计</label>
                <div class="col-xs-12">
                  <input type="tel" class="form-control" id="prodprincipal" name="prodprincipal"></div>
              </div>

              <div id="actions" class="row">
                <div class="col-xs-12">
                  <button class="btn btn-success fileinput-button col-xs-12">
                    <i class="glyphicon glyphicon-plus"></i>
                    <span>Add files...</span>
                  </button>
                  <button type="submit" class="btn btn-primary start col-xs-12" id="uploadfileppp">
                    <i class="glyphicon glyphicon-upload"></i>
                    <span>Start upload</span>
                  </button>
                  <button type="reset" class="btn btn-warning cancel col-xs-12">
                    <i class="glyphicon glyphicon-ban-circle"></i>
                    <span>Cancel upload</span>
                  </button>
                </div>
                <div class="col-xs-12">
                  <span class="fileupload-process" id="filedropzone">
                    <div id="total-progress" class="progress progress-striped active" role="progressbar" aria-valuemin="0" aria-valuemax="100" aria-valuenow="0">
                      <div class="progress-bar progress-bar-success" style="width:0%;" data-dz-uploadprogress></div>
                    </div>
                  </span>
                </div>
              </div>
              <div class="table table-striped files" id="previews">
                <div id="template" class="file-row">
                  <div class="col-xs-4">
                    <span class="preview"><img data-dz-thumbnail /></span>
                  </div>
                  <div class="col-xs-4">
                    <p class="name" data-dz-name></p>
                    <strong class="error text-danger" data-dz-errormessage></strong>
                  </div>
                  <div class="col-xs-4">
                    <p class="size" data-dz-size></p>
                    <div class="progress progress-striped active" role="progressbar" aria-valuemin="0" aria-valuemax="100" aria-valuenow="0">
                      <div class="progress-bar progress-bar-success" style="width:0%;" data-dz-uploadprogress></div>
                    </div>
                  </div>
                  <div class="col-xs-12">
                    <button class="btn btn-primary start">
                      <i class="glyphicon glyphicon-upload"></i>
                      <span>Start</span>
                    </button>
                    <button data-dz-remove class="btn btn-warning cancel">
                      <i class="glyphicon glyphicon-ban-circle"></i>
                      <span>Cancel</span>
                    </button>
                    <button data-dz-remove class="btn btn-danger delete">
                      <i class="glyphicon glyphicon-trash"></i>
                      <span>Delete</span>
                    </button>
                  </div>
                </div>
              </div>

              <!-- <p>请上传文件</p>
              <form id="filedropzone" class="dropzone dz-clickable">
                <div class="dz-default dz-message">
                  <div class="dz-icon icon-wrap icon-circle icon-wrap-md">
                    <i class="fa fa-cloud-upload fa-3x"></i>
                  </div>
                </div>
              </form>
              <div id="dropz" class="dropzone">
                <div id="pro"></div>
                <div id="speed"></div>
                <div id="time"></div>
              </div> -->
            </div>
          </div>
          <div class="modal-footer">
            <!-- <a href="javascript:" id="cancel" class="btn btn-primary radius">取消</a> -->
            <!-- <a href="javascript:" id="qr" class="btn btn-primary radius">提交</a> -->
            <a href="javascript:" class="btn btn-default" data-dismiss="modal">关闭</a>
          </div>
            </div>
          </div>
        </div>
      </div>
      <!-- 添加文章 -->
      <div class="modal fade" id="modalTable2" style="width: 100%">
        <!-- <div class="modal-dialog" style="width: 100%" id="modalDialog2"> -->
        <div class="modal-content">
          <div class="modal-header" style="background-color: #FF5722;">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">添加文章</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <!-- <div class="form-group must" style="float: left;width: 50%;"> -->
              <label class="col-xs-12 control-label" style="padding-left:10px;">编号</label>
              <div class="col-xs-12">
                <input type="text" class="form-control" id="prodcode1" name="prodcode1"></div>
              <!-- </div> -->
              <!-- <div class="form-group must" style="float: left;width: 50%;"> -->
              <label class="col-xs-12 control-label" style="padding-left:10px;">标题</label>
              <div class="col-xs-12">
                <input type="tel" class="form-control" id="prodname1" name="prodname1"></div>
              <!-- </div> -->
              <!-- <div class="form-group must" style="float: left;width: 50%;"> -->
              <label class="col-xs-12 control-label" style="padding-left:10px;">副标题</label>
              <div class="col-xs-12">
                <input type="tel" class="form-control" id="subtext1" name="subtext1"></div>
              <!-- </div> -->
              <!-- <div class="form-group must" style="float: left;width: 50%;"> -->
              <label class="col-xs-12 control-label" style="padding-left:10px;">关键字</label>
              <div class="col-xs-12">
                <input type="tel" class="form-control" id="prodlabel2" name="prodlabel2" placeholder="以英文,号分割"></div>
              <!-- </div> -->
              <!-- <div class="form-group must" style="float: left;width: 50%;"> -->
              <label class="col-xs-12 control-label" style="padding-left:10px;">设计</label>
              <div class="col-xs-12">
                <input type="tel" class="form-control" id="prodprincipal2" name="prodprincipal2"></div>
              <!-- </div> -->
              <!-- <div class="form-group must" style="float: left;width: 50%;"> -->
              <label class="col-xs-12 control-label" style="padding-left:10px;">关联文件</label>
              <div class="col-xs-12">
                <!-- <form name="myform">  -->
                <input type="checkbox" name="box" id="box" value="1" onclick="station_select()" style="width:30px ;height: 24px">
              </div>
              <div class="col-xs-12">
                <input type="tel" class="form-control" id="relevancy" name="relevancy" disabled="true" placeholder="输入文件编号，以英文,号分割">
                <!-- <input type="text" name="aa" id="text">  -->
                <!-- </form> -->
              </div>
            </div>
          </div>
          <label class="control-label">文章正文:</label>
          <div id="editor" style="width: 100%;padding: 10px">
            <div id='edit' style="margin-top: 30px;">
            </div>
          </div>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
          <button type="button" class="btn btn-primary" onclick="save2()">保存</button>
        </div>
        <!-- </div> -->
      </div>
      <!-- <script src="/static/voxes/js/jquery.min.js"></script> -->
      <script src="/static/voxes/js/materialize.min.js"></script>
      <script src="/static/voxes/js/owl.carousel.min.js"></script>
      <!-- <script src="/static/voxes/js/jquery.filterizr.js"></script> -->
      <!-- <script src="/static/voxes/js/jquery.magnific-popup.min.js"></script> -->
      <!-- <script src="/static/voxes/js/portfolio.js"></script> -->
      <!-- <script src="/static/voxes/js/contact-form.js"></script> -->
      <script src="/static/voxes/js/fakeLoader.min.js"></script>
      <script src="/static/voxes/js/main.js"></script>
      <script type="text/javascript" src="/static/froala/js/codemirror.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/xml.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/froala_editor.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/align.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/char_counter.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/code_beautifier.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/code_view.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/colors.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/draggable.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/emoticons.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/entities.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/file.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/font_size.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/font_family.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/fullscreen.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/image.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/image_manager.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/line_breaker.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/inline_style.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/link.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/lists.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/paragraph_format.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/paragraph_style.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/quick_insert.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/quote.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/table.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/save.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/url.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/video.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/help.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/print.min.js"></script>
      <!-- <script type="text/javascript" src="/static/froala/js/third_party/spell_checker.min.js"></script> -->
      <script type="text/javascript" src="/static/froala/js/plugins/special_characters.min.js"></script>
      <script type="text/javascript" src="/static/froala/js/plugins/word_paste.min.js"></script>
      <script src="/static/froala/js/languages/zh_cn.js"></script>
      <script type="text/javascript">
      var projid = {{.Id }}

      $(function() {
        $('#tree').treeview({
          data: [{{.json }}],
          levels: 2,
          showTags: true,
          loadingIcon: "fa fa-hourglass",
          lazyLoad: loaddata,
        });


        $('#tree').on('nodeSelected', function(event, data) {
          document.getElementById("iframepage").src = "/project/{{.Id}}/" + data.id;
          projid = data.id
          $.ajax({
            type: "get",
            url: "/project/navbar/" + data.id,
            success: function(data, status) {
              $(".breadcrumb #nav").remove();
              for (i = 0; i < data.length; i++) {
                $(".breadcrumb").append('<li id="nav"><a href="javascript:gototree(' + data[i].Id + ')">' + data[i].Title + '</a></li>');
              }
            }
          });
        });

        $("#btn").click(function(e) {
          var arr = $('#tree').treeview('getSelected');
          for (var key in arr) {
            c.innerHTML = c.innerHTML + "," + arr[key].id;
          }
        });
      })

      function loaddata(node, func) {
        $.ajax({
          type: "get",
          url: "/project/getprojcate",
          data: { id: node.id },
          success: function(data, status) {
            if (data) {
              func(data);
            }
          }
        });
      }

      function gototree(e) {
        document.getElementById("iframepage").src = "/project/{{.Id}}/" + e;
        var findCheckableNodess = function() {
          return $('#tree').treeview('findNodes', [e, 'id']);
        };
        var checkableNodes = findCheckableNodess();
        $('#tree').treeview('toggleNodeSelected', [checkableNodes, { silent: true }]);
        $('#tree').treeview('toggleNodeExpanded', [checkableNodes, { silent: true }]);
        $('#tree').treeview('revealNode', [checkableNodes, { silent: true }]);
      }

      // 添加附件
      function addattachment() {
        $('#modalTable').modal({
          show: true,
          backdrop: 'static'
        });
      }

      // 添加文章
      function addarticle() {
        $('#modalTable2').modal({
          show: true,
          backdrop: 'static'
        });
      }

      function index1(value, row, index) {
        return index + 1
      }

      function reinitIframe() {
        var iframe = document.getElementById("iframepage");
        try {
          var bHeight = iframe.contentWindow.document.body.scrollHeight;
          var dHeight = iframe.contentWindow.document.documentElement.scrollHeight;
          var height = Math.max(bHeight, dHeight, 800);
          iframe.height = height;

        } catch (ex) {}
      }
      window.setInterval("reinitIframe()", 200);

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

      $(function() {
        //超大屏幕'fullscreen',
        var toolbarButtons = ['bold', 'italic', 'underline', 'strikeThrough', 'subscript', 'superscript', 'fontFamily', 'fontSize', '|', 'color', 'emoticons', 'inlineStyle', 'paragraphStyle', '|', 'paragraphFormat', 'align', 'formatOL', 'formatUL', 'outdent', 'indent', 'quote', 'insertHR', '-', 'insertLink', 'insertImage', 'insertVideo', 'insertFile', 'insertTable', 'undo', 'redo', 'clearFormatting', 'selectAll', 'html', 'help'];
        //大屏幕
        var toolbarButtonsMD = ['bold', 'italic', 'underline', 'strikeThrough', 'subscript', 'superscript', '|', 'fontFamily', 'fontSize', 'color', 'inlineStyle', 'paragraphStyle', '|', 'paragraphFormat', 'align', 'formatOL', 'formatUL', 'outdent', 'indent', '-', 'quote', 'insertLink', 'insertImage', 'insertVideo', 'insertFile', 'insertTable', '|', 'specialCharacters', 'insertHR', 'undo', 'redo', 'clearFormatting', '|', 'html', 'help'];
        //小屏幕'fullscreen',
        var toolbarButtonsSM = ['bold', 'italic', 'underline', 'fontFamily', 'fontSize', 'insertLink', 'insertImage', 'insertTable', 'undo', 'redo'];
        //手机
        var toolbarButtonsXS = ['insertImage', 'insertVideo', 'insertFile', 'bold', 'italic', 'fontFamily', 'fontSize', 'undo', 'redo'];
        // var pid = $('#pid').val();
        //编辑器初始化并赋值 
        $('#edit').froalaEditor({
          placeholderText: '请输入内容',
          charCounterCount: true, //默认
          // charCounterMax         : -1,//默认
          saveInterval: 0, //不自动保存，默认10000
          // theme                    : "red",
          height: "300px",
          toolbarBottom: false, //默认
          toolbarButtonsMD: toolbarButtons, //toolbarButtonsMD,
          toolbarButtonsSM: toolbarButtonsMD, //toolbarButtonsSM,
          toolbarButtonsXS: toolbarButtonsXS,
          toolbarInline: false, //true选中设置样式,默认false
          imageUploadMethod: 'POST',
          heightMin: 450,
          charCounterMax: 3000,
          // imageUploadURL: "uploadImgEditor",
          imageParams: { postId: "123" },
          params: {
            acl: '01',
            AWSAccessKeyId: '02',
            policy: '03',
            signature: '04',
          },
          autosave: true,
          autosaveInterval: 2500,
          saveURL: 'hander/FroalaHandler.ashx',
          saveParams: { postId: '1' },
          spellcheck: false,
          imageUploadURL: '/uploadimg', //上传到本地服务器
          imageUploadParams: { pid: '{{.Id}}' },
          imageDeleteURL: 'lib/delete_image.php', //删除图片
          imagesLoadURL: 'lib/load_images.php', //管理图片
          videoUploadURL: '/uploadvideo',
          videoUploadParams: { pid: '{{.Id}}' },
          fileUploadURL: '/uploadimg',
          fileUploadParams: { pid: '{{.Id}}' },
          enter: $.FroalaEditor.ENTER_BR,
          language: 'zh_cn',
          // toolbarButtons: ['bold', 'italic', 'underline', 'paragraphFormat', 'align','color','fontSize','insertImage','insertTable','undo', 'redo']
        });
      })

      //添加文章
      function save2() {
        // var radio =$("input[type='radio']:checked").val();
        var projectid = projid //$('#pid').val();
        var prodcode = $('#prodcode1').val();
        var prodname = $('#prodname1').val();
        var subtext = $('#subtext1').val();
        var prodprincipal = $('#prodprincipal2').val();
        var prodlabel = $('#prodlabel2').val();
        var relevancy = $('#relevancy2').val();
        var html = $('div#edit').froalaEditor('html.get'); //$('#edit')[0].childNodes[1].innerHTML;
        // $('#myModal').on('hide.bs.modal', function () {  
        if (prodname && prodcode) {
          $.ajax({
            type: "post",
            url: "/project/product/addarticle",
            data: { pid: projectid, code: prodcode, title: prodname, subtext: subtext, label: prodlabel, content: html, principal: prodprincipal, relevancy: relevancy }, //父级id
            success: function(data, status) {
              alert("添加“" + data + "”成功！(status:" + status + ".)");
              $('#modalTable2').modal('hide');
              // $('#table0').bootstrapTable('refresh', { url: '/project/products/' + {{.Id }}});
            },
          });
        } else {
          alert("请填写编号和名称！");
          return;
        }
      }

      // 上传文件
      // Get the template HTML and remove it from the doument
      var previewNode = document.querySelector("#template");
      previewNode.id = "";
      var previewTemplate = previewNode.parentNode.innerHTML;
      previewNode.parentNode.removeChild(previewNode);

      // var myDropzone = new Dropzone(document.body, { // Make the whole body a dropzone
      $("#filedropzone").dropzone({
        url: "/project/product/addattachment", // Set the url
        thumbnailWidth: 80,
        thumbnailHeight: 80,
        parallelUploads: 20,//可修改手动触发一次上传的数量

        maxFiles: 5, //指的是上传目录下的最大文件数
        maxFilesize: 1024,
        acceptedFiles: ".jpg,.doc,.docx,.xlsx,.xls,.ppt,.pptx,.dwg,.dgn,.pdf,.mp3,.mp4,.rar,.7z,.zip",
        // dictDefaultMessage: "拖入需要上传的文件",
        //   uploadMultiple: 默认false,
        dictInvalidFileType: "支持的文件类型是office,.dwg,.dgn,.pdf,.rar,.zip,.7z",
        dictResponseError: "文件上传失败!",
        dictFileTooBig: "文件过大,上传失败!",

        previewTemplate: previewTemplate,
        autoQueue: false, // Make sure the files aren't queued until manually added
        autoProcessQueue: true,//这里用true才能实现单个文件上传，也才能实现多个文件上传
        previewsContainer: "#previews", // Define the container to display the previews
        clickable: ".fileinput-button", // Define the element that should be used as click trigger to select files.
      // });
    init: function() {
      var myDropzone = this;
      myDropzone.on("addedfile", function(file) {
        // Hookup the start button
        file.previewElement.querySelector(".start").onclick = function() { 
          // alert("ok0")
          // alert(file.status)
          // alert(file.accepted)
          myDropzone.enqueueFile(file); 
        };
      });

      // Update the total progress bar
      myDropzone.on("totaluploadprogress", function(progress) {
        document.querySelector("#total-progress .progress-bar").style.width = progress + "%";
      });

      myDropzone.on("sending", function(file, xhr, formData) {
        // Show the total progress bar when upload starts
        document.querySelector("#total-progress").style.opacity = "1";
        // And disable the start button
        file.previewElement.querySelector(".start").setAttribute("disabled", "disabled");
        formData.append('pid', projid);
        formData.append('prodlabel', $('#prodlabel').val());
        formData.append('prodprincipal', $('#prodprincipal').val());
      });

      // Hide the total progress bar when nothing's uploading anymore
      myDropzone.on("queuecomplete", function(progress) {
        document.querySelector("#total-progress").style.opacity = "0";
      });

      // Setup the buttons for all transfers
      // The "add files" button doesn't need to be setup because the config
      // `clickable` has already been specified.
      document.querySelector("#actions .start").onclick = function() {
        // myDropzone.processQueue();
        myDropzone.enqueueFiles(myDropzone.getFilesWithStatus(Dropzone.ADDED));
      };
      document.querySelector("#actions .cancel").onclick = function() {
        myDropzone.removeAllFiles(true);
      };
      // Now fake the file upload, since GitHub does not handle file uploads
      // and returns a 404
      // var minSteps = 6,
      //     maxSteps = 60,
      //     timeBetweenSteps = 100,
      //     bytesPerStep = 100000;

      // myDropzone.uploadFiles = function(files) {
      //   var self = this;
      //   for (var i = 0; i < files.length; i++) {

      //     var file = files[i];
      //     totalSteps = Math.round(Math.min(maxSteps, Math.max(minSteps, file.size / bytesPerStep)));

      //     for (var step = 0; step < totalSteps; step++) {
      //       var duration = timeBetweenSteps * (step + 1);
      //       setTimeout(function(file, totalSteps, step) {
      //         return function() {
      //           file.upload = {
      //             progress: 100 * (step + 1) / totalSteps,
      //             total: file.size,
      //             bytesSent: (step + 1) * file.size / totalSteps
      //           };

      //           self.emit('uploadprogress', file, file.upload.progress, file.upload.bytesSent);
      //           if (file.upload.progress == 100) {
      //             file.status = Dropzone.SUCCESS;
      //             self.emit("success", file, 'success', null);
      //             self.emit("complete", file);
      //             self.processQueue();
      //           }
      //         };
      //       }(file, totalSteps, step), duration);
      //     }
      //   }
      // }
    }
  });
      // var fileArr = new Array();
      // Dropzone.autoDiscover = false;
      // // $(document).ready(function() {
      // $("#filedropzone").dropzone({
      //   url: "/project/product/addattachment",
      //   maxFiles: 5, //指的是上传目录下的最大文件数
      //   maxFilesize: 1024,
      //   acceptedFiles: ".jpg,.doc,.docx,.xlsx,.xls,.ppt,.pptx,.dwg,.dgn,.pdf,.mp3,.mp4,.rar,.7z,.zip",
      //   autoProcessQueue: false,
      //   paramName: "file",
      //   dictDefaultMessage: "拖入需要上传的文件",
      //   uploadMultiple: false,
      //   parallelUploads: 10, //可修改手动触发一次上传的数量
      //   dictInvalidFileType: "支持的文件类型是office,.dwg,.dgn,.pdf,.rar,.zip,.7z",
      //   addRemoveLinks: true,
      //   dictRemoveFile: "移除文件",
      //   dictUploadCanceled: "取消",
      //   dictCancelUploadConfirmation: "取消上传该文件?",
      //   dictDefaultMessage: "<span class='bigger-150 bolder'><i class='icon-caret-right red'></i>拖动文件</span>上传\
      //           <span class='smaller-80 gre'>(或者点击上传)</span> <br /> \
      //           <i class='upload-icon icon-cloud-upload blue icon-3x'></i>",
      //   dictResponseError: "文件上传失败!",
      //   dictFileTooBig: "文件过大,上传失败!",
      //   init: function() {
      //     var myDropzone = this,
      //       submitButton = document.querySelector("#qr"),
      //       cancelButton = document.querySelector("#cancel");
      //     myDropzone.on('addedfile', function(file) {
      //       fileArr.push(file.upload.uuid);
      //       //解决点击时重复发送请求
      //       $(".dz-remove").each(function(index) {
      //         if (!$(".dz-remove:eq(" + index + ")").attr("id")) {
      //           $(".dz-remove:eq(" + index + ")").attr("id", fileArr[index]);
      //         }
      //       })
      //       //添加上传文件的过程，可再次弹出弹框，添加上传文件的信息
      //     });
      //     myDropzone.on('sending', function(data, xhr, formData) {
      //       //向后台发送该文件的参数
      //       // formData.append('watermark', jQuery('#info').val());
      //       formData.append('pid', projid);
      //       formData.append('prodlabel', $('#prodlabel').val());
      //       formData.append('prodprincipal', $('#prodprincipal').val());
      //     });
      //     myDropzone.on('success', function(files, response) {
      //       //文件上传成功之后的操作
      //     });
      //     myDropzone.on('error', function(files, response) {
      //       //文件上传失败后的操作
      //     });
      //     myDropzone.on('totaluploadprogress', function(progress, byte, bytes) {
      //       //progress为进度百分比
      //       $("#pro").text("上传进度：" + parseInt(progress) + "%");
      //       //计算上传速度和剩余时间
      //       var mm = 0;
      //       var byte = 0;
      //       var tt = setInterval(function() {
      //         mm++;
      //         var byte2 = bytes;
      //         var remain;
      //         var speed;
      //         var byteKb = byte / 1024;
      //         var bytesKb = bytes / 1024;
      //         var byteMb = byte / 1024 / 1024;
      //         var bytesMb = bytes / 1024 / 1024;
      //         if (byteKb <= 1024) {
      //           speed = (parseFloat(byte2 - byte) / (1024) / mm).toFixed(2) + " KB/s";
      //           remain = (byteKb - bytesKb) / parseFloat(speed);
      //         } else {
      //           speed = (parseFloat(byte2 - byte) / (1024 * 1024) / mm).toFixed(2) + " MB/s";
      //           remain = (byteMb - bytesMb) / parseFloat(speed);
      //         }
      //         $("#dropz #speed").text("上传速率：" + speed);
      //         $("#dropz #time").text("剩余时间" + arrive_timer_format(parseInt(remain)));
      //         if (bytes >= byte) {
      //           clearInterval(tt);
      //           if (byteKb <= 1024) {
      //             $("#dropz #speed").text("上传速率：0 KB/s");
      //           } else {
      //             $("#dropz #speed").text("上传速率：0 MB/s");
      //           }
      //           $("#dropz #time").text("剩余时间：0:00:00");
      //         }
      //       }, 1000);
      //     });
      //     submitButton.addEventListener('click', function() {
      //       //点击上传文件
      //       myDropzone.processQueue();
      //     });
      //     cancelButton.addEventListener('click', function() {
      //       //取消上传
      //       myDropzone.removeAllFiles();
      //       // if (myDropzone) {
      //       //   myDropzone.destroy();
      //       // }
      //     });
      //     myDropzone.on("complete", function(file) {
      //       myDropzone.removeFile(file);
      //     });
      //   }
      // });
      // //剩余时间格式转换：
      // function arrive_timer_format(s) {
      //   var t;
      //   if (s > -1) {
      //     var hour = Math.floor(s / 3600);
      //     var min = Math.floor(s / 60) % 60;
      //     var sec = s % 60;
      //     var day = parseInt(hour / 24);
      //     if (day > 0) {
      //       hour = hour - 24 * day;
      //       t = day + "day " + hour + ":";
      //     } else t = hour + ":";
      //     if (min < 10) {
      //       t += "0";
      //     }
      //     t += min + ":";
      //     if (sec < 10) {
      //       t += "0";
      //     }
      //     t += sec;
      //   }
      //   return t;
      // }

      // });

      </script>
</body>

</html>
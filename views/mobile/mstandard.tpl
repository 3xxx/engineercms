<!DOCTYPE html>

<title>查阅规范、图集、计算书</title>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css" />
<script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css" />
<script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
<link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
<script src="/static/js/tableExport.js"></script>
<script type="text/javascript" src="/static/js/moment.min.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/webuploader.css">
<script type="text/javascript" src="/static/js/webuploader.min.js"></script>
<script type="text/javascript" src="/static/js/jquery-ui.min.js"></script>
<!-- </head> -->
<!-- 悬浮按钮 -->
<script type="text/javascript" src="/static/float/js/float-module.min.js"></script>
<meta charset="UTF-8">
<!-- <title>Voxes - A Fresh Mobile Template</title> -->
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
/*覆盖bootstrap.min.css里的全局样式，用class+*来匹配，纯class不行，纯*号是全部改变，*/
/*button {margin: 10px;}*/
.fm-ul * {
  box-sizing: content-box
}

#modalTable .modal-header {
  cursor: move;
}

#modalTable1 .modal-header {
  cursor: move;
}

#modalTable2 .modal-header {
  cursor: move;
}

#modalTable3 .modal-header {
  cursor: move;
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
<!-- <div class="container-fill">{{template "/T.navbar.tpl" .}}</div> -->

<body>
  <!-- navbar top -->
  <div class="navbar-top">
    <div class="side-nav-panel-left"><a href="javascript:void(0)" data-activates="slide-out-left" class="side-nav-left"><i class="fa fa-bars"></i></a></div><!-- site brand -->
    <div class="site-brand"><a href="/index"></a>
      <h1>E</h1>
    </div><!-- end site brand -->
    <div class="side-nav-panel-right"><a href="" class="side-nav-right"><i class="fa fa-user">{{.Username}}</i></a></div>
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
  <div class="col-xs-12 col-sm-12">
    <h3>规范查询</h3>
    <div class="section">
      <div class="text-center">
        <h1> <i class="fa fa-terminal fa-2x pb-4 hide-sm"></i>
        </h1>
        <h1>搜索 {{.Length}}个 文件</h1>
        <br>
        <div class="col-xs-12 col-sm-12">
          <!-- <form >   form支持回车，但是不支持ajax，如何做到支持ajax？用ajaxform-->
          <div class="">
            <input type="text" class="form-control" placeholder="请输入关键字或编号进行搜索" name="name" autocomplete="off" size="30" id="name" onkeypress="getKey();">
            <span class="input-group-btn">
              <button class="btn btn-default" type="button" id="search2">
                <i class="glyphicon glyphicon-search"></i>
                Search!
              </button>
            </span>
          </div>
        </div>
      </div>
    </div>

    <div class="col-xs-12 col-sm-12">
      <br>
      <!--SWF在初始化的时候指定，在后面将展示-->
      <div id="uploader" class="wu-example col-xs-12">
        <!--用来存放文件信息-->
        <div id="thelist" class="uploader-list col-xs-12"></div>
        <div id="picker" style="margin:0 auto;width:90px;"><i class="glyphicon glyphicon-plus-sign"></i>选择文件</div>
        <button id="ctlBtn" class="btn btn-default" style="display:block;margin:0 auto"><i class="glyphicon glyphicon-upload"></i></button>
      </div>
      <br>
    </div>
  </div>
  <div class="col-xs-12 col-sm-12">
    <!-- 规范查询结果表 -->
    <div id="details" style="display:none">
      <h3>查询结果</h3>
      <table id="table"></table>
    </div>

  </div>
  <div id="footer">
    <div class="col-xs-12 col-sm-12">
      <br>
      <hr />
    </div>
    <div class="col-xs-12 col-sm-12 text-center">
      <h4>Copyright © 2016~2021 EngineerCMS</h4>
    </div>
  </div>

  <script>
  function actionFormatter(value, row, index) {
    return row.Category + ' ' + row.Number + '-' + row.Year
  }

  function index1(value, row, index) {
    // alert( "Data Loaded: " + index );
    return index + 1
  }

  function localDateFormatter(value) {
    return moment(value, 'YYYY-MM-DD').format('YYYY-MM-DD');
  }

  function setLink(value, row, index) {
    // return '<a href="' + row.Route + '" title="查阅" target="_blank"><i class="fa fa-paperclip"></i></a>';
    pdfUrl = '<a href="/v1/wx/standardpdf?file=' + row.Route + '" title="打开pdf" target="_blank"><i class="fa fa-file-pdf-o fa-lg text-danger"></i></a>';
    return pdfUrl;
  }

  function setLable(value, row, index) {
    return row.LiNumber + row.LibraryTitle
  }

  //项目列表
  $(function() {
    // 初始化【未接受】工作流表格
    $("#table").bootstrapTable({
      url: '/standard/search',
      method: 'get',
      // search: 'true',
      showRefresh: 'true',
      showToggle: 'true',
      // showColumns: 'true',
      // toolbar: '#btn_toolbar',
      pagination: 'true',
      sidePagination: "server",
      queryParamsType: '',
      //请求服务器数据时，你可以通过重写参数的方式添加一些额外的参数，例如 toolbar 中的参数 如果 queryParamsType = 'limit' ,返回参数必须包含
      // limit, offset, search, sort, order 否则, 需要包含: 
      // pageSize, pageNumber, searchText, sortName, sortOrder. 
      // 返回false将会终止请求。
      pageSize: 15,
      pageNumber: 1,
      pageList: [15, 20, 50, 'All'],
      // singleSelect: "true",
      clickToSelect: "true",
      queryParams: function queryParams(params) { //设置查询参数
        var param = {
          limit: params.pageSize, //每页多少条数据
          pageNo: params.pageNumber, // 页码
          searchText: params.searchText // $(".search .form-control").val(),
          // searchText: $("#name").val()
        };
        //搜索框功能
        //当查询条件中包含中文时，get请求默认会使用ISO-8859-1编码请求参数，在服务端需要对其解码
        // if (null != searchText) {
        //   try {
        //     searchText = new String(searchText.getBytes("ISO-8859-1"), "UTF-8");
        //   } catch (Exception e) {
        //     e.printStackTrace();
        //   }
        // }
        return param;
      },
      columns: [
        // {
        //   title: '选择',
        //   checkbox: 'true',
        //   width: '10',
        //   align: "center",
        //   valign: "middle"
        // },
        // {
        //   title: '序号',
        //   formatter: function(value, row, index) {
        //     return index + 1
        //   },
        //   align: "center",
        //   valign: "middle"
        // },
        {
          field: 'Number',
          title: '编号',
          align: "center",
          valign: "middle"
        },
        {
          field: 'Title',
          title: '名称',
          halign: "center",
          valign: "middle"
        },
        {
          field: 'Route',
          title: '链接',
          formatter: setLink,
          align: "center",
          valign: "middle"
        },
        {
          field: 'Uname',
          title: '上传者',
          align: "center",
          valign: "middle"
        },
        {
          field: 'LiNumber',
          title: '有效版本库',
          formatter: setLable,
          halign: "center",
          valign: "middle"
        }
      ]
    });
  });

  // <button class="btn btn-primary" id="export">导出excel</button>
  $(document).ready(function() {
    // 显示规范电子文件命名规则
    $("#about").click(function() {
      $('#modal').modal({
        show: true,
        backdrop: 'static'
      });
    })

    //显示和管理有效版本库
    $("#valid").click(function() {
      $("#details2").show();
      $('#table1').bootstrapTable('refresh', { url: '/standard/valid' });
    })

    $("#search2").click(function() { //这里应该用button的id来区分按钮的哪一个,因为本页有好几个button
      var title = $('#name').val();
      if (title.length >= 2) {
        $('#table').bootstrapTable('refresh', { url: '/standard/search' });
        // $.ajax({
        //   type: "get", 
        //   url: "/standard/search",
        //   data: { searchText: $("#name").val(),pageNo:1 },
        //   success: function(data, status) { //数据提交成功时返回数据
        $("#details").show();
        //     $('#table').bootstrapTable('append', data);
        //     $('#table').bootstrapTable('scrollTo', 'bottom');
        //   }
        // });
      } else {
        alert("请输入2个以上字符");
      }
    });
  });

  //规范表格增删改
  $(document).ready(function() {
    //添加规范
    $("#addButton").click(function() {
      $('#modalTable').modal({
        show: true,
        backdrop: 'static'
      });
    })
    //导入规范数据
    $("#importButton").click(function() {
      if (!{{.IsAdmin }}) {
        alert("非管理员，无权限！")
        return
      }
      $('#importstandardmodal').modal({
        show: true,
        backdrop: 'static'
      });
    })

    $("#editorButton").click(function() {
      if (!{{.IsAdmin }}) {
        alert("非管理员，无权限！")
        return
      }
      var selectRow = $('#table').bootstrapTable('getSelections');
      if (selectRow.length < 1) {
        alert("请先勾选！");
        return;
      }
      if (selectRow.length > 1) {
        alert("请不要勾选一个以上！");
        return;
      }
      $("input#cid").remove();
      var th1 = "<input id='cid' type='hidden' name='cid' value='" + selectRow[0].Id + "'/>"
      $(".modal-body").append(th1); //这里是否要换名字$("p").remove();
      $("#number1").val(selectRow[0].Number);
      $("#title1").val(selectRow[0].Title);
      $("#route1").val(selectRow[0].Route);
      $("#uname1").val(selectRow[0].Uname);
      $('#editorstandardmodal').modal({
        show: true,
        backdrop: 'static'
      });
    })

    $("#deleteButton").click(function() {
      if (!{{.IsAdmin }}) {
        alert("非管理员，无权限！")
        return
      }
      var selectRow = $('#table').bootstrapTable('getSelections');
      if (selectRow.length <= 0) {
        alert("请先勾选！");
        return false;
      }

      if (confirm("确定删除吗？一旦删除将无法恢复！")) {
        var title = $.map(selectRow, function(row) {
          return row.Title;
        })
        var ids = "";
        for (var i = 0; i < selectRow.length; i++) {
          if (i == 0) {
            ids = selectRow[i].Id;
          } else {
            ids = ids + "," + selectRow[i].Id;
          }
        }

        var ids1 = $.map(selectRow, function(row) {
          return row.id;
        })

        $.ajax({
          type: "post",
          url: "/standard/deletestandard",
          data: { ids: ids },
          success: function(data, status) {
            alert("删除“" + data + "”成功！(status:" + status + ".)");
            //删除已选数据
            $('#table').bootstrapTable('remove', {
              field: 'Title',
              values: title
            });
          }
        });
      }
    })
  })

  //有效版本库表格增删改
  $(document).ready(function() {
    $("#addButton1").click(function() {
      if (!{{.IsAdmin }}) {
        alert("非管理员，无权限！")
        return
      }
      $('#modalTable').modal({
        show: true,
        backdrop: 'static'
      });
    })

    $("#importButton1").click(function() {
      if (!{{.IsAdmin }}) {
        alert("非管理员，无权限！")
        return
      }
      $('#importvalidmodal').modal({
        show: true,
        backdrop: 'static'
      });
    })

    $("#editorButton1").click(function() {
      if (!{{.IsAdmin }}) {
        alert("非管理员，无权限！")
        return
      }
      var selectRow = $('#table1').bootstrapTable('getSelections');
      if (selectRow.length < 1) {
        alert("请先勾选！");
        return;
      }
      if (selectRow.length > 1) {
        alert("请不要勾选一个以上！");
        return;
      }
      $("input#cid").remove();
      var th1 = "<input id='cid' type='hidden' name='cid' value='" + selectRow[0].Id + "'/>"
      $(".modal-body").append(th1); //这里是否要换名字$("p").remove();
      $("#category2").val(selectRow[0].Category);
      $("#number2").val(selectRow[0].Number);
      $("#year2").val(selectRow[0].Year);
      $("#title2").val(selectRow[0].Title);
      $('#editorvalidmodal').modal({
        show: true,
        backdrop: 'static'
      });
    })

    $("#deleteButton1").click(function() {
      if (!{{.IsAdmin }}) {
        alert("非管理员，无权限！")
        return
      }
      var selectRow = $('#table1').bootstrapTable('getSelections');
      if (selectRow.length <= 0) {
        alert("请先勾选！");
        return false;
      }
      if (confirm("确定删除吗？一旦删除将无法恢复！")) {
        var title = $.map(selectRow, function(row) {
          return row.Title;
        })
        var ids = "";
        for (var i = 0; i < selectRow.length; i++) {
          if (i == 0) {
            ids = selectRow[i].Id;
          } else {
            ids = ids + "," + selectRow[i].Id;
          }
        }
        var ids1 = $.map(selectRow, function(row) {
          return row.id;
        })
        $.ajax({
          type: "post",
          url: "/standard/deletevalid",
          data: { ids: ids },
          success: function(data, status) {
            alert("删除“" + data + "”成功！(status:" + status + ".)");
            //删除已选数据
            $('#table1').bootstrapTable('remove', {
              field: 'Title',
              values: title
            });
          }
        });
      }
    })
  })

  function updatestandard() {
    if (!{{.IsAdmin }}) {
      alert("非管理员，无权限！")
      return
    }
    // var radio =$("input[type='radio']:checked").val();
    var number1 = $('#number1').val();
    var title1 = $('#title1').val();
    var route1 = $('#route1').val();
    var uname1 = $('#uname1').val();
    var cid = $('#cid').val();
    // $('#myModal').on('hide.bs.modal', function () {  
    if (number1) {
      $.ajax({
        type: "post",
        url: "/standard/updatestandard",
        data: { cid: cid, number: number1, title: title1, route: route1, uname: uname1 },
        success: function(data, status) {
          alert("修改“" + data + "”成功！(status:" + status + ".)");
        }
      });
    }
    // $(function(){$('#myModal').modal('hide')});
    $('#editorstandardmodal').modal('hide');
    $('#table').bootstrapTable('refresh', { url: '/getstandard' });
    // "/category/modifyfrm?cid="+cid
    // window.location.reload();//刷新页面
  }

  // 文件上传
  jQuery(function() {
    var $ = jQuery,
      $list = $('#thelist'),
      $btn = $('#ctlBtn'),
      state = 'pending',
      uploader;
    uploader = WebUploader.create({
      // 不压缩image
      resize: false,
      // swf文件路径
      swf: '/static/fex-team-webuploader/dist/Uploader.swf',
      // 文件接收服务端。
      server: '/standard/standard_one_addbaidu',
      // 选择文件的按钮。可选。
      // 内部根据当前运行是创建，可能是input元素，也可能是flash.
      pick: '#picker'
    });

    // 当有文件添加进来的时候
    uploader.on('fileQueued', function(file) {
      $list.append('<div id="' + file.id + '" class="item">' +
        '<h4 class="info">' + file.name + '</h4>' +
        '<p class="state">等待上传...</p>' +
        '</div>');
    });
    //当某个文件的分块在发送前触发，主要用来询问是否要添加附带参数，
    //大文件在开起分片上传的前提下此事件可能会触发多次。
    // uploader.on( 'fileQueued', function( file ) {
    // do some things.
    // });

    // uploader.on( 'startUpload', function() {
    //    var tnumber = $('#tnumber').val();
    //    var title = $('#title').val();
    //    var categoryid = $('#categoryid').val();
    //    var category = $('#category').val();
    //    var html = ue.getContent();
    //      uploader.option('formData', {
    //        "tnumber":tnumber,
    //        "title":title,
    //        "categoryid":categoryid,
    //        "category":category,
    //        'content':html,
    //      });        
    //    });

    // 文件上传过程中创建进度条实时显示。
    uploader.on('uploadProgress', function(file, percentage) {
      var $li = $('#' + file.id),
        $percent = $li.find('.progress .progress-bar');
      // 避免重复创建
      if (!$percent.length) {
        $percent = $('<div class="progress progress-striped active">' +
          '<div class="progress-bar" role="progressbar" style="width: 0%">' +
          '</div>' +
          '</div>').appendTo($li).find('.progress-bar');
      }

      $li.find('p.state').text('上传中');

      $percent.css('width', percentage * 100 + '%');
    });

    uploader.on('uploadSuccess', function(file) {
      $('#' + file.id).find('p.state').text('已上传');
    });

    uploader.on('uploadError', function(file) {
      $('#' + file.id).find('p.state').text('上传出错');
    });

    uploader.on('uploadComplete', function(file) {
      $('#' + file.id).find('.progress').fadeOut();
    });

    uploader.on('all', function(type) {
      if (type === 'startUpload') {
        state = 'uploading';
      } else if (type === 'stopUpload') {
        state = 'paused';
      } else if (type === 'uploadFinished') {
        state = 'done';
      }

      if (state === 'uploading') {
        $btn.text('暂停上传');
      } else {
        $btn.text('开始上传');
      }
    });

    $btn.on('click', function() {
      if (state === 'uploading') {
        uploader.stop();
      } else {
        uploader.upload();
      }
    });
  });

  function getKey() {
    var title = $('#name').val();
    if (event.keyCode == 13) {
      if (title.length >= 2) {
        $('#table').bootstrapTable('refresh', { url: '/standard/search' });
        // $.ajax({
        //   type: "get",
        //   url: "/standard/search",
        //   data: { searchText: $("#name").val(),pageNo:1 },
        //   success: function(data, status) {
        $("#details").show();
        //     $('#table').bootstrapTable('append', data.rows);
        //     $('#table').bootstrapTable('scrollTo', 'bottom');
        //   }
        // });
      } else {
        alert("请输入2个以上字符");
      }
    }
  }
  </script>
  <!-- scripts -->
  <!-- <script src="/static/voxes/js/jquery.min.js"></script> -->
  <script src="/static/voxes/js/materialize.min.js"></script>
  <script src="/static/voxes/js/owl.carousel.min.js"></script>
  <!-- <script src="/static/voxes/js/jquery.filterizr.js"></script> -->
  <!-- <script src="/static/voxes/js/jquery.magnific-popup.min.js"></script> -->
  <!-- <script src="/static/voxes/js/portfolio.js"></script> -->
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
  </script>
</body>

</html>
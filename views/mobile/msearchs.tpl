<!-- 显示特定项目的搜索页面 -->
<!DOCTYPE html>
<html>

<head>
  <meta charset="UTF-8">
  <title>EngineerCMS</title>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script src="/static/js/bootstrap-treeview.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css" />
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
  <script src="/static/js/tableExport.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <script type="text/javascript" src="/static/js/jquery-ui.min.js"></script>
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

  <!-- 悬浮按钮 -->
  <script type="text/javascript" src="/static/float/js/float-module.min.js"></script>

  <style type="text/css">
    /*覆盖bootstrap.min.css里的全局样式，用class+*来匹配，纯class不行，纯*号是全部改变，*/
    html {
      /*font-family: GillSans, Calibri, Trebuchet, sans-serif;*/
      font-family: "Consolas", "Microsoft Yahei", Arial, monospace;
    }

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

<body>
  <!-- navbar top -->
  <div class="navbar-top">
    <div class="side-nav-panel-left"><a href="#" data-activates="slide-out-left" class="side-nav-left"><i class="fa fa-bars"></i></a></div><!-- site brand -->
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
      <li><a href="/v1/wx/photoswipe"><i class="fa fa-photo"></i>相册</a></li>
      <li><a href="/v1/flv/flvlist"><i class="fa fa-video-camera"></i>视频</a></li>
      <!-- <li><a href="pricing.html"><i class="fa fa-dollar"></i>公告</a></li> -->
      <!-- <li><a href="error404.html"><i class="fa fa-hourglass-half"></i>404</a></li> -->
      <li><a href="/mindoc"><i class="fa fa-book"></i>Book</a></li>
      <!-- <li><a href="about-us.html"><i class="fa fa-user"></i>About Us</a></li> -->
      <!-- <li><a href="contact.html"><i class="fa fa-envelope-o"></i>Contact Us</a></li> -->
      <li><a href="/login"><i class="fa fa-sign-in"></i>Login</a></li>
      <li><a href="" onclick="return logout();"><i class="fa fa-sign-out"></i>Logout</a></li>
      <li><a href="/regist"><i class="fa fa-user-plus"></i>Register</a></li>
      <li><a href="" class="search-trigger"><i class="fa fa-search"></i>Search</a></li>
    </ul>
  </div><!-- end side nav left-->
  <div class="header-search">
    <form role="search" method="POST" class="search-form" action="#">
      <div class="search-group"><input type="text" class="input-search" placeholder="Search ..."><button class="search-submit" type="submit"><i class="fa fa-search"></i></button></div>
    </form>
  </div><!-- . header-search -->
  <div class="col-xs-12 col-sm-6">
    <h3>成果列表</h3>
    <div class="section">
       <table id="table1"></table>
    </div>
  </div>
  <script type="text/javascript">
  $(function() {
    // 初始化【未接受】工作流表格
    $("#table1").bootstrapTable({
      url: '/v1/wx/searchprojectproductdata?keyword={{.Key}}&productid={{.Pid}}',
      method: 'get',
      showRefresh: 'true',
      showToggle: 'true',
      showColumns: 'true',
      pagination: 'true',
      sidePagination: "server",
      queryParamsType: '',
      //请求服务器数据时，你可以通过重写参数的方式添加一些额外的参数，例如 toolbar 中的参数 如果 queryParamsType = 'limit' ,返回参数必须包含
      // limit, offset, search, sort, order 否则, 需要包含: 
      // pageSize, pageNumber, searchText, sortName, sortOrder. 
      // 返回false将会终止请求。
      pageSize: 15,
      pageNumber: 1,
      pageList: [15, 50, 100],
      uniqueId: "id",
      clickToSelect: "true",
      queryParams: function queryParams(params) { //设置查询参数
        var param = {
          limit: params.pageSize, //每页多少条数据
          pageNo: params.pageNumber, // 页码
          searchText: $(".search .form-control").val()
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
        {
          // field: 'Number',
          title: '序号',
          formatter: function(value, row, index) {
            return index + 1
          },
          align: "center",
          valign: "middle"
        },
        // {
        //   field: 'Code',
        //   title: '编号',
        //   halign: "center",
        //   align: "left",
        //   valign: "middle"
        // },
        {
          field: 'Title',
          title: '名称',
          // formatter:setTitle,
          halign: "center",
          align: "left",
          valign: "middle"
        },
        // {
        //   field: 'Label',
        //   title: '标签',
        //   formatter: setLable,
        //   align: "center",
        //   valign: "middle"
        // },
        // {
        //   field: 'Principal',
        //   title: '设计',
        //   align: "center",
        //   valign: "middle"
        // },
        {
          field: 'Articlecontent',
          title: '文章',
          formatter: setArticle,
          events: actionEvents,
          align: "center",
          valign: "middle"
        },
        {
          field: 'Attachmentlink',
          title: '附件',
          formatter: setAttachment,
          events: actionEvents,
          align: "center",
          valign: "middle"
        },
        {
          field: 'Pdflink',
          title: 'PDF',
          formatter: setPdf,
          events: actionEvents,
          align: "center",
          valign: "middle"
        }
      ]
    });
  });
  // 改变点击行颜色
  $(function() {
    $("#table0").on("click-row.bs.table", function(e, row, ele) {
      $(".info").removeClass("info");
      $(ele).addClass("info");
      rowid = row.Id; //全局变量
    });
  });

  // $(document).ready(function(){
  $("#search").click(function() { //这里应该用button的id来区分按钮的哪一个,因为本页有好几个button
    var radio = $("input[type='radio']:checked").val();
    $.ajax({
      type: "post", //这里是否一定要用post，是的，因为get会缓存？？
      url: "/v1/wx/searchprojectproductdata",
      data: { keyword: $("#keyword").val(), radiostring: radio },
      success: function(data, status) { //数据提交成功时返回数据
        //显示结果表
        $("#rowtitle").html("搜寻结果");
        $("#details").show();
        $('#table1').bootstrapTable('append', data);
        $('#table1').bootstrapTable('scrollTo', 'bottom');
      }
    });
  });
  // });

  function getKey() {
    if (event.keyCode == 13) {
      var radio = $("input[type='radio']:checked").val();
      $.ajax({
        type: "post", //这里是否一定要用post，是的，因为get会缓存？？
        url: "/v1/wx/searchprojectproductdata",
        data: { keyword: $("#keyword").val(), radiostring: radio },
        success: function(data, status) { //数据提交成功时返回数据

          //显示结果表
          $("#rowtitle").html("搜寻结果");
          $("#details").show();
          $('#table1').bootstrapTable('append', data);
          $('#table1').bootstrapTable('scrollTo', 'bottom');
        }
      });
    }
  }

  // 成果添加到购物车
  $("#cartButton").click(function() {
    var selectRow = $('#table1').bootstrapTable('getSelections');
    if (selectRow.length <= 0) {
      alert("请先勾选成果！");
      return false;
    }
    if (selectRow[0].Attachmentlink[0]) { //||selectRow[0].Pdflink[0].Link||selectRow[0].Articlecontent[0].Link)
      var site = /http:\/\/.*?\//.exec(selectRow[0].Attachmentlink[0].Link); //非贪婪模式 
    }
    if (selectRow[0].Articlecontent[0]) {
      var site = /http:\/\/.*?\//.exec(selectRow[0].Articlecontent[0].Link); //非贪婪模式 
    }
    if (selectRow[0].Pdflink[0]) {
      var site = /http:\/\/.*?\//.exec(selectRow[0].Pdflink[0].Link); //非贪婪模式 
    }
    if (site) {
      alert("同步成果不允许！");
      return;
    }
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
    $.ajax({
      type: "post",
      url: "/v1/cart/createproductcart",
      data: { ids: ids },
      success: function(data, status) {
        if (data.code == "ERROR") {
          alert(data.msg);
        } else {
          alert("添加“" + data.data[0].Title + "”购物车成功！(status:" + status + "！)");
        }
        // $.toast({
        //   type: TYPES[1],
        //   title: TITLES['info'],
        //   subtitle: '11 mins ago',
        //   content: CONTENT['info'],
        //   delay: 5000
        // });
      }
    });
  })

  function index1(value, row, index) {
    // alert( "Data Loaded: " + index );
    return index + 1
  }

  function localDateFormatter(value) {
    return moment(value, 'YYYY-MM-DD').format('YYYY-MM-DD');
  }
  //关联
  function RelevFormatter(value) {
    if (value) {
      if (value.length == 1) { //'<a href="/project/product/article/'
        var array = value[0].Relevancy.split(",")
        var relevarray = new Array()
        for (i = 0; i < array.length; i++) {
          relevarray[i] = array[i];
        }
        return relevarray.join(",");
        // articleUrl= '<a href="'+value[0].Link+'/'+value[0].Id+'" title="查看" target="_blank"><i class="fa fa-file-text-o"></i></a>';
        // return articleUrl;
      } else if (value.length == 0) {

      } else if (value.length > 1) {
        var relevarray = new Array()
        for (i = 0; i < value.length; i++) {
          relevarray[i] = value[i].Relevancy;
        }
        return relevarray.join(",");
        // articleUrl= "<a class='article' href='javascript:void(0)' title='查看文章列表'><i class='fa fa-list-ol'></i></a>";
        // return articleUrl;
      }
    }
  }

  function setCode(value, row, index) {
    return "<a href='/project/product/attachment/" + row.Id + "'>" + value + "</a>";
  }

  function setLable(value, row, index) {
    // alert(value);
    if (value) { //注意这里如果value未定义则出错，一定要加这个判断。
      var array = value.split(",")
      var labelarray = new Array()
      for (i = 0; i < array.length; i++) {
        labelarray[i] = "<a href='/project/product/keysearch?keyword=" + array[i] + "'>" + array[i] + "</a>";
      }
      return labelarray.join(",");
    }
  }

  function setTitle(value, row, index) {
    return "<a href='/project/product/" + row.Id + "'>" + value + "</a>";
  }

  function setArticle(value, row, index) {
    // return '<a class="article" href="javascript:void(0)" title="article"><i class="fa fa-file-text-o"></i></a>';
    if (value) {
      if (value.length == 1) { //'<a href="/project/product/article/'
        articleUrl = '<a href="' + value[0].Link + '/' + value[0].Id + '" title="查看" target="_blank"><i class="fa fa-file-text-o"></i></a>';
        return articleUrl;
      } else if (value.length == 0) {

      } else if (value.length > 1) {
        articleUrl = "<a class='article' href='javascript:void(0)' title='查看文章列表'><i class='fa fa-list-ol'></i></a>";
        return articleUrl;
      }
    }
  }

  function setAttachment(value, row, index) {
    if (value) {
      if (value.length == 1) {
        attachUrl = '<a href="' + value[0].Link + '/' + value[0].Title + '" title="下载" target="_blank"><i class="fa fa-paperclip"></i></a>';
        return attachUrl;
      } else if (value.length == 0) {

      } else if (value.length > 1) {
        attachUrl = "<a class='attachment' href='javascript:void(0)' title='查看附件列表'><i class='fa fa-list-ol'></i></a>";
        return attachUrl;
      }
    }
  }

  function setPdf(value, row, index) {
    if (value) {
      if (value.length == 1) {
        pdfUrl = '<a href="' + value[0].Link + '/' + value[0].Title + '" title="打开pdf" target="_blank"><i class="fa fa-file-pdf-o"></i></a>';
        return pdfUrl;
      } else if (value.length == 0) {

      } else if (value.length > 1) {
        pdfUrl = "<a class='pdf' href='javascript:void(0)' title='查看pdf列表'><i class='fa fa-list-ol'></i></a>";
        return pdfUrl;
      }
    }
  }

  window.actionEvents = {
    'click .article': function(e, value, row, index) {
      var site = /http:\/\/.*?\//.exec(value[1].Link); //非贪婪模式 
      if (site) {
        $('#articles').bootstrapTable('refresh', { url: '/project/product/syncharticles?site=' + site + '&id=' + row.Id });
      } else {
        $('#articles').bootstrapTable('refresh', { url: '/project/product/articles/' + row.Id });
      }
      $('#modalarticle').modal({
        show: true,
        backdrop: 'static'
      });
    },
    'click .attachment': function(e, value, row, index) {
      // for(var i=0;i<value.length;i++)
      // alert(value[i].Link);
      // var ret=/http:(.*)\:/.exec(value[i].Link);//http://127.0.0.1:
      var site = /http:\/\/.*?\//.exec(value[1].Link); //非贪婪模式 
      if (site) { //跨域
        // alert("1");
        // $.getJSON(ret+'project/product/attachment/'+row.Id,function(){
        // $('#attachs').bootstrapTable('load', randomData());
        // })
        $('#attachs').bootstrapTable('refresh', { url: '/project/product/synchattachment?site=' + site + '&id=' + row.Id });
        // $('#attachs').bootstrapTable('refresh', {url:site+'project/product/attachment/'+row.Id});
      } else {
        // alert("2");
        $('#attachs').bootstrapTable('refresh', { url: '/project/product/attachment/' + row.Id });
      }
      $('#modalattach').modal({
        show: true,
        backdrop: 'static'
      });
    },

    'click .pdf': function(e, value, row, index) {
      var site = /http:\/\/.*?\//.exec(value[1].Link); //非贪婪模式 
      if (site) { //跨域
        $('#pdfs').bootstrapTable('refresh', { url: '/project/product/synchpdf?site=' + site + '&id=' + row.Id });
      } else {
        $('#pdfs').bootstrapTable('refresh', { url: '/project/product/pdf/' + row.Id });
      }
      $('#modalpdf').modal({
        show: true,
        backdrop: 'static'
      });
    },
  };

  //最后面弹出文章列表中用的_根据上面的click，弹出模态框，给模态框中的链接赋值
  function setArticlecontent(value, row, index) {
    articleUrl = '<a href="' + value + '" title="下载" target="_blank"><i class="fa fa-file-text-o"></i></a>';
    return articleUrl;
  }
  //最后面弹出附件列表中用的
  function setAttachlink(value, row, index) {
    attachUrl = '<a href="' + value + '" title="下载" target="_blank"><i class="fa fa-paperclip"></i></a>';
    return attachUrl;
  }
  //最后面弹出pdf列表中用的
  function setPdflink(value, row, index) {
    pdfUrl = '<a href="' + value + '" title="下载" target="_blank"><i class="fa fa-file-pdf-o"></i></a>';
    return pdfUrl;
  }
  </script>
  <!-- 文章列表 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalarticle">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">文章列表</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <!-- <div id="pdfs" style="display:none"> -->
              <!-- <h3>工程目录分级</h3> -->
              <table id="articles" data-toggle="table" data-page-size="5" data-page-list="[5, 25, 50, All]" data-unique-id="id" data-pagination="true" data-side-pagination="client" data-click-to-select="true">
                <thead>
                  <tr>
                    <th data-width="10" data-checkbox="true"></th>
                    <th data-formatter="index1">#</th>
                    <th data-field="Title">名称</th>
                    <th data-field="Subtext">副标题</th>
                    <th data-field="Link" data-formatter="setArticlecontent">查看</th>
                    <th data-field="Created" data-formatter="localDateFormatter">建立时间</th>
                    <th data-field="Updated" data-formatter="localDateFormatter">修改时间</th>
                  </tr>
                </thead>
              </table>
              <!-- </div> -->
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <!-- 除了**pdf**之外的附件列表 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalattach">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">非PDF附件列表</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <!-- <div id="pdfs" style="display:none"> -->
              <!-- <h3>工程目录分级</h3> -->
              <table id="attachs" data-toggle="table" data-page-size="5" data-page-list="[5, 25, 50, All]" data-unique-id="id" data-pagination="true" data-side-pagination="client" data-click-to-select="true">
                <thead>
                  <tr>
                    <th data-width="10" data-checkbox="true"></th>
                    <th data-formatter="index1">#</th>
                    <th data-field="Title">名称</th>
                    <th data-field="FileSize">大小</th>
                    <th data-field="Link" data-formatter="setAttachlink">下载</th>
                    <th data-field="Created" data-formatter="localDateFormatter">建立时间</th>
                    <th data-field="Updated" data-formatter="localDateFormatter">修改时间</th>
                  </tr>
                </thead>
              </table>
              <!-- </div> -->
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <!-- pdf附件列表 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalpdf">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">pdf附件列表</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <!-- <div id="pdfs" style="display:none"> -->
              <!-- <h3>工程目录分级</h3> -->
              <table id="pdfs" data-toggle="table" data-page-size="5" data-page-list="[5, 25, 50, All]" data-unique-id="id" data-pagination="true" data-side-pagination="client" data-click-to-select="true">
                <thead>
                  <tr>
                    <th data-width="10" data-checkbox="true"></th>
                    <th data-formatter="index1">#</th>
                    <th data-field="Title">名称</th>
                    <th data-field="FileSize">大小</th>
                    <th data-field="Link" data-formatter="setPdflink">下载</th>
                    <th data-field="Created" data-formatter="localDateFormatter">建立时间</th>
                    <th data-field="Updated" data-formatter="localDateFormatter">修改时间</th>
                  </tr>
                </thead>
              </table>
              <!-- </div> -->
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
          </div>
        </div>
      </div>
    </div>
  </div>
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
  }
  </script>
</body>

</html>
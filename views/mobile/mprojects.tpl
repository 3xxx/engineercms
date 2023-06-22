<!-- 项目列表页 后端分页-->
<!DOCTYPE html>
<head>
  <meta charset="UTF-8">
  <title>项目列表-EngiCMS</title>
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
<!-- <div class="container-fill">{{template "/T.navbar.tpl" .}}</div> -->

<body>
    <!-- navbar top -->
  <div class="navbar-top">
    <div class="side-nav-panel-left"><a href="javascript:void(0)" data-activates="slide-out-left" class="side-nav-left"><i class="fa fa-bars"></i></a></div><!-- site brand -->
    <div class="site-brand"><a href="/index">
      <h1>E</h1>
    </div><!-- end site brand -->
    <div class="side-nav-panel-right"><a href="" class="side-nav-right"><i class="fa fa-user">{{.Username}}</i></a></div>
  </div><!-- end navbar top -->
  <!-- side nav left-->
  <div class="side-nav-panel-left">
    <ul id="slide-out-left" class="side-nav side-nav-panel collapsible">
      <li class="li-top"><a href="/index"><i class="fa fa-home"></i>主页</a></li>
      <li><a href="/project"><i class="fa fa-product-hunt"></i>项目</a></li>
      <li><a href="/standard"><i class="fa fa-scribd"></i>标准</a></li>
      <li><a href="/v1/wx/photoswipe"><i class="fa fa-photo"></i>相册</a></li>
      <li><a href="/v1/flv/flvlist"><i class="fa fa-video-camera"></i>视频</a></li>
      <!-- <li><a href="pricing.html"><i class="fa fa-dollar"></i>公告</a></li> -->
      <!-- <li><a href="error404.html"><i class="fa fa-hourglass-half"></i>404</a></li> -->
      <li><a href="/mindoc"><i class="fa fa-book"></i>书籍</a></li>
      <!-- <li><a href="about-us.html"><i class="fa fa-user"></i>About Us</a></li> -->
      <!-- <li><a href="contact.html"><i class="fa fa-envelope-o"></i>Contact Us</a></li> -->
      <li><a href="/login"><i class="fa fa-sign-in"></i>登录</a></li>
      <li><a href="" onclick="return logout();"><i class="fa fa-sign-out"></i>退出</a></li>
      <li><a href="/regist"><i class="fa fa-user-plus"></i>注册</a></li>
      <li><a class="search-trigger"><i class="fa fa-search"></i>搜索</a></li>
    </ul>
  </div><!-- end side nav left-->
  <div class="header-search" onkeypress="getKey();">
    <div class="search-form">
      <div class="search-group"><input type="text" class="input-search" placeholder="Search ..." id="keyword" ><button class="search-submit" onclick="searchAllProjectsProduct()"><i class="fa fa-search"></i></button></div>
    </div>
  </div><!-- . header-search -->

  <div class="col-xs-12 col-sm-6">
    <h3>项目列表</h3>
    <div class="section">
    <table id="table0"></table>
    </div>
    <script type="text/javascript">
    //项目列表
    $(function() {
      // 初始化【未接受】工作流表格
      $("#table0").bootstrapTable({
        url: '/project/getprojects',
        method: 'get',
        // search: 'true',
        showRefresh: 'true',
        showToggle: 'true',
        // showColumns: 'true',
        toolbar: '#toolbar1',
        pagination: 'true',
        sidePagination: "server",
        queryParamsType: '',
        //请求服务器数据时，你可以通过重写参数的方式添加一些额外的参数，例如 toolbar 中的参数 如果 queryParamsType = 'limit' ,返回参数必须包含
        // limit, offset, search, sort, order 否则, 需要包含: 
        // pageSize, pageNumber, searchText, sortName, sortOrder. 
        // 返回false将会终止请求。
        pageSize: 15,
        pageNumber: 1,
        pageList: [15, 30, 50, 'All'],
        singleSelect: "true",
        clickToSelect: "true",
        queryParams: function queryParams(params) { //设置查询参数
          var param = {
            limit: params.pageSize, //每页多少条数据
            pageNo: params.pageNumber, // 页码
            searchText: $(".search .form-control").val()
          };
          return param;
        },
        columns: [
          {
            title: '选择',
            radio: 'true',
            width: '0',
            align: "center",
            valign: "middle"
          },
          {
            title: '序号',
            formatter: function(value, row, index) {
              return index + 1
            },
            align: "center",
            valign: "middle"
          },
          {
            field: 'Code',
            title: '编号',
            formatter: setCode,
            align: "center",
            valign: "middle"
          },
          {
            field: 'Title',
            title: '名称',
            formatter: setTitle,
            align: "center",
            valign: "middle"
          },
          // {
          //   field: 'Label',
          //   title: '标签',
          //   formatter: setLable,
          //   align: "center",
          //   valign: "middle"
          // },
          {
            field: 'Principal',
            title: '负责人',
            align: "center",
            valign: "middle"
          }
        ]
      });
    });

    //项目模板中的项目列表
    $(function() {
      // 初始化【未接受】工作流表格
      $("#table2").bootstrapTable({
        url: '/project/getprojects',
        // method: 'get',
        search: 'true',
        showRefresh: 'true',
        showToggle: 'true',
        showColumns: 'true',
        // toolbar:'#toolbar1',
        pagination: 'true',
        sidePagination: "server",
        queryParamsType: '',
        //请求服务器数据时，你可以通过重写参数的方式添加一些额外的参数，例如 toolbar 中的参数 如果 queryParamsType = 'limit' ,返回参数必须包含
        // limit, offset, search, sort, order 否则, 需要包含: 
        // pageSize, pageNumber, searchText, sortName, sortOrder. 
        // 返回false将会终止请求。
        pageSize: 10,
        pageNumber: 1,
        pageList: [10, 20, 50, 100],
        singleSelect: "true",
        clickToSelect: "true",
        queryParams: function queryParams(params) { //设置查询参数
          var param = {
            limit: params.pageSize, //每页多少条数据
            pageNo: params.pageNumber, // 页码
            searchText: $(".search .form-control").val()
          };
          //搜索框功能
          return param;
        },
        columns: [{
            title: '选择',
            radio: 'true',
            width: '10',
            align: "center",
            valign: "middle"
          },
          {
            title: '序号',
            formatter: function(value, row, index) {
              return index + 1
            },
            align: "center",
            valign: "middle"
          },
          {
            field: 'Code',
            title: '编号',
            formatter: setCode,
            align: "center",
            valign: "middle"
          },
          {
            field: 'Title',
            title: '名称',
            formatter: setTitle,
            align: "center",
            valign: "middle"
          }
        ]
      });
    });

    function index1(value, row, index) {
      // alert( "Data Loaded: " + index );
      return index + 1
    }

    function localDateFormatter(value) {
      return moment(value, 'YYYY-MM-DD').format('YYYY-MM-DD');
    }

    function setCode(value, row, index) {
      return "<a href='/project/" + row.Id + "'>" + value + "</a>";
    }

    function setTitle(value, row, index) {
      return "<a href='/project/allproducts/" + row.Id + "'>" + value + "</a>";
    }

    function setLable(value, row, index) {
      if (value) {
        array = value.split(",")
        var labelarray = new Array()
        for (i = 0; i < array.length; i++) {
          labelarray[i] = "<a href='/project/keysearch?keyword=" + array[i] + "'>" + array[i] + "</a>";
        }
        return labelarray.join(",");
      }
    }
    // 改变点击行颜色
    $(function() {
      // $("#table").bootstrapTable('destroy').bootstrapTable({
      //     columns:columns,
      //     data:json
      // });
      $("#table0").on("click-row.bs.table", function(e, row, ele) {
        $(".info").removeClass("info");
        $(ele).addClass("info");
        // rowid=row.Id//全局变量
        // $('#table1').bootstrapTable('refresh', {url:'/admin/category/'+row.Id});
      });
      // $("#get").click(function(){
      //     alert("商品名称：" + getContent().TuanGouName);
      // })
    });
   
    $("#addButton").click(function() {
      if (!{{.IsLogin }}) {
        alert("权限不够！");
        return;
      }
      $("label#info").remove();
      $("#saveproj").removeClass("disabled")
      $('#modalTable').modal({
        show: true,
        backdrop: 'static'
      });
    })

    function editorProjButton() {
      if (!{{.IsLogin }}) {
        alert("未登陆！");
        return;
      }
      var selectRow = $('#table0').bootstrapTable('getSelections');
      if (selectRow.length < 1) {
        alert("请先勾选项目！");
        return;
      }
      if (selectRow.length > 1) {
        alert("请不要勾选一个以上！");
        return;
      }
      // alert({{.IsAdmin}})
      $.ajax({
        type: "get",
        url: "/v1/project/projectuserrole",
        data: { pid: selectRow[0].Id },
        success: function(data, status) {
          // alert(data.userrole)
          if (data.userrole == "isme" || {{.IsAdmin }}) {
            showprojecteditor()
          } else {
            alert("非管理员、非本人，无修改权限")
          }
        },
        complete: function(data) {
          //请求完成的处理
        },
        error: function(data) {
          //请求出错处理
        }
      });
    }

    function showprojecteditor() {
      var selectRow = $('#table0').bootstrapTable('getSelections');
      $("input#cid").remove();
      var th1 = "<input id='cid' type='hidden' name='cid' value='" + selectRow[0].Id + "'/>";
      $(".modal-body").append(th1); //这里是否要换名字$("p").remove();
      $("#projcode1").val(selectRow[0].Code);
      $("#projname1").val(selectRow[0].Title);
      $("#projlabel1").val(selectRow[0].Label);
      $("#projprincipal1").val(selectRow[0].Principal);
      $('#modalTable1').modal({
        show: true,
        backdrop: 'static'
      });
    }
    // 编辑项目目录
    function editorProjTreeButton() {
      if (!{{.IsLogin }}) {
        alert("未登陆！");
        return;
      }
      var selectRow = $('#table0').bootstrapTable('getSelections');
      if (selectRow.length < 1) {
        alert("请先勾选项目！");
        return;
      }
      if (selectRow.length > 1) {
        alert("请不要勾选一个以上！");
        return;
      }
      // alert({{.IsAdmin}})
      $.ajax({
        type: "get",
        url: "/v1/project/projectuserrole",
        data: { pid: selectRow[0].Id },
        success: function(data, status) {
          // alert(data.userrole)
          if (data.userrole == "isme" || {{.IsAdmin }}) {
            //跳转到新页面
            window.location.href = "/v1/project/userprojecteditortree?pid=" + selectRow[0].Id;
          } else {
            alert("非管理员、非本人，无修改权限")
          }
        },
        complete: function(data) {
          //请求完成的处理
        },
        error: function(data) {
          //请求出错处理
        }
      });
    }
    // 设置项目权限
    function ProjPermissionButton() {
      if (!{{.IsLogin }}) {
        alert("未登陆！");
        return;
      }
      var selectRow = $('#table0').bootstrapTable('getSelections');
      if (selectRow.length < 1) {
        alert("请先勾选项目！");
        return;
      }
      if (selectRow.length > 1) {
        alert("请不要勾选一个以上！");
        return;
      }
      // alert({{.IsAdmin}})
      $.ajax({
        type: "get",
        url: "/v1/project/projectuserrole",
        data: { pid: selectRow[0].Id },
        success: function(data, status) {
          // alert(data.userrole)
          if (data.userrole == "isme" || {{.IsAdmin }}) {
            //跳转到新页面
            window.location.href = "/v1/project/userprojectpermission?pid=" + selectRow[0].Id;
          } else {
            alert("非管理员、非本人，无修改权限")
          }
        },
        complete: function(data) {
          //请求完成的处理
        },
        error: function(data) {
          //请求出错处理
        }
      });
    }
    // 删除项目
    $("#deleteButton").click(function() {
      if ({{.role }} != 1) {
        alert("权限不够！");
        return;
      }

      var selectRow = $('#table0').bootstrapTable('getSelections');
      if (selectRow.length <= 0) {
        alert("请先勾选项目！");
        return false;
      }
      if (confirm("确定删除项目吗？第一次提示！")) {} else {
        return false;
      }
      if (confirm("确定删除项目吗？第二次提示！")) {} else {
        return false;
      }
      if (confirm("确定删除项目吗？一旦删除将无法恢复！")) {
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
        //删除前端表格用的
        var ids2 = $.map($('#table0').bootstrapTable('getSelections'), function(row) {
          return row.Id;
        });
        $.ajax({
          type: "post",
          url: "/project/deleteproject",
          data: { ids: ids },
          success: function(data, status) {
            alert("删除“" + data + "”成功！(status:" + status + ".)");
            //删除已选数据
            $('#table0').bootstrapTable('remove', {
              field: 'Id',
              values: ids2
            });
          }
        });
      }
    })
    // })

    /*数据json*/
    var json1 = [{ "Id": "1", "Title": "规划", "Code": "A", "Grade": "1" },
      { "Id": "7", "Title": "可研", "Code": "B", "Grade": "1" },
      { "Id": "2", "Title": "报告", "Code": "B", "Grade": "2" },
      { "Id": "3", "Title": "图纸", "Code": "T", "Grade": "2" },
      { "Id": "4", "Title": "水工", "Code": "5", "Grade": "3" },
      { "Id": "5", "Title": "机电", "Code": "6", "Grade": "3" },
      { "Id": "6", "Title": "施工", "Code": "7", "Grade": "3" }
    ];
    /*初始化table数据*/
    $(function() {
      $("#table1").bootstrapTable({
        data: json1
      });
    });

    //填充select选项
    $(document).ready(function() {
      //   $(array).each(function(index){
      //     alert(this);
      // });
      // $.each(array,function(index){
      //     alert(this);
      // });
      $("#admincategory").append('<option value="a">项目模板</option>');
      if ({{.Select2 }}) { //20171021从meirit修改而来
        $.each({{.Select2 }}, function(i, d) {
          // alert(this);
          // alert(i);
          // alert(d);
          $("#admincategory").append('<option value="' + i + '">' + d + '</option>');
        });
      }
    });

    //根据选择，刷新表格
    function refreshtable() {
      var newcategory = $("#admincategory option:selected").text();
      // alert("你选的是"+newcategory);
      if (newcategory == "项目模板") {
        //根据名字，查到id，或者另外写个categoryname的方法
        // $('#table2').bootstrapTable('refresh', {url:'/admin/categorytitle?title='+newcategory});
        $("#details").hide();
        $("#details1").show();
        // $('#table2').bootstrapTable('refresh', {url:'/project/getprojects'});
      } else {
        //根据名字，查到id，或者另外写个categoryname的方法
        $('#table1').bootstrapTable('refresh', { url: '/admin/categorytitle?title=' + newcategory });
        $("#details1").hide();
        $("#details").show();
      }
    }

    </script>

    <!-- 检索所有项目成果 -->
    <div class="form-horizontal">
      <div class="modal fade" id="modalTable2" style="width: 100%">
        <div class="modal-dialog">
          <div class="modal-content">
            <div class="modal-header">
              <button type="button" class="close" data-dismiss="modal">
                <span aria-hidden="true">&times;</span>
              </button>
              <h3 class="modal-title">检索所有项目成果</h3>
            </div>
            <div class="modal-body">
              <div class="modal-body-content">
                <!-- <div class=""> -->
                  <input type="text" class="form-control" placeholder="请输入关键字进行搜索" autocomplete="off" id="keyword2" onkeypress="getKey2();">
                  <span class="">
                    <button class="btn btn-default" type="button" id="search">
                      <i class="glyphicon glyphicon-search"></i>
                      Search!
                    </button>
                  </span>
                <!-- </div> -->
              </div>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
              <button type="button" class="btn btn-primary" onclick="searchAllProjectsProduct()">检索</button>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- 检索选定项目成果 -->
    <div class="form-horizontal">
      <div class="modal fade" id="modalTable3" style="width: 100%">
        <div class="modal-dialog">
          <div class="modal-content">
            <div class="modal-header">
              <button type="button" class="close" data-dismiss="modal">
                <span aria-hidden="true">&times;</span>
              </button>
              <h3 class="modal-title">检索选定项目成果</h3>
            </div>
            <div class="modal-body">
              <div class="modal-body-content">
                <!-- <div class=""> -->
                  <input type="text" class="form-control" placeholder="请输入关键字进行搜索" autocomplete="off" size="30" id="keyword3" onkeypress="getKey3();">
                  <span class="">
                    <button class="btn btn-default" type="button" id="search3">
                      <i class="glyphicon glyphicon-search"></i>
                      Search!
                    </button>
                  </span>
                <!-- </div> -->
              </div>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
              <button type="button" class="btn btn-primary" onclick="searchProjectProduct()">检索</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <script type="text/javascript">
  $(document).ready(function() {
    $("#modalTable").draggable({ handle: ".modal-header" }); //为模态对话框添加拖拽
    $("#modalTable1").draggable({ handle: ".modal-header" });
    $("#modalTable2").draggable({ handle: ".modal-header" });
    $("#modalTable3").draggable({ handle: ".modal-header" });
    $("#myModal").css("overflow", "hidden"); //禁止模态对话框的半透明背景滚动
  })

  var fm = new FloatModule({
    radius: '50%',
    theme_color: '#56b4f8',
    theme_content_color: '#fff',
    font_size: '18px',
    width_height: '60px',
    margin_screen_x: '30px',
    margin_screen_y: '30px',
    margin_li: '10px',
    animation: 'slide-in',
    position: 'right-bottom',
    icon_css_path: '',
    btn_config: [{
      icon: 'fa fa-search'
    }, {
      icon: 'fa fa-search-plus',
      title: '搜索所有项目',
      click: fnSearch
    }, {
      icon: 'fa fa-search-minus',
      title: '搜索选中项目',
      click: fnSearch2
    }]
  });

  function fnSearch() {
    $('#modalTable2').modal({
      show: true,
      backdrop: 'static'
    });
  }

  function fnSearch2() {
    var selectRow = $('#table0').bootstrapTable('getSelections');
    console.log(selectRow)
    if (selectRow.length < 1) {
      alert("请先勾选成果！");
      return;
    }
    $("input#cid").remove();
    var th1 = "<input id='cid' type='hidden' name='cid' value='" + selectRow[0].Id + "'/>";
    $(".modal-body").append(th1); //这里是否要换名字$("p").remove();
    $('#modalTable3').modal({
      show: true,
      backdrop: 'static'
    });
  }

  // 搜索所有项目成果
  $("#search").click(function() {
    window.location.href="/v1/wx/searchproduct?keyword="+$("#keyword2").val()
  });
  function searchAllProjectsProduct(){
    window.location.href="/v1/wx/searchproduct?keyword="+$("#keyword2").val()
  }
  function getKey2() {
    if (event.keyCode == 13) {
      window.location.href="/v1/wx/searchproduct?keyword="+$("#keyword2").val()
    }
  }
  // 搜索选定项目成果
  $("#search3").click(function() {
    window.location.href="/v1/wx/searchprojectproduct?productid="+$("#cid").val()+"&keyword="+$("#keyword3").val()
  });
  function searchProjectProduct(){
    window.location.href="/v1/wx/searchprojectproduct?productid="+$("#cid").val()+"&keyword="+$("#keyword3").val()
  }
  function getKey3() {
    if (event.keyCode == 13) {
      window.location.href="/v1/wx/searchprojectproduct?productid="+$("#cid").val()+"&keyword="+$("#keyword3").val()
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
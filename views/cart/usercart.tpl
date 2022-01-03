<!-- 视频列表 视频主页-->
<!DOCTYPE html>
<html lang="en">

<head>
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css" />
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css" />
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css" />
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
  <script src="/static/js/tableExport.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <script src="/static/js/jquery.form.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/select2.css" />
  <script type="text/javascript" src="/static/js/select2.js"></script>

  <script type="text/javascript" src="/static/js/jquery-ui.min.js"></script>
  <script type="text/javascript" src="/static/js/clipboard.min.js"></script>
  <style type="text/css">
  #modalDialog10 .modal-header {
    cursor: move;
  }

  #modalDialog11 .modal-header {
    cursor: move;
  }
  </style>
</head>
<!-- {{template "widgets/footer.tpl" .}} -->
<div class="container-fill">{{template "T.navbar1.tpl" .}}</div>
<!-- 引入的文件不能定义difine为模板！！！ -->
<!-- ../views/T.navbar.tpl -->
<!-- ../../views/T.navbar.tpl -->
<!-- ../T.navbar.tpl -->
<!-- ./T.navbar.tpl -->
<!-- /views/T.navbar.tpl -->
<body>
  <!-- <ul class="nav nav-tabs" id="myTab" role="tablist">
    <li class="nav-item" role="presentation">
      <a class="nav-link active" id="home-tab" data-toggle="tab" href="#home" role="tab" aria-controls="home" aria-selected="true">待审批</a>
    </li>
    <li class="nav-item" role="presentation">
      <a class="nav-link" id="profile-tab" data-toggle="tab" href="#profile" role="tab" aria-controls="profile" aria-selected="false">历史</a>
    </li>
    <li class="nav-item" role="presentation">
      <a class="nav-link" id="contact-tab" data-toggle="tab" href="#contact" role="tab" aria-controls="contact" aria-selected="false">Contact</a>
    </li>
  </ul> -->
  <div class="col-lg-12">
    <ul id="myTab" class="nav nav-tabs">
      <li class="active">
        <a href="#approval" data-toggle="tab">待我审批</a>
      </li>
      <li>
        <a href="#apply" data-toggle="tab">我的申请</a>
      </li>
      <li><a href="#approvalhistory" data-toggle="tab">审批历史</a>
      </li>
      <li><a href="#applyhistory" data-toggle="tab">申请历史</a>
      </li>
    </ul>
    <div id="myTabContent" class="tab-content">
      <div class="tab-pane fade in active" id="approval">
        <div id="toolbar" class="btn-group">
          <button {{if .IsAdmin}} style="display:none" {{end}} type="button" data-name="addButton" id="addButton" class="btn btn-default"> <i class="fa fa-envelope-o"> 申请</i>
          </button>
          <button type="button" data-name="shareButton" id="shareButton" class="btn btn-default"> <i class="fa fa-share"> 分享码</i>
          </button>
          <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
            <i class="fa fa-trash"> 删除</i>
          </button>
        </div>
        <table id="table"></table>
      </div>
      <div class="tab-pane fade" id="apply">
        <table id="table1"></table>
      </div>
      <div class="tab-pane fade" id="approvalhistory">
        <table id="table2"></table>
      </div>
      <div class="tab-pane fade" id="applyhistory">
        <table id="table3"></table>
      </div>
    </div>
  </div>
  <!-- 分享过期时间选择 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalShareExpireTime">
      <div class="modal-dialog" id="modalDialog10">
        <div class="modal-content">
          <div class="modal-header" style="background-color: #8bc34a">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">分享</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">有效期</label>
                <div class="col-sm-7">
                  <select name="expiretime" id="expiretime" class="form-control" required placeholder="选择有效期：">
                    <option value="HOUR">1小时</option>
                    <option value="DAY" selected>1天</option>
                    <option value="WEEK">1周</option>
                    <option value="MONTH">1个月</option>
                    <option value="YEAR">1年</option>
                    <option value="INFINITY">永远有效</option>
                  </select>
                </div>
              </div>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-primary btn-sm mr5" onclick="createShare()">分享</button>
            <button type="button" class="btn btn-default  btn-sm mr5" data-dismiss="modal">关闭</button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <!-- 分享结果信息 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalShare">
      <div class="modal-dialog" id="modalDialog11">
        <div class="modal-content">
          <div class="modal-header" style="background-color: #8bc34a">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title"><img src="/static/eyeblue/img/archive.77d78eb7.svg" class="share-icon">分享成果信息</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <!-- <div class="container"> -->
              <dl class="dl-horizontal">
                <dt><span class="italic"><i class="fa fa-check text-success"></i></span>分享成功：</dt>
                <dd id="sharetitle"></dd>
                <dt>分享者：</dt>
                <dd id="username" value=""></dd>
                <dt>失效时间：</dt>
                <dd id="expireTime"></dd>
                <dt>链接:</dt>
                <dd id="uuid"></dd>
                <dt>提取码：</dt>
                <dd id="code"></dd>
              </dl>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-primary copy" id="btncopy" onclick="copyuuid()">复制链接+提取码</button>
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <script type="text/javascript">
  // 待我审批 
  $('#table').bootstrapTable({
    url: '/v1/cart/getapprovalcart',
    search: 'true',

    showSearchClearButton: 'true',

    showRefresh: 'true',
    showColumns: 'true',
    toolbar: '#toolbar',
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
    // singleSelect:"true",
    clickToSelect: "true",
    showExport: "true",
    queryParams: function queryParams(params) { //设置查询参数
      var param = {
        limit: params.pageSize, //每页多少条数据
        pageNo: params.pageNumber, // 页码
        searchText: $(".search .form-control").val(),
        status: "0"
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
    columns: [{
      title: '选择',
      // radio: 'true',
      checkbox: 'true',
      width: '10',
      align: "center",
      valign: "middle"
    }, {
      // field: 'Number',
      title: '序号',
      formatter: function(value, row, index) {
        return index + 1
      },
      align: "center",
      valign: "middle"
    }, {
      field: 'usernickname',
      title: '用户',
      align: "center",
      valign: "middle"
    }, {
      field: 'producttitle',
      title: '成果名称',
      halign: "center",
      valign: "middle"
    }, {
      field: 'projecttitle',
      title: '目录名称',
      align: "center",
      valign: "middle"
    }, {
      field: 'topprojecttitle',
      title: '项目名称',
      align: "center",
      valign: "middle"
    }, {
      field: 'status',
      title: '状态',
      formatter: 'StatusFormatter',
      align: "center",
      valign: "middle"
    }, {
      field: 'updated',
      title: '日期',
      formatter: 'localDateFormatter',
      align: "center",
      valign: "middle"
    }, {
      field: 'action',
      title: '操作',
      formatter: 'actionFormatter',
      events: 'actionEvents',
      align: "center",
      valign: "middle"
    }]
  })
  // 我的申请
  $('#table1').bootstrapTable({
    url: '/v1/cart/getapplycart',
    search: 'true',

    showSearchClearButton: 'true',

    showRefresh: 'true',
    showColumns: 'true',
    // toolbar: '#toolbar',
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
    // singleSelect:"true",
    clickToSelect: "true",
    showExport: "true",
    queryParams: function queryParams(params) { //设置查询参数
      var param = {
        limit: params.pageSize, //每页多少条数据
        pageNo: params.pageNumber, // 页码
        searchText: $(".search .form-control").val(),
        status: "0"
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
    columns: [{
      title: '选择',
      // radio: 'true',
      checkbox: 'true',
      width: '10',
      align: "center",
      valign: "middle"
    }, {
      // field: 'Number',
      title: '序号',
      formatter: function(value, row, index) {
        return index + 1
      },
      align: "center",
      valign: "middle"
    }, {
      field: 'usernickname',
      title: '用户',
      align: "center",
      valign: "middle"
    }, {
      field: 'producttitle',
      title: '成果名称',
      halign: "center",
      valign: "middle"
    }, {
      field: 'projecttitle',
      title: '目录名称',
      align: "center",
      valign: "middle"
    }, {
      field: 'status',
      title: '状态',
      formatter: 'StatusFormatter',
      align: "center",
      valign: "middle"
    }, {
      field: 'updated',
      title: '日期',
      formatter: 'localDateFormatter',
      align: "center",
      valign: "middle"
    }, {
      field: 'topprojecttitle',
      title: '项目名称',
      align: "center",
      valign: "middle"
    }]
  })
  // approval历史
  $('#table2').bootstrapTable({
    url: '/v1/cart/getapprovalcart',
    search: 'true',

    showSearchClearButton: 'true',

    showRefresh: 'true',
    showColumns: 'true',
    // toolbar: '#toolbar',
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
    // singleSelect:"true",
    clickToSelect: "true",
    showExport: "true",
    queryParams: function queryParams(params) { //设置查询参数
      var param = {
        limit: params.pageSize, //每页多少条数据
        pageNo: params.pageNumber, // 页码
        searchText: $(".search .form-control").val(),
        status: "1"
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
    columns: [{
      title: '选择',
      // radio: 'true',
      checkbox: 'true',
      width: '10',
      align: "center",
      valign: "middle"
    }, {
      // field: 'Number',
      title: '序号',
      formatter: function(value, row, index) {
        return index + 1
      },
      align: "center",
      valign: "middle"
    }, {
      field: 'usernickname',
      title: '用户',
      align: "center",
      valign: "middle"
    }, {
      field: 'producttitle',
      title: '成果名称',
      halign: "center",
      valign: "middle"
    }, {
      field: 'projecttitle',
      title: '目录名称',
      align: "center",
      valign: "middle"
    }, {
      field: 'status',
      title: '状态',
      formatter: 'StatusFormatter',
      align: "center",
      valign: "middle"
    }, {
      field: 'updated',
      title: '日期',
      formatter: 'localDateFormatter',
      align: "center",
      valign: "middle"
    }, {
      field: 'topprojecttitle',
      title: '项目名称',
      align: "center",
      valign: "middle"
    }]
  })
  // apply历史
  $('#table3').bootstrapTable({
    url: '/v1/cart/getapplycart',
    search: 'true',

    showSearchClearButton: 'true',

    showRefresh: 'true',
    showColumns: 'true',
    // toolbar: '#toolbar',
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
    // singleSelect:"true",
    clickToSelect: "true",
    showExport: "true",
    queryParams: function queryParams(params) { //设置查询参数
      var param = {
        limit: params.pageSize, //每页多少条数据
        pageNo: params.pageNumber, // 页码
        searchText: $(".search .form-control").val(),
        status: "1"
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
    columns: [{
      title: '选择',
      // radio: 'true',
      checkbox: 'true',
      width: '10',
      align: "center",
      valign: "middle"
    }, {
      // field: 'Number',
      title: '序号',
      formatter: function(value, row, index) {
        return index + 1
      },
      align: "center",
      valign: "middle"
    }, {
      field: 'usernickname',
      title: '用户',
      align: "center",
      valign: "middle"
    }, {
      field: 'producttitle',
      title: '成果名称',
      halign: "center",
      valign: "middle"
    }, {
      field: 'projecttitle',
      title: '目录名称',
      align: "center",
      valign: "middle"
    }, {
      field: 'status',
      title: '状态',
      formatter: 'StatusFormatter',
      align: "center",
      valign: "middle"
    }, {
      field: 'updated',
      title: '日期',
      formatter: 'localDateFormatter',
      align: "center",
      valign: "middle"
    }, {
      field: 'topprojecttitle',
      title: '项目名称',
      align: "center",
      valign: "middle"
    }]
  })  

  // 批量删除
  $("#deleteButton").click(function(e, value, row, index) {
    var selectRow = $('#table').bootstrapTable('getSelections');
    if (selectRow.length <= 0) {
      alert("请先勾选成果！");
      return false;
    }
    if (confirm("确定删除吗？")) {
      // var title = $.map(selectRow, function(row) {
      //   return row.producttitle;
      // })
      var ids = "";
      for (var i = 0; i < selectRow.length; i++) {
        if (i == 0) {
          ids = selectRow[i].id;
        } else {
          ids = ids + "," + selectRow[i].id;
        }
      }
      //删除前端表格用的
      // var ids2 = $.map(selectRow, function(row) {
      //   return row.id;——这个行么？
      // })
      var ids2 = $.map($('#table').bootstrapTable('getSelections'), function(row){
        return row.id;
      });

      // var removeline=$(this).parents("tr")
      //提交到后台进行修改数据库状态修改
      $.ajax({
        type: "post", //这里是否一定要用post？？？
        url: "/v1/cart/deleteusercart",
        // data: JSON.stringify(selectRow3), //JSON.stringify(row),
        data: { ids: ids },
        success: function(data, status) { //数据提交成功时返回数据
          // $('#table').bootstrapTable('remove', {
          //   field: 'producttitle',
          //   values: title
          // });
          //删除已选数据
          $('#table').bootstrapTable('remove', {
            field: 'id',
            values: ids2
          });
          alert("删除“" + data.data + "”成功！(status:" + status + ".)");
        }
      });
    }
  })

  function actionFormatter(value, row, index) {
    return [
      '<a class="btn btn-danger btn-sm" href="javascript:void(0)" title="删除">',
      '<i id="delete" class="fa fa-trash-o"> Delete</i>',
      '</a>'
    ].join('');
  }

  window.actionEvents = {
    'click #delete': function(e, value, row, index) {
      if (confirm("确定删除吗？")) {
        var ids = "";
        ids = row.id
        // 删除前端表格用
        var mycars = new Array()
        mycars[0] = row;
        ids2 = $.map(mycars, function(row) {
          return row.id;
        });
        // var removeline=$(this).parents("tr")
        //提交到后台进行修改数据库状态修改
        $.ajax({
          type: "post", //这里是否一定要用post？？？
          url: "/v1/cart/deleteusercart",
          data: { ids: ids },
          success: function(data, status) { //数据提交成功时返回数据
            $('#table').bootstrapTable('remove', {
              field: 'id',
              values: ids2
            });
            // removeline.remove();
            alert("删除“" + data.data + "”成功！(status:" + status + ".)");
            // $('#table1').bootstrapTable('refresh', {url:'/admin/merit/meritlist/1'});
          }
        });
      }
    }
    // if(confirm("确定删除吗？")){  
    // var removeline=$(this).parents("tr")
    //     $.ajax({
    //     type:"post",//这里是否一定要用post？？？
    //     url:"/admin/merit/deletemeritlist",
    //     data: {CatalogId:row.Id},
    //         success:function(data,status){//数据提交成功时返回数据
    //         removeline.remove();
    //         alert("删除“"+data+"”成功！(status:"+status+".)");
    //         }
    //     });  
    // }
  };

  function localDateFormatter(value) {
    return moment(value, 'YYYY-MM-DD HH:mm:ss').add(8, 'hours').format('YYYY-MM-DD HH:mm:ss');
  }

  function StatusFormatter(value, row, index) {
    // alert(row.Status);
    if (row.status == 0) { //Status
      return '待审批';
    } else if (row.status == 1) {
      return '已审批';
    } else {
      return '失效';
    }
  }

  // 分享成果
  $("#shareButton").click(function() {
    var selectRow = $('#table').bootstrapTable('getSelections');
    if (selectRow.length < 1) {
      alert("请先勾选成果！");
      return;
    }
    // if (selectRow[0].Uid === {{.Uid }} || {{.RoleGet }} == "true") {
      $("div#flowattachment").remove();
      $("div#flowdocname").remove();
      $('#modalShareExpireTime').modal({
        show: true,
        backdrop: 'static'
      });
    // } else {
    //   alert("权限不够！" + selectRow[0].Uid);
    //   return;
    // }
  })

  //分享生成
  function createShare() {
    var selectRow = $('#table').bootstrapTable('getSelections');
    var ids = "";
    var expireInfinity = "";
    var expireTime = "";
    for (var i = 0; i < selectRow.length; i++) {
      if (i == 0) {
        ids = selectRow[i].productid;//这里将Id修改为productid
      } else {
        ids = ids + "," + selectRow[i].productid;//这里将Id修改为productid
      }
    }
    // ?matterUuids=1&expireInfinity=true&expireTime=1
    var expiretime = $('#expiretime').val();
    console.log(expireTime)
    if (expiretime == "INFINITY") {
      expireInfinity = "true";
      expireTime = moment().format('YYYY-MM-DD HH:mm:ss'); //getdate()
    } else if (expiretime == "HOUR") {
      expireInfinity = "false"
      expireTime = moment().add(1, 'hours').format('YYYY-MM-DD HH:mm:ss');
    } else if (expiretime == "DAY") {
      expireInfinity = "false"
      expireTime = moment().add(1, 'days').format('YYYY-MM-DD HH:mm:ss');
    } else if (expiretime == "WEEK") {
      expireInfinity = "false"
      expireTime = moment().add(1, 'weeks').format('YYYY-MM-DD HH:mm:ss');
    } else if (expiretime == "MONTH") {
      expireInfinity = "false"
      expireTime = moment().add(1, 'months').format('YYYY-MM-DD HH:mm:ss');
    } else if (expiretime == "YEAR") {
      expireInfinity = "false"
      expireTime = moment().add(1, 'years').format('YYYY-MM-DD HH:mm:ss');
    }
    $.ajax({
      type: "post",
      url: "/v1/share/create",
      data: { matterUuids: ids, expireInfinityStr: expireInfinity, expireTimeStr: expireTime },
      success: function(data, status) {
        $('#modalShareExpireTime').modal('hide');

        $("input#cid").remove();
        // var th1 = "<input id='cid' type='hidden' name='cid' value='" + selectRow[0].Id + "'/>"
        // $(".modal-body").append(th1); //这里是否要换名字$("p").remove();
        document.getElementById("sharetitle").innerText = data.data.name;
        document.getElementById("username").innerText = data.data.username;
        // document.getElementById("uuid").innerText='https://zsj.itdos.com/share/detail/'+data.data.uuid;
        // document.getElementById("uuid").append('<a title="复制链接" class="mr15"><i class="fa fa-copy"></i></a>');
        document.getElementById("uuid").innerHTML = '<span id="copyuuid">' + {{.Site }} + '/v1/share/detail/' + data.data.uuid + '</span><a title="复制链接" class="mr15" data-clipboard-target="#copyuuid"><i class="fa fa-copy"></i></a>'
        document.getElementById("code").innerHTML = '<span id="copycode">' + data.data.code + '</span><a title="复制提取码" class="mr15" data-clipboard-target="#copycode"><i class="fa fa-copy"></i></a>';

        document.getElementById("expireTime").innerText = moment(data.data.expireTime).format('YYYY-MM-DD HH:mm:ss')
        $('#modalShare').modal({
          show: true,
          backdrop: 'static'
        });
      },
    });
  }

  // 关闭模态框后将cart中数据修改状态
  $('#modalShare').on('hide.bs.modal', function () {
    // alert('模态框关闭了');
    var selectRow = $('#table').bootstrapTable('getSelections');
    var ids = "";
    var expireInfinity = "";
    var expireTime = "";
    for (var i = 0; i < selectRow.length; i++) {
      if (i == 0) {
        ids = selectRow[i].id;
      } else {
        ids = ids + "," + selectRow[i].id;
      }
    }
    //删除前端表格用的
    var ids2 = $.map($('#table').bootstrapTable('getSelections'), function(row){
      return row.id;
    });

    $.ajax({
      type: "post",
      url: "/v1/cart/updateapprovalcart",
      data: { ids: ids },
      success: function(data, status) {
        //删除已选数据
        $('#table').bootstrapTable('remove', {
          field: 'id',
          values: ids2
        });
      },
    });
  });

  new ClipboardJS('.mr15', {
    container: document.getElementById('copyuuid')
  });
  new ClipboardJS('.mr15', {
    container: document.getElementById('copycode')
  });

  function copyuuid() {
    var clipboard = new ClipboardJS("#btncopy", {
      text: function() {
        //寻找被激活的那个div pre元素,同时获取它下面的内容
        return '链接:' + $("span#copyuuid").text() + ' 提取码:' + $("span#copycode").text();
      }
    });
    clipboard.on('success', function(e) {
      alert("复制成功，去粘贴试试吧！");
      //可执行其他代码操作
    });
    clipboard.on('error', function(e) {
      alert("复制失败！请手动复制")
    });
  }

  $(document).ready(function() {
    //为模态对话框添加拖拽
    $("#modalDialog10").draggable({ handle: ".modal-header" });
    $("#modalDialog11").draggable({ handle: ".modal-header" });
    $("#myModal").css("overflow", "hidden"); //禁止模态对话框的半透明背景滚动
  })
  </script>
</body>
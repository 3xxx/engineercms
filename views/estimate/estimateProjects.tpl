<!-- 计算页面，核心-->
<!DOCTYPE html>
<html>

<head>
  <meta charset="UTF-8">
  <!-- 收藏用logo图标 -->
  <link rel="bookmark" type="image/x-icon" href="/static/img/pss.ico" />
  <!-- 网站显示页logo图标 -->
  <link rel="shortcut icon" href="/static/img/pss.ico">
  <meta name="renderer" content="webkit">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title> {{.Username}}</title>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-3.min.css" />
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table1.18.3.min.css"> -->
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css" />
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css" /> -->
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
  <script type="text/javascript" src="/static/js/jquery-3.6.0.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap3.4.1.min.js"></script>
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table1.18.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-editable.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <script type="text/javascript" src="/static/js/tableExport.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <script type="text/javascript" src="/static/js/jquery.form.js"></script>
  <style type="text/css">

  </style>
</head>

<body>
  <div class="col-lg-12">
    <!-- <h5>输入参数表（模板名称： ；作者：；）</h5> -->
    <!-- <div class="input-group">
      <span class="input-group-addon label-info">@</span>
      <input type="text" class="form-control" placeholder="输入20字以内的计算书标题/描述" name="description" id="description">
    </div> -->
    <div id="toolbar" class="btn-group">
      <button type="button" data-name="importButton" id="importButton" class="btn btn-default"> <i class="fa fa-plus">导入</i>
      </button>
      <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default" disabled="disabled">
        <i class="fa fa-trash">删除</i>
      </button>
    </div>
    <!--  data-search="true" data-show-refresh="true" data-show-toggle="true" data-show-columns="true" data-striped="true" data-toolbar="#toolbar" data-query-params="queryParams" data-sort-name="Username" data-sort-order="desc" data-page-size="5" data-page-list="[5, 25, 50, All]" data-unique-id="ID" data-pagination="true" data-side-pagination="client" data-single-select="true" data-click-to-select="true" data-show-export="true" -->
    <table id="table0"></table>
  </div>

  <script type="text/javascript">

  $("#addButton").click(function() {
    $('#modalTable').modal({
      show: true,
      backdrop: 'static'
    });
  })

  //importusers
  $("#importButton").click(function() {
    var url = '/v1/estimate/uploadexcel'
    window.open(url, "_blank", "")
  })

  $('#my_table_id').on('shown', function(e, editable) {
    editable.input.$input.val('overwriting value of input..');
  });


  $('#my_table_id').editable({
    type: 'select',
    pk: 1,
    url: '/post',
    title: 'Enter username',
  });

  $('#my_table_id').bootstrapTable({
  })

  $('#table0').bootstrapTable({
    idField: 'ID', //这个对应editable的pk传值
    url: '/v1/estimate/getestimateprojectsdata',
    method: 'get',
    search: 'true',
    classes: "table table-striped", //这里设置表格样式
    showRefresh: 'true',
    showToggle: 'true',
    showColumns: 'true',
    toolbar:'#toolbar',
    pagination: 'true',
    sidePagination: "server",
    queryParamsType: '',
    classes: "table table-striped", //这里设置表格样式
    //请求服务器数据时，你可以通过重写参数的方式添加一些额外的参数，例如 toolbar 中的参数 如果 queryParamsType = 'limit' ,返回参数必须包含
    // limit, offset, search, sort, order 否则, 需要包含:
    // pageSize, pageNumber, searchText, sortName, sortOrder.
    // 返回false将会终止请求。
    pageSize: 5,
    pageNumber: 1,
    pageList: [15, 20, 50, 100],
    singleSelect: "true",
    clickToSelect: "true",
    // selectItemName: "project",
    queryParams: function queryParams(params) { //设置查询参数
      var param = {
        limit: params.pageSize, //每页多少条数据
        pageNo: params.pageNumber, // 页码
        searchText: params.searchText // $(".search .form-control").val()
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
    // striped: "true",
    columns: [{
      title: '序号',
      formatter: function(value, row, index) {
        return index + 1
      },
      halign: "center",
      align: "center",
      valign: "middle"
    }, {
      field: 'number',
      title: '编号',
      halign: "center",
      align: "center",
      valign: "middle"
    }, {
      field: 'name',
      title: '工程名称',
      halign: "center",
      align: "left",
      valign: "middle",
    }, {
      field: 'grade',
      title: '工程等级',
      halign: "center",
      align: "center",
      valign: "middle",
    }, {
      field: 'period',
      title: '工期',
      halign: "center",
      align: "center",
      valign: "middle"
    }, {
      field: 'profile',
      title: '简介',
      halign: "center",
      align: "left",
      valign: "middle"
    }, {
      field: 'user.Nickname',
      title: '用户名',
      halign: "center",
      align: "center",
      valign: "middle",
    }, {
      field: 'estimateprojphase',
      title: '阶段',
      formatter: phaseFormatter,
      // events: phaseEvents,
      align: "center",
      valign: "middle"
    }, {
      field: 'operate',
      title: '操作',
      formatter: operateFormatter,
      // events: operateEvents,
      align: "center",
      valign: "middle"
    }]
  });

  function selectdata(value, row, index) {
    var data = $('#table0').bootstrapTable('getData')
    return row
  }

  function index1(value, row, index) {
    return index + 1
  }

  function localDateFormatter(value) {
    return moment(value, 'YYYY-MM-DD').format('YYYY-MM-DD');
  }

  function phaseFormatter(value, row, index) {
    var phasearray = new Array()
    for (i = 0; i < value.length; i++) {
      // relevarray[i]=value[i].Relevancy;
      phasearray[i] = '<a href="javascript:void(0);" onclick="getestimatecost(' + value[i].ID + ');return false;" title="查看" target="_blank">' + value[i].phasename + '</a>';
    }
    return phasearray.join(",");
  }

  function getestimatecost(e) {
    // console.log(e)
    var url = '/v1/estimate/getestimatecost/' + e
    window.open(url, "_blank", "")
  }

  window.phaseEvents = {
    'click .cal': function(e, value, row, index) {
      var url = '/v1/excel/excelcal/' + row.ID
      window.open(url, "_blank", "")
    },
    'click .pdf': function(e, value, row, index) {
      // alert("计算书模板预览功能待完善~");
      var url = '/v1/excel/getexceltemplepdf/'+row.ID
      window.open(url, "_blank", "")
    },
    'click .edit': function(e, value, row, index) {
      alert("编辑模板信息功能待完善~");
      var url = '/v1/excel/editor'
    },
    'click .history': function(e, value, row, index) {
      var url = '/v1/excel/gethistoryexcel/' + row.ID
      window.open(url, "_blank", "")
    },
    'click .delete': function(e, value, row, index) {
      // var url = '/v1/excel/deleteexcel/'+row.ID
      var mycars = new Array()
      mycars[0] = row;
      var excelid = $.map(mycars, function(row) {//必须构造这个map，否则删除不了前端
        return row.ID;
      })
      if (confirm("确定删除吗？")) {
        //提交到后台进行修改数据库状态修改
        $.ajax({
          type: "post", //这里是否一定要用post？？？
          url: '/v1/excel/deleteexcel/' + row.ID,
          // data: JSON.stringify(selectRow3), //JSON.stringify(row),
          success: function(data, status) { //数据提交成功时返回数据
            if (data.info == "SUCCESS") {
              $('#table0').bootstrapTable('remove', {
                field: 'ID', //table填充的数据结构中必须提供这个id，否则不能删除某行
                values: excelid//必须构造这个map，否则删除不了前端
              });
              // removeline.remove();
              alert("删除成功！(status:" + status + ".)");
            } else {
              alert("删除失败！(status:" + data.data + ".)");
            }
          }
        });
      }
    }
  };

  function operateFormatter(value, row, index) {
    return [
      '<a class="cal btn btn-xs btn-danger" style="margin-left:10px" href="javascript:void(0)" title="计算">',
      '<i class="fa fa-calculator"></i>',
      '</a>',
      '<a class="pdf btn btn-xs btn-success" style="margin-left:10px" href="javascript:void(0)" title="查看">',
      '<i class="fa fa-file-pdf-o"></i>',
      '</a>',
      '<a class="edit btn btn-xs btn-info" style="margin-left:10px" href="javascript:void(0)" title="编辑">',
      '<i class="fa fa-pencil"></i>',
      '</a>',
      '<a class="history btn btn-xs btn-primary" style="margin-left:10px" href="javascript:void(0)" title="计算历史">',
      '<i class="fa fa-history"></i>',
      '</a>',
      '<a class="delete btn btn-xs btn-danger" style="margin-left:10px" href="javascript:void(0)" title="删除">',
      '<i class="fa fa-trash-o"></i>',
      '</a>'
    ].join('');
  }

  window.operateEvents = {
    'click .cal': function(e, value, row, index) {
      var url = '/v1/excel/excelcal/' + row.ID
      window.open(url, "_blank", "")
    },
    'click .pdf': function(e, value, row, index) {
      // alert("计算书模板预览功能待完善~");
      var url = '/v1/excel/getexceltemplepdf/'+row.ID
      window.open(url, "_blank", "")
    },
    'click .edit': function(e, value, row, index) {
      alert("编辑模板信息功能待完善~");
      var url = '/v1/excel/editor'
    },
    'click .history': function(e, value, row, index) {
      var url = '/v1/excel/gethistoryexcel/' + row.ID
      window.open(url, "_blank", "")
    },
    'click .delete': function(e, value, row, index) {
      // var url = '/v1/excel/deleteexcel/'+row.ID
      var mycars = new Array()
      mycars[0] = row;
      var excelid = $.map(mycars, function(row) {//必须构造这个map，否则删除不了前端
        return row.ID;
      })
      if (confirm("确定删除吗？")) {
        //提交到后台进行修改数据库状态修改
        $.ajax({
          type: "post", //这里是否一定要用post？？？
          url: '/v1/excel/deleteexcel/' + row.ID,
          // data: JSON.stringify(selectRow3), //JSON.stringify(row),
          success: function(data, status) { //数据提交成功时返回数据
            if (data.info == "SUCCESS") {
              $('#table0').bootstrapTable('remove', {
                field: 'ID', //table填充的数据结构中必须提供这个id，否则不能删除某行
                values: excelid//必须构造这个map，否则删除不了前端
              });
              // removeline.remove();
              alert("删除成功！(status:" + status + ".)");
            } else {
              alert("删除失败！(status:" + data.data + ".)");
            }
          }
        });
      }
    }
  };

  // 更新输出表
  // success: function(response, newValue) {
  //   var selectRow3 = $('#table').bootstrapTable('getSelections');
  //   for (var i = 0; i < selectRow3.length; i++) {
  //     $('#table').bootstrapTable('updateByUniqueId', {
  //       id: selectRow3[i].id,
  //       row: {
  //         ProjectNumber: response
  //       }
  //     });
  //   }
  // }
  </script>
</body>

</html>

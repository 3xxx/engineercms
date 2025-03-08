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
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css" />
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table1.18.3.min.css"> -->
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
  #chat-messages {
    min-height: 10vh;
    height: 60vh;
    width: 100%;
    overflow-y: scroll;
  }
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
      <button type="button" data-name="addButton" id="addButton" class="btn btn-default" disabled="disabled"> <i class="fa fa-plus">添加</i>
      </button>
      <button type="button" data-name="importButton" id="importButton" class="btn btn-default" disabled="disabled"><i class="fa fa-plus">导入</i>
      </button>
      <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default" disabled="disabled">
        <i class="fa fa-trash">删除</i>
      </button>
    </div>
    <table id="table0"></table>
  </div>

  <script type="text/javascript">
  $(function() {
    window.onbeforeunload = function() { return ("确认关闭吗？"); }
  })

  $("#addButton").click(function() {
    $('#modalTable').modal({
      show: true,
      backdrop: 'static'
    });
  })

  //importusers
  $("#importButton").click(function() {
    $('#importusers').modal({
      show: true,
      backdrop: 'static'
    });
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
    url: '/v1/estimate/getestimatecostdata/{{.PhaseID}}',
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
      field: 'costname',
      title: '工程或费用名称',
      halign: "center",
      align: "left",
      valign: "middle",
    }, {
      field: 'unit',
      title: '单位',
      halign: "center",
      align: "center",
      valign: "middle",
    }, {
      field: 'quantity',
      title: '数量',
      halign: "center",
      align: "center",
      valign: "middle"
    }, {
      field: 'unitprice',
      title: '单价(元)',
      halign: "center",
      align: "center",
      valign: "middle"
    }, {
      field: 'total',
      title: '合计(万元)',
      halign: "center",
      align: "center",
      valign: "middle",
    }]
  });

  function onMsoNumberFormat(value, row, index) { //科学计数法数字转换为文本，时间格式转换为文本可以显示秒
    // console.log(row.resulttype)//alert无法显示对象，应该用consle.log来显示对象
    // if (row.resulttype == "Real") {
      // var num = new Number(value);
      // return num
      // } else if(row.resulttype == "MATRIX"){
      //   console.log(row.inputvalue)
      //   return row.inputvalue
    // } else {
      return value
    // }
  }

    /**
   * 数字转为科学计数法
   */
  function formatAmount(value){
    if(value){
      var _value = parseFloat(value);
      console.log(_value);
      if(!isNaN(_value)){
        //将数值保存两位小数
        var str = _value.toFixed(2);
        //将整数部分每隔千位分割
        str = str.split('').reverse().join('').replace(/(\d{3})/g, '$1,').replace(/\,$/, '').split('').reverse().join('');
        return str;
      }
    }
    return '0.00' ;
  }

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

  // 点击行显示角色
  $(function() {
    $("#table0").on("click-row.bs.table", function(e, row, ele) {
      $(".info").removeClass("info");
      $(ele).addClass("info");
      userid = row.Id; //全局变量
      rowtitle = row.Nickname
      $("#rowtitle").html("用户角色-" + rowtitle);
      $("#details").show();
      $('#table1').bootstrapTable('refresh', { url: '/admin/role/get/' + row.Id });
    });
  });

  function actionFormatter(value, row, index) {
    return '<button type="button" data-name="addButton" id="addButton" class="btn btn-info btn-xs"> <i class="fa fa-user">角色</i></button>';
  }

  window.actionEvents = {
    //弹出角色选择模态框，选择后保存——未修改
    'click .send': function(e, value, row, index) {
      var selectRow3 = $('#table').bootstrapTable('getSelections');
      if (selectRow3.length == 0) {
        var mycars = new Array()
        mycars[0] = row;
        var selectRow3 = mycars
      }
      if (confirm("确定提交吗？")) {
        var ids = $.map($('#table').bootstrapTable('getSelections'), function(row) {
          return row.id;
        });
        if (ids.length == 0) {
          ids = $.map(mycars, function(row) {
            return row.id;
          });
        }
        // var removeline=$(this).parents("tr")
        //提交到后台进行修改数据库状态修改
        $.ajax({
          type: "post", //这里是否一定要用post？？？
          url: "/achievement/sendcatalog",
          data: JSON.stringify(selectRow3), //JSON.stringify(row),
          success: function(data, status) { //数据提交成功时返回数据
            $('#table').bootstrapTable('remove', {
              field: 'id', //table填充的数据结构中必须提供这个id，否则不能删除某行
              values: ids
            });
            // removeline.remove();
            alert("提交“" + data + "”成功！(status:" + status + ".)");
            // $('#table1').bootstrapTable('refresh', {url:'/admin/merit/meritlist/1'});
          }
        });
      }
    },
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

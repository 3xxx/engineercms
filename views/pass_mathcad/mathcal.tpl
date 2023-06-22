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
  <title> {{substr .UserTemple.TempTitleB 0 6}}...</title>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-3.min.css" />
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table1.18.3.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css" />
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
  <script type="text/javascript" src="/static/js/jquery-3.6.0.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap3.4.1.min.js"></script>
  <!-- <script src="/static/js/bootstrap-treeview.js"></script> -->
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table1.18.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script>
  <!-- <script type="text/javascript" src="/static/js/bootstrap-table-editable1.18.3.min.js"></script> -->
  <script type="text/javascript" src="/static/js/bootstrap-table-editable.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <!-- Latest compiled and minified CSS -->
  <!-- <link rel="stylesheet" href="https://unpkg.com/bootstrap-table@1.18.3/dist/bootstrap-table.min.css"> -->
  <!-- Latest compiled and minified JavaScript -->
  <!-- <script src="https://unpkg.com/bootstrap-table@1.18.3/dist/bootstrap-table.min.js"></script> -->
  <!-- Latest compiled and minified Locales -->
  <!-- <script src="https://unpkg.com/bootstrap-table@1.18.3/dist/locale/bootstrap-table-zh-CN.min.js"></script> -->
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css" /> -->
  <!-- <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script> -->
  <!-- <script type="text/javascript" src="https://unpkg.com/bootstrap-table@1.18.3/dist/extensions/editable/bootstrap-table-editable.min.js"></script>  -->
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <script type="text/javascript" src="/static/js/tableExport.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <script type="text/javascript" src="/static/js/jquery.form.js"></script>
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/select2-d.min.css" /> -->
  <!-- <script type="text/javascript" src="/static/js/select2-d.min.js"></script> -->
  <style type="text/css">
  #chat-messages {
    min-height: 10vh;
    height: 60vh;
    width: 100%;
    overflow-y: scroll;
  }
  </style>
</head>
<div class="container-fill">{{template "navbar" .}}</div>

<body>
  <div class="col-lg-12">
    <h5>输入参数表（模板名称：{{.UserTemple.TempTitleB}} ；作者：{{.UserTemple.User.Nickname}} ；是否有效：{{.UserTemple.Status}}）</h5>
    <div class="input-group">
      <span class="input-group-addon label-info">@</span>
      <input type="text" class="form-control" placeholder="输入20字以内的计算书标题/描述" name="description" id="description">
    </div>
    <div id="toolbar" class="btn-group">
      <button type="button" data-name="addButton" id="addButton" class="btn btn-default" disabled="disabled"> <i class="fa fa-plus">添加</i>
      </button>
      <button type="button" data-name="importButton" id="importButton" class="btn btn-default" disabled="disabled"> <i class="fa fa-plus">导入</i>
      </button>
      <!-- <button type="button" data-name="editorButton" id="editorButton" class="btn btn-default"> <i class="fa fa-edit">编辑</i>
        </button> -->
      <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default" disabled="disabled">
        <i class="fa fa-trash">删除</i>
      </button>
    </div>
    <!--   <table id="my_table_id"
      data-url="/v1/mathcad/getmathcalinput/{{.ID}}"
      data-id-field="ID"
      data-editable-emptytext="Default empty text."
      data-editable-url="/my/editable/update/path">
      <thead>
        <tr>
          <th class="col-md-1" data-field="inputalias" data-sortable="true" data-align="center">#</th>
          <th class="col-md-4" data-field="inputvalue">Name</th>
          <th class="col-md-7" data-field="resulttype" data-editable="true" data-editable-emptytext="Custom empty text.">Description</th>
        </tr>
      </thead>
    </table> -->
    <table class="table table-striped" id="table0" data-search="true" data-show-refresh="true" data-show-toggle="true" data-show-columns="true" data-striped="true" data-toolbar="#toolbar" data-query-params="queryParams" data-sort-name="Username" data-sort-order="desc" data-page-size="5" data-page-list="[5, 25, 50, All]" data-unique-id="ID" data-pagination="true" data-side-pagination="client" data-single-select="true" data-click-to-select="true" data-show-export="true">
    </table>
  </div>
  <div class="col-lg-6">
    <button type="button" data-name="Submit1" id="Submit1" class="btn btn-danger" onclick="oncalculateMC()"> <i class="fa fa-calculator"> 计算</i></button>
    <label><input type="checkbox" id="PrintPDF">生成PDF</label>
    <div class="timer">
      <strong id="minute">0分</strong>
      <strong id="second">00秒</strong>
    </div>
    <a href="" target="_blank" id="pdflink" style="display:none">PDF计算书下载</a>
  </div>
  <!-- websocket反馈信息框 -->
  <form role="form" style="margin: 16px">
    <div class="form-group">
      <textarea id="content" class="form-control" rows="6" placeholder="这个区域是WebSocket回显信息。**注意**：1.要先登录。2.隔几秒再刷新本页面，当看到自己“…加入了 页面 队列” 才可以计算；3.不要多个页面打开同一个模板。"></textarea>
    </div>
  </form>

  <div class="col-lg-12 table-responsive">
    <h3>输出参数表</h3>
    <div id="toolbar2" class="btn-group">
      <button type="button" data-name="addButton" id="addButton" class="btn btn-default" disabled="disabled"> <i class="fa fa-plus">添加</i>
      </button>
      <button type="button" data-name="importButton" id="importButton" class="btn btn-default" disabled="disabled"> <i class="fa fa-plus">导入</i>
      </button>
      <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default" disabled="disabled">
        <i class="fa fa-trash">删除</i>
      </button>
    </div>
    <!-- bootstrap table 自动换行style="word-break:break-all; word-wrap:break-all;" -->
    <!-- 响应式-进度条,<div class="table-responsive"><table class="table table-striped"> -->
    <table class="table table-striped" id="table2" data-search="true" data-show-refresh="true" data-show-toggle="true" data-show-columns="true" data-striped="true" data-toolbar="#toolbar2" data-query-params="queryParams" data-sort-name="Username" data-sort-order="desc" data-page-size="5" data-page-list="[5, 25, 50, All]" data-unique-id="ID" data-pagination="true" data-side-pagination="client" data-single-select="true" data-click-to-select="true" data-show-export="true">
    </table>
  </div>
  <script type="text/javascript">
  $(function() {
    window.onbeforeunload = function() { return ("确认关闭吗？"); }
  })
  // function setCookie (name,value,iDay) {//存储cookie
  //   var oDate=new Date();
  //   oDate.setDate(oDate.getDate()+iDay);
  //   document.cookie=name+'='+value+';expires='+oDate;
  // };
  // //监听浏览器关闭事件(绝对好使，需所有代码放到生命周期mounted中)
  // var userAgent = navigator.userAgent; //取得浏览器的userAgent字符串  
  // var isOpera = userAgent.indexOf("Opera") > -1; //判断是否Opera浏览器  
  // var isIE = userAgent.indexOf("compatible") > -1 && userAgent.indexOf("MSIE") > -1 && !isOpera; //判断是否IE浏览器
  // var isIE11 = userAgent.indexOf("rv:11.0") > -1; //判断是否是IE11浏览器
  // var isEdge = userAgent.indexOf("Edge") > -1 && !isIE; //判断是否IE的Edge浏览器
  // if(!isIE && !isEdge && !isIE11) {//兼容chrome和firefox
  //   var _beforeUnload_time = 0, _gap_time = 0;
  //   var is_fireFox = navigator.userAgent.indexOf("Firefox")>-1;//是否是火狐浏览器
  //   window.onunload = function (){
  //     _gap_time = new Date().getTime() - _beforeUnload_time;
  //     if(_gap_time <= 5){
  //       setCookie('times',true);//谷歌浏览器关闭
  //     }else{//谷歌浏览器刷新
  //       return 'hhh';
  //     }
  //   }
  //   window.onbeforeunload = function (){ 
  //     _beforeUnload_time = new Date().getTime();
  //     if(is_fireFox){//火狐关闭执行
  //       setCookie('times',true);//火狐浏览器关闭
  //     }else{//火狐浏览器刷新
  //       return 'hhh';
  //     }
  //   };
  // }

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


  $(function() {
    if ({{.UserTemple.Status }} == false) {
      alert("本模板为失效版本，请使用新版本！")
    }
  })

  $('#my_table_id').editable({
    type: 'select',
    pk: 1,
    url: '/post',
    title: 'Enter username',
  });

  $('#my_table_id').bootstrapTable({
    // noEditFormatter (value, row, index) {
    //   consle.log(value)
    //   if (value === 'noEdit') {
    //     return 'No Edit'
    //   }
    //   return false
    // }
  })
  // $().editable()
  // })

  $('#table0').bootstrapTable({
    idField: 'ID', //这个对应editable的pk传值
    url: '/v1/mathcad/getmathcalinput/{{.ID}}',
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
      field: 'inputalias',
      title: '别名',
      halign: "center",
      align: "center",
      valign: "middle"
    }, {
      field: 'inputvalue',
      title: '数值',
      halign: "center",
      align: "center",
      valign: "middle",
      editable: {
        type() {
          // console.log(value)
          var pk = $(this).attr("data-pk");
          // console.log(pk)
          var selectRow3 = $('#table0').bootstrapTable('getRowByUniqueId', pk);
          // console.log(selectRow3)
          if (selectRow3.resulttype == "ComboBoxControl") {
            edittype = 'select'
          } else if (selectRow3.resulttype == "MATRIX") {
            edittype = 'textarea'
          } else {
            edittype = 'text'
          }
          return edittype
        },
        source: function() {
          var pk = $(this).attr("data-pk");
          var selectRow3 = $('#table0').bootstrapTable('getRowByUniqueId', pk);
          // console.log(selectRow3.resulttype)
          if (selectRow3.resulttype == "ComboBoxControl") {
            var arr = selectRow3.selectvalue
            // console.log(arr)
            return arr
            // }else if(selectRow3.resulttype == "MATRIX"){//这里是不会执行的，只有select才会执行这里。
            //   var arr = selectRow3.textarealvalue
            //   console.log(arr)
            //   return arr
          }
        },
        select: {
          allowClear: true,
          // width: '150px',
          placeholder: '请选择',
          // multiple: true
        },
        pk: 1,
        // url: 'select',
        // url: function () {
        //   return $.post('/echo/json/', {
        //     json: 'false',
        //     delay: .5
        //   });
        // },
        validate: function(value) {
          var pk4 = $(this).attr("data-pk");
          // console.log(pk)
          var selectRow4 = $('#table0').bootstrapTable('getRowByUniqueId', pk4);
          // console.log(selectRow4)
          // if (value == 'bob') {
          //   return 'bob is not allowed';
          // }
          if ($.trim(value) == '') {
            return '不能为空！';
          }
          if (selectRow4.resulttype == 'Real') { //只能输入数值。如果是string，则可以输入任意，但不能输入xss攻击代码
            // if (!!!value.match(/^[0-9.]{0,20}$/)){
            // const reg = /^(?:-?\d+|-?\d{1,3}(?:,\d{3})+)?(?:\.\d+)?$/;
            const reg = /^[Ee0-9.+-]+$/
            if (!reg.test(value)) {
              return '必须是数值';
            }
          }
          if (selectRow4.realmin != '' || selectRow4.realmax != '') {
            if (parseFloat(value) > parseFloat(selectRow4.realmax) || parseFloat(value) < parseFloat(selectRow4.realmin)) {
              return '超出范围，输入范围：' + selectRow4.realmin + '~' + selectRow4.realmax;
            }
          }
          // 验证矩阵
          if (selectRow4.resulttype == 'MATRIX') {
            // var pattern = /^[0-9=.,:\s-]+$/;
            var pattern = /^[Ee0-9.+-,:\s-]+$/
            if (!pattern.test(value)) {
              return '只能是数值、英文逗号,和换行';
            }
          }
          // 验证文本，不能含<>
          if (selectRow4.resulttype == 'String') {
            var pattern = /\<|\>|\"|\'|\&/g;
            if (pattern.test(value)) {
              return "不能输入<>&";
            }
          }
        },
        success: function(response) {
          if (response === false) {
            console.log('remote error from success');
            return 'remote error';
          }
        },
        error: function(response) {
          console.log('remote error from fail');
          return 'remote error';
        },
        title: '选择或输入',
        params: function(params) {
          //originally params contain pk, name and value
          //取出选择的值
          //与list对比，得到listid
          //要么使用select2方式，type:'select2'
          // var choose = $('#choose option:selected').val();
          params.bb = 'select';
          return params;
        }
      },
      formatter: onMsoNumberFormat
    }, {
      field: 'inputvalue',
      title: '默认值',
      halign: "center",
      align: "center",
      valign: "middle",
      formatter: onMsoNumberFormat
    }, {
      field: 'resulttype',
      title: '值类型',
      halign: "center",
      align: "center",
      valign: "middle"
    }, {
      field: 'units',
      title: '单位',
      halign: "center",
      align: "center",
      valign: "middle"
    }, {
      field: 'realmin',
      title: '最小值',
      halign: "center",
      align: "center",
      valign: "middle",
      editable: {
        noEditFormatter(value, row, index) {
          // console.log(value)
          var bools = {{.RoleUpdate }}
          if (!bools) {
            return value
          }
          return false
        },
        type: 'text',
        pk: 1,
        url: '/v1/mathcad/putmathcalinput',
        title: '输入小值',
        // emptytext: "点击输入",
        // noeditFormatter: true,
        // function (value,row,index) {
        //   // if(row.isParent == "Y"){
        //   // var bool={{.RoleUpdate}}
        //   // if(bool){
        //     console.log(value)
        //   return "false";
        //   // }
        //   //  else{
        //   //   return "+";
        //   //   // return false;
        //   // }
        // }
      },
    }, {
      field: 'realmax',
      title: '最大值',
      halign: "center",
      align: "center",
      valign: "middle",
      editable: {
        noEditFormatter(value, row, index) {
          // console.log(value)
          var bools = {{.RoleUpdate }}
          if (!bools) {
            return value
          }
          return false
        },
        type: 'text',
        pk: 1,
        url: '/v1/mathcad/putmathcalinput',
        title: '输入大值'
      },
    }, {
      field: 'comment',
      title: '说明',
      halign: "center",
      // align: "center",
      valign: "middle",
      width: "200",
      editable: {
        noEditFormatter(value, row, index) {
          // console.log(value)
          var bools = {{.RoleUpdate }}
          if (!bools) {
            return value
          }
          return false
        },
        type: 'text',
        pk: 1,
        url: '/v1/mathcad/putmathcalinput',
        title: 'Enter MathComment'
      },
    }]
  });

  $('#table2').bootstrapTable({
    // uniqueId: "outputalias", //每一行的唯一标识，一般为主键列
    idField: 'ID', //可以通过种方式设置主键
    url: '/v1/mathcad/getmathcaloutput/{{.ID}}',
    // striped: "true",
    columns: [{
        title: '序号',
        formatter: function(value, row, index) {
          return index + 1
        },
        halign: "center",
        align: "center",
        valign: "middle"
      },
      // {
      //   field: 'ID',
      //   title: 'ID',
      //   halign: "center",
      //   align: "center",
      //   valign: "middle"
      // },
      {
        field: 'outputalias',
        title: '别名',
        halign: "center",
        align: "center",
        valign: "middle"
      }, {
        //   field: 'outputvalue',
        //   title: '数值',
        //   halign: "center",
        //   align: "center",
        //   valign: "middle",
        //   editable: {
        //     type: 'text',
        //   },
        //   formatter: onMsoNumberFormat
        // }, {
        field: 'outputvalue',
        title: '数值',
        halign: "center",
        align: "center",
        valign: "middle",
        editable: {
          type() {
            var pk = $(this).attr("data-pk");
            // console.log(pk)
            var selectRow3 = $('#table2').bootstrapTable('getRowByUniqueId', pk);
            if (selectRow3.resulttype == "ComboBoxControl") {
              edittype = 'select'
            } else if (selectRow3.resulttype == "MATRIX") {
              edittype = 'textarea'
            } else {
              edittype = 'text'
            }
            return edittype
          },
          pk: 1,
        },
        formatter: onMsoNumberFormat
      }, {
        field: 'outputvalue',
        title: '默认值',
        halign: "center",
        align: "center",
        valign: "middle",
        editable: {
          type() {
            var pk = $(this).attr("data-pk");
            // console.log(pk)
            var selectRow3 = $('#table2').bootstrapTable('getRowByUniqueId', pk);
            if (selectRow3.resulttype == "MATRIX") {
              edittype = 'textarea'
            }
            return edittype
          },
          pk: 1,
          // disabled: false,不起作用
        },
        // disabled: true,
        formatter: onMsoNumberFormat,
        visible: "true"
      }, {
        field: 'resulttype',
        title: '值类型',
        halign: "center",
        align: "center",
        valign: "middle"
      }, {
        field: 'units',
        title: '单位',
        halign: "center",
        align: "center",
        valign: "middle"
      }, {
        field: 'comment',
        title: '说明',
        halign: "center",
        // align: "center",
        valign: "middle",
        width: "200",
        editable: {
          noEditFormatter(value, row, index) {
            // console.log(value)
            var bools = {{.RoleUpdate }}
            if (!bools) {
              return value
            }
            return false
          },
          type: 'text',
          pk: 1,
          url: '/v1/mathcad/putmathcaloutput',
          title: 'Enter MathComment'
        },
      }
    ]
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
  function formatAmount(value) {
    if (value) {
      var _value = parseFloat(value);
      console.log(_value);
      if (!isNaN(_value)) {
        //将数值保存两位小数
        var str = _value.toFixed(2);
        //将整数部分每隔千位分割
        str = str.split('').reverse().join('').replace(/(\d{3})/g, '$1,').replace(/\,$/, '').split('').reverse().join('');
        return str;
      }
    }
    return '0.00';
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
    // $("#table").bootstrapTable('destroy').bootstrapTable({
    //     columns:columns,
    //     data:json
    // });
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

  //删除按钮与修改按钮的出现与消失
  //只勾选一项，可以修改删除
  //勾选多项只能删除
  $('.bootstrap-table').change(function() {
    var dataArr = $('#mytab .selected');
    if (dataArr.length == 1) {
      $('#btn_edit').css('display', 'block').removeClass('fadeOutRight').addClass('animated fadeInRight');
    } else {
      $('#btn_edit').addClass('fadeOutRight');
      setTimeout(function() {
        $('#btn_edit').css('display', 'none');
      }, 400);
    }
    if (dataArr.length >= 1) {
      $('#btn_delete').css('display', 'block').removeClass('fadeOutRight').addClass('animated fadeInRight');
    } else {
      $('#btn_delete').addClass('fadeOutRight');
      setTimeout(function() {
        $('#btn_delete').css('display', 'none');
      }, 400);
    }
  });

  // 用户打开页面，建立ws长连接
  $(function() {
    document.getElementById("content").value="这个区域是WebSocket回显信息。**注意**：1.要先登录。2.隔几秒再刷新本页面，当看到 自己“…加入了 页面 队列” 才可以计算；3.不要多个页面打开同一个模板。"+ '\n'
    var ws = new WebSocket('ws://' + window.location.host + '/v1/mathcad/wsmathcal/{{.ID }}');
    ws.onmessage = function(event) {
      var msg = JSON.parse(event.data);
      var ele = document.getElementById("content");
      ele.value = ele.value + msg.message + '\n';

      ele.scrollTop = ele.scrollHeight;
      //添加返回消息 告诉服务器 此页面还存在
      var obj = {};
      obj.message = 'index.html 还活着 ' + new Date().toLocaleString();
      obj.templeid = {{.ID }}
      // obj.description = document.getElementById("description").value
      // console.log(obj);

      // var data = $('#table0').bootstrapTable('getData')
      // var data2 = JSON.stringify(data)
      // data: { templeid: {{.ID }}, inputdata: data2 },
      // obj.inputdata = data

      obj = JSON.stringify(obj); //将JSON对象转化为JSON字符
      ws.send(obj);
      // alert(msg.info)
      // if (msg.info == "ERROR") {
      //   alert("错误！(status:" + msg.message + ".)");
      // }else if(msg.info == "SUCCESS"){
      //   alert("信息！(status:" + msg.message + ".)");
      // }
    };
  });

  //计时器
  var intDiff = parseInt(1); //倒计时总秒数量
  function oncalculateMC() {
    //两种设置disabled属性
    $('#Submit1').attr("disabled", true);
    $("#pdflink").css("display", "none");
    console.log(intDiff);
    timer(intDiff); //计时器开始
    // var $ul = $('#msg-list');
    var ws = new WebSocket('ws://' + window.location.host + '/v1/mathcad/postmath2/{{.ID }}');
    ws.onmessage = function(event) {
      // var last = JSON.parse(event.data); //使用JSON.parse() 将JSON字符串转为JS对象;
      // $('<li>').text(last.message+last.usernickname+',前面有：'+last.size+'人').appendTo($ul);
      // console.log(last)
      var msg = JSON.parse(event.data);
      // $('<li>').text('<div class="chip">'+ msg.message+msg.usernickname+',前面有：'+msg.size+'人'+ '</div>'+ '<br/>').appendTo($ul);
      var ele = document.getElementById("content");
      // ele.value = ele.value.replace(/\r?\n/g, '<br />');+msg.usernickname+',前面有：'+msg.size+'人'
      ele.value = ele.value + msg.message + '\n';

      ele.scrollTop = ele.scrollHeight;
      // console.log(last)
      // $('#chat-messages').text +=
      // + '<img src="' + self.gravatarURL(msg.email) + '">' // Avatar
      // + emojione.toImage(msg.message) + '<br/>'; // Parse emojis
      // var element = document.getElementById('chat-messages');
      // element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
      //添加返回消息 告诉服务器 此页面还存在
      var obj = {};
      obj.message = 'index.html 还活着 ' + new Date().toLocaleString();
      obj.templeid = {{.ID }}
      obj.description = document.getElementById("description").value

      obj.printpdf = $('#PrintPDF').is(':checked');
      // console.log(obj);
      // obj.printpdf = $('#PrintPDF').attr('checked');这个方法不行
      // console.log(obj);
      // obj.printpdf = document.getElementById('PrintPDF').checked;这个方法可以
      console.log(obj);

      var data = $('#table0').bootstrapTable('getData')
      // var data2 = JSON.stringify(data)
      // data: { templeid: {{.ID }}, inputdata: data2 },
      obj.inputdata = data
      obj = JSON.stringify(obj); //将JSON对象转化为JSON字符
      ws.send(obj);
      console.log(msg)
      if (msg.info == "ERROR") {
        clearInterval(time),
          function() {
            timer();
          }
        // alert("错误！(status:" + msg.message + ".)");
        $('#Submit1').attr("disabled", false);
      } else if (msg.info == "SUCCESS") {
        alert("计算成功！(status:" + msg.message + ".)");
        // 表格刷新
        // for (var i = 0; i < msg.data.length; i++) {
        //   $('#table2').bootstrapTable('updateByUniqueId', {
        //     id: msg.data[i].ID, //outputalias,
        //     row: {
        //       outputvalue: msg.data[i].outputvalue
        //     }
        //   });
        // }
        $('#table2').bootstrapTable('load',msg.data)
        if (msg.pdflink!=""){
          $("#pdflink").css("display", "block");
          $("#pdflink").attr("href", "/v1/wx/getmathpdf/" + msg.pdflink);
        }  
        $('#Submit1').attr("disabled", false);
        clearInterval(time),
        function() {
          timer();
        }
      }
    };
    // });
  };

  function oncalculateMC_back() {
    //两种设置disabled属性
    $('#Submit1').attr("disabled", true);
    $("#pdflink").css("display", "none");
    // $('#uid').attr("disabled","disabled");
    //三种移除disabled属性
    // $('#Submit1').attr("disabled",false);
    // $('#uid').removeAttr("disabled");
    // $('#uid').attr("disabled","");
    timer(intDiff); //计时器开始
    var data = $('#table0').bootstrapTable('getData')
    // 坑：x-editable修改数值后，只能是text等的string，见它的源码Type of input. Can be <code>text|textarea|select|date|checklist</code>
    // 需要转成数值，jquery中只要*1即可将文本转为数值
    // for(var i=0;i<data.length;i++){
    //   data[i].realvalue=data[i].realvalue*1;
    // }
    // 全部按string传到服务端
    var data2 = JSON.stringify(data)
    // var data = str*1;
    $.ajax({
      type: "post",
      url: "/v1/mathcad/postmath",
      // data: { firstInput: Number(Select1.value), secondInput: Number(Text1.value), thirdInput: Number(Text2.value) },
      data: { templeid: {{.ID }}, inputdata: data2 },
      success: function(data, status) {
        if (data.info == "ERROR") {
          clearInterval(time),
            function() {
              timer();
            }
          alert("错误！(status:" + data.data + ".)");
          $('#Submit1').attr("disabled", false);
        } else if (data.info == "SUCCESS") {
          alert("计算成功！(status:" + status + ".)");
          // Text3.value = data.data;
          // 表格刷新
          for (var i = 0; i < data.data.length; i++) {
            // alert(data.data[i].ID);
            // alert(data.data[i].realvalue);
            $('#table2').bootstrapTable('updateByUniqueId', {
              id: data.data[i].outputalias,
              row: {
                outputvalue: data.data[i].outputvalue
              }
            });
          }
          // var html ="<input type=\"button\" value=\"下载pdf\" onclick=\"newButton();\">";
          // var html2 ='<a href="'+ data.pdflink + ' target="_blank">查看pdf</a>'
          // document.getElementById("toolbar").innerHTML=html;
          // $("#toolbar").append(html2);
          // var obj = document.getElementById("pdflink");
          // obj.setAttribute("display", "block");
          $("#pdflink").css("display", "block");
          $("#pdflink").attr("href", "/v1/wx/getmathpdf/" + data.pdflink);
          $('#Submit1').attr("disabled", false);
          clearInterval(time),
            function() {
              timer();
            }
          // var rows = []
          // rows.push({
          //   outputalias: data.pdflink
          // })
          // $('#table2').bootstrapTable('append', rows)
        }
      },
      error: function(data, status) {
        alert(data);
        $('#Submit1').attr("disabled", false);
        $("#pdflink").css("display", "none");
        // revertFunc();
      }
    });
    // worksheet.SetRealValue(firstInput, , "");
    // worksheet.SetRealValue(secondInput, , "in");
    // worksheet.SetRealValue(thirdInput, , "in");
    // worksheet.Synchronize();
    // var outputs = worksheet.Outputs;
    // firstOutput = outputs.GetAliasByIndex(0);
    // var val = worksheet.OutputGetStringValue(firstOutput);
  };

  //计时器
  var intDiff = parseInt(0); //倒计时总秒数量
  function timer(intDiff) {
    time = window.setInterval(function() {
      var minute = 0,
        second = 1; //时间默认值
      if (intDiff > 0) {
        // minute = Math.floor(intDiff / 60);
        second = Math.floor(intDiff);
      }
      // if (minute <= 9) minute = '0' + minute;
      if (second <= 9) second = '0' + second;
      // $('#minute').html('<s></s>' + minute + '分');
      $('#second').html('<s></s>' + second + '秒');
      intDiff++;
      // if (intDiff <= 0) {
      //   clearInterval(timer);   //定时器清除；
      //   history.back(-1);
      // }
    }, 1000);
    // window.setInterval(function () {
    //   var minute = 0,
    //     second = 0;//时间默认值
    //   if (intDiff > 0) {
    //     minute = Math.floor(intDiff / 60);
    //     second = Math.floor(intDiff) - (minute * 60);
    //   }
    //   if (minute <= 9) minute = '0' + minute;
    //   if (second <= 9) second = '0' + second;
    //   $('#minute').html('<s></s>' + minute + '分');
    //   $('#second').html('<s></s>' + second + '秒');
    //   intDiff--;
    //   if (intDiff <= 0) {
    //     clearInterval(timer);   //定时器清除；
    //     history.back(-1);
    //   }
    // }, 1000);
  }

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
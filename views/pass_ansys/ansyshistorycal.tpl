<!-- 计算历史参数记录-->
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
  <title>历史计算</title>
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
  <script type="text/javascript" src="/static/js/bootstrap-table-editable1.18.3.min.js"></script>
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
    <h3>History-输入参数表</h3>
    <div id="toolbar" class="btn-group">
      <button type="button" data-name="addButton" id="addButton" class="btn btn-default"> <i class="fa fa-plus">添加</i>
      </button>
      <button type="button" data-name="importButton" id="importButton" class="btn btn-default"> <i class="fa fa-plus">导入</i>
      </button>
      <!-- <button type="button" data-name="editorButton" id="editorButton" class="btn btn-default"> <i class="fa fa-edit">编辑</i>
        </button> -->
      <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
      </button>
    </div>
    <table id="table0" data-search="true" data-show-refresh="true" data-show-toggle="true" data-show-columns="true" data-striped="true" data-toolbar="#toolbar" data-query-params="queryParams" data-sort-name="Username" data-sort-order="desc" data-page-size="5" data-page-list="[5, 25, 50, All]" data-unique-id="ID" data-pagination="true" data-side-pagination="client" data-single-select="true" data-click-to-select="true" data-show-export="true">
    </table>
  </div>
  <div class="col-lg-6">
    <button type="button" data-name="Submit1" id="Submit1" class="btn btn-danger" onclick="oncalculateMC()"> <i class="fa fa-calculator"> 计算</i></button>
    <div class="timer">
      <strong id="minute">0分</strong>
      <strong id="second">00秒</strong>
    </div>
    <a href="" target="_blank" id="pdflink" style="display:none">PDF计算书下载</a>
  </div>
  <form role="form">
    <div class="form-group">
      <textarea id="content" class="form-control" rows="6"></textarea>
    </div>
  </form>
  <div class="col-lg-12">
    <h3>History-输出参数表</h3>
    <div id="toolbar2" class="btn-group">
      <button type="button" data-name="addButton" id="addButton" class="btn btn-default"> <i class="fa fa-plus">添加</i>
      </button>
      <button type="button" data-name="importButton" id="importButton" class="btn btn-default"> <i class="fa fa-plus">导入</i>
      </button>
      <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
      </button>
    </div>
    <table id="table2" data-search="true" data-show-refresh="true" data-show-toggle="true" data-show-columns="true" data-striped="true" data-toolbar="#toolbar2" data-query-params="queryParams" data-sort-name="Username" data-sort-order="desc" data-page-size="5" data-page-list="[5, 25, 50, All]" data-unique-id="id" data-pagination="true" data-side-pagination="client" data-single-select="true" data-click-to-select="true" data-show-export="true">
    </table>
  </div>
  <script type="text/javascript">
  // $(document).ready(function() {
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
  // })
  $('#my_table_id').on('shown', function(e, editable) {
    editable.input.$input.val('overwriting value of input..');
  });

  $(function() {
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
  })

  $('#table0').bootstrapTable({
    idField: 'ID',//这个对应editable的pk传值
    url: '/v1/mathcad/gethistoryinput/{{.HistoryID}}?templeid={{.TempleID}}',
    // striped: "true",
    columns: [{
      title: '序号',
      formatter: function(value, row, index) {
        return index + 1
      },
      halign: "center",
      align: "center",
      valign: "middle"
    },{
      field: 'inputalias',
      title: '别名',
      halign: "center",
      align: "center",
      valign: "middle"
    },{
      field: 'historyinputvalue.inputvalue',
      title: '历史数值',
      halign: "center",
      align: "center",
      valign: "middle",
      editable: {
        // noEditFormatter (value, row, index) {
        //   console.log(value)
        //   if (value === 'noEdit') {
        //     return 'No Edit'
        //   }
        //   return false
        // },
        type() {
          var pk = $(this).attr("data-pk");
          // console.log(pk)
          var selectRow3 = $('#table0').bootstrapTable('getRowByUniqueId', pk);
          // console.log(selectRow3)
          if (selectRow3.resulttype == "ComboBoxControl") {
            edittype = 'select'
          } else if(selectRow3.resulttype == "MATRIX") {
            edittype = 'textarea'
          }else {
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
            return arr
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
        title: '选择',
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
    },{
      field: 'inputvalue',
      title: '默认值',
      halign: "center",
      align: "center",
      valign: "middle",
      formatter: onMsoNumberFormat
    },{
      field: 'resulttype',
      title: '值类型',
      halign: "center",
      align: "center",
      valign: "middle"
    },{
      field: 'units',
      title: '单位',
      halign: "center",
      align: "center",
      valign: "middle"
    },{
      field: 'comment',
      title: '说明',
      halign: "center",
      align: "center",
      valign: "middle",
      editable: {
        type: 'text',
        pk: 1,
        url: '/v1/mathcad/putmathcalinput',
        title: 'Enter MathComment'
      },
      // formatter: onMsoNumberFormat
    }
    ]
  });

  $('#table2').bootstrapTable({
    uniqueId: "outputalias", //每一行的唯一标识，一般为主键列
    idField: 'ID', //可以通过种方式设置主键
    url: '/v1/mathcad/gethistoryoutput/{{.HistoryID}}?templeid={{.TempleID}}',
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
        field: 'historyoutputvalue.outputvalue',
        title: '历史数值',
        halign: "center",
        align: "center",
        valign: "middle",
        editable: {
          type: 'text',
        },
        formatter: onMsoNumberFormat
      }, {
        field: 'outputvalue',
        title: '默认值',
        halign: "center",
        align: "center",
        valign: "middle",
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
      },{
      field: 'comment',
      title: '说明',
      halign: "center",
      align: "center",
      valign: "middle",
      editable: {
        type: 'text',
        pk: 1,
        url: '/v1/mathcad/putmathcaloutput',
        title: 'Enter MathComment'
      },
    }]
  });

  function onMsoNumberFormat(value, row, index) { //科学计数法数字转换为文本，时间格式转换为文本可以显示秒
    // console.log(row.resulttype)//alert无法显示对象，应该用consle.log来显示对象
    if (row.resulttype == "Real") {
      var num = new Number(value);
      return num
    // } else if(row.resulttype == "MATRIX"){
    //   console.log(row.textarealvalue.value)
    //   return row.textarealvalue.value
    }else{
      return value
    }
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



  // $(function() {
  function oncalculateMC() {
    //两种设置disabled属性
    $('#Submit1').attr("disabled", true);
    $("#pdflink").css("display", "none");
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
      // console.log(obj);

      var data = $('#table0').bootstrapTable('getData')
      // var data2 = JSON.stringify(data)
      // data: { templeid: {{.ID }}, inputdata: data2 },
      obj.inputdata = data

      obj = JSON.stringify(obj); //将JSON对象转化为JSON字符
      ws.send(obj);

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
        for (var i = 0; i < msg.data.length; i++) {
          $('#table2').bootstrapTable('updateByUniqueId', {
            id: msg.data[i].outputalias,
            row: {
              outputvalue: msg.data[i].outputvalue
            }
          });
        }
        $("#pdflink").css("display", "block");
        $("#pdflink").attr("href", msg.pdflink);
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
          $("#pdflink").attr("href", data.pdflink);
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
        second = 0; //时间默认值
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
<!-- 文档列表 -->
<!DOCTYPE html>

<head>
  <title>出差统计</title>
  <meta name="renderer" content="webkit">
  <!-- 加上这句，360等浏览器就会默认使用google内核，而不是IE内核 。 -->
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <!-- 加上这一句，如果被用户强行使用IE浏览器，就会使用IE的最高版本渲染内核 -->
  <!-- <link type='text/css' href='/static/oo/files-3TmaoIbj3PAed78NYLoa7w2.css' rel='stylesheet' />
    <link type='text/css' href='/static/oo/common-HLDWebQ4QDcrVRYNq4-rWA2.css' rel='stylesheet' />
    <link type='text/css' href='/static/oo/files-CUBYqoHsKUGuN7k-PidXtQ2.css' rel='stylesheet' /> -->
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css" />
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css" />
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
  <script src="/static/js/tableExport.js"></script>
  <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
  <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.min.js"> </script>
  <!--建议手动加在语言，避免在ie下有时因为加载语言失败导致编辑器加载失败-->
  <!--这里加载的语言文件会覆盖你在配置项目里添加的语言类型，比如你在配置项目里配置的是英文，这里加载的中文，那最后就是中文-->
  <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/webuploader.css">
  <script type="text/javascript" src="/static/js/webuploader.min.js"></script>
  <script type="text/javascript" src="/static/js/jquery-ui.min.js"></script>
  <!-- <script type="text/javascript" src="/static/bootstrap-datepicker/bootstrap-datepicker.js"></script> -->
  <script type="text/javascript" src="/static/bootstrap-datepicker/bootstrap-datepicker.js"></script>
  <script type="text/javascript" src="/static/bootstrap-datepicker/bootstrap-datepicker.zh-CN.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/bootstrap-datepicker/bootstrap-datepicker.css" />
  <!-- <link rel="stylesheet" type="text/css" href="/static/bootstrap-datepicker/bootstrap-datepicker.min.css"/> -->
  <link rel="stylesheet" type="text/css" href="/static/css/select2.css" />
  <script type="text/javascript" src="/static/js/select2.js"></script>
  <!-- <script language="javascript" src="/static/oo/ga-teamlab.js" type="text/javascript"></script> -->
  <style type="text/css">
  #modalDialog .modal-header {
    cursor: move;
  }

  #modalDialog1 .modal-header {
    cursor: move;
  }

  #modalDialog2 .modal-header {
    cursor: move;
  }

  #modalDialog3 .modal-header {
    cursor: move;
  }

  #modalDialog4 .modal-header {
    cursor: move;
  }

  #modalDialog5 .modal-header {
    cursor: move;
  }

  #modalDialog6 .modal-header {
    cursor: move;
  }

  #modalDialog7 .modal-header {
    cursor: move;
  }

  #modalDialog8 .modal-header {
    cursor: move;
  }

  #modalDialog9 .modal-header {
    cursor: move;
  }

  #modalDialog10 .modal-header {
    cursor: move;
  }

  /*.form-group .datepicker{
        z-index: 9999;
      }*/
  /*模态框效果*/
  /*.modal-header {*/
  /*background: #00FF00;*/
  /*min-height: 16.42857143px;
      padding: 15px;
      border-bottom: 1px solid #e5e5e5;*/
  /*}*/
  /*.col-sm-1 input[type=checkbox]{
　　display: inline-block;
　　vertical-align: middle;
　　margin-bottom: 2px; 
    }*/
  </style>
</head>
<div class="container-fill">{{template "/T.navbar.tpl" .}}</div>

<body>
  <div class="col-lg-12">
    <h3>考勤列表</h3>
    <ul id="myTab" class="nav nav-tabs">
      <li class="active">
        <a href="#checkdate" data-toggle="tab">打卡日期</a>
      </li>
      <li>
        <a href="#location" data-toggle="tab">定位情况</a>
      </li>
      <li><a href="#checktime" data-toggle="tab">打卡时间</a>
      </li>
    </ul>
    <div id="toolbar" class="btn-group">
      <span style="position: relative;z-index: 9999;">
        <input type='text' placeholder='选择月份' class='datepicker btn btn-default' id='Date' value='' />
      </span>
      <span>
        <select name="activity" id='activity' class="btn btn-default" required onchange="activity()">
          <!--<option value="1">活动1</option>
      <option value="2">活动2</option>
      <option value="0">活动3</option> -->
        </select>
      </span>
    </div>

    <div id="myTabContent" class="tab-content">
      <div class="tab-pane fade in active" id="checkdate">
        <table id="table"></table>
      </div>
      <div class="tab-pane fade" id="location">
        <table id="table1"></table>
      </div>
      <div class="tab-pane fade" id="checktime">
        <table id="table2"></table>
      </div>
    </div>
    
    <script type="text/javascript">
    // *********日期转换*************
    //https://www.jb51.net/article/80599.htm
    //http://blog.sina.com.cn/s/blog_9c2960490101watm.html
    /** 
     * 字符串转时间（yyyy-MM-dd HH:mm:ss） 
     * result （分钟） 
     */
    // stringToDate : function(fDate){  
    //   var fullDate = fDate.split("-");  
    //   return new Date(fullDate[0], fullDate[1]-1, fullDate[2], 0, 0, 0);  
    // }
    // var myDate = new Date();
    // myDate.getYear();        //获取当前年份(2位)
    // myDate.getFullYear();    //获取完整的年份(4位)
    // myDate.getMonth();       //获取当前月份(0-11,0代表1月)
    // myDate.getMonth()+1;       //获取正确月份
    // *********日期转换*************

    $(document).ready(function() {
      var now = new Date();
      myDate = new Date().getFullYear() + '年' + (new Date().getMonth() + 1) + '月' + new Date().getDate() + '日';
      $("#Date").val(myDate);
      // var vs = $('select option:selected').val();
      // console.log(vs)
      // var myDate = new Date();
      // myDate.getYear();        //获取当前年份(2位)
      var b = now.getFullYear(); //获取完整的年份(4位)
      // myDate.getMonth();       //获取当前月份(0-11,0代表1月)
      var d = now.getMonth() + 1;

      $.ajax({
        type: "GET", //这里是否一定要用post？？？
        // url: "/v1/wx/getbusiness/25001",
        url: "/v1/project/getwxprojects",
        dataType: 'json', //dataType:JSON,这种是jquerylatest版本的表达方法。不支持新版jquery。
        success: function(data, status) {
          console.log(data)
          $(".option").remove();
          $.each(data.rows, function(i, d) {
            // console.log(data.processing[0])
            // console.log(i)
            // console.log(d)
            $("#activity").append('<option class="option" value="' + data.rows[i].Id + '">' + data.rows[i].Title + '</option>');
          });
          var page = $("li.page-number.active").text()
          // alert(page)
          var limit = $("span.page-size:first").text()

          $("#table").bootstrapTable({
            ajax: function(request) {
              $.ajax({
                type: "GET",
                url: "/v1/wx/businessmonthcheck3/"+data.rows[0].Id+"?&year=" + b + "&month=" + d,
                contentType: "application/json;charset=utf-8",
                dataType: "json",
                data: { page: 1, limit: 500 },
                json: 'callback',
                success: function(json) {
                  var columnsArray = [];
                  columnsArray.push({ field: "state", checkbox: true });
                  for (var i = 0; i < (Object.keys(json[0])).length; i++) { //Object.keys(obj) 获取key名称
                    var property = (Object.keys(json[0]))[i]; //id   username
                    if (property == 0) {
                      columnsArray.push({
                        "title": "name",
                        "field": property,
                        "align": 'center',
                        "valign": 'middle',
                        switchable: true,
                      });
                    } else {
                      columnsArray.push({
                        "title": property,
                        "field": property,
                        "align": 'center',
                        "valign": 'middle',
                        switchable: true,
                      });
                    }
                  }
                  $('#table').bootstrapTable('destroy').bootstrapTable({
                    data: json,
                    toolbar: '#toolbar',
                    singleSelect: true,
                    clickToSelect: true,
                    sortName: "create_time",
                    sortOrder: "desc",
                    pageSize: 10,
                    pageNumber: 1,
                    pageList: "[10, 25, 50, 100, All]",
                    showToggle: true,
                    showRefresh: true,
                    showColumns: true,
                    search: true,
                    showExport: true,
                    pagination: true,
                    striped: true,
                    // sidePagination: "server",
                    queryParams: function queryParams(params) { //设置查询参数
                      var param = {
                        limit: params.pageSize, //每页多少条数据
                        page: params.pageNumber, // 页码
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
                    columns: columnsArray
                  });
                },
                error: function() {
                  alert("错误");
                }
              });
            }
          });

        }
      });
    });

    function activity() {
      //获取被选中的option标签
      var vs = $('select option:selected').val();
      var date = $('#Date').val();
      var a = document.getElementById("Date").value;
      // console.log(a)
      // var ttDate = "2013年12月20日 14:20:20";  
      var ttDate = a.match(/\d{4}.\d{1,2}.\d{1,2}/mg).toString();
      var ttDate = ttDate.replace(/[^0-9]/mg, '-');
      // alert(ttDate);
      var myDate = new Date(ttDate);
      // var str="sfsfsfdf2011年sfsdfsf05月随碟附送lfs23日";
      // str.match(/\d{4}年[01]?\d月[0123]?\d日/g)[0];
      // var myDate=new Date(a);
      // console.log(myDate)
      var b = myDate.getFullYear(); //获取完整的年份(4位)
      // var c=myDate.getMonth();       //获取当前月份(0-11,0代表1月)
      var d = myDate.getMonth() + 1;
      // console.log(b)
      // console.log(c)
      // console.log(d)
      $.ajax({
        type: "GET",
        url: "/v1/wx/businessmonthcheck3/"+vs+"?year=" + b + "&month=" + d,
        contentType: "application/json;charset=utf-8",
        dataType: "json",
        data: { page: 1, limit: 500 },
        json: 'callback',
        success: function(json) {
          var columnsArray = [];
          columnsArray.push({ field: "state", checkbox: true });
          if (json[0]) {
            for (var i = 0; i < (Object.keys(json[0])).length; i++) { //Object.keys(obj) 获取key名称
              var property = (Object.keys(json[0]))[i]; //id username
              if (property == 0) {
                columnsArray.push({
                  "title": "name",
                  "field": property,
                  "align": 'center',
                  "valign": 'middle',
                  switchable: true,
                });
              } else {
                columnsArray.push({
                  "title": property,
                  "field": property,
                  "align": 'center',
                  "valign": 'middle',
                  switchable: true,
                });
              }
            }
          }
          $('#table').bootstrapTable('destroy').bootstrapTable({
            data: json,
            toolbar: '#toolbar',
            singleSelect: true,
            clickToSelect: true,
            sortName: "create_time",
            sortOrder: "desc",
            pageSize: 10,
            pageNumber: 1,
            pageList: "[10, 25, 50, 100, All]",
            showToggle: true,
            showRefresh: true,
            showColumns: true,
            showExport: true,
            search: true,
            pagination: true,
            striped: true,
            // sidePagination: "server",
            queryParams: function queryParams(params) { //设置查询参数
              var param = {
                limit: params.pageSize, //每页多少条数据
                page: params.pageNumber, // 页码
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
            columns: columnsArray
          });
        },
        rror: function() {
          alert("错误");
        }
      });
    }

    $("#Date").datepicker({
      startView: 1,
      minViewMode: 1,
      maxViewMode: 2,
      // weekStart: 1,
      language: "zh-CN",
      autoclose: true, //选中之后自动隐藏日期选择框
      altFormat: "yy-mm",
    })

    $('#Date').datepicker().on('hide', function(e) {
      activity()
    });

    function getColumns() {//试验
      // 加载动态表格
      $.ajax({
        url: '/v1/wx/businessmonthcheck3?year=2019&month=04',
        type: 'get',
        dataType: "json",
        async: false,
        success: function(returnValue) {
          // 未查询到相应的列，展示默认列
          if (returnValue.retCode == "0") {
            //没查到列的时候把之前的列再给它
            myColumns = $table.bootstrapTable('getOptions').columns[0];
          } else {
            // 异步获取要动态生成的列
            var arr = returnValue.data;
            $.each(arr, function(i, item) {
              myColumns.push({
                "field": item.labelColumnCode,
                "title": item.labelColumnName,
                "hide": true,
                "align": 'center',
                "valign": 'middle'
              });
            });
          }
          // console.log(myColumns);
          return myColumns;
        }
      });
    }
    // $table.bootstrapTable("refreshOptions", {  
    //   url : path + "/api/peopledataInfo/getPeopleInfoList", // 获取数据的地址
    //   columns : myColumns,
    // });
    </script>
  </div>
</body>

</html>
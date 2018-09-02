<!-- 项目列表页 后端分页-->
<!DOCTYPE html>
{{template "header"}}
<title>项目列表-EngiCMS</title>
<!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/> -->
<!-- <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script> -->
  <!-- <script type="text/javascript" src="/static/js/bootstrap.min.js"></script> -->
  <script src="/static/js/bootstrap-treeview.js"></script>
  <!-- <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script> -->
  
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css"/> -->
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <!-- <script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script> -->
  <!-- <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script> -->
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/font-awesome.min.css"/> -->
  <script src="/static/js/tableExport.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
</head>

<!-- <div class="navbar navbar-default navbar-static-top"> -->
  <div class="container-fill">{{template "navbar" .}}</div>
<!-- </div> -->
<!-- <nav class="navbar navbar-default navbar-static-top">
  <div class="container">
    ...
  </div>
</nav> -->

<body>
<div class="col-lg-12">
  <h3>项目列表</h3>
<div id="toolbar1" class="btn-group">
        <button type="button" id="addButton" class="btn btn-default"> <i class="fa fa-plus">添加</i>
        </button>
        <button type="button" id="editorButton" class="btn btn-default"> <i class="fa fa-edit">编辑</i>
        </button>
        <button type="button" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
        </button>
</div>

<table id="table0"></table>
        <!-- data-url="/project/getprojects"
        data-search="true"
        data-show-refresh="true"
        data-show-toggle="true"
        data-show-columns="true"
        data-toolbar="#toolbar1"
        data-query-params="queryParams"
        data-sort-name="Created"
        data-sort-order="desc"
        data-page-size="15"
        data-page-list="[15,20, 50, 100, All]"
        data-unique-id="id"
        data-pagination="true"
        data-side-pagination="server"
        data-single-select="true"
        data-click-to-select="true" -->
    <!-- <thead>        
      <tr>
        <th data-width="10" data-radio="true" data-align="center" data-valign="middle"></th>
        <th data-formatter="index1" data-align="center" data-valign="middle">#</th>
        <th data-field="Code" data-formatter="setCode" data-align="center" data-valign="middle">编号</th>
        <th data-field="Title" data-formatter="setTitle" data-align="center" data-valign="middle">名称</th>
        <th data-field="Label" data-formatter="setLable" data-align="center" data-valign="middle">标签</th>
        <th data-field="Principal" data-align="center" data-valign="middle">负责人</th>
        <th data-field="Number" data-title-tooltip="A=B=C=D=E=F=G=" data-align="center" data-valign="middle">成果数量</th>
        <th data-field="Number" data-formatter="actionFormatter0" data-title-tooltip="A=B=C=D=E=F=G=" data-align="center" data-valign="middle">阶段</th>
        <th data-field="Number" data-formatter="actionFormatter1" data-title-tooltip="A=B=C=D=E=F=G=" data-align="center" data-valign="middle">文件</th>
        <th data-field="Number" data-formatter="actionFormatter2" data-title-tooltip="A=B=C=D=E=F=G=" data-align="center" data-valign="middle">团队</th>
        <th data-field="Number" data-formatter="actionFormatter3" data-title-tooltip="本项目合同" data-align="center" data-valign="middle">合同</th>
        <th data-field="action" data-formatter="actionFormatter" data-events="actionEvents" data-align="center" data-valign="middle">时间轴</th>
        <th data-field="Created" data-formatter="localDateFormatter" data-align="center" data-valign="middle">建立时间</th>
      </tr>
    </thead> -->

<!-- <div class="gridview2"></div> -->

<script type="text/javascript">
  /*数据json*/
  var json =
     [{"Id":"1","Code":"SL0001","Title":"***水利枢纽工程","Label":"水电站,枢纽","Principal":"秦晓川","Product":"8","Created":"2016-11-26"},
      {"Id":"2","Code":"SL0002","Title":"***电力工程","Label":"输电,高压","Principal":"秦晓川","Product":"8","Created":"2016-11-26"},
      {"Id":"3","Code":"SL0003","Title":"***市政工程","Label":"管网,给水","Principal":"秦晓川","Product":"8","Created":"2016-11-26"},
      {"Id":"4","Code":"SL0004","Title":"***建筑工程","Label":"写字楼,高层","Principal":"秦晓川","Product":"8","Created":"2016-11-26"},
      {"Id":"5","Code":"SL0005","Title":"***交通工程","Label":"进场路,一级","Principal":"秦晓川","Product":"8","Created":"2016-11-26"},
      {"Id":"6","Code":"SL0006","Title":"***境外工程","Label":"水电站,枢纽","Principal":"秦晓川","Product":"8","Created":"2016-11-26"}];
  /*初始化table数据*/
  // $(function(){
  //     $("#table0").bootstrapTable({
  //         data:json
  //         // onClickRow: function (row, $element) {
  //           // alert( "选择了行Id为: " + row.Id );
  //           // rowid=row.Id//全局变量
  //           // $('#table1').bootstrapTable('refresh', {url:'/admincategory?pid='+row.Id});
  //         // }
  //     });
  // });
  //项目列表
  $(function () {
    // 初始化【未接受】工作流表格
    $("#table0").bootstrapTable({
        url : '/project/getprojects',
        method: 'get',
        search:'true',
        showRefresh:'true',
        showToggle:'true',
        showColumns:'true',
        toolbar:'#toolbar1',
        pagination: 'true',
        sidePagination: "server",
        queryParamsType:'',
        //请求服务器数据时，你可以通过重写参数的方式添加一些额外的参数，例如 toolbar 中的参数 如果 queryParamsType = 'limit' ,返回参数必须包含
        // limit, offset, search, sort, order 否则, 需要包含: 
        // pageSize, pageNumber, searchText, sortName, sortOrder. 
        // 返回false将会终止请求。
        pageSize: 15,
        pageNumber: 1,
        pageList: [15,20, 50, 100],
        singleSelect:"true",
        clickToSelect:"true",
        queryParams:function queryParams(params) {   //设置查询参数
          var param = {
              limit: params.pageSize,   //每页多少条数据
              pageNo: params.pageNumber, // 页码
              searchText:$(".search .form-control").val()
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
            title: '选择',
            radio: 'true',
            width: '10',
            align:"center",
            valign:"middle"
          },
          {
            // field: 'Number',
            title: '序号',
            formatter:function(value,row,index){
              return index+1
            },
            align:"center",
            valign:"middle"
          },
          {
            field: 'Code',
            title: '编号',
            formatter:setCode,
            align:"center",
            valign:"middle"
          },
          {
            field: 'Title',
            title: '名称',
            formatter:setTitle,
            align:"center",
            valign:"middle"
          },
          {
            field: 'Label',
            title: '标签',
            formatter:setLable,
            align:"center",
            valign:"middle"
          },
          {
            field: 'Principal',
            title: '负责人',
            align:"center",
            valign:"middle"
          },
          {
            field: 'Number',
            title: '成果数量',
            formatter:setCode,
            align:"center",
            valign:"middle"
          },
          {
            field: 'action',
            title: '时间轴',
            formatter:actionFormatter,
            events:actionEvents,
            align:"center",
            valign:"middle"
          },
          {
            field: 'Created',
            title: '建立时间',
            formatter:localDateFormatter,
            align:"center",
            valign:"middle"
          }
            // {
            //     field: 'dContMainEntity.createTime',
            //     title: '发起时间',
            //     formatter: function (value, row, index) {
            //         return new Date(value).toLocaleString().substring(0,9);
            //     }
            // },
            // {
            //     field: 'dContMainEntity.operate',
            //     title: '操作',
            //     formatter: operateFormatter
            // }
        ]
    });
  });

  //项目模板中的项目列表
  $(function () {
    // 初始化【未接受】工作流表格
    $("#table2").bootstrapTable({
        url : '/project/getprojects',
        // method: 'get',
        search:'true',
        showRefresh:'true',
        showToggle:'true',
        showColumns:'true',
        // toolbar:'#toolbar1',
        pagination: 'true',
        sidePagination: "server",
        queryParamsType:'',
        //请求服务器数据时，你可以通过重写参数的方式添加一些额外的参数，例如 toolbar 中的参数 如果 queryParamsType = 'limit' ,返回参数必须包含
        // limit, offset, search, sort, order 否则, 需要包含: 
        // pageSize, pageNumber, searchText, sortName, sortOrder. 
        // 返回false将会终止请求。
        pageSize: 10,
        pageNumber: 1,
        pageList: [10,20, 50, 100],
        singleSelect:"true",
        clickToSelect:"true",
        queryParams:function queryParams(params) {   //设置查询参数
          var param = {
              limit: params.pageSize,   //每页多少条数据
              pageNo: params.pageNumber, // 页码
              searchText:$(".search .form-control").val()
          };
          //搜索框功能
            return param;
        },
        columns: [
          {
            title: '选择',
            radio: 'true',
            width: '10',
            align:"center",
            valign:"middle"
          },
          {
            title: '序号',
            formatter:function(value,row,index){
              return index+1
            },
            align:"center",
            valign:"middle"
          },
          {
            field: 'Code',
            title: '编号',
            formatter:setCode,
            align:"center",
            valign:"middle"
          },
          {
            field: 'Title',
            title: '名称',
            formatter:setTitle,
            align:"center",
            valign:"middle"
          }
        ]
    });
  });

  function index1(value,row,index){
    // alert( "Data Loaded: " + index );
    return index+1
  }

  //专业快捷进入

  //阶段快捷进入{font-weight:bold
  function actionFormatter0(value, row, index) {
    return [
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';" title="规划">',
        'A</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';" title="项建">',        
        'B</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';" title="可研">',
        'C</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';" title="初设">',
        'D</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';" title="招标">',
        'E</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';" title="施工图">',
        'F</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';" title="竣工图">',
        'J</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';" title="其他：文章、图纸……">',
        '<i class="fa fa-ellipsis-h"></i></a>'
    ].join('');
  }
  //文件快捷进入
  function actionFormatter1(value, row, index) {
    return [
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';" title="批文">',
        'A</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';" title="日志">',        
        'B</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';" title="强条">',
        'C</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';" title="大纲">',
        'D</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';" title="报告">',
        'E</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';" title="综合说明">',
        'F</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';" title="总布置图">',
        'G</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';" title="项目简介">',
        'H</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';" title="宣传动画">',
        'I</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';" title="其他：文章、图纸、PPT……">',
        '<i class="fa fa-ellipsis-h"></i></a>'
    ].join('');
  }
  //成员
  function actionFormatter2(value, row, index) {
    return [
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Georgia',', serif;color:',getRandomColor(),';" title="设总">',
        '1</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Georgia',', serif;color:',getRandomColor(),';" title="规划">',        
        '2</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Georgia',', serif;color:',getRandomColor(),';" title="地质">',
        '3</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Georgia',', serif;color:',getRandomColor(),';" title="">',
        '4</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Georgia',', serif;color:',getRandomColor(),';" title="水工">',
        '5</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Georgia',', serif;color:',getRandomColor(),';" title="机电">',
        '6</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Georgia',', serif;color:',getRandomColor(),';" title="施工">',
        '7</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Georgia',', serif;color:',getRandomColor(),';" title="征地">',
        '8</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Georgia',', serif;color:',getRandomColor(),';" title="水保">',
        '9</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-family:','Georgia',', serif;color:',getRandomColor(),';" title="其他：预算、环保……">',
        '<i class="fa fa-ellipsis-h"></i></a>'
    ].join('');
  }

  //合同
  function actionFormatter3(value, row, index) {
    return [
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';"title="收款合同">',
        'C</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';"title="付款合同">',        
        'P</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';"title="合同进度">',
        'P</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';"title="收款进度">',
        'GP</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-weight:bold;font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';"title="付款进度">',
        'PP</a>&nbsp;&nbsp;',
        '<a href="javascript:void(0)" style="font-family:','Comic Sans MS',',cursive;color:',getRandomColor(),';"title="其他：……">',
        '<i class="fa fa-ellipsis-h"></i></a>'
    ].join('');
  }
  //详细，提交，删除
  function actionFormatter(value, row, index) {
    return [
        '<a class="memorabilia" href="javascript:void(0)" title="大事记">',
        '&nbsp;&nbsp;<i class="fa fa-ellipsis-v"></i>&nbsp;&nbsp;</a>&nbsp;&nbsp;',
        '<a class="log" href="javascript:void(0)" title="日志">',
        '<i class="fa fa-calendar"></i>',
        '</a>&nbsp;&nbsp;',
        '<a class="operate" href="javascript:void(0)" title="操作记录">',
        '<i id="delete" class="fa fa-wrench"></i>',
        '</a>'
    ].join('');
  }
  window.actionEvents = {
    //大事记
    'click .memorabilia': function (e, value, row, index) {
      window.open('/project/'+row.Id+'/gettimeline');
    },
    //日志
    'click .log': function (e, value, row, index) {
      window.open('/project/'+row.Id+'/getcalendar');
    },
    //操作记录
    'click .operate': function (e, value, row, index) {
        // alert('You click remove icon, row: ' + JSON.stringify(row));
        // console.log(value, row, index);
    }
  };
  function localDateFormatter(value) {
    return moment(value, 'YYYY-MM-DD').format('YYYY-MM-DD');
  }
  function setCode(value,row,index){
    return "<a href='/project/"+row.Id+"'>" + value + "</a>";
  }
  function setTitle(value,row,index){
    return "<a href='/project/allproducts/"+row.Id+"'>" + value + "</a>";
  }
  function setLable(value,row,index){
    if (value){
      array=value.split(",")
      var labelarray = new Array() 
      for (i=0;i<array.length;i++)
      {
        labelarray[i]="<a href='/project/keysearch?keyword="+array[i]+"'>" + array[i] + "</a>";
      }
      return labelarray.join(",");
    }
  } 
  // 改变点击行颜色
  $(function(){
     // $("#table").bootstrapTable('destroy').bootstrapTable({
     //     columns:columns,
     //     data:json
     // });
     $("#table0").on("click-row.bs.table",function(e,row,ele){
         $(".info").removeClass("info");
         $(ele).addClass("info");
         // rowid=row.Id//全局变量
         // $('#table1').bootstrapTable('refresh', {url:'/admin/category/'+row.Id});
     });
     // $("#get").click(function(){
     //     alert("商品名称：" + getContent().TuanGouName);
     // })
  });
  // $(function () {
  //     var $result = $('#eventsResult');
  //     var selectRow=$('#table').bootstrapTable('getSelections');
  
  //     $('#table').on('all.bs.table', function (e, name, args) {
  //         console.log('Event:', name, ', data:', args);
  //     })
  //     .on('click-row.bs.table', function (e, row, $element) {
  //       alert("选择！"+row.Id);
  //       if (selectRow.length<1){
  //         selectRow=$('#table').bootstrapTable('getSelections');
  //         alert("请选择"+selectRow.length);
  //         // return;
  //         }
  //         $result.text('Event: click-row.bs.table');
  //     })
  //     .on('dbl-click-row.bs.table', function (e, row, $element) {
  //         $result.text('Event: dbl-click-row.bs.table');
  //     })
  //     .on('sort.bs.table', function (e, name, order) {
  //         $result.text('Event: sort.bs.table');
  //     })
  //     .on('check.bs.table', function (e, row) {
  //         $result.text('Event: check.bs.table');
  //     })
  //     .on('uncheck.bs.table', function (e, row) {
  //         $result.text('Event: uncheck.bs.table');
  //     })
  //     .on('check-all.bs.table', function (e) {
  //         $result.text('Event: check-all.bs.table');
  //     })
  //     .on('uncheck-all.bs.table', function (e) {
  //         $result.text('Event: uncheck-all.bs.table');
  //     })
  //     .on('load-success.bs.table', function (e, data) {
  //         $result.text('Event: load-success.bs.table');
  //     })
  //     .on('load-error.bs.table', function (e, status) {
  //         $result.text('Event: load-error.bs.table');
  //     })
  //     .on('column-switch.bs.table', function (e, field, checked) {
  //         $result.text('Event: column-switch.bs.table');
  //     })
  //     .on('page-change.bs.table', function (e, number, size) {
  //         $result.text('Event: page-change.bs.table');
  //     })
  //     .on('search.bs.table', function (e, text) {
  //         $result.text('Event: search.bs.table');
  //     });
  // });


  // $(document).ready(function() {
  // $("#addButton").click(function() {
    
  // var selectRow=$('#table').bootstrapTable('getSelections');  
  // if (selectRow.length<1){
  // selectRow=$('#table').bootstrapTable('getSelections');
  // alert("请选择"+selectRow.length);
  // return;
  // }
        // $('#modalTable').modal({
        // show:true,
        // backdrop:'static'
        // });
    // })
  // })

  $(document).ready(function() {
    $("#addButton").click(function() {
      if ({{.role}}!=1){
        alert("权限不够！");
        return;
      }
      $("label#info").remove();
      $("#saveproj").removeClass("disabled")
      $('#modalTable').modal({
      show:true,
      backdrop:'static'
      });
    })

    $("#editorButton").click(function() {
      if ({{.role}}!=1){
        alert("权限不够！");
        return;
      }
      var selectRow=$('#table0').bootstrapTable('getSelections');
      if (selectRow.length<1){
        alert("请先勾选类别！");
        return;
      }
      if (selectRow.length>1){
        alert("请不要勾选一个以上！");
        return;
      }
      $("input#cid").remove();
      var th1="<input id='cid' type='hidden' name='cid' value='" +selectRow[0].Id+"'/>"
      $(".modal-body").append(th1);//这里是否要换名字$("p").remove();
      $("#projcode1").val(selectRow[0].Code);
      $("#projname1").val(selectRow[0].Title);
      $("#projlabel1").val(selectRow[0].Label);
      $("#projprincipal1").val(selectRow[0].Principal);

        $('#modalTable1').modal({
        show:true,
        backdrop:'static'
        });
    })

    // 删除项目
    $("#deleteButton").click(function() {
      if ({{.role}}!=1){
        alert("权限不够！");
        return;
      }
      
      var selectRow=$('#table0').bootstrapTable('getSelections');
      if (selectRow.length<=0) {
        alert("请先勾选项目！");
        return false;
      }
      if(confirm("确定删除项目吗？第一次提示！")){
      }else{
        return false;
      }
      if(confirm("确定删除项目吗？第二次提示！")){
      }else{
        return false;
      }
      if(confirm("确定删除项目吗？一旦删除将无法恢复！")){
        var title=$.map(selectRow,function(row){
          return row.Title;
        })
        var ids="";
        for(var i=0;i<selectRow.length;i++){
          if(i==0){
            ids=selectRow[i].Id;
          }else{
            ids=ids+","+selectRow[i].Id;
          }  
        }
        $.ajax({
          type:"post",
          url:"/project/deleteproject",
          data: {ids:ids},
          success:function(data,status){
            alert("删除“"+data+"”成功！(status:"+status+".)");
            //删除已选数据
            $('#table0').bootstrapTable('remove',{
              field:'Title',
              values:title
            });
          }
        });
      }  
    })
  })

   /*数据json*/
  var json1 = [{"Id":"1","Title":"规划","Code":"A","Grade":"1"},
              {"Id":"7","Title":"可研","Code":"B","Grade":"1"},
              {"Id":"2","Title":"报告","Code":"B","Grade":"2"},
              {"Id":"3","Title":"图纸","Code":"T","Grade":"2"},
              {"Id":"4","Title":"水工","Code":"5","Grade":"3"},
              {"Id":"5","Title":"机电","Code":"6","Grade":"3"},
              {"Id":"6","Title":"施工","Code":"7","Grade":"3"}];
  /*初始化table数据*/
  $(function(){
      $("#table1").bootstrapTable({
          data:json1
      });
  });

  //填充select选项
  $(document).ready(function(){
    //   $(array).each(function(index){
    //     alert(this);
    // });
    // $.each(array,function(index){
    //     alert(this);
    // });
    $("#admincategory").append('<option value="a">项目模板</option>');
    if ({{.Select2}}){//20171021从meirit修改而来
      $.each({{.Select2}},function(i,d){
      // alert(this);
      // alert(i);
      // alert(d);
        $("#admincategory").append('<option value="' + i + '">'+d+'</option>');
      });
    }
  });
  
  //根据选择，刷新表格
  function refreshtable(){
      var newcategory = $("#admincategory option:selected").text();
      // alert("你选的是"+newcategory);
      if (newcategory=="项目模板"){
        //根据名字，查到id，或者另外写个categoryname的方法
        // $('#table2').bootstrapTable('refresh', {url:'/admin/categorytitle?title='+newcategory});
        $("#details").hide();
        $("#details1").show();
        // $('#table2').bootstrapTable('refresh', {url:'/project/getprojects'});
      }else{
        //根据名字，查到id，或者另外写个categoryname的方法
        $('#table1').bootstrapTable('refresh', {url:'/admin/categorytitle?title='+newcategory});
        $("#details1").hide();
        $("#details").show();
      }
  }

  //保存
  function save(){
    var projcode = $('#projcode').val();
    var projname = $('#projname').val();
    if (projname&&projcode){
      var newcategory = $("#admincategory option:selected").text();
      // alert("你选的是"+newcategory);
      // var radio =$("input[type='radio']:checked").val();
      var projlabel = $('#projlabel').val();
      var projprincipal = $('#projprincipal').val();
      if (newcategory=="项目模板"){
        var selectRow3=$('#table2').bootstrapTable('getSelections');
        if (selectRow3.length<1){
          alert("请先勾选项目！");
          return;
        }
        var ispermission=document.getElementById("ispermission").checked;
        var lab="<label id='info'>建立项目中，请耐心等待几秒/分钟……</label>"
        $(".modal-footer").prepend(lab);//这里是否要换名字$("p").remove();
        $("#saveproj").addClass("disabled");
        $.ajax({
          type:"post",
          url:"/project/addprojtemplet",
          data: {code:projcode,name:projname,label:projlabel,principal:projprincipal,projid:selectRow3[0].Id,ispermission:ispermission},//
          success:function(data,status){
            alert("添加“"+data+"”成功！(status:"+status+".)");
            //按确定后再刷新
            $('#modalTable').modal('hide');
            $('#table0').bootstrapTable('refresh', {url:'/project/getprojects'});
          }
        });
      }else{
        var selectRow3=$('#table1').bootstrapTable('getSelections');
        if (selectRow3.length<1){
          alert("请先勾选目录！");
          return;
        }
        var lab="<label id='info'>建立项目中，请耐心等待几秒/分钟……</label>"
        $(".modal-footer").prepend(lab);//这里是否要换名字$("p").remove();
        $("#saveproj").addClass("disabled");
        var ids="";
        for(var i=0;i<selectRow3.length;i++){
          if(i==0){
            ids=selectRow3[i].Id;
          }else{
            ids=ids+","+selectRow3[i].Id;
          }  
        }
        // $('#myModal').on('hide.bs.modal', function () {  
        $.ajax({
          type:"post",
          url:"/project/addproject",
          data: {code:projcode,name:projname,label:projlabel,principal:projprincipal,ids:ids},//
          success:function(data,status){
            alert("添加“"+data+"”成功！(status:"+status+".)");
            //按确定后再刷新
            $('#modalTable').modal('hide');
            $('#table0').bootstrapTable('refresh', {url:'/project/getprojects'});
          }
        });   
      }
    }else{
      alert("请填写编号和名称！");
      return;
    }
    // $(function(){$('#myModal').modal('hide')}); 
    // "/category/modifyfrm?cid="+cid
    // window.location.reload();//刷新页面
  }

  //修改项目
  function update(){
      // var radio =$("input[type='radio']:checked").val();
      var pid = $('#cid').val();
      var projcode = $('#projcode1').val();
      var projname = $('#projname1').val();
      var projlabel = $('#projlabel1').val();
      var projprincipal = $('#projprincipal1').val();

      if (projname&&projcode)
        {  
          $.ajax({
            type:"post",
            url:"/project/updateproject",
            data: {code:projcode,name:projname,label:projlabel,principal:projprincipal,pid:pid},//
            success:function(data,status){
              alert("修改“"+data+"”成功！(status:"+status+".)");
              //按确定后再刷新
              $('#modalTable1').modal('hide');
              $('#table0').bootstrapTable('refresh', {url:'/project/getprojects'});
            }
          });  
        }else{
          alert("请填写编号和名称！");
          return;
        } 
  }
  //模态框可拖曳—要引入ui-jquery.js
  // $("#modalTable").draggable({
  //   handle:".modal-header",
  //   cusor:"move",
  //   refreshPositions:false
  // });
  // 来自群，保留，批量
  // var rows= $('#account-table').bootstrapTable('getSelections');
  //       if(rows.length==0) {
  //           layer.alert('请您选择要删除的子账号！', {
  //               title:'提示信息',
  //               closeBtn: 0,
  //               icon: 0,
  //               skin: 'layui-layer-lan',
  //               shift:0 //动画类型
  //           });
  //           return false;
  //       }
  //           var ids="";
  //           for(var i=0;i<rows.length;i++){
  //               if(i==0){
  //                   ids=rows[i].frontUserId;
  //               }else{
  //                   ids=ids+","+rows[i].frontUserId;
  //               }
  //           }
</script>
  <!-- 添加项目 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalTable">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">添加项目</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">编号</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="projcode" name="projcode"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">名称</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="projname" name="projname"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">标签</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="projlabel" name="projlabel"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">负责人</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="projprincipal" name="projprincipal"></div>
              </div>
              <!-- <div class="form-group">
                <label class="col-sm-3 control-label">标签</label>
                <div class="col-sm-7">
                  <input type="number" class="form-control digits" name="label" maxlength="20" placeholder="至多20个字符" required></div>
              </div> -->
              <!-- <div class="form-group must">
                <label class="col-sm-3 control-label">负责人</label>
                  <div class="col-sm-7">
                    <input type="password" class="form-control" name="password" id="password" maxlength="32" placeholder="至多32个字符" required></div>
              </div> -->

            <!-- <div class="form-group must">
              <label class="col-sm-3 control-label">确认密码</label>
              <div class="col-sm-7">
                <input type="password" class="form-control equalto" name="password2" maxlength="32" placeholder="至多32个字符" required data-rule-equalto="#password" data-msg-equalto="密码不一致"></div>
            </div> -->
              <div class="form-group must">
                <label class="col-sm-3 control-label">类别</label>
                <div class="col-sm-7">
                  <select name="admincategory" id="admincategory" class="form-control" required onchange="refreshtable()">
                    <option>选择类别：</option>
                    <!-- <option value="1">SL</option>
                    <option value="2">DL</option>
                    <option value="0">SZ</option> -->
                  </select>
                </div>
              </div>
            </div>

            <div id="details" style="display:none">
              <h3>工程目录分级</h3>
              <table id="table1"
                    data-page-size="5"
                    data-page-list="[5, 25, 50, All]"
                    data-unique-id="id"
                    data-sort-name="Grade"
                    data-pagination="true"
                    data-side-pagination="client"
                    data-click-to-select="true">
                  <thead>        
                  <tr>
                    <th data-width="10" data-checkbox="true"></th>
                    <th data-formatter="index1">#</th>
                    <th data-field="Title">名称</th>
                    <th data-field="Code">代码</th>
                    <th data-field="Grade" data-sortable="true">级别</th>
                  </tr>
                </thead>
              </table>
            </div>

            <div id="details1" style="display:none">
              <h3>项目模板</h3>
                <div class="col-sm-3 checkbox">
                  <label><input type="checkbox" checked="checked" value="true" id="ispermission">权限继承</label>
                </div>
              <table id="table2"></table>
            </div>

          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
            <button type="button" class="btn btn-primary" id="saveproj" onclick="save()">保存</button><!--  style="display:none" -->
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- 编辑项目 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalTable1">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">编辑项目</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">编号</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="projcode1" name="projcode"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">名称</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="projname1" name="projname"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">标签</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="projlabel1" name="projlabel"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">负责人</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="projprincipal1" name="projprincipal"></div>
              </div>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
            <button type="button" class="btn btn-primary" id="updateproj" onclick="update()">更新</button>
          </div>
        </div>
      </div>
    </div>
  </div>

</div>

<script type="text/javascript"> 
  // $(function(){ 
  //   $("ul li").each(function(){ 
  //     $(this).css("background-color",getRandomColor()); 
  //   }); 
  // }) 
  function getRandomColor(){ 
    var c = '#'; 
    var cArray = ['0','1','2','3','4','5','6','7','8','9','A','B','C','D','E','F']; 
    for(var i = 0; i < 6;i++){ 
      var cIndex = Math.round(Math.random()*15); 
      c += cArray[cIndex]; 
    } 
    return c; 
  } 
</script>
<!-- <div>
  <ul>
    <li>第一个色块</li>
    <li>第二个色块</li>
    <li>第三个色块</li>
    <li>第四个色块</li>
  </ul>
</div> -->
</body>
</html>
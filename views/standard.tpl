 <!DOCTYPE html>
{{template "header"}}
<title>查阅规范、图集、计算书</title>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  <script src="/static/js/tableExport.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/webuploader.css">
  <script type="text/javascript" src="/static/js/webuploader.min.js"></script>
  <script type="text/javascript" src="/static/js/jquery-ui.min.js"></script>
 
</head>

<body>
<div class="navbar navba-default navbar-fixed-top">
  <div class="container-fill">{{template "navbar" .}}</div>
</div>


<div class="text-center">
  <h1 > <i class="glyphicon glyphicon-chevron-right"></i> <i class="glyphicon glyphicon-minus"></i>
  </h1>
  <h1 >搜索 {{.Length}}个 文件</h1>
  <p class="large">
    工程师在设计过程中接触最多的是规范、图集和计算书。
  </p>
  <p class="large">
    规范与计算本应是配套的，因国情特殊，二者未结合，造成失误、人力浪费。
  </p>
  <p class="large">
  充分利用这个平台查阅和共享它们<a href="#" id="about"><i class="glyphicon glyphicon-question-sign"></i></a>。
  <a href="#" id="valid" title="管理有效版本库">《规范目录有效版本》</a>数据来自OA系统。</p>

  <div class="col-lg-4">
  </div>
  
  <div class="col-lg-4">
  <!-- <form >   form支持回车，但是不支持ajax，如何做到支持ajax？用ajaxform-->
    <div class="input-group">
      <input type="text" class="form-control" placeholder="请输入关键字或编号进行搜索" name="name" autocomplete="off" size="30" id="name" onkeypress="getKey();">
      <span class="input-group-btn">
        <button class="btn btn-default" type="button" id="search">
          <i class="glyphicon glyphicon-search"></i>
          Search!
        </button>
      </span>
    </div>
  </div>

    <div class="col-lg-12">
    <br>
    <!--SWF在初始化的时候指定，在后面将展示-->
    <div id="uploader" class="wu-example">
      <!--用来存放文件信息-->
      <div id="thelist" class="uploader-list"></div>
      <div id="picker"><i class="glyphicon glyphicon-plus-sign"></i>选择文件</div>
      <button id="ctlBtn" class="btn btn-default"><i class="glyphicon glyphicon-upload"></i></button>
    </div>
    <br>
  </div>
 </div>

  <div class="col-lg-12">
  <!-- 规范查询结果表 -->
  <toolbar id="btn_toolbar" class="toolbar">
    <div class="btn-group">
        <button type="button" id="addButton" class="btn btn-default" data-target="modal"><i class="fa fa-plus" aria-hidden="true"></i>添加</button>
        <button type="button" id="editorButton" class="btn btn-default" data-target="modal"><i class="fa fa-edit" aria-hidden="true"></i>编辑</button>
        <button type="button" id="deleteButton" class="btn btn-default" data-target="default"><i class="fa fa-trash" aria-hidden="true"></i>删除</button>
        <button type="button" id="importButton" class="btn btn-default" data-target="default"><i class="fa fa-trash" aria-hidden="true" title="导入规范数据"></i>导入</button>
    </div>
  </toolbar>
  <!-- data-query-params="queryParams" data-content-type="application/json"-->
  <div id="details" style="display:none">
    <h3>查询结果</h3>
  <!-- data-url="/admin/category/2" 没有了这个，当然table1表格无法支持刷新了！！！data-show-refresh="true"-->
    <table id="table"
          data-toggle="table"
          data-search="true"
          data-show-toggle="true"
          data-show-columns="true"
          data-toolbar="#btn_toolbar"
          data-sort-name="Grade"
          data-page-size="20"
          data-page-list="[15, 25, 50, All]"
          data-unique-id="id"
          data-pagination="true"
          data-side-pagination="client"
          data-click-to-select="false"
          data-halign="center"
          data-striped="true"
          >
      <thead>        
        <tr>
          <th data-width="10" data-checkbox="true"></th>
          <th data-formatter="index1">#</th>
          <th data-field="Number">编号</th>
          <th data-field="Title">名称</th>
          <th data-field="Route" data-formatter="setLink">链接</th>
          <th data-field="Uname">上传者</th>
          <th data-field="LiNumber" data-formatter="setLable">有效版本库</th>
        </tr>
      </thead>
    </table>
  </div>

  <!-- 有效版本库管理表 -->
  <toolbar id="btn_toolbar1" class="toolbar">
    <div class="btn-group">
        <button type="button" id="addButton1" class="btn btn-default" data-target="modal"><i class="fa fa-plus" aria-hidden="true"></i>添加</button>
        <button type="button" id="editorButton1" class="btn btn-default" data-target="modal"><i class="fa fa-edit" aria-hidden="true"></i>编辑</button>
        <button type="button" id="deleteButton1" class="btn btn-default" data-target="default"><i class="fa fa-trash" aria-hidden="true"></i>删除</button>
        <button type="button" id="importButton1" class="btn btn-default" data-target="default"><i class="fa fa-trash" aria-hidden="true" title="导入有效库数据"></i>导入</button>
    </div>
  </toolbar>
  <div id="details2" style="display:none">
    <h3>管理有效版本库</h3>
    <table id="table1"
          data-toggle="table"
          data-search="true"
          data-show-toggle="true"
          data-show-columns="true"
          data-toolbar="#btn_toolbar1"
          data-sort-name=""
          data-page-size="20"
          data-page-list="[15, 25, 50, All]"
          data-unique-id="id"
          data-pagination="true"
          data-side-pagination="client"
          data-click-to-select="false"
          data-halign="center"
          data-striped="true"
          >
      <thead>        
        <tr>
          <th data-width="10" data-checkbox="true"></th>
          <th data-formatter="index1">#</th>
          <th data-field="Category" data-sortable="true">行业</th>
          <th data-field="Number" data-sortable="true">编号</th>
          <th data-field="Year">年份</th>
          <th data-field="action" data-formatter="actionFormatter">完整编号</th>
          <th data-field="Title">名称</th>
        </tr>
      </thead>
    </table>
  </div>
    
  </div>

<div id="footer">
  <div class="col-lg-12">
    <br>
    <hr/>
  </div>

  <div class="col-lg-6">
    <h4>Copyright © 2016~2018 EngineerCMS</h4>
    <p>
      网站由 <i class="user icon"></i>
      <a target="_blank" href="https://github.com/3xxx">@3xxx</a>
      建设，并由
      <a target="_blank" href="http://golang.org">golang</a>
      和
      <a target="_blank" href="http://beego.me">beego</a>
      提供动力。
    </p>

    <p>
      请给 <i class="glyphicon glyphicon-envelope"></i>
      <a class="email" href="mailto:504284@qq.com">我</a>
      发送反馈信息或提交
      <i class="tasks icon"></i>
      <a target="_blank" href="https://github.com/3xxx/engineercms/issues">应用问题</a>
      。
    </p>
  </div>
  <div class="col-lg-6">
    <h4 >更多项目</h4>
    <div >
      <p>
        <a href="https://github.com/3xxx/hydrows">水利供水管线设计工具</a>
      </p>
      <p>
        <a href="https://github.com/3xxx/meritms">成果与价值管理系统</a>
      </p>
      <p>
        <a href="https://github.com/3xxx/engineercms">工程师知识管理系统</a>
      </p>
    </div>
  </div>
</div>


<!-- 使用说明 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modal">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">使用规则</h3>
            <label></label>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">一、查询：</label>
                <label class="col-sm-8">1、显示全部规范，输入allstandard；2、显示全部图集，输入allatlas；3、显示全部计算书，输入allcompute</label>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">二、电子规范命名：</label>
                <label class="col-sm-8">GB 50007-2011建筑地基基础设计规范.pdf</label>
                <label class="col-sm-8">SL 677-2014水工混凝土施工规范.pdf</label>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">三、电子图集命名：</label>
                <label class="col-sm-8">07MS101市政给水管道工程-消火栓、附属构筑物、架空钢管、防水套管.pdf</label>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">四、计算书命名：</label>
                <label class="col-sm-8">波浪爬高.xlsx</label>
              </div>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
          </div>
        </div>
      </div>
    </div>
  </div>

<!-- 修改规范 -->
<div class="container form-horizontal">
  <div class="modal fade" id="editorstandardmodal">
    <div class="modal-dialog">
      <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">修改</h3>
          </div>
          <div class="modal-body">
            <div class="form-group must">
              <label class="col-sm-3 control-label">编号</label>
              <div class="col-sm-7">
                <input type="text" class="form-control" id="number1"></div>
            </div>
            <div class="form-group must">
              <label class="col-sm-3 control-label">名称</label>
              <div class="col-sm-7">
                <input type="text" class="form-control" id="title1"></div>
            </div>
            <div class="form-group must">
              <label class="col-sm-3 control-label">链接</label>
              <div class="col-sm-7">
                <input type="text" class="form-control" id="route1"></div>
            </div>
            <div class="form-group must">
              <label class="col-sm-3 control-label">上传者</label>
              <div class="col-sm-7">
                <input type="tel" class="form-control" id="uname1"></div>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
            <button type="button" class="btn btn-primary" onclick="updatestandard()">修改</button>
          </div>
        </div>
    </div>
  </div>
</div>

<!-- 批量导入规范数据 -->
<div class="container form-horizontal">
  <div class="modal fade" id="importstandardmodal">
    <div class="modal-dialog">
      <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">导入规范数据表</h3>
          </div>
          <div class="modal-body">
              <!-- <div class="form-group"> -->
                <form method="post" action="/standard/importexcel" enctype="multipart/form-data">
                  <label>
                    <input type="file" name="excel" id="excel" class="btn btn-default"/>
                  </label>
                  <button type="submit" class="btn btn-default" onclick="import()">导入规范数据</button>
                </form>
              <!-- </div> -->
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
          </div>
        </div>
    </div>
  </div>
</div>

<!-- 修改有效版本库 -->
<div class="container form-horizontal">
  <div class="modal fade" id="editorvalidmodal">
    <div class="modal-dialog">
      <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">修改</h3>
          </div>
          <div class="modal-body">
            <div class="form-group must">
              <label class="col-sm-3 control-label">行业</label>
              <div class="col-sm-7">
                <input type="text" class="form-control" id="category2"></div>
            </div>
            <div class="form-group must">
              <label class="col-sm-3 control-label">编号</label>
              <div class="col-sm-7">
                <input type="text" class="form-control" id="number2"></div>
            </div>
            <div class="form-group must">
              <label class="col-sm-3 control-label">年份</label>
              <div class="col-sm-7">
                <input type="text" class="form-control" id="year2"></div>
            </div>
            <div class="form-group must">
              <label class="col-sm-3 control-label">名称</label>
              <div class="col-sm-7">
                <input type="text" class="form-control" id="title2"></div>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
            <button type="button" class="btn btn-primary" onclick="update()">修改</button>
          </div>
        </div>
    </div>
  </div>
</div>

<!-- 批量导入有效版本库 -->
<div class="container form-horizontal">
  <div class="modal fade" id="importvalidmodal">
    <div class="modal-dialog">
      <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">导入有效版本数据</h3>
          </div>
          <div class="modal-body">
            <form method="post" action="/importlibrary" enctype="multipart/form-data">
              <label>
                <input type="file" name="excel2" id="excel2" class="btn btn-default"/>
              </label>
              <button type="submit" class="btn btn-default" onclick="import()">导入有效库</button>
            </form>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
          </div>
        </div>
    </div>
  </div>
</div>

<script>
  function actionFormatter(value, row, index) {
    return row.Category+' '+row.Number+'-'+row.Year
  }
  function index1(value,row,index){
  // alert( "Data Loaded: " + index );
    return index+1
  }

  function localDateFormatter(value) {
    return moment(value, 'YYYY-MM-DD').format('YYYY-MM-DD');
  }
  function setLink(value,row,index){
    return '<a href="'+row.Route+'" title="下载" target="_blank"><i class="fa fa-paperclip"></i></a>';
  }
  function setLable(value,row,index){
    return row.LiNumber+row.LibraryTitle
  }

// <button class="btn btn-primary" id="export">导出excel</button>
$(document).ready(function(){
    // 显示规范电子文件命名规则
    $("#about").click(function() {
      $('#modal').modal({
        show:true,
        backdrop:'static'
      });
    })

    //显示和管理有效版本库
    $("#valid").click(function() {
      $("#details2").show();
      $('#table1').bootstrapTable('refresh', {url:'/standard/valid'});
    })

  $("#search").click(function(){//这里应该用button的id来区分按钮的哪一个,因为本页有好几个button
    var title = $('#name').val();
    if (title.length>=2){
          $.ajax({
          type:"post",//这里是否一定要用post？？？
          url:"/standard/search",
          data: {name: $("#name").val()},
          success:function(data,status){//数据提交成功时返回数据
            $("#details").show();
            //追加数据
            $('#table').bootstrapTable('append', data);
            $('#table').bootstrapTable('scrollTo', 'bottom');
            }       
          });
        }else{
          alert("请输入2个以上字符");
        }
    });
  });

  //规范表格增删改
  $(document).ready(function() {
    //添加规范
    $("#addButton").click(function() {
        $('#modalTable').modal({
        show:true,
        backdrop:'static'
        });
    })
    //导入规范数据
    $("#importButton").click(function() {
        $('#importstandardmodal').modal({
        show:true,
        backdrop:'static'
        });
    })

    $("#editorButton").click(function() {
      var selectRow=$('#table').bootstrapTable('getSelections');
      if (selectRow.length<1){
        alert("请先勾选！");
        return;
      }
      if (selectRow.length>1){
      alert("请不要勾选一个以上！");
      return;
      }
      $("input#cid").remove();
      var th1="<input id='cid' type='hidden' name='cid' value='" +selectRow[0].Id+"'/>"
      $(".modal-body").append(th1);//这里是否要换名字$("p").remove();
      $("#number1").val(selectRow[0].Number);
      $("#title1").val(selectRow[0].Title);
      $("#route1").val(selectRow[0].Route);
      $("#uname1").val(selectRow[0].Uname);
        $('#editorstandardmodal').modal({
        show:true,
        backdrop:'static'
        });
    })

    $("#deleteButton").click(function() {
      var selectRow=$('#table').bootstrapTable('getSelections');
      if (selectRow.length<=0) {
        alert("请先勾选！");
        return false;
      }

      if(confirm("确定删除吗？一旦删除将无法恢复！")){
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

        var ids1=$.map(selectRow,function(row){
        return row.id;
        })

        $.ajax({
          type:"post",
          url:"/standard/deletestandard",
          data: {ids:ids},
          success:function(data,status){
            alert("删除“"+data+"”成功！(status:"+status+".)");
            //删除已选数据
            $('#table').bootstrapTable('remove',{
              field:'Title',
              values:title
            });
          }
        });
      }
    })
  })  

  //有效版本库表格增删改
  $(document).ready(function() {
    $("#addButton1").click(function() {
        $('#modalTable').modal({
        show:true,
        backdrop:'static'
        });
    })

    $("#importButton1").click(function() {
        $('#importvalidmodal').modal({
        show:true,
        backdrop:'static'
        });
    })

    $("#editorButton1").click(function() {
      var selectRow=$('#table1').bootstrapTable('getSelections');
      if (selectRow.length<1){
        alert("请先勾选！");
        return;
      }
      if (selectRow.length>1){
      alert("请不要勾选一个以上！");
      return;
      }
      $("input#cid").remove();
      var th1="<input id='cid' type='hidden' name='cid' value='" +selectRow[0].Id+"'/>"
      $(".modal-body").append(th1);//这里是否要换名字$("p").remove();
      $("#category2").val(selectRow[0].Category);
      $("#number2").val(selectRow[0].Number);
      $("#year2").val(selectRow[0].Year);
      $("#title2").val(selectRow[0].Title);
        $('#editorvalidmodal').modal({
        show:true,
        backdrop:'static'
        });
    })

    $("#deleteButton1").click(function() {
      var selectRow=$('#table1').bootstrapTable('getSelections');
      if (selectRow.length<=0) {
        alert("请先勾选！");
        return false;
      }
      if(confirm("确定删除吗？一旦删除将无法恢复！")){
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
        var ids1=$.map(selectRow,function(row){
        return row.id;
        })
        $.ajax({
          type:"post",
          url:"/standard/deletevalid",
          data: {ids:ids},
          success:function(data,status){
            alert("删除“"+data+"”成功！(status:"+status+".)");
            //删除已选数据
            $('#table1').bootstrapTable('remove',{
              field:'Title',
              values:title
            });
          }
        });
      } 
    })
  })

  function updatestandard(){
    // var radio =$("input[type='radio']:checked").val();
    var number1 = $('#number1').val();
    var title1 = $('#title1').val();
    var route1 = $('#route1').val();
    var uname1 = $('#uname1').val();
    var cid = $('#cid').val();
    // $('#myModal').on('hide.bs.modal', function () {  
    if (number1)
      {  
        $.ajax({
          type:"post",
          url:"/standard/updatestandard",
          data: {cid:cid,number:number1,title:title1,route:route1,uname:uname1},
          success:function(data,status){
            alert("添加“"+data+"”成功！(status:"+status+".)");
          }
        });  
      } 
      // $(function(){$('#myModal').modal('hide')});
        $('#editorstandardmodal').modal('hide');
        $('#table').bootstrapTable('refresh', {url:'/getstandard'});
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
      uploader.on( 'fileQueued', function( file ) {
          $list.append( '<div id="' + file.id + '" class="item">' +
              '<h4 class="info">' + file.name + '</h4>' +
              '<p class="state">等待上传...</p>' +
          '</div>' );
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
    uploader.on( 'uploadProgress', function( file, percentage ) {
        var $li = $( '#'+file.id ),
            $percent = $li.find('.progress .progress-bar');
        // 避免重复创建
        if ( !$percent.length ) {
            $percent = $('<div class="progress progress-striped active">' +
              '<div class="progress-bar" role="progressbar" style="width: 0%">' +
              '</div>' +
            '</div>').appendTo( $li ).find('.progress-bar');
        }

        $li.find('p.state').text('上传中');

        $percent.css( 'width', percentage * 100 + '%' );
    });

    uploader.on( 'uploadSuccess', function( file ) {
        $( '#'+file.id ).find('p.state').text('已上传');
    });

    uploader.on( 'uploadError', function( file ) {
        $( '#'+file.id ).find('p.state').text('上传出错');
    });

    uploader.on( 'uploadComplete', function( file ) {
        $( '#'+file.id ).find('.progress').fadeOut();
    });

    uploader.on( 'all', function( type ) {
        if ( type === 'startUpload' ) {
            state = 'uploading';
        } else if ( type === 'stopUpload' ) {
            state = 'paused';
        } else if ( type === 'uploadFinished' ) {
            state = 'done';
        }

        if ( state === 'uploading' ) {
            $btn.text('暂停上传');
        } else {
            $btn.text('开始上传');
        }
    });

    $btn.on( 'click', function() {
        if ( state === 'uploading' ) {
            uploader.stop();
        } else {
            uploader.upload();
        }
    });
  });

  function getKey(){
    var title = $('#name').val();
    if(event.keyCode==13){ 
      if (title.length>=2){
        $.ajax({
        type:"post",//这里是否一定要用post？？？
        url:"/standard/search",
        data: {name: $("#name").val()},
        success:function(data,status){//数据提交成功时返回数据
          $("#details").show();
          //追加数据
          $('#table').bootstrapTable('append', data);
          $('#table').bootstrapTable('scrollTo', 'bottom');
          }       
        });
      }else{
        alert("请输入2个以上字符");
      }
    }   
  }

</script>

</body>
</html> 
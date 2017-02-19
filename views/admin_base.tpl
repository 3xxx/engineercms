<!-- iframe里系统基本设置-->
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>EngineerCMS</title>
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <!-- <script src="/static/js/bootstrap-treeview.js"></script> -->
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>

  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css"/> -->
  
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <!-- <script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script> -->
  <!-- <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script> -->
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>

  <link rel="stylesheet" type="text/css" href="/static/fex-team-webuploader/css/webuploader.css">
  <script type="text/javascript" src="/static/fex-team-webuploader/dist/webuploader.min.js"></script>

  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <script src="/static/js/tableExport.js"></script>
</head>
<body>

<script type="text/javascript">
  function index1(value,row,index){
    return index+1
  }

  function localDateFormatter(value) {
    return moment(value, 'YYYY-MM-DD').format('YYYY-MM-DD');
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
         rowid=row.Id;//全局变量
         $("#details").show();
         $('#table1').bootstrapTable('refresh', {url:'/admin/base/carousel'});
     });
     // $("#get").click(function(){
     //     alert("商品名称：" + getContent().TuanGouName);
     // })
  });
</script>

<div class="col-lg-12">
<h3>系统基本设置</h3>
<div id="toolbar1" class="btn-group">
        <button type="button" data-name="addButton" id="addButton" class="btn btn-default"> <i class="fa fa-plus">添加</i>
        </button> 
        <button type="button" data-name="editorButton" id="editorButton" class="btn btn-default"> <i class="fa fa-edit">编辑</i>
        </button>
        <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
        </button>
</div>
<!--         data-url="/admin/base" -->
<table id="table0"

        data-search="true"
        data-show-refresh="true"
        data-show-toggle="true"
        data-show-columns="true"
        data-toolbar="#toolbar1"
        data-query-params="queryParams"
        data-sort-name="ProjectName"
        data-sort-order="desc"
        data-page-size="5"
        data-page-list="[5, 25, 50, All]"
        data-unique-id="id"
        data-pagination="true"
        data-side-pagination="client"
        data-single-select="true"
        data-click-to-select="true"
        >
    <thead>        
      <tr>
        <!-- radiobox data-checkbox="true"-->
        <th data-width="10" data-radio="true"></th>
        <th data-formatter="index1">#</th>
        <th data-field="Title">名称</th>
        <th data-field="Content">内容</th>
        <th data-field="Tips">备注</th>
      </tr>
    </thead>
</table>
<!-- <div class="gridview2"></div> -->

<script type="text/javascript">
        /*数据json*/
        var json =  [{"Id":"1","Title":"首页轮播图片","Content":"添加宽幅图片","Tips":"最新的10张"},
                      {"Id":"2","Title":"网站昵称","Content":"平行世界","Tips":"网站昵称"},
                      {"Id":"3","Title":"管理员昵称","Content":"无影","Tips":"昵称"},
                      {"Id":"4","Title":"管理员头像","Content":"怪物","Tips":"图像"},
                      {"Id":"5","Title":"导航条定制","Content":"wiki","Tips":"选择导航条"},
                      {"Id":"6","Title":"网站简介","Content":"来自虫洞","Tips":"介绍"},
                      {"Id":"7","Title":"远程访问","Content":"允许","Tips":"vpn或nat123"}];
        /*初始化table数据*/
        $(function(){
            $("#table0").bootstrapTable({
                data:json,
                // onClickRow: function (row, $element) {
                  // alert( "选择了行Id为: " + row.Id );
                  // rowid=row.Id//全局变量
                  // $('#table1').bootstrapTable('refresh', {url:'/admincategory?pid='+row.Id});
                // }
            });
        });

  $(document).ready(function() {
    $("#addButton").click(function() {
        $('#modalTable').modal({
        show:true,
        backdrop:'static'
        });
    })

    $("#editorButton").click(function() {
      var selectRow=$('#table0').bootstrapTable('getSelections');
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
      $("#Iptitle1").val(selectRow[0].Title);
      $("#StartIp1").val(selectRow[0].StartIp);
      $("#EndIp1").val(selectRow[0].EndIp);
      $("#Iprole1").val(selectRow[0].Iprole);
      // alert(JSON.stringify(selectRow));
      // alert(selectRow[0].Id);
      // var title = $('#'+id).attr("value");
      // var title = $('#'+id).attr("href");
      // var categoryid = $('#categoryid').val();
        $('#modalTable1').modal({
        show:true,
        backdrop:'static'
        });
    })

    $("#deleteButton").click(function() {
      var selectRow=$('#table0').bootstrapTable('getSelections');
      // if (selectRow.length<1){
      //   alert("请先勾选类别！");
      //   return;
      // }
      if (selectRow.length<=0) {
        alert("请先勾选！");
        return false;
      }
      var ids=$.map(selectRow,function(row){
        return row.id;
      })
      //删除已选数据
      $('$table0').bootstrapTable('remove',{
        field:'id',
        values:ids
      });
    })

  })

</script>

<!-- 添加 -->
<div class="container">
  <form class="form-horizontal">
    <div class="modal fade" id="modalTable">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">添加</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">名称</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="title"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">内容</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="content"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">备注</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="about"></div>
              </div>
          </div>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
          <button type="button" class="btn btn-primary" onclick="save()">保存</button>
        </div>
      </div>
    </div>
  </div>
</form>
</div>
<script type="text/javascript">
  function save(){
      // var radio =$("input[type='radio']:checked").val();
      var title = $('#title').val();
      var content = $('#content').val();
      var about = $('#about').val();
      // $('#myModal').on('hide.bs.modal', function () {  
      if (title)
        {  
            $.ajax({
                type:"post",
                url:"/admin/base/addbase",
                data: {title:title,content:content,about:about},
                success:function(data,status){
                  alert("添加“"+data+"”成功！(status:"+status+".)");
                 }
            });  
        }else{
          alert("名称不能为空");
        }  
        // $(function(){$('#myModal').modal('hide')}); 
          $('#modalTable').modal('hide');
          $('#table0').bootstrapTable('refresh', {url:'/admin/base'});
          // "/category/modifyfrm?cid="+cid
          // window.location.reload();//刷新页面
  }

  function update(){
    // var radio =$("input[type='radio']:checked").val();
    var title = $('#title1').val();
    var content = $('#content1').val();
      var about = $('#about1').val();
    var cid = $('#cid').val();
    // $('#myModal').on('hide.bs.modal', function () {  
    if (Iptitle1)
      {  
          $.ajax({
              type:"post",
              url:"/admin/base/updatebase",
              data: {cid:cid,title:title,content:content,about:about},
              success:function(data,status){
                alert("添加“"+data+"”成功！(status:"+status+".)");
               }
          });  
      }else{
        alert("名称不能为空");
      } 
      // $(function(){$('#myModal').modal('hide')});
        $('#modalTable1').modal('hide');
        $('#table0').bootstrapTable('refresh', {url:'/admin/base'});
        // "/category/modifyfrm?cid="+cid
        // window.location.reload();//刷新页面
  }
</script>
<!-- 修改 -->
<div class="container">
  <form class="form-horizontal">
    <div class="modal fade" id="modalTable1">
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
              <label class="col-sm-3 control-label">名称</label>
              <div class="col-sm-7">
                <input type="text" class="form-control" id="title1"></div>
            </div>
            <div class="form-group must">
              <label class="col-sm-3 control-label">内容</label>
              <div class="col-sm-7">
                <input type="text" class="form-control" id="content1"></div>
            </div>
            <div class="form-group must">
              <label class="col-sm-3 control-label">备注</label>
              <div class="col-sm-7">
                <input type="tel" class="form-control" id="about1"></div>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
            <button type="button" class="btn btn-primary" onclick="update()">修改</button>
          </div>
        </div>
    </div>
  </div>
</form>
</div>

<!-- ********点击某一行，显示所有ip对应用户名 -->
<!-- onClickRow  click-row.bs.table  row, $element 当用户点击某一行的时候触发，参数包括：
row：点击行的数据，
$element：tr 元素，
field：点击列的 field 名称 -->

<toolbar id="btn_toolbar1" class="toolbar">
<div class="btn-group">
        <button type="button" data-name="addButton1" id="addButton1" class="btn btn-default" data-target="modal"><i class="fa fa-plus" aria-hidden="true"> </i>添加</button>
        <button type="button" data-name="editorButton1" id="editorButton1" class="btn btn-default" data-target="modal"><i class="fa fa-edit" aria-hidden="true"> </i>编辑</button>
        <button type="button" data-name="deleteButton1" id="deleteButton1" class="btn btn-default" data-target="default"><i class="fa fa-trash" aria-hidden="true"> </i>删除</button>
    </div>
</toolbar>
<!-- data-query-params="queryParams" data-content-type="application/json"-->
<div id="details" style="display:none">
<h3>子表格</h3>
<!-- data-url="/admin/category/2" 没有了这个，当然table1表格无法支持刷新了！！！data-show-refresh="true"-->
<table id="table1"
        data-toggle="table"
        data-search="true"
        data-show-toggle="true"
        data-show-columns="true"
        data-toolbar="#btn_toolbar1"
        data-sort-name="ProjectName"
        data-sort-order="desc"
        data-page-size="5"
        data-page-list="[5, 25, 50, All]"
        data-unique-id="id"
        data-pagination="true"
        data-side-pagination="client"
        data-click-to-select="true">
    <thead>        
      <tr>
        <th data-width="10" data-checkbox="true"></th>
        <th data-formatter="index1">#</th>
        <th data-field="Title">名称</th>
        <th data-field="Url">路径</th>
        <th data-field="Created" data-formatter="localDateFormatter">建立</th>
        <th data-field="Updated" data-formatter="localDateFormatter">修改</th>
      </tr>
    </thead>
</table>
</div>

<!-- 添加工程目录分级 -->
<script type="text/javascript">
$(document).ready(function() {
    // 批量上传宽幅图片
    $("#addButton1").click(function() {
      if ({{.role}}!=1){
        alert("权限不够！"+{{.role}});
        return;
      }
        $('#uploaderTable').modal({
        show:true,
        backdrop:'static'
        });
    })

    var uploader;
    $('#uploaderTable').on('shown.bs.modal',function(){
      // var $ = jQuery,
      $list = $('#thelist'),
      $btn = $('#ctlBtn'),
      state = 'pending',
      // uploader;
      uploader = WebUploader.create({
        // 不压缩image
        resize: false,
        // swf文件路径
        swf: '/static/fex-team-webuploader/dist/Uploader.swf',
        // 文件接收服务端。
        server: '/admin/base/addcarousel',
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

      //传递参数——成果id
      uploader.on( 'startUpload', function() {//uploadBeforeSend——这个居然不行？
      // if (prodlabel){
        var pid = $('#pid').val();
        var prodlabel = $('#prodlabel').val();
        var prodprincipal = $('#prodprincipal').val();
        uploader.option('formData', {
          "pid":pid,
          "prodlabel":prodlabel,
          "prodprincipal":prodprincipal
        }); 
      });

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
        $('#table1').bootstrapTable('refresh', {url:'/admin/base/carousel'});
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
    })

    $('#uploaderTable').on('hide.bs.modal',function(){
      $list.text("");
      uploader.destroy();//销毁uploader
    })
})
</script>
  <!-- 批量上传宽幅图片 -->
  <div class="form-horizontal">
    <div class="modal fade" id="uploaderTable">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">批量添加图片</h3>
            <label>**请选择标准宽幅4:1图片上传**</label>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <!--SWF在初始化的时候指定，在后面将展示-->
              <div id="uploader" style="position:relative;text-align: center;">
              <!--用来存放文件信息-->
                <div id="thelist"></div>
                <div class="btns">
                  <div id="picker">选择文件</div>
                  <button id="ctlBtn" class="btn btn-default">开始上传</button>
                </div>
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

</div>

</body>
</html>
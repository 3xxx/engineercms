<!-- 采用bootstrap treeview编辑目录、设置同步ip、设置公开私有 -->
<!DOCTYPE html>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-treeview.css"/>
  <script src="/static/js/bootstrap-treeview.js"></script>
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css"/>
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/font-awesome.min.css"/> -->
  <script src="/static/js/tableExport.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>

</head>

<body>

<div class="col-lg-12">
  <h3>项目列表</h3>
<!-- <div id="toolbar1" class="btn-group">
        <button type="button" data-name="editorcateButton" id="editorcateButton" class="btn btn-default"> <i class="fa fa-edit">编辑目录</i>
        </button>
        <button type="button" data-name="editorpubButton" id="editorpubButton" class="btn btn-default"> <i class="fa fa-edit">公开、私有</i>
        </button>
         <button type="button" data-name="editoripButton" id="editoripButton" class="btn btn-default"> 
        <i class="fa fa-edit">同步IP</i>
        </button>
</div> -->

<table id="table0"
        data-toggle="table"  
        data-url="/project/getprojects"
        data-search="true"
        data-show-refresh="true"
        data-show-toggle="true"
        data-show-columns="true"
        data-query-params="queryParams"
        data-sort-name="ProjectName"
        data-sort-order="desc"
        data-page-size="5"
        data-page-list="[5,20, 50, 100, All]"
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
        <th data-field="Code">编号</th>
        <th data-field="Title">名称</th>
        <th data-field="Label">标签</th>
        <th data-field="Principal">负责人</th>
        <th data-field="Product">成果数量</th>
        <th data-field="Created" data-formatter="localDateFormatter">建立时间</th>
      </tr>
    </thead>
</table>

<script type="text/javascript">
  function index1(value,row,index){
  // alert( "Data Loaded: " + index );
            return index+1
  }
  function localDateFormatter(value) {
     return moment(value, 'YYYY-MM-DD').format('YYYY-MM-DD');
  }
  // 改变点击行颜色
  $(function(){
     $("#table0").on("click-row.bs.table",function(e,row,ele){
         $(".info").removeClass("info");
         $(ele).addClass("info");
         rowid=row.Id;//全局变量
         rowtitle=row.Title
         $("#rowtitle").html("项目同步IP表-"+rowtitle);
         $("#details").show();
         $('#iptable').bootstrapTable('refresh', {url:'/admin/project/synchip/'+row.Id});
     });
  });
</script>

<!-- *******同步IP开始 -->
<script type="text/javascript">
$(document).ready(function() {
  //添加同步ip
  $("#addipButton").click(function() {
        // alert("添加pip"+rowid);
        // if (rowid=""){
        //   alert("请先点击类别！");
        //   return;
        // }
        $("input#pid").remove();
        var th1="<input id='pid' type='hidden' name='pid' value='" +rowid+"'/>"
        $(".modal-body").append(th1);
        $('#addipmodal').modal({
          show:true,
          backdrop:'static'
        });
  })
  //编辑同步ip
  $("#editoripButton").click(function() {
    var selectipRow=$('#iptable').bootstrapTable('getSelections');
    if (selectipRow.length<1){
      alert("请先勾选！");
      return;
    }
    if (selectipRow.length>1){
      alert("请不要勾选一个以上！");
      return;
    }
    $("input#cid").remove();
    var th2="<input id='cid' type='hidden' name='cid' value='" +selectipRow[0].Id+"'/>"
      $(".modal-body").append(th2);

      $("#UserName1").val(selectipRow[0].UserName);
      $("#Ip1").val(selectipRow[0].SynchIp);
      $("#Port1").val(selectipRow[0].Port);

    $('#editoripmodal').modal({
      show:true,
      backdrop:'static'
    });
  })
  //删除同步ip
  $("#deleteipButton").click(function() {
     var selectipRow=$('#iptable').bootstrapTable('getSelections');
     if (selectipRow.length<=0) {
       alert("请先勾选！");
       return false;
     }
     var titles=$.map(selectipRow,function(row){
       return row.UserName;//注意这里要更换！！
     })
     var ids="";
     for(var i=0;i<selectipRow.length;i++){
       if(i==0){
         ids=selectipRow[i].Id;
       }else{
         ids=ids+","+selectipRow[i].Id;
       }  
     }

      if(confirm("确定删除吗？一旦删除将无法恢复！")){
        $.ajax({
          type:"post",
          url:"/admin/project/deletesynchip",
          data: {ids:ids},
          success:function(data,status){
            alert("删除“"+data+"”成功！(status:"+status+".)");
            //删除已选数据
            $('#iptable').bootstrapTable('remove',{
              field:'UserName',
              values:titles
            });
          }
        });
      } 
  })

  // ******试验提交选择的表格************
  $("#submitButton1").click(function() {
    var selectRow3=$('#table1').bootstrapTable('getSelections');
    // var obj = selectRow3.parseJSON();
    // var obj = jQuery.parseJSON(selectRow3);
    console.log(selectRow3[0].Code);
    if (selectRow3.length<1){
      alert("请先勾选目录！");
      return;
    }
    var obj = JSON.stringify(selectRow3);
    alert(selectRow3);
    alert(obj);
    console.log(obj);
    // var ids=$.map(selectRow3,function(row){
    //     return row.Id;
    //   })
    // alert(ids);
    var ids="";
    for(var i=0;i<selectRow3.length;i++){
      if(i==0){
        ids=selectRow3[i].Id;
      }else{
        ids=ids+","+selectRow3[i].Id;
      }
        
    }
    $.ajax({
      type:"post",
      url:"/project/category/addcategory",
      data: {rows:selectRow3},
      success:function(data,status){
        alert("添加“"+data+"”成功！(status:"+status+".)");
      }
    });  
  })
})

</script>
<!-- 工程项目同步ip表 -->
<toolbar id="ip_toolbar" class="toolbar">
<div class="btn-group">
        <button type="button" data-name="addipButton" id="addipButton" class="btn btn-default" data-target="modal"><i class="fa fa-plus" aria-hidden="true"> </i>添加</button>
        <button type="button" data-name="editoripButton" id="editoripButton" class="btn btn-default" data-target="modal"><i class="fa fa-edit" aria-hidden="true"> </i>编辑</button>
        <button type="button" data-name="deleteipButton" id="deleteipButton" class="btn btn-default" data-target="default"><i class="fa fa-trash" aria-hidden="true"> </i>删除</button>
        <!-- <button type="button" data-name="submitButton1" id="submitButton1" class="btn btn-default" data-target="default"><i class="fa fa-cog" aria-hidden="true"> </i>提交</button> -->
    </div>
</toolbar>
<!-- data-query-params="queryParams" data-content-type="application/json"-->
<div id="details" style="display:none">
<h3 id="rowtitle"></h3>
<!-- data-url="/admin/category/2" 没有了这个，当然table1表格无法支持刷新了！！！data-show-refresh="true"-->
<table id="iptable"
        data-toggle="table"
        data-search="true"
        data-show-toggle="true"
        data-show-columns="true"
        data-toolbar="#ip_toolbar"
        data-sort-name="Grade"
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
        <th data-field="UserName">用户名</th>
        <th data-field="SynchIp">IP</th>
        <th data-field="Port">端口</th>
      </tr>
    </thead>
</table>
</div>

<!-- 添加项目同步ip -->
<div class="container">
  <form class="form-horizontal">
    <div class="modal fade" id="addipmodal">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">添加项目同步IP</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">用户名称</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="UserName"></div>
                </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">同步IP</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="Ip"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">端口号</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="Port"></div>
              </div>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
            <!-- <button type="submit" class="btn btn-primary">保存</button> -->
            <button type="button" class="btn btn-primary" onclick="saveip()">保存</button>
          </div>
        </div>
      </div>
    </div>
  </form>
</div>
<script type="text/javascript">
  // 添加同步ip
  function saveip(){
      // var radio =$("input[type='radio']:checked").val();
      var UserName = $('#UserName').val();
      var Ip = $('#Ip').val();
      var parentid = $('#pid').val();
      var Port = $('#Port').val();
      // $('#myModal').on('hide.bs.modal', function () {  
      if (Ip)
        {  
            $.ajax({
                type:"post",
                url:"/admin/project/addsynchip",
                data: {pid:parentid,username:UserName,ip:Ip,port:Port},//父级id
                success:function(data,status){
                  alert("添加“"+data+"”成功！(status:"+status+".)");
                 }
            });  
        } 
        // $(function(){$('#myModal').modal('hide')}); 
          $('#addipmodal').modal('hide');
          $('#iptable').bootstrapTable('refresh', {url:'/admin/project/synchip/'+parentid});
  }
  //更新项目同步ip
  function updateip(){
      // var radio =$("input[type='radio']:checked").val();
      var UserName = $('#UserName1').val();
      var Ip = $('#Ip1').val();
      var Port = $('#Port1').val();
      var cid = $('#cid').val();
      // $('#myModal').on('hide.bs.modal', function () {  
      if (Ip)
        {  
            $.ajax({
                type:"post",
                url:"/admin/project/updatesynchip",
                data: {cid:cid,username:UserName,ip:Ip,port:Port},
                success:function(data,status){
                  alert("修改“"+data+"”成功！(status:"+status+".)");
                 }
            });  
        } 
        // $(function(){$('#myModal').modal('hide')});
          $('#editoripmodal').modal('hide');
          // alert("添加“"+rowid);
          $('#iptable').bootstrapTable('refresh', {url:'/admin/project/synchip/'+rowid});
  } 

</script>
<!-- 修改项目同步IP -->
<div class="form-horizontal" class="container">
    <div class="modal fade" id="editoripmodal">
      <div class="modal-dialog">
        <div  class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">修改项目同步IP</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">用户名称</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="UserName1"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">同步IP</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="Ip1"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">端口号</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="Port1"></div>
              </div>
            </div>
          </div>
          <div class="modal-footer">
          <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
          <button type="submit" class="btn btn-primary" onclick="updateip()">修改</button>
          </div>
        </div>
      </div>
    </div>
</div>

</div>
</body>
</html>
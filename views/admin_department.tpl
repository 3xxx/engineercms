<!-- iframe里组织结构代码表-->
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>Merit</title>
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script src="/static/js/bootstrap-treeview.js"></script>
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>

  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css"/>
  
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  <script src="/static/js/tableExport.js"></script>

</head>
<body>

<script type="text/javascript">
  function index1(value,row,index){
  // alert( "Data Loaded: " + index );
            return index+1
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
         rowtitle=row.Title
         $("#rowtitle").html("部门-"+rowtitle);
         $("#details").show();
         $('#table1').bootstrapTable('refresh', {url:'/admin/department/'+row.Id});
     });
     // $("#get").click(function(){
     //     alert("商品名称：" + getContent().TuanGouName);
     // })
  });
</script>

<div class="col-lg-12">
<h3>部门名称</h3>
<div id="toolbar1" class="btn-group">
        <button type="button" data-name="addButton" id="addButton" class="btn btn-default"> <i class="fa fa-plus">添加</i>
        </button>
        <button type="button" data-name="editorButton" id="editorButton" class="btn btn-default"> <i class="fa fa-edit">编辑</i>
        </button>
        <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
        </button>
</div>

<table id="table0"
        data-url="/admin/department"
        data-search="true"
        data-show-refresh="true"
        data-show-toggle="true"
        data-show-columns="true"
        data-toolbar="#toolbar1"
        data-query-params="queryParams"
        data-sort-name="DepartName"
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
        <!-- <th data-field="Id">Id</th> -->
        <th data-field="Title">部门名称</th>
        <th data-field="Code">部门代码</th>
      </tr>
    </thead>
</table>
<!-- <div class="gridview2"></div> -->

<script type="text/javascript">
        /*数据json*/
        var json =  [{"Id":"1","Title":"水利","Code":"SL"},
                     {"Id":"2","Title":"电力","Code":"DL"},
                     {"Id":"3","Title":"市政","Code":"CJ"},
                     {"Id":"4","Title":"建筑","Code":"JG"},
                     {"Id":"5","Title":"交通","Code":"JT"},
                     {"Id":"6","Title":"境外","Code":"JW"}];
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
        alert("请先勾选部门！");
        return;
      }
      if (selectRow.length>1){
      alert("请不要勾选一个以上部门！");
      return;
      }
      $("input#cid").remove();
      var th1="<input id='cid' type='hidden' name='cid' value='" +selectRow[0].Id+"'/>"
      $(".modal-body").append(th1);//这里是否要换名字$("p").remove();
      $("#projcatename1").val(selectRow[0].Title);
      $("#projcatecode1").val(selectRow[0].Code);
        $('#modalTable1').modal({
        show:true,
        backdrop:'static'
        });
    })

    $("#deleteButton").click(function() {
      var selectRow=$('#table0').bootstrapTable('getSelections');
      if (selectRow.length<=0) {
        alert("请先勾选部门！");
        return false;
      }
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
        url:"/admin/department/deletedepartment",
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
    })
  })

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

<!-- 添加部门代码 -->
<div class="container form-horizontal">
    <div class="modal fade" id="modalTable">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">添加部门</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">部门名称</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="projcatename" name="projcatename"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">部门代码</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="projcatecode" name="projcatecode"></div>
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
</div>
    <script type="text/javascript">
      function save(){
        // var radio =$("input[type='radio']:checked").val();
        var projcatename = $('#projcatename').val();
        var projcatecode = $('#projcatecode').val();
        // $('#myModal').on('hide.bs.modal', function () {  
        if (projcatename){  
          $.ajax({
              type:"post",
              url:"/admin/department/adddepartment",
              data: {title:projcatename,code:projcatecode},
              success:function(data,status){
                alert("添加“"+data+"”成功！(status:"+status+".)");
               }
          });  
        } 
        // $(function(){$('#myModal').modal('hide')}); 
        $('#modalTable').modal('hide');
        $('#table0').bootstrapTable('refresh', {url:'/admin/department'});
        // "/category/modifyfrm?cid="+cid
        // window.location.reload();//刷新页面
      }

      function update(){
        // var radio =$("input[type='radio']:checked").val();
        var projcatename1 = $('#projcatename1').val();
        var projcatecode1 = $('#projcatecode1').val();
        var cid = $('#cid').val();
        // $('#myModal').on('hide.bs.modal', function () {  
        if (projcatename1){  
          $.ajax({
            type:"post",
            url:"/admin/department/updatedepartment",
            data: {cid:cid,title:projcatename1,code:projcatecode1},
            success:function(data,status){
              alert("添加“"+data+"”成功！(status:"+status+".)");
             }
          });  
        } 
        // $(function(){$('#myModal').modal('hide')});
        $('#modalTable1').modal('hide');
        $('#table0').bootstrapTable('refresh', {url:'/admin/department'});
        // "/category/modifyfrm?cid="+cid
        // window.location.reload();//刷新页面
      }
    </script>
<!-- 修改部门代码 -->
<div class="container" class="form-horizontal">
  <!-- <form > -->
    <div class="modal fade" id="modalTable1">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">修改部门</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">部门</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="projcatename1"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">代码</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="projcatecode1"></div>
              </div>
            </div>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
          <button type="button" class="btn btn-primary" onclick="update()">修改</button>
        </div>
      </div>
    </div>
  </div>
<!-- </form> -->
</div>
<!-- onClickRow  click-row.bs.table  row, $element 当用户点击某一行的时候触发，参数包括：
row：点击行的数据，
$element：tr 元素，
field：点击列的 field 名称 -->
<script type="text/javascript">
 /*初始化table数据*/
 /*数据json*/
  var json1 = [{"Id":"1","ProjCateName":"规划","ProjCateCode":"A","ProjCateGrade":"1"},
              {"Id":"2","ProjCateName":"报告","ProjCateCode":"B","ProjCateGrade":"2"},
              {"Id":"3","ProjCateName":"图纸","ProjCateCode":"T","ProjCateGrade":"2"},
              {"Id":"4","ProjCateName":"水工","ProjCateCode":"5","ProjCateGrade":"3"},
              {"Id":"5","ProjCateName":"机电","ProjCateCode":"6","ProjCateGrade":"3"},
              {"Id":"6","ProjCateName":"施工","ProjCateCode":"7","ProjCateGrade":"3"}];

function format_status(status,row,index) {
  if(status == 1){
    return '显示'
  }else if(status == 2){
    return  '隐藏'
  }else if(status == 0){
    return  '禁止'
  }
}

$(document).ready(function() {
  //添加科室
  $("#addButton1").click(function() {
        $("input#pid").remove();
        var th1="<input id='pid' type='hidden' name='pid' value='" +rowid+"'/>"
        $(".modal-body").append(th1);
        $('#modalTable2').modal({
          show:true,
          backdrop:'static'
        });
  })
  //编辑科室
  $("#editorButton1").click(function() {
    var selectRow3=$('#table1').bootstrapTable('getSelections');
    if (selectRow3.length<1){
      alert("请先勾选科室！");
      return;
    }
    if (selectRow3.length>1){
      alert("请不要勾选一个以上科室！");
      return;
    }
    $("input#cid").remove();
    var th2="<input id='cid' type='hidden' name='cid' value='" +selectRow3[0].Id+"'/>"
      $(".modal-body").append(th2);//这里是否要换名字$("p").remove();
      $("#projcatename3").val(selectRow3[0].Title);
      $("#projcatecode3").val(selectRow3[0].Code);
    $('#modalTable3').modal({
      show:true,
      backdrop:'static'
    });
  })
  //删除科室
  $("#deleteButton1").click(function() {
     var selectRow=$('#table1').bootstrapTable('getSelections');
     if (selectRow.length<=0) {
       alert("请先勾选科室！");
       return false;
     }
     var titles=$.map(selectRow,function(row){
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
       url:"/admin/department/deletedepartment",
       data: {ids:ids},
       success:function(data,status){
         alert("删除“"+data+"”成功！(status:"+status+".)");
         //删除已选数据
         $('#table1').bootstrapTable('remove',{
           field:'Title',
           values:titles
         });
       }
     }); 
  })
  //科室绑定价值
  $("#bindButton1").click(function() {
    var selectRow3=$('#table1').bootstrapTable('getSelections');
    if (selectRow3.length<1){
      alert("请先勾选！");
      return;
    }
    if (selectRow3.length>1){
      alert("请不要勾选一个以上！");
      return;
    }
    
    $("input#cid").remove();
    var th2="<input id='cid' type='hidden' name='cid' value='" +selectRow3[0].Id+"'/>"
    $(".modal-body").append(th2);
    $("#secoffice").html("绑定科室价值-"+selectRow3[0].Title);
    $('#modalTable4').modal({
      show:true,
      backdrop:'static'
    });
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

<!-- 科室 -->
<toolbar id="btn_toolbar1" class="toolbar">
<div class="btn-group">
        <button type="button" data-name="addButton1" id="addButton1" class="btn btn-default" data-target="modal"><i class="fa fa-plus" aria-hidden="true"></i>添加</button>
        <button type="button" data-name="editorButton1" id="editorButton1" class="btn btn-default" data-target="modal"><i class="fa fa-edit" aria-hidden="true"></i>编辑</button>
        <button type="button" data-name="deleteButton1" id="deleteButton1" class="btn btn-default" data-target="default"><i class="fa fa-trash" aria-hidden="true"> </i>删除</button>
        <button type="button" data-name="bindButton1" id="bindButton1" class="btn btn-default" data-target="modal"><i class="fa fa-edit" aria-hidden="true"></i>价值</button>
        <!-- <button type="button" data-name="submitButton1" id="submitButton1" class="btn btn-default" data-target="default"><i class="fa fa-cog" aria-hidden="true"> </i>提交</button> -->
    </div>
</toolbar>
<!-- data-query-params="queryParams" data-content-type="application/json"-->
<!-- 显示科室 -->
<div id="details" style="display:none">
  <h3 id="rowtitle"></h3>
  <!-- data-url="/admin/category/2" 没有了这个，当然table1表格无法支持刷新了！！！data-show-refresh="true"-->
  <table id="table1"
        data-toggle="table"
        data-search="true"
        data-show-toggle="true"
        data-show-columns="true"
        data-toolbar="#btn_toolbar1"
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
        <th data-field="Title">名称</th>
        <th data-field="Code">代码</th>
      </tr>
    </thead>
  </table>
</div>
<script type="text/javascript">
    // 改变点击行颜色
  $(function(){
     // $("#table").bootstrapTable('destroy').bootstrapTable({
     //     columns:columns,
     //     data:json
     // });
     $("#table1").on("click-row.bs.table",function(e,row,ele){
         $(".info").removeClass("info");
         $(ele).addClass("info");
         rowid=row.Id;//全局变量
         rowtitle=row.Title
         $("#rowtitle").html("部门-"+rowtitle);
         $("#details").show();
         $('#table2').bootstrapTable('refresh', {url:'/admin/department/'+row.Id});
     });
     // $("#get").click(function(){
     //     alert("商品名称：" + getContent().TuanGouName);
     // })
  });
</script>
<!-- 添加科室 -->
<div class="container">
  <form  class="form-horizontal">
    <div class="modal fade" id="modalTable2">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">添加科室</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">科室名称</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="projcatename2"></div>
                </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">代码</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="projcatecode2"></div>
              </div>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
            <button type="button" class="btn btn-primary" onclick="save2()">保存</button>
          </div>
        </div>
      </div>
    </div>
  </form>
</div>
<script type="text/javascript">
  function save2(){
      // var radio =$("input[type='radio']:checked").val();
      var projcatename2 = $('#projcatename2').val();
      var projcatecode2 = $('#projcatecode2').val();
      var parentid = $('#pid').val();
      // $('#myModal').on('hide.bs.modal', function () {  
      if (projcatename2)
        {  
            $.ajax({
                type:"post",
                url:"/admin/department/adddepartment",
                data: {pid:parentid,title:projcatename2,code:projcatecode2},
                success:function(data,status){
                  alert("添加“"+data+"”成功！(status:"+status+".)");
                 }
            });  
        } 
        // $(function(){$('#myModal').modal('hide')}); 
          $('#modalTable2').modal('hide');
          $('#table1').bootstrapTable('refresh', {url:'/admin/department/'+parentid});
  }
  function update2(){
      // var radio =$("input[type='radio']:checked").val();
      var projcatename3 = $('#projcatename3').val();
      var projcatecode3 = $('#projcatecode3').val();
      var cid = $('#cid').val(); 
      if (projcatename3)
        {  
            $.ajax({
                type:"post",
                url:"/admin/department/updatedepartment",
                data: {cid:cid,title:projcatename3,code:projcatecode3},
                success:function(data,status){
                  alert("添加“"+data+"”成功！(status:"+status+".)");
                 }
            });  
        } 
        // $(function(){$('#myModal').modal('hide')});
          $('#modalTable3').modal('hide');
          // alert("添加“"+rowid);
          $('#table1').bootstrapTable('refresh', {url:'/admin/department/'+rowid});
  } 

</script>
<!-- 修改科室 -->
<div class="container">
  <form class="form-horizontal">
    <div class="modal fade" id="modalTable3">
      <div class="modal-dialog">
        <div  class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">修改科室</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">科室名称</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="projcatename3"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">代码</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="projcatecode3"></div>
              </div>
            </div>
          </div>
          <div class="modal-footer">
          <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
          <button type="submit" class="btn btn-primary" onclick="update2()">修改</button>
          </div>
        </div>
      </div>
    </div>
  </form>
</div>
<!-- 绑定科室价值 -->
<div class="container">
  <form class="form-horizontal">
    <div class="modal fade" id="modalTable4">
      <div class="modal-dialog">
        <div  class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title" id="secoffice"></h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <table id="table0"
                      data-toggle="table"
                      data-url="/admin/merit"
                      data-page-size="5"
                      data-page-list="[5, 25, 50, All]"
                      data-unique-id="id"
                      data-pagination="true"
                      data-side-pagination="client"
                      data-click-to-select="true"
                      >
                  <thead>        
                    <tr>
                      <!-- radiobox data-checkbox="true"-->
                      <th data-width="10" data-checkbox="true"></th>
                      <th data-formatter="index1">#</th>
                      <th data-field="Title">价值名称</th>
                      <th data-field="Mark">价值分值</th>
                      <th data-field="List">价值选项</th>
                      <th data-field="ListMark">选项分值</th>
                      <!-- <th data-field="Iprole" data-title-tooltip="1-管理员;2-下载任意附件;3-下载pdf;4-查看成果">权限等级</th> -->
                    </tr>
                  </thead>
              </table>
            </div>
          </div>
          <div class="modal-footer">
          <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
          <button type="submit" class="btn btn-primary" onclick="savedepmerit()">保存</button>
          </div>
        </div>
      </div>
    </div>
  </form>
</div>
<br/>
<br/>
</div>

</body>
</html>
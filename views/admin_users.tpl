<!-- iframe里用户列表-->
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>EngineerCMS</title>
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
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <script src="/static/js/jquery.form.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/select2.css"/>
  <script type="text/javascript" src="/static/js/select2.js"></script>
</head>
<body>

<div class="col-lg-12">
<h3>用户表</h3>
  <div id="toolbar1" class="btn-group">
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
  <!-- data-url="/admin/user"
   -->
  <table id="table0"
        data-search="true"
        data-show-refresh="true"
        data-show-toggle="true"
        data-show-columns="true"
        data-striped="true"
        data-toolbar="#toolbar1"
        data-query-params="queryParams"
        data-sort-name="Username"
        data-sort-order="desc"
        data-page-size="5"
        data-page-list="[5, 25, 50, All]"
        data-unique-id="id"
        data-pagination="true"
        data-side-pagination="client"
        data-single-select="true"
        data-click-to-select="true"
        data-show-export="true"
        >
    <!-- <thead>        
      <tr>       
        radiobox data-checkbox="true"
        <th data-width="10" data-radio="true"></th>
        <th data-formatter="index1">#</th>
        <th data-field="Username">用户名</th>
        <th data-field="Nickname">昵称</th>
        <th data-field="Password">密码</th>
        <th data-field="Email">邮箱</th>
        <th data-field="Department">部门</th>
        <th data-field="Secoffice">科室</th>
        <th data-field="Ip">IP</th>
        <th data-field="Status">状态</th>
        <th data-field="Lastlogintime" data-formatter="localDateFormatter">最后登陆</th>
        <th data-field="Createtime" data-formatter="localDateFormatter">建立</th>
        <th data-field="Role">权限</th>
      </tr>
    </thead> -->
  </table>

<script type="text/javascript">
        /*数据json*/
        var json =  [{"Id":"1","UserName":"水利","UserNickname":"SL","Lastlogintime":"2016-01-05"},
                     {"Id":"2","UserName":"电力","UserNickname":"DL"},
                     {"Id":"3","UserName":"市政","UserNickname":"CJ"},
                     {"Id":"4","UserName":"建筑","UserNickname":"JG"},
                     {"Id":"5","UserName":"交通","UserNickname":"JT"},
                     {"Id":"6","UserName":"境外","UserNickname":"JW"}];
        /*初始化table数据*/
        //不能用这个，否则editable无法显示数据
        // $(function(){
        //     $("#table0").bootstrapTable({
        //         data:json,
        //     });
        // });
  $(document).ready(function() {
    $("#addButton").click(function() {
        $('#modalTable').modal({
        show:true,
        backdrop:'static'
        });
    })

    //importusers
    $("#importButton").click(function() {
        $('#importusers').modal({
        show:true,
        backdrop:'static'
        });
    })
    //用表格在线编辑比较好
    // $("#editorButton").click(function() {
    //   var selectRow=$('#table0').bootstrapTable('getSelections');
    //   if (selectRow.length<1){
    //     alert("请先勾选类别！");
    //     return;
    //   }
    //   if (selectRow.length>1){
    //   alert("请不要勾选一个以上类别！");
    //   return;
    //   }
    //   $("input#cid").remove();
    //   var th1="<input id='cid' type='hidden' name='cid' value='" +selectRow[0].Id+"'/>"
    //   $(".modal-body").append(th1);//这里是否要换名字$("p").remove();
    //   $("#Username1").val(selectRow[0].Username);
    //   $("#Nickname1").val(selectRow[0].Nickname);
    //   $("#Password1").val(selectRow[0].Password);
    //   $("#Repassword1").val(selectRow[0].Repassword);
    //   $("#Email1").val(selectRow[0].Email);
    //   $("#Department1").val(selectRow[0].Department);
    //   $("#Secoffice1").val(selectRow[0].Secoffice);
    //   $("#Ip1").val(selectRow[0].Ip);
    //   $("#Status1").val(selectRow[0].Status);
    //   $("#Role1").val(selectRow[0].Role);
    //     $('#modalTable1').modal({
    //     show:true,
    //     backdrop:'static'
    //     });
    // })

    // $("#deleteButton").click(function() {
    //   var selectRow=$('#table0').bootstrapTable('getSelections');
    //   if (selectRow.length<=0) {
    //     alert("请先勾选类别！");
    //     return false;
    //   }
    //   var title=$.map(selectRow,function(row){
    //     return row.Title;
    //   })
    //   var ids="";
    //   for(var i=0;i<selectRow.length;i++){
    //     if(i==0){
    //       ids=selectRow[i].Id;
    //     }else{
    //       ids=ids+","+selectRow[i].Id;
    //     }  
    //   }
    //   $.ajax({
    //     type:"post",
    //     url:"/admin/category/deletecategory",
    //     data: {ids:ids},
    //     success:function(data,status){
    //       alert("删除“"+data+"”成功！(status:"+status+".)");
    //       //删除已选数据
    //       $('#table0').bootstrapTable('remove',{
    //         field:'Title',
    //         values:title
    //       });
    //     }
    //   });  
    // })
  })

</script>

<!-- 添加用户 -->
<div class="container">
  <form class="form-horizontal">
    <div class="modal fade" id="modalTable">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">添加用户</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">        
              <div class="form-group must">
                <label class="col-sm-3 control-label">用户名</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="Username"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">昵称</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="Nickname"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">密码</label>
                <div class="col-sm-7">
                  <input type="password" class="form-control" id="password" maxlength="32" placeholder="至多32个字符"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">确认密码</label>
                <div class="col-sm-7">
                  <input type="password" class="form-control equalto" name="password2" maxlength="32" placeholder="至多32个字符" data-rule-equalto="#password" data-msg-equalto="密码不一致"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">邮箱</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="Email"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">部门</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="Department"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">科室</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="Secoffice"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">IP</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="Ip"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">CMS端口Port</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="Port"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">状态</label>
                <div class="col-sm-7">
                  <select id="Status" class="form-control">
                    <option value="1" >显示</option>
                    <option value="2" >隐藏</option>
                    <option value="0" >禁用</option>
                  </select>
                </div>
              </div>
              <div class="form-group">
                <label class="col-sm-3 control-label">权限</label>
                <div class="col-sm-7">
                <input type="number" class="form-control digits" id="Role" maxlength="20" placeholder="至多20个字符"></div>
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
<!-- 导入用户数据 -->
<div class="container form-horizontal">
    <div class="modal fade" id="importusers">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">添加用户</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content"> 
              <div class="form-group">
                <form method="post" id="form1" action="/admin/user/importusers" enctype="multipart/form-data">
                  <div class="form-inline" class="form-group">
                    <label>选择用户数据文件(Excel)：
                      <input type="file" class="form-control" name="usersexcel" id="usersexcel"/> </label>
                    <br/>          
                  </div>
                  <!-- <button type="submit" class="btn btn-default">提交</button> -->
                </form>
              </div>
            </div>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
          <button type="submit" class="btn btn-primary" onclick="return importusers();">导入</button>
          <!-- <button type="submit" class="btn btn-primary" onclick="return import_xls_catalog();">提交</button> -->
        </div>
      </div>
    </div>
  </div>
</div>
<script type="text/javascript">
    function save(){
      // var radio =$("input[type='radio']:checked").val();        
      var Username   = $('#Username').val();
      var Nickname   = $('#Nickname').val();
      var Password   = $('#Password').val();
      var Repassword = $('#Repassword').val();
      var Email      = $('#Email').val();
      var Department = $('#Department').val();
      var Secoffice  = $('#Secoffice').val();
      var Ip         = $('#Ip').val();
      var Port         = $('#Port').val();
      // var Status     = $('#Status option:selected').text();
      var Status     = $('#Status option:selected').val();
      var Role       = $('#Role').val();
      if (Username)
        {  
            $.ajax({
                type:"post",
                url:"/admin/user/adduser",
                data: {username:Username,nickname:Nickname,password:Password,repassword:Repassword,email:Email,department:Department,secoffice:Secoffice,ip:Ip,port:Port,status:Status,role:Role},
                success:function(data,status){
                  alert("添加“"+data+"”成功！(status:"+status+".)");
                 }
            });  
        } else{
          alert("用户名等不能为空！");
        }
        // $(function(){$('#myModal').modal('hide')}); 
          $('#modalTable').modal('hide');
          $('#table0').bootstrapTable('refresh', {url:'/admin/user'});
          // "/category/modifyfrm?cid="+cid
          // window.location.reload();//刷新页面
    }
      //导入用户数据表
    function importusers(){
        var file=$("#usersexcel").val();
        if(file!=""){  
            var form = $("form[id=form1]");
            var options  = {    
                url:'/admin/user/importusers',    
                type:'post', 
                success:function(data)    
                {    
                  alert("导入数据："+data+"！")
                }    
            };
           form.ajaxSubmit(options);
           return false;
        }else{
            alert("请选择文件！");
            return false; 
        }
    }

  $(function () {
    $('#table0').bootstrapTable({
        idField: 'Id',
        url: '/admin/user',
        // striped: "true",
        columns: [
          {
            radio: 'true',
            width: '10'
          },
          {
            // field: 'Number',
            title: '序号',
            formatter:function(value,row,index){
            return index+1
            }
          },{
            field: 'Username',
            title: '用户名',
            sortable:'true',
            editable: {
                type: 'text',
                pk: 1,
                url: '/admin/user/updateuser',
                title: 'Enter ProjectNumber' 
            }
          },{
            field: 'Nickname',
            title: '昵称',
            editable: {
                type: 'text',
                pk: 1,
                url: '/admin/user/updateuser',
                title: 'Enter ProjectName'  
            }
          },{
            field: 'Password',
            title: '密码',
            editable: {
              type: 'text',
                // type: 'select',
                // source: ["规划", "项目建议书", "可行性研究", "初步设计", "招标设计", "施工图"],
                pk: 1,
                url: '/admin/user/updateuser',
                title: 'Enter Password'  
            }
          },{
            field: 'Email',
            title: '邮箱',
            // sortable:'true',
            editable: {
                type: 'text',
                pk: 1,
                url: '/admin/user/updateuser',
                title: 'Enter Email'  
            }
          },{
            field: 'Department',
            title: '部门',
            editable: {
                type: 'text',
                pk: 1,
                url: '/admin/user/updateuser',
                title: 'Enter Department'  
            }
          },{
            field: 'Secoffice',
            title: '科室',
            sortable:'true',
            editable: {
                type: 'text',
                // source: {{.Select2}},//["$1", "$2", "$3"],
                pk: 1,
                url: '/admin/user/updateuser',
                title: 'Enter Category' 
            }
          },{
            field: 'Ip',
            title: 'IP',
            editable: {
                type: 'text',
                pk: 1,
                url: '/admin/user/updateuser',
                title: 'Enter Count'  
            }
          },{
            field: 'Port',
            title: 'CMS端口Port',
            editable: {
                type: 'text',
                pk: 1,
                url: '/admin/user/updateuser',
                title: 'Enter Port'  
            }
          },{
            field: 'Status',
            title: '状态',
            editable: {
                type: 'select2', 
                // source:{{.Userselect}},//'/regist/getuname1',
                source: [
                      {id: '1', text: '显示',value:1},
                      {id: '2', text: '隐藏',value:2},
                      {id: '3', text: '禁止',value:3}
                   ],
        //'[{"id": "1", "text": "One"}, {"id": "2", "text": "Two"}]'
                select2: {
                  allowClear: true,
                  width: '150px',
                  placeholder: '请选择状态',
                  // multiple: true
                },//'/regist/getuname1',//这里用get方法，所以要换一个
                pk: 1,
                url: '/admin/user/updateuser',
                title: 'Enter Status'  
            }
          },{
            field: 'Lastlogintime',
            title: '最后登录',
            formatter:localDateFormatter,
          },{
            field: 'Createtime',
            title: '建立',
            formatter:localDateFormatter,
          },{
            field: 'Role',
            title: '权限',
            editable: {
                type: 'select2', 
                // source:{{.Userselect}},//'/regist/getuname1',
                source: [
                      {id: '1', text: '1级',value:1},
                      {id: '2', text: '2级',value:2},
                      {id: '3', text: '3级',value:3},
                      {id: '4', text: '4级',value:4}
                   ],
        //'[{"id": "1", "text": "One"}, {"id": "2", "text": "Two"}]'
                select2: {
                  allowClear: true,
                  width: '150px',
                  placeholder: '请选择权限',
                  // multiple: true
                },//'/regist/getuname1',//这里用get方法，所以要换一个
                pk: 1,
                url: '/admin/user/updateuser',
                title: 'Enter Status'  
            }
          }
        ]
    });
  });

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
         rowtitle=row.Nickname
         $("#rowtitle").html("用户详情-"+rowtitle);
         $("#details").show();
         $('#table1').bootstrapTable('refresh', {url:'/admin/user/'+row.Id});
     });
  });
    //下面这个用在线表格编辑代替
    // function update(){
    // var radio =$("input[type='radio']:checked").val();
      // var Username   = $('#Username1').val();
      // var Nickname   = $('#Nickname1').val();
      // var Password   = $('#Password1').val();
      // var Repassword = $('#Repassword1').val();
      // var Email      = $('#Email1').val();
      // var Department = $('#Department1').val();
      // var Secoffice  = $('#Secoffice1').val();
      // var Ip         = $('#Ip').val();
      // var Status     = $('#Status1 option:selected').text();
      // var Role       = $('#Role1').val();
      // var cid = $('#cid').val();
    // $('#myModal').on('hide.bs.modal', function () {  
    //   if (projcatename1)
    //     {  
    //       $.ajax({
    //           type:"post",
    //           url:"/admin/user/updateuser",
    //           data: {cid:cid,username:Username,nickname:Nickname,password:Password,repassword:Repassword,email:Email,department:Department,secoffice:Secoffice,ip:Ip,statusStatus,role:Role},
    //           success:function(data,status){
    //             alert("添加“"+data+"”成功！(status:"+status+".)");
    //            }
    //       });  
    //     } 
    //     $('#modalTable1').modal('hide');
    //     $('#table0').bootstrapTable('refresh', {url:'/admin/user'});
    // }
</script>

<!-- onClickRow  click-row.bs.table  row, $element 当用户点击某一行的时候触发，参数包括：
row：点击行的数据，
$element：tr 元素，
field：点击列的 field 名称 -->
<script type="text/javascript">
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
  //添加分级
  $("#addButton1").click(function() {
        $("input#pid").remove();
        var th1="<input id='pid' type='hidden' name='pid' value='" +rowid+"'/>"
        $(".modal-body").append(th1);

        $('#modalTable2').modal({
          show:true,
          backdrop:'static'
        });
  })
  //编辑分级目录
  $("#editorButton1").click(function() {
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
      $(".modal-body").append(th2);//这里是否要换名字$("p").remove();

      $("#projcatename3").val(selectRow3[0].Title);
      $("#projcatecode3").val(selectRow3[0].Code);
      $("#projcategrade3").val(selectRow3[0].Grade);

    $('#modalTable3').modal({
      show:true,
      backdrop:'static'
    });
  })
  //删除分级
  $("#deleteButton1").click(function() {
     var selectRow=$('#table1').bootstrapTable('getSelections');
     if (selectRow.length<=0) {
       alert("请先勾选！");
       return false;
     }
     var titles=$.map(selectRow,function(row){
       // alert(row.Id);
       // return row.Id;
       // alert(row.Title);
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
       url:"/admin/category/deletecategory",
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

<!-- 用户详情表 -->
<toolbar id="btn_toolbar1" class="toolbar">
<div class="btn-group">
        <button type="button" data-name="addButton1" id="addButton1" class="btn btn-default" data-target="modal"><i class="fa fa-plus" aria-hidden="true"> </i>添加</button>
        <button type="button" data-name="editorButton1" id="editorButton1" class="btn btn-default" data-target="modal"><i class="fa fa-edit" aria-hidden="true"> </i>编辑</button>
        <button type="button" data-name="deleteButton1" id="deleteButton1" class="btn btn-default" data-target="default"><i class="fa fa-trash" aria-hidden="true"> </i>删除</button>
        <button type="button" data-name="submitButton1" id="submitButton1" class="btn btn-default" data-target="default"><i class="fa fa-cog" aria-hidden="true"> </i>提交</button>
    </div>
</toolbar>
<!-- data-query-params="queryParams" data-content-type="application/json"-->
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
        <th data-field="Username">用户名</th>
        <th data-field="Nickname">昵称</th>
        <th data-field="Password">密码</th>
        <th data-field="Email">邮箱</th>
        <th data-field="Department">部门</th>
        <th data-field="Secoffice">科室</th>
        <th data-field="Ip">IP</th>
        <th data-field="Status">状态</th>
        <th data-field="Lastlogintime" data-formatter="localDateFormatter">最后登陆</th>
        <th data-field="Createtime" data-formatter="localDateFormatter">建立</th>
        <th data-field="Role">权限</th>
      </tr>
    </thead>
</table>
</div>

<!-- 添加用户详细 -->
<div class="container">
  <form class="form-horizontal">
    <div class="modal fade" id="modalTable2">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">添加详细</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">名称</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="projcatename2"></div>
                </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">代码</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="projcatecode2"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">级别</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="projcategrade2"></div>
              </div>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
            <!-- <button type="submit" class="btn btn-primary">保存</button> -->
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
      var projcategrade2 = $('#projcategrade2').val();
      // $('#myModal').on('hide.bs.modal', function () {  
      if (projcatename2)
        {  
            $.ajax({
                type:"post",
                url:"/admin/category/addcategory",
                data: {pid:parentid,title:projcatename2,code:projcatecode2,grade:projcategrade2},//父级id
                success:function(data,status){
                  alert("添加“"+data+"”成功！(status:"+status+".)");
                 }
            });  
        } 
        // $(function(){$('#myModal').modal('hide')}); 
          $('#modalTable2').modal('hide');
          $('#table1').bootstrapTable('refresh', {url:'/admin/category/'+parentid});
          // "/category/modifyfrm?cid="+cid
          // window.location.reload();//刷新页面
  }
  function update2(){
      // var radio =$("input[type='radio']:checked").val();
      var projcatename3 = $('#projcatename3').val();
      var projcatecode3 = $('#projcatecode3').val();
      var projcategrade3 = $('#projcategrade3').val();
      var cid = $('#cid').val();
      // $('#myModal').on('hide.bs.modal', function () {  
      if (projcatename3)
        {  
            $.ajax({
                type:"post",
                url:"/admin/category/updatecategory",
                data: {cid:cid,title:projcatename3,code:projcatecode3,grade:projcategrade3},
                success:function(data,status){
                  alert("添加“"+data+"”成功！(status:"+status+".)");
                 }
            });  
        } 
        // $(function(){$('#myModal').modal('hide')});
          $('#modalTable3').modal('hide');
          // alert("添加“"+rowid);
          $('#table1').bootstrapTable('refresh', {url:'/admin/category/'+rowid});
  } 

</script>
<!-- 修改工程目录分级 -->
<div class="form-horizontal" class="container">
  <!-- <form class="form-horizontal"> -->
    <div class="modal fade" id="modalTable3">
      <div class="modal-dialog">
        <div  class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">修改目录分级</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">目录名称</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="projcatename3"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">代码</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="projcatecode3"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">级别</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="projcategrade3"></div>
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
  <!-- </form> -->
</div>
<br/>
<br/>
</div>

</body>
</html>
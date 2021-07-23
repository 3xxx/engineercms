<!-- iframe里merit基本信息-->
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>EngineerCMS</title>
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <!-- <script src="/static/js/bootstrap-treeview.js"></script> -->
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
  <!-- <script src="/static/js/jquery.form.js"></script> -->
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/select2.min.css"/> -->
  <!-- <script type="text/javascript" src="/static/js/select2.js"></script> -->
</head>
<body>

<div class="col-lg-12">
<h3>merit基本信息表</h3>
  <!-- <div id="toolbar1" class="btn-group"> -->
        <!-- <button type="button" data-name="addButton" id="addButton" class="btn btn-default"> <i class="fa fa-plus">添加</i>
        </button>
        <button type="button" data-name="importButton" id="importButton" class="btn btn-default"> <i class="fa fa-plus">导入</i>
        </button>
        <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
        </button> -->
  <!-- </div> -->
  <!-- data-url="/v1/wx/user"
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
  </table>

<script type="text/javascript">
        /*数据json*/
        var json =  [{"Id":"1","UserName":"水利","UserNickname":"SL","Lastlogintime":"2016-01-05"},
                     {"Id":"2","UserName":"电力","UserNickname":"DL"},
                     {"Id":"3","UserName":"市政","UserNickname":"CJ"},
                     {"Id":"4","UserName":"建筑","UserNickname":"JG"},
                     {"Id":"5","UserName":"交通","UserNickname":"JT"},
                     {"Id":"6","UserName":"境外","UserNickname":"JW"}];
</script>

<!-- 添加用户 -->
<!-- <div class="container">
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
</div> -->
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
                <form method="post" id="form1" action="/v1/wx/importusers" enctype="multipart/form-data">
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
                url:"/v1/wx/adduser",
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
          $('#table0').bootstrapTable('refresh', {url:'/v1/wx/user/0'});
          // "/category/modifyfrm?cid="+cid
          // window.location.reload();//刷新页面
    }
      //导入用户数据表
    function importusers(){
        var file=$("#usersexcel").val();
        if(file!=""){  
            var form = $("form[id=form1]");
            var options  = {    
                url:'/v1/wx/importusers',    
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
        url: '/admin/merit/meritbasic',
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
                url: '/admin/merit/updatemeritbasic',
                title: 'Enter ProjectNumber' 
            }
          },{
            field: 'Nickname',
            title: '昵称',
            editable: {
                type: 'text',
                pk: 1,
                url: '/admin/merit/updatemeritbasic',
                title: 'Enter ProjectName'  
            }
          },{
            field: 'Password',
            title: '密码',
            editable: {
              type: 'text',
                pk: 1,
                url: '/admin/merit/updatemeritbasic',
                title: 'Enter Password'  
            }
          },{
            field: 'Ip',
            title: 'Merit IP',
            editable: {
                type: 'text',
                pk: 1,
                url: '/admin/merit/updatemeritbasic',
                title: 'Enter Count'  
            }
          },{
            field: 'Port',
            title: 'Merit端口Port',
            editable: {
                type: 'text',
                pk: 1,
                url: '/admin/merit/updatemeritbasic',
                title: 'Enter Port'  
            }
          },{
            field: 'EcmsIp',
            title: 'Ecms的IP',
            editable: {
                type: 'text',
                pk: 1,
                url: '/admin/merit/updatemeritbasic',
                title: 'Enter Port'  
            }
          },{
            field: 'EcmsPort',
            title: 'Ecms的端口',
            editable: {
                type: 'text',
                pk: 1,
                url: '/admin/merit/updatemeritbasic',
                title: 'Enter Port'  
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
     });
  });
</script>

</div>

</body>
</html>
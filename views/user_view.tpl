<!-- 用户登录后自己的资料列表-->
<!DOCTYPE html>
<html>

<head>
  <meta charset="UTF-8">
  <title>EngineerCMS</title>
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <!-- <script src="/static/js/bootstrap-treeview.js"></script> -->
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css" />
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css" />
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css" />
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
  <script src="/static/js/tableExport.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <!-- <script src="/static/js/jquery.form.js"></script> -->
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/select2.min.css"/> -->
  <!-- <script type="text/javascript" src="/static/js/select2.js"></script> -->
</head>
<!-- <div class="navbar navba-default navbar-fixed-top"> -->
<div class="container-fill">{{template "navbar" .}}</div>
<!-- </div> -->

<body>
  <div class="col-lg-12">
    <h3>用户表-{{.Username}}</h3>
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
    <table id="table0" data-search="true" data-show-refresh="true" data-show-toggle="true" data-show-columns="true" data-striped="true" data-toolbar="#toolbar1" data-query-params="queryParams" data-sort-name="Username" data-sort-order="desc" data-page-size="5" data-page-list="[5, 25, 50, All]" data-unique-id="id" data-pagination="true" data-side-pagination="client" data-single-select="true" data-click-to-select="true" data-show-export="true">
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
                      <input type="password" class="form-control" id="Password" maxlength="32" placeholder="至多32个字符" required></div>
                  </div>
                  <div class="form-group must">
                    <label class="col-sm-3 control-label">确认密码</label>
                    <div class="col-sm-7">
                      <input type="password" class="form-control equalto" name="password2" maxlength="32" placeholder="至多32个字符" required data-rule-equalto="#password" data-msg-equalto="密码不一致"></div>
                  </div>
                  <div class="form-group must">
                    <label class="col-sm-3 control-label">邮箱</label>
                    <div class="col-sm-7">
                      <input type="text" class="form-control" id="Email"></div>
                  </div>
                  <div class="form-group must">
                    <label class="col-sm-3 control-label">性别</label>
                    <div class="col-sm-7">
                      <!-- <input type="text" class="form-control" id="Sex"> -->
                      <select id="Sex" class="form-control">
                        <option value="男">男</option>
                        <option value="女">女</option>
                      </select>
                    </div>
                  </div><div class="form-group must">
                    <label class="col-sm-3 control-label">是否党员</label>
                    <div class="col-sm-7">
                      <select id="IsPartyMember" class="form-control">
                        <option value="true">是</option>
                        <option value="false">否</option>
                      </select>
                    </div>
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
                      <select id="Statusadd" class="form-control" required>
                        <option value="1">显示</option>
                        <option value="2">隐藏</option>
                        <option value="0">禁用</option>
                      </select>
                    </div>
                  </div>
                  <div class="form-group">
                    <label class="col-sm-3 control-label">权限</label>
                    <div class="col-sm-7">
                      <input type="number" class="form-control digits" id="Role" maxlength="20" placeholder="至多20个字符" required></div>
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
    $(function() {
      $('#table0').bootstrapTable({
        idField: 'Id',
        url: '/usermyself',
        // striped: "true",
        columns: [{
            radio: 'true',
            width: '10',
            valign: 'middle',
            align: 'center',
          },
          {
            // field: 'Number',
            title: '序号',
            valign: 'middle',
            align: 'center',
            formatter: function(value, row, index) {
              return index + 1
            }
          }, {
            field: 'name', //这里用user数据库json字段，不能是username
            title: '用户名',
            valign: 'middle',
            align: 'center',
            sortable: 'true',
            // editable: {
            //     type: 'text',
            //     pk: 1,
            //     url: '/v1/wx/updateuser',
            //     title: 'Enter ProjectNumber' 
            // }
          }, {
            field: 'Nickname',
            title: '昵称',
            valign: 'middle',
            align: 'center',
            editable: {
              type: 'text',
              pk: 1,
              url: '/v1/wx/updateuser',
              title: 'Enter ProjectName'
            }
          }, {
            field: 'Password',
            title: '密码',
            valign: 'middle',
            align: 'center',
            editable: {
              type: 'text',
              // type: 'select',
              // source: ["规划", "项目建议书", "可行性研究", "初步设计", "招标设计", "施工图"],
              pk: 1,
              url: '/v1/wx/updateuser',
              title: 'Enter Password'
            }
          }, {
            field: 'Email',
            title: '邮箱',
            valign: 'middle',
            align: 'center',
            // sortable:'true',
            editable: {
              type: 'text',
              pk: 1,
              url: '/v1/wx/updateuser',
              title: 'Enter Email'
            }
          },{
            field: 'Sex',
            valign: 'middle',
            align: 'center',
            visible: true,
            title: '性别',
            // sortable:'true',
            editable: {
              type: 'text',
              pk: 1,
              url: '/v1/wx/updateuser',
              title: 'Enter Sex'
            }
          }, {
            field: 'IsPartyMember',
            valign: 'middle',
            align: 'center',
            visible: true,
            title: '是否党员',
            // sortable:'true',
            editable: {
              type: 'select',
              source: [
                {text: '是', value: true},
                {text: '否', value: false}
              ],
              // type: 'select2',
              // source: [
              //   {text: '是', value: true},
              //   {text: '否', value: false}
              // ],
              // select2: {
              //   allowClear: true,
              //   // width: '150px',
              //   placeholder: '请选择状态',
              //   multiple: true
              // },
              pk: 1,
              url: '/v1/wx/updateuser',
              title: 'Enter IsPartyMember'
            },  
          }, {
            field: 'Department',
            title: '部门',
            valign: 'middle',
            align: 'center',
            editable: {
              type: 'text',
              pk: 1,
              url: '/v1/wx/updateuser',
              title: 'Enter Department'
            }
          }, {
            field: 'Secoffice',
            title: '科室',
            valign: 'middle',
            align: 'center',
            sortable: 'true',
            editable: {
              type: 'text',
              // source: {{.Select2}},//["$1", "$2", "$3"],
              pk: 1,
              url: '/v1/wx/updateuser',
              title: 'Enter Category'
            }
          }, {
            field: 'Ip',
            title: 'IP',
            valign: 'middle',
            align: 'center',
            editable: {
              type: 'text',
              pk: 1,
              url: '/v1/wx/updateuser',
              title: 'Enter Count'
            }
          }, {
            field: 'Port',
            title: '端口号',
            valign: 'middle',
            align: 'center',
            editable: {
              type: 'text',
              pk: 1,
              url: '/v1/wx/updateuser',
              title: 'Enter port'
            }
          }, {
            field: 'Status',
            title: '状态',
            valign: 'middle',
            align: 'center',
            // editable: {
            // type: 'select2',
            //   // source:{{.Userselect}},//'/regist/getuname1',
            // source: [
            //   {id: '1', text: '显示',value:1},
            //   {id: '2', text: '隐藏',value:2},
            //   {id: '3', text: '禁止',value:3}
            // ],
            //   //'[{"id": "1", "text": "One"}, {"id": "2", "text": "Two"}]'
            //   select2: {
            //     allowClear: true,
            //     width: '150px',
            //     placeholder: '请选择状态',
            //     // multiple: true
            //   },//'/regist/getuname1',//这里用get方法，所以要换一个
            //   pk: 1,
            //   url: '/v1/wx/updateuser',
            // title: 'Enter Status'  
            // }
          }, {
            field: 'Lastlogintime',
            title: '最后登录',
            valign: 'middle',
            align: 'center',
            formatter: localDateFormatter,
          }, {
            field: 'Createtime',
            title: '建立',
            valign: 'middle',
            align: 'center',
            formatter: localDateFormatter,
          }, {
            field: 'role',
            title: '权限',
            valign: 'middle',
            align: 'center',
            // editable: {
            //   type: 'select2', 
            //   // source:{{.Userselect}},//'/regist/getuname1',
            //   source: [
            //     {id: '1', text: '1级',value:1},
            //     {id: '2', text: '2级',value:2},
            //     {id: '3', text: '3级',value:3}
            //   ],
            //   //'[{"id": "1", "text": "One"}, {"id": "2", "text": "Two"}]'
            //   select2: {
            //     allowClear: true,
            //     width: '150px',
            //     placeholder: '请选择权限',
            //     // multiple: true
            //   },//'/regist/getuname1',//这里用get方法，所以要换一个
            //   pk: 1,
            //   url: '/v1/wx/updateuser',
            //   title: 'Enter Status'  
            // }
          }
        ]
      });
    });

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
    </script>
    <!-- 显示用户角色表 -->
    <div id="details" style="display:none">
      <div class="row">
        <div id="h-role-info" class="col-sm-6 col-md-6 col-lg-6">
          <h3 id="rowtitle">角色表</h3>
          <div id="toolbar1" class="btn-group">
            <!-- <button type="button" id="editorButton" class="btn btn btn-primary btn-sm"> <i class="fa fa-edit">保存修改</i>
        </button> -->
            <!-- <button type="button" data-name="editorButton" id="editorButton" class="btn btn btn-primary btn-sm"> <i class="fa fa-edit">编辑</i>
        </button>
        <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn btn-danger btn-sm">
        <i class="fa fa-trash">删除</i>
        </button> -->
          </div>
          <table id="table1" data-toggle="table" data-striped="true" data-click-to-select="false" data-page-list="[5, 25, 50, All]" data-search="false">
            <thead>
              <tr>
                <th data-width="10" data-checkbox="true" data-formatter="stateFormatter"></th>
                <th data-formatter="index1">#</th>
                <th data-field="Rolenumber">角色编码</th>
                <th data-field="name">角色名称</th>
                <th data-align="center" data-formatter="StatusFormatter">状态</th>
                <!-- <th data-field="domain_desc">所属域</th> -->
                <!-- <th data-align="center"
                    data-field="create_user">创建人</th> -->
                <!-- <th data-align="center"
                    data-field="Createtime" data-formatter="localDateFormatter">创建时间</th> -->
                <!-- <th data-align="center"
                    data-field="modify_user">修改人</th> -->
                <!-- <th data-align="center"
                    data-field="Updated" data-formatter="localDateFormatter">修改时间</th> -->
                <!-- <th data-field="state-handle"
                    data-align="center"
                    data-formatter="RoleObj.formatter">资源操作</th> -->
              </tr>
            </thead>
          </table>
        </div>
      </div>
    </div>
    <script type="text/javascript">
    function stateFormatter(value, row, index) {
      if (row.Level === "1") {
        return {
          disabled: true,
          checked: true
        }
      } else {
        return {
          disabled: true
        }
      }
      return value;
    }

    function StatusFormatter(value, row, index) {
      // alert(row.Status);
      if (row.status == "0") {
        return '正常';
      } else {
        return '失效';
      }
    }
    </script>
  </div>
</body>

</html>
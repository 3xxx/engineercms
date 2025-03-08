<!-- 用户登录后自己的资料列表-->
<!DOCTYPE html>
<html>

<head>
  <meta charset="UTF-8">
  <title>EngineerCMS</title>
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>

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

</head>
<div class="container-fill">{{template "navbar" .}}</div>
<body>
  <div class="col-lg-12">
    <h3>用户表-{{.Username}}</h3>
    <div id="toolbar1" class="btn-group">
      <button type="button" data-name="addButton" id="addButton" class="btn btn-default"> <i class="fa fa-plus">添加</i>
      </button>
      <button type="button" data-name="importButton" id="importButton" class="btn btn-default"> <i class="fa fa-plus">导入</i>
      </button>
      <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
      </button>
    </div>
    <table id="table0" data-search="true" data-show-refresh="true" data-show-toggle="true" data-show-columns="true" data-striped="true" data-toolbar="#toolbar1" data-query-params="queryParams" data-sort-name="Username" data-sort-order="desc" data-page-size="5" data-page-list="[5, 25, 50, All]" data-unique-id="id" data-pagination="true" data-side-pagination="client" data-single-select="true" data-click-to-select="true" data-show-export="true">
    </table>
    <!-- 显示用户角色表 -->
    <div id="details" style="display:none">
      <div class="row">
        <div id="h-role-info" class="col-sm-6 col-md-6 col-lg-6">
          <h3 id="rowtitle">角色表</h3>
          <div id="toolbar1" class="btn-group">
          </div>
          <table id="table1" data-toggle="table" data-striped="true" data-click-to-select="false" data-page-list="[5, 25, 50, All]" data-search="false">
            <thead>
              <tr>
                <th data-width="10" data-checkbox="true" data-formatter="stateFormatter"></th>
                <th data-formatter="index1">#</th>
                <th data-field="Rolenumber">角色编码</th>
                <th data-field="name">角色名称</th>
                <th data-align="center" data-formatter="StatusFormatter">状态</th>
              </tr>
            </thead>
          </table>
        </div>
      </div>
    </div>
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
            editable: {
                type: 'text',
                pk: 1,
                url: '/v1/wx/updateuser',
                title: 'Enter UserName',
                // mode: "popup",
                // emptytext: "--"
                success: function(response, newValue) {
                  // console.log(response.data)
                  if(response.info=='ERROR'){
                    return response.data;
                  }
                  // if(response=='7'){
                  //   return "对不起，您无此操作权限！";
                  // }
                  if(response.status =='error') {
                    return response.msg;
                  }
                }
            }
          }, {
            field: 'Nickname',
            title: '昵称',
            valign: 'middle',
            align: 'center',
            editable: {
              type: 'text',
              pk: 1,
              url: '/v1/wx/updateuser',
              title: 'Enter NickName'
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
              title: 'Enter Secoffice'
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
              title: 'Enter IP'
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
        ],
        // onEditableSave: function (field, row, oldValue, $el) {
        //   $.ajax({
        //     type: "post",
        //     url: "/v1/wx/updateuser",
        //     data: {"pk":row.id,
                        // "name":field,
                        // "oldValue":oldValue,
                        // "newValue":row[field]},
        //     dataType: 'JSON',
        //     success: function (data, status) {
        //       if (status == "success" && data.info == "SUCCESS") {
        //         alert('修改成功');
        //       }else if (status == "success" && data.info == "ERROR"){
        //         alert(data.data);
        //       }
        //     },
        //     error: function () {
        //       alert('返回失败，编辑失败');
        //     },
        //     complete: function () {
        //     }
        //   });
        // }
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
</body>

</html>
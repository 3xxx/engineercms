<!DOCTYPE html>
<html lang="zh-CN">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no">
  <title>用户管理后台</title>
  <meta name="theme-color" content="#2C2E2F">
  <meta name="keywords" content="">
  <meta name="description" content="">
  <link rel="shortcut icon" href="/static/pass_project/favicon.png">
  <link rel="stylesheet" id="font-awesome-css" href="/static/pass_project/font-awesome.min.css" type="text/css" media="all">
  <link rel="stylesheet" id="bootstrap-css" href="/static/pass_project/bootstrap.css" type="text/css" media="all">
  <link rel="stylesheet" id="nav-css" href="/static/pass_project/nav.css" type="text/css" media="all">
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css"/>
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
</head>

<body class="page-body white">
  <div class="page-container">
    <div class="sidebar-menu toggle-others fixed" style="">
      <div class="sidebar-menu-inner ps-container ps-active-y">
        <header class="logo-env">
          <!-- logo -->
          <div class="logo">
            <a href="/v1/passproject/getpassproject" class="logo-expanded">
              <img src="/static/pass_project/logo@2x.png" height="40" alt="EngineerCMS">
            </a>
            <a href="/v1/passproject/getpassproject" class="logo-collapsed">
              <img src="/static/pass_project/logo-collapsed@2x.png" height="40" alt="EngineerCMS">
            </a>
          </div>
          <div class="mobile-menu-toggle visible-xs">
            <a href="/v1/passproject/getpassproject#" data-toggle="mobile-menu">
              <i class="fa fa-bars"></i>
            </a>
          </div>
        </header>
        <ul id="main-menu" class="main-menu">
          <li class="active">
            <a href="#" class="smooth">
              <i class="fa fa-user-circle-o fa-fw"></i>
              <span class="title">个人资料</span>
            </a>
          </li>
          <li>
            <a href="/v1/wx/getuserwxpaymath" class="smooth">
              <i class="fa fa-jpy fa-fw"></i>
              <span class="title">收支记录</span>
            </a> 
          </li>
          <li>
            <a href="/v1/wx/wxrecharge" class="smooth">
              <i class="fa fa-podcast fa-fw"></i>
              <span class="title">充值记录</span>
            </a>
          </li>
          <li>
            <a href="/v1/wx/wxbalance" class="smooth">
              <i class="fa fa-podcast fa-fw"></i>
              <span class="title">余额表</span>
            </a>
          </li>
          <li>
            <a href="/v1/wx/wxincome" class="smooth">
              <i class="fa fa-podcast fa-fw"></i>
              <span class="title">收益记录</span>
            </a>
          </li>
          <li>
            <a href="/v1/wx/collection" class="smooth">
              <i class="fa fa-star-o fa-fw"></i>
              <span class="title">收藏</span>
            </a>
          </li>
          <li class="has-sub">
            <a>
              <i class="fa fa-lightbulb-o fa-fw"></i>
              <span class="title">我的成果</span>
            </a>
            <ul>
              <li>
                <a href="/v1/passproject/getpassproject" class="smooth">PASS计算书-MathCAD</a>
              </li>
              <li>
                <a href="/v1/passproject/getpassproject" class="smooth">PASS计算书-Excel</a>
              </li>
              <li>
                <a href="/v1/passproject/getpassproject" class="smooth">PASS计算书-Ansys APDL</a>
              </li>
              <li>
                <a href="/v1/passproject/getpassproject" class="smooth">FreeCAD参数化模型</a>
              </li>
            </ul>
          </li>
          <li class="has-sub">
            <a>
              <i class="fa fa-folder-open-o fa-fw"></i>
              <span class="title">供水工程</span>
            </a>
            <ul>
              <li>
                <a href="/v1/passproject/getpassproject" class="smooth">水力学</a>
              </li>
              <li>
                <a href="/v1/passproject/getpassproject#Mockup" class="smooth">岩土工程</a>
              </li>
            </ul>
          </li>
          <li class="has-sub">
            <a>
              <i class="fa fa-bullseye fa-fw"></i>
              <span class="title">项目策划</span>
            </a>
            <ul>
              <li>
                <a href="/v1/passproject/getpassproject" class="smooth">项目建议书</a>
              </li>
              <li>
                <a href="/v1/passproject/getpassproject" class="smooth">可行性研究</a>
              </li>
            </ul>
          </li>
          <li>
            <a href="/v1/passproject/getpassproject" class="smooth">
              <i class="fa fa-heart-o fa-fw"></i>
              <span class="title">结构团队</span>
            </a>
          </li>
          <li class="submit-tag">
          </li>
          <li id="menu-item-837" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-837"><a target="_blank" rel="noopener" href="/v1/math/uploadtemple/26997">
              <i class="fa fa-download fa-fw"></i>
              <span class="smooth">下载计算书</span>
              <span class="label label-Primary pull-right hidden-collapsed">♥</span></a></li>
          <li id="menu-item-801" class="menu-item menu-item-type-post_type menu-item-object-page menu-item-801"><a href="/v1/math/getmath/26997">
              <i class="fa fa-newspaper-o fa-fw"></i>
              <span class="smooth">上传计算书</span>
              <span class="label label-Primary pull-right hidden-collapsed">♥</span></a></li>
        </ul>
        <div class="ps-scrollbar-x-rail" style="left: 0px; bottom: 3px;">
          <div class="ps-scrollbar-x" style="left: 0px; width: 0px;"></div>
        </div>
        <div class="ps-scrollbar-y-rail" style="top: 0px; height: 438px; right: 2px;">
          <div class="ps-scrollbar-y" style="top: 0px; height: 300px;"></div>
        </div>
      </div>
    </div>
    <div class="main-content" style="">
      <nav class="navbar user-info-navbar" role="navigation">
        <div class="navbar-content">
          <ul class="user-info-menu left-links list-inline list-unstyled">
            <li class="hidden-xs">
              <a href="/v1/passproject/getpassproject#" data-toggle="sidebar">
                <i class="fa fa-bars"></i>
              </a>
            </li>
            <!-- 天气 -->
            <!-- <li>
              <div data-v-41ba7e2c="" id="he-plugin-simple" style="position: relative;">
                <div data-v-41ba7e2c="" class="s-sticker" style="font-size: 1.4em; cursor: pointer; padding: 30px 10px; background-color: transparent; border-radius: 5px;">
                </div>
                <div data-v-41ba7e2c="">
                </div>
              </div>
              <script>
                WIDGET = { CONFIG: { "modules": "12034", "background": 5, "tmpColor": "aaa", "tmpSize": 16, "cityColor": "aaa", "citySize": 16, "aqiSize": 16, "weatherIconSize": 24, "alertIconSize": 18, "padding": "30px 10px 30px 10px", "shadow": "1", "language": "auto", "borderRadius": 5, "fixed": "false", "vertical": "middle", "horizontal": "left", "key": "a922adf8928b4ac1ae7a31ae7375e191" } }
              </script>
              <script src="/static/pass_project/he-simple-common.js"></script>
            </li> -->
            <!-- 天气 end -->
          </ul>
        </div>
      </nav>

      <div class="sites-list" style="margin-bottom: 8.5rem;">
        <!-- <div id="search" class="s-search">搜索</div> -->
        <h4 class="text-gray" style="display: inline-block;"><i class="icon-io-tag" style="margin-right: 27px;" id="个人信息"></i>个人信息</h4>
        <div class="row">
          <div class="xe-card col-sm-12 col-md-12 col-lg-12">
            <h3>用户表-{{.Username}}</h3>
            <div id="toolbar1" class="btn-group">
              <button type="button" data-name="addButton" id="addButton" class="btn btn-default" disabled="true"> <i class="fa fa-plus">添加</i>
              </button>
              <button type="button" data-name="importButton" id="importButton" class="btn btn-default" disabled="true"> <i class="fa fa-plus">导入</i>
              </button>
              <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default" disabled="true"><i class="fa fa-trash">删除</i></button>
            </div>
            <table id="table0" data-search="true" data-show-refresh="true" data-show-toggle="true" data-show-columns="true" data-striped="true" data-toolbar="#toolbar1" data-query-params="queryParams" data-sort-name="Username" data-sort-order="desc" data-page-size="5" data-page-list="[5, 25, 50, All]" data-unique-id="id" data-pagination="true" data-side-pagination="client" data-single-select="true" data-click-to-select="true" data-show-export="true">
            </table>
            <!-- 显示用户角色表 -->
            <div id="details" style="display:none">
              <!-- <div class="row"> -->
                <!-- <div id="h-role-info" class="col-sm-6 col-md-6 col-lg-6"> -->
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
                <!-- </div> -->
              <!-- </div> -->
            </div>
          </div>
        </div>
      </div>
    </div>

  </div>
  <!-- <script type="text/javascript" src="/static/pass_project/jquery-1.11.1.min.js" id="jquery-js"></script> -->
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <!-- <script type="text/javascript" src="/static/pass_project/bootstrap.min.js" id="bootstrap-js"></script> -->
  <script type="text/javascript" src="/static/pass_project/TweenMax.min.js" id="TweenMax-js"></script>
  <script type="text/javascript" id="appjs-js-extra">
    var theme = { "ajaxurl": "http:\/\/webstack.iotheme.cn\/wp-admin\/admin-ajax.php", "addico": "http:\/\/webstack.iotheme.cn\/wp-content\/themes\/WebStack-1.1422\/images\/add.png", "version": "1.1522" };
  </script>
  <script type="text/javascript" src="/static/pass_project/app.js" id="appjs-js"></script>
  <script type="text/javascript" src="/static/pass_project/lazyload.min.js" id="lazyload-js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>

  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <script type="text/javascript" src="/static/js/tableExport.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <script type="text/javascript">
    $(function() {
      $(document).on('click', '.has-sub', function() {
        var _this = $(this)
        if (!$(this).hasClass('expanded')) {
          setTimeout(function() {
            _this.find('ul').attr("style", "")
          }, 300);
        } else {
          $('.has-sub ul').each(function(id, ele) {
            var _that = $(this)
            if (_this.find('ul')[0] != ele) {
              setTimeout(function() {
                _that.attr("style", "")
              }, 300);
            }
          })
        }
      })
      $('.user-info-menu .hidden-xs').click(function() {
        if ($('.sidebar-menu').hasClass('collapsed')) {
          $('.has-sub.expanded > ul').attr("style", "")
        } else {
          $('.has-sub.expanded > ul').show()
        }
      })
      $("#main-menu li ul li").click(function() {
        $(this).siblings('li').removeClass('active'); // 删除其他兄弟元素的样式
        $(this).addClass('active'); // 添加当前元素的样式
      });
      $("a.smooth").click(function(ev) {
        ev.preventDefault();
        if ($("#main-menu").hasClass('mobile-is-visible') != true)
          return;
        public_vars.$mainMenu.add(public_vars.$sidebarProfile).toggleClass('mobile-is-visible');
        ps_destroy();
        $("html, body").animate({
          scrollTop: $($(this).attr("href")).offset().top - 80
        }, {
          duration: 500,
          easing: "swing"
        });
      });
      return false;
    });

    var href = "";
    var pos = 0;
    $("a.smooth").click(function(e) {
      e.preventDefault();
      if ($("#main-menu").hasClass('mobile-is-visible') === true)
        return;
      $("#main-menu li").each(function() {
        $(this).removeClass("active");
      });
      $(this).parent("li").addClass("active");
      href = $(this).attr("href");
      window.location.href=href;
      // window.open(href)
      // pos = $(href).position().top - 100;
      // $("html,body").animate({
      //   scrollTop: pos
      // }, 500);
    });

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
              title: 'Enter NickName',
              success: function(response, newValue) {
                // console.log(response.data)
                if(response.info=='ERROR'){
                  return response.data;
                }
                // if(response=='7'){
                //   return "对不起，您无此操作权限！";
                // }
                if(response.status =='ERROR') {
                  return response.msg;
                }
              }
            }
          }, {
            field: 'Password',
            title: '密码',
            valign: 'middle',
            align: 'center',
            visible: false,
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
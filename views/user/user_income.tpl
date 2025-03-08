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
          <li>
            <a href="/user" class="smooth">
              <i class="fa fa-user-circle-o fa-fw"></i>
              <span class="title">个人资料</span>
            </a>
          </li>
          <li>
            <a href="#" class="smooth">
              <i class="fa fa-jpy fa-fw"></i>
              <span class="title">消费记录</span>
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
          <li class="active">
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
        <h4 class="text-gray" style="display: inline-block;"><i class="icon-io-tag" style="margin-right: 27px;" id="个人信息"></i>个人收益</h4>
        <div class="row">
          <div class="xe-card col-sm-12 col-md-12 col-lg-12">
            <h3>收益记录</h3>
            <div id="toolbar1" class="btn-group">
              <button type="button" data-name="addButton" id="addButton" class="btn btn-default" disabled="true"> <i class="fa fa-plus">添加</i>
              </button>
              <button type="button" data-name="importButton" id="importButton" class="btn btn-default" disabled="true"> <i class="fa fa-plus">导入</i>
              </button>
              <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default" disabled="true"><i class="fa fa-trash">删除</i></button>
            </div>
            <table id="table0"></table>
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
        // idField: 'ID',
        url: '/v1/wx/wxincomedata',
        method: 'get',
        search: 'true',
        classes: "table table-striped", //这里设置表格样式
        showRefresh: 'true',
        showColumns: 'true',
        toolbar: '#toolbar1',
        pagination: 'true',
        sidePagination: "server",
        queryParamsType: '',
          //请求服务器数据时，你可以通过重写参数的方式添加一些额外的参数，例如 toolbar 中的参数 如果 queryParamsType = 'limit' ,返回参数必须包含
          // limit, offset, search, sort, order 否则, 需要包含:
          // pageSize, pageNumber, searchText, sortName, sortOrder.
          // 返回false将会终止请求。
        pageSize: 15,
        pageNumber: 1,
        pageList: [15, 50, 100, 'All'],
        uniqueId: "ID",
        idField: 'ID',
        // singleSelect:"true",
        clickToSelect: "true",
        showExport: "true",
        queryParams: function queryParams(params) { //设置查询参数
          var param = {
            limit: params.pageSize, //每页多少条数据
            pageNo: params.pageNumber, // 页码
            searchText: params.searchText // $(".search .form-control").val()
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
        // striped: "true",
        columns: [{
            radio: 'true',
            width: '10',
            valign: 'middle',
            align: 'center',
          },
          {
            title: '序号',
            valign: 'middle',
            align: 'center',
            formatter: function(value, row, index) {
              return index + 1
            }
          }, {
            field: 'TempleUser.Nickname', //这里用user数据库json字段，不能是username
            title: '模板作者',
            valign: 'middle',
            align: 'center',
            sortable: 'true'
          }, {
            field: 'Temple.number',
            title: '模板编号',
            valign: 'middle',
            align: 'center',
          }, {
            field: 'Temple.titleb',
            title: '模板名称',
            valign: 'middle',
            align: 'center'
          }, {
            field: 'Temple.version',
            title: '模板版本号',
            valign: 'middle',
            align: 'center'
          }, {
            field: 'Temple.price',
            title: '模板单价（元）',
            valign: 'middle',
            align: 'center',
            formatter: amountFormatter,
          }, {
            field: 'PayUser.name',
            title: '支付者用户名',
            valign: 'middle',
            align: 'center'
          }, {
            field: 'PayUser.Nickname',
            title: '支付者昵称',
            valign: 'middle',
            align: 'center'
          }, {
            field: 'payer_openid',
            title: '微信支付openid',
            valign: 'middle',
            align: 'center',
            visible: false,
          }, {
            field: 'amount',
            valign: 'middle',
            align: 'center',
            visible: true,
            title: '支付数额（元）',
            formatter: amountFormatter,
            // sortable:'true'
          }, {
            field: 'out_trade_no',
            valign: 'middle',
            align: 'center',
            visible: true,
            title: '商家订单号',
            // sortable:'true' 
          }, {
            field: 'transaction_id',
            title: '微信支付订单号',
            valign: 'middle',
            align: 'center',
            visible: false,
            sortable: 'true'
          }, {
            field: 'attach',
            title: '订单信息',
            valign: 'middle',
            visible: false,
            align: 'center'
          }, {
            field: 'success_time',
            title: '订单支付时间',
            valign: 'middle',
            align: 'center'
          }, {
            field: 'trade_state',
            title: '支付状态',
            valign: 'middle',
            align: 'center'
          }, {
            field: 'cal_finished',
            title: '完成计算？',
            valign: 'middle',
            align: 'center'
          }, {
            field: 'CreatedAt',
            title: '记录时间',
            valign: 'middle',
            align: 'center',
            visible: false,
          }
        ],
      });
    });

    function index1(value, row, index) {
      return index + 1
    }

    /**
   * 数字除以100
   */
  function amountFormatter(value) {
    return parseFloat(value/100);
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
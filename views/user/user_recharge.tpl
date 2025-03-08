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
  <!-- <link rel="stylesheet" id="font-awesome-css" href="/static/pass_project/font-awesome.min.css" type="text/css" media="all"> -->
  <!-- <link rel="stylesheet" id="bootstrap-css" href="/static/pass_project/bootstrap.css" type="text/css" media="all"> -->
  <link rel="stylesheet" id="nav-css" href="/static/pass_project/nav.css" type="text/css" media="all">
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css" />
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css" />
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css" /> -->
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
  <style>
    /*:root {
            --primary-color: #07c160;
            --bg-color: #f5f5f5;
            --card-bg: #fff;
            --text-color: #333;
        }*/

        /** {
            box-sizing: border-box;
            margin: 0;
            padding: 0;
        }*/

        html { font-size: 1em; }

        .recharge-body {
          --primary-color: #07c160;
            --bg-color: #f5f5f5;
            --card-bg: #fff;
            --text-color: #333;

            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 
                'Helvetica Neue', Arial, sans-serif;
            /*background: var(--bg-color);*/
            color: var(--text-color);
            /*min-height: 100vh;*/
            display: flex;
            justify-content: center;
            align-items: center;
        }

        .recharge-container {
            background: var(--card-bg);
            padding: 2rem;
            border-radius: 12px;
            box-shadow: 0 4px 12px rgba(0,0,0,0.1);
            width: 100%;
            max-width: 400px;
            text-align: center;
        }

        .recharge-title {
            font-size: 1.5rem;
            margin-bottom: 1.5rem;
            color: var(--primary-color);
        }

        .recharge-input-group {
            margin-bottom: 1.5rem;
        }

        .recharge-amount-input {
            width: 100%;
            padding: 0.8rem;
            border: 2px solid #eee;
            border-radius: 8px;
            font-size: 1.1rem;
            text-align: center;
            transition: border-color 0.3s ease;
        }

        .recharge-amount-input:focus {
            outline: none;
            border-color: var(--primary-color);
        }

        .recharge-submit-btn {
            background: var(--primary-color);
            color: white;
            border: none;
            padding: 0.8rem 1.5rem;
            border-radius: 8px;
            font-size: 1rem;
            cursor: pointer;
            transition: opacity 0.3s ease;
            width: 100%;
        }

        .recharge-submit-btn:disabled {
            opacity: 0.7;
            cursor: not-allowed;
        }

        .recharge-qrcode-container {
            margin: 1.5rem 0;
            padding: 1rem;
            background: #f8f8f8;
            border-radius: 8px;
            display: none;
        }

        .recharge-qrcode-img {
            width: 200px;
            height: 200px;
            margin: 0 auto;
        }

        .recharge-status-message {
            margin-top: 3rem;
            color: #666;
            font-size: 0.9rem;
        }

        .recharge-loading {
            display: inline-block;
            width: 20px;
            height: 20px;
            border: 3px solid rgba(0,0,0,0.1);
            border-radius: 50%;
            border-top-color: var(--primary-color);
            animation: spin 1s ease-in-out infinite;
        }

        @keyframes spin {
            to { transform: rotate(360deg); }
        }
  </style>
</head>

<body class="page-body white">
  <div class="page-container">
    <div class="sidebar-menu toggle-others fixed" style="">
      <div class="sidebar-menu-inner ps-container ps-active-y">
        <header class="logo-env">
          <!-- logo -->
          <div class="logo">
            <a href="#" class="logo-expanded">
              <img src="/static/pass_project/logo@2x.png" height="40" alt="EngineerCMS">
            </a>
            <a href="#" class="logo-collapsed">
              <img src="/static/pass_project/logo-collapsed@2x.png" height="40" alt="EngineerCMS">
            </a>
          </div>
          <div class="mobile-menu-toggle visible-xs">
            <a href="##" data-toggle="mobile-menu">
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
            <a href="/v1/wx/getuserwxpaymath" class="smooth">
              <i class="fa fa-jpy fa-fw"></i>
              <span class="title">消费记录</span>
            </a> 
          </li>
          <li class="active">
            <a href="#" class="smooth">
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
                <a href="#" class="smooth">PASS计算书-Excel</a>
              </li>
              <li>
                <a href="#" class="smooth">PASS计算书-Ansys</a>
              </li>
              <li>
                <a href="#" class="smooth">FreeCAD参数化模型</a>
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
                <a href="#" class="smooth">水力学</a>
              </li>
              <li>
                <a href="#" class="smooth">岩土工程</a>
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
                <a href="#" class="smooth">项目建议书</a>
              </li>
              <li>
                <a href="#" class="smooth">可行性研究</a>
              </li>
            </ul>
          </li>
          <li>
            <a href="https://www.iotheme.cn/store/webstack.html" class="smooth">
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
              <a href="#" data-toggle="sidebar">
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

    <!-- 模态框 支付二维码 -->
  <div class="container">
    <form class="form-horizontal">
      <div class="modal fade" id="qrcode_modal" style="margin-top: 100px;">
        <div class="modal-dialog" style="width:640px">
          <div class="modal-content">
            <div class="modal-header">
              <button type="button" class="close" data-dismiss="modal">
                <span aria-hidden="true">&times;</span>
              </button>
              <h3 class="modal-title">微信扫码支付</h3>
            </div>
            <div class="modal-body">
              <div class="modal-body-content">
                <!-- <h1 class="mod-title">
                    <span class="ico-wechat"></span><span class="text">微信扫码支付</span>
                  </h1> -->
                <div class="recharge-body">
                  <div class="recharge-container">
                    <h1 class="recharge-title">微信扫码充值</h1>
                    <div class="recharge-input-group">
                      <input type="number" class="recharge-amount-input" placeholder="请输入充值金额（元）" min="0.1" step="0.1">
                    </div>
                    <button class="recharge-submit-btn" onclick="handleRecharge()">
                      生成充值二维码
                    </button>
                    <div class="recharge-qrcode-container">
                      <!-- <img class="qrcode-img" id="qrcodeImg"> -->
                      <div class="recharge-qrcode-img" id="qrcodeImg"></div>
                      <p class="recharge-status-message" id="statusMessage"></p>
                    </div>
                  </div>
                </div>
                <!-- <div class="foot">
                  <div class="inner">
                    <p>手机用户可保存上方二维码到手机中</p>
                    <p>在微信扫一扫中选择“相册”即可</p>
                  </div>
                </div> -->
              </div>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
              <!-- <button type="button" class="btn btn-primary" onclick="save()">保存</button> -->
            </div>
          </div>
        </div>
      </div>
    </form>
  </div>

      <div class="sites-list" style="margin-bottom: 8.5rem;">
        <!-- <div id="search" class="s-search">搜索</div> -->
        <h4 class="text-gray" style="display: inline-block;"><i class="icon-io-tag" style="margin-right: 27px;" id="个人信息"></i>个人信息</h4>
        <div class="row">
          <div class="xe-card col-sm-12 col-md-12 col-lg-12">
            <h3>充值记录表-{{.Username}}</h3>
            <div id="toolbar1" class="btn-group">
              <button type="button" data-name="addButton" id="addButton" class="btn btn-default" onclick="qrcodemodal()"> <i class="fa fa-plus">充值</i>
              </button>
              <button type="button" data-name="importButton" id="importButton" class="btn btn-default"> <i class="fa fa-minus">退款</i>
              </button>
              <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
                <i class="fa fa-trash">删除</i>
              </button>
            </div>
            <table id="table0">
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
  <script src="/static/vue.js/qrcode.min.js"></script>
  <script type="text/javascript">
  function qrcodemodal() {
    // 先查询服务端是否够费用，如果够，则进行计算和扣费；如果不够，则弹出二维码
    // 销毁二维码和取消订单
    if (qrcode) {
      qrcode.clear();
    }

    $('#qrcode_modal').modal({
      show: true,
      backdrop: 'static'
    });
    // getqrcode();
  }

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
    window.location.href = href;
    // window.open(href)
    // pos = $(href).position().top - 100;
    // $("html,body").animate({
    //   scrollTop: pos
    // }, 500);
  });

  $(function() {
    $('#table0').bootstrapTable({
      // idField: 'ID',
      url: '/v1/wx/wxrechargelist',
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
            // field: 'Number',
            title: '序号',
            valign: 'middle',
            align: 'center',
            formatter: function(value, row, index) {
              return index + 1
            }
          }, {
            field: 'User.Nickname',
            title: '支付者昵称',
            valign: 'middle',
            align: 'center'
          }, {
            field: 'User.name',
            title: '支付者用户名',
            valign: 'middle',
            align: 'center'
          }, {
            field: 'amount',
            valign: 'middle',
            align: 'center',
            visible: true,
            title: '充值数额（元）',
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
            field: 'payer_openid',
            title: '微信支付openid',
            valign: 'middle',
            align: 'center',
            visible: false,
          }, {
            field: 'transaction_id',
            title: '微信支付订单号',
            valign: 'middle',
            align: 'center',
            sortable: 'true'
          }, {
            field: 'success_time',
            title: '订单支付时间',
            valign: 'middle',
            align: 'center'
          }, {
            field: 'attach',
            title: '订单信息',
            valign: 'middle',
            align: 'center'
          }, {
            field: 'CreatedAt',
            title: '时间',
            valign: 'middle',
            align: 'center'
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

  // DOM元素
  const amountInput = document.querySelector('.recharge-amount-input');
  const qrcodeContainer = document.querySelector('.recharge-qrcode-container');
  const qrcodeImg = document.getElementById('qrcodeImg');
  const statusMessage = document.getElementById('statusMessage');
  const submitBtn = document.querySelector('.recharge-submit-btn');

  let currentOrderId = null;
  let checkInterval = null;

  var out_trade_no = "";

  var qrcode = new QRCode(document.getElementById("qrcodeImg"), {
    width: 230,
    height: 230,
    foreground: "#000000",
    background: "#ffffff",
    typeNumber: QRCode.CorrectLevel.L
  });
    // 由低到高
// QRCode.CorrectLevel.L
// QRCode.CorrectLevel.M
// QRCode.CorrectLevel.Q
// QRCode.CorrectLevel.H
  var timer = null;

  async function handleRecharge() {
    const amount = parseFloat(amountInput.value);
    // 输入验证
    if (!amount || amount < 0.01) {
      alert('请输入有效的充值金额，不能小于1分钱');
      return;
    }

    // 显示加载状态
    submitBtn.disabled = true;
    submitBtn.innerHTML = '<div class="loading"></div>';

    $.ajax({
      url: '/v1/wx/wxnativepay',
      type: 'post',
      dataType: 'json',
      data: { payamount: parseFloat(amount * 100), out_trade_no: "recharge", description: "recharge", attach: "recharge" }, //post数据
      success: function(data, textStatus) {
        //从服务器得到数据，显示数据并继续查询
        if (data.CodeUrl && data.CodeUrl !== '') {
          if (qrcode !== null) {
            qrcode.clear();
          }
          qrcode.makeCode(data.CodeUrl);
          // qrcodeImg.src = response.data.qrcodeUrl;
          qrcodeContainer.style.display = 'block';
          // 赋值订单号和生成订单时间
          // document.getElementById("billId").innerHTML = data.out_trade_no;
          out_trade_no = data.out_trade_no;
          // document.getElementById("createTime").innerHTML =data.createtime;
        };

        // loadmsg();
        window.timer = setInterval(() => {
          setTimeout("loadmsg()", 0)
        }, 2000)
      },
      //Ajax请求超时，继续查询
      error: function(XMLHttpRequest, textStatus, errorThrown) {
        if (textStatus == "timeout") {
          // setTimeout("loadmsg()", 1000);
        } else { //异常
          // setTimeout("loadmsg()", 4000);
        }
      },
      complete: function(xhr, status) {
        submitBtn.disabled = false;
        submitBtn.textContent = '生成充值二维码';
      }
      // success(result,status,xhr)
    })

    // try {
    //   // 调用获取二维码接口
    //   const response = await mockServer.getRechargeQrcode(amount);

    //   if (response.code === 200) {
    //     // 显示二维码
    //     qrcodeImg.src = response.data.qrcodeUrl;
    //     qrcodeContainer.style.display = 'block';
    //     currentOrderId = response.data.orderId;

    //     // 开始轮询支付状态
    //     startPaymentPolling();
    //   }
    // } catch (error) {
    //   console.error('充值请求失败:', error);
    //   alert('生成二维码失败，请重试');
    // } finally {
    //   submitBtn.disabled = false;
    //   submitBtn.textContent = '生成充值二维码';
    // }
  }

  function startPaymentPolling() {
    checkInterval = setInterval(async () => {
      try {
        const res = await mockServer.checkPaymentStatus(currentOrderId);

        if (res.paid) {
          clearInterval(checkInterval);
          statusMessage.innerHTML = '✅ 支付成功！页面即将跳转...';
          setTimeout(() => {
            alert('充值成功！');
            location.reload();
          }, 1500);
        } else {
          statusMessage.textContent = '等待扫码支付中...';
        }
      } catch (error) {
        console.error('支付状态查询失败:', error);
        statusMessage.textContent = '状态查询失败，请稍后重试';
      }
    }, 2000);
  }

  // 检查是否支付完成
  function loadmsg() {
    $.ajax({
      type: "GET",
      dataType: "json",
      url: "/v1/wx/queryorderbyouttradeno",
      timeout: 10000, //ajax请求超时时间10s
      data: { type: "wxNativePay", out_trade_no: out_trade_no }, //post数据
      success: function(data, textStatus) {
        //从服务器得到数据，显示数据并继续查询
        if (data.result.trade_state && data.result.trade_state == "SUCCESS") {
          console.log("支付成功！");

          statusMessage.innerHTML = '✅ 支付成功！页面即将跳转...';
          setTimeout(() => {
            alert('充值成功！');
            location.reload();
          }, 1500);

        } else {
          statusMessage.textContent = '等待扫码支付中...';
          //   // setTimeout("loadmsg()", 3000);
          //   window.timer = setInterval(() => {
          //     setTimeout("loadmsg()", 0)
          //   }, 2000)
        }
      },
      //Ajax请求超时，继续查询
      error: function(XMLHttpRequest, textStatus, errorThrown) {
        if (textStatus == "timeout") {
          // setTimeout("loadmsg()", 1000);
          console.error('超时:', textStatus);
          statusMessage.textContent = '超时，请稍后重试';
        } else { //异常
          // setTimeout("loadmsg()", 4000);
          console.error('支付状态查询失败:', XMLHttpRequest.status);
          statusMessage.textContent = '支付状态查询失败，请稍后重试';
        }
      }
    });
  }

  // 输入验证
  amountInput.addEventListener('input', (e) => {
    let value = e.target.value;
    // 限制只能输入数字和一个小数点
    value = value.replace(/[^0-9.]/g, '');
    // 去除多余的小数点
    value = value.replace(/\.{2,}/g, '.');
    // 限制小数点后两位
    if (value.indexOf('.') >= 0) {
      value = value.substring(0, value.indexOf('.') + 3);
    }
    e.target.value = value;
  });

  // 关闭模态框后，模态框完全消失后触发
  $('#qrcode_modal').on('hidden.bs.modal', function() {
    // 执行一些动作...
    console.log("关闭模态框")
    clearInterval(window.timer)
    // 销毁二维码和取消订单
    if (qrcode !== null) {
      // console.log(qrcode)
      // qrcode.clear();
      $("#qrcodeImg").html("");
    }

  })

  </script>
</body>

</html>
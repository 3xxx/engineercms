<!-- 显示所有项目的搜索结果页面，分页 -->
<!DOCTYPE html>
<html>

<head>
  <meta charset="UTF-8">
  <title>EngineerCMS</title>
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css" />
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css" />
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
  <script src="/static/js/tableExport.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
</head>
<div class="container-fill">{{template "navbar" .}}</div>
<body>
  <div class="col-lg-12">
    <h3 id="rowtitle">搜索结果</h3>
    <div id="toolbar1" class="btn-toolbar" role="toolbar" aria-label="...">
      <div class="btn-group">
        <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="true" title="添加资料">
          <i class="fa fa-plus">&nbsp;&nbsp;添加</i>
          <span class="caret"></span>
        </button>
        <ul class="dropdown-menu" aria-labelledby="">
          <li>
            <a href="javascript:void(0)" onclick="addButton()"><i class="fa fa-plus">&nbsp;&nbsp;单附件模式</i></a>
          </li>
          <li>
            <a href="javascript:void(0)" onclick="addButton1()"><i class="fa fa-plus-square-o">&nbsp;&nbsp;多附件模式</i></a>
          </li>
          <li>
            <a href="javascript:void(0)" onclick="addButton2()"><i class="fa fa-plus-square">&nbsp;&nbsp;文章模式</i></a>
          </li>
        </ul>
      </div>
      <div class="btn-group">
        <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="true" title="编辑">
          <i class="fa fa-edit">&nbsp;&nbsp;编辑</i>
          <span class="caret"></span>
        </button>
        <ul class="dropdown-menu" aria-labelledby="">
          <li>
            <a href="javascript:void(0)" onclick="editorProdButton()"><i class="fa fa-pencil">&nbsp;&nbsp;编辑成果信息</i></a>
          </li>
          <li>
            <a href="javascript:void(0)" onclick="editorAttachButton()"><i class="fa fa-edit">&nbsp;&nbsp;编辑成果附件</i></a>
          </li>
        </ul>
      </div>
      <div class="btn-group">
        <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default" title="删除">
          <i class="fa fa-trash">&nbsp;&nbsp;删除</i>
        </button>
        <button type="button" data-name="shareButton" id="shareButton" class="btn btn-default" title="分享文件">
          <i class="fa fa-share">&nbsp;&nbsp;分享</i>
        </button>
        <button type="button" data-name="flowButton" id="flowButton" class="btn btn-default" title="流程、状态">
          <i class="fa fa-share-alt">&nbsp;&nbsp;Flow</i>
        </button>
      </div>
      <div class="btn-group">
        <button type="button" data-name="cartButton" id="cartButton" class="btn btn-default" title="购物车">
          <i class="fa fa-shopping-cart">&nbsp;&nbsp;Cart</i>
        </button>
      </div>
    </div>
    <table id="table1"></table>
  </div>
  <script type="text/javascript">
  $(function() {
    // 初始化【未接受】工作流表格
    $("#table1").bootstrapTable({
      url: '/v1/wx/searchproductdata?keyword={{.Key}}',
      method: 'get',
      search: 'true',
      classes: "table table-striped", //这里设置表格样式
      showRefresh: 'true',
      showToggle: 'true',
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
      pageList: [15, 50, 100],
      uniqueId: "id",
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
      columns: [{
          title: '选择',
          checkbox: 'true',
          width: '10',
          align: "center",
          valign: "middle"
        },
        {
          // field: 'Number',
          title: '序号',
          formatter: function(value, row, index) {
            return index + 1
          },
          align: "center",
          valign: "middle"
        },
        {
          field: 'Code',
          title: '编号',
          // formatter:setCode,
          halign: "center",
          align: "left",
          valign: "middle"
        },
        {
          field: 'Title',
          title: '名称',
          // formatter:setTitle,
          halign: "center",
          align: "left",
          valign: "middle"
        },
        {
          field: 'Label',
          title: '标签',
          formatter: setLable,
          align: "center",
          valign: "middle"
        },
        {
          field: 'Principal',
          title: '设计',
          align: "center",
          valign: "middle"
        },
        {
          field: 'Articlecontent',
          title: '文章',
          formatter: setArticle,
          events: actionEvents,
          align: "center",
          valign: "middle"
        },
        {
          field: 'Attachmentlink',
          title: '附件',
          formatter: setAttachment,
          events: actionEvents,
          align: "center",
          valign: "middle"
        },
        {
          field: 'Pdflink',
          title: 'PDF',
          formatter: setPdf,
          events: actionEvents,
          align: "center",
          valign: "middle"
        },
        {
          field: 'Created',
          title: '建立时间',
          formatter: localDateFormatter,
          visible: "false",
          align: "center",
          valign: "middle"
        },
        {
          field: 'Updated',
          title: '更新时间',
          formatter: localDateFormatter,
          visible: "false",
          align: "center",
          valign: "middle"
        },
        {
          field: 'Relevancy',
          title: '关联',
          formatter: RelevFormatter,
          // events:actionRelevancy,
          // visible："false",
          align: "center",
          valign: "middle"
        }
      ]
    });
  });
  // 改变点击行颜色
  $(function() {
    $("#table0").on("click-row.bs.table", function(e, row, ele) {
      $(".info").removeClass("info");
      $(ele).addClass("info");
      rowid = row.Id; //全局变量
    });
  });

    $("#search").click(function() { //这里应该用button的id来区分按钮的哪一个,因为本页有好几个button
      var radio = $("input[type='radio']:checked").val();
      $.ajax({
        type: "post", //这里是否一定要用post，是的，因为get会缓存？？
        url: "/v1/wx/searchproductdata",
        data: { keyword: $("#keyword").val(), radiostring: radio },
        success: function(data, status) { //数据提交成功时返回数据
          //显示结果表
          $("#rowtitle").html("搜寻结果");
          $("#details").show();
          $('#table1').bootstrapTable('append', data);
          $('#table1').bootstrapTable('scrollTo', 'bottom');
        }
      });
    });

  function getKey() {
    if (event.keyCode == 13) {
      var radio = $("input[type='radio']:checked").val();
      $.ajax({
        type: "post", //这里是否一定要用post，是的，因为get会缓存？？
        url: "/v1/wx/searchproductdata",
        data: { keyword: $("#keyword").val(), radiostring: radio },
        success: function(data, status) { //数据提交成功时返回数据

          //显示结果表
          $("#rowtitle").html("搜寻结果");
          $("#details").show();
          $('#table1').bootstrapTable('append', data);
          $('#table1').bootstrapTable('scrollTo', 'bottom');
        }
      });
    }
  }

    // 成果添加到购物车
    $("#cartButton").click(function() {
      var selectRow = $('#table1').bootstrapTable('getSelections');
      if (selectRow.length <= 0) {
        alert("请先勾选成果！");
        return false;
      }
      if (selectRow[0].Attachmentlink[0]) { //||selectRow[0].Pdflink[0].Link||selectRow[0].Articlecontent[0].Link)
        var site = /http:\/\/.*?\//.exec(selectRow[0].Attachmentlink[0].Link); //非贪婪模式 
      }
      if (selectRow[0].Articlecontent[0]) {
        var site = /http:\/\/.*?\//.exec(selectRow[0].Articlecontent[0].Link); //非贪婪模式 
      }
      if (selectRow[0].Pdflink[0]) {
        var site = /http:\/\/.*?\//.exec(selectRow[0].Pdflink[0].Link); //非贪婪模式 
      }
      if (site) {
        alert("同步成果不允许！");
        return;
      }
      var title = $.map(selectRow, function(row) {
        return row.Title;
      })
      var ids = "";
      for (var i = 0; i < selectRow.length; i++) {
        if (i == 0) {
          ids = selectRow[i].Id;
        } else {
          ids = ids + "," + selectRow[i].Id;
        }
      }
      $.ajax({
        type: "post",
        url: "/v1/cart/createproductcart",
        data: { ids: ids },
        success: function(data, status) {
          if (data.code=="ERROR"){
            alert(data.msg);
          }else{
            alert("添加“" + data.data[0].Title + "”购物车成功！(status:" + status + "！)");
          }
          // $.toast({
          //   type: TYPES[1],
          //   title: TITLES['info'],
          //   subtitle: '11 mins ago',
          //   content: CONTENT['info'],
          //   delay: 5000
          // });
        }
      });
    })

  function index1(value, row, index) {
    // alert( "Data Loaded: " + index );
    return index + 1
  }

  function localDateFormatter(value) {
    return moment(value, 'YYYY-MM-DD').format('YYYY-MM-DD');
  }
  //关联
  function RelevFormatter(value) {
    if (value) {
      if (value.length == 1) { //'<a href="/project/product/article/'
        var array = value[0].Relevancy.split(",")
        var relevarray = new Array()
        for (i = 0; i < array.length; i++) {
          relevarray[i] = array[i];
        }
        return relevarray.join(",");
        // articleUrl= '<a href="'+value[0].Link+'/'+value[0].Id+'" title="查看" target="_blank"><i class="fa fa-file-text-o"></i></a>';
        // return articleUrl;
      } else if (value.length == 0) {

      } else if (value.length > 1) {
        var relevarray = new Array()
        for (i = 0; i < value.length; i++) {
          relevarray[i] = value[i].Relevancy;
        }
        return relevarray.join(",");
        // articleUrl= "<a class='article' href='javascript:void(0)' title='查看文章列表'><i class='fa fa-list-ol'></i></a>";
        // return articleUrl;
      }
    }
  }

  function setCode(value, row, index) {
    return "<a href='/project/product/attachment/" + row.Id + "'>" + value + "</a>";
  }

  function setLable(value, row, index) {
    // alert(value);
    if (value) { //注意这里如果value未定义则出错，一定要加这个判断。
      var array = value.split(",")
      var labelarray = new Array()
      for (i = 0; i < array.length; i++) {
        labelarray[i] = "<a href='/project/product/keysearch?keyword=" + array[i] + "'>" + array[i] + "</a>";
      }
      return labelarray.join(",");
    }
  }

  function setTitle(value, row, index) {
    return "<a href='/project/product/" + row.Id + "'>" + value + "</a>";
  }

  function setArticle(value, row, index) {
    // return '<a class="article" href="javascript:void(0)" title="article"><i class="fa fa-file-text-o"></i></a>';
    if (value) {
      if (value.length == 1) { //'<a href="/project/product/article/'
        articleUrl = '<a href="' + value[0].Link + '/' + value[0].Id + '" title="查看" target="_blank"><i class="fa fa-file-text-o"></i></a>';
        return articleUrl;
      } else if (value.length == 0) {

      } else if (value.length > 1) {
        articleUrl = "<a class='article' href='javascript:void(0)' title='查看文章列表'><i class='fa fa-list-ol'></i></a>";
        return articleUrl;
      }
    }
  }

  function setAttachment(value, row, index) {
    if (value) {
      if (value.length == 1) {
        attachUrl = '<a href="/downloadattachment?id='+ value[0].Id + '" title="下载" target="_blank"><i class="fa fa-paperclip"></i></a>';
        return attachUrl;
      } else if (value.length == 0) {

      } else if (value.length > 1) {
        attachUrl = "<a class='attachment' href='javascript:void(0)' title='查看附件列表'><i class='fa fa-list-ol'></i></a>";
        return attachUrl;
      }
    }
  }

  function setPdf(value, row, index) {
    if (value) {
      if (value.length == 1) {
        pdfUrl = '<a href="/pdf?id=' + value[0].Id + '" title="打开pdf" target="_blank"><i class="fa fa-file-pdf-o"></i></a>';
        return pdfUrl;
      } else if (value.length == 0) {

      } else if (value.length > 1) {
        pdfUrl = "<a class='pdf' href='javascript:void(0)' title='查看pdf列表'><i class='fa fa-list-ol'></i></a>";
        return pdfUrl;
      }
    }
  }

  window.actionEvents = {
    'click .article': function(e, value, row, index) {
      var site = /http:\/\/.*?\//.exec(value[1].Link); //非贪婪模式 
      if (site) {
        $('#articles').bootstrapTable('refresh', { url: '/project/product/syncharticles?site=' + site + '&id=' + row.Id });
      } else {
        $('#articles').bootstrapTable('refresh', { url: '/project/product/articles/' + row.Id });
      }
      $('#modalarticle').modal({
        show: true,
        backdrop: 'static'
      });
    },
    'click .attachment': function(e, value, row, index) {
      // for(var i=0;i<value.length;i++)
      // alert(value[i].Link);
      // var ret=/http:(.*)\:/.exec(value[i].Link);//http://127.0.0.1:
      var site = /http:\/\/.*?\//.exec(value[1].Link); //非贪婪模式 
      if (site) { //跨域
        // alert("1");
        // $.getJSON(ret+'project/product/attachment/'+row.Id,function(){
        // $('#attachs').bootstrapTable('load', randomData());
        // })
        $('#attachs').bootstrapTable('refresh', { url: '/project/product/synchattachment?site=' + site + '&id=' + row.Id });
        // $('#attachs').bootstrapTable('refresh', {url:site+'project/product/attachment/'+row.Id});
      } else {
        // alert("2");
        $('#attachs').bootstrapTable('refresh', { url: '/project/product/attachment/' + row.Id });
      }
      $('#modalattach').modal({
        show: true,
        backdrop: 'static'
      });
    },

    'click .pdf': function(e, value, row, index) {
      var site = /http:\/\/.*?\//.exec(value[1].Link); //非贪婪模式 
      if (site) { //跨域
        $('#pdfs').bootstrapTable('refresh', { url: '/project/product/synchpdf?site=' + site + '&id=' + row.Id });
      } else {
        $('#pdfs').bootstrapTable('refresh', { url: '/project/product/pdf/' + row.Id });
      }
      $('#modalpdf').modal({
        show: true,
        backdrop: 'static'
      });
    },
  };

  //最后面弹出文章列表中用的_根据上面的click，弹出模态框，给模态框中的链接赋值
  function setArticlecontent(value, row, index) {
    articleUrl = '<a href="' + value + '" title="下载" target="_blank"><i class="fa fa-file-text-o"></i></a>';
    return articleUrl;
  }
  //最后面弹出附件列表中用的
  function setAttachlink(value, row, index) {
    attachUrl = '<a href="' + value + '" title="下载" target="_blank"><i class="fa fa-paperclip"></i></a>';
    return attachUrl;
  }
  //最后面弹出pdf列表中用的
  function setPdflink(value, row, index) {
    pdfUrl = '<a href="/pdf?id=' + row.Id + '" title="下载" target="_blank"><i class="fa fa-file-pdf-o"></i></a>';
    return pdfUrl;
  }
  </script>
  <!-- 文章列表 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalarticle">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">文章列表</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <!-- <div id="pdfs" style="display:none"> -->
              <!-- <h3>工程目录分级</h3> -->
              <table id="articles" data-toggle="table" data-page-size="5" data-page-list="[5, 25, 50, All]" data-unique-id="id" data-pagination="true" data-side-pagination="client" data-click-to-select="true">
                <thead>
                  <tr>
                    <th data-width="10" data-checkbox="true"></th>
                    <th data-formatter="index1">#</th>
                    <th data-field="Title">名称</th>
                    <th data-field="Subtext">副标题</th>
                    <th data-field="Link" data-formatter="setArticlecontent">查看</th>
                    <th data-field="Created" data-formatter="localDateFormatter">建立时间</th>
                    <th data-field="Updated" data-formatter="localDateFormatter">修改时间</th>
                  </tr>
                </thead>
              </table>
              <!-- </div> -->
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <!-- 除了**pdf**之外的附件列表 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalattach">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">非PDF附件列表</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <!-- <div id="pdfs" style="display:none"> -->
              <!-- <h3>工程目录分级</h3> -->
              <table id="attachs" data-toggle="table" data-page-size="5" data-page-list="[5, 25, 50, All]" data-unique-id="id" data-pagination="true" data-side-pagination="client" data-click-to-select="true">
                <thead>
                  <tr>
                    <th data-width="10" data-checkbox="true"></th>
                    <th data-formatter="index1">#</th>
                    <th data-field="Title">名称</th>
                    <th data-field="FileSize">大小</th>
                    <th data-field="Link" data-formatter="setAttachlink">下载</th>
                    <th data-field="Created" data-formatter="localDateFormatter">建立时间</th>
                    <th data-field="Updated" data-formatter="localDateFormatter">修改时间</th>
                  </tr>
                </thead>
              </table>
              <!-- </div> -->
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <!-- pdf附件列表 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalpdf">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">pdf附件列表</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <!-- <div id="pdfs" style="display:none"> -->
              <!-- <h3>工程目录分级</h3> -->
              <table id="pdfs" data-toggle="table" data-page-size="5" data-page-list="[5, 25, 50, All]" data-unique-id="id" data-pagination="true" data-side-pagination="client" data-click-to-select="true">
                <thead>
                  <tr>
                    <th data-width="10" data-checkbox="true"></th>
                    <th data-formatter="index1">#</th>
                    <th data-field="Title">名称</th>
                    <th data-field="FileSize">大小</th>
                    <th data-field="Link" data-formatter="setPdflink">下载</th>
                    <th data-field="Created" data-formatter="localDateFormatter">建立时间</th>
                    <th data-field="Updated" data-formatter="localDateFormatter">修改时间</th>
                  </tr>
                </thead>
              </table>
              <!-- </div> -->
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</body>

</html>
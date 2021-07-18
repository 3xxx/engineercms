<!-- 管理员查询所有充值申请 -->
<!DOCTYPE html>

<head>
  <title>用户充值申请</title>
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css" />
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css" />
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
  <script src="/static/js/tableExport.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/webuploader.css">
  <script type="text/javascript" src="/static/js/webuploader.min.js"></script>
  <script type="text/javascript" src="/static/js/jquery-ui.min.js"></script>
  <script type="text/javascript" src="/static/js/clipboard.min.js"></script>
  <link rel="stylesheet" href="/static/froala/css/froala_editor.css">
  <link rel="stylesheet" href="/static/froala/css/froala_style.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/code_view.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/draggable.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/colors.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/emoticons.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/image_manager.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/image.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/line_breaker.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/table.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/char_counter.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/video.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/fullscreen.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/file.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/quick_insert.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/help.css">
  <!-- <link rel="stylesheet" href="/static/froala/css/third_party/spell_checker.css"> -->
  <link rel="stylesheet" href="/static/froala/css/plugins/special_characters.css">
  <link rel="stylesheet" href="/static/froala/js/codemirror.min.css">
  <link rel="stylesheet" href="/static/froala/css/themes/red.css">
  <!-- <link rel="stylesheet" href="/static/css/magnific-popup.css" /> -->
  <!-- <script type="text/javascript" src="/static/js/jquery.magnific-popup.min.js"></script> -->
  <!-- <script src="/static/toast/toast.min.js"></script> -->
  <!-- <link rel="stylesheet" href="/static/toast/toast.min.css"> -->
  <style type="text/css">
  #imgmodalDialog .modal-header {
    cursor: move;
  }

  #modalDialog .modal-header {
    cursor: move;
  }

  #modalDialog1 .modal-header {
    cursor: move;
  }

  #modalDialog2 .modal-header {
    cursor: move;
  }

  #modalDialog3 .modal-header {
    cursor: move;
  }

  #modalDialog4 .modal-header {
    cursor: move;
  }

  #modalDialog5 .modal-header {
    cursor: move;
  }

  #modalDialog6 .modal-header {
    cursor: move;
  }

  #modalDialog7 .modal-header {
    cursor: move;
  }

  #modalDialog8 .modal-header {
    cursor: move;
  }

  #modalDialog9 .modal-header {
    cursor: move;
  }

  #modalDialog10 .modal-header {
    cursor: move;
  }

  #modalDialog11 .modal-header {
    cursor: move;
  }

  /*#modalNewDwg .modal-header {cursor: move;}*/
  /*#modalFlow .modal-header {cursor: move;}*/
  /*body {
    text-align: center;
  }*/
  div#editor {
    width: 81%;
    margin: auto;
    text-align: left;
  }

  .ss {
    background-color: red;
  }

  div#modalTable2 {
    /*.modal .fade .in*/
    z-index: 3;
  }

  h3 .share-icon {
    width: 30px;
    height: 30px
  }
  </style>
</head>

<body>
  <div class="col-lg-12">
    <div id="toolbar1" class="btn-toolbar" role="toolbar" aria-label="...">
      <div class="btn-group">
        <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="true" title="添加资料">
          <i class="fa fa-plus">&nbsp;&nbsp;添加</i>
          <span class="caret"></span>
        </button>
        <ul class="dropdown-menu" aria-labelledby="">
          <li>
            <a href="#" onclick="addButton()"><i class="fa fa-plus">&nbsp;&nbsp;单附件模式</i></a>
          </li>
        </ul>
      </div>
      <div class="btn-group">
        <a href="https://zsj.itdos.net/docs/pss" target="_blank" type="button"><i class="fa fa-question-circle-o">&nbsp;&nbsp;帮助</i></a>
      </div>
      <div class="btn-group">
        <a href="/wiki" target="_blank" type="button"><i class="fa fa-envelope-o">&nbsp;&nbsp;留言</i></a>
      </div>
      <div class="btn-group">
        <a href="/v1/wx/applyrecharge" target="_blank" type="button"><i class="fa fa-money">&nbsp;&nbsp;充值&nbsp;&nbsp;</i></a>
      </div>
    </div>
    <h1 id="amount">用户申请充值记录</h1>
    <table id="table0"></table>
    <div class="modal fade" id="imgModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
      <div class="modal-dialog modal-lg" style="display: inline-block; width: auto;" id="imgmodalDialog">
        <div class="modal-content">
          <div class="modal-header" style="background-color: #1E90FF">
            <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;
            </button>
            <h4 class="modal-title" id="myModalLabel" style="color: #FFFFFF"> 图片预览</h4>
          </div>
          <div class="modal-body text-center">
            <!-- -text-cente  bootstrap子元素居中--->
            <span id="myImg">
              <!--预览图片位置，默认图片-->
              <!-- <img src="./img/notlogin.jpg" class="img-circle"> -->
              <img id="imgInModalID" src="">
            </span>
          </div>
          <div class="modal-footer">
          </div>
        </div>
      </div>
    </div>
    <script type="text/javascript">

    $(function() {
      // 初始化【未接受】工作流表格showToggle:'true',
      $("#table0").bootstrapTable({
        url: '/v1/wx/getapplyrechargedata',
        method: 'get',
        search: 'true',
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
            searchText: $(".search .form-control").val()
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
            // radio: 'true',
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
            field: 'User.Nickname',
            title: '姓名',
            align: "center",
            valign: "middle"
          },
          {
            field: 'Amount',
            title: '数值',
            halign: "center",
            align: "center",
            valign: "middle",
            
          },
          {
            field: 'CreatedAt',
            title: '建立时间',
            formatter: localDateFormatter,
            visible: "false",
            align: "center",
            valign: "middle"
          },
          {
            field: 'UpdatedAt',
            title: '更新时间',
            formatter: localDateFormatter,
            visible: "false",
            align: "center",
            valign: "middle"
          },
          {
            field: 'operate',
            title: '操作',
            formatter: operateFormatter,
            events: operateEvents,
            align: "center",
            valign: "middle"
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

    function operateFormatter(value, row, index) {
      return [
        '<a class="check btn btn-xs btn-success" style="margin-left:10px" href="javascript:void(0)" title="同意">',
        '<i class="fa fa-check-square-o"></i>',
        '</a>',
        '<a class="edit btn btn-xs btn-info" style="margin-left:10px" href="javascript:void(0)" title="编辑">',
        '<i class="fa fa-pencil"></i>',
        '</a>',
        '<a class="delete btn btn-xs btn-danger" style="margin-left:10px" href="javascript:void(0)" title="删除">',
        '<i class="fa fa-trash-o"></i>',
        '</a>'
      ].join('');
    }

    window.operateEvents = {
      'click .check': function(e, value, row, index) {
        $.ajax({
          type: "post", //这里是否一定要用post？？？
          url: '/v1/wx/adduserrecharge/' + row.ID,
          data: { amount: row.Amount },
          // data: JSON.stringify(selectRow3), //JSON.stringify(row),
          success: function(data, status) { //数据提交成功时返回数据
            if (data.info == "SUCCESS") {
              $('#table0').bootstrapTable('refresh', { url: '/v1/wx/getapplyrechargedata'});
            } else {
              alert("更新失败！(status:" + data.data + ".)");
            }
          }
        });
      },
      'click .edit': function(e, value, row, index) {
        alert("编辑功能待完善~");
      },
      'click .delete': function(e, value, row, index) {
        // var url = '/v1/mathcad/deletetemple/'+row.ID
        var mycars = new Array()
        mycars[0] = row;
        var templeid = $.map(mycars, function(row) {//必须构造这个map，否则删除不了前端
          return row.ID;
        })
        if (confirm("确定删除吗？")) {
          return
          //提交到后台进行修改数据库状态修改
          // $.ajax({
          //   type: "post", //这里是否一定要用post？？？
          //   url: '/v1/mathcad/deletetemple/' + row.ID,
          //   // data: JSON.stringify(selectRow3), //JSON.stringify(row),
          //   success: function(data, status) { //数据提交成功时返回数据
          //     if (data.info == "SUCCESS") {
          //       $('#table0').bootstrapTable('remove', {
          //         field: 'ID', //table填充的数据结构中必须提供这个id，否则不能删除某行
          //         values: templeid//必须构造这个map，否则删除不了前端
          //       });
          //       alert("删除成功！(status:" + status + ".)");
          //     } else {
          //       alert("删除失败！(status:" + data.data + ".)");
          //     }
          //   }
          // });
        }
      }
    };

    // 关联成果跳转到对应的树状目录下
    function gototree(e) {
      parent.gototree(e); // pClick 为父页面 js 方法
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

    function setCodetest(value, row, index) {
      //保留，数组和字符串以及循环的处理
      // array=value.split(",")
      // var labelarray = new Array() 
      // for (i=0;i<value.length;i++)//value是数组"Code":[数组"SL0001-510-08","SL0001-510-08"],
      // {
      //   labelarray[i]="<a href='/project/product/attachment/"+value[i]+"'>" + value[i] + "</a>";
      // }
      // if (value.match(",")!=null){
      if (value) {
        array = value.split(",")
        var labelarray = new Array()
        for (i = 0; i < array.length; i++) {
          labelarray[i] = "<a href='/project/product/attachment/" + array[i] + "'>" + array[i] + "</a>";
        }
        return labelarray.join(",");
      }
    }

    function setTitle(value, row, index) {
      return "<a href='/project/product/" + row.Id + "'>" + value + "</a>";
    }

    function setArticle(value, row, index) {
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

    // 编辑成果信息
    // $("#editorProdButton").click(function() {
    function editorProdButton() {
      var selectRow = $('#table0').bootstrapTable('getSelections');
      if (selectRow.length < 1) {
        alert("请先勾选成果！");
        return;
      }
      if (selectRow.length > 1) {
        alert("请不要勾选一个以上成果！");
        return;
      }
      if (selectRow[0].Uid === {{.Uid }} || {{.RoleUpdate }} == "true") {
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

        $("input#cid").remove();
        var th1 = "<input id='cid' type='hidden' name='cid' value='" + selectRow[0].Id + "'/>"
        $(".modal-body").append(th1); //这里是否要换名字$("p").remove();
        $("#prodcode3").val(selectRow[0].Code);
        $("#prodname3").val(selectRow[0].Title);
        $("#prodlabel3").val(selectRow[0].Label);
        $("#prodprincipal3").val(selectRow[0].Principal);
        var value = selectRow[0].Relevancy
        var relevancy3
        if (value) {
          if (value.length == 1) { //'<a href="/project/product/article/'
            var array = value[0].Relevancy.split(",")
            var relevarray = new Array()
            for (i = 0; i < array.length; i++) {
              relevarray[i] = array[i];
            }
            relevancy3 = relevarray.join(",");
          } else if (value.length > 1) {
            var relevarray = new Array()
            for (i = 0; i < value.length; i++) {
              relevarray[i] = value[i].Relevancy;
            }
            relevancy3 = relevarray.join(",");
          }
        }
        $("#relevancy3").val(relevancy3);
        $('#modalProdEditor').modal({
          show: true,
          backdrop: 'static'
        });
      } else {
        alert("权限不够！" + selectRow[0].Uid);
        return;
      }
    }

    // 删除成果
    $("#deleteButton").click(function() {
      var selectRow = $('#table0').bootstrapTable('getSelections');
      if (selectRow.length <= 0) {
        alert("请先勾选成果！");
        return false;
      }
      //问题：如果多选，而其中有自己的，也有自己不具备权限的********
      if ({{.Uid }} == 0) {
        alert("权限不足！");
        return;
      }
      if (selectRow[0].Uid === {{.Uid }} || {{.RoleDelete }} == "true") {
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
        if (confirm("确定删除成果吗？一旦删除将无法恢复！")) {
          // var title = $.map(selectRow, function(row) {
          //   return row.Title;
          // })
          var ids = "";
          for (var i = 0; i < selectRow.length; i++) {
            if (i == 0) {
              ids = selectRow[i].Id;
            } else {
              ids = ids + "," + selectRow[i].Id;
            }
          }
          //删除前端表格用的
          var ids2 = $.map($('#table0').bootstrapTable('getSelections'), function(row) {
            return row.Id;
          });

          $.ajax({
            type: "post",
            url: "/project/product/deleteproduct",
            data: { ids: ids },
            success: function(data, status) {
              alert("删除“" + data + "”成功！(status:" + status + ".)");
              //删除已选数据
              $('#table0').bootstrapTable('remove', {
                field: 'Id',
                values: ids2
              });
            }
          });
        }
      } else {
        alert("权限不够！" + selectRow[0].Uid);
        return;
      }
    })

  </script>

    <!-- 编辑成果名称等信息 -->
    <div class="form-horizontal">
      <div class="modal fade" id="modalProdEditor">
        <div class="modal-dialog" id="modalDialog6">
          <div class="modal-content">
            <div class="modal-header" style="background-color: #8bc34a">
              <button type="button" class="close" data-dismiss="modal">
                <span aria-hidden="true">&times;</span>
              </button>
              <h3 class="modal-title">编辑成果信息</h3>
            </div>
            <div class="modal-body">
              <div class="modal-body-content">
                <div class="form-group must">
                  <label class="col-sm-3 control-label">编号</label>
                  <div class="col-sm-7">
                    <input type="text" class="form-control" id="prodcode3" name="prodcode3"></div>
                </div>
                <div class="form-group must">
                  <label class="col-sm-3 control-label">标题</label>
                  <div class="col-sm-7">
                    <input type="tel" class="form-control" id="prodname3" name="prodname3"></div>
                </div>
                <div class="form-group must">
                  <label class="col-sm-3 control-label">关键字</label>
                  <div class="col-sm-7">
                    <input type="tel" class="form-control" id="prodlabel3" name="prodlabel3" placeholder="以英文,号分割"></div>
                </div>
                <div class="form-group must">
                  <label class="col-sm-3 control-label">设计</label>
                  <div class="col-sm-7">
                    <input type="tel" class="form-control" id="prodprincipal3" name="prodprincipal3"></div>
                </div>
                <div class="form-group must">
                  <label class="col-sm-3 control-label">关联文件</label>
                  <div class="col-sm-7">
                    <input type="tel" class="form-control" id="relevancy3" name="relevancy3" placeholder="输入文件编号，以英文,号分割">
                  </div>
                </div>
              </div>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
              <button type="button" class="btn btn-primary" onclick="updateprod()">保存</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <!-- <script type="text/javascript" src="/static/froala/js/jquery.min.1.11.0.js"></script> -->
  <script type="text/javascript" src="/static/froala/js/codemirror.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/xml.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/froala_editor.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/align.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/char_counter.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/code_beautifier.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/code_view.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/colors.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/draggable.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/emoticons.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/entities.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/file.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/font_size.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/font_family.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/fullscreen.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/image.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/image_manager.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/line_breaker.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/inline_style.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/link.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/lists.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/paragraph_format.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/paragraph_style.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/quick_insert.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/quote.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/table.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/save.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/url.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/video.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/help.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/print.min.js"></script>
  <!-- <script type="text/javascript" src="/static/froala/js/third_party/spell_checker.min.js"></script> -->
  <script type="text/javascript" src="/static/froala/js/plugins/special_characters.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/word_paste.min.js"></script>
  <script src="/static/froala/js/languages/zh_cn.js"></script>
  <script>
  // $(function(){
  //   $('#edit').froalaEditor()
  // });
  $(function() {
    //超大屏幕'fullscreen',
    var toolbarButtons = ['bold', 'italic', 'underline', 'strikeThrough', 'subscript', 'superscript', 'fontFamily', 'fontSize', '|', 'color', 'emoticons', 'inlineStyle', 'paragraphStyle', '|', 'paragraphFormat', 'align', 'formatOL', 'formatUL', 'outdent', 'indent', 'quote', 'insertHR', '-', 'insertLink', 'insertImage', 'insertVideo', 'insertFile', 'insertTable', 'undo', 'redo', 'clearFormatting', 'selectAll', 'html', 'help'];
    //大屏幕
    var toolbarButtonsMD = ['bold', 'italic', 'underline', 'strikeThrough', 'subscript', 'superscript', '|', 'fontFamily', 'fontSize', 'color', 'inlineStyle', 'paragraphStyle', '|', 'paragraphFormat', 'align', 'formatOL', 'formatUL', 'outdent', 'indent', '-', 'quote', 'insertLink', 'insertImage', 'insertVideo', 'insertFile', 'insertTable', '|', 'specialCharacters', 'insertHR', 'undo', 'redo', 'clearFormatting', '|', 'html', 'help'];
    //小屏幕'fullscreen',
    var toolbarButtonsSM = ['bold', 'italic', 'underline', 'fontFamily', 'fontSize', 'insertLink', 'insertImage', 'insertVideo', 'insertTable', 'undo', 'redo'];
    //手机
    var toolbarButtonsXS = ['insertImage', 'insertVideo', 'bold', 'italic', 'fontSize', 'undo', 'redo'];
    var pid = $('#pid').val();
    //编辑器初始化并赋值 
    $('#edit').froalaEditor({
      placeholderText: '请输入内容',
      charCounterCount: true, //默认
      // charCounterMax         : -1,//默认
      saveInterval: 0, //不自动保存，默认10000
      // theme                    : "red",
      height: "300px",
      toolbarBottom: false, //默认
      toolbarButtonsMD: toolbarButtons, //toolbarButtonsMD,
      toolbarButtonsSM: toolbarButtonsMD, //toolbarButtonsSM,
      toolbarButtonsXS: toolbarButtonsXS,
      toolbarInline: false, //true选中设置样式,默认false
      imageUploadMethod: 'POST',
      heightMin: 450,
      charCounterMax: 3000,
      // imageUploadURL: "uploadImgEditor",
      imageParams: { postId: "123" },
      params: {
        acl: '01',
        AWSAccessKeyId: '02',
        policy: '03',
        signature: '04',
      },
      autosave: true,
      autosaveInterval: 2500,
      saveURL: 'hander/FroalaHandler.ashx',
      saveParams: { postId: '1' },
      spellcheck: false,
      imageUploadURL: '/uploadimg', //上传到本地服务器
      imageUploadParams: { pid: '{{.Id}}' },
      imageDeleteURL: 'lib/delete_image.php', //删除图片
      imagesLoadURL: 'lib/load_images.php', //管理图片
      videoUploadURL: '/uploadvideo',
      videoUploadParams: { pid: '{{.Id}}' },
      fileUploadURL: '/uploadimg',
      fileUploadParams: { pid: '{{.Id}}' },
      enter: $.FroalaEditor.ENTER_BR,
      language: 'zh_cn',
      // toolbarButtons: ['bold', 'italic', 'underline', 'paragraphFormat', 'align','color','fontSize','insertImage','insertTable','undo', 'redo']
    });
  })

  // 编辑成果信息
  function updateprod() {
    // var radio =$("input[type='radio']:checked").val();
    var projectid = $('#cid').val();
    var prodcode = $('#prodcode3').val();
    var prodname = $('#prodname3').val();
    var prodlabel = $('#prodlabel3').val();
    var prodprincipal = $('#prodprincipal3').val();
    var relevancy = $('#relevancy3').val();
    if (prodname && prodcode) {
      $.ajax({
        type: "post",
        url: "/project/product/updateproduct",
        data: { pid: projectid, code: prodcode, title: prodname, label: prodlabel, principal: prodprincipal, relevancy: relevancy }, //父级id
        success: function(data, status) {
          alert("添加“" + data + "”成功！(status:" + status + ".)");
          $('#modalProdEditor').modal('hide');
          $('#table0').bootstrapTable('refresh', { url: '/project/products/' + {{.Id }} });
        },
      });
    } else {
      alert("请填写编号和名称！");
      return;
    }
  }

  //******表格追加项目同步ip中的数据*******
  $(function() {
    $('#synchIP').click(function() {
      // alert("ha ");
      $.ajax({
        type: "get",
        url: "/project/synchproducts/" + {{.Id }},
        // data: {ids:ids},
        success: function(data, status) {
          alert("同步成功！(status:" + status + ".)");
          //追加数据
          $('#table0').bootstrapTable('append', data);
          $('#table0').bootstrapTable('scrollTo', 'bottom');
        }
      });
    });
  });

  function randomData() {
    var startId = ~~(Math.random() * 100),
      rows = [];
    for (var i = 0; i < 10; i++) {
      rows.push({
        Id: startId + i,
        Code: startId + i,
        Title: 'test' + (startId + i),
        // id: startId + i,
        // name: 'test' + (startId + i),
        // price: '$' + (startId + i)
      });
    }
    return rows;
  }

  //勾选后输入框可用
  function station_select1() {
    if (box1.checked) {
      document.getElementById("relevancy1").disabled = false;
    } else {
      document.getElementById("relevancy1").disabled = true;
    }
  }

  function station_select2() {
    if (box2.checked) {
      document.getElementById("relevancy2").disabled = false;
    } else {
      document.getElementById("relevancy2").disabled = true;
    }
  }

  $(document).ready(function() {
    $("#imgmodalDialog").draggable({ handle: ".modal-header" });
    $("#modalDialog").draggable({ handle: ".modal-header" }); //为模态对话框添加拖拽
    $("#modalDialog1").draggable({ handle: ".modal-header" });
    $("#modalDialog2").draggable({ handle: ".modal-header" });
    $("#modalDialog3").draggable({ handle: ".modal-header" });
    $("#modalDialog4").draggable({ handle: ".modal-header" });
    $("#modalDialog5").draggable({ handle: ".modal-header" });
    $("#modalDialog6").draggable({ handle: ".modal-header" });
    $("#modalDialog7").draggable({ handle: ".modal-header" });
    $("#modalDialog8").draggable({ handle: ".modal-header" });
    $("#modalDialog9").draggable({ handle: ".modal-header" });
    $("#modalDialog10").draggable({ handle: ".modal-header" });
    $("#modalDialog11").draggable({ handle: ".modal-header" });
    // $("#modalNewDwg").draggable({ handle: ".modal-header" });
    // $("#modalFlow").draggable({ handle: ".modal-header" });
    $("#myModal").css("overflow", "hidden"); //禁止模态对话框的半透明背景滚动
  })
  </script>
</body>

</html>
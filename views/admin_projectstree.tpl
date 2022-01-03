<!-- 采用bootstrap treeview编辑目录、设置同步ip、设置公开私有 -->
<!DOCTYPE html>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css" />
<script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-treeview.css" />
<script src="/static/js/bootstrap-treeview.js"></script>
<script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css" />
<!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css"/> -->
<script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
<!-- <script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script> -->
<!-- <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script> -->
<script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
<link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
<!-- <link rel="stylesheet" type="text/css" href="/static/css/font-awesome.min.css"/> -->
<script src="/static/js/tableExport.js"></script>
<script type="text/javascript" src="/static/js/moment.min.js"></script>
</head>

<body>
  <div class="col-lg-12">
    <h3>项目列表</h3>
    <div id="toolbar1" class="btn-group">
      <!-- <button type="button" data-name="editorcateButton" id="editorcateButton" class="btn btn-default"> <i class="fa fa-edit">编辑目录</i>
        </button> -->
      <button type="button" data-name="editorpubButton" id="editorpubButton" class="btn btn-default"> <i class="fa fa-edit">公开、私有</i>
      </button>
    </div>
    <table id="table0"></table>
    <script type="text/javascript">
    function index1(value, row, index) {
      // alert( "Data Loaded: " + index );
      return index + 1
    }

    function localDateFormatter(value) {
      return moment(value, 'YYYY-MM-DD').format('YYYY-MM-DD');
    }

    $(function() {
      // 初始化【未接受】工作流表格
      $("#table0").bootstrapTable({
        url: '/project/getprojects',
        method: 'get',
        search: 'true',
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
        pageSize: 5,
        pageNumber: 1,
        pageList: [15, 20, 50, 100],
        singleSelect: "true",
        clickToSelect: "true",
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
            radio: 'true',
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
            align: "center",
            valign: "middle"
          },
          {
            field: 'Title',
            title: '名称',
            // formatter:setTitle,
            align: "center",
            valign: "middle"
          },
          // {
          //   field: 'Label',
          //   title: '标签',
          //   formatter:setLable,
          //   align:"center",
          //   valign:"middle"
          // },
          {
            field: 'Principal',
            title: '负责人',
            align: "center",
            valign: "middle"
          },
          // {
          //   field: 'Number',
          //   title: '成果数量',
          //   formatter:setCode,
          //   align:"center",
          //   valign:"middle"
          // },
          // {
          //   field: 'action',
          //   title: '时间轴',
          //   formatter:actionFormatter,
          //   events:actionEvents,
          //   align:"center",
          //   valign:"middle"
          // },
          {
            field: 'Created',
            title: '建立时间',
            formatter: localDateFormatter,
            align: "center",
            valign: "middle"
          },
          // {
          //     field: 'dContMainEntity.createTime',
          //     title: '发起时间',
          //     formatter: function (value, row, index) {
          //         return new Date(value).toLocaleString().substring(0,9);
          //     }
          // },
          {
            field: 'action',
            title: '操作',
            formatter: 'actionFormatter',
            events: 'actionEvents',
          }
        ]
      });
    });
    // 点击出现目录
    $(function() {
      $("#table0").on("click-row.bs.table", function(e, row, ele) {
        $(".info").removeClass("info");
        $(ele).addClass("info");
        rowid = row.Id; //全局变量
        rowtitle = row.Title
        $("#rowtitle").html("项目目录-" + rowtitle);

        // $('#iptable').bootstrapTable('refresh', {url:'/admin/project/synchip/'+row.Id});
        $("input#cid").remove();
        var th1 = "<input id='cid' type='hidden' name='cid' value='" + rowid + "'/>"
        $(".modal-body").append(th1);
        //初始化树   
        $.ajax({ //JQuery的Ajax  
          type: 'POST',
          dataType: "json", //返回数据类型  
          // async:false, //同步会出现警告：Synchronous XMLHttpRequest on the main thread is deprecated because of its detrimental effects to the end user's experience 
          url: "/admin/project/getprojectcate/" + rowid, //请求的action路径  
          // data: {"flag":true},  //同步请求将锁住浏览器，用户其它操作必须等待请求完成才可以执行  
          error: function() { //请求失败处理函数    
            alert('请求失败');
          },
          success: function(data) { //请求成功后处理函数。取到Json对象data
            // alert(data);
            $('#tree').treeview({
              showCheckbox: true,
              data: [data], // data is not optional
              levels: 5,
              enableLinks: true,
              // hierarchicalCheck:true,有效！！
              // propagateCheckEvent:true,
              state: {
                checked: true,
                disabled: true,
                expanded: true,
                selected: true
              }
            });
            $('#tree').treeview('expandAll'); //这句为何无效？
            $("#details").show();
          }
        });
      });
    });

    function actionFormatter(value, row, index) {
      return '<button type="button" data-name="editorjsonButton" id="editorjsonButton" class="btn btn-info btn-xs"> <i class="fa fa-user">设置json数据</i></button>';
    }
    window.actionEvents = {
      'click #editorjsonButton': function(e, value, row, index) {
        // alert(row.Id);
        window.open("/v1/admin/jsoneditor?projectid="+row.Id,"_blank");
      },
    };

    $(document).ready(function() {
      // $("#editorcateButton").click(function() {
      //   var selectRow=$('#table0').bootstrapTable('getSelections');
      //   if (selectRow.length<1){
      //     alert("请先勾选！");
      //     return;
      //   }
      //   if (selectRow.length>1){
      //   alert("请不要勾选一个以上！");
      //   return;
      //   }
      //   $("input#cid").remove();
      //   var th1="<input id='cid' type='hidden' name='cid' value='" +selectRow[0].Id+"'/>"
      //   $(".modal-body").append(th1);//这里是否要换名字$("p").remove();
      //   //初始化树   
      //     $.ajax({  //JQuery的Ajax  
      //         type: 'POST',    
      //         dataType : "json",//返回数据类型  
      //         async:false,  
      //         url: "/admin/project/getprojectcate/"+selectRow[0].Id,//请求的action路径  
      //         // data: {"flag":true},  //同步请求将锁住浏览器，用户其它操作必须等待请求完成才可以执行  
      //         error: function () {//请求失败处理函数    
      //             alert('请求失败');    
      //         },  
      //         success:function(data){ //请求成功后处理函数。取到Json对象data
      //           // alert(data);
      //           $('#tree').treeview({
      //             showCheckbox: true,
      //             data: [data], // data is not optional
      //             // levels: 2,
      //             enableLinks: true,
      //             // hierarchicalCheck:false,
      //             // propagateCheckEvent:false,
      //             state: {
      //               checked: true,
      //               disabled: true,
      //               expanded: true,
      //               selected: true
      //             }
      //           });
      //           $('#tree').treeview('expandAll'); 
      //         }
      //     });
      //     $('#modalTable').modal({
      //     show:true,
      //     backdrop:'static'
      //     });
      //   })

      //模态框中添加节点
      $("#addcateButton").click(function() {
        //得到选择的节点
        var arr = new Array();
        arr = $('#tree').treeview('getChecked');
        if (arr.length == 0) {
          alert("请先勾选！");
          return;
        }

        if (arr.length >= 2) {
          alert("请不要勾选一个以上！");
          return;
        }
        // var nodeData = $tree.treeview('getSelected');
        //模态框填入节点名称和代码
        // $("input#pid").remove();
        // $("input#nid").remove();
        // var th1="<input id='pid' type='hidden' name='pid' value='" +arr[0].id+"'/>"
        // var th2="<input id='nid' type='hidden' name='nid' value='" +arr[0].nodeId+"'/>"
        // $(".modal-body").append(th1);//这里是否要换名字$("p").remove();
        // $(".modal-body").append(th2);
        $('#modalTable2').modal({
          show: true,
          backdrop: 'static'
        });
      })

      //模态框中编辑节点
      $("#updateButton").click(function() {
        //得到选择的节点
        var arr = new Array();
        arr = $('#tree').treeview('getChecked');
        // arr=$('#tree').treeview('getSelected',0);
        // alert(arr[0].nodeId);//节点顺序号0.0.0.1这样的
        // alert(arr[0].code);
        if (arr.length == 0) {
          alert("请先勾选！");
          return;
        }

        if (arr.length >= 2) {
          alert("请不要勾选一个以上！");
          return;
        }
        $("#projcatename3").val(arr[0].text);
        $("#projcatecode3").val(arr[0].code);
        // var nodeData = $tree.treeview('getSelected');
        //模态框填入节点名称和代码
        // $("input#pid").remove();
        // $("input#nid").remove();
        // var th1="<input id='pid' type='hidden' name='pid' value='" +arr[0].id+"'/>"
        // var th2="<input id='nid' type='hidden' name='nid' value='" +arr[0].nodeId+"'/>"
        // $(".modal-body").append(th1);//这里是否要换名字$("p").remove();
        // $(".modal-body").append(th2);
        $('#modalTable3').modal({
          show: true,
          backdrop: 'static'
        });
      })
    })
    //保存添加节点
    function savecate() {
      //获取数据，添加到树中
      var projcatename2 = $('#projcatename2').val();
      var projcatecode2 = $('#projcatecode2').val();
      arr = $('#tree').treeview('getChecked');
      // var nid = $('#nid').val();//节点顺序号nodeId
      // alert(projcatename2);
      // alert(nid);
      // $.getJSON("treeAddData.json", function (data) {
      // }); 
      if (projcatename2) {
        $.ajax({
          type: "post",
          url: "/admin/project/addprojectcate",
          data: { id: arr[0].id, name: projcatename2, code: projcatecode2 }, //父级id
          success: function(data, status) {
            alert("添加“" + data + "”成功！(status:" + status + ".)");
            var singleNode = {
              text: projcatename2,
              id: data,
              code: projcatecode2
            };
            $("#tree").treeview("addNode", [singleNode, arr]);
            $('#modalTable2').modal('hide');
          }
        });
      }
    }
    //编辑节点
    function updatecate() {
      var projcatename3 = $('#projcatename3').val();
      var projcatecode3 = $('#projcatecode3').val();
      arr = $('#tree').treeview('getChecked');
      if (projcatename3) {
        $.ajax({
          type: "post",
          url: "/admin/project/updateprojectcate",
          data: { id: arr[0].id, name: projcatename3, code: projcatecode3 },
          success: function(data, status) {
            alert("编辑成功！(status:" + status + ".)");
            var singleNode = {
              text: projcatename3,
              id: data,
              code: projcatecode3
            };
            $('#tree').treeview('updateNode', [arr, singleNode, { silent: true }]);
            // $('#tree').treeview('updateNode', [ node, newNode, { silent: true } ]);
            $('#modalTable3').modal('hide');
          }
        });
      }
    }
    </script>
    <!-- 编辑项目目录 -->
    <!-- <div class="form-horizontal">
    <div class="modal fade" id="modalTable">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">编辑项目目录</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div  class="col-lg-3" class="btn-group">
                <button type="button" data-name="addcateButton" id="addcateButton" class="btn btn-default"> <i class="fa fa-edit">添加下级</i>
                </button>
                <button type="button" data-name="updateButton" id="updateButton" class="btn btn-default"> <i class="fa fa-edit">修改名称</i>
                </button>
                <button type="button" data-name="removeButton" id="removeButton" class="btn btn-default">
                <i class="fa fa-edit">删除</i>
                </button>
              </div>
              <div class="col-lg-3">
                <div id="tree"></div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div> -->
    <!-- data-query-params="queryParams" data-content-type="application/json"-->
    <!-- 编辑项目目录 -->
    <div id="details" style="display:none">
      <h3 id="rowtitle"></h3>
      <div class="btn-group">
        <button type="button" data-name="addcateButton" id="addcateButton" class="btn btn-default" data-target="modal"><i class="fa fa-plus" aria-hidden="true"> </i>添加</button>
        <button type="button" data-name="updateButton" id="updateButton" class="btn btn-default" data-target="modal"><i class="fa fa-edit" aria-hidden="true"> </i>编辑</button>
        <button type="button" data-name="removeButton" id="removeButton" class="btn btn-default" data-target="default"><i class="fa fa-trash" aria-hidden="true"> </i>删除</button>
      </div>
      <div class="modal-body">
        <!-- <div class="modal-body-content"> -->
        <!-- <div  class="col-lg-3" class="btn-group">
                <button type="button" data-name="addcateButton" id="addcateButton" class="btn btn-default"> <i class="fa fa-edit">添加下级</i>
                </button>
                <button type="button" data-name="updateButton" id="updateButton" class="btn btn-default"> <i class="fa fa-edit">修改名称</i>
                </button>
                <button type="button" data-name="removeButton" id="removeButton" class="btn btn-default">
                <i class="fa fa-edit">删除</i>
                </button>
              </div> -->
        <!-- <div class="col-lg-3"> -->
        <div id="tree"></div>
        <!-- </div> -->
        <!-- </div> -->
      </div>
      <!-- </div> -->
    </div>
    <!-- 添加项目目录 -->
    <div class="container form-horizontal">
      <div class="modal fade" id="modalTable2">
        <div class="modal-dialog">
          <div class="modal-content">
            <div class="modal-header">
              <button type="button" class="close" data-dismiss="modal">
                <span aria-hidden="true">&times;</span>
              </button>
              <h3 class="modal-title">添加目录</h3>
            </div>
            <div class="modal-body">
              <div class="modal-body-content">
                <div class="form-group must">
                  <label class="col-sm-3 control-label">目录名称</label>
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
              <!-- <button type="submit" class="btn btn-primary">保存</button> -->
              <button type="button" class="btn btn-primary" onclick="savecate()">保存</button>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- 编辑项目目录 -->
    <div class="container form-horizontal">
      <div class="modal fade" id="modalTable3">
        <div class="modal-dialog">
          <div class="modal-content">
            <div class="modal-header">
              <button type="button" class="close" data-dismiss="modal">
                <span aria-hidden="true">&times;</span>
              </button>
              <h3 class="modal-title">编辑目录</h3>
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
                    <input type="tel" class="form-control" id="projcatecode3"></div>
                </div>
              </div>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
              <!-- <button type="submit" class="btn btn-primary">保存</button> -->
              <button type="button" class="btn btn-primary" onclick="updatecate()">保存</button>
            </div>
          </div>
        </div>
      </div>
    </div>
    <script type="text/javascript">
    //删除节点
    $("#removeButton").click(function() {
      arr = $('#tree').treeview('getChecked');
      if (arr.length == 0) {
        alert("请先勾选！");
        return;
      }
      var ids = "";
      for (var i = 0; i < arr.length; i++) {
        if (i == 0) {
          ids = arr[i].id;
        } else {
          ids = ids + "," + arr[i].id;
        }
      }
      if (confirm("确定删除吗？第一次提示！")) {} else {
        return false;
      }
      if (confirm("确定删除吗？一旦删除将无法恢复！")) {
        $.ajax({
          type: "post",
          url: "/admin/project/deleteprojectcate",
          data: { ids: ids },
          success: function(data, status) {
            if (data == "no") {
              alert("删除出错！");
            } else {
              alert("删除“" + data + "”成功！(status:" + status + ".)");
              $('#tree').treeview('removeNode', [arr, { silent: true }]);
              // $('#tree').treeview('revealNode', [ nodes, { silent: true } ]);
            }
          }
        });
      }
    })
    </script>
  </div>
</body>

</html>
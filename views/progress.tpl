<!-- 节点进度展示设计测试 -->
<!DOCTYPE html>
<html>
  {{template "header"}}
    <meta charset="UTF-8">
    <title>待处理成果</title>

    <!-- <script type="text/javascript" src="/static/js/moment.min.js"></script> -->
    <!-- <script type="text/javascript" src="/static/js/daterangepicker.js"></script> -->
    <!-- <link rel="stylesheet" type="text/css" href="/static/css/daterangepicker.css"/> -->
    <!-- <script type="text/javascript" src="/static/bootstrap-datepicker/bootstrap-datepicker.js"></script> -->
    <!-- <script type="text/javascript" src="/static/bootstrap-datepicker/bootstrap-datepicker.zh-CN.js"></script> -->
    <!-- <link rel="stylesheet" type="text/css" href="/static/bootstrap-datepicker/bootstrap-datepicker3.css"/> -->

    <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
    <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css"/>

    <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
    <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
    <script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script>
    <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script>
    <!-- <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script> -->

    <link rel="stylesheet" type="text/css" href="/static/css/select2.min.css"/>
    <script type="text/javascript" src="/static/js/select2.js"></script>
    <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
    <!-- <script src="/static/js/tableExport.js"></script> -->
    <!-- <script src="/static/js/jquery.form.js"></script> -->
<!-- <script type="text/javascript" src="/static/js/jquery.mockjax.js"></script> -->

    <link rel="stylesheet" type="text/css" href="/static/css/zb-main-ce9797224b9c240256d3159bf553be0cc4fc84bc.css"/>

    <link type="text/css" rel="stylesheet" href="/static/css/zb-svg-440758a60a537e595528613553c5d25aa114cd30.css">
    <style>
      /*.form-group .datepicker{
        z-index: 9999;
      }*/
      i#delete
        {
          color:#C71585;
        }
    </style>
  </head>
<!-- <div class="navbar navba-default navbar-fixed-top"> -->
      <div class="container-fill">{{template "navbar" .}}</div>
    <!-- </div> -->
<body>
    
  <div class="col-lg-7">

    <h3>我发起，待提交</h3>
    <div id="toolbar" class="btn-group">
        <button type="button" data-name="addButton" id="addButton" class="btn btn-default"> <i class="fa fa-plus">添加</i>
        </button>
        <button type="button" data-name="editorButton" id="editorButton" class="btn btn-default"> <i class="fa fa-edit">编辑</i>
        </button>
        <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
        </button>
    </div>
    <table id="table"
      data-query-params="queryParams"
      data-url="/getprogress"
      data-toolbar="#toolbar"
      data-search="true"
      data-show-refresh="true"
      data-show-toggle="true"
      data-show-columns="true"
      data-striped="true"
      data-clickToSelect="true"
      data-show-export="true"
      data-filter-control="true"
      >
      <thead>
    <tr>
      <th data-align="center" data-valign="middle">选择</th>
      <th data-field="id" data-width="1" data-align="center" data-valign="middle">序号</th>
      <!-- <th data-field="name" data-editable="true" data-editable-emptytext="For free." data-editable-display-formatter="nameFormatter">Price</th> -->
      <th data-field="tasktitle" data-editable-display-formatter="taskFormatter" data-width="150" data-valign="middle">任务名称</th>
      <th data-field="name1" data-editable="true" data-editable-emptytext="For free." data-editable-display-formatter="priceFormatter" data-width="10" data-align="center" data-align="center" data-valign="middle" data-title-tooltip="content">设计</th>
      <th data-field="path1" data-editable="true" data-editable-emptytext="For free." data-editable-display-formatter="priceFormatter1" data-width="80" data-align="center" data-valign="middle">流程</th>
      <th data-field="name2" data-editable="true" data-editable-emptytext="For free." data-editable-display-formatter="priceFormatter" data-width="10" data-align="center" data-valign="middle">校核</th>
      <th data-field="path2" data-editable="true" data-editable-emptytext="For free." data-editable-display-formatter="priceFormatter1" data-width="80" data-align="center" data-valign="middle">流程</th>
      <th data-field="name3" data-editable="true" data-editable-emptytext="For free." data-editable-display-formatter="priceFormatter" data-width="10" data-align="center" data-valign="middle">审查</th>
    </tr>
  </thead>
    </table>

<div class="reward-progress">
    <!-- <div class="z-nav-container"> -->
        <div class="step-contain box">
            <div class="step box-aw"> <i class="ic-svg ic-home-box-accept"></i>
                <div>合同建立</div>
                <div class="text-gray">2017-05-22</div>
            </div>

            <div class="step box-aw"> <i class="ic-svg ic-home-box-accept"></i>
                <div>合同生效</div>
                <div class="text-gray">2017-05-22</div>
            </div>

            <div class="step box-aw">
                <i class="ic-svg ic-home-box-accept"></i>
                <div>预付款请款</div>
                <div class="text-gray">2017-05-22</div>
            </div>

            <div class="step box-aw">
                <i class="ic-svg ic-home-box-accept"></i>
                <div>预付款到帐</div>
                <div class="text-gray">2017-05-22</div>
            </div>

            <div class="step box-aw">
                <i class="step-item  step-on "></i>
                <div>第二期请款</div>
                <div class="text-gray">2017-06-18</div>
            </div>

            <div class="step box-aw">
                <i class="step-item "></i>
                <div>第二期到帐</div>
                <div class="text-gray"></div>
            </div>

            <div class="step box-aw">
                <i class="step-item "></i>
                <div>第三期请款</div>
                <div class="text-gray"></div>
            </div>

            <div class="step box-aw">
                <i class="ic-svg ic-home-box-accept"></i>
                <div>第三期到帐</div>
                <div class="text-gray">2017-05-22</div>
            </div>

            <div class="step box-aw">
                <i class="step-item  step-on "></i>
                <div>合同结束</div>
                <div class="text-gray">2017-06-18</div>
            </div>
        </div>
    <!-- </div> -->
</div>

    <!-- 添加成果 -->
    <div class="container">
      <div class="form-horizontal">
        <div class="modal fade" id="modalTable">
          <div class="modal-dialog" style="width: 800px">
            <div class="modal-content">
              <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal">
                  <span aria-hidden="true">&times;</span>
                </button>
                <h3 class="modal-title">添加成果</h3>
              </div>
              <div class="modal-body">
                <div class="modal-body-content">
                  <div class="form-group must">
                    <label class="col-sm-3 control-label">项目编号/名称</label>
                    <div class="col-sm-2">
                      <input type='text' placeholder='项目编号' class="form-control" id='Pnumber' value='' size='4'/>
                    </div> 
                  <!-- </div> -->
                  <!-- <div class="form-group must"> -->
                    <!-- <label class="col-sm-3 control-label">项目名称</label> -->
                    <div class="col-sm-6">
                      <input type='text' placeholder='项目名称' class="form-control" id='Pname' value='' size='20'/>
                    </div>    
                  </div>
                  <div class="form-group must">
                    <label class="col-sm-3 control-label">项目阶段/专业</label>
                    <div class="col-sm-4">
                      <select class="form-control" id='Stage'>
                        <option>阶段：</option>
                        <option>规划</option>
                        <option>项目建议书</option>
                        <option>可行性研究</option>
                        <option>初步设计</option>
                        <option>招标设计</option>
                        <option>施工图</option>
                      </select>
                    </div>    
                  <!-- </div> -->
                  <!-- <div class="form-group must"> -->
                    <!-- <label class="col-sm-3 control-label">项目专业</label> -->
                    <div class="col-sm-4">
                      <select class="form-control" id='Section'>
                        <option>专业：</option>
                        <option>水工</option>
                        <option>施工</option>
                        <option>预算</option>
                      </select>
                    </div>    
                  </div>
                  <div class="form-group must">
                    <label class="col-sm-3 control-label">成果编号/名称</label>
                    <div class="col-sm-3">
                      <input type='text' placeholder='成果编号' class="form-control" id='Tnumber' value='' size='10'/>
                    </div>    
                  <!-- </div> -->
                  <!-- <div class="form-group must"> -->
                    <!-- <label class="col-sm-3 control-label">成果名称</label> -->
                    <div class="col-sm-5">
                      <input type='text' placeholder='成果名称' class="form-control" id='Name' value='' size='25'/>
                    </div>    
                  </div>
                  <div class="form-group must">
                    <label class="col-sm-3 control-label">成果类型/数量</label>
                    <div class="col-sm-4">
                      <select class="form-control" id='Category'>
                        <option>成果类型：</option>
                      </select>
                    </div>    
                  <!-- </div> -->
                  <!-- <div class="form-group must"> -->
                    <!-- <label class="col-sm-3 control-label">数量</label> -->
                    <div class="col-sm-4">
                      <input type='text' placeholder='数量' class="form-control" id='Count' value='' size='2'/>
                    </div>    
                  </div>
                  <div class="form-group must">
                    <label class="col-sm-3 control-label">人员名称(拼音)</label>
                    <div class="col-sm-2">
                      <input type='text' placeholder='绘制/编制' class="form-control" id="uname1" value='' list="cars1" size='7'/>
                    </div>    
                  <!-- </div> -->
                  <!-- <div class="form-group must"> -->
                    <!-- <label class="col-sm-3 control-label">成果类型</label> -->
                    <div class="col-sm-2">
                      <input type='text' placeholder='设计' class="form-control" id="uname2" value='' list="cars2" size='7'/>
                    </div>    
                  <!-- </div> -->
                  <!-- <div class="form-group must"> -->
                    <!-- <label class="col-sm-3 control-label">成果类型</label> -->
                    <div class="col-sm-2">
                      <input type='text' placeholder='校核' class="form-control" id="uname3" value='' list="cars3" size='7'/>
                    </div>    
                  <!-- </div> -->
                  <!-- <div class="form-group must"> -->
                    <!-- <label class="col-sm-3 control-label">成果类型</label> -->
                    <div class="col-sm-2">
                      <input type='text' placeholder='审查' class="form-control" id="uname4" value='' list="cars4" size='7'/>
                    </div>    
                  </div>
                  <div class="form-group must">
                    <label class="col-sm-3 control-label">绘制/设计系数</label>
                    <div class="col-sm-4">
                      <input type='text' placeholder='绘制/编制系数' class="form-control" id='Drawnratio' value='' size='4'/>
                    </div>    
                  <!-- </div> -->
                  <!-- <div class="form-group must"> -->
                    <!-- <label class="col-sm-3 control-label">成果类型</label> -->
                    <div class="col-sm-4">
                      <input type='text' placeholder='设计系数' class="form-control" id='Designdratio' value='' size='4'/>
                    </div>    
                  </div>
                  <!-- <div class="form-group must">
                    <label class="col-sm-3 control-label">附件链接</label>
                    <div class="col-sm-8">
                      <input type='text' placeholder='http://' class="form-control" id='Link1' name='Links' value='http://' size='4'/>
                    </div>    
                  </div> -->
                  <div class="form-group must">
                    <label class="col-sm-3 control-label">附件链接</label>
                    <div class="col-sm-8">
                      <textarea placeholder='http://，多个链接地址用,号隔开' class="form-control" rows="3" id='Link'></textarea>
                    </div>    
                  </div>
                  <div class="form-group must">
                    <label class="col-sm-3 control-label">成果说明</label>
                    <div class="col-sm-8">
                      <textarea class="form-control" rows="3" id='Content'></textarea>
                    </div>    
                  </div>
                  <div class="form-group">
                    <label class="col-sm-3 control-label">出版日期</label>
                    <div class="col-sm-3">
                      <span style="position: relative;z-index: 9999;">
                        <input type="text" class='datepicker' id='Date'/>
                      </span>
                    </div>    
                  </div>
                </div>
              </div>
              <div class="modal-footer">
                <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
                <button type="button" class="btn btn-primary" onclick="saveAddRow()">保存</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- <script type="text/javascript">
      $("#Date").datepicker({
        weekStart: 1,
        language: "zh-CN",
        autoclose: true,//选中之后自动隐藏日期选择框
        clearBtn: true,//清除按钮
        todayBtn: 'linked',//今日按钮
        setDate:moment(),
        todayHighlight:true,
        format: "yyyy-mm-dd"//日期格式，详见 http://bootstrap-datepicker.readthedocs.org/en/release/options.html#format
      });

      $(document).ready(function() {
        var now = new Date(); 
        myDate=new Date().getFullYear() + '-' + (new Date().getMonth() + 1) + '-' + new Date().getDate(); 
        $("#Date").val(myDate);
      })
    </script> -->

  </div>

<script type="text/javascript">
  function index1(value,row,index){
  // alert( "Data Loaded: " + index );
    return index+1
  }

  function setAttachment(value,row,index){
    if (value){
      // if (value.length==1){
        // attachUrl= '<a href="'+value[0].Url+'" title="下载" target="_blank"><i class="fa fa-paperclip"></i></a>';
        // return attachUrl;
      // }else if(value.length==0){
                    
      // }else if(value.length>1){
        attachUrl= "<a class='attachment' href='javascript:void(0)' title='查看附件列表'><i class='fa fa-list-ol'></i></a>";
        return attachUrl;
      // }
    }
  }

  function setContent(value,row,index){
    if (row.State==1||row.State==2){
      // if (value.length==1){
      //   attachUrl= '<a href="http://'+value[0].Url+'" title="下载" target="_blank"><i class="fa fa-paperclip"></i></a>';
      //   return attachUrl;
      // }else 
      // if(value.length==0){
                    
      // }else if(value.length>=1){
        contentUrl= "<a class='content' href='javascript:void(0)' title='查看设计说明'><i class='fa fa-list-alt'></i></a>";
        return contentUrl;
      // }
    }else if(row.State==3){
      contentUrl= "<a class='content' href='javascript:void(0)' title='查看意见'><i class='fa fa-list-alt'></i></a>";
        return contentUrl;
    }else if(row.State==4){
      contentUrl= "<a class='content' href='javascript:void(0)' title='查看意见'><i class='fa fa-list-alt'></i></a>";
        return contentUrl;
    }
  }

  function actionFormatter(value, row, index) {
    return [
        '<a class="send" href="javascript:void(0)" title="提交">',
        '<i class="glyphicon glyphicon-step-forward"></i>',
        '</a>&nbsp;',
        '<a class="downsend" href="javascript:void(0)" title="退回">',
        '<i class="glyphicon glyphicon-step-backward"></i>',
        '</a>&nbsp;',
        '<a class="remove" href="javascript:void(0)" title="删除">',
        '<i id="delete" class="glyphicon glyphicon-remove"></i>',
        '</a>'
    ].join('');
  }

  //最后面弹出附件列表中用的
  function setAttachlink(value,row,index){
    attachUrl= '<a href="'+value+'" title="下载" target="_blank"><i class="fa fa-paperclip"></i></a>';
      return attachUrl;
  }

  //别人发起，我设计，不提供删除功能的操作
  function actionFormatter1(value, row, index) {
    return [
        '<a class="send" href="javascript:void(0)" title="提交">',
        '<i class="glyphicon glyphicon-step-forward"></i>',
        '</a>',
        '<a class="downsend" href="javascript:void(0)" title="退回">',
        '<i class="glyphicon glyphicon-step-backward"></i>',
        '</a>',
    ].join('');
  }

    //待选择的修改*******不要删除
    //我发起
    $(function (value, sourceData) {
      $('#table').bootstrapTable({
        idField: 'id',
        uniqueId:'id',
        // url: '/getprogress',///achievement/send/1
        // striped: "true",
        columns: [
          {
            checkbox:true,
            // width:8
          },
          {
            // field: 'Number',
            // title: '序号',
            formatter:function(value,row,index){
              return index+1
            }
          },{
            visible: true,
            // width:30
          },{
            visible: true,
            // width:5,
            editable: {
              params:function(params) {
                //originally params contain pk, name and value
                var selectRow3=$('#table').bootstrapTable('getSelections');
                  var ids="";
                  for(var i=0;i<selectRow3.length;i++){
                    if(i==0){
                      ids=selectRow3[i].id;
                    }else{
                      ids=ids+","+selectRow3[i].id;
                    }  
                  }
                params.ids = ids;
                return params;
              },
              type: 'select2',
              select2: {
                allowClear: true,
                width: '150px',
                placeholder: '请选择',
                // multiple: true
                placeholder: 'Select Task Progress',
                // minimumInputLength: 0,
                id: function (item) {
                    return item.Id;
                },
                ajax: {
                    url: '/getselect',//这里要修改一下
                    dataType: 'json',
                    data: function (term, page) {
                        return { query: term };
                    },
                    results: function (data, page) {
                      // alert(JSON.stringify(data));
                        return { results: data };
                    }
                },
                formatResult: function (item) {
                  // return item.Nickname;
                  // alert(item.name);
                  var baseUrl = "/static/img";
                  var $state = $('<span><img src="' + baseUrl + '/' + item.name1 + '.png" class="img-flag" width="20%"/> ' + item.content + '</span>');
                  return $state;
                },
                formatSelection: function (item) {
                  // alert(item.name);
                  // return item.name;
                  var baseUrl = "/static/img";
                  var $state = $('<span><img src="' + baseUrl + '/' + item.name1 + '.png" class="img-flag" width="20%"/> ' + item.content + '</span>');
                  return $state;
                },
                initSelection: function (element, callback) {//这个是默认值
                  // alert(JSON.stringify(element));
                  // alert(element.val());
                  // return {"Id":1,"name":"0","content":"未启动"}
                  return $.get('/getprogress1', { query:  element.val()}, function (data) {
                      // alert(JSON.stringify(data)); 
                    callback(data);
                  });
                }
              },//'/regist/getuname1',//这里用get方法，所以要换一个
              pk: 1,
              url: '/modifyprogress',
              title: 'Enter Progress',
              success: function(response, newValue) {
                // alert(response.name);//这个是反馈回来的值
                // alert(newValue);//这个是选择的顺序值
                var selectRow3=$('#table').bootstrapTable('getSelections');
                for(var i=0;i<selectRow3.length;i++){
                  alert(selectRow3[i].Id);
                  $('#table').bootstrapTable('updateByUniqueId', {
                      id: selectRow3[i].Id,
                      row: {
                          name1: response.name1
                      }
                  });
                }
              }  
            }
          }
        ],
      });
    });

  function localDateFormatter(value) {
     return moment(value, 'YYYY-MM-DD').format('L');
  }
  function nameFormatter(value) {
    return '<a href="https://github.com/wenzhixin/' + value + '">' + value + '</a>';
  }
  //这个是显示时间选择
  function datepicker(value) {
    $(".datepicker").datepicker({
      language: "zh-CN",
      autoclose: true,//选中之后自动隐藏日期选择框
      clearBtn: true,//清除按钮
      todayBtn: 'linked',//今日按钮
      format: "yyyy-mm-dd"//日期格式，详见 http:// atepicker.readthedocs.org/en/release/options.html#format
    });
  }

    function queryParams(params) {
      // var newPage = $("#txtPage").val();
      var date=$("#datefilter").val();
      params.datefilter=date;//"2016-09-10 - 2016-09-15";
        // params.your_param1 = 1; // add param1
        // params.your_param2 = 2; // add param2
        // console.log(JSON.stringify(params));
        // {"limit":10,"offset":0,"order":"asc","your_param1":1,"your_param2":2}
        return params;
    }

    // var $table = $('#table'),
    // $button = $('#button');
    $(function () {
        $('#button').click(function () {
            $('#table').bootstrapTable('refresh', {url:'/achievement/send/1'});
            $('#table1').bootstrapTable('refresh', {url:'/achievement/send/2'});
            $('#table2').bootstrapTable('refresh', {url:'/achievement/send/3'});
            $('#table3').bootstrapTable('refresh', {url:'/achievement/send/4'});
        });
    });    

    $(document).ready(function(){
        $("#sel_Province").change(function(){
          $.ajax({
            url: '<%=basePath%>areaAjax/getCity.do',
            data: "procode="+$("#sel_Province").val(),
            type: 'get',
            dataType:'json',
            error: function(data)
            {
              alert("加载json 文件出错！");
            },
            success: function(data)
            {
              for (var one in data){
                var name = data[one].name;
                var code = data[one].code;
                $("#sel_City").append("<option value="+code+">"+name+"</option>");
              }
            },
          });
        });
    });

    $(document).ready(function(){
    $.each({{.Select2}},function(i,d){
      $("#Category").append('<option value="' + i + '">'+d+'</option>');
      });
    });

</script>

<script>
$(function(){
    $('#country').editable({
        select2: {
            placeholder: 'Select Country',
            width: '150px',
            allowClear: true,
            minimumInputLength: 3,
            id: function (item) {
                return item.Id;
            },
            ajax: {
                url: '/regist/getuname',
                dataType: 'json',
                data: function (term, page) {
                    return { query: term };
                },
                results: function (data, page) {
                    return { results: data };
                }
            },
            formatResult: function (item) {
              // return item.Nickname;
              var baseUrl = "/static/img";
              var $state = $('<span><img src="' + baseUrl + '/' + '0.75.png" class="img-flag" width="20%"/> ' + item.Nickname + '</span>');
              return $state;
            },
            formatSelection: function (item) {
              // return item.Nickname;
              var baseUrl = "/static/img";
              var $state = $('<span><img src="' + baseUrl + '/' + '0.75.png" class="img-flag" width="20%"/> ' + item.Nickname + '</span>');
              return $state;
            },
            initSelection: function (element, callback) {
              return $.get('/regist/getuname', { query: element.val() }, function (data) {
                  callback(data);
                });
            } 
        }  
    });
});

var scope = this;

scope.taskFormatter = function(value) {
  var icon = $(this).data('pk') % 2 === 0 ? 'glyphicon-star' : 'glyphicon-star-empty';
  var markup = '<i class="glyphicon ' + icon + '"></i> ' + value;
  $(this).html(markup);
};

scope.nameFormatter = function(value) {
  var icon = $(this).data('pk') % 2 === 0 ? 'glyphicon-star' : 'glyphicon-star-empty';
  var markup = '<i class="glyphicon ' + icon + '"></i> ' + value;
  $(this).html(markup);
};

scope.priceFormatter = function(value) {
  // alert(value);
  if (value=="0"){
    title="未开始"
  }else if(value=="0.125"){
    title="完成1/8"
  }else if(value=="0.25"){
    title="完成1/4"
  }else if(value=="0.375"){
    title="完成3/8"
  }else if(value=="0.5"){
    title="完成1/2"
  }else if(value=="0.625"){
    title="完成5/8"
  }else if(value=="0.75"){
    title="完成3/4"
  }else if(value=="0.875"){
    title="完成7/8"
  }else if(value=="1"){
    title="任务完成"
  };
  var color = '#' + Math.floor(Math.random() * 6777215).toString(16);
  var markup = '<div title="' + title + '" style="color: ' + color + '">' + '<span><img src="/static/img/'  + value + '.png" class="img-flag" width="80%"/> </span></div>';
  $(this).html(markup);
};

scope.priceFormatter1 = function(value) {
if (value=="dashed"){
    title="未开始"
  }else if(value=="solid"){
    title="已开始"
  }
  var color = '#' + Math.floor(Math.random() * 6777215).toString(16);
  var markup = '<div title="' + title + '" style="color: ' + color + '">' + '<span><img src="/static/img/'  + value + 'arrow-redd.png" class="img-flag" width="100%"/> </span></div>';
  $(this).html(markup);
};

// $.fn.editable.defaults.mode = 'inline';
$.fn.editable.defaults.display = function(value, sourceData) {
  var displayFormatterFunction = $(this).data('display-formatter');
  if (displayFormatterFunction) {
    scope[displayFormatterFunction].call(this, value, sourceData);
  } else {
    $(this).html(value);
  }
};

</script>

</body>
</html>


<!-- $(".js-example-data-ajax").select2({
  ajax: {
    url: "https://api.github.com/search/repositories",
    dataType: 'json',
    delay: 250,
    data: function (params) {
      return {
        q: params.term, // search term
        page: params.page
      };
    },
    processResults: function (data, params) {
      // parse the results into the format expected by Select2
      // since we are using custom formatting functions we do not need to
      // alter the remote JSON data, except to indicate that infinite
      // scrolling can be used
      params.page = params.page || 1;

                //     return {
                //       results: data.items,
                //       pagination: {
                //         more: (params.page * 30) < data.total_count
                //       }
                //     };
                //   },
                //   cache: true
                // },
                // placeholder: 'Search for a repository',
                // escapeMarkup: function (markup) { return markup; }, // let our custom formatter work
                // minimumInputLength: 1,
                // templateResult: formatRepo,
                // templateSelection: formatRepoSelection

                // },
 
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter DesignStage', 
            }
          },
        ],
      });


  });

function formatRepo (repo) {
  if (repo.loading) {
    return repo.text;
  }

  var markup = "<div class='select2-result-repository clearfix'>" +
    "<div class='select2-result-repository__avatar'><img src='" + repo.owner.avatar_url + "' /></div>" +
    "<div class='select2-result-repository__meta'>" +
      "<div class='select2-result-repository__title'>" + repo.full_name + "</div>";

  if (repo.description) {
    markup += "<div class='select2-result-repository__description'>" + repo.description + "</div>";
  }

  markup += "<div class='select2-result-repository__statistics'>" +
    "<div class='select2-result-repository__forks'><i class='fa fa-flash'></i> " + repo.forks_count + " Forks</div>" +
    "<div class='select2-result-repository__stargazers'><i class='fa fa-star'></i> " + repo.stargazers_count + " Stars</div>" +
    "<div class='select2-result-repository__watchers'><i class='fa fa-eye'></i> " + repo.watchers_count + " Watchers</div>" +
  "</div>" +
  "</div></div>";

  return markup;
}

function formatRepoSelection (repo) {
  return repo.full_name || repo.text;
}



***********templateting
function formatState (state) {
  if (!state.id) {
    return state.text;
  }
  var baseUrl = "/user/pages/images/flags";
  var $state = $(
    '<span><img src="' + baseUrl + '/' + state.element.value.toLowerCase() + '.png" class="img-flag" /> ' + state.text + '</span>'
  );
  return $state;
};

$(".js-example-templating").select2({
  templateSelection: formatState
});


***********x-editable
<a href="javascript:void(0)" id="country" data-type="select2" data-pk="1" data-value="ru" data-url="/post" data-title="Select country"></a>
<script>
$(function(){
    //local source
    $('#country').editable({
        source: [
              {id: 'gb', text: 'Great Britain'},
              {id: 'us', text: 'United States'},
              {id: 'ru', text: 'Russia'}
           ],
        select2: {
           multiple: true
        }
    });
    //remote source (simple)
    $('#country').editable({
        source: '/getCountries',
        select2: {
            placeholder: 'Select Country',
            minimumInputLength: 1
        }
    });
    //remote source (advanced)
    $('#country').editable({
        select2: {
            placeholder: 'Select Country',
            allowClear: true,
            minimumInputLength: 3,
            id: function (item) {
                return item.CountryId;
            },
            ajax: {
                url: '/getCountries',
                dataType: 'json',
                data: function (term, page) {
                    return { query: term };
                },
                results: function (data, page) {
                    return { results: data };
                }
            },
            formatResult: function (item) {
                return item.CountryName;
            },
            formatSelection: function (item) {
                return item.CountryName;
            },
            initSelection: function (element, callback) {
                return $.get('/getCountryById', { query: element.val() }, function (data) {
                    callback(data);
                });
            } 
        }  
    });
});
</script> -->
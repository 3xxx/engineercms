<!-- 文档列表 -->
<!DOCTYPE html>
<head>

  <title>fei-ONLYOFFICE</title>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  <script src="/static/js/tableExport.js"></script>
  <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
  <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.min.js"> </script>
    <!--建议手动加在语言，避免在ie下有时因为加载语言失败导致编辑器加载失败-->
    <!--这里加载的语言文件会覆盖你在配置项目里添加的语言类型，比如你在配置项目里配置的是英文，这里加载的中文，那最后就是中文-->
  <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/webuploader.css">
  <script type="text/javascript" src="/static/js/webuploader.min.js"></script>
  
  <script type="text/javascript" src="/static/js/jquery-ui.min.js"></script>
  <script type="text/javascript" src="/static/bootstrap-datepicker/bootstrap-datepicker.js"></script>
  <script type="text/javascript" src="/static/bootstrap-datepicker/bootstrap-datepicker.zh-CN.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/bootstrap-datepicker/bootstrap-datepicker3.css"/>
  <style type="text/css">
      /*.form-group .datepicker{
        z-index: 9999;
      }*/
  /*模态框效果*/
    /*.modal-header {*/
      /*background: #00FF00;*/
      /*min-height: 16.42857143px;
      padding: 15px;
      border-bottom: 1px solid #e5e5e5;*/
    /*}*/
    /*.col-sm-1 input[type=checkbox]{
　　display: inline-block;
　　vertical-align: middle;
　　margin-bottom: 2px; 
    }*/
  </style>
</head>

<div class="container-fill">{{template "navbar" .}}</div>

<body>

<div class="col-lg-12">
  <h3>文档列表</h3>
  <div id="toolbar1" class="btn-group">
        <!-- 多文件批量上传 -->
        <button type="button" data-name="addButton" id="addButton" class="btn btn-default" title="批量上传模式"> <i class="fa fa-plus">添加</i>
        </button>
        <!-- 多附件上传 -->
        <!-- <button type="button" data-name="addButton1" id="addButton1" class="btn btn-default"> <i class="fa fa-plus-square-o" title="多附件模式">添加</i>
        </button> -->
        <!-- 添加文章 -->
        <!-- <button type="button" data-name="addButton2" id="addButton2" class="btn btn-default"> <i class="fa fa-plus-square" title="文章模式">添加</i>
        </button> -->
        <button type="button" data-name="editorProdButton" id="editorProdButton" class="btn btn-default"> <i class="fa fa-edit" title="修改成果信息">编辑</i>
        </button>
        <button type="button" data-name="editorAttachButton" id="editorAttachButton" class="btn btn-default"> <i class="fa fa-edit" title="修改成果附件">编辑</i>
        </button>
        <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
        </button>
        <!-- <button type="button" data-name="synchIP" id="synchIP" class="btn btn-default">
        <i class="fa fa-refresh">同步</i>
        </button> -->
  </div>
<!--data-click-to-select="true" -->
  <table id="table0" 
        data-toggle="table" 
        data-url="/onlyoffice/data"
        data-search="true"
        data-show-refresh="true"
        data-show-toggle="true"
        data-show-columns="true"
        data-toolbar="#toolbar1"
        data-query-params="queryParams"
        data-sort-name="Code"
        data-sort-order="desc"
        data-page-size="15"
        data-page-list="[10,15, 50, 100, All]"
        data-unique-id="id"
        data-pagination="true"
        data-side-pagination="client"
        data-single-select="true"
        data-click-to-select="true"
        data-show-export="true"
        >
    <thead>        
      <tr>
        <!-- radiobox data-checkbox="true" data-formatter="setCode" data-formatter="setTitle"-->
        <th data-width="10" data-radio="true"></th>
        <th data-formatter="index1" data-align="center">#</th>
        <th data-field="Code" data-halign="center">编号</th>
        <th data-field="Title" data-halign="center">名称</th>
        <th data-field="Label" data-formatter="setLable" data-halign="center" data-align="center">关键字</th>
        <th data-field="Principal" data-halign="center" data-align="center">负责人</th>
        <th data-field="Docxlink" data-formatter="setDocx" data-events="actionEvents" data-halign="center" data-align="center">协作</th>
        <!-- <th data-field="Xlsxlink" data-formatter="setXlsx" data-events="actionEvents" data-halign="center" data-align="center">XLSX</th> -->
        <!-- <th data-field="Pptxlink" data-formatter="setPptx" data-events="actionEvents" data-halign="center" data-align="center">PPTX</th> -->
        <th data-field="End" data-formatter="localDateFormatter" data-halign="center" data-align="center">结束时间</th>
        <th data-field="Created" data-formatter="localDateFormatter" data-halign="center" data-visible="false" data-align="center">建立时间</th>
        <th data-field="Updated" data-formatter="localDateFormatter" data-halign="center" data-align="center">更新时间</th>
      </tr>
    </thead>
  </table>

<script type="text/javascript">
  function index1(value,row,index){
    return index+1
  }

  function localDateFormatter(value) {
    return moment(value, 'YYYY-MM-DD').format('YYYY-MM-DD');
  }

  function setCode(value,row,index){
    return "<a href='/attachment/onlyoffice/"+value+"'>" + value + "</a>";
  }

  function setLable(value,row,index){
    // alert(value);
    if (value){//注意这里如果value未定义则出错，一定要加这个判断。
      var array=value.split(",")
      var labelarray = new Array() 
      for (i=0;i<array.length;i++)
      {
        labelarray[i]="<a href='/project/product/keysearch?keyword="+array[i]+"'>" + array[i] + "</a>";
      }
        return labelarray.join(",");
      }
  } 

  function setCodetest(value,row,index){
    //保留，数组和字符串以及循环的处理
    // array=value.split(",")
    // var labelarray = new Array() 
    // for (i=0;i<value.length;i++)//value是数组"Code":[数组"SL0001-510-08","SL0001-510-08"],
    // {
    //   labelarray[i]="<a href='/project/product/attachment/"+value[i]+"'>" + value[i] + "</a>";
    // }
    // if (value.match(",")!=null){
    if (value){
      array=value.split(",")
      var labelarray = new Array() 
      for (i=0;i<array.length;i++)
      {
        labelarray[i]="<a href='/project/product/attachment/"+array[i]+"'>" + array[i] + "</a>";
      }
      return labelarray.join(",");
    }
  }

  function setTitle(value,row,index){
    return "<a href='/attachment/onlyoffice/"+value+"'>" + value + "</a>";
  }

  function setDocx(value,row,index){
    if (value){
      if (value.length==1){
        if (value[0].Suffix=="docx"){
          docUrl= '<a href=/onlyoffice/'+value[0].Id+' title="协作" target="_blank"><i class="fa fa-file-word-o fa-lg"></i></a>';
          return docUrl;
        }else if(value[0].Suffix=="xlsx"){
          xlsUrl= '<a href=/onlyoffice/'+value[0].Id+' title="协作" target="_blank"><i class="fa fa-file-excel-o fa-lg" style="color:LimeGreen;"></i></a>';
          return xlsUrl;
        }else if(value[0].Suffix=="pptx"){
          pptUrl= '<a href=/onlyoffice/'+value[0].Id+' title="协作" target="_blank"><i class="fa fa-file-powerpoint-o fa-lg" style="color:Red;"></i></a>';
          return pptUrl;
        }else if(value[0].Suffix=="pdf"){
          pdfUrl= '<a href=/onlyoffice/'+value[0].Id+' title="协作" target="_blank"><i class="fa fa-file-pdf-o fa-lg" style="color:Brown;"></i></a>';
          return pdfUrl;
        }else if(value[0].Suffix=="txt"){
          txtUrl= '<a href=/onlyoffice/'+value[0].Id+' title="协作" target="_blank"><i class="fa fa-file-text-o fa-lg" style="color:black;"></i></a>';
          return txtUrl;
        }
        
      }else if(value.length==0){
                    
      }else if(value.length>1){
        fileUrl= "<a class='Docx' href='javascript:void(0)' title='查看文档列表'><i class='fa fa-list-ol'></i></a>";
        return fileUrl;
      }
    }
  }

  // function setXlsx(value,row,index){
  //   if (value){
  //     if (value.length==1){
  //       attachUrl= '<a href=/onlyoffice/'+value[0].Id+' title="下载" target="_blank"><i class="fa fa-file-excel-o"></i></a>';
  //       return attachUrl;
  //     }else if(value.length==0){
                    
  //     }else if(value.length>1){
  //       attachUrl= "<a class='Xlsx' href='javascript:void(0)' title='查看Xlsx列表'><i class='fa fa-list-ol'></i></a>";
  //       return attachUrl;
  //     }
  //   }
  // }

  // function setPptx(value,row,index){
  //   if (value){
  //     if (value.length==1){
  //       pdfUrl= '<a href=/onlyoffice/'+value[0].Id+' title="打开pdf" target="_blank"><i class="fa fa-file-powerpoint-o"></i></a>';
  //       return pdfUrl;
  //     }else if(value.length==0){
                    
  //     }else if(value.length>1){
  //       pdfUrl= "<a class='Pptx' href='javascript:void(0)' title='查看Pptx列表'><i class='fa fa-list-ol'></i></a>";
  //       return pdfUrl;
  //     }
  //   }
  // }

  window.actionEvents = {
    'click .Docx': function (e, value, row, index) {
        $('#docx').bootstrapTable('refresh', {url:'/onlyoffice/docx/'+row.Id});
      $('#modaldocx').modal({
        show:true,
        backdrop:'static'
      }); 
    },

    'click .Xlsx': function (e, value, row, index) {
        $('#xlsx').bootstrapTable('refresh', {url:'/onlyoffice/xlsx/'+row.Id});
      $('#modalxlsx').modal({
        show:true,
        backdrop:'static'
      }); 
    },

    'click .Pptx': function (e, value, row, index) {
        $('#pptx').bootstrapTable('refresh', {url:'/onlyoffice/pptx/'+row.Id});
      $('#modalpptx').modal({
        show:true,
        backdrop:'static'
      }); 
    },
  };

  //最后面弹出doc列表中用的_根据上面的click，弹出模态框，给模态框中的链接赋值
  function setDocxlink(value,row,index){
    docxUrl= '<a href="'+value+'" title="下载" target="_blank"><i class="fa fa-file-text-o"></i></a>';
      return docxUrl;
  }
  //最后面弹出附件列表中用的<a href="'+value+
  function setXlsxlink(value,row,index){
    xlsxUrl= '<a href="/onlyoffice?id='+row.Id+'" title="下载" target="_blank"><i class="fa fa-paperclip"></i></a>';
      return xlsxUrl;
  }
  //最后面弹出pdf列表中用的'&file='+value+
  function setPptxlink(value,row,index){
    pptxUrl= '<a href="/onlyoffice?id='+row.Id+'" title="下载" target="_blank"><i class="fa fa-file-pdf-o"></i></a>';
      return pptxUrl;
  }

  // 批量上传
  $("#addButton").click(function() {
      // if ({{.RoleAdd}}!="true"){
      //   alert("权限不够！");
      //   return;
      // }
      $("input#pid").remove();
      var th1="<input id='pid' type='hidden' name='pid' value='" +{{.Id}}+"'/>"
        $(".modal-body").append(th1);
        $('#modalTable').modal({
        show:true,
        backdrop:'static'
        });
  })

  $(document).ready(function() {
    $list1 = $('#thelist');
    $btn = $('#ctlBtn');
    state = 'pending';
    // $('#modalTable').on('shown.bs.modal',function(e){
      var allMaxSize = 100;
      var uploader=WebUploader.create({
        // 不压缩image
        resize: false,
        fileSingleSizeLimit: 10*1024*1024,//限制大小10M，单文件
        fileSizeLimit: allMaxSize*1024*1024,//限制大小10M，所有被选文件，超出选择不上
        // swf文件路径
        swf: '/static/js/Uploader.swf',
        // 文件接收服务端。
        server: '/onlyoffice/addattachment',
        // 选择文件的按钮。可选。
        // 内部根据当前运行是创建，可能是input元素，也可能是flash.
        pick: '#picker',
        // 只允许选择规定文件类型。
        accept: {
            title: 'Images',
            extensions: 'doc,docx,xls,xlsx,ppt,pptx,txt,pdf',
            mimeTypes: '*/*'
        }
      });
      /**
     * 验证文件格式以及文件大小
     */
      uploader.on("error",function (type){
        if (type == "F_DUPLICATE") {
              alert("请不要重复选择文件！");
         } else if (type == "Q_EXCEED_SIZE_LIMIT"){
              alert("所选附件总大小不可超过" + allMaxSize + "M！多分几次传吧！");
         }else if (type=="Q_TYPE_DENIED"){
          alert("请上传文档格式文件");
        }else if(type=="F_EXCEED_SIZE"){
          alert("单个文件大小不能超过10M");
        }
      });

      // 当有文件添加进来的时候
      uploader.on( 'fileQueued', function( file ) {
       $list1.append( '<div id="' + file.id + '" class="item">' +
            '<h4 class="info">' + file.name + '</h4>' +
            '<p class="state">等待上传...</p>' +
        '</div>');
      }); 

      //传递参数——成果id
      uploader.on( 'startUpload', function() {//uploadBeforeSend——这个居然不行？
      // if (prodlabel){
        var pid = $('#pid').val();
        var prodlabel = $('#prodlabel').val();
        var prodprincipal = $('#prodprincipal').val();
        var newDate = $("#Date").val();
        uploader.option('formData', {
          "pid":pid,
          "prodlabel":prodlabel,
          "prodprincipal":prodprincipal,
          "proddate":newDate
        }); 
      });

      // 文件上传过程中创建进度条实时显示。
      uploader.on( 'uploadProgress', function( file, percentage ) {
        var $li = $( '#'+file.id ),
            $percent = $li.find('.progress .progress-bar');
        // 避免重复创建
        if ( !$percent.length ) {
            $percent = $('<div class="progress progress-striped active">' +
              '<div class="progress-bar" role="progressbar" style="width: 0%">' +
              '</div>' +
            '</div>').appendTo( $li ).find('.progress-bar');
        }
        $li.find('p.state').text('上传中');
        $percent.css( 'width', percentage * 100 + '%' );
      });

      uploader.on( 'uploadSuccess', function( file ) {
        $( '#'+file.id ).find('p.state').text('已上传');
      });

      uploader.on( 'uploadError', function( file ) {
        $( '#'+file.id ).find('p.state').text('上传出错');
      });

      uploader.on( 'uploadComplete', function( file ) {
        $( '#'+file.id ).find('.progress').fadeOut();
        $('#table0').bootstrapTable('refresh', {url:'/onlyoffice/data'});
      });

      uploader.on( 'all', function( type ) {
        if ( type === 'startUpload' ) {
            state = 'uploading';
        } else if ( type === 'stopUpload' ) {
            state = 'paused';
        } else if ( type === 'uploadFinished' ) {
            state = 'done';
        }
        if ( state === 'uploading' ) {
            $btn.text('暂停上传');
        } else {
            $btn.text('开始上传');
        }
      });
    
      $btn.on( 'click', function() {
        if ( state === 'uploading' ) {
            uploader.stop();
        } else {
            uploader.upload();
        }
      });

      $('#picker').mouseenter(function(){
        uploader.refresh();
      })

      $('#modalTable').on('hide.bs.modal',function(){
        $list1.text("");
        // uploader.destroy();//销毁uploader
      })
  })

  // 编辑成果信息
  $("#editorProdButton").click(function() {
      var selectRow=$('#table0').bootstrapTable('getSelections');
      if (selectRow.length<1){
        alert("请先勾选成果！");
        return;
      }
      if (selectRow.length>1){
        alert("请不要勾选一个以上成果！");
        return;
      }

      $("input#cid").remove();
      var th1="<input id='cid' type='hidden' name='cid' value='" +selectRow[0].Id+"'/>"
      $(".modal-body").append(th1);//这里是否要换名字$("p").remove();
      $("#prodcode3").val(selectRow[0].Code);
      $("#prodname3").val(selectRow[0].Title);
      $("#prodlabel3").val(selectRow[0].Label);
      $("#prodprincipal3").val(selectRow[0].Principal);
      //selectRow[0].End
      $("#proddate3").val(moment(selectRow[0].End, 'YYYY-MM-DD').format('YYYY-MM-DD'));

      $('#modalProdEditor').modal({
      show:true,
      backdrop:'static'
      });
  })

  // 编辑成果附件——删除附件、文章或追加附件
  var selectrowid;
  $("#editorAttachButton").click(function() {
      var selectRow=$('#table0').bootstrapTable('getSelections');
      if (selectRow.length<1){
        alert("请先勾选成果！");
        return;
      }
      if (selectRow.length>1){
      alert("请不要勾选一个以上成果！");
      return;
      }

      selectrowid=selectRow[0].Id;
      $("input#pid").remove();
      var th1="<input id='pid' type='hidden' name='pid' value='" +selectRow[0].Id+"'/>"
      $(".modal-body").append(th1);//这里是否要换名字$("p").remove();
      $('#attachments').bootstrapTable('refresh', {url:'/onlyoffice/'+selectRow[0].Id});//取得所有附件列表和文章列表
      $('#modalAttachEditor').modal({
      show:true,
      backdrop:'static'
      });
      // }else{
      //   alert("权限不够！"+selectRow[0].Uid);
      //   return;
      // }
  })

  $(document).ready(function() {
    var uploader;
    $('#modalAttachEditor').on('shown.bs.modal',function(){
      // var $ = jQuery,
      $list2 = $('#thelist2'),
      $btn = $('#ctlBtn2'),
      state = 'pending',
      // uploader;
      uploader = WebUploader.create({
        // 不压缩image
        resize: false,
        // swf文件路径
        swf: '/static/js/Uploader.swf',
        // 文件接收服务端。
        server: '/onlyoffice/updateattachment',
        // 选择文件的按钮。可选。
        // 内部根据当前运行是创建，可能是input元素，也可能是flash.
        pick: '#picker2'
      });
      // 当有文件添加进来的时候
      uploader.on( 'fileQueued', function( file ) {
        $list2.append( '<div id="' + file.id + '" class="item">' +
              '<h4 class="info">' + file.name + '</h4>' +
              '<p class="state">等待上传...</p>' +
          '</div>' );
      }); 

      //传递参数——成果id
      uploader.on( 'startUpload', function() {//uploadBeforeSend——这个居然不行？
        var pid = $('#pid').val();
        uploader.option('formData', {
          "pid":pid,
        });        
      });

      // 文件上传过程中创建进度条实时显示。
      uploader.on( 'uploadProgress', function( file, percentage ) {
        var $li = $( '#'+file.id ),
        $percent = $li.find('.progress .progress-bar');
        // 避免重复创建
        if ( !$percent.length ) {
          $percent = $('<div class="progress progress-striped active">' +
                '<div class="progress-bar" role="progressbar" style="width: 0%">' +
              '</div>' +
              '</div>').appendTo( $li ).find('.progress-bar');
        }
        $li.find('p.state').text('上传中');
        $percent.css( 'width', percentage * 100 + '%' );
      });

      uploader.on( 'uploadSuccess', function( file ) {
          $( '#'+file.id ).find('p.state').text('已上传');
      });

      uploader.on( 'uploadError', function( file ) {
          $( '#'+file.id ).find('p.state').text('上传出错');
      });

      uploader.on( 'uploadComplete', function( file ) {
          $( '#'+file.id ).find('.progress').fadeOut();
          $('#attachments').bootstrapTable('refresh', {url:'/onlyoffice/'+selectrowid});
          $('#table0').bootstrapTable('refresh', {url:'/onlyoffice/data'});
      });

      uploader.on( 'all', function( type ) {
        if ( type === 'startUpload' ) {
            state = 'uploading';
        } else if ( type === 'stopUpload' ) {
            state = 'paused';
        } else if ( type === 'uploadFinished' ) {
            state = 'done';
        }
        if ( state === 'uploading' ) {
            $btn.text('暂停上传');
        } else {
            $btn.text('开始上传');
        }
      });
    
      $btn.on( 'click', function() {
          if ( state === 'uploading' ) {
              uploader.stop();
          } else {
              uploader.upload();
          }
      });
    })

    $('#modalAttachEditor').on('hide.bs.modal',function(){
      $list2.text("");
      uploader.destroy();//销毁uploader
    })
  })
    
  // 删除成果
  $("#deleteButton").click(function() {
    if ({{.IsAdmin}}){
      var selectRow=$('#table0').bootstrapTable('getSelections');
      if (selectRow.length<=0) {
        alert("请先勾选成果！");
        return false;
      }
     //问题：如果多选，而其中有自己的，也有自己不具备权限的********
      // if (selectRow[0].Uid==={{.Uid}}||{{.RoleDelete}}=="true"){
      if(confirm("确定删除成果吗？一旦删除将无法恢复！")){
        var title=$.map(selectRow,function(row){
          return row.Title;
        })
        var ids="";
        for(var i=0;i<selectRow.length;i++){
          if(i==0){
            ids=selectRow[i].Id;
          }else{
            ids=ids+","+selectRow[i].Id;
          }  
        }
        $.ajax({
          type:"post",
          url:"/onlyoffice/deletedoc",
          data: {ids:ids},
          success:function(data,status){
            alert("删除“"+data+"”成功！(status:"+status+".)");
            //删除已选数据
            $('#table0').bootstrapTable('remove',{
              field:'Title',
              values:title
            });
          }
        });
      }
    }else{
      alert("权限不够！");
      return;
    }
  })

</script>
  <!-- 批量上传 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalTable">
      <div class="modal-dialog" id="modalDialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">批量添加成果</h3>
            
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <label>**请选择标准命名的电子文件上传：编号+名称**</label>
              <div class="form-group must">
                <label class="col-sm-3 control-label">关键字</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodlabel" placeholder="以英文,号分割"></div>
              </div>
              <div class="form-group">
                  <label class="col-sm-3 control-label">结束时间</label>
                  <div class="col-sm-3">
                    <span style="position: relative;z-index: 9999;">
                      <input type="text" class='datepicker' id='Date'/>
                    </span>
                  </div>    
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">负责人</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodprincipal" placeholder="输入姓名"></div>
              </div>
              <!--SWF在初始化的时候指定，在后面将展示-->
              <div id="uploader" style="position:relative;text-align: center;">
              <!--用来存放文件信息-->
                <div id="thelist"></div>
                <div class="btns">
                  <div id="picker">选择文件</div>
                  <button id="ctlBtn" class="btn btn-default">开始上传</button>
                </div>
              </div>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- 编辑成果名称等信息 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalProdEditor">
      <div class="modal-dialog" id="modalDialog6">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">编辑文档信息</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">编号</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="prodcode3"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">标题</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodname3"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">关键字</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodlabel3" placeholder="以英文,号分割"></div>
              </div>
              <div class="form-group">
                  <label class="col-sm-3 control-label">结束时间</label>
                  <div class="col-sm-3">
                    <span style="position: relative;z-index: 9999;">
                      <input type="text" class='datepicker' id='proddate3'/>
                    </span>
                  </div>    
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">负责人</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodprincipal3"></div>
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
  <!-- 编辑成果附件 删除附件或追加附件-->
  <div class="form-horizontal">
    <div class="modal fade" id="modalAttachEditor">
      <div class="modal-dialog" id="modalDialog7">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">编辑成果附件</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div id="attachtoolbar" class="btn-group">
                <button type="button" data-name="deleteAttachButton" id="deleteAttachButton" class="btn btn-default">
                <i class="fa fa-trash">删除</i>
                </button>
              </div>
              <table id="attachments"
                    data-toggle="table"
                    data-toolbar="#attachtoolbar"
                    data-page-size="5"
                    data-page-list="[5, 25, 50, All]"
                    data-unique-id="id"
                    data-pagination="true"
                    data-side-pagination="client"
                    data-click-to-select="true">
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
              <!--SWF在初始化的时候指定，在后面将展示-->
              <div id="uploader1" style="position:relative;text-align: center;">
              <!--用来存放文件信息-->
                <div id="thelist2"></div>
                <div class="btns">
                  <div id="picker2">选择文件</div>
                  <button id="ctlBtn2" class="btn btn-default">开始上传</button>
                </div>
              </div>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
          </div>
        </div>
      </div>
    </div>
  </div>

</div>

<script type="text/javascript">
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

    $("#proddate3").datepicker({
        weekStart: 1,
        language: "zh-CN",
        autoclose: true,//选中之后自动隐藏日期选择框
        clearBtn: true,//清除按钮
        todayBtn: 'linked',//今日按钮
        setDate:moment(),
        todayHighlight:true,
        format: "yyyy-mm-dd"//日期格式，详见 http://bootstrap-datepicker.readthedocs.org/en/release/options.html#format
    });
  // 编辑成果信息
  function updateprod(){
    // var radio =$("input[type='radio']:checked").val();
    var projectid = $('#cid').val();
    var prodcode = $('#prodcode3').val();
    var prodname = $('#prodname3').val();
    var prodlabel = $('#prodlabel3').val();
    var prodprincipal = $('#prodprincipal3').val();
    var proddate = $('#proddate3').val();
 	  //alert(proddate);
    if (prodname&&prodcode){  
      $.ajax({
        type:"post",
        url:"onlyoffice/updatedoc",
        data: {pid:projectid,code:prodcode,title:prodname,label:prodlabel,principal:prodprincipal,proddate:proddate},//父级id
        success:function(data,status){
          alert("添加“"+data+"”成功！(status:"+status+".)");
          $('#modalProdEditor').modal('hide');
          $('#table0').bootstrapTable('refresh', {url:'/onlyoffice/data'});
        },
        
      });
    }else{
      alert("请填写编号和名称！");
      return;
    }
  }

  // 删除附件
  $("#deleteAttachButton").click(function() {
      // if ({{.role}}!=1){
      //   alert("权限不够！");
      //   return;
      // }
      if ({{.RoleDelete}}!="true"){
        alert("权限不够！");
        return;
      }
      var selectRow=$('#attachments').bootstrapTable('getSelections');
      if (selectRow.length<=0) {
        alert("请先勾选！");
        return false;
      }

      if ({{.RoleDelete}}!="true"){
        alert("权限不够！"+selectRow[0].Uid);
        return;
      }

      if(confirm("确定删除吗？一旦删除将无法恢复！")){
        var title=$.map(selectRow,function(row){
          return row.Title;
        })
        var ids="";
        for(var i=0;i<selectRow.length;i++){
          if(i==0){
            ids=selectRow[i].Id;
          }else{
            ids=ids+","+selectRow[i].Id;
          }  
        }
        $.ajax({
          type:"post",
          url:"/onlyoffice/deleteonlyattachment",
          data: {ids:ids},
          success:function(data,status){
            alert("删除“"+data+"”成功！(status:"+status+".)");
            //删除已选数据
            $('#attachments').bootstrapTable('remove',{
              field:'Title',
              values:title
            });
          }
        });
      }  
  })

    //******表格追加项目同步ip中的数据*******
    $(function () {
        $('#synchIP').click(function () {
          // alert("ha ");
          $.ajax({
            type:"get",
            url:"/project/synchproducts/"+{{.Id}},
            // data: {ids:ids},
            success:function(data,status){
              alert("同步成功！(status:"+status+".)");
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
              Code:startId + i,
              Title: 'test' + (startId + i),
                // id: startId + i,
                // name: 'test' + (startId + i),
                // price: '$' + (startId + i)
            });
        }
        return rows;
    }

    //勾选后输入框可用
    function station_select(){ 
      if(box.checked){ 
        document.getElementById("relevancy").disabled=false; 
      } else{ 
        document.getElementById("relevancy").disabled=true; 
      } 
    }

    $(document).ready(function(){
        $("#modalDialog").draggable();//为模态对话框添加拖拽
        $("#modalDialog1").draggable();
        $("#modalDialog2").draggable();
        $("#modalDialog3").draggable();
        $("#modalDialog4").draggable();
        $("#modalDialog5").draggable();
        $("#modalDialog6").draggable();
        $("#modalDialog7").draggable();
        $("#myModal").css("overflow", "hidden");//禁止模态对话框的半透明背景滚动
    })
</script>

</body>
</html>
<!-- 具体一个项目侧栏id下所有成果，不含子目录下的成果 -->
<!DOCTYPE html>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css"/>
  
  <script src="/static/js/bootstrap-treeview.js"></script>
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/font-awesome.min.css"/> -->
  <script src="/static/js/tableExport.js"></script>

  <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
  <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.js"> </script>
    <!--建议手动加在语言，避免在ie下有时因为加载语言失败导致编辑器加载失败-->
    <!--这里加载的语言文件会覆盖你在配置项目里添加的语言类型，比如你在配置项目里配置的是英文，这里加载的中文，那最后就是中文-->
  <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>

  <script type="text/javascript" src="/static/js/moment.min.js"></script>

  <!-- <link rel="stylesheet" type="text/css" href="/static/css/webuploader.css"> -->
  <!-- <script type="text/javascript" src="/static/js/webuploader.js"></script> -->
  <link rel="stylesheet" type="text/css" href="/static/fex-team-webuploader/css/webuploader.css">
<script type="text/javascript" src="/static/fex-team-webuploader/dist/webuploader.min.js"></script>
</head>

<body>

<div class="col-lg-12">
  <h3>成果列表</h3>
              <!--{{.Id}}SWF在初始化的时候指定，在后面将展示-->
              <!-- <div id="uploader" class="wu-example"> -->
              <!--用来存放文件信息-->
                <!-- <div id="thelist" class="uploader-list"></div> -->
                <!-- <div class="btns"> -->
                  <!-- <div id="picker">选择文件</div> -->
                  <!-- <button id="ctlBtn" class="btn btn-default">开始上传</button> -->
                <!-- </div> -->
              <!-- </div> -->
<div id="toolbar1" class="btn-group">
        <!-- 多文件批量上传 -->
        <button type="button" data-name="addButton" id="addButton" class="btn btn-default" title="批量上传模式"> <i class="fa fa-plus">添加</i>
        </button>
        <!-- 多附件上传 -->
        <button type="button" data-name="addButton1" id="addButton1" class="btn btn-default"> <i class="fa fa-plus-square-o" title="多附件模式">添加</i>
        </button>
        <!-- 添加文章 -->
        <button type="button" data-name="addButton2" id="addButton2" class="btn btn-default"> <i class="fa fa-plus-square" title="文章模式">添加</i>
        </button>
        <button type="button" data-name="editorProdButton" id="editorProdButton" class="btn btn-default"> <i class="fa fa-edit" title="修改成果信息">编辑</i>
        </button>
        <button type="button" data-name="editorAttachButton" id="editorAttachButton" class="btn btn-default"> <i class="fa fa-edit" title="修改成果附件">编辑</i>
        </button>
        <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
        </button>
</div>
<!--        data-click-to-select="true" -->
<table id="table0" 
        data-toggle="table" 
        data-url="/project/products/{{.Id}}"
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
        <th data-formatter="index1">#</th>
        <th data-field="Code">编号</th>
        <th data-field="Title">名称</th>
        <th data-field="Label" data-formatter="setLable">关键字</th>
        <th data-field="Principal">设计</th>
        <th data-field="Articlecontent" data-formatter="setArticle" data-events="actionEvents">文章</th>
        <th data-field="Attachmentlink" data-formatter="setAttachment" data-events="actionEvents">附件</th>
        <th data-field="Pdflink" data-formatter="setPdf" data-events="actionEvents">PDF</th>
        <th data-field="Created" data-formatter="localDateFormatter">建立时间</th>
        <!-- <th data-field="Created" data-formatter="actionFormatter" events="actionEvents">操作</th> -->
      </tr>
    </thead>
</table>
<!-- <table id="tabletest" 
        data-search="true"
        data-show-refresh="true"
        data-show-toggle="true"
        data-show-columns="true"
        data-query-params="queryParams"
        data-sort-name="ProjectName"
        data-sort-order="desc"
        data-page-size="5"
        data-page-list="[5,20, 50, 100, All]"
        data-unique-id="id"
        data-pagination="true"
        data-side-pagination="client"
        data-single-select="true"
        >
    <thead>        
      <tr>
        <th data-width="10" data-radio="true"></th>
        <th data-formatter="index1">#</th>
        <th data-field="Code" data-formatter="setCodetest">编号</th>
        <th data-field="Title">名称</th>
        <th data-field="Label" data-formatter="setLable">关键字</th>
        <th data-field="Principal">设计</th>
        <th data-field="Content">文章</th>
        <th data-field="Attachment">附件</th>
        <th data-field="PDF">PDF</th>
        <th data-field="Created" data-formatter="localDateFormatter">建立时间</th>
      </tr>
    </thead>
</table> -->
<!-- <div class="gridview2"></div> -->

<script type="text/javascript">
  /*数据json使用json数据要删除data-toggle="table"*/
  // var json =
     // 保留***[{"Id":"1","Code":[这个数组也行"SL0001-510-08","SL0001-510-08"],"Title":"水利枢纽布置图","Label":"水电站","Principal":"秦晓川","Product":"8","Created":"2016-11-26"},
      // [{"Id":"2","Code":"SL0002-530-10,SL0002-530-10","Title":"电力布置图","Label":"水电站","Principal":"秦晓川","Product":"8","Created":"2016-11-26"},
      // {"Id":"3","Code":"SL0003-650-20","Title":"市政布置图","Label":"水电站","Principal":"秦晓川","Product":"8","Created":"2016-11-26"},
      // {"Id":"4","Code":"SL0004-750-60","Title":"建筑平面图","Label":"水电站","Principal":"秦晓川","Product":"8","Created":"2016-11-26"},
      // {"Id":"5","Code":"SL0005-870-20","Title":"交通纵面图","Label":"水电站","Principal":"秦晓川","Product":"8","Created":"2016-11-26"},
      // {"Id":"6","Code":"SL0006-230-25","Title":"境外鸟瞰图","Label":"水电站","Principal":"秦晓川","Product":"8","Created":"2016-11-26"}];
        /*初始化table数据*/
        // $(function(){
        //     $("#tabletest").bootstrapTable({
        //         data:json
                // onClickRow: function (row, $element) {
                //   alert( "选择了行Id为: " + row.Id );
                //   rowid=row.Id//全局变量
                //   $('#table1').bootstrapTable('refresh', {url:'/admincategory?pid='+row.Id});
                // }
        //     });
        // });
  function index1(value,row,index){
  // alert( "Data Loaded: " + index );
            return index+1
  }

  function localDateFormatter(value) {
    return moment(value, 'YYYY-MM-DD').format('YYYY-MM-DD');
  }
  function setCode(value,row,index){
    return "<a href='/project/product/attachment/"+row.Id+"'>" + value + "</a>";
  }
  function setLable(value,row,index){
    array=value.split(",")
    var labelarray = new Array() 
    for (i=0;i<array.length;i++)
    {
      labelarray[i]="<a href='/project/product/search/"+array[i]+"'>" + array[i] + "</a>";
    }
    return labelarray.join(",");
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
    array=value.split(",")
    var labelarray = new Array() 
    for (i=0;i<array.length;i++)
    {
      labelarray[i]="<a href='/project/product/attachment/"+array[i]+"'>" + array[i] + "</a>";
    }
    return labelarray.join(",");
  // }else{
    // return "<a href='/project/product/attachment/"+value+"'>" + value + "</a>";
  // }
    // var x   
    // for (x in array)
    // {
    // array[x];
    // }
          // <a href="/category?op=viewlabel&label=value">
          // <span class="label label-info"></span></a>
  }
  function setTitle(value,row,index){
    return "<a href='/project/product/"+row.Id+"'>" + value + "</a>";
  }
  function setArticle(value,row,index){
    // return '<a class="article" href="javascript:void(0)" title="article"><i class="fa fa-file-text-o"></i></a>';
    if (value.length==1){
      articleUrl= '<a href="/project/product/article/'+value[0].Id+'" title="查看" target="_blank"><i class="fa fa-file-text-o"></i></a>';
      return articleUrl;
    }else if(value.length==0){
                    
    }else if(value.length>1){
      articleUrl= "<a class='article' href='javascript:void(0)' title='查看文章列表'><i class='fa fa-list-ol'></i></a>";
      return articleUrl;
    }
  }
// var bb;
  function setAttachment(value,row,index){
    if (value.length==1){
      attachUrl= '<a href="'+value[0].Link+'/'+value[0].Title+'" title="下载" target="_blank"><i class="fa fa-paperclip"></i></a>';
      return attachUrl;
    }else if(value.length==0){
                    
    }else if(value.length>1){
      attachUrl= "<a class='attachment' href='javascript:void(0)' title='查看附件列表'><i class='fa fa-list-ol'></i></a>";
      return attachUrl;
    }
          // $.ajax({//这种同步加载行不通，会混乱。异步又无法传出返回值data
            // type:"get",//这里是否一定要用post？？？
            // url:"/project/product/attachment/"+row.Id,
            // data: {CatalogId:row.Id},
            // dataType:'json',
            // async:false,//必须加这个异步为否定，即同步，否则bb传不出去
            // success:function(data,status){//数据提交成功时返回数据
                // $.each(data,function(i,d){
                //     $("#cars").append('<option value="' + data[i].Username + '"></option>');
                  // });
          //         if (data.length==1){
          //           bb= '<a href="'+data[0].Link+'/'+data[0].Title+'" title="下载" target="_blank"><i class="fa fa-paperclip"></i></a>';
          //         }else if(data.length==0){
                    
          //         }else if(data.length>1){
          //           bb= '<a class="attachment" href="javascript:void(0)" title="查看附件列表"><i class="fa fa-list-ol"></i></a>';
          //         } 
          //   }
          // });
          // return bb; 
    // return '<a class="attachment" href="javascript:void(0)" title="attachment"><i class="fa fa-paperclip"></i></a>';
  }

  // var pdfUrl;
  function setPdf(value,row,index){
    // return value;
    // alert(value[0].Link);
    if (value.length==1){
      pdfUrl= '<a href="'+value[0].Link+'/'+value[0].Title+'" title="打开pdf" target="_blank"><i class="fa fa-file-pdf-o"></i></a>';
      return pdfUrl;
    }else if(value.length==0){
                    
    }else if(value.length>1){
      pdfUrl= "<a class='pdf' href='javascript:void(0)' title='查看pdf列表'><i class='fa fa-list-ol'></i></a>";
      return pdfUrl;
    } 
    // $.ajax({
            // type:"get",//这里是否一定要用post？？？
            // url:"/project/product/pdf/"+row.Id,
            // data: {CatalogId:row.Id},
            // dataType:'json',
            // async:false,//必须加这个异步为否定，即同步，否则bb传不出去
            // success:function(data,status){//数据提交成功时返回数据
                // $.each(data,function(i,d){
                //     $("#cars").append('<option value="' + data[i].Username + '"></option>');
                  // });
            //       if (data.length==1){
            //         pdfUrl= '<a href="'+data[0].Link+'/'+data[0].Title+'" title="下载" target="_blank"><i class="fa fa-file-pdf-o"></i></a>';
            //         return pdfUrl;
            //       }else if(data.length==0){
                    
            //       }else if(data.length>1){
            //         pdfUrl= "<a class='pdf' href='javascript:void(0)' title='查看pdf列表'><i class='fa fa-list-ol'></i></a>";
            //         return pdfUrl;
            //       } 
            // }
    // });
  }
  //最后面弹出文章列表中用的
  function setArticlecontent(value,row,index){
    articleUrl= '<a href="'+value+'" title="下载" target="_blank"><i class="fa fa-paperclip"></i></a>';
      return articleUrl;
  }
  //最后面弹出附件列表中用的
  function setAttachlink(value,row,index){
    attachUrl= '<a href="'+value+'" title="下载" target="_blank"><i class="fa fa-paperclip"></i></a>';
      return attachUrl;
  }
  //最后面弹出pdf列表中用的
  function setPdflink(value,row,index){
    pdfUrl= '<a href="'+value+'" title="下载" target="_blank"><i class="fa fa-file-pdf-o"></i></a>';
      return pdfUrl;
  }
  //这个没用
  // function actionFormatter(value, row, index) {
  //   return [
  //       '<a class="send" href="javascript:void(0)" title="提交">',
  //       '<i class="glyphicon glyphicon-step-forward"></i>',
  //       '</a>&nbsp;',
  //       '<a class="downsend" href="javascript:void(0)" title="退回">',
  //       '<i class="glyphicon glyphicon-step-backward"></i>',
  //       '</a>&nbsp;',
  //       '<a class="remove" href="javascript:void(0)" title="删除">',
  //       '<i id="delete" class="glyphicon glyphicon-remove"></i>',
  //       '</a>'
  //   ].join('');
  // }
// '<a class="edit ml10" href="javascript:void(0)" title="退回">','<i class="glyphicon glyphicon-edit"></i>','</a>'
  window.actionEvents = {
    'click .article': function (e, value, row, index) {
        // if(confirm("暂时不支持查看文章！土豪~")){
        // }
      $('#articles').bootstrapTable('refresh', {url:'/project/product/articles/'+row.Id});
      $('#modalarticle').modal({
        show:true,
        backdrop:'static'
      }); 
    },
    'click .attachment': function (e, value, row, index) {
      $('#attachs').bootstrapTable('refresh', {url:'/project/product/attachment/'+row.Id});
      $('#modalattach').modal({
        show:true,
        backdrop:'static'
      });
    },
    'click .pdf': function (e, value, row, index) {
      // alert("商品名称：" +row.Id);
        // if(confirm("确定提交该行吗？")){
        //后台获取附件——如果大于一个，则打开模态框，否则直接打开/下载
        // var removeline=$(this).parents("tr")
          // $.ajax({
          // type:"get",//这里是否一定要用post？？？
          // url:"/project/product/pdf/"+row.Id,
          // data: {CatalogId:row.Id},
            // success:function(data,status){//数据提交成功时返回数据
              $('#pdfs').bootstrapTable('refresh', {url:'/project/product/pdf/'+row.Id});
              $('#modalpdf').modal({
                show:true,
                backdrop:'static'
              }); 
              // alert("提交“"+data+"”成功！(status:"+status+".)");
            // }
          // });  
        // }
    },
  };
  // 改变点击行颜色
  // $(function(){
     // $("#table").bootstrapTable('destroy').bootstrapTable({
     //     columns:columns,
     //     data:json
     // });
     // $("#table0").on("click-row.bs.table",function(e,row,ele){
         // $(".info").removeClass("info");
         // $(ele).addClass("info");
         // rowid=row.Id//全局变量
         // $('#table1').bootstrapTable('refresh', {url:'/admin/category/'+row.Id});
     // });
     // $("#get").click(function(){
     //     alert("商品名称：" + getContent().TuanGouName);
     // })
  // });
  // $(function () {
  //     var $result = $('#eventsResult');
  //     var selectRow=$('#table').bootstrapTable('getSelections');
  
  //     $('#table').on('all.bs.table', function (e, name, args) {
  //         console.log('Event:', name, ', data:', args);
  //     })
  //     .on('click-row.bs.table', function (e, row, $element) {
  //       alert("选择！"+row.Id);
  //       if (selectRow.length<1){
  //         selectRow=$('#table').bootstrapTable('getSelections');
  //         alert("请选择"+selectRow.length);
  //         // return;
  //         }
  //         $result.text('Event: click-row.bs.table');
  //     })
  //     .on('dbl-click-row.bs.table', function (e, row, $element) {
  //         $result.text('Event: dbl-click-row.bs.table');
  //     })
  //     .on('sort.bs.table', function (e, name, order) {
  //         $result.text('Event: sort.bs.table');
  //     })
  //     .on('check.bs.table', function (e, row) {
  //         $result.text('Event: check.bs.table');
  //     })
  //     .on('uncheck.bs.table', function (e, row) {
  //         $result.text('Event: uncheck.bs.table');
  //     })
  //     .on('check-all.bs.table', function (e) {
  //         $result.text('Event: check-all.bs.table');
  //     })
  //     .on('uncheck-all.bs.table', function (e) {
  //         $result.text('Event: uncheck-all.bs.table');
  //     })
  //     .on('load-success.bs.table', function (e, data) {
  //         $result.text('Event: load-success.bs.table');
  //     })
  //     .on('load-error.bs.table', function (e, status) {
  //         $result.text('Event: load-error.bs.table');
  //     })
  //     .on('column-switch.bs.table', function (e, field, checked) {
  //         $result.text('Event: column-switch.bs.table');
  //     })
  //     .on('page-change.bs.table', function (e, number, size) {
  //         $result.text('Event: page-change.bs.table');
  //     })
  //     .on('search.bs.table', function (e, text) {
  //         $result.text('Event: search.bs.table');
  //     });
  // });


  // $(document).ready(function() {
  // $("#addButton").click(function() {
    
  // var selectRow=$('#table').bootstrapTable('getSelections');  
  // if (selectRow.length<1){
  // selectRow=$('#table').bootstrapTable('getSelections');
  // alert("请选择"+selectRow.length);
  // return;
  // }
        // $('#modalTable').modal({
        // show:true,
        // backdrop:'static'
        // });
    // })
  // })

  $(document).ready(function() {
    // 批量上传
    $("#addButton").click(function() {
      if ({{.role}}!=1){
        alert("权限不够！"+{{.role}});
        return;
      }
      $("input#pid").remove();
      var th1="<input id='pid' type='hidden' name='pid' value='" +{{.Id}}+"'/>"
        $(".modal-body").append(th1);

        $('#modalTable').modal({
        show:true,
        backdrop:'static'
        });
    })

    var uploader;
    $('#modalTable').on('shown.bs.modal',function(){
      // var $ = jQuery,
      $list = $('#thelist'),
      $btn = $('#ctlBtn'),
      state = 'pending',
      // uploader;
      uploader = WebUploader.create({
        // 不压缩image
        resize: false,
        // swf文件路径
        swf: '/static/fex-team-webuploader/dist/Uploader.swf',
        // 文件接收服务端。
        server: '/project/product/addattachment',
        // 选择文件的按钮。可选。
        // 内部根据当前运行是创建，可能是input元素，也可能是flash.
        pick: '#picker'
      });
      // 当有文件添加进来的时候
      uploader.on( 'fileQueued', function( file ) {
      $list.append( '<div id="' + file.id + '" class="item">' +
            '<h4 class="info">' + file.name + '</h4>' +
            '<p class="state">等待上传...</p>' +
        '</div>' );
      }); 

      //传递参数——成果id
      uploader.on( 'startUpload', function() {//uploadBeforeSend——这个居然不行？
      // if (prodlabel){
        var pid = $('#pid').val();
        var prodlabel = $('#prodlabel').val();
        var prodprincipal = $('#prodprincipal').val();
        // var html = ue.getContent();
        // alert(html);
        uploader.option('formData', {
          "pid":pid,
          "prodlabel":prodlabel,
          "prodprincipal":prodprincipal
          // 'content':html,
          // {'tnumber':a,'title':b,'categoryid':c,'category':d,'content':e}
        });
        // }else{
        //     alert("名称不能为空");
        // }         
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
        $('#table0').bootstrapTable('refresh', {url:'/project/products/'+{{.Id}}});
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

    $('#modalTable').on('hide.bs.modal',function(){
      $list.text("");
      uploader.destroy();//销毁uploader
    })

    // 多附件模式
    $("#addButton1").click(function() {
      if ({{.role}}!=1){
        alert("权限不够！");
        return;
      }
      $("input#pid").remove();
      var th1="<input id='pid' type='hidden' name='pid' value='" +{{.Id}}+"'/>"
        $(".modal-body").append(th1);
        $('#modalTable1').modal({
          show:true,
          backdrop:'static'
        });
    })

    var uploader;
    $('#modalTable1').on('shown.bs.modal',function(){
      // var $ = jQuery,
      $list = $('#thelist1'),
      $btn = $('#ctlBtn1'),
      state = 'pending',
      // uploader;
      uploader = WebUploader.create({
        // 不压缩image
        resize: false,
        // swf文件路径
        swf: '/static/fex-team-webuploader/dist/Uploader.swf',
        // 文件接收服务端。
        server: '/project/product/addattachment2',
        // 选择文件的按钮。可选。
        // 内部根据当前运行是创建，可能是input元素，也可能是flash.
        pick: '#picker1'
      });
      // 当有文件添加进来的时候
      uploader.on( 'fileQueued', function( file ) {
        $list.append( '<div id="' + file.id + '" class="item">' +
              '<h4 class="info">' + file.name + '</h4>' +
              '<p class="state">等待上传...</p>' +
          '</div>' );
      }); 

      //传递参数——成果id
      uploader.on( 'startUpload', function() {//uploadBeforeSend——这个居然不行？
        var pid = $('#pid').val();
        var prodcode = $('#prodcode').val();
        var prodname = $('#prodname').val();
        var prodlabel = $('#prodlabel').val();
        var prodprincipal = $('#prodprincipal').val();
        // var html = ue.getContent();
        // alert(html);
        uploader.option('formData', {
          "pid":pid,
          "prodcode":prodcode,
          "prodname":prodname,
          "prodlabel":prodlabel,
          "prodprincipal":prodprincipal
          // 'content':html,
          // {'tnumber':a,'title':b,'categoryid':c,'category':d,'content':e}
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
          $('#table0').bootstrapTable('refresh', {url:'/project/products/'+{{.Id}}});
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
        var prodcode = $('#prodcode').val();
        var prodname = $('#prodname').val();
        if (prodcode&&prodname){
          if ( state === 'uploading' ) {
              uploader.stop();
          } else {
              uploader.upload();
          }
        }else{
            alert("编号和名称不能为空"+prodcode+prodname);
        }
      });
    })

    $('#modalTable1').on('hide.bs.modal',function(){
      $list.text("");
      uploader.destroy();//销毁uploader
    })

    //****添加文章
    $("#addButton2").click(function() {
      if ({{.role}}!=1){
        alert("权限不够！");
        return;
      }
      $("input#pid").remove();
      var th1="<input id='pid' type='hidden' name='pid' value='" +{{.Id}}+"'/>"
        $(".modal-body").append(th1);

        $('#modalTable2').modal({
        show:true,
        backdrop:'static'
        });
    })


    // 编辑成果信息
    $("#editorProdButton").click(function() {
      if ({{.role}}!=1){
        alert("权限不够！");
        return;
      }
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

      $('#modalProdEditor').modal({
      show:true,
      backdrop:'static'
      });
    })

    // 编辑成果附件——删除附件、文章或追加附件
    var selectrowid;
    $("#editorAttachButton").click(function() {
      if ({{.role}}!=1){
        alert("权限不够！");
        return;
      }
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
      $('#attachments').bootstrapTable('refresh', {url:'/project/product/allattachments/'+selectRow[0].Id});//取得所有附件列表和文章列表
      $('#modalAttachEditor').modal({
      show:true,
      backdrop:'static'
      });
    })

    var uploader;
    $('#modalAttachEditor').on('shown.bs.modal',function(){
      // var $ = jQuery,
      $list = $('#thelist2'),
      $btn = $('#ctlBtn2'),
      state = 'pending',
      // uploader;
      uploader = WebUploader.create({
        // 不压缩image
        resize: false,
        // swf文件路径
        swf: '/static/fex-team-webuploader/dist/Uploader.swf',
        // 文件接收服务端。
        server: '/project/product/updateattachment',
        // 选择文件的按钮。可选。
        // 内部根据当前运行是创建，可能是input元素，也可能是flash.
        pick: '#picker2'
      });
      // 当有文件添加进来的时候
      uploader.on( 'fileQueued', function( file ) {
        $list.append( '<div id="' + file.id + '" class="item">' +
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
          $('#attachments').bootstrapTable('refresh', {url:'/project/product/allattachments/'+selectrowid});
          $('#table0').bootstrapTable('refresh', {url:'/project/products/'+{{.Id}}});
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
      $list.text("");
      uploader.destroy();//销毁uploader
    })
    // 删除成果
    $("#deleteButton").click(function() {
      if ({{.role}}!=1){
        alert("权限不够！");
        return;
      }
      var selectRow=$('#table0').bootstrapTable('getSelections');
      if (selectRow.length<=0) {
        alert("请先勾选成果！");
        return false;
      }
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
          url:"/project/product/deleteproduct",
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
    })

  })

   /*数据json*/
  // var json1 = [{"Id":"1","Title":"规划","Code":"A","Grade":"1"},
  //             {"Id":"7","Title":"可研","Code":"B","Grade":"1"},
  //             {"Id":"2","Title":"报告","Code":"B","Grade":"2"},
  //             {"Id":"3","Title":"图纸","Code":"T","Grade":"2"},
  //             {"Id":"4","Title":"水工","Code":"5","Grade":"3"},
  //             {"Id":"5","Title":"机电","Code":"6","Grade":"3"},
  //             {"Id":"6","Title":"施工","Code":"7","Grade":"3"}];
  /*初始化table数据*/
  // $(function(){
  //     $("#table1").bootstrapTable({
  //         data:json1
  //     });
  // });


  //模态框可拖曳—要引入ui-jquery.js
  // $("#modalTable").draggable({
  //   handle:".modal-header",
  //   cusor:"move",
  //   refreshPositions:false
  // });
  // 来自群，保留，批量
  // var rows= $('#account-table').bootstrapTable('getSelections');
  //       if(rows.length==0) {
  //           layer.alert('请您选择要删除的子账号！', {
  //               title:'提示信息',
  //               closeBtn: 0,
  //               icon: 0,
  //               skin: 'layui-layer-lan',
  //               shift:0 //动画类型
  //           });
  //           return false;
  //       }
  //           var ids="";
  //           for(var i=0;i<rows.length;i++){
  //               if(i==0){
  //                   ids=rows[i].frontUserId;
  //               }else{
  //                   ids=ids+","+rows[i].frontUserId;
  //               }
  //           }
</script>
  <!-- 批量上传 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalTable">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">批量添加成果</h3>
            <label>**请选择标准命名的电子文件上传**</label>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">关键字</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodlabel" name="prodlabel"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">设计</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodprincipal" name="prodprincipal"></div>
              </div>
              <!-- <div style="display: inline-block;">
                <span id="filePicker" onclick="create()">上传</span>
                <span id="responseText" style="display: inline-block;"></span>
              </div> -->
              <!--SWF在初始化的时候指定，在后面将展示-->
              <div id="uploader" style="position:relative;text-align: center;">
              <!--用来存放文件信息-->
                <div id="thelist"></div>
                <div class="btns">
                  <div id="picker">选择文件</div>
                  <button id="ctlBtn" class="btn btn-default">开始上传</button>
                </div>
              </div>
              <!-- <div class="form-group">
                <label class="col-sm-3 control-label">标签</label>
                <div class="col-sm-7">
                  <input type="number" class="form-control digits" name="label" maxlength="20" placeholder="至多20个字符" required></div>
              </div> -->
              <!-- <div class="form-group must">
                <label class="col-sm-3 control-label">负责人</label>
                  <div class="col-sm-7">
                    <input type="password" class="form-control" name="password" id="password" maxlength="32" placeholder="至多32个字符" required></div>
              </div> -->
            <!-- <div class="form-group must">
              <label class="col-sm-3 control-label">确认密码</label>
              <div class="col-sm-7">
                <input type="password" class="form-control equalto" name="password2" maxlength="32" placeholder="至多32个字符" required data-rule-equalto="#password" data-msg-equalto="密码不一致"></div>
            </div> -->
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
            <!-- <button type="button" class="btn btn-primary" onclick="save()">保存</button> -->
          </div>
        </div>
      </div>
    </div>
  </div>

<!-- 多附件 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalTable1">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">添加成果——多附件模式</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">编号</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="prodcode" name="prodcode"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">名称</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodname" name="prodname"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">关键字</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodlabel1" name="prodlabel1"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">设计</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodprincipal1" name="prodprincipal1"></div>
              </div>
              <!--SWF在初始化的时候指定，在后面将展示-->
              <div id="uploader1" style="position:relative;text-align: center;">
              <!--用来存放文件信息-->
                <div id="thelist1"></div>
                <div class="btns1">
                  <div id="picker1">选择文件</div>
                  <button id="ctlBtn1" class="btn btn-default">开始上传</button>
                </div>
              </div>
            </div>
            <!-- <label>文件简介:</label>
              <div>
                <script id="container" type="text/plain" style="height:200px;"></script>width:1024px;
              </div> -->
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
            <!-- <button type="button" class="btn btn-primary" onclick="save1()">保存</button> -->
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- 添加文章 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalTable2">
      <div class="modal-dialog" style="width: 100%">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">添加文章</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">编号</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="prodcode1" name="prodcode1"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">标题</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodname1" name="prodname1"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">副标题</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="subtext1" name="subtext1"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">关键字</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodlabel2" name="prodlabel2"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">设计</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodprincipal2" name="prodprincipal2"></div>
              </div>
            </div>
            <label>文章正文:</label>
              <div>
                <script id="container" type="text/plain" style="height:200px;"></script><!-- width:1024px; -->
              </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
            <button type="button" class="btn btn-primary" onclick="save2()">保存</button>
          </div>
        </div>
      </div>
    </div>
  </div>
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
                <table id="articles"
                      data-toggle="table"
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
                      <th data-field="Content" data-formatter="setArticlecontent">查看</th>
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
                <table id="attachs"
                      data-toggle="table"
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
                <table id="pdfs"
                      data-toggle="table"
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
  <!-- 编辑成果名称等信息 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalProdEditor">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
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
                  <input type="tel" class="form-control" id="prodlabel3" name="prodlabel3"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">设计</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodprincipal3" name="prodprincipal3"></div>
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
  <!-- 编辑成果附件 删除附件、文章或追加附件-->
  <div class="form-horizontal">
    <div class="modal fade" id="modalAttachEditor">
      <div class="modal-dialog">
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
//实例化编辑器
    var ue = UE.getEditor('container', {
      autoHeightEnabled: true,
      autoFloatEnabled: true
    });
  /* 2.传入参数表,添加到已有参数表里 通过携带参数，实现不同的页面使用不同controllers*/
    ue.ready(function () {
        ue.addListener('focus', function () {//startUpload start-upload startUpload beforeExecCommand是在插入图片之前触发
            var pid = $('#pid').val();
            // var html = ue.getContent();
            ue.execCommand('serverparam', {
              "pid":pid 
            });
        });
    });

  //添加文章
  function save2(){
    // var radio =$("input[type='radio']:checked").val();
    var projectid = $('#pid').val();
    var prodcode = $('#prodcode1').val();
    var prodname = $('#prodname1').val();
    var subtext = $('#subtext1').val();
    var prodprincipal = $('#prodprincipal2').val();
    var prodlabel = $('#prodlabel2').val();
    var html = ue.getContent();
    // $('#myModal').on('hide.bs.modal', function () {  
    if (prodname&&prodcode){  
      $.ajax({
        type:"post",
        url:"/project/product/addarticle",
        data: {pid:projectid,code:prodcode,title:prodname,subtext:subtext,label:prodlabel,content:html,principal:prodprincipal},//父级id
        success:function(data,status){
          alert("添加“"+data+"”成功！(status:"+status+".)");
          $('#modalTable2').modal('hide');
          $('#table0').bootstrapTable('refresh', {url:'/project/products/'+{{.Id}}});
        },
        
      });
    }else{
      alert("请填写编号和名称！");
      return;
    }
  }
  // 编辑成果信息
  function updateprod(){
    // var radio =$("input[type='radio']:checked").val();
    var projectid = $('#cid').val();
    var prodcode = $('#prodcode3').val();
    var prodname = $('#prodname3').val();
    var prodlabel = $('#prodlabel3').val();
    var prodprincipal = $('#prodprincipal3').val();
 
    if (prodname&&prodcode){  
      $.ajax({
        type:"post",
        url:"/project/product/updateproduct",
        data: {pid:projectid,code:prodcode,title:prodname,label:prodlabel,principal:prodprincipal},//父级id
        success:function(data,status){
          alert("添加“"+data+"”成功！(status:"+status+".)");
          $('#modalProdEditor').modal('hide');
          $('#table0').bootstrapTable('refresh', {url:'/project/products/'+{{.Id}}});
        },
        
      });
    }else{
      alert("请填写编号和名称！");
      return;
    }
  }
  // 删除附件
    $("#deleteAttachButton").click(function() {
      if ({{.role}}!=1){
        alert("权限不够！");
        return;
      }
      var selectRow=$('#attachments').bootstrapTable('getSelections');
      if (selectRow.length<=0) {
        alert("请先勾选成果！");
        return false;
      }
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
          url:"/project/product/deleteattachment",
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

</script>

</body>
</html>
<!-- 设代财务登记的列表页面-->
<!DOCTYPE html>
<head>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <!-- <script src="/static/js/bootstrap-treeview.js"></script> -->
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  <script src="/static/js/tableExport.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/webuploader.css">
  <script type="text/javascript" src="/static/js/webuploader.min.js"></script>
  <script type="text/javascript" src="/static/js/jquery-ui.min.js"></script>
  
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

  <style type="text/css">
    #modalDialog .modal-header {cursor: move;}
    #modalDialog1 .modal-header {cursor: move;}
    #modalDialog2 .modal-header {cursor: move;}
    #modalDialog3 .modal-header {cursor: move;}
    #modalDialog4 .modal-header {cursor: move;}
    #modalDialog5 .modal-header {cursor: move;}
    #modalDialog6 .modal-header {cursor: move;}
    #modalDialog7 .modal-header {cursor: move;}
    #modalDialog8 .modal-header {cursor: move;}
    #modalDialog9 .modal-header {cursor: move;}
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
      div#modalTable2 {/*.modal .fade .in*/
        z-index: 3;
      }
      /*.form-horizontal .control-label{
        padding-left:10px; 
      }
      .form-horizontal .form-group{
        float: left;
        width: 50%;
      }*/    
  </style>
</head>

<body>

<div class="col-lg-12">
  <h3>财务登记列表</h3>
<!-- <a href="/htmltodoc" class="btn btn-warning btn-xs">导出</a> -->
<button type="button" class="btn btn-danger btn-xs" id="exporthtml">导出</button>

<table id="table0"></table>

<script type="text/javascript">
  $(function () {
    // 初始化【未接受】工作流表格
    $("#table0").bootstrapTable({
      url : '/v1/wx/getwxfinance2/170592',
      method: 'get',
      search:'true',
      classes: "table table-striped", //这里设置表格样式
      showRefresh:'true',
      showToggle:'true',
      showColumns:'true',
      toolbar:'#toolbar1',
      pagination: 'true',
      sidePagination: "server",
      queryParamsType:'',
      //请求服务器数据时，你可以通过重写参数的方式添加一些额外的参数，例如 toolbar 中的参数 如果 queryParamsType = 'limit' ,返回参数必须包含
      // limit, offset, search, sort, order 否则, 需要包含: 
      // pageSize, pageNumber, searchText, sortName, sortOrder. 
      // 返回false将会终止请求。
      pageSize: 15,
      pageNumber: 1,
      pageList: [15, 50, 100],
      uniqueId:"id",
      // singleSelect:"true",
      clickToSelect:"true",
      showExport:"true",
      queryParams:function queryParams(params) {   //设置查询参数
        var param = {
          limit: params.pageSize,   //每页多少条数据
          page: params.pageNumber, // 页码
          skey:$(".search .form-control").val()
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
      columns: [
        {
          title: '选择',
          // radio: 'true',
          checkbox:'true',
          width: '10',
          align:"center",
          valign:"middle"
        },
        {
          // field: 'Number',
          title: '序号',
          formatter:function(value,row,index){
            return index+1
          },
          width: '20',
          align:"center",
          valign:"middle"
        },
        {
          field: 'Financedate',
          title: '日期',
          width: '100',
          formatter:setCode,
          halign:"center",
          align:"center",
          valign:"middle"
        },
        {
          field: 'amount',
          title: '数额',
          formatter:setTitle,
          halign:"center",
          align:"right",
          valign:"middle"
        },
        {
          field: 'consider',
          title: '是否统筹',
          formatter:setConsider,
          halign:"center",
          align:"right",
          valign:"middle"
        },
        {
          field: 'Nickname',
          title: '登记',
          width: '100',
          align:"center",
          valign:"middle"
        },
        {
          field: 'Created',
          title: '建立时间',
          width: '200',
          formatter:localDateFormatter,
          visible:"false",
          align:"center",
          valign:"middle"
        }
      ]
    });
  });

  // 导出word
  $("#exporthtml").click(function() {
    //获取page-number active
    //获取page-size
    // var lis = document.getElementsByClassName('page-size');
    var page=$("li.page-number.active").text()
    // alert(page)
    var limit=$("span.page-size:first").text()
    // alert(limit)

    $.ajax({
      type:"get",
      url:"/htmltodoc",
      data: {page:page,limit:limit},
      success:function(data,status){
        alert("导出成功！"+data.filename);
        var btn="<a href='/static/" + data.filename + "' class='btn btn-danger btn-xs' id='testBtn' value='下载'>下载</a>";
        $("#exporthtml").after(btn);
        // addBtnEvent("testBtn");
      }
      //添加下载按钮
    }); 
  })

  // function addBtnEvent(id){
  //   $("#"+id).bind("click",function(){
  //     alert("Test");
  //   });
  // }

  function index1(value,row,index){
    return index+1
  }

  function localDateFormatter(value) {
    return moment(value, 'YYYY-MM-DD h:mm:ss a').format('YYYY-MM-DD h:mm:ss a');
  }

  function RelevFormatter(value) {
    if (value){
      if (value.length==1){//'<a href="/project/product/article/'
        var array=value[0].Relevancy.split(",")
        var relevarray = new Array() 
        for (i=0;i<array.length;i++)
        {
          // relevarray[i]=array[i];
            relevarray[i]='<a href="javascript:gototree(' + value[i].ProjectId + ')" title="查看" target="_blank">'+value[i].Relevancy+'</a>';
        }
        return relevarray.join(",");
        // articleUrl= '<a href="'+value[0].Link+'/'+value[0].Id+'" title="查看" target="_blank"><i class="fa fa-file-text-o"></i></a>';
        // return articleUrl;
      }else if(value.length==0){
                    
      }else if(value.length>1){
        var relevarray = new Array()
        for (i=0;i<value.length;i++)
          {
            // relevarray[i]=value[i].Relevancy;
            relevarray[i]='<a href="javascript:gototree(' + value[i].ProjectId + ')" title="查看" target="_blank">'+value[i].Relevancy+'</a>';
          }
        return relevarray.join(",");
        // articleUrl= "<a class='article' href='javascript:void(0)' title='查看文章列表'><i class='fa fa-list-ol'></i></a>";
        // return articleUrl;
      }
    }
  }

  // 关联成果跳转到对应的树状目录下
  function gototree(e){
    // $(window.parent.document).find("input:radio").attr("checked","true");
    // $('#objId', parent.document);
    // 格式：$("#父页面元素id" , parent.document);
    parent.gototree(e); // pClick 为父页面 js 方法
    // window.parent.document.getElementById("父窗口元素ID");
    // window.parent.document.getElementById("iframepage").src="/project/{{.Id}}/"+e;
    // var findCheckableNodess = function() {
    //   return $('#tree',parent.document).treeview('findNodes', [e, 'id']);
    // }; 
    // var checkableNodes = findCheckableNodess();
    //   $('#tree',parent.document).treeview('toggleNodeSelected', [ checkableNodes, { silent: true } ]);
    //   $('#tree',parent.document).treeview('toggleNodeExpanded', [ checkableNodes, { silent: true } ]);
    //   $('#tree',parent.document).treeview('revealNode', [ checkableNodes, { silent: true } ]);
  }

  function setCode(value,row,index){
    return "<a href='/getwxfinance2/"+row.id+"'>" + value + "</a>";
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
    return "<a href='/getwxfinance2/"+row.id+"'>" + value + "</a>";
  }

  function setConsider(value,row,index){
    if (value) {
      return "是";
    }else{
      return "否";
    }
  }

  function setDocState(value,row,index){
    // if (value.Name){
      // return "<a href='/cms/#/flow/documentdetail2?docid="+ row.ProdDoc.DocumentId +"&dtid="+row.ProdDoc.DocTypeId+"'target='_blank'>" + value.Name + "</a>";
    // }
  }

</script>

</div>

  <!-- <script type="text/javascript" src="/static/froala/js/jquery.min.1.11.0.js"></script> -->
  <script type="text/javascript" src="/static/froala/js/codemirror.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/xml.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/froala_editor.min.js" ></script>
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
$(function (){
  //超大屏幕'fullscreen',
  var toolbarButtons   = [ 'bold', 'italic', 'underline', 'strikeThrough', 'subscript', 'superscript', 'fontFamily', 'fontSize', '|', 'color', 'emoticons', 'inlineStyle', 'paragraphStyle', '|', 'paragraphFormat', 'align', 'formatOL', 'formatUL', 'outdent', 'indent', 'quote', 'insertHR', '-', 'insertLink', 'insertImage', 'insertVideo', 'insertFile', 'insertTable', 'undo', 'redo', 'clearFormatting', 'selectAll', 'html','help'];
  //大屏幕
  var toolbarButtonsMD = ['bold', 'italic', 'underline','strikeThrough', 'subscript', 'superscript', '|','fontFamily', 'fontSize', 'color','inlineStyle', 'paragraphStyle','|', 'paragraphFormat', 'align', 'formatOL', 'formatUL', 'outdent', 'indent', '-', 'quote', 'insertLink', 'insertImage', 'insertVideo', 'insertFile', 'insertTable', '|','specialCharacters','insertHR','undo', 'redo', 'clearFormatting','|','html','help'];
  //小屏幕'fullscreen',
  var toolbarButtonsSM = [ 'bold', 'italic', 'underline', 'fontFamily', 'fontSize', 'insertLink', 'insertImage', 'insertTable', 'undo', 'redo'];
  //手机
  var toolbarButtonsXS = ['bold', 'italic', 'fontFamily', 'fontSize', 'undo', 'redo'];
  var pid = $('#pid').val();
  //编辑器初始化并赋值 
  $('#edit').froalaEditor({
    placeholderText: '请输入内容',
    charCounterCount       : true,//默认
    // charCounterMax         : -1,//默认
    saveInterval            : 0,//不自动保存，默认10000
    // theme                    : "red",
    height                   : "300px",
    toolbarBottom           : false,//默认
    toolbarButtonsMD        : toolbarButtons,//toolbarButtonsMD,
    toolbarButtonsSM        : toolbarButtonsMD,//toolbarButtonsSM,
    toolbarButtonsXS        : toolbarButtonsXS,
    toolbarInline           : false,//true选中设置样式,默认false
    imageUploadMethod       : 'POST',
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
    saveParams: { postId: '1'},
    spellcheck: false,
    imageUploadURL: '/uploadimg',//上传到本地服务器
    imageUploadParams: {pid: '{{.Id}}'},
    imageDeleteURL: 'lib/delete_image.php',//删除图片
    imagesLoadURL: 'lib/load_images.php',//管理图片
    videoUploadURL:'/uploadvideo',
    videoUploadParams: {pid: '{{.Id}}'},
    fileUploadURL: '/uploadimg',
    fileUploadParams: {pid: '{{.Id}}'},
    enter: $.FroalaEditor.ENTER_BR,
    language: 'zh_cn',
    // toolbarButtons: ['bold', 'italic', 'underline', 'paragraphFormat', 'align','color','fontSize','insertImage','insertTable','undo', 'redo']
  });
})

  //添加文章
  function save2(){
    // var radio =$("input[type='radio']:checked").val();
    var projectid = $('#pid').val();
    var prodcode = $('#prodcode1').val();
    var prodname = $('#prodname1').val();
    var subtext = $('#subtext1').val();
    var prodprincipal = $('#prodprincipal2').val();
    var prodlabel = $('#prodlabel2').val();
    var relevancy = $('#relevancy2').val();
    var html = $('div#edit').froalaEditor('html.get');//$('#edit')[0].childNodes[1].innerHTML;
    // $('#myModal').on('hide.bs.modal', function () {  
    if (prodname&&prodcode){  
      $.ajax({
        type:"post",
        url:"/project/product/addarticle",
        data: {pid:projectid,code:prodcode,title:prodname,subtext:subtext,label:prodlabel,content:html,principal:prodprincipal,relevancy:relevancy},//父级id
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

  //添加流程flow
  function saveFlowDoc(){
    var doctype = $('#doctype').val();
    var accesscontext = $('#accesscontext').val();
    var group = $('#group').val();
    var docdata = $('#flowdata_attachment').val();
    var docname = $('#flowdata_docname').val();
    if (doctype&&group){  
      $.ajax({
        type:"post",
        url:"/v1/admin/flowdoc",
        data: {dtid:doctype,acid:accesscontext,gid:group,docname:docname,docdata:docdata},
        success:function(data,status){
          alert("添加“"+data.err+"”成功！(status:"+status+".)");
          $('#modalFlow').modal('hide');
        },
      });
    }else{
      alert("请填写编号和名称！");
      return;
    }
  }

  // 编辑成果信息——这个没用到
  function updateprod(){
    // var radio =$("input[type='radio']:checked").val();
    var projectid = $('#cid').val();
    var prodcode = $('#prodcode3').val();
    var prodname = $('#prodname3').val();
    var prodlabel = $('#prodlabel3').val();
    var prodprincipal = $('#prodprincipal3').val();
    var relevancy = $('#relevancy3').val();
    if (prodname&&prodcode){  
      $.ajax({
        type:"post",
        url:"/project/product/updateproduct",
        data: {pid:projectid,code:prodcode,title:prodname,label:prodlabel,principal:prodprincipal,relevancy:relevancy},//父级id
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
    // $(document).ready(function() {
    //   $('#table0').bootstrapTable({
    //     // onLoadSuccess: function(){
    //       onPostBody: function(){
    //       alert("加载成功");
    //     $('#table0').bootstrapTable('append', randomData());
    //     $('#table0').bootstrapTable('scrollTo', 'bottom'); 
    //          return false;
    //       }
    //     });
    // });
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
    function station_select1(){ 
      if(box1.checked){ 
        document.getElementById("relevancy1").disabled=false; 
      } else{ 
        document.getElementById("relevancy1").disabled=true; 
      } 
    }

    function station_select2(){ 
      if(box2.checked){ 
        document.getElementById("relevancy2").disabled=false; 
      } else{ 
        document.getElementById("relevancy2").disabled=true; 
      } 
    }

    $(document).ready(function(){
      $("#modalDialog").draggable({ handle: ".modal-header" });//为模态对话框添加拖拽
      $("#modalDialog1").draggable({ handle: ".modal-header" });
      $("#modalDialog2").draggable({ handle: ".modal-header" });
      $("#modalDialog3").draggable({ handle: ".modal-header" });
      $("#modalDialog4").draggable({ handle: ".modal-header" });
      $("#modalDialog5").draggable({ handle: ".modal-header" });
      $("#modalDialog6").draggable({ handle: ".modal-header" });
      $("#modalDialog7").draggable({ handle: ".modal-header" });
      $("#modalDialog8").draggable({ handle: ".modal-header" });
      $("#modalDialog9").draggable({ handle: ".modal-header" });
      // $("#modalNewDwg").draggable({ handle: ".modal-header" });
      // $("#modalFlow").draggable({ handle: ".modal-header" });
      $("#myModal").css("overflow", "hidden");//禁止模态对话框的半透明背景滚动
    })
</script>

</body>
</html>
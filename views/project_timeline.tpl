<!-- 测试时间轴 -->
<!doctype html>
<html lang="en-US">
<head>
  <meta charset="utf-8">
  <meta http-equiv="Content-Type" content="text/html">
  <title>{{.ProjectTile}}——大事记</title>
  <meta name="author" content="Jake Rocheleau">
  <link rel="stylesheet" type="text/css" media="all" href="/static/css/bootstrap.min.css">
  <link rel="stylesheet" type="text/css" media="all" href="/static/css/timeline.css">
	<link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <script type="text/javascript" src="/static/js/html2canvas.js"></script> 
  <!--需要注意的是,这里一定要等待js库和网页加载完毕后再执行函数-->  
  <!-- html2canvas()中,第一个参数是要截图的Dom对象，第二个参数时渲染完成后回调的canvas对象。-->  
  <script type="text/javascript">  

  function takeScreenshot() {
  var getPixelRatio = function(context) {
    var backingStore = context.backingStorePixelRatio ||
          context.webkitBackingStorePixelRatio ||
          context.mozBackingStorePixelRatio ||
          context.msBackingStorePixelRatio ||
          context.oBackingStorePixelRatio ||
          context.backingStorePixelRatio || 1;
    return (window.devicePixelRatio || 1) / backingStore;
  };
  var canvas = document.createElement("canvas");
  var context = canvas.getContext("2d");
  var ratio=getPixelRatio(context);
  var w = $("#canv").width();
  var h = $("#canv").height();
  canvas.width =document.body.clientWidth * 1;// w*1;ratio;
  canvas.height = document.body.clientHeight * 1;//h*1;ratio;//h * 2;
  //然后将画布缩放，将图像放大两倍画到画布上
  context.scale(1,1);
  html2canvas(document.getElementById("canv"), {
    // "logging": true, //Enable log (use Web Console for get Errors and Warnings)
    // "proxy": basePath + "canvas/proxy.do",
    // allowTaint: false,
    // taintTest: true,
    canvas: canvas,
    "onrendered": function(canvas) {
      canvas.style.width=document.body.clientWidth+ "px";
      canvas.style.height=document.body.clientHeight+ "px";
      var img = new Image();
      img.crossOrigin = "*";
      img.onload = function() {
        img.onload = null;
        //document.body.appendChild(img);
        // $("#canv").append(img);
      };
      img.onerror = function() {
        img.onerror = null;
        if(window.console.log) {
          window.console.log("Not loaded image from canvas.toDataURL");
        } else {
        alert("Not loaded image from canvas.toDataURL");
        }
      };
      img.src = canvas.toDataURL("image/png");
      var triggerDownload = $("<a>").attr("href", img.src).attr("download", "img.png").appendTo("body"); 
      triggerDownload[0].click(); 
      triggerDownload.remove();
    }
  });
  }

      // $(function(){     
      //     print();  
      // });  
      // function print(){//不完整   
      //     html2canvas(document.body, {            
      //         onrendered: function(canvas){  
      //             $('#down_button').attr('href',canvas.toDataURL("image/png")) ;  
      //             $('#down_button').attr('download','myjobdeer.png') ;  
      //         }  
      //     });  
      // }
      // html2canvas(document.body, {
      //   onrendered: function(canvas) {
      //     document.body.appendChild(canvas);
      //   }
      // });
    //   function print(){  
    //     html2canvas(document.body, {            
    //         onrendered: function(canvas){  
    //             var canvas = document.getElementById("<span style="font-family: Arial, Helvetica, sans-serif;">canvas"</span><span style="font-family: Arial, Helvetica, sans-serif;">),</span> 
    //               url = canvas.toDataURL(); 
    //               //以下代码为下载此图片功能 
    //               var triggerDownload = $("<a>").attr("href", url).attr("download", "img.png").appendTo("body"); 
    //               triggerDownload[0].click(); 
    //               triggerDownload.remove();
    //         },  
    //     // height:9000  
    //     });  
    // }


    // .on('click','.download',function(){
    //             $('#mycanvas').remove();
    //             var _height=$('.skinReport').height();
    //             //滚到顶部
    //             $('html, body').animate({scrollTop:0});
    //             if(confirm('是否下载肌肤检测报告？'))
    //             {
    //                 setTimeout(function(){
    //                     var canvas = document.createElement("canvas"),
    //                         w=$('#skinReport').width(),
    //                         h=$('#skinReport').height();
    //                     canvas.width = w * 2;
    //                     canvas.height = h * 2;
    //                     canvas.style.width = w + "px";
    //                     canvas.style.height = h + "px";
    //                     var context = canvas.getContext("2d");
    //                     //然后将画布缩放，将图像放大两倍画到画布上
    //                     context.scale(2,2);
    //                     html2canvas(document.getElementById('skinReport'), {
    //                         allowTaint: false,
    //                         taintTest: true,
    //                         canvas: canvas,
    //                         onrendered: function(canvas) {
    //                             canvas.id = "mycanvas";
    //                             canvas.style.display = 'none';
    //                             document.body.appendChild(canvas);
    //                             //生成base64图片数据
    //                             imgData = canvas.toDataURL(type);
    //                             //var newImg = document.createElement("img");
    //                             //newImg.src =  dataUrl;
    //                             //document.body.appendChild(newImg);
    //                             //console.log(imgData);
    //                             var _fixType = function(type) {
    //                                 type = type.toLowerCase().replace(/jpg/i, 'jpeg');
    //                                 var r = type.match(/png|jpeg|bmp|gif/)[0];
    //                                 return 'image/' + r;
    //                             };
    //                             // 加工image data，替换mime type
    //                             imgData = imgData.replace(_fixType(type),'image/octet-stream');
    //                             /**
    //                              * 在本地进行文件保存
    //                              * @param  {String} data     要保存到本地的图片数据
    //                              * @param  {String} filename 文件名
    //                              */
    //                             var saveFile = function(data, filename){
    //                                 var save_link = document.createElementNS('http://www.w3.org/1999/xhtml', 'a');
    //                                 save_link.href = data;
    //                                 save_link.download = filename;

    //                                 var event = document.createEvent('MouseEvents');
    //                                 event.initMouseEvent('click', true, false, window, 0, 0, 0, 0, 0, false, false, false, false, 0, null);
    //                                 save_link.dispatchEvent(event);
    //                             };

    //                             // 下载后的问题名
    //                             var filename = aname+'肌肤检测报告' + (new Date()).getTime() + '.' + type;
    //                             // download
    //                             saveFile(imgData,filename);
    //                         },
    //                         width:1512,
    //                         height:15000
    //                     })
    //                 },2500)
    //             }
    //             else
    //             {
    //                 return;
    //             }

    //         })
  </script>
  <style>
    .right-top 
    {
        /*width: 40px;*/
        /*height:40px;*/
        position: fixed;/*这是必须的*/
        z-index: 999;
        left:95%;/*这是必须的*/
        top: 40px;/*这是必须的*/
        /*background:red;*/
    }
  </style>
  </head>
 	<!-- Dark Responsive Timeline with Bootstrap黑色Bootstrap响应式时间轴 -->
<body>
<input type="button" class='btn btn-warning right-top' value="截图" onclick="takeScreenshot()" id="down">
	<div class="container" id="canv">
	  <header class="page-header">
	    <h1>{{.ProjectTile}}</h1>
	    <h1><small>——大事记</small></h1>
      <!-- <a type="button" id="down_button">download</a> -->
      
	  </header>
	  
	  <ul class="timeline">
  	</ul>

	</div>

<script type="text/javascript">
var i=1;
var j=1;
var d=new Date();
var year;//=d.getFullYear();
var month;//=d.getMonth() + 1;
	$(function(){
  	// var _timeline_date_ = $("<li><div class='tldate'>Apr 2017<div><li>");
  	// $(".timeline").append(_timeline_date_);
  	var loadData=function(ii){
  		// alert(ii);
    	$.getJSON("/project/"+{{.ProjectId}}+"/timeline?p="+ii, function (data) {
      	$.each(data, function (i, tl) {
					// alert(moment(tl["start"], 'YYYY-MM-DD').format('YYYY-MM-DD'));
        	// var d=new Date();
        	// var day=d.getDate();
					// var month=d.getMonth() + 1;
					// var year=d.getFullYear();
					// alert(tl["start"]);
					// var ret=/http:(.*)\:/.exec(value[i].Link);//http://127.0.0.1:
      		// var site=/-.*?-/.exec(tl["start"]);//非贪婪模式 
					array=tl["start"].split("-");
					// alert(array[0]);
					// alert(array[1]);
					// var year=array[0];
					// var month=array[1];
					if (array[0]!=year||array[1]!=month){
						var _timeline_date_ = $("<li><div class='tldate'>"+array[0]+"-"+array[1]+"<div><li>");
  					$(".timeline").append(_timeline_date_);
						year=array[0];
						month=array[1];
					}

					if(i%2==1){
						var _timeline_invert_ = $("<li></li>");
					}else{
						var _timeline_invert_ = $("<li></li>").addClass("timeline-inverted");
					}

        	$(".timeline").append(_timeline_invert_);

        	var _timeline_icon_ = $("<div></div>").addClass("timeline-icon timeline-icon-hide-border");
        	var _timeline_fa_ = $("<i style='color:#c23b22'></i>").addClass("fa fa-github fa-lg");
        	_timeline_icon_.append(_timeline_fa_);
        	_timeline_invert_.append(_timeline_icon_);
        	/**
        	 * 设置显示内容
        	 */
					var _timeline_panel_ = $("<div></div>").addClass("timeline-panel");
					var _timeline_head_ = $("<div></div>").addClass("tl-heading");
					var _timeline_body_ = $("<div></div>").addClass("tl-body");

					var _timeline_body_p1_ = $("<p></p>").text(tl["content"]);
					_timeline_body_.append(_timeline_body_p1_);
					var _timeline_body_p2_ = $("<p></p>");
					_timeline_body_p2_.append($("<img/>").attr("src","/"+tl["image"]));
					_timeline_body_.append(_timeline_body_p1_);
					_timeline_body_.append(_timeline_body_p2_);

        	var _timeline_head_h4_ = $("<h4></h4>").text(tl["title"]);
					var _timeline_head_p_ = $("<p></p>");
					var _timeline_head_p_small_ = $("<small class='text-muted'></small>");
					var _timeline_head_p_small_i_ = $("<i class='glyphicon glyphicon-time'></i>");
					_timeline_head_p_small_.append(_timeline_head_p_small_i_);
					_timeline_head_p_small_.append(moment(tl["start"], 'YYYY-MM-DD').format('YYYY-MM-DD'));

					_timeline_head_p_.append(_timeline_head_p_small_);

					_timeline_head_.append(_timeline_head_h4_);
					_timeline_head_.append(_timeline_head_p_);

					_timeline_panel_.append(_timeline_head_);
					_timeline_panel_.append(_timeline_body_);

					_timeline_invert_.append(_timeline_panel_);
      	});

				if ({{.Count}}>=i*2){
      		if($(window).height()>=document.documentElement.scrollHeight){
      	  	//没有出现滚动条,继续加载下一页
      	  	loadData(i);
      	  	i=i+1;
      		}
				}else if({{.Count}}==i*2-1){//当count是单数的时候
          loadData(i);
          i=i+1;
        }

        // alert("第一个"+{{.Count}}+i)
    	});
  	}

  	var tcScroll=function(){
  	  $(window).on('scroll', function () {
  	    var scrollTop = $(this).scrollTop();
  	    var scrollHeight = $(document).height();
  	    var windowHeight = $(this).height();
  	    //得到总长度，如果超出长度，则不再触发
  	    // alert(i);
  	    // alert(i*2);
  	    // alert({{.Count}});
  	    if ({{.Count}}>=i*2){
  	    	if (scrollTop + windowHeight == scrollHeight) {
        		//此处是滚动条到底部时候触发的事件，在这里写要加载的数据，或者是拉动滚动条的操作
        		loadData(i);
        		i=i+1;
      		}
      	}else if({{.Count}}==i*2-1){
          loadData(i);
          i=i+1;
        }
    	})
  	}
  	loadData(i);
  	tcScroll();
  	i=i+1;
  	// j=j+1;

	});

	</script>

	</body>
</html>
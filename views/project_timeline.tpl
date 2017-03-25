<!-- 测试时间轴 -->
<!doctype html>
<html lang="en-US">
<head>
  <meta charset="utf-8">
  <meta http-equiv="Content-Type" content="text/html">
  <title>Vertical Responsive Timeline UI - Template Monster Demo</title>
  <meta name="author" content="Jake Rocheleau">
  <link rel="stylesheet" type="text/css" media="all" href="/static/css/bootstrap.min.css">
  <link rel="stylesheet" type="text/css" media="all" href="/static/css/timeline.css">
	<link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
 	<!-- Dark Responsive Timeline with Bootstrap黑色Bootstrap响应式时间轴 -->
<body>
	<div class="container">
	  <header class="page-header">
	    <h1>{{.ProjectTile}}</h1>
	    <h1><small>——大事记</small></h1>
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
					_timeline_body_p2_.append($("<img/>").attr("src",tl["image"]));
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
				}
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
<!-- iframe里日历-->
<!DOCTYPE html>
<html>
<head>
<link rel='stylesheet' href='/static/css/fullcalendar.min.css' />
<script src='/static/js/jquery-2.1.3.min.js'></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
<script src='/static/js/moment.min.js'></script>
<script src='/static/js/fullcalendar.min.js'></script>
<script src='/static/js/fullcalendar.zh-cn.js'></script>
<script src='/static/js/bootstrap-datetimepicker.min.js'></script>
<script src='/static/js/bootstrap-datetimepicker.zh-CN.js'></script>

<link rel='stylesheet' href='/static/css/bootstrap-datetimepicker.min.css'/>
<link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>

<style>
	/*body {
		margin: 0;
		padding: 0;
		font-family: "Lucida Grande",Helvetica,Arial,Verdana,sans-serif;
		font-size: 14px;
	}*/

	#script-warning {
		display: none;
		background: #eee;
		border-bottom: 1px solid #ddd;
		padding: 0 10px;
		line-height: 40px;
		text-align: center;
		font-weight: bold;
		font-size: 12px;
		color: red;
	}

	#loading {
		display: none;
		position: absolute;
		top: 10px;
		right: 10px;
	}

	#calendar {
		max-width: 900px;
		margin: 40px auto;
		padding: 0 10px;
	}

	/*body {这个导致窗口挪动，不好
		margin: 40px 10px;
		padding: 0;
		font-family: "Lucida Grande",Helvetica,Arial,Verdana,sans-serif;
		font-size: 14px;
	}*/

	/*#calendar {
		max-width: 900px;
		margin: 0 auto;
	}*/
		  .fc-color-picker {
    		list-style: none;
    		margin: 0;
    		padding: 0;
    	}
    	.fc-color-picker > li {
    	  float: left;
    	  font-size: 30px;
    	  margin-right: 5px;
    	  line-height: 30px;
    	}
    	.fc-color-picker > li .fa {
    	  -webkit-transition: -webkit-transform linear 0.3s;
    	  -moz-transition: -moz-transform linear 0.3s;
    	  -o-transition: -o-transform linear 0.3s;
    	  transition: transform linear 0.3s;
    	}
    	.fc-color-picker > li .fa:hover {
    	  -webkit-transform: rotate(30deg);
    	  -ms-transform: rotate(30deg);
    	  -o-transform: rotate(30deg);
    	  transform: rotate(30deg);
    	}
    	.text-red {
  			color: #dd4b39 !important;
		  }
		  .text-yellow {
		    color: #f39c12 !important;
		  }
		  .text-aqua {
		    color: #00c0ef !important;
		  }
		  .text-blue {
		    color: #0073b7 !important;
		  }
		  .text-black {
		    color: #111111 !important;
		  }
		  .text-light-blue {
		    color: #3c8dbc !important;
		  }
		  .text-green {
		    color: #00a65a !important;
		  }
		  .text-gray {
		    color: #d2d6de !important;
		  }
		  .text-navy {
		    color: #001f3f !important;
		  }
		  .text-teal {
		    color: #39cccc !important;
		  }
		  .text-olive {
		    color: #3d9970 !important;
		  }
		  .text-lime {
		    color: #01ff70 !important;
		  }
		  .text-orange {
		    color: #ff851b !important;
		  }
		  .text-fuchsia {
		    color: #f012be !important;
		  }
		  .text-purple {
		    color: #605ca8 !important;
		  }
		  .text-maroon {
		    color: #d81b60 !important;
		  }
</style>
</head>
<!-- <body> -->
<script type="text/javascript">
$(document).ready(function() {
    // page is now ready, initialize the calendar...
    $('#calendar').fullCalendar({
        // put your options and callbacks here
      header: {
				left: 'prev,next today',
				center: 'title',
				right: 'month,agendaWeek,agendaDay,listMonth'
			},
			// defaultDate: '2017-01-12',
			navLinks: true, // can click day/week names to navigate views
			editable: true,
			eventLimit: true, // allow "more" link when too many events
			businessHours: true, // display business hours
			selectable: true,
			selectHelper: true,
			// select: function(start, end) {
			// 	var title = prompt('Event Title:');
			// 	var eventData;
			// 	if (title) {
			// 		eventData = {
			// 			title: title,
			// 			start: start,
			// 			end: end,
			// 			// Color: getRandomColor(),
			// 			textColor: getRandomColor(),
			// 			backgroundColor: getRandomColor(),
			// 			// borderColor: getRandomColor(),
   			//             className: 'done',
			// 		};
			// 		$('#calendar').fullCalendar('renderEvent', eventData, true); // stick? = true
			// 	}
			// 	$('#calendar').fullCalendar('unselect');
			// },
			select: function(start, end, jsEvent, view){  
                    //添加日程事件
                    // $("input#cid").remove();
      				// var th1="<input id='cid' type='hidden' name='cid' value='" +selectRow[0].Id+"'/>"
      				// $(".modal-body").append(th1);//这里是否要换名字$("p").remove();
      				$("#start").val(start.format('YYYY-MM-DD HH:mm'));
      				$("#end").val(end.format('YYYY-MM-DD HH:mm'));
        			$('#modalTable').modal({
        				show:true,
        				backdrop:'static'
        			});
            },
			editable: true,
			// events: '/admin/calendar',
			// eventSources: [  
      //      		'/feed1.php',  
      //      		'/feed2.php'  
      //  		]
			events: {  
        		url: '/admin/calendar',  
        		type: 'post'  
    		},
			// events: {
			// 	url: '/admin/getcalendar',
			// 	error: function() {
			// 		$('#script-warning').show();
			// 	}
			// },
			loading: function(bool) {
				$('#loading').toggle(bool);
			},
			// events: [
			// 	{
			// 		title: 'All Day Event',
			// 		start: '2017-01-01'
			// 	},
			// 	{
			// 		title: 'Long Event',
			// 		start: '2017-01-07',
			// 		end: '2017-01-10'
			// 	},
			// 	{
			// 		id: 999,
			// 		title: 'Repeating Event',
			// 		start: '2017-01-09T16:00:00'
			// 	},
			// 	{
			// 		id: 999,
			// 		title: 'Repeating Event',
			// 		start: '2017-01-16T16:00:00'
			// 	},
			// 	{
			// 		title: 'Conference',
			// 		start: '2017-01-11',
			// 		end: '2017-01-13'
			// 	},
			// 	{
			// 		title: 'Meeting',
			// 		start: '2017-01-12T10:30:00',
			// 		end: '2017-01-12T12:30:00'
			// 	},
			// 	{
			// 		title: 'Lunch',
			// 		start: '2017-01-12T12:00:00'
			// 	},
			// 	{
			// 		title: 'Meeting',
			// 		start: '2017-01-12T14:30:00'
			// 	},
			// 	{
			// 		title: 'Happy Hour',
			// 		start: '2017-01-12T17:30:00'
			// 	},
			// 	{
			// 		title: 'Dinner',
			// 		start: '2017-01-12T20:00:00'
			// 	},
			// 	{
			// 		title: 'Birthday Party',
			// 		start: '2017-01-13T07:00:00'
			// 	},
			// 	{
			// 		title: 'Click for Google',
			// 		url: 'http://google.com/',
			// 		start: '2017-01-28'
			// 	}
			// ],
			dayClick: function(date, jsEvent, view) {
        		// alert('Clicked on: ' + date.format());
        		// alert('Coordinates: ' + jsEvent.pageX + ',' + jsEvent.pageY);
        		// alert('Current view: ' + view.name);
        		// change the day's background color just for fun
        		// $(this).css('background-color', getRandomColor());
    	},
    	eventClick: function(data, jsEvent, view){  
                    //修改日程事件 
        $("input#cid").remove();
      	var th1="<input id='cid' type='hidden' name='cid' value='" +data.id+"'/>"
      				$(".modal-body").append(th1);
      				$("#title1").val(data.title);
      				$("#content1").val(data.content);
              $("#isallday1").prop('checked',data.allDay);
                $("#ispublic1").prop('checked',false);
                // $("#ispublic1[name='private']").prop('checked',false);
              if (data.Public==true){
                $("#ispublic1[value='true']").prop('checked',true);
              }else{
                $("#ispublic1[value='false']").prop('checked',true);
              }
              // $("#ispublic1").prop('checked',data.Public);

      				$("#start1").val(data.start.format('YYYY-MM-DD HH:mm'));
      				$("#end1").val(data.end.format('YYYY-MM-DD HH:mm'));
					    $('#add-new-event1').css({"background-color": data.color, "border-color": data.color});
        			$('#modalTable1').modal({
        				show:true,
        				backdrop:'static'
        			});  
            },
        eventDrop: function(event,delta,revertFunc) {
          // alert(event.id+event.title+delta.days());
          // var url = "/admin/calendar/dropcalendar";
          // $.post(url,{id:event.id,dalta:delta.days()},function(msg){
          //   });
          $.ajax({
                type:"post",
                url:"/admin/calendar/dropcalendar",
                data: {id:event.id,delta:delta.days()},
                success:function(data,status){
                  alert("修改“"+data+"”成功！(status:"+status+".)");
                },
                error:function(data,status){
                  alert(data);
                  revertFunc();
                }
            });
        },
        eventResize: function(event,delta,revertFunc) {
          // alert(delta.asHours());
          $.ajax({
                type:"post",
                url:"/admin/calendar/resizecalendar",
                data: {id:event.id,delta:delta.asHours()},
                success:function(data,status){
                  alert("修改“"+data+"”成功！(status:"+status+".)");
                },
                error:function(data,status){
                  alert(data);
                  revertFunc();
                }
            });     
          }
});
});
// RGB 转16进制
  var rgbToHex = function(rgb) {
    // rgb(x, y, z)
    var color = rgb.toString().match(/\d+/g); // 把 x,y,z 推送到 color 数组里
    var hex = "#";
    for (var i = 0; i < 3; i++) {
      // 'Number.toString(16)' 是JS默认能实现转换成16进制数的方法.
      // 'color[i]' 是数组，要转换成字符串.
      // 如果结果是一位数，就在前面补零。例如： A变成0A
      hex += ("0" + Number(color[i]).toString(16)).slice(-2);
    }
    return hex;
  }
  // 16进制 转 RGB
  
  // 能处理 #axbycz 或 #abc 形式
  var hexToRgb = function(hex) {
    var color = [], rgb = [];
    hex = hex.replace(/#/,"");
    if (hex.length == 3) { // 处理 "#abc" 成 "#aabbcc"
      var tmp = [];
      for (var i = 0; i < 3; i++) {
        tmp.push(hex.charAt(i) + hex.charAt(i));
      }
      hex = tmp.join("");
    }
    for (var i = 0; i < 3; i++) {
      color[i] = "0x" + hex.substr(i+2, 2);
      rgb.push(parseInt(Number(color[i])));
    }
    return "rgb(" + rgb.join(",") + ")";
  }

	function getRandomColor(){ 
		var c = '#'; 
		var cArray = ['0','1','2','3','4','5','6','7','8','9','A','B','C','D','E','F']; 
		for(var i = 0; i < 6;i++){ 
			var cIndex = Math.round(Math.random()*15); 
			c += cArray[cIndex]; 
		} 
			return c; 
	}

	function save(){
    var title = $('#title').val();
    var content = $('#content').val();
    var start = $('#start').val();
    var end = $('#end').val();
    var allday=document.getElementById("isallday").checked;
    var public=document.getElementById("ispublic").checked;
    // alert(allday);
    // alert(public);
      if (title){  
            $.ajax({
                type:"post",
                url:"/admin/calendar/addcalendar",
                data: {title:title,content:content,allday:allday,public:public,start:start,end:end,color:rgbToHex(currColor)},
                success:function(data,status){
                  alert("添加“"+data+"”成功！(status:"+status+".)");
                  var eventData;
					       if (title) {
					       eventData = {
					       	title: title,
					       	content: content,
					       	start: start,
					       	end: end,
                  color:rgbToHex(currColor),
					       	// textColor: getRandomColor(),
					       	// backgroundColor: rgbToHex(currColor),
					       	// borderColor: rgbToHex(currColor),
   			            className: 'done',
					         };
					         // $('#calendar').fullCalendar('renderEvent', eventData, true); // stick? = true要用下面这个，否则添加后立即删除，无法删除
            $('#calendar').fullCalendar('refetchEvents'); //重新获取所有事件数据
				    }
				    $('#calendar').fullCalendar('unselect'); 
          			$('#modalTable').modal('hide');
                }
            });  
      }
    }

  function update(){
  	  var title = $('#title1').val();
      var content = $('#content1').val();
      var start = $('#start1').val();
      var end = $('#end1').val();
      var cid = $('#cid').val();
      var allday=document.getElementById("isallday1").checked;
      var public=document.getElementById("ispublic1").checked;
      var currColor=$('#add-new-event1').css("background-color");
      if (title){  
            $.ajax({
                type:"post",
                url:"/admin/calendar/updatecalendar",
                data: {cid:cid,title:title,content:content,allday:allday,public:public,start:start,end:end,color:rgbToHex(currColor)},
                success:function(data,status){
                  alert("修改“"+data+"”成功！(status:"+status+".)");
					       $('#calendar').fullCalendar('refetchEvents'); //重新获取所有事件数据 // stick? = true 
          			   $('#modalTable1').modal('hide');
                }
            });  
      }
  }
  //删除事件
    function delete_event(){
        if(confirm("您确定要删除吗？")){
            var cid = $("#cid").val();  
              $.ajax({
                type:"post",
                url:"/admin/calendar/deletecalendar",
                data: {cid:cid},
                success:function(data,status){
                  alert("删除“"+data+"”成功！(status:"+status+".)");
					 //从日程视图中删除该事件
                     $("#calendar").fullCalendar("removeEvents",cid); // stick? = true
          			$('#modalTable1').modal('hide');
                }
              });  
        }
    } 
</script>
<!-- <div class="col-lg-12"> -->
	<div id='calendar'></div>
<!-- </div> -->

	<!-- 新建日程窗口 -->
<div class="container">
  <form class="form-horizontal">
    <div class="modal fade" id="modalTable">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">添加日程</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">标题</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="title"></div>
              </div>
              <div class="form-group must">
                    <label class="col-sm-3 control-label">内容</label>
                    <div class="col-sm-7">
                      <textarea class="form-control" rows="3" id='content'></textarea>
                    </div>    
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">全天事件</label>
                <div class="col-sm-7 checkbox">
                  <label>
                  <input type="checkbox" value="true" id="isallday"></label>
                  <label>
                  <input type="radio" id="ispublic" value="true" name="public" checked="checked">公开</label>
                  <label>
                  <input type="radio" id="ispublic" value="false" name="public">私有</label>
                </div>
              </div>
              <!-- $('input:radio:checked').val()；
              $("input[type='radio']:checked").val();
              $("input[name='rd']:checked").val(); --> 
              <div class="form-group must">
                <label class="col-sm-3 control-label">开始时间</label>
                <div class="input-group date form_datetime col-sm-7" data-link-field="dtp_input1">
                    <input class="form-control" type="text" id="start" readonly="">
                    <span class="input-group-addon"><span class="glyphicon glyphicon-remove"></span></span>
					         <span class="input-group-addon"><span class="glyphicon glyphicon-th"></span></span>
                </div>
				        <input type="hidden" id="dtp_input1" value="">
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">结束时间</label>
                <div class="input-group date form_datetime col-sm-7" data-link-field="dtp_input1">
                    <input class="form-control" type="text" id="end" readonly="">
                    <span class="input-group-addon"><span class="glyphicon glyphicon-remove"></span></span>
					           <span class="input-group-addon"><span class="glyphicon glyphicon-th"></span></span>
                </div>
				        <input type="hidden" id="dtp_input1" value="">
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">选择背景色</label>
                <div class="col-sm-7">
                  <div class="btn-group" style="width: 100%; margin-bottom: 10px;">
                    <ul class="fc-color-picker" id="color-chooser">
                      <li><a class="text-aqua" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-blue" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-light-blue" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-teal" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-yellow" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-orange" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-green" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-lime" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-red" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-purple" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-fuchsia" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-muted" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-navy" href="#"><i class="fa fa-square"></i></a></li>
                    </ul>
                  </div>
                </div>
              </div>
          </div>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
          <button type="button" id="add-new-event" class="btn btn-primary" onclick="save()">保存</button>
        </div>
      </div>
    </div>
  </div>
</form>
</div>

<!-- 编辑日程窗口 -->
<div class="container">
  <div class="form-horizontal">
    <div class="modal fade" id="modalTable1">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">编辑日程</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">标题</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="title1"></div>
              </div>
              <div class="form-group must">
                    <label class="col-sm-3 control-label">内容</label>
                    <div class="col-sm-7">
                      <textarea class="form-control" rows="3" id='content1'></textarea>
                    </div>    
               </div>
               <div class="form-group must">
                <label class="col-sm-3 control-label">全天事件</label>
                <div class="col-sm-7 checkbox">
                  <label>
                  <input type="checkbox" value="true" id="isallday1"></label>
                  <label>
                  <input type="radio" id="ispublic1" value="true" name="public1" checked="checked">公开</label>
                  <label>
                  <input type="radio" id="ispublic1" value="false" name="public1">私有</label>
                </div>
              </div> 
               <div class="form-group must">
                <label class="col-sm-3 control-label">开始时间</label>
                <!-- <div class="col-sm-7">
                  <input type="text" class="form-control" id="start">
                </div> -->
                <div class="input-group date form_datetime col-sm-7" data-link-field="dtp_input1">
                    <input class="form-control" type="text" id="start1" readonly="">
                    <span class="input-group-addon"><span class="glyphicon glyphicon-remove"></span></span>
					           <span class="input-group-addon"><span class="glyphicon glyphicon-th"></span></span>
                </div>
				        <input type="hidden" id="dtp_input1" value="">
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">结束时间</label>
                <!-- <div class="col-sm-7">
                  <input type="tel" class="form-control" id="end">
                </div> -->

                <div class="input-group date form_datetime col-sm-7" data-link-field="dtp_input1">
                    <input class="form-control" type="text" id="end1" readonly="">
                    <span class="input-group-addon"><span class="glyphicon glyphicon-remove"></span></span>
					           <span class="input-group-addon"><span class="glyphicon glyphicon-th"></span></span>
                </div>
				        <input type="hidden" id="dtp_input1" value="">

              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">选择背景色</label>
                <div class="col-sm-7">
                  <div class="btn-group" style="width: 100%; margin-bottom: 10px;">
                  <!--<button type="button" id="color-chooser-btn" class="btn btn-info btn-block dropdown-toggle" data-toggle="dropdown">Color <span class="caret"></span></button>-->
                    <ul class="fc-color-picker" id="color-chooser1">
                      <li><a class="text-aqua" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-blue" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-light-blue" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-teal" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-yellow" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-orange" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-green" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-lime" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-red" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-purple" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-fuchsia" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-muted" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-navy" href="#"><i class="fa fa-square"></i></a></li>
                    </ul>
                  </div>
                </div>
              </div>
          </div>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
          <!-- <button type="button" class="btn btn-primary" onclick="update()">修改</button> -->
          <button type="button" id="add-new-event1" class="btn btn-primary" onclick="update()">修改</button>
          <button type="button" class="btn btn-danger" onclick="delete_event()">删除</button>
        </div>
      </div>
    </div>
  </div>
</div>
</div> 

<script type="text/javascript">
    $('.form_datetime').datetimepicker({
        language:  'zh-CN',
        weekStart: 1,
        todayBtn:  1,
		autoclose: 1,
		todayHighlight: 1,
		startView: 2,
		forceParse: 0,
        showMeridian: 1
    });
	$('.form_date').datetimepicker({
        language:  'zh-CN',
        weekStart: 1,
        todayBtn:  1,
		autoclose: 1,
		todayHighlight: 1,
		startView: 2,
		minView: 2,
		forceParse: 0
    });
	$('.form_time').datetimepicker({
        language:  'zh-CN',
        weekStart: 1,
        todayBtn:  1,
		autoclose: 1,
		todayHighlight: 1,
		startView: 1,
		minView: 0,
		maxView: 1,
		forceParse: 0
    });
</script>
<script>
  var currColor = "#3c8dbc"; //Red by default
  $(function () {
    /* ADDING EVENTS */
    //Color chooser button
    // var colorChooser = $("#color-chooser-btn");
    $("#color-chooser > li > a").click(function (e) {
      e.preventDefault();
      //Save color
      currColor = $(this).css("color");
      //Add color effect to button
      $('#add-new-event').css({"background-color": currColor, "border-color": currColor});
    });
    $("#color-chooser1 > li > a").click(function (e) {
      e.preventDefault();
      //Save color
      currColor = $(this).css("color");
      //Add color effect to button
      $('#add-new-event1').css({"background-color": currColor, "border-color": currColor});
    });


  //   $("#isallday").click(function(){//是否是全天事件
  //     if($("#sel_start").css("display")=="none"){
  //       $("#sel_start,#sel_end").show();
  //     }else{
  //        $("#sel_start,#sel_end").hide();
  //     }
  //   });
  //   $("#isend").click(function(){//是否有结束时间
  //     if($("#p_endtime").css("display")=="none"){
  //       $("#p_endtime").show();
  //     }else{
  //       $("#p_endtime").hide();
  //     }
  // });

    // $("#add-new-event").click(function (e) {
    //   e.preventDefault();
    //   //Get value and make sure it is not null
    //   var val = $("#new-event").val();
    //   if (val.length == 0) {
    //     return;
    //   }

    //   //Create events
    //   var event = $("<div />");
    //   event.css({"background-color": currColor, "border-color": currColor, "color": "#fff"}).addClass("external-event");
    //   event.html(val);
    //   $('#external-events').prepend(event);

    //   //Add draggable funtionality
    //   ini_events(event);

    //   //Remove event from text input
    //   $("#new-event").val("");
    // });
  });
</script>

<!-- </body> -->
</html>
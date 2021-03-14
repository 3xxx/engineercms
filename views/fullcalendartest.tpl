<!DOCTYPE html>
<html>
<link rel="stylesheet" href="$!{rc.contextPath}/modules/devappwithfullcanlendar/css/mainstructure.css"> 
<link rel="stylesheet" href="$!{rc.contextPath}/modules/devappwithfullcanlendar/css/maincontent.css"> 
<!-- Jquery and Jquery UI -->

<script type="text/javascript" src="$!{rc.contextPath}/modules/devappwithfullcanlendar/js/jquery-1.4.2.min.js"></script>

<script type="text/javascript" src="$!{rc.contextPath}/modules/devappwithfullcanlendar/js/jquery-ui-1.8.6.custom.min.js"></script>

<script type="text/javascript" src="$!{rc.contextPath}/modules/devappwithfullcanlendar/js/jquery-ui-timepicker-addon.js"></script>

<link rel="stylesheet" href="$!{rc.contextPath}/modules/devappwithfullcanlendar/css/redmond/jquery-ui-1.8.1.custom.css"> 

<!-- Jquery and Jquery UI -->

<script src="$!{rc.contextPath}/modules/devappwithfullcanlendar/js/formValidator/js/jquery.validationEngine.js" type="text/javascript"></script>

<script src="$!{rc.contextPath}/modules/devappwithfullcanlendar/js/formValidator/js/jquery.validationEngine-en.js" type="text/javascript"></script>

<link rel="stylesheet" href="$!{rc.contextPath}/modules/devappwithfullcanlendar/js/formValidator/css/validationEngine.jquery.css" type="text/css" media="screen" charset="utf-8" />

<!-- FullCalender -->

<link rel='stylesheet' type='text/css' href='$!{rc.contextPath}/modules/devappwithfullcanlendar/js/fullcal/css/fullcalendar.css' />
<script type='text/javascript' src='$!{rc.contextPath}/modules/devappwithfullcanlendar/js/fullcal/fullcalendar.js'></script>
<BODY>

<STYLE type=text/css>#loading {
    TOP: 0px; RIGHT: 0px
}
.tooltip {
    PADDING-BOTTOM: 25px; PADDING-LEFT: 25px; WIDTH: 160px; PADDING-RIGHT: 25px; DISPLAY: none; BACKGROUND: url(images/black_arrow.png); HEIGHT: 70px; COLOR: #fff; FONT-SIZE: 12px; PADDING-TOP: 25px; z-order: 100
}
</STYLE>

<DIV id=wrap>
<SCRIPT type=text/javascript>
$(document).ready(function() {
    
    $("#groupName").click(function(){
            dialog({
                content: 'url:$!{rc.contextPath}/duty/dutyGroup/lookup?txtId=groupId&txtName=groupName&from=canlendar',
                zIndex:20000,
                title: '选择',
                lock: false,
                width: 300,
                height: 300,
 
            });
        }); 

    $("#reserveformID").validationEngine({
        validationEventTriggers:"keyup blur", // 键盘按键触发验证
        openDebug: true
    }) ;


    $("#addhelper").hide();        
    
    var calendar =$('#calendar').fullCalendar({
      header:{
            right: 'prev,next today',
            center: 'title',
            left: 'month,agendaWeek,agendaDay'
      },
      monthNames: ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"],
      monthNamesShort: ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"],
      dayNames: ["周日", "周一", "周二", "周三", "周四", "周五", "周六"],
      dayNamesShort: ["周日", "周一", "周二", "周三", "周四", "周五", "周六"],
      today: ["今天"],
      firstDay: 1,                
         buttonText: {
              today: '今天',
              month: '月',
              week: '周',
              day: '日',
              prev: '上一月',
              next: '下一月'
       },

      theme: true,
      editable: false,//日历项拖拽
      allDaySlot : false,
      events:  function(start, end , callback){//生成日历
          //alert(calendar.fullCalendar('getDate'));
        var events = [];
        $.ajax({
            'url':"$!{rc.contextPath}/duty/dutyScheduling/canlendarModel?time="+new Date().getTime(),
            'data': {                
                timeStart:$.fullCalendar.formatDate(start,"yyyy-MM-dd HH:mm:ss"),
                timeEnd:$.fullCalendar.formatDate(end,"yyyy-MM-dd HH:mm:ss")       
            },
            'dataType': 'json',
            'type': 'post',
            'error': function(data){
                alert("保存失败");
                 return false;
             },
            'success': function(doc) {
                $(doc).each(function(i) {
                    events.push({
                        sid: doc[i].id,
                        uid: doc[i].id,
                        title: 'Daily Scrum meeting',
                        start: doc[i].startTime,
                        end: doc[i].endTime,
                        fullname: doc[i].groupName,// 此处改变，列表和明细都会改变
                        confname: doc[i].orderName,
                        groupPersons: doc[i].personNames,
                        confshortname: doc[i].orderName.substring(0, 1),
                        confcolor: '#ff3f3f',
                        confid: 'test1',
                        allDay: false,
                        topic: doc[i].orderName,
                        description : doc[i].remark,
                        id: 1
                    });    
                });
                callback(events);//
            }
        });                        
        
      },
      dayClick: function(date, allDay, jsEvent, view) {// 单机日历内空白 新增事件
                var selectdate = $.fullCalendar.formatDate(date, "yyyy-MM-dd");    
                $( "#dutyDate" ).val(selectdate);
                // dutyDate
                $( "#reservebox" ).dialog({
                    autoOpen: false,
                    height: 450,
                    width: 400,
                    title: '新增排班' + selectdate ,// 此处声明title，会将reservebox中的title属性覆盖
                    modal: true,
                    position: "center",
                    draggable: true,// 可拖拽
                    beforeClose: function(event, ui) {
                            $.validationEngine.closePrompt("#orderId");
                            $.validationEngine.closePrompt("#groupId");
                            $.validationEngine.closePrompt("#groupName");                                
                    },
                    buttons: {// 弹出窗右下角的按钮
                        "关 闭": function() {
                            $( this ).dialog( "close" );
                        },
                        "保 存": function() {                
                            if($("#reserveformID").validationEngine({returnIsValid:true})){
                                var orderId = $("#orderId").val();    
                                var dutyDate = $("#dutyDate").val();
                                var groupId = $("#groupId").val();    
                                var groupName = $("#groupName").val();
                                var remark = $("#remark").val();
                                var schdata = {orderId:orderId, dutyDate:dutyDate, groupId:groupId, groupName:groupName,remark:remark};
                                $.ajax({
                                        'url':"$!{rc.contextPath}/duty/dutyScheduling/save?time="+new Date().getTime(),
                                        'data': schdata,
                                        'dataType': 'json',
                                        'type': 'post',
                                        'error': function(data){
                                        alert("保存失败");
                                         return false;
                                            },
                                        'success': function(data) {
                                        window.location.reload();
                                        }
                                 });
                            }    
                        }
                    }
                });
                $( "#reservebox" ).dialog( "open" );
            return false;
      },
      timeFormat: 'HH:mm{ - HH:mm}',
      eventClick: function(event) {// 单机日历内已有事件
              var tempStart = $.fullCalendar.formatDate(event.start, "yyyy/MM/dd");
            var tempEnd = $.fullCalendar.formatDate(event.end, "yyyy/MM/dd");
            if(tempStart==tempEnd){//若在同一天，结束日期省略
                var fstart  = $.fullCalendar.formatDate(event.start, "yyyy/MM/dd HH:mm");
                var fend  = $.fullCalendar.formatDate(event.end, "HH:mm");    
            }else{
                var fstart  = $.fullCalendar.formatDate(event.start, "yyyy/MM/dd HH:mm");
                var fend  = $.fullCalendar.formatDate(event.end, "yyyy/MM/dd HH:mm");    
            }
                          
            var schdata = {sid:event.sid, deleted:1, uid:event.uid};
            
            $( "#reserveinfo" ).dialog({
                autoOpen: false,
                height: 280,
                width: 400,
                modal: true,
                position: "center",
                draggable: true,
                buttons: {// 这里貌似不可以自定义添加其他按钮
                    "close": function() {
                        $( this ).dialog( "close" );
                    }
                }
            });
                                            
            if(1==1||2==schdata.uid){
                $("#reserveinfo").dialog("option", "buttons", {
                    "关闭": function() {
                        $( this ).dialog( "close" );
                    }
                });
            }
            
            var showtopic = '';
            
            if(event.topic.length>15){// 题目过长处理
                showtopic = event.topic.substring(0, 15) + '...';
            }else{
                showtopic = event.topic;
            }
            
            //明细弹出窗  描述
            $("#revdesc").html('<div style="font-weight:bold;color:#5383c2;border-bottom: 1px dotted #5383c2; padding: 3px 0px 3px;">'+showtopic+'</div>'
                                +'<table style="width:100%;font-family:'+"宋体"+'; line-height:28px;"><tr height="40px;"><td colspan="2">'
                                +'<div style="padding:0px 5px;color:#1d5987;font-weight:bold;font-size:9px; text-align:left; font-size:14px;background:#A4C3E3; ">'
                                +'<span style="background:#ff3f3f;width:14px;height:14px;color:#E3E3E3;font-size:10px;position:relative;left:0;top:0;font-size:14px;">'+event.confshortname
                                +'</span>&nbsp;'+event.confname+'值班'+event.fullname
                                   +'</div></td></tr>'
                                   +'<tr height="40px;"><td style="width:70px;color:#4b4b4b;">'+'值班人员'+'</td><td style="padding-left:10px;">'+event.groupPersons+'</td></tr>'
                                   +'<tr height="40px;"><td valign="top" style="width:70px;color:#4b4b4b;">'+'值班备注'+'</td><td style="padding-left:10px;"><textarea readonly rows=3 cols=35>'+event.description+'</textarea></td></tr>'
                                +'</table>');
                                
            
            $( "#reserveinfo" ).dialog( 
                { title:  fstart + "-" + fend + " " + showtopic }
            );
            
            $( "#reserveinfo" ).dialog( "open" );
            return false;
      },
      loading: function(bool) {
            if (bool) $('#loading').show();
            else $('#loading').hide();
      },
      eventMouseover: function(calEvent, jsEvent, view) {
            var fstart  = $.fullCalendar.formatDate(calEvent.start, "yyyy/MM/dd HH:mm");
            var fend  = $.fullCalendar.formatDate(calEvent.end, "HH:mm");    
            $(this).attr('title', fstart + " - " + fend + " " + calEvent.topic + " : " + calEvent.description);
            $(this).css('font-weight', 'normal');                
            $(this).tooltip({
                effect:'toggle',
                cancelDefault: true
            });
      },
      eventMouseout: function(calEvent, jsEvent, view) {
            $(this).css('font-weight', 'normal');
      },
      eventRender: function(event, element) {
        var fstart  = $.fullCalendar.formatDate(event.start, "HH:mm");
        var fend  = $.fullCalendar.formatDate(event.end, "HH:mm");    
        // Bug in IE8
        // element.html('<a href=#>' + fstart + "-" + fend + '<div
        // style=color:#E5E5E5>' + event.title + "</div></a>");
      },
      eventAfterRender : function(event, element, view) {
        // alert($.fullCalendar.formatDate($('#calendar').fullCalendar('getView').start,"yyyy-MM-dd"),);
        var fstart  = $.fullCalendar.formatDate(event.start, "HH:mm");
        var fend  = $.fullCalendar.formatDate(event.end, "HH:mm");        
        // element.html('<a href=#><div>Time: ' + fstart + "-" + fend +
        // '</div><div>Room:' + event.confname + '</div><div
        // style=color:#E5E5E5>Host:' + event.fullname + "</div></a>");
        
        
        var confbg='';
        if(event.confid==1){
            confbg = confbg + '<span class="fc-event-bg">什么？</span>';
        }else if(event.confid==2){
            confbg = confbg + '<span class="fc-event-bg">什么？</span>';
        }else if(event.confid==3){
            confbg = confbg + '<span class="fc-event-bg">什么？</span>';
        }else if(event.confid==4){
            confbg = confbg + '<span class="fc-event-bg">什么？</span>';
        }else if(event.confid==5){
            confbg = confbg + '<span class="fc-event-bg">什么？</span>';
        }else if(event.confid==6){
            confbg = confbg + '<span class="fc-event-bg">什么？</span>';
        }else{
            confbg = confbg + '<span class="fc-event-bg">什么？</span>';
        }
        
        var titlebg =  '<span class="fc-event-conf" style="background:'+  event.confcolor +'">' + event.confshortname + '</span>';
        
        // if(event.repweeks>0){
        // titlebg = titlebg + '<span class="fc-event-conf"
        // style="background:#fff;top:0;right:15;color:#3974BC;font-weight:bold">R</span>';
        // }
        if(view.name=="month"){
            var evtcontent = '<div class="fc-event-vert"><a>';
            evtcontent = evtcontent + confbg;
            evtcontent = evtcontent + '<span class="fc-event-titlebg">' +  fstart + " - " +  fend + titlebg + '</span>';
            evtcontent = evtcontent + '<span>班次: ' +  event.confname + '&nbsp;&nbsp;&nbsp;&nbsp;班组:' +  event.fullname + '</span>';
            //evtcontent = evtcontent + '<span>班组: ' +  event.fullname + '</span>';
            evtcontent = evtcontent + '</a><div class="ui-resizable-handle ui-resizable-e"></div></div>';
            element.html(evtcontent);
        }else if(view.name=="agendaWeek"){
            var evtcontent = '<a>';
            evtcontent = evtcontent + confbg;
            evtcontent = evtcontent + '<span class="fc-event-time">' +  fstart + "-" +  fend + titlebg + '</span>';
            evtcontent = evtcontent + '<span>' +  event.confname + ' by ' + event.fullname + '</span>';
            // evtcontent = evtcontent + '<span>' + event.fullname + '</span>';
            evtcontent = evtcontent + '</a><span class="ui-icon ui-icon-arrowthick-2-n-s"><div class="ui-resizable-handle ui-resizable-s"></div></span>';
            element.html(evtcontent);                        
        }else if(view.name=="agendaDay"){
            var evtcontent = '<a>';
            evtcontent = evtcontent + confbg;
            evtcontent = evtcontent + '<span class="fc-event-time">' +  fstart + " - " +  fend + titlebg + '</span>';
            evtcontent = evtcontent + '<span>班次: ' +  event.confname + '</span>';
            evtcontent = evtcontent + '<span>班组: ' +  event.fullname + '</span>';
            evtcontent = evtcontent + '<span>组员: ' +  event.groupPersons + '</span>';
            evtcontent = evtcontent + '</a><span class="ui-icon ui-icon-arrow-2-n-s"><div class="ui-resizable-handle ui-resizable-s"></div></span>';
            element.html(evtcontent);                                
        }
      },
      eventDragStart: function( event, jsEvent, ui, view ) {
        ui.helper.draggable("option", "revert", true);
      },
      eventDragStop: function( event, jsEvent, ui, view ) {
      },
      eventDrop: function( event, dayDelta, minuteDelta, allDay, revertFunc, jsEvent, ui, view ) { 
        if(1==1||2==event.uid){
            var schdata = {startdate:event.start, enddate:event.end, confid:event.confid, sid:event.sid};
        }else{
            revertFunc();
        }
        
      },
      eventResizeStart:  function( event, jsEvent, ui, view ) {
      
        alert('开始调整大小');
      
      },
      eventResize: function(event,dayDelta,minuteDelta,revertFunc) {

        if(1==1||2==event.uid){
            var schdata = {startdate:event.start, enddate:event.end, confid:event.confid, sid:event.sid};

        }else{
            revertFunc();
        }

      }

    });
    
    // goto date function
    if($.browser.msie){// 转到某一日
        $("#calendar .fc-header-right table td:eq(0)").before('<td><div class="ui-state-default ui-corner-left ui-corner-right" style="border-right:0px;padding:1px 3px 2px;" >'
                                                                    +'<input type="text" id="selecteddate" size="10" style="padding:0px;"></div></td>'
                                                             +'<td><div class="ui-state-default ui-corner-left ui-corner-right">'
                                                                     +'<a><span id="selectdate" class="ui-icon ui-icon-search">goto</span></a></div></td>'
                                                             +'<td><span class="fc-header-space"></span></td>');
    }else{
        $("#calendar .fc-header-right table td:eq(0)").before('<td><div class="ui-state-default ui-corner-left ui-corner-right" style="border-right:0px;padding:3px 2px 4px;" >'
                                                                    +'<input type="text" id="selecteddate" size="10" style="padding:0px;"></div></td>'
                                                             +'<td><div class="ui-state-default ui-corner-left ui-corner-right">'
                                                                     +'<a><span id="selectdate" class="ui-icon ui-icon-search">goto</span></a></div></td>'
                                                             +'<td><span class="fc-header-space"></span></td>');
    }
    
    $("#selecteddate").datepicker({
        dateFormat:'yy-mm-dd',
        beforeShow: function (input, instant) {  
            setTimeout(
                function () {
                    $('#ui-datepicker-div').css("z-index", 15);
                }, 100
            );
        }
    });
                    

    
    $("#selectdate").click(function() {
        var selectdstr =     $("#selecteddate").val();    
        var selectdate = $.fullCalendar.parseDate(selectdstr, "yyyy-mm-dd");
        alert(selectdate.getFullYear()+"年"+selectdate.getMonth()+"月"+selectdate.getDate()+"日");
        var now = new Date();
        alert(now.getFullYear()+"年"+now.getMonth()+"月"+now.getDate()+"日");//月份0开始
        $('#calendar').fullCalendar( 'gotoDate', selectdate.getFullYear(), selectdate.getMonth(), selectdate.getDate());
    });
        
    //var view = $('#calendar').fullCalendar('getView');
    //alert("The view's title is " + view.title);
});
</SCRIPT>

    <!-- 日历背景 -->
    <DIV id=calendar></DIV>
    <!-- 明细弹出窗 -->
    <DIV id=reserveinfo title=Details>
    <DIV id=revtitle></DIV>
    <DIV id=revdesc></DIV></DIV>
    <!-- 新增弹出窗 -->
    <DIV style="DISPLAY: none" id=reservebox>
    <FORM id="reserveformID" action="$!{rc.contextPath}/duty/dutyScheduling/save">
      <DIV class=sysdesc>&nbsp;</DIV>
        <DIV class=rowElem><LABEL>日期:</LABEL> 
            <input type="text" id="dutyDate" name="dutyDate"  class="ui-input Wdate validate[required]" readonly="readonly"  onclick="WdatePicker({dateFmt:'yyyy-MM-dd'})" value="$!{dutyScheduling.dutyDate}"  size=22 />
        </DIV>
        <DIV class=rowElem><LABEL>班次:</LABEL> 
            <select name="orderId" id="orderId" class=validate[required]>
                    <option value="">请选择班次</option>
                    #foreach($!list in $!listDutyOrder)
                    <option value="$!{list.id}">$!{list.orderName}</option>
                    #end
             </select>
        </DIV>
        <DIV class=rowElem><LABEL>班组:</LABEL> 
            <input class="ui-input" type="hidden" readonly="readonly" id="groupId" name="groupId" value="$!{dutyScheduling.groupId}"  />
            <input class="ui-input validate[required]" type="text" readonly="readonly" id="groupName" name="groupName" value="$!{dutyScheduling.groupName}"  />
        </DIV>
        <DIV class=rowElem><LABEL>备注:</LABEL> 
            <TEXTAREA id="remark" rows=3 cols=43 name=remark>$!{dutyScheduling.remark}</TEXTAREA> 
        </DIV>
        <DIV class=rowElem> </DIV>
        <DIV class=rowElem> </DIV>
        <DIV id=addhelper class=ui-widget>
            <DIV 
            style="PADDING-BOTTOM: 5px; PADDING-LEFT: 5px; PADDING-RIGHT: 5px; PADDING-TOP: 5px" 
            class="ui-state-error ui-corner-all">
            <DIV id=addresult></DIV>
            </DIV>
        </DIV>
    </FORM>
    </DIV>
</DIV>

</BODY>

</html>
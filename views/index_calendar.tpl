<!-- iframe里日历-->
<!DOCTYPE html>
{{template "header"}}
<title>EngineerCMS</title>

<link rel='stylesheet' href='/static/css/fullcalendar.min.css' />
<!-- <script src='/static/js/jquery-2.1.3.min.js'></script> -->
  <!-- <script type="text/javascript" src="/static/js/bootstrap.min.js"></script> -->
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/> -->
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

	/*body {
		margin: 40px 10px;
		padding: 0;
		font-family: "Lucida Grande",Helvetica,Arial,Verdana,sans-serif;
		font-size: 14px;
	}*/
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
			editable: false,
			events: {  
        		url: '/admin/calendar',  
        		type: 'post'  
    		},
			loading: function(bool) {
				$('#loading').toggle(bool);
			},
    });
  });
</script>
<body>
  <div class="navbar navba-default navbar-fixed-top">
    <div class="container-fill">{{template "navbar" .}}</div>
  </div>
<!-- <div class="col-lg-12"> -->
	<div id='calendar'></div>
<!-- </div> -->
</body>
</html>
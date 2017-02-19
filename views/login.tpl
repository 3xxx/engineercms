<!-- 用户登录页面 -->
 <!DOCTYPE html>
<html>
<head>
 <meta charset="UTF-8">
  <title>EngineerCMS</title>
<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
 <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
 <!-- <script src="/static/js/bootstrap-treeview.js"></script> -->
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
</head>
<body>

<div id="content" class="col-md-8 col-md-offset-2">
    <div class="col-md-6 auth-page">
      <h3 class="title">
        <span class="glyphicon glyphicon-user"></span>
        登录
      </h3>
    <form method="POST" action="/login">
        <span style="color: #ff0000;"><input type="hidden" name="url" value="{{.Url}}"/></span>
      <div class="form-group">
        <label class="control-label" for="LoginForm-UserName">用户名 或 邮箱</label>
        <input id="uname" name="uname" type="text" value="" class="form-control" placeholder="Enter account" list="cars"></div>
        <div id='datalistDiv'>
          <datalist id="cars" name="cars">
          </datalist>
        </div>
      <div class="form-group">
        <label class="control-label" for="LoginForm-Password">密码</label>
        <input id="pwd" name="pwd" type="password" value="" class="form-control" placeholder="Password"></div>

      <div class="checkbox">
        <label>
          <input type="checkbox">自动登陆</label>
      </div>
      <button type="submit" class="btn btn-default" onclick="return checkInput();">
        登录&nbsp;&nbsp; <span class="glyphicon glyphicon-circle-arrow-right"></span>
      </button>
      <a href="./forgot" class="pull-right">
        <span class="glyphicon glyphicon-question-sign"></span>
        忘记密码
      </a>
      <button class="btn btn-default" onclick="return backToHome();">返回&nbsp;&nbsp; <span class="glyphicon glyphicon-circle-arrow-left"></span></button>

  </form>
</div>

<div class="col-md-6 auth-page">
  <div class="auth-page">
    <h3 class="title">
      <span class="glyphicon glyphicon-question-sign"></span>
      帮助
    </h3>
    <p class="well">如果您还没有注册帐户的话，请先注册。</p>
    <p>
      <a href="./regist" class="btn btn-default">
        立即注册&nbsp;&nbsp; <i class="icon-chevron-sign-right"></i>
      </a>
    </p>
  </div>
</div>

</div>

<script type="text/javascript">
function checkInput(){
  var uname=document.getElementById("uname");
  if (uname.value.length==0){
    alert("请输入账号");
    return false;
  }
    var pwd=document.getElementById("pwd");
  if (pwd.value.length==0){
    alert("请输入密码");
    return false;
    }
    return true
}
function backToHome(){
  window.location.href="/";
  return false;
}
</script>

<script type="text/javascript">
//禁用chrome和firefox浏览器自带的自动提示 
$('#uname').attr("autocomplete","off"); 
$(document).ready(function(){
  $("#uname").keyup(function(event){
    var uname1=document.getElementById("uname");
    // alert(event.keyCode);
   if (event.keyCode != 38 && event.keyCode != 40 && uname1.value.length==2){
    $.ajax({
                type:"post",//这里是否一定要用post？？？
                url:"/regist/getuname",
                data: { uname: $("#uname").val()},
                dataType:'json',//dataType:JSON,这种是jquerylatest版本的表达方法。不支持新版jquery。
                success:function(data,status){
                  $(".option").remove();
                  $.each(data,function(i,d){
                      $("#cars").append('<option class="option" value="' + data[i].Username + '">' + data[i].Nickname + '</option>');
                  });
                }
      });
                // $("#uname").keydown(function(){
                //   $("option").remove();
                // }); 
    }
 });
});   
  // $("#uname").focus(function(){
  // $("#uname").blur(function(){
  // $("#uname").keydown(function(){用keydown不行，因为当按键放开后才传值
  
    
  // $(function(){
            
                //控制器中用jsonserver的话就用下面这个形式，不用eval转换，如果用writestring方式，就要用eval转换。
                      // 方式一：this.Ctx.WriteString(string(b))
                      //document.write(data);
                      // [{"Id":533,"Username":"zhong.cq","Password":"2d243e65cb10f17d2a953ea3bb539449","Repassword":"","Nickname":"钟春强","Email":"zhong.cq@gpdiwe.com","Remark":"","Status":0,"Lastlogintime":"2015-08-17T15:18:41.2317219Z","Createtime":"2015-08-17T15:18:41.2327219Z","Roles":null},{"Id":535,"Username":"cai.sq","Password":"f30db3e6af22028c4cf078e6dcc69050","Repassword":"","Nickname":"蔡顺情","Email":"cai.sq@gpdiwe.com","Remark":"","Status":0,"Lastlogintime":"2015-08-17T15:18:42.1157724Z","Createtime":"2015-08-17T15:18:42.1157724Z","Roles":null},
                  //方式一：
                          // var obj = eval(data);//将json字符串转换成对象?                  
                          //  $.each(obj,function(i,d){
                          //    $("#cars").append('<option value="' + obj[i].Username + '"></option>');
                          //   }); 

                      
                      //方式二：
                      // 控制器里：this.Data["json"] = string(b)
                      // this.ServeJSON()
                      // document.write(data);
                      // "[{\"Id\":533,\"Username\":\"zhong.cq\",\"Password\":\"2d243e65cb10f17d2a953ea3bb539449\",\"Repassword\":\"\",\"Nickname\":\"钟春强\",\"Email\":\"zhong.cq@gpdiwe.com\",\"Remark\":\"\",\"Status\":0,\"Lastlogintime\":\"2015-08-17T15:18:41.2317219Z\",\"Createtime\":\"2015-08-17T15:18:41.2327219Z\",\"Roles\":null},{\"Id\":535,\"Username\":\"cai.sq\",\"Password\":\"f30db3e6af22028c4cf078e6dcc69050\",\"Repassword\":\"\",\"Nickname\":\"蔡顺情\",\"Email\":\"cai.sq@gpdiwe.com\",\"Remark\":\"\",\"Status\":0,\"Lastlogintime\":\"2015-08-17T15:18:42.1157724Z\",\"Createtime\":\"2015-08-17T15:18:42.1157724Z\",\"Roles\":null},

                      // var obj = eval(data);
                      // document.write(obj);
                      // [{"Id":533,"Username":"zhong.cq","Password":"2d243e65cb10f17d2a953ea3bb539449","Repassword":"","Nickname":"钟春强","Email":"zhong.cq@gpdiwe.com","Remark":"","Status":0,"Lastlogintime":"2015-08-17T15:18:41.2317219Z","Createtime":"2015-08-17T15:18:41.2327219Z","Roles":null},{"Id":535,"Username":"cai.sq","Password":"f30db3e6af22028c4cf078e6dcc69050","Repassword":"","Nickname":"蔡顺情","Email":"cai.sq@gpdiwe.com","Remark":"","Status":0,"Lastlogintime":"2015-08-17T15:18:42.1157724Z","Createtime":"2015-08-17T15:18:42.1157724Z","Roles":null},
                      // alert(data);
                       // var obj = eval(data);//将json字符串转换成对象? 去掉斜杠\                 
                         // alert(obj); 
                          // document.write(obj);
                         // var obj1=eval(obj) //再转成json对象？
                         // alert(obj1[0].Username);只有firefox支持
                           
                           // for (i=0;i<=10;i++){取10个也不行
                           //  $("#cars").append('<option value="'+obj1[i].Username+'"></option>');
                           // }

                         // document.getElementById('datalistDiv').innerHTML=obj[i].Username;
                           // $("ul").append("<li>"+d.Username+"</li>");
                           // document.write(d.Username)
                          // var returned = xmlhttp.responseText;
                          // 接着,你就可以这样访问:
                          // var person1 = obj[0]; var person2 = obj[1];
                          // alert(person1.Username);
                          // alert(person2.Username);
                          // document.write(data);
                          // var i=0
                          // for (i=0;i<=obj.length;i++)
                          // {
                            // $("#cars").append('<option value="' + obj[i].Id + '">' + obj[i].Username + '</option>');
                            // $("ul").append("<li>"+obj[i].Username+"</li>");
                          // document.write(obj[i].Username)
                          // document.write("<br />")
                          // }
                          // $.each(data.suburbs,function(i,v) 
                          //   { 
                          //     $suburbs.append('<option data-value=' + i + ' data-postcode=' + v.id + ' data-state=' + v.Username + '>' + i + '</option>' );
                          //     });
                        //   $.each(data.list, function(key, val) {
                        //     $("#cars").append('<option value="' + val.id + '">' + val.Username + '</option>');
                        // });
                 //装载数据
                     // document.getElementById('datalistDiv').innerHTML=obj[i].Username;
                    // $("ul").append("<li>"+d.Username+"</li>");
                    // document.write(d.Username)
                    //         $.each(data.list, function (i, item) {
                    //     $("#cars").append("<option value='" + item.id + "'>" + item.Username> +"</option>");
                    //     document.write(item.Username)
                    //     // alert(item.id)
                    // });
                          // var x
                          // var mycars = new Array()
                          // mycars[0] = "Saab"
                          // mycars[1] = "Volvo"
                          // mycars[2] = "BMW"
                          
                          // for (x in mycars)
                          // {
                          // document.write(mycars[x] + "<br />")
                          // }
                         // }

// $('input').keydown(function(event){ 
   // alert(event.keyCode); 
 // });
// 上面代码中的,event.keyCode就可以帮助我们获取到我们按下了什么按键,他返回的是ascII码,比如说上下左右键,分别是38,40,37,39
// $(document).keypress(function(e) { 
      // if (e.ctrlKey && e.which == 13) 
      // $("form").submit(); 
 // })

</script>


<img src="data:image/png;base64,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
QmCC" width="10%" align ="left"> 

</body>
</html>
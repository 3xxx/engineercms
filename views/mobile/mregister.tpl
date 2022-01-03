<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <title>EngineerCMS - A Fresh Mobile System</title>
  <meta name="viewport" content="width=device-width, initial-scale=1  maximum-scale=1 user-scalable=no">
  <meta name="mobile-web-app-capable" content="yes">
  <meta name="apple-mobile-web-app-capable" content="yes">
  <meta name="apple-touch-fullscreen" content="yes">
  <meta name="HandheldFriendly" content="True">
  <link rel="stylesheet" href="/static/voxes/css/materialize.css">
  <link rel="stylesheet" href="/static/voxes/font-awesome/css/font-awesome.min.css">
  <link rel="stylesheet" href="/static/voxes/css/normalize.css">
  <link rel="stylesheet" href="/static/voxes/css/owl.carousel.css">
  <link rel="stylesheet" href="/static/voxes/css/owl.theme.css">
  <link rel="stylesheet" href="/static/voxes/css/owl.transitions.css">
  <link rel="stylesheet" href="/static/voxes/css/fakeLoader.css">
  <link rel="stylesheet" href="/static/voxes/css/style.css">
  <link rel="shortcut icon" href="/static/voxes/img/favicon.png">
</head>

<body class="login-register-wrap">
  <!-- login register -->
  <div class="login-register-wrap-home">
    <div class="container">
      <div class="content">
        <h1>E</h1>
        <h6>EmgineerCMS</h6>
        <form method="POST" action="/regist">
          <input type="text" id="uname" name="uname" placeholder="Full Name" class="validate" required>
          <input type="text" id="nickname" name="nickname" placeholder="Nick Name" class="validate" required>
          <input type="email" id="email" name="email" placeholder="Email" class="validate" required>
          <input type="password" id="pwd" name="pwd" placeholder="password" class="validate" required>
          <button class="button-default" type="submit" οnclick="return checkInput()">Sing up</button>
          <h6>Already hava an account ? <a href="/login">Sign in</a></h6>
        </form>
      </div>
    </div>
  </div><!-- end login register -->
  <!-- scripts -->
  <script src="/static/voxes/js/jquery.min.js"></script>
  <script src="/static/voxes/js/materialize.min.js"></script>
  <script src="/static/voxes/js/owl.carousel.min.js"></script>
  <script src="/static/voxes/js/contact-form.js"></script>
  <script src="/static/voxes/js/fakeLoader.min.js"></script>
  <script src="/static/voxes/js/main.js"></script>
  <script type="text/javascript">
  $("#uname").blur(function() {
    $.ajax({
      type: "post",
      url: "/regist/checkuname",
      data: { uname: $("#uname").val() },
      success: function(data, status) {
        if ($("#uname").val().length > 0 && data == "false") {
          // $("#regis").html("用户名已存在！");
          // $("#regis").css("color", "red");
          Materialize.toast('用户名已存在！', 3000, 'rounded')
        } else if ($("#uname").val().length > 0 && data == "true") {
          // $("#regis").html("用户名可用！");
          // $("#regis").css("color", "green");
          Materialize.toast('用户名可用！', 3000, 'rounded')
        } else if ($("#uname").val().length == 0) {
          // $("#regis").html("用户名");
          // $("#regis").css("color", "black");
          Materialize.toast('用户名', 3000, 'rounded')
        }
      }
    });
  });
  // });

  function checkInput() {
    var uname = document.getElementById("uname");
    // if (uname.value.length == 0) {
    //   alert("请输入账号");
    //   return false;
    // }
    var pwd = document.getElementById("pwd");
    // if (pwd.value.length == 0) {
    //   alert("请输入密码");
    //   return false;
    // }
    if ($("#uname").val().length > 0 && $("#regis").html() == "用户名已存在！") {
      Materialize.toast('用户名已存在！', 3000, 'rounded')
      return false;
    }
  }
  </script>
</body>

</html>
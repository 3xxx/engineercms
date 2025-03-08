<!DOCTYPE html>
<html lang="en" xmlns:th="http://www.w3.org/1999/xhtml">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <script src="/static/vue.js/js/jquery.min.js"></script>
  <script src="/static/vue.js/qrcode2.min.js"></script>
  <title>微信扫码登录</title>
  <style>
  * {
    padding: 0;
    margin: 0;
  }

  body {
    /*background:  aqua url("/images/bodybg.png");*/
  }

  #container {
    margin: 0px auto;
    display: flex;
    justify-content: center;
    margin-top: 100px;
    /*color: rgb(140,0,0);*/
    font-size: 20px;
  }

  #container span {
    font-family: "楷体";
  }

  #bottom {
    display: flex;
    /*margin: 0 auto;*/
    justify-content: space-around;
  }

  #bottom button {
    padding: 5px 20px;
    background-color: rgb(140, 0, 0);
    color: #ffffff;
    font-size: 18px;
    border: none;
    border-radius: 5px;
  }

  #bottom button:hover {
    padding: 5px 20px;
    background-color: #ffffff;
    color: rgb(140, 0, 0);
    font-size: 18px;
    border: 2px solid rgb(140, 0, 0);
    border-radius: 5px;
  }

  #emPower {
    margin-right: 20px;
  }

  #right,
  #left {
    width: 230px;
    height: 230px;
    padding: 10px;
    /*background: url("/static/images/4ww.jpg");*/
  }

  /*.qrCodeImgId {
    display: none;
  }*/
  </style>
</head>

<body>
  <div id="container">
    <div id="bottom">
      <div>
        <!-- <input type="hidden" id="projectId" value="{{.donateVo.projectId}}">
        <input type="hidden" id="projectName" value="{{.donateVo.projectName}}">
        <input type="hidden" id="payMoney" value="{{.donateVo.payMoney}}">
        <input type="hidden" id="projectPicture" value="{{.donateVo.projectPicture}}"> -->
        <!-- <button id="emPower" onclick="getQrCode()">扫码授权</button> -->
        <button id="emPower" onclick="getQrCode()">扫码授权</button>
        <button id="cancelEmPower">跳过授权</button>
      </div>
    </div>
  </div>
  <div id="container">
    <div id="right">
      <h4>温馨提示：</h4>
      <span>微信扫码登录页面，首次扫描会自动注册。扫码前先微信关注下列测试公众号</span>
      <img src="/static/img/测试公众号二维码.png">
    </div>
    <div id="left">
      <div id="qrcode" style="width:300px; height:300px; margin-top:15px;"></div>
    </div>
  </div>
  <script>
  // 存储二维码标识,用于验证是否扫码成功
  // var sceneStr;
  var t;
  var qrcode = new QRCode(document.getElementById("qrcode"), {
    width: 230,
    height: 230,
    foreground: "#000000",
    background: "#ffffff",
    typeNumber: QRCode.CorrectLevel.L
  });

// QRCode.CorrectLevel.LMQH
  // function makeCode() {
  //   var elText = document.getElementById("text");

  //   if (!elText.value) {
  //     alert("Input a text");
  //     elText.focus();
  //     return;
  //   }

  //   qrcode.makeCode(elText.value);
  // }

  // makeCode();

  // $("#text").
  // on("blur", function() {
  //   makeCode();
  // }).
  // on("keydown", function(e) {
  //   if (e.keyCode == 13) {
  //     makeCode();
  //   }
  // });
  // 获取登录二维码
  function getQrCode() {
    $.ajax({
      url: '/v1/wx/wxloginqrcode',
      success: function(data) {
        console.log(data);
        if (data.code == 200) {
          if (data.CodeUrl && data.CodeUrl !== '') {
            if (qrcode !== null){
              qrcode.clear();
            }
            qrcode.makeCode(data.CodeUrl);
          }
          // sceneStr = data.data.sceneStr;
          // $('#qrCodeImgId').attr('src', "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" + data.data.ticket);
          // $('#qrCodeImgId').show();
          t = window.setInterval(getOpenId, 2000);
        } else {
          alert(data.msg);
        }
      }
    });
  };
  // 扫码成功，获取用户openId=>为获取用户信息做准备
  function getOpenId() {
    $.ajax({
      url: '/v1/wx/getwxuserlogin',
      type: "GET",
      // dataType: "json",
      success: function(data) {
        if (data.code == 200) {
          console.log(data),
          console.log("========getOpenId==========");
          console.log(data.data);
          window.clearInterval(t);
          window.location.href = "/" // + data.data.openid
          alert("登录成功");
          /**
           * 1、第一次扫码登录进行账号绑定
           * 2、以后根据openId获取用户信息
           */
        }
      }
    });
  };

  $("#cancelEmPower").on("click", function() {
    var data = {
      payMoney: $("#payMoney").val(),
      projectName: $("#projectName").val(),
      projectId: $("#projectId").val(),
      projectPicture: $("#projectPicture").val()
    };
    window.location.href = "/"; // + $.param(data)
  });
  </script>
</body>

</html>
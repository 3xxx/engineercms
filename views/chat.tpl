<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <title>Simple Chat</title>
  <!--<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.97.8/css/materialize.min.css">
    <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/emojione/2.2.6/assets/css/emojione.min.css"/>
    <script src="https://twemoji.maxcdn.com/v/latest/twemoji.min.js" crossorigin="anonymous"></script>
     -->
  <link rel="stylesheet" href="/static/chat/materialize.min.css">
  <link rel="stylesheet" href="/static/chat/icon?family=Material+Icons">
  <link rel="stylesheet" href="/static/chat/emojione.min.css" />
  <link rel="stylesheet" href="/static/chat/style.css">
  <script src="/static/js/twemoji.min.js" crossorigin="anonymous"></script>
  <style>
    ul.emoji-list * {
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
  }

  ul.emoji-list li {
    font-size: 36px;
    float: left;
    display: inline-block;
    padding: 2px;
    margin: 4px;
  }

  img.emoji {
    cursor: pointer;
    height: 1em;
    width: 1em;
    margin: 0 .05em 0 .1em;
    vertical-align: -0.1em;
  }
  </style>
</head>

<body>
  <header>
    <nav>
      <div class="nav-wrapper">
        <a href="/v1/chat/chat" class="brand-logo right">Simple Chat</a>
      </div>
    </nav>
  </header>
  <main id="app">
    <div class="row">
      <div class="col s12">
        <div class="card horizontal">
          <div id="chat-messages" class="card-content" v-html="chatContent">
          </div>
        </div>
      </div>
    </div>
    <div class="row" v-if="joined">
      <div class="input-field col s8">
        <input type="text" v-model="newMsg" v-html="chatMsg" id="chat" @keyup.enter="send">
      </div>
      <div class="input-field col s4">
        <button class="waves-effect waves-light btn" @click="send">
          <i class="material-icons right">chat</i>
          Send
        </button>
      </div>
    </div>
    <div class="row" v-if="!joined">
      <div class="input-field col s8">
        <input type="email" v-model.trim="email" placeholder="Email">
      </div>
      <div class="input-field col s8">
        <input type="text" v-model.trim="username" placeholder="Username">
      </div>
      <div class="input-field col s4">
        <button class="waves-effect waves-light btn" @click="join()">
          <i class="material-icons right">done</i>
          Join
        </button>
      </div>
    </div>
  </main>
  <footer class="page-footer">
    <ul class="emoji-list">
      <li>&#x1F004;&#xFE0F;</li>
      <li>&#x1F600;</li>
      <li>&#x1F601;</li>
      <li>&#x1F602;</li>
      <li>&#x1F603;</li>
      <li>&#x1F604;</li>
      <li>&#x1F605;</li>
      <li>&#x1F606;</li>
      <li>&#x1F607;</li>
      <li>&#x1F608;</li>
      <li>&#x1F609;</li>
      <li>&#x1F60A;</li>
      <li>&#x1F60B;</li>
      <li>&#x1F60C;</li>
      <li>&#x1F60D;</li>
      <li>&#x1F60E;</li>
      <li>&#x1F60F;</li>
      <li>&#x1F610;</li>
      <li>&#x1F611;</li>
      <li>&#x1F612;</li>
      <li>&#x1F613;</li>
      <li>&#x1F614;</li>
      <li>&#x1F615;</li>
      <li>&#x1F616;</li>
      <li>&#x1F617;</li>
      <li>&#x1F618;</li>
      <li>&#x1F619;</li>
      <li>&#x1F61A;</li>
      <li>&#x1F61B;</li>
      <li>&#x1F61C;</li>
      <li>&#x1F61D;</li>
      <li>&#x1F61E;</li>
      <li>&#x1F61F;</li>
      <li>&#x1F620;</li>
      <li>&#x1F621;</li>
      <li>&#x1F622;</li>
      <li>&#x1F623;</li>
      <li>&#x1F624;</li>
      <li>&#x1F625;</li>
      <li>&#x1F626;</li>
      <li>&#x1F627;</li>
      <li>&#x1F628;</li>
      <li>&#x1F629;</li>
      <li>&#x1F62A;</li>
      <li>&#x1F62B;</li>
      <li>&#x1F62C;</li>
      <li>&#x1F62D;</li>
      <li>&#x1F62E;</li>
      <li>&#x1F62F;</li>
      <li>&#x1F630;</li>
      <li>&#x1F631;</li>
      <li>&#x1F632;</li>
      <li>&#x1F633;</li>
      <li>&#x1F634;</li>
      <li>&#x1F635;</li>
      <li>&#x1F636;</li>
      <li>&#x1F637;</li>
      <li>&#x1F638;</li>
      <li>&#x1F639;</li>
      <li>&#x1F63A;</li>
      <li>&#x1F63B;</li>
      <li>&#x1F63C;</li>
      <li>&#x1F63D;</li>
      <li>&#x1F63E;</li>
      <li>&#x1F63F;</li>
      <li>&#x1F640;</li>
      <li>&#x1F641;</li>
      <li>&#x1F642;</li>
      <li>&#x1F643;</li>
      <li>&#x1F644;</li>
      <li>&#x2639;&#xFE0F;</li>
      <li>&#x263A;&#xFE0F;</li>
      <li>&#x1F910;</li>
      <li>&#x1F911;</li>
      <li>&#x1F912;</li>
      <li>&#x1F913;</li>
      <li>&#x1F914;</li>
      <li>&#x1F915;</li>
      <li>&#x1F916;</li>
      <li>&#x1F917;</li>
      <li>&#x1F920;</li>
      <li>&#x1F921;</li>
      <li>&#x1F922;</li>
      <li>&#x1F923;</li>
      <li>&#x1F924;</li>
      <li>&#x1F925;</li>
      <li>&#x1F927;</li>
      <li>&#x1F928;</li>
      <li>&#x1F929;</li>
      <li>&#x1F92A;</li>
      <li>&#x1F92B;</li>
      <li>&#x1F92C;</li>
      <li>&#x1F92D;</li>
      <li>&#x1F92E;</li>
      <li>&#x1F92F;</li>
      <li>&#x1F970;</li>
      <li>&#x1F971;</li>
      <li>&#x1F973;</li>
      <li>&#x1F974;</li>
      <li>&#x1F975;</li>
      <li>&#x1F976;</li>
      <li>&#x1F97A;</li>
      <li>&#x1F90F;</li>
      <li>&#x1F44C;&#x1F3FB;</li>
      <!-- <li>&#x1F44C;&#x1F3FC;</li>
    <li>&#x1F44C;&#x1F3FD;</li>
    <li>&#x1F44C;&#x1F3FE;</li>
    <li>&#x1F44C;&#x1F3FF;</li> -->
      <!-- <li>&#x1F44C;</li> -->
      <li>&#x1F44D;&#x1F3FB;</li>
      <!-- <li>&#x1F44D;&#x1F3FC;</li>
    <li>&#x1F44D;&#x1F3FD;</li>
    <li>&#x1F44D;&#x1F3FE;</li>
    <li>&#x1F44D;&#x1F3FF;</li> -->
      <!-- <li>&#x1F44D;</li> -->
      <li>&#x1F44E;&#x1F3FB;</li>
      <!-- <li>&#x1F44E;&#x1F3FC;</li>
    <li>&#x1F44E;&#x1F3FD;</li>
    <li>&#x1F44E;&#x1F3FE;</li>
    <li>&#x1F44E;&#x1F3FF;</li> -->
      <!-- <li>&#x1F44E;</li> -->
      <li>&#x1F44F;&#x1F3FB;</li>
      <!-- <li>&#x1F44F;&#x1F3FC;</li>
    <li>&#x1F44F;&#x1F3FD;</li>
    <li>&#x1F44F;&#x1F3FE;</li>
    <li>&#x1F44F;&#x1F3FF;</li> -->
      <!-- <li>&#x1F44F;</li> -->
      <!-- <li>&#x1F918;&#x1F3FB;</li>
    <li>&#x1F918;&#x1F3FC;</li>
    <li>&#x1F918;&#x1F3FD;</li>
    <li>&#x1F918;&#x1F3FE;</li>
    <li>&#x1F918;&#x1F3FF;</li>
    <li>&#x1F918;</li>
    <li>&#x1F919;&#x1F3FB;</li>
    <li>&#x1F919;&#x1F3FC;</li>
    <li>&#x1F919;&#x1F3FD;</li>
    <li>&#x1F919;&#x1F3FE;</li>
    <li>&#x1F919;&#x1F3FF;</li>
    <li>&#x1F919;</li>
    <li>&#x1F91A;&#x1F3FB;</li>
    <li>&#x1F91A;&#x1F3FC;</li>
    <li>&#x1F91A;&#x1F3FD;</li>
    <li>&#x1F91A;&#x1F3FE;</li>
    <li>&#x1F91A;&#x1F3FF;</li>
    <li>&#x1F91A;</li>
    <li>&#x1F91B;&#x1F3FB;</li>
    <li>&#x1F91B;&#x1F3FC;</li>
    <li>&#x1F91B;&#x1F3FD;</li>
    <li>&#x1F91B;&#x1F3FE;</li>
    <li>&#x1F91B;&#x1F3FF;</li>
    <li>&#x1F91B;</li>
    <li>&#x1F91C;&#x1F3FB;</li>
    <li>&#x1F91C;&#x1F3FC;</li>
    <li>&#x1F91C;&#x1F3FD;</li>
    <li>&#x1F91C;&#x1F3FE;</li>
    <li>&#x1F91C;&#x1F3FF;</li>
    <li>&#x1F91C;</li>
    <li>&#x1F91D;</li>
    <li>&#x1F91E;&#x1F3FB;</li>
    <li>&#x1F91E;&#x1F3FC;</li>
    <li>&#x1F91E;&#x1F3FD;</li>
    <li>&#x1F91E;&#x1F3FE;</li>
    <li>&#x1F91E;&#x1F3FF;</li>
    <li>&#x1F91E;</li>
    <li>&#x1F91F;&#x1F3FB;</li>
    <li>&#x1F91F;&#x1F3FC;</li>
    <li>&#x1F91F;&#x1F3FD;</li>
    <li>&#x1F91F;&#x1F3FE;</li>
    <li>&#x1F91F;&#x1F3FF;</li>
    <li>&#x1F91F;</li> -->
      <li>&#x261D;&#x1F3FB;</li>
      <!--<li>&#x261D;&#x1F3FC;</li>
    <li>&#x261D;&#x1F3FD;</li>
    <li>&#x261D;&#x1F3FE;</li>
    <li>&#x261D;&#x1F3FF;</li>
    <li>&#x261D;&#xFE0F;</li> -->
      <li>&#x1F926;&#x1F3FB;&#x200D;&#x2640;&#xFE0F;</li>
      <li>&#x1F926;&#x1F3FB;&#x200D;&#x2642;&#xFE0F;</li>
      <li>&#x1F926;&#x1F3FB;</li>
      <li>&#x1F926;&#x1F3FC;&#x200D;&#x2640;&#xFE0F;</li>
      <li>&#x1F926;&#x1F3FC;&#x200D;&#x2642;&#xFE0F;</li>
      <li>&#x1F926;&#x1F3FC;</li>
      <li>&#x1F926;&#x1F3FD;&#x200D;&#x2640;&#xFE0F;</li>
      <li>&#x1F926;&#x1F3FD;&#x200D;&#x2642;&#xFE0F;</li>
      <li>&#x1F926;&#x1F3FD;</li>
      <!-- <li>&#x1F926;&#x1F3FE;&#x200D;&#x2640;&#xFE0F;</li>
    <li>&#x1F926;&#x1F3FE;&#x200D;&#x2642;&#xFE0F;</li>
    <li>&#x1F926;&#x1F3FE;</li>
    <li>&#x1F926;&#x1F3FF;&#x200D;&#x2640;&#xFE0F;</li>
    <li>&#x1F926;&#x1F3FF;&#x200D;&#x2642;&#xFE0F;</li>
    <li>&#x1F926;&#x1F3FF;</li> -->
      <li>&#x1F926;&#x200D;&#x2640;&#xFE0F;</li>
      <li>&#x1F926;&#x200D;&#x2642;&#xFE0F;</li>
      <li>&#x1F926;</li>
      <li>&#x2614;&#xFE0F;</li>
      <li>&#x2615;&#xFE0F;</li>
      <li>&#x2618;&#xFE0F;</li>
      <li>&#x2620;&#xFE0F;</li>
      <li>&#x2622;&#xFE0F;</li>
      <li>&#x2623;&#xFE0F;</li>
      <li>&#x2626;&#xFE0F;</li>
      <li>&#x262A;&#xFE0F;</li>
      <li>&#x262E;&#xFE0F;</li>
      <li>&#x262F;&#xFE0F;</li>
      <li>&#x2638;&#xFE0F;</li>
      <li>&#x3297;&#xFE0F;</li>
      <li>&#x3299;&#xFE0F;</li>
      <li>&#x2764;&#xFE0F;</li>
      <li>&#x1F21A;&#xFE0F;</li>
      <li>&#x1F22F;&#xFE0F;</li>
      <li>&#x1F232;</li>
      <li>&#x1F233;</li>
      <li>&#x1F234;</li>
      <li>&#x1F235;</li>
      <li>&#x1F236;</li>
      <li>&#x1F237;&#xFE0F;</li>
      <li>&#x1F238;</li>
      <li>&#x1F239;</li>
      <li>&#x1F23A;</li>
      <li>&#x1F250;</li>
      <li>&#x1F251;</li>
      <li>&#x1F1E8;&#x1F1F3;</li>
      <li>&#x1F192;</li>
      <li>&#x1F193;</li>
      <li>&#x1F194;</li>
      <li>&#x1F195;</li>
      <li>&#x1F196;</li>
      <li>&#x1F197;</li>
      <li>&#x1F198;</li>
      <li>&#x1F199;</li>
      <li>&#x1F19A;</li>
    </ul>
  </footer>
  <script type="text/javascript">
  var ul = document.getElementsByTagName('ul')[0];
  // var total = ul.getElementsByTagName('li').length;
  // var elapsed = +new Date;
  // twemoji.parse(ul, { "size": 72 });
  twemoji.parse(ul, { "folder": "/static/assets/svg", "ext": ".svg", "base": "" });
  // elapsed = (+new Date) - elapsed;
  // document.body.insertBefore(
  //   document.createTextNode(total + ' emoji parsed in ' + elapsed + 'ms'),
  //   document.body.firstChild
  // );

  (function(img, metaKey, i) {
    // function copyToClipboard(e) {
    //   prompt('Copy to clipboard via ' + metaKey + '+C and Enter', this.alt);
    // }
    function addnewMsg(e) {
      // this.chatMsg +=this.alt;
      // this.newMsg +=this.alt;
      document.getElementById('chat').value += this.alt;
      // document.getElementById('chat').html(this.alt);
      // document.getElementById('chat').innerText += this.alt;
      // document.getElementById('chat').attr("value",this.alt);
      // $('<p>').html(this.newMsg).text()+= this.alt;
      // alert(this.alt);
      // console.log(document.getElementById('chat').value)
      // console.log($('<p>').html(this.newMsg).text())
    }

    // for (i = 0; i < img.length; img[i++].onclick = copyToClipboard) {}
    for (i = 0; i < img.length; img[i++].onclick = addnewMsg) {}
  }(
    document.getElementsByTagName('img'),
    /\b(?:Mac |i)OS\b/i.test(navigator.userAgent) ? 'Command' : 'Ctrl'
  ));

  // var emoji = twemoji.parse('🌽')
  // var emoji2 = twemoji.parse('\ud83c\udf3d')
  // console.log(emoji)
  // console.log(emoji2)

  // var el = document.createElement('div')
  // el.innerHTML = emoji
  // document.body.appendChild(el)

  // var emoji = twemoji.parse('\ud83c\udf3d', {
  //   base: '/static/js/twemoji.min.js',
  //   folder: '/svg', // 子路径，可选。 folder 会拼接到 base 后面，也可以都写在 base 上，不使用 folder
  //   ext: '.svg' // 可选 .png 或 .png ，官方提供了两套格式的 emoji
  // })

  // var reg = /<[\s\S]*?src="([^\s]+)".*/
  // var result = emoji.match(reg)
  // console.log(result[1])
  </script>
  <!-- <script src="https://unpkg.com/vue@2.1.3/dist/vue.min.js"></script>
<script src="https://cdn.jsdelivr.net/emojione/2.2.6/lib/js/emojione.min.js"></script>
<script src="https://code.jquery.com/jquery-2.1.1.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/crypto-js/3.1.2/rollups/md5.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.97.8/js/materialize.min.js"></script> -->
  <script src="/static/chat/vue.min.js"></script>
  <script src="/static/chat/emojione.min.js"></script>
  <script src="/static/chat/jquery-2.1.1.min.js"></script>
  <script src="/static/chat/md5.js"></script>
  <script src="/static/chat/materialize.min.js"></script>
  <script src="/static/chat/app.js"></script>
</body>

</html>
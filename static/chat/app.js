// const CryptoJS = require('crypto-js');  //引用AES源码js
// import CryptoJS from 'crypto-js';
new Vue({
  el: '#app',

  data: {
    ws: null, // Our websocket
    newMsg: '', // Holds new messages to be sent to the server
    chatMsg: '',
    chatContent: '', // A running list of chat messages displayed on the screen
    email: null, // Email address used for grabbing an avatar
    username: null, // Our username
    joined: false // True if email and username have been filled in
  },

  created: function() {
    var self = this;
    this.ws = new WebSocket('ws://' + window.location.host + '/v1/chat/wschat');
    this.ws.addEventListener('message', function(e) {
      var msg = JSON.parse(e.data);
      self.chatContent += '<div class="chip">' +
        // '<img src="' + self.gravatarURL(msg.email) + '">' // Avatar
        '<img src="' + self.gravatarURL(msg.username) + '">'
        +
        msg.email +
        '</div>' +
        emojione.toImage(msg.message) + '<br/>'; // Parse emojis

      var element = document.getElementById('chat-messages');
      element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
    });
  },

  methods: {
    send: function() {
      // if (this.newMsg != '') {
        if (document.getElementById('chat').value != '') {
        // console.log($('<p>').html(this.newMsg).text())
        this.ws.send(
          JSON.stringify({
            email: this.email,
            username: this.username,
            // message: $('<p>').html(this.newMsg).text() // Strip out html
            message: document.getElementById('chat').value
          }));
        this.newMsg = ''; // Reset newMsg
      }
    },

    join: function() {
      if (!this.email) {
        Materialize.toast('You must enter an email', 2000);
        return
      }
      if (!this.username) {
        Materialize.toast('You must choose a username', 2000);
        return
      }
      this.email = $('<p>').html(this.email).text();
      this.username = $('<p>').html(this.username).text();
      this.joined = true;
    },

    // gravatarURL: function(email) {
    gravatarURL: function(username) {
      // return 'http://www.gravatar.com/avatar/' + CryptoJS.MD5(email);
      // return 'http://www.gravatar.com/avatar/' + calcMD5(email);这个可以
      return '/v1/chat/avatar/' + username
      // return '/go.jpg';
    }
  }
});
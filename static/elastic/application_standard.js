/* globals Vue */

(function() {
  'use strict'

  var app = new Vue({
    el: '#app',

    data: {
      state: {
        error: null,
        loading: true,
        fetching: false,
        replaceResults: true
      },

      isBottom: false,

      query: '',
      results: [],
      total: 0,
      search_after: []
    },

    methods: {
      loadResults: function() {
        var self = this

        if (self.state.fetching) { return }
        self.state.fetching = true

        if (self.query === '') {
          window.history.pushState({}, '', '?')
        } else {
          window.history.pushState({}, '', '?q=' + encodeURIComponent(self.query))
        }

        if (self.state.replaceResults) { self.search_after = [] }

        // window.fetch(`/v1/elastic/search?q=${encodeURIComponent(self.query)}&a=${self.search_after.join(',')}`)
        window.fetch(`/v1/wx/elasticsearch?q=${encodeURIComponent(self.query)}&a=${self.search_after.join(',')}`)
          .then(function(response) {
            if (!response.ok) { return Promise.reject(response) }
            return response.json()
          })
          .then(function(response) {
            var results = []

            response.hits.forEach(function(r) {
              var author = {
                first_name: r.author.first_name
              }
              var result = {
                id: r.id,
                url: r.url,//+'#page=3&keyword='+${encodeURIComponent(self.query)}
                image_url: r.image_url,
                published: r.published,
                body: r.body,
                author: author
              }

              if (r.highlights && r.highlights.title) {
                result.title = r.highlights.title[0]
              } else {
                result.title = r.title
              }

              if (r.highlights && r.highlights.alt) {
                result.alt = r.highlights.alt[0]
              }

              if (r.highlights && r.highlights.body) {
                // 获得<em>位置——截取前100和后100个字符，不满100按实际算
                var myString=""; //= r.highlights.body[0];
                r.highlights.body.forEach((elem, index) => {
                  console.log(elem, index);
                  myString=myString+'……'+elem;
                });

                // var myString = r.highlights.body[0];
                // var w = myString.indexOf("<em>");
                // console.log(w)
                // var start = 0
                // if (Number(w) >= 100) {
                //   start = w - 100
                // } else {
                //   start = w
                // }
                // var myString2 = myString.substr(start, 318)
                // if (myString2.length>317){
                //   myString=myString.substr(start, 317)+"......"
                // }
                result.body = myString//r.highlights.body[0]
              }

              if (r.highlights && r.highlights.transcript) {
                result.transcript = r.highlights.transcript.join('&hellip;')
              }

              results.push(result)
            })

            self.total = response.total

            if (self.state.replaceResults) {
              self.results = results
            } else {
              self.results = self.results.concat(results)
            }

            if (response.hits.length > 0) {
              self.search_after = response.hits[response.hits.length - 1].sort
            }
          })
          .then(function() {
            self.state.loading = false
            self.state.fetching = false
          })
          .catch(function(response) {
            self.state.loading = false
            self.state.fetching = false
            self.state.error = response
          })
      },

      toggle: function(event) {
        event.currentTarget.closest('div.result').classList.toggle('expanded')
      }
    },

    watch: {
      query: function() {
        this.state.replaceResults = true
        this.loadResults()
      },

      isBottom: function() {
        if (this.total > this.results.length) {
          this.state.replaceResults = false
          this.loadResults()
        }
      }
    },

    created: function() {
      var self = this

      var q = document.location.search.split('q=')[1]
      if (q) { self.query = decodeURIComponent(q) }
      self.loadResults()

      window.onscroll = function() {
        self.isBottom = (document.documentElement.scrollTop || document.body.scrollTop) + window.innerHeight === document.documentElement.scrollHeight
      }
    }
  })

  return app
})()

// vue.js提示Vue is not a constructor或Vue.createApp is not a function解决方法 -->
/*Vue 3*/
// Vue.createApp({
//   data() {
//     return {
//       items: [{ message: 'Foo' }, { message: 'Bar' }]
//     }
//   }
// }).mount('#array-rendering')

/*Vue 2*/
// var example1 = new Vue({
//   el: '#example-1',
//   data: {
//     items: [
//       { message: 'Foo' },
//       { message: 'Bar' }
//     ]
//   }
// })
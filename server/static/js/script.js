$(document).ready(function() {
  /** Serialize form data into JSON 

-- Modified from Original Code by Gabriel R.
https://jsfiddle.net/gabrieleromanato/bynaK/
   **/

  $.fn.serializeFormJSON = function () {
    var o = {};
    var a = this.serializeArray();
    $.each(a, function () {
      if (o[this.name]) {
        if (!o[this.name].push) {
          o[this.name] = [o[this.name]];
        }
        o[this.name].push(this.value || '');
      } else {
        o[this.name] = this.value || '';
      }
    });
    return o;
  };
});

$(document).ready(function() {
  // Splitter functions
  $('button#split').click(function() {
    $('form#form').submit(function (e) {
      e.preventDefault();
      var data = $(this).serializeFormJSON();

      // Perform a POST to split the tweet
      $.post(
        "split",
        JSON.stringify(data),
        function(res) {
          $('.row').remove()
          $.each(res, function(key, value) {
            $.each(this, function(k, v) {
              $('.footer').before( "<div class='row marketing'> <div class='col-log-6'> <hr> <p>" + v + "</p></div></div>" )
            });
          });
        },
        'json'
      )

    });

    $('button#tweet').remove()
    $('.jumbotron').append(" <button class='btn btn-lg btn-info' id='tweet' type='button'>Tweet!</button> ")
    // Tweet functions
    $('button#tweet').click(function(e) {
      e.preventDefault()

      $.post(
        "tweet",
        ""
      )
    });
  });
});

/** Serialize form data into JSON 

-- Code by Gabriel R.
https://jsfiddle.net/gabrieleromanato/bynaK/
 **/

(function ($) {
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
})(jQuery);

$(document).ready(function() {
  $('input#split').click(function() {
    $('form#form').submit(function (e) {
      e.preventDefault();
      var data = $(this).serializeFormJSON();

      $.post(
        "post",
        JSON.stringify(data),
        function(res) {
          console.log(res)
        },
        'json'
      )
    })
  })
});

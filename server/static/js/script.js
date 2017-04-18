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

$('form').submit(function (e) {
  e.preventDefault();
  var data = $(this).serializeFormJSON();
  console.log(data);

  //$.post("post", data);

  /* Object
        email: "value"
        name: "value"
        password: "value"
        */
});

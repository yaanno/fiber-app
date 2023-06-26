$(document).ready(function () {
  // If js is enabled prevent submit, otherwise the form should handle the submit
  $("#form").on("submit", function (e) {
    e.preventDefault();
  });
  $("#search").on("keyup", function () {
    $.get(`search?name=${encodeURIComponent($(this).val())}`, (newTable) => {
      // Do not follow this example in production code, if the source of the data is not trusted.
      $("#table").html(newTable);
    });
  });
});

$(function () {
  $("#submit-btn").click(function () {
    // Collect data from form inputs
    var packSizes = [];
    $('input[name="packsize"]').each(function () {
      packSizes.push($(this).val());
    });

    var packSizes = packSizes.map(Number);
    console.log(packSizes);

    $.ajax({
      type: "POST",
      url: "/submitPackSettings",
      contentType: "application/json",
      data: JSON.stringify(packSizes),
      success: function (response) {
        console.log(response);
        // Handle success response
      },
      error: function (error) {
        console.error(error);
        // Handle error
      }
    });
  });


  $("#calculate-btn").click(function () {
    // Collect data from form input
    var newPackSize = $("input[name='newpacksize']").val();

    // Create data object
    var data = {
      totalNumberOfPacks: Number(newPackSize)
    };

    // Perform AJAX request
    $.ajax({
      type: "POST",
      url: "/calculatePacks",
      data: JSON.stringify(data),
      success: function (response) {
        // Update the table with the calculated packs
       // $("tbody").html(""); // Clear existing rows
        response.forEach(function (pack) {
          $("#pack-calculate-form > table > tbody").append("<tr><td>" + pack.num + "</td>" + "<td>" + pack.size + "</td></tr>");
        });
      },
      error: function (error) {
        console.error(error);
        // Handle error
      }
    });
  });

});
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
        $("#packsizes").empty();
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
        $("#calc-result").empty(); // Clear existing rows
        console.log('response : ', response);
        response.forEach(function (pack) {
          if (pack.num > 0) {
            $("#calc-result").append("<tr><td>" + pack.size + "</td>" + "<td>" + pack.num + "</td></tr>");
          }
        });
      },
      error: function (error) {
        console.error(error);
        // Handle error
      }
    });
  });

});
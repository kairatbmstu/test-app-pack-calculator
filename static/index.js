$(function () {

  $("#packsizes").on("click", ".delete-row-btn", function () {
    $(this).closest("tr").remove();
  });

  $("#add-row-btn").click(function () {
    // Get the last pack size value
    var lastPackSize = 0;

    // Create a new row with an input field, copying the last pack size value
    var newRow = $("<tr>").html('<td><input type="number" name="packsize" value="' + lastPackSize + '" /><td><button class=\"delete-row-btn\" type=\"button\">Delete</button></td></td>');

    // Append the new row to the tbody
    $("#packsizes").append(newRow);
  });

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

        response.forEach(function (packSize) {
          if (packSize > 0) {
            $("#packsizes")
            .append("<tr><td><input type=\"number\" name=\"packsize\" value=" + packSize + " />" + 
            "</td><td><button class=\"delete-row-btn\" type=\"button\">Delete</button></td></tr>");
          }
        });
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
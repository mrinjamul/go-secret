/*
 main.js
*/

function check() {
  String.prototype.trim = function () {
    return this.replace(/^\s+|\s+$/g, "");
  };
  var usernameValue = document.getElementById("username").value.trim();
  var textAreaValue = document.getElementById("message").value;
  var trimmedTextAreaValue = textAreaValue.trim();
  if (trimmedTextAreaValue != "") {
    document.forms["form"].submit();
  } else {
    // document.body.style.background = "#FF4365";
    // document.body.style.color = "#FFFFF3";
    var body = document.getElementById("container");
    // remove class named from-cyan-600 to-cyan-400
    body.classList.remove("from-cyan-600");
    body.classList.remove("to-cyan-400");
    // add class named bg-gradient-to-t from-red-600 to-red-400
    body.classList.add("from-red-600");
    body.classList.add("to-red-400");
    return false;
  }
}

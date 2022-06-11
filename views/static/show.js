var count = 15;
document.getElementById("counter").innerHTML = count;

var interval = setInterval(function () {
  if (count === 0) {
    clearInterval(interval);
    document.getElementById("cover").classList.remove("hidden");
    document.getElementById("content").innerHTML = "";
  } else {
    count--;
  }
  document.getElementById("counter").innerHTML = count;
}, 1000);

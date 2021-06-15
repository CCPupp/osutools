var acc = document.getElementsByClassName("accordion");
var i;

for (i = 0; i < acc.length; i++) {
  acc[i].addEventListener("click", function() {
    this.classList.toggle("active");
    var panel = this.nextElementSibling;
    if (panel.style.maxHeight) {
      panel.style.maxHeight = null;
    } else {
      panel.style.maxHeight = panel.scrollHeight + "px";
    }
  });
}

function req(username) {
  const xhr = new XMLHttpRequest();
  const url = `https://api.github.com/users/${username}/repos`;

  xhr.open('GET', url, true);

  xhr.onload = function() {
      const data = JSON.parse(this.response);
      console.log(data);

      for (i = 0; i < data.length; i++) {
        if (data[i].name == "osutools") {
          document.getElementById("lastModified").innerHTML = "Last modified: " + data[i].updated_at.slice(0, 10) + " at " + data[i].updated_at.slice(11, 19) + " utc";
        }
      }
  }
  xhr.send();
}
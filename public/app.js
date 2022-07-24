window.addEventListener("DOMContentLoaded", (_) => {
  let websocket = new WebSocket("ws://" + window.location.host + "/websocket");
  let room = document.getElementById("chat-text");

  websocket.addEventListener("message", function (e) {
    let data = JSON.parse(e.data);
    // creating html element
    let p = document.createElement("p");
    console.log(data);
    let date = new Date(data.time);
    p.innerHTML = `<strong>${date.toLocaleString()} - ${data.userName}</strong>: ${data.text}`;

    room.append(p);
    room.scrollTop = room.scrollHeight; // Auto scroll to the bottom
  });

  let form = document.getElementById("input-form");
  form.addEventListener("submit", function (event) {
    event.preventDefault();
    let username = document.getElementById("input-username");
    let text = document.getElementById("input-text");
    websocket.send(
      JSON.stringify({
        user_name: username.value,
        message: text.value,
      })
    );
    text.value = "";
  });
});

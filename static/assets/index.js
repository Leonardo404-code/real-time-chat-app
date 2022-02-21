let socket = new WebSocket("ws://127.0.0.1:8080/ws");
const chat = document.querySelector("#chat");
const input = document.querySelector("#messages");
const send = document.querySelector("#send");

console.log("Attempting Connection...");

const insertMessage = (msg, chat) => {
  const message = document.createElement("p");

  message.setAttribute("class", "chat-message");

  message.textContent = `${msg}`;

  chat.appendChild(message);
};

socket.onopen = () => {
  socket.send("Novo usuário conectado");
};

socket.onclose = (event) => {
  console.log("Socket Closed Connection: ", event);
  socket.send("Client Closed!");
};

socket.onerror = (error) => {
  console.log("Socket Error: ", error);
};

socket.onmessage = (msg) => {
  const { data } = msg;

  insertMessage(data, chat);
};

send.onclick = () => {
  const message = input.value;

  socket.send(message);
  input.value = "";
};

input.addEventListener("keyup", (e) => {
  if (e.keyCode === 13) {
    send.click();
  }
});

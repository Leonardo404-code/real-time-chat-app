const chat = document.querySelector("#chat");
const input = document.querySelector("#messages");
const send = document.querySelector("#send");
const userModal = document.querySelector("#user");
const usernameInput = document.querySelector("#username");
const addUserButton = document.querySelector("#addUser");
const addUserClient = document.querySelector("#newClient");
const username = sessionStorage.getItem("user");

const insertMessage = (msg, chat) => {
  const message = document.createElement("p");

  message.setAttribute("class", "chat-message");

  message.textContent = `${msg}`;

  chat.appendChild(message);
};

const Connect = (username) => {
  console.log("Attempting Connection...");

  let socket = new WebSocket("ws://127.0.0.1:8080/ws");

  socket.onopen = () => {
    socket.send(`${username} Entrou`);
  };

  socket.addEventListener("close", () => {
    socket.send(`${username} Saiu`, chat);
    insertMessage(`${username} Saiu`, chat);
  });

  socket.onerror = (error) => {
    console.log("Socket Error: ", error);
  };

  socket.onmessage = (msg) => {
    const { data } = msg;

    insertMessage(data, chat);
  };

  send.onclick = () => {
    const message = input.value;

    if (message === "") {
      return;
    }

    socket.send(`${username}: ${message}`);
    input.value = "";
  };

  input.addEventListener("keyup", (e) => {
    if (e.keyCode === 13) {
      send.click();
    }
  });
};

if (username === null) {
  addUserButton.addEventListener("click", (e) => {
    e.preventDefault();

    const value = usernameInput.value;

    if (e.keyCode === 13) {
      sessionStorage.setItem("user", value);
      userModal.style.display = "none";
      Connect(value);
    }

    sessionStorage.setItem("user", value);
    userModal.style.display = "none";
    Connect(value);
  });

  addUserClient.addEventListener("click", () => {
    sessionStorage.setItem("user", "Usuário");
    userModal.style.display = "none";
    Connect("Usuário");
  });
} else {
  userModal.style.display = "none";
  Connect(username);
}

let socket = new WebSocket("ws://127.0.0.1:6000/ws");
console.log("Attempting connection to server..");

socket.onopen = (event) => {
  console.log('The websocket connection is connect.');
  socket.send("Hello, Server.");
}

socket.onmessage = (event) => {
  alert(`[message] message sent from server: ${event.data}`);
}

socket.onclose = (event) => {
 if (event.wasClean) {
   console.log(`[closed] connection is closed successfully. code=${event.code} reason=${event.reason}`);
 } else {
   console.log(`[closed] connection is dead.`);
 }
}

socket.onerror = (err) => {
  console.log(`[error] ${err.message}`)
}
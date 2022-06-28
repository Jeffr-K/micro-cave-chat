console.log("실행이 안되나?")

let socket = new WebSocket("ws://localhost:6700/ws");
console.log("소켓 연결을 시도합니다.");
socket.onopen = (event) => {
  console.log('소켓 연결을 시도하고 메세지를 보냅니다.');
  let message = "Hello, Server?"
  socket.send(`[message] 클라이언트로부터 온 메세지: ${message}.`);
}
socket.onmessage = (event) => {
  console.log(`[message] 서버로부터 온 메세지: ${event.data}`);
}
socket.onclose = (event) => {
  if (event.wasClean) {
    console.log(`[closed] 커넥션이 성공적으로 종료되었습니다. code=${event.code} reason=${event.reason}`);
  } else {
    console.log(`[closed] 커넥션이 종료 됩니다.`);
  }
}
socket.onerror = (err) => {
  console.log(`[error] ${err.message}`)
}
#### Websocket

##### 클라이언트 Request 의 이해

소켓 통신은 서버와 클라이언트 측 미리 커넥션의 핸드 쉐이크 과정을 통해 계약하고, 파이프라인을 구축해 데이터를 양방향 지속적으로 교류하는 기술이다.
클라이언트에서 어떤 계약서를 보내면 서버가 이를 받아들이고 계약을 하는지 공부하자.

아래는 클라이언트가 서버로 커넥션 계약을 맺기 위해 보내는 요청이다:

```text
GET ws://domain.com:8080/ws HTTP/1.1
Host: 127.0.0.1/8080
Connection: Upgrade
Pragma: no-cache
Cache-Control: no-cache
Upgrade: websocket
Sec-WebSocket-Version: 13
Sec-WebSocket-Key: q4xkc032u266gidTuKaSow==
```

서버는 이 요청에 대해 아래와 같이 응답한다.

```text
HTTP/1.1 101 Switching Protocols
Upgrade: websocket
Connection: Upgrade
Sec-WebSocket-Accept: fA9dggdnMPU89Ijqw1tkTUIL=
```

- 클라이언트가 받고자 하는 Response Header 는 HTTP/1.1 101 Switching Protocols 이다.
- Upgrade 와 Connection 정보는 Response 로 변경 없이 그대로 보낸다.
- Sec-Websocket-Accept: 클라이언트가 보낸 Sec-WebSocket-Key 문자열에 상수값 `258EAFA5-E914-47DA-95CA-C5AB0DC85B11` 를 붙여서
  RPC 6455 SHA-1 Hasing + Base64 encoding 을 해준 값을 보낸다.

##### 커넥션이 이루어진 다음엔?
Websocket Connection 이 적절히 이루어지면, 클라이언트와 서버는 UTF-8 text 또는 byte 메세지를 주고 받는다. WebSocket 은 Framed protocol 이며,
메세지를 프레임에 담아 보내고, 메세지가 크다면 여러 프레임으로 나누어 보내게 된다.
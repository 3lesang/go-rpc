import useWebSocket from "react-use-websocket";

const socketURL = "ws://localhost:3000/rpc/v1";

function App() {
  const { sendMessage } = useWebSocket(socketURL, {
    share: true,
    onOpen: () => console.log("WebSocket connection opened"),
    onMessage: (event) => {
      console.log("Received message:", event.data);
    },
  });

  const handleSend = () => {
    sendMessage(
      JSON.stringify({ method: "sayHello", params: "Hello, server!" })
    );
  };
  return <button onClick={handleSend}>send</button>;
}

export default App;

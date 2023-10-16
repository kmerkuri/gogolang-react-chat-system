import React, {useEffect, useState} from "react";
import "./App.css";
import { connect, sendMsg } from "./api";
import Header from './components/Header/Header';
import ChatHistory from "./components/ChatHistory";
import Chat from "./components/Chat/Chat";

function App() {
  const [chatHistory, setChatHistory] = useState([]);
  const [message, setMessage] = useState("");

  useEffect(() => {
    connect((msg) => {
      setChatHistory([...chatHistory, msg])
    })
  }, [connect, chatHistory]);

  console.log("hi", message);

  const handleSend = () => {
      sendMsg(message);

      setMessage('');
  }

  return (
      <div className="App">
        <Header />
        <Chat setMessage={setMessage} handleSend={handleSend} message={message} />
        <ChatHistory chatHistory={chatHistory} />
        {/*<button onClick={handleSend}>Hit</button>*/}
      </div>
  )
}

export default App;
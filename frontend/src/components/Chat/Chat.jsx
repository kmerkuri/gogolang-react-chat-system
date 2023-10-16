import React from "react";

function Chat({setMessage, handleSend, message}){
    return (
        <div style={{display: "flex", flexDirection: "column"}}>
        <label>
            Write your message: </label>
            <textarea value={message} onChange={(event) => setMessage(event.target.value)} name="postContent" rows={4} cols={40} />
            <button onClick={handleSend}>Hit</button>
        </div>
    );
}

export default Chat
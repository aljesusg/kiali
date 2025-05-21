
import ReactDOM from 'react-dom';
import React, { useRef, useState } from 'react';
import { Spinner } from '@patternfly/react-core';
import { ChatbotWindow } from './ChatbotWindow';

export function Chatbot({...props}) {
    const chatbotFrameRef = useRef<HTMLIFrameElement>(null);
    const [loading, setLoading] = useState(true);

    const injectStyles = () => {
        console.log("Inject styles")
    }

    const sendChatbotState = () => {
        sendMessage(props?.username);
    };

    const sendMessage = (message: string) => {
        const www = chatbotFrameRef.current?.contentWindow;
        www?.postMessage(JSON.stringify(message));
    };

    const chatbotLoadCompleted = () => {
        injectStyles();
        setLoading(false);
        sendChatbotState();
    };
    
    return (
        <>
        {loading ? <Spinner /> : <></>}
            <iframe
                title="Ansible Chatbot IFrame"
                srcDoc="<!DOCTYPE html>"
                ref={chatbotFrameRef}
                onLoad={chatbotLoadCompleted}
                {...props}
                style={{
                height: '100%',
                width: '100%',
                position: 'sticky',
                padding: '0px',
                margin: '0px',
                border: '0px',
                overflow: 'hidden',
                }}
            >
                {chatbotFrameRef.current !== null &&
                ReactDOM.createPortal(
                    <ChatbotWindow />,
                    // @ts-expect-error IFrame for chatbot must exist
                    chatbotFrameRef.current?.contentWindow.document.body
                )}
            </iframe>
        </>
    )
}
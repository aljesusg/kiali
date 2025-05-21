import { connect } from 'react-redux';
import { App as KialiChat } from "kiali_ai_chatbot";
import { KialiAppState } from 'store/Store';


const THEME_CLASS_LIGHT = 'pf-v6-theme-light';
const THEME_CLASS_DARK = 'pf-v6-theme-dark';

type ChatbotWProps = {
    username: string;
    theme: string;
}


export const ChatbotW: React.FC<ChatbotWProps> = (props: ChatbotWProps) => {
    
    return <KialiChat username={props.username}/>
}

const mapStateToProps = (state: KialiAppState): ChatbotWProps => {
  return {
    username: state.authentication.session?.username || 'anonymous',
    theme: state.globalState.theme
  };
};

export const ChatbotWindow = connect(mapStateToProps)(ChatbotW);
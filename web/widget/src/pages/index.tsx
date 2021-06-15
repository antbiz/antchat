import React, { useEffect, useState } from 'react';
import Chat, { Bubble, useMessages } from '@chatui/core';
import useWebSocket from "react-use-websocket";
import ProForm, { ProFormCaptcha } from '@ant-design/pro-form';
import { getChatHistory, sendMsg } from '@/services/api';
import { getApiSid } from '@/utils/authority';
import '@chatui/core/dist/index.css';
import './index.less';


const initialMessages = [
  {
    type: 'text',
    content: { text: '您好！请问有什么可以帮助您？' },
    user: { avatar: '//gw.alicdn.com/tfs/TB1DYHLwMHqK1RjSZFEXXcGMXXa-56-62.svg' },
  },
];

// 默认快捷短语，可选
const defaultQuickReplies = [
  {
    icon: 'message',
    name: '老哥！666啊',
    isNew: true,
    isHighlight: true,
  },
];

export default function() {
  // 消息列表
  const { messages, appendMsg, setTyping, resetList } = useMessages(initialMessages);
  const [opened, changeOpened] = useState<boolean>(false);
  const [signined, changeSignined] = useState<boolean>(false);

  useWebSocket(
    `ws://${document.location.host}/api/v1/visitor/chat?x-api-sid=${getApiSid()}`,
    {
      onOpen: () => console.log("Connection Opened"),
      onClose: () => console.log("Websocket Connection Closed"),
      onError: (event: any) => console.log(event),
      onMessage: (event: any) => {
        const response = JSON.parse(event.data);
        console.log("recevied: ", response);
        appendMsg({
          ...response,
          position: 'left',
        });
      },
      share: true,
      //Will attempt to reconnect on all close events, such as server shutting down
      shouldReconnect: (_closeEvent) => signined,
      reconnectAttempts: 10,
      reconnectInterval: 3000,
    },
    true
  );

  const handleOpenChatBtnClick = () => {
    changeOpened(!opened);
  };

  // 快捷短语回调，可根据 item 数据做出不同的操作，这里以发送文本消息为例
  const handleQuickReplyClick = (item: any) => {
    handleSend('text', item.name);
  };

  function renderMessageContent(msg: any) {
    const { type, content } = msg;

    // 根据消息类型来渲染
    switch (type) {
      case 'text':
        return <Bubble content={content.text} />;
      case 'image':
        return (
          <Bubble type="image">
            <img src={content.picUrl} alt="" />
          </Bubble>
        );
      default:
        return null;
    }
  };

  const handleSend = async (type: string, val: string) => {
    if (type !== 'text' || !val.trim()) return;
    appendMsg({
      type: 'text',
      content: { text: val },
      position: 'right',
    });
    try {
      await sendMsg({
        type: 'text',
        content: { text: val },
      });
    } catch (e) {
      // do nothing
    }
    // setTyping(true);
  };

  const fetchChatHistory = async () => {
    await getChatHistory().then(resp => {
      changeSignined(true);
      const { data = [] } = resp;
      resetList(data.map(function(item) {
        return {
          _id: item.id,
          type: item.type,
          content: item.content,
          position: item.senderID === item.visitorID ? 'right' : 'left'
        }
      }));
    }).catch(e => {
      // do nothing
    });
  };

  useEffect(() => {
    fetchChatHistory();
  }, []);

  return (
    <div className={`widget ${opened ? "opened" : ''}`}>
      <div className="chat-dialog">
        <div className="close-btn">
          <a onClick={() => {
            changeOpened(false);
          }}>
            <img src="//gw.alicdn.com/tfs/TB1lWlNOkvoK1RjSZPfXXXPKFXa-29-29.svg" />
          </a>
        </div>
        {
          signined ? (
            <Chat
              navbar={{ title: '智能助理' }}
              messages={messages}
              renderMessageContent={renderMessageContent}
              quickReplies={defaultQuickReplies}
              onQuickReplyClick={handleQuickReplyClick}
              onSend={handleSend}
            />
          ) : (
            <div className="chatApp" style={{ margin: '80px 30px' }}>
              <ProForm
                submitter={{
                  searchConfig: {
                    submitText: '开始对话',
                  },
                  render: (_, dom) => dom.pop(),
                  submitButtonProps: {
                    size: 'large',
                    style: {
                      width: '100%',
                    },
                  },
                }}
              >
                <ProFormCaptcha
                  fieldProps={{
                    size: 'large',
                  }}
                  captchaProps={{
                    size: 'large',
                  }}
                  name="captcha"
                  rules={[
                    {
                      required: true,
                      message: '请输入验证码',
                    },
                  ]}
                  placeholder="请输入验证码"
                  captchaTextRender={() => (
                    <img style={{ height: '100%' }} src="data:image/png;base64,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" />
                  )}
                  onGetCaptcha={async () => {

                  }}
                />
              </ProForm>
            </div>
          )
        }
      </div>
      <div className="open-chat-btn" onClick={handleOpenChatBtnClick}>
        <img
          src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADAAAAAwCAMAAABg3Am1AAAAUVBMVEUAAAD///////////////////////////////////////////////////////////////////////////////////////////////////////8IN+deAAAAGnRSTlMAAwgJEBk0TVheY2R5eo+ut8jb5OXs8fX2+cjRDTIAAADsSURBVHgBldZbkoMgFIThRgQv8SKKgGf/C51UnJqaRI30/9zfe+NQUQ3TvG7bOk9DVeCmshmj/CuOTYnrdBfkUOg0zlOtl9OWVuEk4+QyZ3DIevmSt/ioTvK1VH/s5bY3YdM9SBZ/mUUyWgx+U06ycgp7D8msxSvtc4HXL9BLdj2elSEfhBJAI0QNgJEBI1BEBsQClVBVGDgwYOLAhJkDM1YOrNg4sLFAsLJgZsHEgoEFFQt0JAFGFjQsKAMJ0LFAexKgZYFyJIDxJIBNJEDNAtSJBLCeBDCOBFAPzwFA94ED+zmhwDO9358r8ANtIsMXi7qVAwAAAABJRU5ErkJggg=="
          className="open-chat-btnimg"
        />
        <span
          className="open-chat-btntxt"
        >Chat with us</span>
      </div>
    </div>
  );
};

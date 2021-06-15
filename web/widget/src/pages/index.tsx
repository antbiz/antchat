import React, { useEffect, useState } from 'react';
import Chat, { Bubble, useMessages } from '@chatui/core';
import useWebSocket from "react-use-websocket";
import ProForm, { ProFormCaptcha } from '@ant-design/pro-form';
import { getCaptcha, getChatHistory, sendMsg, signin } from '@/services/api';
import { getApiSid, getVid, setApiSid, setVid } from '@/utils/authority';
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
  const [captcha, setCaptcha] = useState<API.CaptchaGenReply>();

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
      // 未登录
      getCaptcha().then(resp => {
        setCaptcha(resp);
      })
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
              navbar={{ title: 'AntChat客服' }}
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
                onFinish={async (values) => {
                  await signin({
                    captchaID: captcha?.id,
                    captcha: values.captcha,
                    domain: window.location.host,
                    visitorID: getVid(),
                  }).then(resp => {
                    const { id, sid } = resp;
                    setApiSid(sid);
                    setVid(id);
                    changeSignined(true);
                  })
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
                    <img style={{ height: '100%' }} src={captcha?.base64} />
                  )}
                  onGetCaptcha={async () => {
                    const resp = await getCaptcha();
                    setCaptcha(resp);
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

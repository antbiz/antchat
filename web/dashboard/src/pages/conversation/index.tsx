import React, { useState } from 'react';
import { List, Result, Comment, Tooltip } from 'antd';
import { MessageOutlined } from '@ant-design/icons';
import ProCard from '@ant-design/pro-card';
import Chat, { Bubble, useMessages } from '@chatui/core';
import useWebSocket from "react-use-websocket";
import moment from 'moment';
import { getChatHistory, sendMsg } from '@/services/agent';
import { getApiSid } from '@/utils/authority';
import '@chatui/core/dist/index.css';
import './index.less';


// 默认快捷短语，可选
// TODO: 可配置
const defaultQuickReplies = [
  {
    icon: 'message',
    name: '您好！请问有什么可以帮助您？',
    isNew: true,
    isHighlight: true,
  },
];

export default (): React.ReactNode => {
  const [conversations, setSonversations] = useState<API.Conversation[]>();
  const [activeConversation, setActiveConversation] = useState<API.Conversation>();
  // 消息列表
  const { messages, appendMsg, setTyping, resetList } = useMessages();
  const [chatBoxLoading, setChatBoxLoading] = useState<Boolean>(false);

  // TODO: 客户端发送心跳？目前服务端有个45s的ping
  useWebSocket(
    `ws://${document.location.host}/api/v1/agent/chat?x-api-sid=${getApiSid()}`,
    {
      onOpen: () => console.log("Connection Opened"),
      onClose: () => console.log("Websocket Connection Closed"),
      onError: (event: any) => console.log(event),
      onMessage: (event: any) => {
        const response = JSON.parse(event.data);
        console.log("recevied: ", response);
        if (response.type === "cmd" && response.content && response.content.code === "incoming_update") {
          setSonversations(response.content.list);
        } else if (response.visitorID === activeConversation?.id) {
          appendMsg({
            ...response,
            position: 'left',
          })
        }
      },
      share: true,
      //Will attempt to reconnect on all close events, such as server shutting down
      shouldReconnect: (_closeEvent) => true,
      reconnectAttempts: 10,
      reconnectInterval: 3000
    },
    true
  );

  const handleSend = async (type: string, val: string) => {
    if (!activeConversation || type !== 'text' || !val.trim()) return;
    appendMsg({
      type: 'text',
      content: { text: val },
      position: 'right',
    });
    try {
      await sendMsg({
        receiverID: activeConversation.id,
        receiverNick: activeConversation.nickname,
        type: 'text',
        content: { text: val },
      })
    } catch (e) {
      // do nothing
    }
    setTyping(true);

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

  const handleSwitchConversation = async (item: API.Conversation) => {
    setActiveConversation(item);
    setChatBoxLoading(true);
    try {
      const historyMsgList = await getChatHistory(item.id);
      resetList(historyMsgList);
    } catch (e) {
      // do nothing
    }
    setChatBoxLoading(false);
  }

  return (
    <ProCard split="vertical">
      <ProCard
        title="Conversations"
        colSpan="30%"
      >
        <List<API.Conversation>
          pagination={{
            onChange: page => {
              console.log(page);
            },
            pageSize: 10,
          }}
          itemLayout="horizontal"
          dataSource={conversations}
          renderItem={item => (
            <div className={`conversation ${activeConversation?.id === item.id ? 'active': '' }`}
              onClick={(e) => {
                handleSwitchConversation(item)
              }}
            >
              <Comment
                author={item.nickname}
                // avatar={item.avatar}
                content={item.content ? item.content.text || '' : ''}
                datetime={() => {
                  const activeAgo = item.activeAt ? moment(item.activeAt as string).fromNow() : '';
                  return activeAgo ? (
                      <Tooltip title={activeAgo}>
                        <span>2 hours ago</span>
                      </Tooltip>
                    ) : ''
                  }
                }
              />
            </div>
          )}
        />
      </ProCard>
      <ProCard
        loading={chatBoxLoading}
      >
        {
          activeConversation ? (
            <Chat
              navbar={{ title: `匿名用户 ${activeConversation.nickname}` }}
              messages={messages}
              renderMessageContent={renderMessageContent}
              quickReplies={defaultQuickReplies}
              onQuickReplyClick={handleQuickReplyClick}
              onSend={handleSend}
            />
          ) : (
            <Result
              icon={<MessageOutlined />}
              title="无选中对话"
              subTitle="没有需要处理的对话，休息一下吧！"
            />
          )
        }
      </ProCard>
    </ProCard>
  );
};

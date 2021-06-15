import React, { useState, useEffect } from 'react';
import { List, Result, Comment } from 'antd';
import { MessageOutlined } from '@ant-design/icons';
import ProCard from '@ant-design/pro-card';
import Chat, { Bubble, useMessages } from '@chatui/core';
import useWebSocket from "react-use-websocket";
import dayjs from 'dayjs';
import relativeTime from 'dayjs/plugin/relativeTime';
import VisitorForm from './VisitorForm';
import { getChatHistory, getConversations, sendMsg } from '@/services/agent';
import { getApiSid } from '@/utils/authority';
import '@chatui/core/dist/index.css';
import './index.less';
dayjs.extend(relativeTime);


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
  const [conversations, setSonversations] = useState<API.Conversation[]>([]);
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
          if (response.visitorID !== activeConversation?.id) {
            conversations.push(response.content.data);
            setSonversations(conversations);
          };
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
        visitorID: activeConversation.id,
        type: 'text',
        content: { text: val },
      })
    } catch (e) {
      // do nothing
    }
    // setTyping(true);
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
      const { data = [] } = await getChatHistory(item.id);
      resetList(data.map(function(item) {
        return {
          _id: item.id,
          type: item.type,
          content: item.content,
          position: item.senderID === item.visitorID ? 'left' : 'right'
        }
      }));
    } catch (e) {
      // do nothing
    }
    setChatBoxLoading(false);
  };

  const fetchConversations = async () => {
    const { data } = await getConversations();
    setSonversations(data || []);
    if (data) {
      handleSwitchConversation(data[0]);
    };
  };

  useEffect(() => {
    fetchConversations();
  }, []);

  return (
    <ProCard split="vertical">
      <ProCard
        title="Conversations"
        colSpan="20%"
        className="conversation-container"
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
                avatar="https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"
                content={item.content ? item.content.text || '' : ''}
                datetime={dayjs(item.activeAt as string).fromNow()}
              />
            </div>
          )}
          className="conversation-list"
        />
      </ProCard>
      <ProCard
        loading={chatBoxLoading}
      >
        {
          activeConversation ? (
            <Chat
              navbar={{ title: `${activeConversation.nickname || '匿名用户'}` }}
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
      {activeConversation && (
        <ProCard
          title="客户信息"
          colSpan="20%"
          loading={chatBoxLoading}
        >
          <VisitorForm visitorID={activeConversation.id} />
        </ProCard>
      )}
    </ProCard>
  );
};

import React, { useState } from 'react';
import { List, Result, Comment, Tooltip } from 'antd';
import { MessageOutlined } from '@ant-design/icons';
import ProCard from '@ant-design/pro-card';
import Chat, { Bubble, useMessages } from '@chatui/core';
import useWebSocket from "react-use-websocket";
import '@chatui/core/dist/index.css';
import './index.less';

export type Conversation = {
  id: string;
  active?: boolean;
  avatar?: string;
  contactName?: string;
  msg?: string;
}

// 默认快捷短语，可选
const defaultQuickReplies = [
  {
    icon: 'message',
    name: '联系人工服务',
    isNew: true,
    isHighlight: true,
  },
  {
    name: '短语1',
    isNew: true,
  },
  {
    name: '短语2',
    isHighlight: true,
  },
  {
    name: '短语3',
  },
];

export default (): React.ReactNode => {
  const [tab, setTab] = useState('mine');
  const [mineContacts, setMineContacts] = useState<Conversation[]>();
  const [activeConversation, setActiveConversation] = useState<String>('');
  // 消息列表
  const { messages, appendMsg, setTyping } = useMessages();

  const {
    sendJsonMessage
  } = useWebSocket(
    'ws://localhost:8199',
    {
      onOpen: () => console.log("Connection Opened"),
      onClose: () => console.log("Websocket Connection Closed"),
      onError: (event: any) => console.log(event),
      onMessage: (event: any) => {
        const response = JSON.parse(event.data);
        console.log("recevied: ", response);
        if (response.type === "cmd" && response.content && response.content.code === "mine_contacts") {
          setMineContacts(response.content.list);
          setActiveConversation(response.content.list[0].id);
        } else {
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

  // 发送回调
  function handleSend(type: any, val: any) {
    if (type === 'text' && val.trim()) {
      // TODO: 发送请求
      appendMsg({
        type: 'text',
        content: { text: val },
        position: 'right',
      });
      sendJsonMessage({
        type: 'text',
        content: { text: val },
        position: 'right',
      })

      setTyping(true);
    }
  }

  // 快捷短语回调，可根据 item 数据做出不同的操作，这里以发送文本消息为例
  function handleQuickReplyClick(item: any) {
    handleSend('text', item.name);
  }

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
  }

  return (
    <ProCard split="vertical">
      <ProCard
        title="Conversations"
        colSpan="30%"
        tabs={{
          tabPosition: 'top',
          activeKey: tab,
          onChange: (key) => {
            setTab(key);
          },
        }}
      >
        <ProCard.TabPane key="mine" tab="Mine" className="conversation-list">
          <List
            pagination={{
              onChange: page => {
                console.log(page);
              },
              pageSize: 10,
            }}
            itemLayout="horizontal"
            dataSource={mineContacts}
            renderItem={item => (
              <div className={`conversation ${activeConversation === item.id ? 'active': '' }`}
                onClick={(e) => {
                  setActiveConversation(item.id)
                }}
              >
                <Comment
                  author={item.contactName}
                  avatar={item.avatar}
                  content={item.msg}
                  datetime={
                    <Tooltip title="2 hours ago">
                      <span>2 hours ago</span>
                    </Tooltip>
                  }
                />
              </div>
            )}
          />
        </ProCard.TabPane>
        <ProCard.TabPane key="unassigned" tab="Unassigned">
          Unassigned
        </ProCard.TabPane>
        <ProCard.TabPane key="all" tab="All">
          All
        </ProCard.TabPane>
      </ProCard>
      <ProCard>
        {
          tab === "all" ? (
            <Result
              icon={<MessageOutlined />}
              title="无选中对话"
              subTitle="没有需要处理的对话，休息一下吧！"
            />
          ) : (
            <Chat
              navbar={{ title: `匿名用户 ${activeConversation}` }}
              messages={messages}
              renderMessageContent={renderMessageContent}
              quickReplies={defaultQuickReplies}
              onQuickReplyClick={handleQuickReplyClick}
              onSend={handleSend}
            />
          )
        }
      </ProCard>
    </ProCard>
  );
};

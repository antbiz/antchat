import React, { useState } from 'react';
import { List, Result, Comment, Tooltip } from 'antd';
import { MessageOutlined } from '@ant-design/icons';
import ProCard from '@ant-design/pro-card';
import Chat, { Bubble, useMessages } from '@chatui/core';
import '@chatui/core/dist/index.css';
import './index.less';

export type Conversation = {
  id?: string;
  active?: boolean;
  avatar?: string;
  contactName?: string;
  msg?: string;
}

const initialMessages = [
  {
    type: 'text',
    content: { text: '主人好，我是智能助理，你的贴心小助手~' },
    user: { avatar: '//gw.alicdn.com/tfs/TB1DYHLwMHqK1RjSZFEXXcGMXXa-56-62.svg' },
  },
  {
    type: 'image',
    content: {
      picUrl: 'https://gw.alipayobjects.com/zos/rmsportal/mqaQswcyDLcXyDKnZfES.png',
    },
  },
];

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

const dataSource = [
  {
    id: 'q',
    contactName: '语雀的天空',
    avatar:
      'https://gw.alipayobjects.com/zos/antfincdn/efFD%24IOql2/weixintupian_20170331104822.jpg',
    desc: '我是一条测试的描述',
    msg: '主人好，我是智能助理，你的贴心小助手~',
  },
  {
    id: 'w',
    contactName: 'Ant Design',
    avatar:
      'https://gw.alipayobjects.com/zos/antfincdn/efFD%24IOql2/weixintupian_20170331104822.jpg',
    desc: '我是一条测试的描述',
    msg: '主人好，我是智能助理，你的贴心小助手~',
  },
  {
    id: 'e',
    contactName: '蚂蚁金服体验科技',
    avatar:
      'https://gw.alipayobjects.com/zos/antfincdn/efFD%24IOql2/weixintupian_20170331104822.jpg',
    desc: '我是一条测试的描述',
    msg: '主人好，我是智能助理，你的贴心小助手~',
  },
  {
    id: 't',
    contactName: 'TechUI',
    avatar:
      'https://gw.alipayobjects.com/zos/antfincdn/efFD%24IOql2/weixintupian_20170331104822.jpg',
    desc: '我是一条测试的描述',
    msg: '主人好，我是智能助理，你的贴心小助手~',
  },
];

export default (): React.ReactNode => {
  const [tab, setTab] = useState('mine');
  const [activeConversation, setActiveConversation] = useState<String>(dataSource[0].id);
  // 消息列表
  const { messages, appendMsg, setTyping } = useMessages(initialMessages);

  // 发送回调
  function handleSend(type: any, val: any) {
    if (type === 'text' && val.trim()) {
      // TODO: 发送请求
      appendMsg({
        type: 'text',
        content: { text: val },
        position: 'right',
      });

      setTyping(true);

      // 模拟回复消息
      setTimeout(() => {
        appendMsg({
          type: 'text',
          content: { text: '亲，您遇到什么问题啦？请简要描述您的问题~' },
        });
      }, 1000);
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
              pageSize: 3,
            }}
            itemLayout="horizontal"
            dataSource={dataSource}
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
              title="No Chat!"
            />
          ) : (
            <Chat
              navbar={{ title: '匿名用户 127.0.0.1' }}
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

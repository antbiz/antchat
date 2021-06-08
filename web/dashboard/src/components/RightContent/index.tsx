import React, { useState } from 'react';
import { Tag, Space } from 'antd';
import { QuestionCircleOutlined } from '@ant-design/icons';
import useWebSocket from "react-use-websocket";
import { useModel, SelectLang } from 'umi';
import Avatar from './AvatarDropdown';
import styles from './index.less';

export type SiderTheme = 'light' | 'dark';

const ENVTagColor = {
  dev: 'orange',
  test: 'green',
  pre: '#87d068',
};

const GlobalHeaderRight: React.FC = () => {
  const { initialState } = useModel('@@initialState');
  const [connect] = useState(true);

  if (!initialState || !initialState.settings) {
    return null;
  }

  useWebSocket(
    'ws://localhost:8199',
    {
      onOpen: () => console.log("Connection Opened"),
      onClose: () => console.log("Websocket Connection Closed"),
      onError: (event: any) => console.log(event),
      share: true,
      //Will attempt to reconnect on all close events, such as server shutting down
      shouldReconnect: (_closeEvent) => true,
      reconnectAttempts: 10,
      reconnectInterval: 3000
    },
    connect
  );

  const { navTheme, layout } = initialState.settings;
  let className = styles.right;

  if ((navTheme === 'dark' && layout === 'top') || layout === 'mix') {
    className = `${styles.right}  ${styles.dark}`;
  }

  return (
    <Space className={className}>
      <span
        className={styles.action}
        onClick={() => {
          window.open('https://github.com/antbiz/antchat');
        }}
      >
        <QuestionCircleOutlined />
      </span>
      <Avatar menu={true} />
      {REACT_APP_ENV && (
        <span>
          <Tag color={ENVTagColor[REACT_APP_ENV]}>{REACT_APP_ENV}</Tag>
        </span>
      )}
      <SelectLang className={styles.action} />
    </Space>
  );
};

export default GlobalHeaderRight;

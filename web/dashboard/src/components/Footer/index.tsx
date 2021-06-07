import { useIntl } from 'umi';
import { GithubOutlined } from '@ant-design/icons';
import { DefaultFooter } from '@ant-design/pro-layout';

export default () => {
  const intl = useIntl();
  const defaultMessage = intl.formatMessage({
    id: 'app.copyright.produced',
    defaultMessage: 'AntBiz',
  });

  return (
    <DefaultFooter
      copyright={`2021 ${defaultMessage}`}
      links={[
        {
          key: 'antchat',
          title: 'AntApi',
          href: 'https://github.com/antbiz/antchat',
          blankTarget: true,
        },
        {
          key: 'github',
          title: <GithubOutlined />,
          href: 'https://github.com/BeanWei',
          blankTarget: true,
        },
        {
          key: 'antbiz',
          title: 'AntBiz',
          href: 'https://github.com/antbiz',
          blankTarget: true,
        },
      ]}
    />
  );
};

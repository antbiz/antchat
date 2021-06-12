import React from 'react';
import ProForm, { ProFormText } from '@ant-design/pro-form';
import { getVisitor, updateVisitor } from '@/services/visitor';

const VisitorForm: React.FC<{visitorID: string}> = (props) => {
  return (
    <ProForm<API.Visitor>
      onFinish={async (values) => {
        updateVisitor(props.visitorID, values);
      }}
      request={async () => {
        const data = await getVisitor(props.visitorID);
        return data;
      }}
      submitter={{
        searchConfig: {
          submitText: '保存',
        },
        render: (_, dom) => dom.pop(),
        submitButtonProps: {
          // size: 'large',
          style: {
            width: '100%',
          },
        },
      }}
    >
      <ProFormText
        name="nickname"
        label="昵称"
        placeholder="客户昵称"
        rules={[
          {
            required: true,
            message: '请输入客户昵称!',
          },
        ]}
      />
      <ProFormText
        name="phone"
        label="手机号"
        placeholder="客户手机号"
        rules={[
          {
            pattern: /^1\d{10}$/,
            message: '不合法的手机号格式!',
          },
        ]}
      />
      <ProFormText
        name="email"
        label="邮箱"
        placeholder="客户邮箱"
        rules={[
          {
            pattern: /^([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+\.[a-zA-Z]{2,3}$/,
            message: '不合法的邮箱格式!',
          },
        ]}
      />
    </ProForm>
  );
};

export default VisitorForm;

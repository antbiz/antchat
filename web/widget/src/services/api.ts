// @ts-ignore
/* eslint-disable */
import { getApiSid } from '@/utils/authority';
import request from 'umi-request';

/** 登录接口 POST /api/sys/account/signin */
export async function signin(body: API.SigninRequest, options?: { [key: string]: any }) {
  return request<API.SigninReply>('/api/v1/visitor/signin', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
};

/** 访客发送消息 GET /api/v1/visitor/send */
export async function sendMsg(body: API.SendMsgRequest, options?: { [key: string]: any }) {
  return request('/api/v1/visitor/send', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'x-api-sid': getApiSid(),
    },
    data: body,
    ...(options || {}),
  });
};

/** 获取聊天记录 GET /api/v1/visitor/history */
export async function getChatHistory(options?: { [key: string]: any }) {
  return request<{data: API.Message[]}>(`/api/v1/visitor/history`, {
    method: 'GET',
    headers: {
      'x-api-sid': getApiSid(),
    },
    ...(options || {}),
  });
};

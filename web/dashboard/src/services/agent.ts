// @ts-ignore
/* eslint-disable */
import { request } from 'umi';


/** 客服发送消息 GET /api/v1/agent/send */
export async function sendMsg(body: API.SendMsgRequest, options?: { [key: string]: any }) {
  return request('/api/v1/agent/send', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 获取指定访客的聊天记录 GET /api/v1/agent/history?visitorID= */
export async function getChatHistory(visitorID: string, options?: { [key: string]: any }) {
  return request<{data: API.Message[]}>(`/api/v1/agent/history?visitorID=${visitorID}`, {
    method: 'GET',
    ...(options || {}),
  });
}

/** 获取当前对话列表 GET /api/v1/agent/conversations */
export async function getConversations(options?: { [key: string]: any }) {
  return request<API.Conversation[]>('/api/v1/agent/conversations', {
    method: 'GET',
    ...(options || {}),
  });
}

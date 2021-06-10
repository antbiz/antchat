// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

/** 登录接口 POST /api/v1/account/signin */
export async function signin(body: API.SigninRequest, options?: { [key: string]: any }) {
  return request<API.SigninReply>('/api/v1/account/signin', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 退出登录接口 POST /api/v1/account/logout */
export async function logout(options?: { [key: string]: any }) {
  return request<Record<string, any>>('/api/v1/account/logout', {
    method: 'POST',
    ...(options || {}),
  });
}

/** 获取当前的用户 GET /api/v1/account/info */
export async function currentUser(options?: { [key: string]: any }) {
  return request<API.CurrentUser>('/api/v1/account/info', {
    method: 'GET',
    ...(options || {}),
  });
}

// TODO: 后端未实现
/** 获取通知 GET /api/notices */
export async function getNotices(options?: { [key: string]: any }) {
  return request<API.NoticeIconList>('/api/v1/account/notices', {
    method: 'GET',
    ...(options || {}),
  });
}

// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

/** 登录接口 POST /api/account/signin */
export async function signin(body: API.SigninRequest, options?: { [key: string]: any }) {
  return request<API.SigninReply>('/api/account/signin', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 退出登录接口 POST /api/account/logout */
export async function logout(options?: { [key: string]: any }) {
  return request<Record<string, any>>('/api/account/logout', {
    method: 'POST',
    ...(options || {}),
  });
}

/** 获取当前的用户 GET /api/currentUser */
export async function currentUser(options?: { [key: string]: any }) {
  return request<API.CurrentUser>('/api/user/info', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 此处后端没有提供注释 GET /api/notices */
export async function getNotices(options?: { [key: string]: any }) {
  return request<API.NoticeIconList>('/api/user/notices', {
    method: 'GET',
    ...(options || {}),
  });
}

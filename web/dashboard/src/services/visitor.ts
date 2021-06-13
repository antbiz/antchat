// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

/** 获取访客详情 GET /api/sys/visitor/{visitorID} */
export async function getVisitor(visitorID: string) {
  return request<API.Visitor>(`/api/sys/visitor/${visitorID}`, {
    method: 'GET',
  });
};

/** 更新访客详情 GET /api/sys/visitor/{visitorID} */
export async function updateVisitor(visitorID: string, data: API.Visitor) {
  return request(`/api/sys/visitor/${visitorID}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    data: data,
  });
};

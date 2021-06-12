// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

/** 获取访客详情 GET /api/v1/visitor/{visitorID} */
export async function getVisitor(visitorID: string) {
  return request<API.Visitor>(`/api/v1/visitor/${visitorID}`, {
    method: 'GET',
  });
};

/** 更新访客详情 GET /api/v1/visitor/{visitorID} */
export async function updateVisitor(visitorID: string, data: API.Visitor) {
  return request(`/api/v1/visitor/${visitorID}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    data: data,
  });
};

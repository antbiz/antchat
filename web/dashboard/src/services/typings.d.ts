// @ts-ignore
/* eslint-disable */

declare namespace API {
  type CurrentUser = {
    _id?: string;
    username?: string;
    nickname?: string;
    phone?: string;
    email?: string;
    avatar?: string;
    language?: string;
    isAdmin?: boolean;
  };

  type SigninRequest = {
    account?: string;
    password?: string;
    autoLogin?: boolean;
  };

  type SigninReply = {
    sid: string;
    status?: string;
    type?: string;
    currentAuthority?: string;
  };

  type PageParams = {
    pageNum?: number;
    pageSize?: number;
  };

  type FakeCaptcha = {
    code?: number;
    status?: string;
  };

  type ErrorResponse = {
    /** 业务约定的错误码 */
    errorCode: string;
    /** 业务上的错误信息 */
    errorMessage?: string;
    /** 业务上的请求是否成功 */
    success?: boolean;
  };

  type NoticeIconList = {
    data?: NoticeIconItem[];
    /** 列表的内容总数 */
    total?: number;
    success?: boolean;
  };

  type NoticeIconItemType = 'notification' | 'message' | 'event';

  type NoticeIconItem = {
    id?: string;
    extra?: string;
    key?: string;
    read?: boolean;
    avatar?: string;
    title?: string;
    status?: string;
    datetime?: string;
    description?: string;
    type?: NoticeIconItemType;
  };
}

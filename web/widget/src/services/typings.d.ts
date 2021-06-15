// @ts-ignore
/* eslint-disable */

declare namespace API {
  type ErrorResponse = {
    /** 业务约定的错误码 */
    errorCode: string;
    /** 业务上的错误信息 */
    errorMessage?: string;
    /** 业务上的请求是否成功 */
    success?: boolean;
  };

  type PageParams = {
    pageNum?: number;
    pageSize?: number;
  };

  type FakeCaptcha = {
    code?: number;
    status?: string;
  };

  type SigninRequest = {
    captcha: string;
    domian: string;
    visitorID?: string;
  };

  type SigninReply = {
    id: string;
    sid: string;
  };

  type SendMsgRequest = {
    type?: string;
    content?: any;
    createdAt?: number;
  }

  type Message = {
    id?: string;
    createdAt?: string;
    agentID?: string;
    visitorID?: string;
    senderID?: string;
    senderNick?: string;
    type?: string;
    content?: any;
    status?: number;
  }

  type Visitor = {
    id?: string;
    nickname: string;
    email?: string;
    phone?: string;
  }
}

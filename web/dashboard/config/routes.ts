export default [
  {
    name: 'exception.not-find',
    path: '/404',
    layout: false,
    hideInMenu: true,
    component: './exception/404',
  },
  {
    name: 'account.signin',
    path: '/signin',
    layout: false,
    hideInMenu: true,
    component: './signin',
  },
  {
    path: '/',
    redirect: '/conversation',
  },
  {
    path: '/conversation',
    name: 'conversation',
    icon: 'message',
    component: './conversation',
  },
];

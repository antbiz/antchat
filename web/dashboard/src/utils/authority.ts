export function getApiSid(): string {
  return localStorage.getItem('x-api-sid') || '';
};

export function setApiSid(sid: string) {
  localStorage.setItem('x-api-sid', sid);
};

export interface IResponse<T> {
  data: T;
}
export interface ILoginResponseData {
  message: string;
}
export interface IFileInfo {
  name: string;
  hash: string;
  size: number;
  download_count: number;
  created_at: string;
  type: string;
}

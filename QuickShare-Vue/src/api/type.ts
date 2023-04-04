export interface IResponse<T> {
  data: T;
}
export interface ILoginResponseData {
  message: string;
}
export interface IFileInfo {
  name: string;
  hash: string;
  type: string;
}

export interface IFileInfoDetail {
  name: string;
  hash: string;
  type: string;
  size: number;
  create_time: string;
  download_count: number;
}

export interface IResponse<T> {
  data: T;
  message: string;
}

export interface ILoginResponse {
  message: string;
}

export interface IUploadResponse {
  hash: string;
}

export interface IDeleteResponse {
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

export interface IShareCodeResponse {
  shareCode: string;
}

export interface IGetHashByShareCodeResponse {
  hash: string;
}

export interface IGetFileListRespose {
  data: IFileInfo[];
}

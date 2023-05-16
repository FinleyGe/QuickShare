import { AxiosProgressEvent } from "axios";
import { BaseUrl } from "../const";
import { request } from "./request";
import { IUploadResponse, IFileInfo, IFileInfoDetail, IShareCodeResponse, IDeleteResponse, ILoginResponse, IGetFileListRespose, IGetHashByShareCodeResponse } from "./type";
import "./type";

export async function UploadFile(file: File,
  onUploadProgress: (e: AxiosProgressEvent) => void)
  : Promise<IUploadResponse> {
  const formData = new FormData();
  formData.append("file", file);
  let res = await request<IUploadResponse>({
    url: "/upload",
    method: "POST",
    data: formData,
    onUploadProgress: onUploadProgress,
  })

  return res.data;
}

export function GetFile(hash: string) {
  window.location.href = `${BaseUrl}get/${hash}`;
}

export async function GetFileList(): Promise<IFileInfo[]> {
  let res = await request<IGetFileListRespose>({
    url: "/all_info",
    method: "GET",
  });

  return res.data.data;
}

export async function
  LoginAPI(username: string, password: string):
  Promise<ILoginResponse> {
  let res = await request({
    url: "/login",
    method: "POST",
    data: { username, password },
  });

  return res;
}

export async function DeleteFile(hash: string):
  Promise<IDeleteResponse> {
  let res = await request({
    url: `/delete/${hash}`,
    method: "DELETE",
  });

  return res.data;
}

export async function GetFileDetail(hash: string):
  Promise<IFileInfoDetail> {
  let res = await request({
    url: `/info/${hash}`,
    method: "GET",
  });
  console.log(res.data.data);

  return res.data.data;
}

export async function
  GetShareCode(hash: string):
  Promise<IShareCodeResponse> {

  let res = await request({
    url: `/share/${hash}`,
    method: "GET",
  });

  return res.data;
}

export async function
  GetHashByShareCode(shareCode: string):
  Promise<IGetHashByShareCodeResponse> {
  let res = await request({
    url: `/getshare/${shareCode}`,
    method: "GET",
  });

  return res.data;
}

export async function
  GetFileListByType(type: string):
  Promise<IFileInfo[]> {
  let res = await request({
    url: `/all_info_type`,
    params: { type },
    method: "GET",
  });

  return res.data.data;
}

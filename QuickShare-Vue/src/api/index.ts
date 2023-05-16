import axios, { AxiosProgressEvent } from "axios";
import { BaseUrl } from "../const";
import { IResponse, IUploadResponseData, IFileInfo, IFileInfoDetail } from "./type";

export const api = axios.create({
  baseURL: BaseUrl,
  withCredentials: true,
});

export function UploadFile(file: File,
  onUploadProgress: (e: AxiosProgressEvent) => void)
  : Promise<IResponse<IUploadResponseData>> {
  const formData = new FormData();
  formData.append("file", file);
  return api.post("/upload", formData, {
    onUploadProgress: onUploadProgress,
  });
}

export function GetFile(hash: string) {
  window.location.href = `${BaseUrl}get/${hash}`;
}

export async function GetFileList(): Promise<IFileInfo[]> {
  const response = await api.get("/all_info");
  return response.data.data.data;
}

export async function LoginAPI(username: string, password: string): Promise<Boolean> {
  const response = await api.post("/login", { username, password });
  return response.data.message === "OK";
}

export async function DeleteFile(hash: string): Promise<Boolean> {
  const response = await api.delete(`/delete/${hash}`);
  return response.data.message === "OK";
}

export async function GetFileDetail(hash: string): Promise<IFileInfoDetail> {
  const response = await api.get(`/info/${hash}`);
  return response.data.data.data as IFileInfoDetail;
}

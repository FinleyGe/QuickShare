import axios from "axios";
import { BaseUrl } from "../const";
import { IResponse, ILoginResponseData, IFileInfo } from "./type";

export const api = axios.create({
  baseURL: BaseUrl,
  withCredentials: true,
});

export function UploadFile(file: File) {
  const formData = new FormData();
  formData.append("file", file);
  return api.post("/upload?temporary=true", formData);
}

export function GetFile(hash: string) {
  // return api.get(`/get/${hash}`, { responseType: "blob" });
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


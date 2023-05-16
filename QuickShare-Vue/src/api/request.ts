import axios, { AxiosRequestConfig } from "axios";
import { BaseUrl } from "../const";
import { IResponse } from "./type";

const api = axios.create({
  baseURL: BaseUrl,
  withCredentials: true,
});

export async function request<T = any>
  (config: AxiosRequestConfig):
  Promise<IResponse<T>> {
  const response = await api.request(config);
  return response.data as IResponse<T>;
}

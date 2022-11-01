import axios from "axios";
import {BaseUrl} from "../const";

export const api = axios.create({
  baseURL: BaseUrl,
});

export function UploadFile(file: File) {
  const formData = new FormData();
  formData.append("file", file);
  return api.post("/upload?temporary=true", formData);
}

export const ENV = "dev"; // dev | prod | test

export var BaseUrl: string = "http://localhost:8888/";
if (ENV === "dev") {
  BaseUrl = "http://localhost:8888/";
} else if (ENV === "prod" || ENV === "test") {
  BaseUrl = "https://qs.f1nley.xyz/api";
}

import { defineStore, createPinia } from 'pinia';

export type DeviceSize = "small" | "medium" | "large";

interface State {
  isLogin: Boolean;
  deviceSize: DeviceSize;
}

export const useStore = defineStore({
  id: 'app',
  state: (): State => ({
    isLogin: false,
    deviceSize: "small",
  }),
  actions: {
    setLogin(isLogin: Boolean) {
      this.isLogin = isLogin;
    },
    setDeviceSize(deviceSize: "small" | "medium" | "large") {
      this.deviceSize = deviceSize;
    },
  },
});

export const pinia = createPinia();

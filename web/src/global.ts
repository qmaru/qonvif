declare global {
  interface Window {
    api: string,
    dateVer: string
    commitVer: string
  }
}

// API 地址
window.api = "http://127.0.0.1:8373"
// 版本号
window.dateVer = "20000829"
window.commitVer = "minami"

export { }

/**********************************
 * @FilePath: index.js
 * @Author: Ronnie Zhang
 * @LastEditor: Ronnie Zhang
 * @LastEditTime: 2023/12/04 22:46:28
 * @Email: zclzone@outlook.com
 *  | https://isme.top
 **********************************/

import axios from 'axios'
import { setupInterceptors } from './interceptors'

export function createAxios(options = {}) {
  // 初始化请求这里使用了配置文件的这个参数 VITE_AXIOS_BASE_URL
  const defaultOptions = {
    // baseURL: import.meta.env.VITE_AXIOS_BASE_URL,
    baseURL: '/api',
    timeout: 12000,
  }
  const service = axios.create({
    ...defaultOptions,
    ...options,
  })
  setupInterceptors(service)
  return service
}

export const request = createAxios()

export const mockRequest = createAxios({
  baseURL: '/mock-api',
})

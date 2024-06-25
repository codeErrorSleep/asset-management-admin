/**********************************
 * @Author: Ronnie Zhang
 * @LastEditor: Ronnie Zhang
 * @LastEditTime: 2023/12/05 21:29:51
 * @Email: zclzone@outlook.com
 *  | https://isme.top
 **********************************/

import { request } from '@/utils'

export default {
  create: data => request.post('/api/tower-details', data),
  read: (params = {}) => request.get('/api/tower-details', { params }),
  update: data => request.patch(`/api/tower-details/${data.id}`, data),
  delete: id => request.delete(`/api/tower-details/${id}`),
  resetPwd: (id, data) => request.patch(`/user/password/reset/${id}`, data),

  getAllRoles: () => request.get('/role?enable=1'),

  // 绑定设备
  // readDeviceList: (params = {}) => request.get(`/api/tower-equipment/list/`, { params }),
  readDeviceList: id => request.get(`/api/tower-equipment/list/${id}`),
}

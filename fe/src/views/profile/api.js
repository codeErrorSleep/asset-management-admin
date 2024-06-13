/**********************************
 * @Author: Ronnie Zhang
 * @LastEditor: Ronnie Zhang
 * @LastEditTime: 2023/12/05 21:30:03
 * @Email: zclzone@outlook.com
 *  | https://isme.top
 **********************************/

import { request } from '@/utils'

export default {
  changePassword: data => request.post('/auth/password', data),
  updateProfile: data => request.patch(`/user/profile/${data.id}`, data),
}
